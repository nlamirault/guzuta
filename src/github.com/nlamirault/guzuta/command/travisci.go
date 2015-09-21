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
	"github.com/nlamirault/guzuta/providers/travis"
	"github.com/nlamirault/guzuta/utils"
)

// TravisCICommand represents CLI command for TravisCI provider
type TravisCICommand struct {
	UI cli.Ui
}

func (c *TravisCICommand) Help() string {
	helpText := `
Usage: guzuta travisci [options]
	Check projects status from TravisCI
Options:
	--debug                       Debug mode enabled
	--travis-name=name            Project name
	--travis-namespace=namespace  Namespace
`
	return strings.TrimSpace(helpText)
}

func (c *TravisCICommand) Synopsis() string {
	return "Check projects status from TravisCI"
}

func (c *TravisCICommand) Run(args []string) int {
	var debug bool
	var name, namespace string
	f := flag.NewFlagSet("travis", flag.ContinueOnError)
	f.Usage = func() { c.UI.Output(c.Help()) }
	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&name, "travis-name", "", "TravisCI project's name")
	f.StringVar(&namespace, "travis-namespace", "", "TravisCI namespace")

	if err := f.Parse(args); err != nil {
		return 1
	}
	if debug {
		c.UI.Info("Debug mode enabled.")
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	client := travis.NewClient(utils.Getenv("GUZUTA_TRAVIS_GITHUB_TOKEN"))
	if len(name) > 0 {
		travisRepositoryStatus(client, name)
		return 0
	}
	if len(namespace) > 0 {
		travisRepositoriesStatus(client, namespace)
		return 0
	}
	return 0
}

func travisRepositoryStatus(client *travis.Client, name string) {
	resp, err := client.GetRepository(name)
	if err != nil {
		colorstring.Printf("[red] Travis : %s", err.Error())
		return
	}
	status := "[green] OK"
	if resp.Repository.LastBuildState == "failed" {
		status = "[red] KO"
	}
	fmt.Printf(colorstring.Color(status) + "\t" + resp.Repository.Slug + "\n")
}

func travisRepositoriesStatus(client *travis.Client, namespace string) {
	resp, err := client.GetRepositories(namespace)
	if err != nil {
		colorstring.Printf("[red] Travis : %s", err.Error())
		return
	}
	for _, repo := range resp.Repositories {
		status := "[green] OK"
		if repo.LastBuildState == "failed" {
			status = "[red] KO"
		}
		fmt.Printf(colorstring.Color(status) + "\t" + repo.Slug + "\n")
	}
}
