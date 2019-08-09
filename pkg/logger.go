package pkg

import (
	"github.com/cronohub/go-cronohub/config"
	"log"
)

// LogIfVerbose is a simple utility to log a message if verbose
// logging is enabled.
func LogIfVerbose(msg string, params ...interface{}) {
	if config.Config.Verbose {
		log.Printf(msg, params...)
	}
}
