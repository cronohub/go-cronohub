package main

import (
	"fmt"
	"log"

	"github.com/cronohub/go-cronohub/config"
	"gopkg.in/alecthomas/kingpin.v3-unstable"
)

func init() {
	config.Config.Verbose = *kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	config.Config.Fetcher = *kingpin.Arg("fetcher", "Plugin to use for fetching.").Default("git").String()
	config.Config.Archiver = *kingpin.Arg("archiver", "Plugin to use for archiving.").Default("s3").String()
	kingpin.Parse()
}

func main() {
	fmt.Println(`
	_______  ______    _______  __    _  _______  __   __  __   __  _______
	|       ||    _ |  |       ||  |  | ||       ||  | |  ||  | |  ||  _    |
	|       ||   | ||  |   _   ||   |_| ||   _   ||  |_|  ||  | |  || |_|   |
	|       ||   |_||_ |  | |  ||       ||  | |  ||       ||  |_|  ||       |
	|      _||    __  ||  |_|  ||  _    ||  |_|  ||       ||       ||  _   |
	|     |_ |   |  | ||       || | |   ||       ||   _   ||       || |_|   |
	|_______||___|  |_||_______||_|  |__||_______||__| |__||_______||_______|
	`)

	log.Println("Loading all plugins...")
	//main2.loadPlugins()
	// repos := getRepositoryList()
	// list := download(*parallel, repos)
	//main2.archive(*archiver, *aParallel, list)
}
