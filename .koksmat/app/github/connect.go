package github

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/v68/github"
	"github.com/jferrl/go-githubauth"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func GetAppClient() (*github.Client, error) {
	privateKey := []byte(viper.GetString("GITHUB_PRIVATEKEY"))
	applicationId := viper.GetInt64("GITHUB_APPID")
	appTokenSource, err := githubauth.NewApplicationTokenSource(applicationId, privateKey)
	if err != nil {
		fmt.Println("Error creating application token source:", err)
		return nil, err
	}

	installationTokenSource := githubauth.NewInstallationTokenSource(55863281, appTokenSource)

	// oauth2.NewClient uses oauth2.ReuseTokenSource to reuse the token until it expires.
	// The token will be automatically refreshed when it expires.
	// InstallationTokenSource has the mechanism to refresh the token when it expires.
	httpClient := oauth2.NewClient(context.Background(), installationTokenSource)

	client := github.NewClient(httpClient)
	return client, nil
}

/*
GetClient returns a new GitHub client.

*/

func GetClient() (*github.Client, error) {
	if !viper.IsSet("GITHUB_PAT") {
		return nil, errors.New("GITHUB_PAT not set")
	}

	pat := viper.GetString("GITHUB_PAT")
	client := github.NewClient(nil).WithAuthToken(pat)
	return client, nil
}
