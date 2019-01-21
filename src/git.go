package main

import (
	"context"
	"log"

	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
)

func getRepositoryList() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(ctx, "", opt)
		if err != nil {
			log.Fatal("error retrieving repositories: ", err)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	LogIfVerbose("Retrieved %d repositories.", len(allRepos))
}
