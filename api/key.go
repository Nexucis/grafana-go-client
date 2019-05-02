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

package api

import (
	"github.com/nexucis/grafana-go-client/grafanahttp"
	"strconv"

	"github.com/nexucis/grafana-go-client/api/types"
)

const keyAPI = "/api/auth/keys"

type KeyInterface interface {
	Get() ([]*types.GetAPIKeyResponse, error)
	Create(*types.APIKeyForm) (*types.CreateAPIKeyResponse, error)
	Delete(int64) error
}

func newKey(client *grafanahttp.RESTClient) KeyInterface {
	return &key{
		client: client,
	}
}

type key struct {
	KeyInterface
	client *grafanahttp.RESTClient
}

func (c *key) Get() ([]*types.GetAPIKeyResponse, error) {
	var result []*types.GetAPIKeyResponse
	err := c.client.Get(keyAPI).
		Do().
		SaveAsObj(&result)
	return result, err
}

func (c *key) Create(key *types.APIKeyForm) (*types.CreateAPIKeyResponse, error) {
	result := &types.CreateAPIKeyResponse{}
	err := c.client.Post(keyAPI).
		Body(key).
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
