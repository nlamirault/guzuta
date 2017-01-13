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

package github

import (
	"fmt"
	"strings"
	"testing"

	"github.com/nlamirault/guzuta/utils"
)

var (
	username    = "nlamirault"
	name        = "guzuta"
	description = "A CLI to manage personal open source contributions"
)

func getGithubClient() *Client {
	return NewClient(utils.Getenv("GUZUTA_GITHUB_TOKEN"))
}

func TestRetrieveGithubRepositories(t *testing.T) {
	client := getGithubClient()
	resp, _ := client.GetRepositories(username)
	for _, repo := range *resp {
		if !strings.HasPrefix(repo.FullName, username) {
			t.Fatalf("Invalid full name : %#v", repo)
		}
	}
}

func TestRetrieveGithubUnknownUsername(t *testing.T) {
	client := getGithubClient()
	_, err := client.GetRepositories("azecslcklnlsdcnsjkdn")
	if err == nil {
		t.Fatalf("No error with unknown username")
	}
}

func TestRetrieveGithubRepository(t *testing.T) {
	client := getGithubClient()
	repo, _ := client.GetRepository(username, name)
	fullname := fmt.Sprintf("%s/%s", username, name)
	if !strings.HasPrefix(repo.FullName, fullname) {
		t.Fatalf("Invalid full name : %#v", repo)
	}
}

func TestRetrieveGithubUnknownRepository(t *testing.T) {
	client := getGithubClient()
	_, err := client.GetRepository(username, "azecslcklnlsdcnsjkdn")
	if err == nil {
		t.Fatalf("No error with unknown repository")
	}
}
