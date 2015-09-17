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
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/nlamirault/guzuta/version"
)

const (
	apiURL       = "https://api.travis-ci.org/"
	acceptHeader = "application/vnd.travis-ci.2+json"
	mediaType    = "application/json"
)

var (
	userAgent = fmt.Sprintf("guzuta/%s", version.Version)
)

// Client is the Travis API client
type Client struct {

	// The User Agent of the client
	UserAgent string

	// Base URL for API requests.
	BaseURL *url.URL

	// The token used to communicate
	Token string

	// The HTTP client to use when sending requests.
	HTTPClient *http.Client
}

// NewClient returns a new Travis API client instance
func NewClient(token string) *Client {
	log.Printf("[DEBUG] [travis] Client creation : %s", token)
	baseURL, _ := url.Parse(apiURL)
	client := &Client{
		Token:      token,
		UserAgent:  userAgent,
		HTTPClient: http.DefaultClient,
		BaseURL:    baseURL,
	}
	log.Printf("[DEBUG] [travis] Client created : %v", client)
	return client
}

// Response is a Travis response.
type Response struct {
	*http.Response
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// Error message
	Message string
}

// Do perform a HTTP request
func (c *Client) Do(method, urlStr string, body interface{}) (*http.Response, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", acceptHeader)
	req.Header.Add("User-Agent", userAgent)
	log.Printf("[DEBUG] [travis] Perform request : %v", req)
	return c.HTTPClient.Do(req)
}

// func (c *Client) perform(req *http.Request) (*http.Response, error) {
// 	resp, err := c.HTTPClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// }
