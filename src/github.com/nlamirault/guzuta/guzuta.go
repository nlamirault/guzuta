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

package main

import (
	"flag"
	"fmt"
	//"log"

	"github.com/mitchellh/colorstring"

	"github.com/nlamirault/guzuta/logging"
	"github.com/nlamirault/guzuta/providers/circleci"
	"github.com/nlamirault/guzuta/providers/travis"
	"github.com/nlamirault/guzuta/utils"
	"github.com/nlamirault/guzuta/version"
)

const (
	// APP is the application name
	APP string = "guzuta"
)

var (
	debug              bool
	showVersion        bool
	travisRepository   string
	travisRepositories string
	circleciProject    string
	circleciUsername   string
)

func init() {
	// parse flags
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.BoolVar(&showVersion, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	// Travis
	flag.StringVar(&travisRepository, "travis-name", "", "Status of repository")
	flag.StringVar(&travisRepositories, "travis-namespace", "", "Status of repositories")
	flag.StringVar(&circleciProject, "circleci-project", "", "Project name")
	flag.StringVar(&circleciUsername, "circleci-username", "", "Username")
	flag.Parse()
}

func getConfigurationFile() string {
	return fmt.Sprintf("%s/.config/guzuta/guzuta.yml", utils.UserHomeDir())
}

func travisRepositoryStatus(travis *travis.Client, name string) {
	resp, err := travis.GetRepository(name)
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

func travisRepositoriesStatus(travis *travis.Client, namespace string) {
	resp, err := travis.GetRepositories(namespace)
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

func circleciProjectStatus(client *circleci.Client, username string, project string) {
	resp, err := client.GetProject(&circleci.ProjectInput{
		Username: username,
		Project:  project,
		Limit:    1,
	})
	if err != nil {
		colorstring.Printf("[red] Circleci : %s", err.Error())
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
		colorstring.Printf("[red] Circleci : %s", err.Error())
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

func main() {
	if debug {
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	if showVersion {
		fmt.Printf("%s v%s\n", APP, version.Version)
		return
	}
	travis := travis.NewClient(utils.Getenv("GUZUTA_TRAVIS_GITHUB_TOKEN"))
	circleci := circleci.NewClient(utils.Getenv("GUZUTA_CIRCLECI_TOKEN"))
	// err := travis.Authenticate()
	// if err != nil {
	// 	colorstring.Printf("[red] Travis error : %s", err.Error())
	// 	return
	// }
	if len(travisRepository) > 0 {
		travisRepositoryStatus(travis, travisRepository)
		return
	}
	if len(travisRepositories) > 0 {
		travisRepositoriesStatus(travis, travisRepositories)
		return
	}
	if len(circleciProject) > 0 && len(circleciUsername) > 0 {
		circleciProjectStatus(circleci, circleciUsername, circleciProject)
		return
	}
	if len(circleciUsername) > 0 {
		circleciProjectsStatus(circleci)
		return
	}
}
