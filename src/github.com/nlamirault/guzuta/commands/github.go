// Copyright (C) 2015  Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func getGithubClient(c *cli.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.GlobalString("github-token")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	log.Debugf("Client Github: %v", client)
	return client
}

var commandGithubListRepositories = cli.Command{
	Name:        "list",
	Usage:       "List user repositories",
	Description: ``,
	Action:      doListRepositories,
	Flags:       []cli.Flag{},
}

func doListRepositories(c *cli.Context) {
	log.Debugf("List Github repositories")
	client := getGithubClient(c)
	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List("", nil)
	if err != nil {
		log.Errorf("Retrieving repositories %v", err)
		return
	}
	log.Info("Github repositories")
	log.Infof("----------------------------------------------")
	for _, repo := range repos {
		name := *repo.Name
		log.Infof("%s", name)
	}

}
