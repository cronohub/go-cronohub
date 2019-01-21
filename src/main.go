package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v3-unstable"
)

var (
	verbose  = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	parallel = kingpin.Flag("parallel", "Number of parallel threads.").Short('p').Int()
	archiver = kingpin.Arg("archiver", "Method to use for archiving.").String()
)

func main() {
	kingpin.Parse()
	fmt.Println(`
	_______  ______    _______  __    _  _______  __   __  __   __  _______
	|       ||    _ |  |       ||  |  | ||       ||  | |  ||  | |  ||  _    |
	|       ||   | ||  |   _   ||   |_| ||   _   ||  |_|  ||  | |  || |_|   |
	|       ||   |_||_ |  | |  ||       ||  | |  ||       ||  |_|  ||       |
	|      _||    __  ||  |_|  ||  _    ||  |_|  ||       ||       ||  _   |
	|     |_ |   |  | ||       || | |   ||       ||   _   ||       || |_|   |
	|_______||___|  |_||_______||_|  |__||_______||__| |__||_______||_______|
	`)
	LogIfVerbose(*verbose, "Archiving with %s\n", *archiver)
}
