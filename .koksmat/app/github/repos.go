package github

import (
	"context"

	"github.com/google/go-github/v68/github"
)

func GetRepos(client *github.Client, orgName string) ([]*github.Repository, error) {

	repos, _, err := client.Repositories.ListByOrg(context.Background(), orgName, nil)
	if err != nil {
		return nil, err
	}
	println(repos)

	return repos, err
}
