package main

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
)

// GitClient is a client with an authenticated http client a context
// and a github.Client.
type GitClient struct {
	ctx     context.Context
	client  *http.Client
	gClient *github.Client
}

var gClient *GitClient

// NewClient creates a new Archiver client. This will return an existing client
// if it already was created.
func NewClient() *GitClient {
	LogIfVerbose("Creating github client\n")
	if gClient != nil {
		return gClient
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: go_cronohub.token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	gClient = &GitClient{
		ctx:     ctx,
		client:  tc,
		gClient: client,
	}
	return gClient
}

func getRepositoryList() []*github.Repository {
	client := NewClient()
	user, _, _ := client.gClient.Users.Get(client.ctx, "")
	LogIfVerbose("Client created...\n")
	// list all repositories for the authenticated user
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	var allRepos []*github.Repository
	LogIfVerbose("Gathering repositories...\n")
	for {
		repos, resp, err := client.gClient.Repositories.List(client.ctx, user.GetLogin(), opt)
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
