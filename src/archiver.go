package main

// archive will get the necessary plugin and pass on
// a file name to archive. This method can be called
// concurrently.
func archive(p string) {
	LogIfVerbose("Archiving using plugin %s\n", p)
}
