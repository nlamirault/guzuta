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
	"github.com/nlamirault/guzuta/providers/gitlab"
	"github.com/nlamirault/guzuta/utils"
)

// GitlabCICommand represents CLI command for Gitlab provider
type GitlabCommand struct {
	UI cli.Ui
}

func (c *GitlabCommand) Help() string {
	helpText := `
Usage: guzuta gitlab [options]
	Check projects status from Gitlab
Options:
	--debug                   Debug mode enabled
	--name=name               Project name
	--namespace=namespace     Gitlab namespace`
	return strings.TrimSpace(helpText)
}

func (c *GitlabCommand) Synopsis() string {
	return "Check projects status from Gitlab"
}

func (c *GitlabCommand) Run(args []string) int {
	var debug bool
	var name, namespace, token string
	f := flag.NewFlagSet("github", flag.ContinueOnError)
	f.Usage = func() { c.UI.Output(c.Help()) }
	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&name, "name", "", "Gitlab project's name")
	f.StringVar(&namespace, "namespace", "", "Gitlab namespace")
	f.StringVar(&token, "token", utils.Getenv("GUZUTA_GITLAB_TOKEN"), "API token")

	if err := f.Parse(args); err != nil {
		return 1
	}
	if debug {
		c.UI.Info("Debug mode enabled.")
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	client := gitlab.NewClient(utils.Getenv("GUZUTA_GITLAB_TOKEN"))
	if len(name) > 0 && len(namespace) > 0 {
		gitlabProjectStatus(client, namespace, name)
		return 0
	}
	if len(namespace) > 0 {
		gitlabProjectsStatus(client, namespace)
		return 0
	}
	return 0
}

func gitlabProjectStatus(client *gitlab.Client, namespace string, name string) {
	project, err := client.GetProject(namespace, name)
	if err != nil {
		colorstring.Printf("[red] Gitlab : %s\n", err.Error())
		return
	}
	gitlabPrintProject(project)
}

func gitlabProjectsStatus(client *gitlab.Client, namespace string) {
	resp, err := client.GetProjects()
	if err != nil {
		colorstring.Printf("[red] Gitlab : %s\n", err.Error())
		return
	}
	for _, repo := range *resp {
		gitlabPrintProject(&repo)
	}
}

func gitlabPrintProject(project *gitlab.Project) {
	fmt.Printf("* %s - %s\n", project.Name, project.Description)
}
