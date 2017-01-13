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

package gitlab

import (
	"fmt"
	"strings"
	"testing"

	"github.com/nlamirault/guzuta/utils"
)

var (
	namespace   = "nicolas-lamirault"
	name        = "scame"
	description = "An Emacs configuration"
)

func getGitlabClient() *Client {
	return NewClient(utils.Getenv("GUZUTA_GITLAB_TOKEN"))
}

func TestRetrieveGitLabProjects(t *testing.T) {
	client := getGitlabClient()
	projects, _ := client.GetProjects()
	for _, project := range *projects {
		path := fmt.Sprintf("%s/%s", namespace, project.Path)
		if !strings.HasPrefix(project.PathWithNamespace, path) {
			t.Fatalf("Invalid project name : %s %s",
				project.PathWithNamespace, path)
		}
	}
}

// FIXME
// func TestRetrieveGitLabProject(t *testing.T) {
// 	client := getGitlabClient()
// 	project, _ := client.GetProject(namespace, name)
// 	if !strings.HasPrefix(project.Name, name) {
// 		t.Fatalf("Invalid project name : %#v", project)
// 	}
// }

func TestRetrieveGitlabUnknownProject(t *testing.T) {
	client := getGitlabClient()
	_, err := client.GetProject(namespace, "aaaaaaaaaa")
	if err == nil {
		t.Fatalf("No error with unknown username")
	}
}
