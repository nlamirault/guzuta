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
	//"encoding/json"
	"errors"
	"fmt"
	//"io"
	//"io/ioutil"
	"log"
	"net/http"

	"github.com/nlamirault/guzuta/utils"
)

// User represents a GitHub user.
type User struct {
	Login             string `json:"login,omitempty"`
	ID                int    `json:"id,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	HTMLURL           string `json:"html_url,omitempty"`
	GravatarID        string `json:"gravatar_id,omitempty"`
	Name              string `json:"name,omitempty"`
	Company           string `json:"company,omitempty"`
	Blog              string `json:"blog,omitempty"`
	Location          string `json:"location,omitempty"`
	Email             string `json:"email,omitempty"`
	Hireable          bool   `json:"hireable,omitempty"`
	Bio               string `json:"bio,omitempty"`
	PublicRepos       int    `json:"public_repos,omitempty"`
	PublicGists       int    `json:"public_gists,omitempty"`
	Followers         int    `json:"followers,omitempty"`
	Following         int    `json:"following,omitempty"`
	CreatedAt         string `json:"created_at,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	Type              string `json:"type,omitempty"`
	SiteAdmin         bool   `json:"site_admin,omitempty"`
	TotalPrivateRepos int    `json:"total_private_repos,omitempty"`
	OwnedPrivateRepos int    `json:"owned_private_repos,omitempty"`
	PrivateGists      int    `json:"private_gists,omitempty"`
	DiskUsage         int    `json:"disk_usage,omitempty"`
	Collaborators     int    `json:"collaborators,omitempty"`

	// API URLs
	URL               string `json:"url,omitempty"`
	EventsURL         string `json:"events_url,omitempty"`
	FollowingURL      string `json:"following_url,omitempty"`
	FollowersURL      string `json:"followers_url,omitempty"`
	GistsURL          string `json:"gists_url,omitempty"`
	OrganizationsURL  string `json:"organizations_url,omitempty"`
	ReceivedEventsURL string `json:"received_events_url,omitempty"`
	ReposURL          string `json:"repos_url,omitempty"`
	StarredURL        string `json:"starred_url,omitempty"`
	SubscriptionsURL  string `json:"subscriptions_url,omitempty"`
}

// Repository represents a GitHub repository.
type Repository struct {
	ID               int    `json:"id,omitempty"`
	Owner            User   `json:"owner,omitempty"`
	Name             string `json:"name,omitempty"`
	FullName         string `json:"full_name,omitempty"`
	Description      string `json:"description,omitempty"`
	Homepage         string `json:"homepage,omitempty"`
	DefaultBranch    string `json:"default_branch,omitempty"`
	MasterBranch     string `json:"master_branch,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	PushedAt         string `json:"pushed_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	HTMLURL          string `json:"html_url,omitempty"`
	CloneURL         string `json:"clone_url,omitempty"`
	GitURL           string `json:"git_url,omitempty"`
	MirrorURL        string `json:"mirror_url,omitempty"`
	SSHURL           string `json:"ssh_url,omitempty"`
	SVNURL           string `json:"svn_url,omitempty"`
	Language         string `json:"language,omitempty"`
	Fork             bool   `json:"fork"`
	ForksCount       int    `json:"forks_count,omitempty"`
	NetworkCount     int    `json:"network_count,omitempty"`
	OpenIssuesCount  int    `json:"open_issues_count,omitempty"`
	StargazersCount  int    `json:"stargazers_count,omitempty"`
	SubscribersCount int    `json:"subscribers_count,omitempty"`
	WatchersCount    int    `json:"watchers_count,omitempty"`
	Size             int    `json:"size,omitempty"`

	// API URLs
	URL              string `json:"url,omitempty"`
	ArchiveURL       string `json:"archive_url,omitempty"`
	AssigneesURL     string `json:"assignees_url,omitempty"`
	BlobsURL         string `json:"blobs_url,omitempty"`
	BranchesURL      string `json:"branches_url,omitempty"`
	CollaboratorsURL string `json:"collaborators_url,omitempty"`
	CommentsURL      string `json:"comments_url,omitempty"`
	CommitsURL       string `json:"commits_url,omitempty"`
	CompareURL       string `json:"compare_url,omitempty"`
	ContentsURL      string `json:"contents_url,omitempty"`
	ContributorsURL  string `json:"contributors_url,omitempty"`
	DownloadsURL     string `json:"downloads_url,omitempty"`
	EventsURL        string `json:"events_url,omitempty"`
	ForksURL         string `json:"forks_url,omitempty"`
	GitCommitsURL    string `json:"git_commits_url,omitempty"`
	GitRefsURL       string `json:"git_refs_url,omitempty"`
	GitTagsURL       string `json:"git_tags_url,omitempty"`
	HooksURL         string `json:"hooks_url,omitempty"`
	IssueCommentURL  string `json:"issue_comment_url,omitempty"`
	IssueEventsURL   string `json:"issue_events_url,omitempty"`
	IssuesURL        string `json:"issues_url,omitempty"`
	KeysURL          string `json:"keys_url,omitempty"`
	LabelsURL        string `json:"labels_url,omitempty"`
	LanguagesURL     string `json:"languages_url,omitempty"`
	MergesURL        string `json:"merges_url,omitempty"`
	MilestonesURL    string `json:"milestones_url,omitempty"`
	NotificationsURL string `json:"notifications_url,omitempty"`
	PullsURL         string `json:"pulls_url,omitempty"`
	ReleasesURL      string `json:"releases_url,omitempty"`
	StargazersURL    string `json:"stargazers_url,omitempty"`
	StatusesURL      string `json:"statuses_url,omitempty"`
	SubscribersURL   string `json:"subscribers_url,omitempty"`
	SubscriptionURL  string `json:"subscription_url,omitempty"`
	TagsURL          string `json:"tags_url,omitempty"`
	TreesURL         string `json:"trees_url,omitempty"`
	TeamsURL         string `json:"teams_url,omitempty"`
}

// GetRepositories retrieve all projects for user
func (c *Client) GetRepositories(username string) (*[]Repository, error) {
	log.Printf("[DEBUG] Get repositories : %s", username)
	var repositories *[]Repository
	resp, err := c.Do(
		"GET",
		fmt.Sprintf("users/%s/repos", username),
		nil)
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
	} else {
		var apiError *APIError
		err = utils.DecodeResponse(resp, &apiError)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(apiError.Message)
	}
	return repositories, nil
}

// GetRepositories retrieve all projects for user
func (c *Client) GetRepository(username string, name string) (*Repository, error) {
	log.Printf("[DEBUG] Get repository : %s %s", username, name)
	var repository *Repository
	resp, err := c.Do(
		"GET",
		fmt.Sprintf("repos/%s/%s", username, name),
		nil)
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
	} else {
		var apiError *APIError
		err = utils.DecodeResponse(resp, &apiError)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(apiError.Message)
	}
	return repository, nil
}
