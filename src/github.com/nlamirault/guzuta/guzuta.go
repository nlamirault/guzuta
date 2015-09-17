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

	"github.com/nlamirault/guzuta/ci/travis"
	"github.com/nlamirault/guzuta/logging"
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
)

func init() {
	// parse flags
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.BoolVar(&showVersion, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	// Travis
	flag.StringVar(&travisRepository, "travis-name", "", "Status of repository")
	flag.StringVar(&travisRepositories, "travis-namespace", "", "Status of repositories")
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
	color := "green"
	if resp.Repository.LastBuildState == "failed" {
		color = "red"
	}
	fmt.Printf(resp.Repository.Slug + "\t" +
		colorstring.Color("["+color+"]"+resp.Repository.LastBuildState) + "\n")
}

func travisRepositoriesStatus(travis *travis.Client, namespace string) {
	resp, err := travis.GetRepositories(namespace)
	if err != nil {
		colorstring.Printf("[red] Travis : %s", err.Error())
		return
	}
	for _, repo := range resp.Repositories {
		color := "green"
		if repo.LastBuildState == "failed" {
			color = "red"
		}
		fmt.Printf(repo.Slug + "\t" +
			colorstring.Color("["+color+"]"+repo.LastBuildState) + "\n")
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
	// fmt.Println("Guzuta")
	travis := travis.NewClient(utils.Getenv("GUZUTA_TRAVIS_GITHUB_TOKEN"))
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
}
