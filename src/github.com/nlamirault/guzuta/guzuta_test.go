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
	"strings"
	"testing"

	"github.com/codegangsta/cli"
)

func TestCLICommands(t *testing.T) {
	app := makeApp()
	if len(app.Commands) != 1 {
		t.Errorf("Invalid CLI number of commands")
	}
}

func checkGlobalArgument(flags []cli.Flag, name string) int {
	for i, flag := range flags {
		fmt.Printf("Flag: %v\n", flag.String())
		if strings.HasPrefix(flag.String(), name) {
			return i
		}
	}
	return -1
}

func TestGithubTokenArgument(t *testing.T) {
	app := makeApp()
	if checkGlobalArgument(app.Flags, "--github-token") == -1 {
		t.Errorf("No token flag")
	}
}
