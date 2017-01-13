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
