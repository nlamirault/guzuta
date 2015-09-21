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
	// "bytes"
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/nlamirault/guzuta/providers"
)

const (
	apiURL       = "https://api.travis-ci.org/"
	acceptHeader = "application/vnd.travis-ci.2+json"
)

// Client is the Travis API client
type Client struct {

	// The User Agent of the client
	UserAgent string

	// Endpoint is the base URL for API requests.
	Endpoint *url.URL

	// The token used to authenticate
	Token string

	// The token used to send API requests
	AccessToken string

	// The HTTP client to use when sending requests.
	HTTPClient *http.Client
}

// NewClient returns a new Travis API client instance
func NewClient(token string) *Client {
	log.Printf("[DEBUG] [travis] Client creation : %s", token)
	baseURL, _ := url.Parse(apiURL)
	client := &Client{
		Token:      token,
		UserAgent:  providers.UserAgent,
		HTTPClient: http.DefaultClient,
		Endpoint:   baseURL,
	}
	log.Printf("[DEBUG] [travis] Client created : %v", client)
	return client
}

func (c *Client) SetupHeaders(request *http.Request) {
	request.Header.Add("Content-Type", providers.MediaType)
	request.Header.Add("Accept", acceptHeader)
	request.Header.Add("User-Agent", c.UserAgent)
	if len(c.AccessToken) > 0 {
		request.Header.Add(
			"Authorization", fmt.Sprintf("token %s", c.AccessToken))
	}
}

func (c *Client) Do(method, urlStr string, body interface{}) (*http.Response, error) {
	u, err := providers.GetURL(c.Endpoint, urlStr)
	if err != nil {
		return nil, err
	}
	req, err := providers.CreateRequest(method, u.String(), body) //, c.BaseURL, urlStr, body)
	if err != nil {
		return nil, err
	}
	c.SetupHeaders(req)
	return c.HTTPClient.Do(req)
}
