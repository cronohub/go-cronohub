package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cronohub/sdk"
	"google.golang.org/grpc/grpclog"

	plugin "github.com/hashicorp/go-plugin"
)

var pluginMap = make(map[string]plugin.Plugin, 0)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion: 1,
	MagicCookieKey:  "CRONOHUB_PLUGINS",
	// Never ever change this.
	MagicCookieValue: "9f4c000c-dd07-4968-a33a-a42337c5f479",
}

func init() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)
}

func loadPlugins() {
	ps, _ := discoverPlugins("crono_*_provider")
	for _, v := range ps {
		pluginName := filepath.Base(v)
		pluginMap[pluginName] = &sdk.ArchiveGRPCPlugin{}
	}
}

func runPlugin(name string, filename string) (bool, error) {
	raw := getRawForPlugin(name)

	p := raw.(sdk.Archive)
	fmt.Println("FILENAME: ", filename)
	fmt.Println("NAME: ", name)
	fmt.Println(pluginMap)
	ret := p.Execute(filename)
	if !ret {
		LogIfVerbose("A plugin with name '%s' prevented archive to run.\n", name)
		err := errors.New("plugin prevented archive to run")
		return false, err
	}
	return true, nil
}

func discoverPlugins(postfix string) (p []string, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return p, err
	}
	plugs, err := plugin.Discover(postfix, dir)
	if err != nil {
		return nil, err
	}
	return plugs, nil
}

func getRawForPlugin(v string) interface{} {
	var cmd *exec.Cmd
	dir, _ := os.Getwd()
	ext := filepath.Ext(v)
	switch ext {
	case ".py":
		python := getExecutionBinary("python")
		cmd = exec.Command(python, filepath.Join(dir, v))
	case ".rb":
		ruby := getExecutionBinary("ruby")
		cmd = exec.Command(ruby, filepath.Join(dir, v))
	default:
		cmd = exec.Command(filepath.Join(dir, v))
	}
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: Handshake,
		Plugins:         pluginMap,
		Cmd:             cmd,
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC},
	})

	defer client.Kill()
	grpcClient, err := client.Client()
	if err != nil {
		log.Println("Error creating client:", err.Error())
		os.Exit(1)
	}

	pluginName := filepath.Base(v)
	// Request the plugin
	raw, err := grpcClient.Dispense(pluginName)
	if err != nil {
		log.Println("Error requesting plugin:", err.Error())
		os.Exit(1)
	}
	return raw
}

func getExecutionBinary(want string) string {
	binary, err := exec.LookPath(want)
	if err != nil {
		log.Printf("Could not locate binary for %s on PATH.\n", want)
		os.Exit(1)
	}
	return binary
}
