package main

import (
	"io/ioutil"

	"github.com/cronohub/sdk"
	plugin "github.com/hashicorp/go-plugin"
)

type Archive struct{}

// Execute is the entry point to this plugin.
func (Archive) Execute(filename string) bool {
	ioutil.WriteFile("test.txt", []byte(filename), 0766)
	return true
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: sdk.Handshake,
		Plugins: map[string]plugin.Plugin{
			"crono_scp_provider": &sdk.ArchiveGRPCPlugin{Impl: &Archive{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
