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

	"github.com/nlamirault/guzuta/providers/travisci"
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
	--debug                   Debug mode enabled
	--name=name               Project name
	--namespace=namespace     Namespace
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
	f.StringVar(&name, "name", "", "TravisCI project's name")
	f.StringVar(&namespace, "namespace", "", "TravisCI namespace")

	if err := f.Parse(args); err != nil {
		return 1
	}
	setupLogging(debug)
	client := travisci.NewClient(utils.Getenv("GUZUTA_TRAVIS_GITHUB_TOKEN"))
	if len(name) > 0 && len(namespace) > 0 {
		c.travisRepositoryStatus(client, namespace, name)
		return 0
	}
	if len(namespace) > 0 {
		c.travisRepositoriesStatus(client, namespace)
		return 0
	}
	return 0
}

func (c *TravisCICommand) travisRepositoryStatus(client *travisci.Client, namespace string, name string) {
	resp, err := client.GetRepository(fmt.Sprintf("%s/%s", namespace, name))
	if err != nil {
		c.UI.Error(colorstring.Color("[red] Travis : " + err.Error()))
		return
	}
	c.travisPrintRepository(&resp.Repository)
}

func (c *TravisCICommand) travisRepositoriesStatus(client *travisci.Client, namespace string) {
	resp, err := client.GetRepositories(namespace)
	if err != nil {
		c.UI.Error(colorstring.Color("[red] Travis : " + err.Error()))
		return
	}
	for _, repo := range resp.Repositories {
		c.travisPrintRepository(&repo)
	}
}

func (c *TravisCICommand) travisPrintRepository(repo *travisci.Repository) {
	status := ""
	if repo.LastBuildState == "passed" {
		status = "[green]OK"
	} else if repo.LastBuildState == "failed" {
		status = "[red]KO"
	}
	c.UI.Info(colorstring.Color(status + "\t" + repo.Slug))
}
