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
	"fmt"
	"log"
	"net/http"
	"net/url"

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

	// SetupHeaders add customer headers
	SetupHeaders(request *http.Request)

	// Do perform a HTTP request
	Do(method, urlStr string, body interface{}) (*http.Response, error)
}

func GetURL(base *url.URL, urlStr string) (*url.URL, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	return base.ResolveReference(rel), nil
}

// MakeRequest creates a HTTP request given a method, URL, and optional body.
// func MakeRequest(method, base *url.URL, urlStr string, body interface{}) (*http.Request, error) {
// 	_, err := getURL(base, urlStr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	buf := new(bytes.Buffer)
// 	if body != nil {
// 		err := json.NewEncoder(buf).Encode(body)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	// req, err := http.NewRequest(method, u.String(), buf)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// log.Printf("[DEBUG] Request : %v", req)
// 	// return req, nil
// 	return nil, nil
// }

func CreateRequest(method, uri string, body interface{}) (*http.Request, error) {
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
