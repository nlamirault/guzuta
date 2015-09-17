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
	"log"

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
	debug       bool
	showVersion bool
)

func init() {
	// parse flags
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.BoolVar(&showVersion, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.Parse()
}

func getConfigurationFile() string {
	return fmt.Sprintf("%s/.config/guzuta/guzuta.yml", utils.UserHomeDir())
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
	log.Printf("[INFO] Guzuta")
	travis := travis.NewClient("64c3acc2c2a010d18e6314ba9db85df51d4f7ea2")
	token, err := travis.Authenticate()
	if err != nil {
		log.Printf("[INFO] Travis error : %#v", err)
		return
	}
	log.Printf("[INFO] Done %v", token)
}
