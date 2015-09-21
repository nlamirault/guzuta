// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/colorstring"

	"github.com/nlamirault/guzuta/logging"
	"github.com/nlamirault/guzuta/providers/github"
	"github.com/nlamirault/guzuta/utils"
)

// GithubCICommand represents CLI command for Github provider
type GithubCommand struct {
	UI cli.Ui
}

func (c *GithubCommand) Help() string {
	helpText := `
Usage: guzuta github [options]
	Check projects status from Github
Options:
	--debug                   Debug mode enabled
	--name=name               Project name
	--username=username       Github username`
	return strings.TrimSpace(helpText)
}

func (c *GithubCommand) Synopsis() string {
	return "Check projects status from Github"
}

func (c *GithubCommand) Run(args []string) int {
	var debug bool
	var name, username, token string
	f := flag.NewFlagSet("github", flag.ContinueOnError)
	f.Usage = func() { c.UI.Output(c.Help()) }
	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&name, "name", "", "GithubCI project's name")
	f.StringVar(&username, "username", "", "Github username")
	f.StringVar(&token, "token", utils.Getenv("GUZUTA_GITHUB_TOKEN"), "API token")

	if err := f.Parse(args); err != nil {
		return 1
	}
	if debug {
		c.UI.Info("Debug mode enabled.")
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	client := github.NewClient(utils.Getenv("GUZUTA_GITHUB_TOKEN"))
	if len(name) > 0 && len(username) > 0 {
		githubRepositoryStatus(client, username, name)
		return 0
	}
	if len(username) > 0 {
		githubRepositoriesStatus(client, username)
		return 0
	}
	return 0
}

func githubRepositoryStatus(client *github.Client, username string, name string) {
	repo, err := client.GetRepository(username, name)
	if err != nil {
		colorstring.Printf("[red] Github : %s\n", err.Error())
		return
	}
	githubPrintRepository(repo)
}

func githubRepositoriesStatus(client *github.Client, username string) {
	resp, err := client.GetRepositories(username)
	if err != nil {
		colorstring.Printf("[red] Github : %s\n", err.Error())
		return
	}
	for _, repo := range *resp {
		githubPrintRepository(&repo)
	}
}

func githubPrintRepository(repo *github.Repository) {
	fmt.Printf("* %s - %s\n", repo.Name, repo.Description)
}
