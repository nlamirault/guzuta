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

package providers

import (
	"bytes"
	"encoding/json"
	//"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/nlamirault/guzuta/utils"
	"github.com/nlamirault/guzuta/version"
)

const (
	// AcceptHeader is the default Accept Header : application/json
	AcceptHeader = "application/json"

	// MediaType is the default mediaType header : application/json
	MediaType = "application/json"
)

var (
	// UserAgent represents the user agent used
	UserAgent = fmt.Sprintf("guzuta/%s", version.Version)
)

// APIClient represents a client for a REST API
type APIClient interface {

	// EndPoint returns the API base URL
	EndPoint() *url.URL

	// GetHTTPClient returns the HTTP client to use
	GetHTTPClient() *http.Client

	// SetupHeaders add customer headers
	SetupHeaders(request *http.Request)

	// Do perform a HTTP request
	// Do(method, urlStr string, body interface{}) (*http.Response, error)
}

func getURL(base *url.URL, urlStr string) (*url.URL, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	return base.ResolveReference(rel), nil
}

func createRequest(method, uri string, body interface{}) (*http.Request, error) {
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri, buf)
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] Request : %v", req)
	return req, nil
}

func performRequest(client APIClient, method, urlStr string, body interface{}) (*http.Response, error) {
	u, err := getURL(client.EndPoint(), urlStr)
	if err != nil {
		return nil, err
	}
	req, err := createRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}
	client.SetupHeaders(req)
	return client.GetHTTPClient().Do(req)
}

// APIError represents an error from REST API
type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (a *APIError) Error() string {
	return fmt.Sprintf("%d / %s", a.StatusCode, a.Message)
}

// Do perform a HTTP request using the REST API Client.
// body is used for the content of the request
// result contains the JSON decoded response
// apiError contains the JSON error response if HTTP status code isn't OK.
func Do(client APIClient, method, urlStr string, body interface{}, result interface{}) error {
	resp, err := performRequest(client, method, urlStr, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return utils.DecodeResponse(resp, result)
	}
	content, err := utils.GetResponseBody(resp)
	if err != nil {
		return fmt.Errorf("Can't read API Error : %v", err)
	}
	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    content,
	}
}
