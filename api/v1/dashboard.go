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
	"github.com/nexucis/grafana-go-client/api/v1/types"
	"strconv"
)

const dashboardAPI = "/api/dashboards"

type DashboardInterface interface {
	GetByUID(string)
	DeleteByUID(string) error
	// GetBySlug is deprecated since Grafana 5.0, please use GetByUID instead
	GetBySlug(string)
	// DeleteBySlug is deprecated since Grafana 5.0, please use DeleteByUID instead
	DeleteBySlug(string) error
	CalculateDiff()
	Create()
	GetHome()
	GetTags() ([]*types.DashboardTags, error)
	Import()
	GetVersion(int64) ([]*types.DashboardVersion, error)
	GetVersionByID(int64, int) (*types.DashboardVersionMeta, error)
	RestoreVersion(int64, int) (*types.SimpleDashboard, error)
	GetPermissions(int64) ([]*types.FolderOrDashboardPermission, error)
	UpdatePermissions(int64, items []*types.DashboardAclUpdateItem) error
}

func newDashboard(client *http.RESTClient) DashboardInterface {
	return &dashboard{
		client: client,
	}
}

type dashboard struct {
	DashboardInterface
	client *http.RESTClient
}

func (c *dashboard) GetByUID(string) {
	// not yet implemented, need to specify the dashboard struct
}

func (c *dashboard) DeleteByUID(uid string) error {
	return c.client.Delete(dashboardAPI).
		SetSubPath("/uid/:uid").
		SetPathParam("uid", uid).
		Do().
		Error()
}

// GetBySlug is deprecated since Grafana 5.0, please use GetByUID instead
func (c *dashboard) GetBySlug(string) {
	// not yet implemented, need to specify the dashboard struct
}

// DeleteBySlug is deprecated since Grafana 5.0, please use DeleteByUID instead
func (c *dashboard) DeleteBySlug(slug string) error {
	return c.client.Delete(dashboardAPI).
		SetSubPath("/db/:slug").
		SetPathParam("slug", slug).
		Do().
		Error()
}

func (c *dashboard) CalculateDiff() {

}

func (c *dashboard) Create() {
	// not yet implemented, need to specify the dashboard struct
}

func (c *dashboard) GetHome() {
	// not yet implemented, need to specify the dashboard struct
}

func (c *dashboard) GetTags() ([]*types.DashboardTags, error) {
	var result []*types.DashboardTags
	err := c.client.Get(dashboardAPI).
		SetSubPath("/tags").
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *dashboard) Import() {

}

func (c *dashboard) GetVersion(dashboardID int64) ([]*types.DashboardVersion, error) {
	var result []*types.DashboardVersion
	err := c.client.Get(dashboardAPI).
		SetSubPath("/id/:dashboardId/versions").
		SetPathParam("dashboardId", strconv.FormatInt(dashboardID, 10)).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *dashboard) GetVersionByID(dashboardID int64, versionID int) (*types.DashboardVersionMeta, error) {
	result := &types.DashboardVersionMeta{}
	err := c.client.Get(dashboardAPI).
		SetSubPath("/id/:dashboardId/versions/:id").
		SetPathParam("dashboardId", strconv.FormatInt(dashboardID, 10)).
		SetPathParam("id", strconv.Itoa(versionID)).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *dashboard) RestoreVersion(dashboardID int64, versionID int) (*types.SimpleDashboard, error) {
	body := struct {
		Version int `json:"version" binding:"Required"`
	}{
		Version: versionID,
	}
	result := &types.SimpleDashboard{}
	err := c.client.Post(dashboardAPI).
		SetSubPath("/id/:dashboardId/versions/restore").
		SetPathParam("dashboardId", strconv.FormatInt(dashboardID, 10)).
		Body(body).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *dashboard) GetPermissions(dashboardID int64) ([]*types.FolderOrDashboardPermission, error) {
	var result []*types.FolderOrDashboardPermission
	err := c.client.Get(dashboardAPI).
		SetSubPath("/id/:dashboardId/permissions").
		SetPathParam("dashboardId", strconv.FormatInt(dashboardID, 10)).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *dashboard) UpdatePermissions(dashboardID int64, items []*types.DashboardAclUpdateItem) error {
	body := struct {
		Items []*types.DashboardAclUpdateItem `json:"items"`
	}{Items: items}

	return c.client.Post(dashboardAPI).
		SetSubPath("/id/:dashboardId/permissions").
		SetPathParam("dashboardId", strconv.FormatInt(dashboardID, 10)).
		Body(body).
		Do().
		Error()
}
