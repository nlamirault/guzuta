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
	//"encoding/json"
	//"errors"
	"fmt"
	//"io"
	//"io/ioutil"
	"log"
	//"net/http"

	"github.com/nlamirault/guzuta/providers"
)

// ProjectInput represents project parameters
type ProjectInput struct {
	Username string `json:"username"`
	Project  string `json:"project"`
	Limit    int    `json:"limit"`
}

type Build struct {
	AddedAt     string `json:"added_at"`
	PushedAt    string `json:"pushed_at"`
	VcsRevision string `json:"vcs_revision"`
	BuildNum    int    `json:"build_num"`
	Status      string `json:"status"`
	Outcome     string `json:"outcome"`
}

type Branch struct {
	LastNonSuccess Build    `json:"last_non_success"`
	LastSuccess    Build    `json:"last_success"`
	RecentBuilds   []Build  `json:"recent_builds"`
	RunningBuilds  []Build  `json:"running_builds"`
	PusherLogins   []string `json:"pushing_logins"`
}

type BranchBuilds struct {
	Master   Branch `json:"master"`
	Develop  Branch `json:"develop"`
	Unstable Branch `json:"unstable"`
}

// ProjectOutput represents project content
type Project struct {
	Username string       `json:"username"`
	Reponame string       `json:"reponame"`
	Branches BranchBuilds `json:"branches"`
}

// GetProjects retrieve all projects for user
func (c *Client) GetProjects() (*[]Project, error) {
	log.Printf("[DEBUG] Get projects")
	var projects *[]Project
	err := providers.Do(
		c,
		"GET",
		fmt.Sprintf("projects?circle-token=%s", c.Token),
		nil,
		&projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// GetProject retrieve repositorys by name
func (c *Client) GetProject(input *ProjectInput) (*[]Build, error) {
	log.Printf("[DEBUG] Get project: %s %s", input.Username, input.Project)
	var project *[]Build
	err := providers.Do(
		c,
		"GET",
		fmt.Sprintf("project/%s/%s?limit=%d&circle-token=%s",
			input.Username, input.Project, input.Limit, c.Token),
		nil,
		&project)
	if err != nil {
		return nil, err
	}
	return project, nil
}
