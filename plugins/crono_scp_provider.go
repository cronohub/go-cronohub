package main

import (
	"log"

	"github.com/cronohub/cronohub/sdk"
	plugin "github.com/hashicorp/go-plugin"
)

// Archive is a concrete implementation of the archive plugin.
type Archive struct{}

// Execute is the entry point to this plugin.
func (Archive) Execute(filename string) (bool, error) {
	// ioutil.WriteFile("test.txt", []byte(filename), 0766)
	log.Println("I'm alive!")
	return true, nil
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
