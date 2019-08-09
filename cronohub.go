package main

import (
	"fmt"
	main2 "github.com/cronohub/go-cronohub/pkg"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v3-unstable"
)

var (
	config.Config.Verbose   = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	config.Config.Fetcher = kingpin.Arg("fetcher", "Plugin to use for fetching.").Default("git").String()
	config.Config.Archiver  = kingpin.Arg("archiver", "Plugin to use for archiving.").Default("s3").String()
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

	pkg.LogIfVerbose("Loading all plugins...\n")
	main2.loadPlugins()
	main2.LogIfVerbose("Archiving with %s using %d parallel threads.\n", *archiver, *parallel)
	// repos := getRepositoryList()
	// list := download(*parallel, repos)
	list := []string{"test", "test2"}
	main2.LogIfVerbose("Downloaded %d repositories.\n", len(list))
	main2.archive(*archiver, *aParallel, list)
	main2.LogIfVerbose("Finished archiving. Good bye.")
}
