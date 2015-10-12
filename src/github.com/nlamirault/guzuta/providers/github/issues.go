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
	"log"
	"time"

	"github.com/nlamirault/guzuta/providers"
)

// Label represents a GitHub label on an Issue
type Label struct {
	URL   string `json:"url,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

// Milestone represents a Github repository milestone.
type Milestone struct {
	URL          string `json:"url,omitempty"`
	Number       int    `json:"number,omitempty"`
	State        string `json:"state,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Creator      User   `json:"creator,omitempty"`
	OpenIssues   int    `json:"open_issues,omitempty"`
	ClosedIssues int    `json:"closed_issues,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	DueOn        string `json:"due_on,omitempty"`
}

// PullRequestLinks object is added to the Issue object when it's an issue included
// in the IssueCommentEvent webhook payload, if the webhooks is fired by a comment on a PR
type PullRequestLinks struct {
	URL      *string `json:"url,omitempty"`
	HTMLURL  *string `json:"html_url,omitempty"`
	DiffURL  *string `json:"diff_url,omitempty"`
	PatchURL *string `json:"patch_url,omitempty"`
}

// Issue represents a GitHub issue on a repository.
type Issue struct {
	Number           int              `json:"number,omitempty"`
	State            string           `json:"state,omitempty"`
	Title            string           `json:"title,omitempty"`
	Body             string           `json:"body,omitempty"`
	User             User             `json:"user,omitempty"`
	Labels           []Label          `json:"labels,omitempty"`
	Assignee         User             `json:"assignee,omitempty"`
	Comments         int              `json:"comments,omitempty"`
	ClosedAt         string           `json:"closed_at,omitempty"`
	CreatedAt        string           `json:"created_at,omitempty"`
	UpdatedAt        string           `json:"updated_at,omitempty"`
	URL              string           `json:"url,omitempty"`
	HTMLURL          string           `json:"html_url,omitempty"`
	Milestone        Milestone        `json:"milestone,omitempty"`
	PullRequestLinks PullRequestLinks `json:"pull_request,omitempty"`

	// TextMatches is only populated from search results that request text matches
	// See: search.go and https://developer.github.com/v3/search/#text-match-metadata
	// TextMatches []TextMatch `json:"text_matches,omitempty"`
}

// IssueListOptions specifies the optional parameters to the IssuesService.List
// and IssuesService.ListByOrg methods.
type IssueListOptions struct {
	// Filter specifies which issues to list.  Possible values are: assigned,
	// created, mentioned, subscribed, all.  Default is "assigned".
	Filter string `url:"filter,omitempty"`

	// State filters issues based on their state.  Possible values are: open,
	// closed.  Default is "open".
	State string `url:"state,omitempty"`

	// Labels filters issues based on their label.
	Labels []string `url:"labels,comma,omitempty"`

	// Sort specifies how to sort issues.  Possible values are: created, updated,
	// and comments.  Default value is "created".
	Sort string `url:"sort,omitempty"`

	// Direction in which to sort issues.  Possible values are: asc, desc.
	// Default is "asc".
	Direction string `url:"direction,omitempty"`

	// Since filters issues by time.
	Since time.Time `url:"since,omitempty"`

	ListOptions
}

func (c *Client) ListIssues(opt *IssueListOptions) (*[]Issue, error) {
	log.Printf("[DEBUG] List issues : %v", opt)
	uri, err := addOptions("issues", opt)
	if err != nil {
		return nil, err
	}
	var issues *[]Issue
	err = providers.Do(c, "GET", uri, nil, &issues)
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func (c *Client) ListRepositoryIssues(username string, name string, opt *IssueListOptions) (*[]Issue, error) {
	log.Printf("[DEBUG] List issues for : %s/%s %v", username, name, opt)
	uri, err := addOptions(
		fmt.Sprintf("repos/%s/%s/issues", username, name), opt)
	if err != nil {
		return nil, err
	}
	var issues *[]Issue
	err = providers.Do(c, "GET", uri, nil, &issues)
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func (c *Client) GetRepositoryIssue(username string, name string, issueID int) (*Issue, error) {
	log.Printf("[DEBUG] Get issue for : %s/%s %d", username, name, issueID)
	var issue *Issue
	err := providers.Do(
		c,
		"GET",
		fmt.Sprintf("repos/%s/%s/issues/%d", username, name, issueID),
		nil,
		&issue)
	if err != nil {
		return nil, err
	}
	return issue, nil
}
