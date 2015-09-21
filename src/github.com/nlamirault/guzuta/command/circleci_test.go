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
	//"fmt"
	//"strings"
	"testing"

	"github.com/mitchellh/cli"
)

func TestCircleCI_WithoutArgs(t *testing.T) {
	ui := new(cli.MockUi)
	c := &CircleCICommand{UI: ui}
	args := []string{}
	if code := c.Run(args); code != 0 {
		t.Fatalf("bad: \n%s", ui.ErrorWriter.String())
	}
}

// func TestCircleCI_WithoutToken(t *testing.T) {
// 	ui := new(cli.MockUi)
// 	c := &CircleCICommand{UI: ui}
// 	args := []string{"--username", "nlamirault", "--name", "guzuta"}
// 	if code := c.Run(args); code != 0 {
// 		t.Fatalf("Error with invalid args: %s \n%s",
// 			code, ui.OutputWriter.String())
// 	}
// 	fmt.Printf(ui.OutputWriter.String())
// }
