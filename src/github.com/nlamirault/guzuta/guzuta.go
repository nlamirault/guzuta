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
	//"flag"
	"fmt"
	//"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/colorstring"

	"github.com/nlamirault/guzuta/command"
	"github.com/nlamirault/guzuta/version"
)

func main() {
	cli := &cli.CLI{
		Args:       os.Args[1:],
		Commands:   command.Commands,
		HelpFunc:   cli.BasicHelpFunc("guzuta"),
		HelpWriter: os.Stdout,
		Version:    version.Version,
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Printf(colorstring.Color(
			"[red] Error executing CLI: " + err.Error() + "\n"))
		return
	}

	os.Exit(exitCode)
}
