package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v3-unstable"
)

var (
	verbose   = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	parallel  = kingpin.Flag("parallel", "Number of parallel threads.").Default("5").Short('p').Int()
	aParallel = kingpin.Flag("aparallel", "Number of parallel threads to when archiving.").Default("5").Short('a').Int()
	archiver  = kingpin.Arg("archiver", "Method to use for archiving.").Default("s3").String()
)

var token string

func main() {
	token = os.Getenv("CRONO_GITHUB_TOKEN")
	if len(token) < 1 {
		log.Fatal("Please set CRONO_GITHUB_TOKEN to a valid token in order to authenticate to github.")
	}
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

	LogIfVerbose("Loading all plugins...\n")
	loadPlugins()
	LogIfVerbose("Archiving with %s using %d parallel threads.\n", *archiver, *parallel)
	// repos := getRepositoryList()
	// list := download(*parallel, repos)
	list := []string{"test", "test2"}
	LogIfVerbose("Downloaded %d repositories.\n", len(list))
	archive(*archiver, *aParallel, list)
	LogIfVerbose("Finished archiving. Good bye.")
}
