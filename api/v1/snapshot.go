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
	"github.com/nexucis/grafana-go-client/api/v1/types"
	"github.com/nexucis/grafana-go-client/http"
)

const snapshotAPI = "/api/snapshots"

type SnapshotInterface interface {
	Create(*types.DashboardSnapshot) (*types.CreateSnaphostResponse, error)
	Get()
	GetSharingOptions() (*types.SharedOptionSnaphost, error)
	GetByKey(string) (*types.DashboardWithMeta, error)
	DeleteByKey(string) error
	DeleteByDeleteKey(string) error
}

func newSnapshot(client *http.RESTClient) SnapshotInterface {
	return &snapshot{
		client: client,
	}
}

type snapshot struct {
	SnapshotInterface
	client *http.RESTClient
}

func (c *snapshot) Create(snapshot *types.DashboardSnapshot) (*types.CreateSnaphostResponse, error) {
	result := &types.CreateSnaphostResponse{}
	err := c.client.Post(snapshotAPI).
		Body(snapshot).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *snapshot) Get() {
	c.client.Get("/api/dashboard/snapshot")
}

func (c *snapshot) GetSharingOptions() (*types.SharedOptionSnaphost, error) {
	result := &types.SharedOptionSnaphost{}
	err := c.client.Get(snapshotAPI).
		SetSubPath("/shared-options").
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *snapshot) GetByKey(key string) (*types.DashboardWithMeta, error) {
	result := &types.DashboardWithMeta{}
	err := c.client.Get(snapshotAPI).
		SetSubPath("/:key").
		SetPathParam("key", key).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *snapshot) DeleteByKey(key string) error {
	return c.client.Delete(snapshotAPI).
		SetSubPath("/:key").
		SetPathParam("key", key).
		Do().
		Error()
}

func (c *snapshot) DeleteByDeleteKey(deleteKey string) error {
	return c.client.Get("/api/snapshots-delete").
		SetSubPath("/:deleteKey").
		SetPathParam("deleteKey", deleteKey).
		Do().
		Error()
}
