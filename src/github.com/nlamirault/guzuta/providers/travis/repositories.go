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

package travis

import (
	//"encoding/json"
	"fmt"
	//"io"
	//"io/ioutil"
	"log"
	"net/http"

	"github.com/nlamirault/guzuta/utils"
)

// type RepositoryInput struct {
// }

type Repository struct {
	ID                  int64  `json:"id"`
	Slug                string `json:"slug"`
	Description         string `json:"description,omitempty"`
	LastBuildID         int64  `json:"last_build_id,omitempty"`
	LastBuildNumber     string `json:"last_build_number,omitempty"`
	LastBuildState      string `json:"last_build_state,omitempty"`
	LastBuildDuration   int64  `json:"last_build_duration,omitempty"`
	LastBuildStartedAt  string `json:"last_build_started_at,omitempty"`
	LastBuildFinishedAt string `json:"last_build_finished_at,omitempty"`
	GithubLanguage      string `json:"github_language,omitempty"`
}

type RepositoryOutput struct {
	Repository Repository `json:"repo"`
}

type RepositoriesOutput struct {
	Repositories []Repository `json:"repos"`
}

// GetRepository retrieve repository by ID or Slug
func (c *Client) GetRepository(name string) (*RepositoryOutput, error) {
	log.Printf("[DEBUG] Get repository: %s", name)
	var repository *RepositoryOutput
	resp, err := c.Do(
		"GET",
		fmt.Sprintf("repos/%s", name),
		nil) //&AuthenticateInput{Token: c.Token})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		err = utils.DecodeResponse(resp, &repository)
		if err != nil {
			return nil, err
		}
		log.Printf("[DEBUG] Repository: %v", repository)
	}
	return repository, nil
}

// GetRepositories retrieve repository by ID or Slug
func (c *Client) GetRepositories(namespace string) (*RepositoriesOutput, error) {
	log.Printf("[DEBUG] Get repository: %s", namespace)
	var repositories *RepositoriesOutput
	resp, err := c.Do(
		"GET",
		fmt.Sprintf("repos/%s", namespace),
		nil) //&AuthenticateInput{Token: c.Token})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		err = utils.DecodeResponse(resp, &repositories)
		if err != nil {
			return nil, err
		}
		log.Printf("[DEBUG] Repositories: %v", repositories)
	}
	return repositories, nil
}
