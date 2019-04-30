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
	"github.com/nexucis/grafana-go-client/api/types"
	"github.com/nexucis/grafana-go-client/http"

	"strconv"
)

const alertAPI = "/api/alerts"

type AlertInterface interface {
	Get(*QueryParamAlert) ([]*types.ResponseGetAlert, error)
	GetByID(int64) (*types.ResponseGetAlert, error)
	GetStatesForDashboard(int64) (*types.ResponseGetStatesForDashboard, error)
	CreateTest(*types.PostAlertTest) (*types.ResponsePostAlertTest, error)
	PauseAlert(int64, bool) error
}

func newAlert(client *http.RESTClient) AlertInterface {
	return &alert{
		client: client,
	}
}

type alert struct {
	AlertInterface
	client *http.RESTClient
}

func (c *alert) Get(query *QueryParamAlert) ([]*types.ResponseGetAlert, error) {
	var response []*types.ResponseGetAlert
	request := c.client.Get(alertAPI)

	setQueryParamAlert(request, query)

	err := request.Do().
		SaveAsObj(&response)

	return response, err
}

func (c *alert) GetByID(alertID int64) (*types.ResponseGetAlert, error) {
	response := &types.ResponseGetAlert{}
	err := c.client.Get(alertAPI).
		SetSubPath("/:alertId").
		SetPathParam("alertId", strconv.FormatInt(alertID, 10)).
		Do().
		SaveAsObj(response)

	return response, err
}

func (c *alert) GetStatesForDashboard(alertID int64) (*types.ResponseGetStatesForDashboard, error) {
	response := &types.ResponseGetStatesForDashboard{}
	err := c.client.Get(alertAPI).
		SetSubPath("/states-for-dashboard").
		Do().
		SaveAsObj(response)
	return response, err
}

func (c *alert) CreateTest(alertTest *types.PostAlertTest) (*types.ResponsePostAlertTest, error) {
	response := &types.ResponsePostAlertTest{}
	err := c.client.Post(alertAPI).
		SetSubPath("/test").
		Body(alertTest).
		Do().
		SaveAsObj(response)
	return response, err
}

func (c *alert) PauseAlert(alertID int64, paused bool) error {
	body := &types.PostPauseAlert{AlertId: alertID, Paused: paused}

	return c.client.Post(alertAPI).
		SetSubPath("/:alertId/pause").
		SetPathParam("alertId", strconv.FormatInt(alertID, 10)).
		Body(body).
		Do().
		Error()
}

func setQueryParamAlert(request *http.Request, query *QueryParamAlert) {
	if query.panelId > 0 {
		request.AddQueryParam("panelId", strconv.FormatInt(query.panelId, 10))
	}

	if len(query.dashboardQuery) > 0 {
		request.AddQueryParam("dashboardQuery", query.dashboardQuery)
	}

	if len(query.query) > 0 {
		request.AddQueryParam("query", query.query)
	}

	if query.limit > 0 {
		request.AddQueryParam("limit", strconv.FormatInt(query.limit, 10))
	}

	if query.states != nil {
		for _, state := range query.states {
			request.AddQueryParam("state", string(state))
		}
	}

	if query.dashboardIds != nil {
		for _, dashboardId := range query.dashboardIds {
			request.AddQueryParam("dashboardId", strconv.FormatInt(dashboardId, 10))
		}
	}

	if query.folderIds != nil {
		for _, folderId := range query.folderIds {
			request.AddQueryParam("folderId", strconv.FormatInt(folderId, 10))
		}
	}

	if query.states != nil {
		for _, dashboardTag := range query.dashboardTags {
			request.AddQueryParam("dashboardTag", dashboardTag)
		}
	}
}
