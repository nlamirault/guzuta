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

package travisci

import (
	"fmt"
	"strings"
	"testing"
)

var (
	namespace   = "kubernetes"
	name        = "kubernetes"
	description = "Container Cluster Manager from Google"
)

func TestRetrieveRepositories(t *testing.T) {
	client := NewClient("0246813579")
	resp, _ := client.GetRepositories(namespace)
	for _, repo := range resp.Repositories {
		if !strings.HasPrefix(repo.Slug, namespace) {
			t.Fatalf("Invalid Slug : %s", repo)
		}
	}
}

func TestRetrieveRepository(t *testing.T) {
	client := NewClient("0246813579")
	slug := fmt.Sprintf("%s/%s", namespace, name)
	resp, _ := client.GetRepository(slug)
	fmt.Printf("=> %#v", resp.Repository)
	if !strings.HasPrefix(resp.Repository.Slug, slug) {
		t.Fatalf("Invalid Slug : %s", resp.Repository)
	}
	if resp.Repository.Description != description {
		t.Fatalf("Invalid Description : %s", resp.Repository)
	}
}
