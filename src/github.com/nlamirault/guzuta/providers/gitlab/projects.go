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
	//"encoding/json"
	//"errors"
	"fmt"
	//"io"
	//"io/ioutil"
	"log"
	//"net/http"
	"time"

	"github.com/nlamirault/guzuta/providers"
)

type User struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	Name             string    `json:"name"`
	State            string    `json:"state"`
	CreatedAt        time.Time `json:"created_at"`
	Bio              string    `json:"bio"`
	Skype            string    `json:"skype"`
	Linkedin         string    `json:"linkedin"`
	Twitter          string    `json:"twitter"`
	WebsiteURL       string    `json:"website_url"`
	ExternUID        string    `json:"extern_uid"`
	Provider         string    `json:"provider"`
	ThemeID          int       `json:"theme_id"`
	ColorSchemeID    int       `json:"color_scheme_id"`
	IsAdmin          bool      `json:"is_admin"`
	AvatarURL        string    `json:"avatar_url"`
	CanCreateGroup   bool      `json:"can_create_group"`
	CanCreateProject bool      `json:"can_create_project"`
	ProjectsLimit    int       `json:"projects_limit"`
	CurrentSignInAt  time.Time `json:"current_sign_in_at"`
	TwoFactorEnabled bool      `json:"two_factor_enabled"`
}

type Namespace struct {
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	OwnerID     int       `json:"owner_id"`
	Path        string    `json:"path"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Project struct {
	ID                   int         `json:"id"`
	Description          interface{} `json:"description"`
	DefaultBranch        string      `json:"default_branch"`
	Public               bool        `json:"public"`
	VisibilityLevel      int         `json:"visibility_level"`
	SSHURLToRepo         string      `json:"ssh_url_to_repo"`
	HTTPURLToRepo        string      `json:"http_url_to_repo"`
	WebURL               string      `json:"web_url"`
	TagList              []string    `json:"tag_list"`
	Owner                *User       `json:"owner"`
	Name                 string      `json:"name"`
	NameWithNamespace    string      `json:"name_with_namespace"`
	Path                 string      `json:"path"`
	PathWithNamespace    string      `json:"path_with_namespace"`
	IssuesEnabled        bool        `json:"issues_enabled"`
	MergeRequestsEnabled bool        `json:"merge_requests_enabled"`
	WikiEnabled          bool        `json:"wiki_enabled"`
	SnippetsEnabled      bool        `json:"snippets_enabled"`
	CreatedAt            time.Time   `json:"created_at"`
	LastActivityAt       time.Time   `json:"last_activity_at"`
	CreatorID            int         `json:"creator_id"`
	Namespace            Namespace   `json:"namespace"`
	Archived             bool        `json:"archived"`
	AvatarURL            string      `json:"avatar_url"`
}

// GetProjects retrieve all repositories for user
func (c *Client) GetProjects() (*[]Project, error) {
	log.Printf("[DEBUG] Get projects")
	var projects *[]Project
	err := providers.Do(c, "GET", "projects", nil, &projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// GetProject retrieve project
func (c *Client) GetProject(namespace string, name string) (*Project, error) {
	log.Printf("[DEBUG] Get project: %s %s", namespace, name)
	var project *Project
	err := providers.Do(
		c,
		"GET",
		fmt.Sprintf("projects/%s/%s", namespace, name),
		nil,
		&project)
	if err != nil {
		return nil, err
	}
	return project, nil
}
