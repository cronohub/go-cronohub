package main

import (
	"log"

	"github.com/cronohub/sdk"
	plugin "github.com/hashicorp/go-plugin"
)

type Archive struct{}

// Execute is the entry point to this plugin.
func (Archive) Execute(filename string) bool {
	log.Println("got filename: ", filename)
	return true
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			"crono_scp_provider": &sdk.ArchiveGRPCPlugin{Impl: &Archive{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
