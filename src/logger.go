package main

import "log"

// LogIfVerbose is a simple utility to log a message if verbose
// logging is enabled.
func LogIfVerbose(msg string, params ...interface{}) {
	if *verbose {
		log.Printf(msg, params...)
	}
}
