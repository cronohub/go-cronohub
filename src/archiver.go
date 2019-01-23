package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/google/go-github/v21/github"
)

// download will download archives of the users repositories.
// only owned repositories are archived, organization repositories are left out.
func download(p int, repos []*github.Repository) []string {
	LogIfVerbose("Downloading %d number of repos with %d parallel threads.\n", len(repos), p)
	sema := make(chan struct{}, p)
	archiveList := make([]string, 0)
	var wg sync.WaitGroup
	client := NewClient()
	for _, r := range repos {
		wg.Add(1)
		go func(repo *github.Repository) {
			defer wg.Done()
			sema <- struct{}{}
			defer func() { <-sema }()
			LogIfVerbose("Downloading archive for repository: %s\n", *repo.URL)
			downloadURL := "https://github.com/" + repo.GetOwner().GetLogin() + "/" + repo.GetName() + "/archive/master.zip"
			archiveName := repo.GetName() + ".zip"
			out, err := os.Create(archiveName)
			if err != nil {
				log.Println("failed downloading repo: ", repo.GetName())
				return
			}
			defer out.Close()
			LogIfVerbose("Started downloading archive for: %s\n", downloadURL)
			resp, err := client.client.Get(downloadURL)
			if err != nil {
				log.Println("failed to get zip archive for repo: ", repo.GetName())
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				log.Println("status was not 200: ", resp.Status)
			}
			io.Copy(out, resp.Body)
			archiveList = append(archiveList, archiveName)
		}(r)
	}
	wg.Wait()
	return archiveList
}

func archive(plugin string, parallel int, list []string) {
	LogIfVerbose("Running with plugin %s parallel %d threads on %d items.\n", plugin, parallel, len(list))
	_, err := runPlugin("crono_"+plugin+"_provider", list[0])
	if err != nil {
		log.Fatal("error while running archiver: ", err)
	}
}
