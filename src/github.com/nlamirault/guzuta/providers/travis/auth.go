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
	"log"
	"net/http"

	"github.com/nlamirault/guzuta/utils"
)

type AuthenticateInput struct {
	Token string `json:"github_token"`
}

type AuthenticateOutput struct {
	Token string `json:"access_token"`
}

// Authenticate retrieve access token
func (c *Client) Authenticate() error {
	log.Printf("[DEBUG] [travis] Authenticate")
	var token *AuthenticateOutput
	resp, err := c.Do("POST", "auth/github", AuthenticateInput{Token: c.Token})
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] [travis] Response : %#v %s", resp, resp.StatusCode)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil
	}
	utils.DecodeResponse(resp, &token)
	log.Printf("[DEBUG] [travis] Authorization: %v", token)
	c.AccessToken = token.Token
	return nil
}
