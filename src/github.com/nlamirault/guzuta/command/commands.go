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
	"os"

	"github.com/mitchellh/cli"

	"github.com/nlamirault/guzuta/logging"
)

// Commands is the mapping of all the available Terraform commands.
var (
	Commands map[string]cli.CommandFactory
	UI       cli.Ui
)

type Meta struct {
	UI cli.Ui
}

func init() {
	UI = &cli.ColoredUi{
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			Reader:      os.Stdin,
			ErrorWriter: os.Stderr,
		},
		OutputColor: cli.UiColorNone,
		InfoColor:   cli.UiColorNone,
		ErrorColor:  cli.UiColorRed,
	}

	Commands = map[string]cli.CommandFactory{
		"travisci": func() (cli.Command, error) {
			return &TravisCICommand{
				UI: UI,
			}, nil
		},
		"circleci": func() (cli.Command, error) {
			return &CircleCICommand{
				UI: UI,
			}, nil
		},
		"github": func() (cli.Command, error) {
			return &GithubCommand{
				UI: UI,
			}, nil
		},
		"gitlab": func() (cli.Command, error) {
			return &GitlabCommand{
				UI: UI,
			}, nil
		},
	}
}

func errorMessage(ui cli.Ui, msg string, help string) {
	ui.Error(msg)
	ui.Error("")
	ui.Error(help)
}

func setupLogging(debug bool) {
	if debug {
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
}
