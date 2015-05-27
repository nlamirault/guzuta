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
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"github.com/nlamirault/guzuta/commands"
	mylog "github.com/nlamirault/guzuta/log"
	"github.com/nlamirault/guzuta/version"
)

func makeApp() *cli.App {
	app := cli.NewApp()
	app.Name = "scaleway"
	app.Version = version.Version
	app.Usage = "A CLI for Scaleway"
	app.Author = "Nicolas Lamirault"
	app.Email = "nicolas.lamirault@gmail.com"
	app.CommandNotFound = cmdNotFound
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: fmt.Sprintf("Log level (options: debug, info, warn, error, fatal, panic)"),
		},
		cli.StringFlag{
			Name:   "github-token",
			Usage:  "Github access token",
			Value:  "",
			EnvVar: "GUZUTA_GITHUB_TOKEN",
		},
	}
	app.Before = func(c *cli.Context) error {
		log.SetFormatter(&mylog.CustomFormatter{})
		log.SetOutput(os.Stderr)
		level, err := log.ParseLevel(c.String("log-level"))
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.SetLevel(level)
		return nil
	}

	app.Commands = commands.Commands
	//app.Flags = commands.Flags
	return app
}

func cmdNotFound(c *cli.Context, command string) {
	log.Fatalf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		c.App.Name,
	)
}

func main() {
	app := makeApp()
	app.Run(os.Args)
}
