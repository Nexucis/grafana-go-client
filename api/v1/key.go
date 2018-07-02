// Copyright 2018 Augustin Husson
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"github.com/nexucis/grafana-go-client/http"
	"strconv"
)

const keyAPI = "/api/auth/keys"

type KeyInterface interface {
	Get() ([]*GetAPIKeyResponse, error)
	Create(*APIKeyForm) (*CreateAPIKeyResponse, error)
	Delete(int64) error
}

func newKey(client *http.RESTClient) KeyInterface {
	return &key{
		client: client,
	}
}

type key struct {
	KeyInterface
	client *http.RESTClient
}

func (c *key) Get() ([]*GetAPIKeyResponse, error) {
	var result []*GetAPIKeyResponse
	err := c.client.Get(keyAPI).
		Do().
		SaveAsObj(&result)
	return result, err
}

func (c *key) Create(*APIKeyForm) (*CreateAPIKeyResponse, error) {
	result := &CreateAPIKeyResponse{}
	err := c.client.Post(keyAPI).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *key) Delete(id int64) error {
	return c.client.Delete(keyAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(id, 10)).
		Do().
		Error()
}
