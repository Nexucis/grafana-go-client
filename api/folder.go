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
	"strconv"

	"github.com/nexucis/grafana-go-client/api/types"
	"github.com/nexucis/grafana-go-client/grafanahttp"
)

const folderAPI = "/api/folders"

type FolderInterface interface {
	Get(int) ([]*types.SimpleFolder, error)
	GetByID(id int64) (*types.Folder, error)
	GetByUID(string) (*types.Folder, error)
	Create(string, string) (*types.Folder, error)
	Update(string, *types.UpdateFolder) (*types.Folder, error)
	Delete(string) error
	GetPermissions(string) ([]*types.FolderOrDashboardPermission, error)
	UpdatePermissions(string, []*types.DashboardACLUpdateItem) error
}

func newFolder(client *grafanahttp.RESTClient) FolderInterface {
	return &folder{
		client: client,
	}
}

type folder struct {
	FolderInterface
	client *grafanahttp.RESTClient
}

func (c *folder) Get(limit int) ([]*types.SimpleFolder, error) {
	if limit == 0 {
		limit = 1000
	}

	var result []*types.SimpleFolder
	err := c.client.Get(folderAPI).
		AddQueryParam("limit", strconv.Itoa(limit)).
		Do().
		SaveAsObj(&result)
	return result, err
}

func (c *folder) GetByID(id int64) (*types.Folder, error) {
	result := &types.Folder{}
	err := c.client.Get(folderAPI).
		SetSubPath("/id/:id").
		SetPathParam("id", strconv.FormatInt(id, 10)).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *folder) GetByUID(uid string) (*types.Folder, error) {
	result := &types.Folder{}
	err := c.client.Get(folderAPI).
		SetSubPath("/:uid").
		SetPathParam("uid", uid).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *folder) Create(title string, uid string) (*types.Folder, error) {
	body := struct {
		Title string `json:"title"`
		UID   string `json:"uid"`
	}{Title: title, UID: uid}

	result := &types.Folder{}
	err := c.client.Post(folderAPI).
		Body(body).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *folder) Update(uid string, folder *types.UpdateFolder) (*types.Folder, error) {
	result := &types.Folder{}
	err := c.client.Put(folderAPI).
		SetSubPath("/:uid").
		SetPathParam("uid", uid).
		Body(folder).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *folder) Delete(uid string) error {
	return c.client.Delete(folderAPI).
		SetSubPath("/:uid").
		SetPathParam("uid", uid).
		Do().
		Error()
}

func (c *folder) GetPermissions(uid string) ([]*types.FolderOrDashboardPermission, error) {
	var result []*types.FolderOrDashboardPermission
	err := c.client.Get(folderAPI).
		SetSubPath("/:uid/permissions").
		SetPathParam("uid", uid).
		Do().
		SaveAsObj(&result)
	return result, err
}

func (c *folder) UpdatePermissions(uid string, items []*types.DashboardACLUpdateItem) error {
	body := struct {
		Items []*types.DashboardACLUpdateItem `json:"items"`
	}{Items: items}

	return c.client.Post(folderAPI).
		SetSubPath("/:uid/permissions").
		SetPathParam("uid", uid).
		Body(body).
		Do().
		Error()
}
