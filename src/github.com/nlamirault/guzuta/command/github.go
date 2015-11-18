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

	"github.com/nlamirault/guzuta/providers/github"
	"github.com/nlamirault/guzuta/utils"
)

// GithubCICommand represents CLI command for Github provider
type GithubCommand struct {
	UI cli.Ui
}

func (c *GithubCommand) Help() string {
	helpText := `
Usage: guzuta github [options] actions

   Manage projects from Github

Actions:
       repo     : Describe a repository
       repos    : List all repositories
       issue    : Describe an issue
       issues   : List all issues

Options:
	--debug                   Debug mode enabled
	--name=name               Project name
	--username=username       Github username
`
	return strings.TrimSpace(helpText)
}

func (c *GithubCommand) Synopsis() string {
	return "Manage projects from Github"
}

func (c *GithubCommand) Run(args []string) int {
	var debug bool
	var name, username, token string
	var issueID int
	f := flag.NewFlagSet("github", flag.ContinueOnError)
	f.Usage = func() { c.UI.Output(c.Help()) }
	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&name, "name", "", "Github project's name")
	f.StringVar(&username, "username", "", "Github username")
	f.IntVar(&issueID, "issueid", 0, "Issue number")
	f.StringVar(&token, "token", utils.Getenv("GUZUTA_GITHUB_TOKEN"), "API token")

	if err := f.Parse(args); err != nil {
		return 1
	}
	action := f.Args()
	// fmt.Printf("Args : %v\n", action)
	if len(action) != 1 {
		errorMessage(
			c.UI,
			"At least one action to github must be specified.",
			c.Help())
		return 1
	}
	setupLogging(debug)
	if action[0] == "repo" {
		if len(name) > 0 && len(username) > 0 {
			githubRepositoryStatus(getGithubClient(), username, name)
			return 0
		}
		errorMessage(c.UI, "Please specify name and username.", c.Help())
		return 1
	} else if action[0] == "repos" {
		if len(username) > 0 {
			githubRepositoriesStatus(getGithubClient(), username)
			return 0
		}
		errorMessage(c.UI, "Please specify username.", c.Help())
		return 1
	} else if action[0] == "issue" {
		if len(name) > 0 && len(username) > 0 {
			githubGetRepositoryIssue(
				getGithubClient(), username, name, issueID)
			return 0
		}
		errorMessage(
			c.UI, "Please specify username, name and issueID.", c.Help())
		return 1
	} else if action[0] == "issues" {
		opt := &github.IssueListOptions{
			Filter: "all",
		}
		if len(name) > 0 && len(username) > 0 {
			githubListRepositoryIssues(
				getGithubClient(), username, name, opt)
		} else {
			githubListIssues(getGithubClient(), opt)
		}
		return 0
	}
	return 0
}

func getGithubClient() *github.Client {
	return github.NewClient(utils.Getenv("GUZUTA_GITHUB_TOKEN"))
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

func githubListIssues(client *github.Client, opt *github.IssueListOptions) {
	issues, err := client.ListIssues(opt)
	if err != nil {
		colorstring.Printf("[red] Github : %s\n", err.Error())
		return
	}
	for _, issue := range *issues {
		githubPrintIssue(&issue)
	}
}

func githubListRepositoryIssues(client *github.Client, username string, name string, opt *github.IssueListOptions) {
	issues, err := client.ListRepositoryIssues(username, name, opt)
	if err != nil {
		colorstring.Printf("[red] Github : %s\n", err.Error())
		return
	}
	for _, issue := range *issues {
		githubPrintIssue(&issue)
	}
}

func githubGetRepositoryIssue(client *github.Client, username string, name string, issueID int) {
	issue, err := client.GetRepositoryIssue(username, name, issueID)
	if err != nil {
		colorstring.Printf("[red] Github : %s\n", err.Error())
		return
	}
	fmt.Printf("Number:     %d\nState:      %s\nTitle:      %s\n",
		issue.Number, issue.State, issue.Title)
	fmt.Printf("Creation:   %s\nUpdated:    %s\n",
		issue.CreatedAt, issue.UpdatedAt)
}

func githubPrintIssue(issue *github.Issue) {
	fmt.Printf("- [%d] %s - %s\n", issue.Number, issue.State, issue.Title)
}
