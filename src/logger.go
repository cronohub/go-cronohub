package main

import "log"

// LogIfVerbose is a simple utility to log a message if verbose
// logging is enabled.
func LogIfVerbose(v bool, msg string, params ...interface{}) {
	if v {
		log.Printf(msg, params...)
	}
}
