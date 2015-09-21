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
	"github.com/nlamirault/guzuta/providers/circleci"
	"github.com/nlamirault/guzuta/utils"
)

type CircleCICommand struct {
	UI cli.Ui
}

func (c *CircleCICommand) Help() string {
	helpText := `
Usage: guzuta circleci [options]
	Check projects status from CircleCI
Options:
	--debug                       Debug mode enabled
	--name=name                   Project name
	--username=username           Username
        --token=xxxx                  API Token
`
	return strings.TrimSpace(helpText)
}

func (c *CircleCICommand) Synopsis() string {
	return "Check projects status from CircleCI"
}

func (c *CircleCICommand) Run(args []string) int {
	var debug bool
	var name, username, token string
	f := flag.NewFlagSet("circle", flag.ContinueOnError)
	f.Usage = func() { c.UI.Output(c.Help()) }
	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&name, "name", "", "CircleCI project's name")
	f.StringVar(&username, "username", "", "CircleCI username")
	f.StringVar(&token, "token", utils.Getenv("GUZUTA_CIRCLECI_TOKEN"), "API Token")

	if err := f.Parse(args); err != nil {
		return 1
	}
	if debug {
		c.UI.Info("Debug mode enabled.")
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	// token := utils.Getenv("GUZUTA_CIRCLECI_TOKEN")
	if len(token) <= 0 {
		c.UI.Error("CircleCI token invalid. Set CLI argument or GUZUTA_CIRCLECI_TOKEN environment variable.")
		return 1
	}
	client := circleci.NewClient(token)
	if len(name) > 0 && len(username) > 0 {
		circleciProjectStatus(client, username, name)
		return 0
	}
	if len(username) > 0 {
		circleciProjectsStatus(client)
		return 0
	}
	return 0
}

func circleciProjectStatus(client *circleci.Client, username string, project string) {
	resp, err := client.GetProject(&circleci.ProjectInput{
		Username: username,
		Project:  project,
		Limit:    1,
	})
	if err != nil {
		colorstring.Printf("[red] CircleCI : %s\n", err.Error())
		return
	}
	for _, p := range *resp {
		status := "[green] OK"
		if p.Outcome == "failed" {
			status = "[red] KO"
		}
		fmt.Printf(colorstring.Color(status) + "\t" +
			fmt.Sprintf("%s/%s", username, project) + "\n")
	}
}

func circleciProjectsStatus(client *circleci.Client) {
	resp, err := client.GetProjects()
	if err != nil {
		colorstring.Printf("[red] CircleCI : %s\n", err.Error())
		return
	}
	for _, p := range *resp {
		status := "[green] OK"
		if p.Branches.Master.RecentBuilds[0].Outcome == "failed" {
			status = "[red] KO"
		}
		fmt.Printf(colorstring.Color(status) + "\t" +
			fmt.Sprintf("%s/%s", p.Username, p.Reponame) + "\n")
	}
}
