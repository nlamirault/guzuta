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
	// "bytes"
	// "encoding/json"
	//"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/nlamirault/guzuta/providers"
)

const (
	apiURL = "https://circleci.com/api/v1/"
)

// Client is the CircleCI API client
type Client struct {

	// The User Agent of the client
	UserAgent string

	// Endpoint is the base URL for API requests.
	Endpoint *url.URL

	// The token used to send API requests
	Token string

	// The HTTP client to use when sending requests.
	HTTPClient *http.Client
}

type APIError struct {
	Message string `json:"message"`
}

// NewClient returns a new CircleCI API client instance
func NewClient(token string) *Client {
	log.Printf("[DEBUG] [circleci] Client creation : %s", token)
	baseURL, _ := url.Parse(apiURL)
	client := &Client{
		Token:      token,
		UserAgent:  providers.UserAgent,
		HTTPClient: http.DefaultClient,
		Endpoint:   baseURL,
	}
	log.Printf("[DEBUG] [circleci] Client created : %v", client)
	return client
}

func (c *Client) SetupHeaders(request *http.Request) {
	request.Header.Add("Content-Type", providers.MediaType)
	request.Header.Add("Accept", providers.AcceptHeader)
	request.Header.Add("User-Agent", c.UserAgent)
}

func (c *Client) EndPoint() *url.URL {
	return c.Endpoint
}

func (c *Client) GetHTTPClient() *http.Client {
	return c.HTTPClient
}

func (c *Client) Do(method, urlStr string, body interface{}) (*http.Response, error) {
	return providers.PerformRequest(c, method, urlStr, body)
}
