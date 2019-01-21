package main

import (
	"context"
	"log"

	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
)

var client *github.Client

func getRepositoryList() []*github.Repository {
	LogIfVerbose("Creating github client\n")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)
	user, _, _ := client.Users.Get(ctx, "")
	LogIfVerbose("Client created...\n")
	// list all repositories for the authenticated user
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	var allRepos []*github.Repository
	LogIfVerbose("Gathering repositories...\n")
	for {
		repos, resp, err := client.Repositories.List(ctx, user.GetLogin(), opt)
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
	return allRepos
}
