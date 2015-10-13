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

package circleci

import (
	//"fmt"
	//"strings"
	"testing"

	"github.com/nlamirault/guzuta/utils"
)

var (
	username    = "nlamirault"
	name        = "guzuta"
	description = "A CLI to manage personal open source contributions"
)

func TestRetrieveCircleCIProjects(t *testing.T) {
	client := NewClient(utils.Getenv("GUZUTA_CIRCLECI_TOKEN"))
	projects, _ := client.GetProjects()
	for _, p := range *projects {
		if len(p.Username) == 0 {
			t.Fatalf("Invalid project : %#v", p)
		}
	}
}

func TestRetrieveCircleCIProject(t *testing.T) {
	client := NewClient(utils.Getenv("GUZUTA_CIRCLECI_TOKEN"))
	builds, _ := client.GetProject(&ProjectInput{
		Username: username,
		Project:  name,
		Limit:    1,
	})
	for _, build := range *builds {
		if len(build.VcsRevision) == 0 {
			t.Fatalf("Invalid project : %#v", build)
		}
	}
}

func TestRetrieveCircleCIUnknownProject(t *testing.T) {
	client := NewClient(utils.Getenv("GUZUTA_CIRCLECI_TOKEN"))
	_, err := client.GetProject(&ProjectInput{
		Username: username,
		Project:  "cddedsddcsdcsd",
		Limit:    1,
	})
	if err == nil {
		t.Fatalf("No error with unknown repository")
	}

}
