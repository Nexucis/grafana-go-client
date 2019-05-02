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
	"github.com/nexucis/grafana-go-client/grafanahttp"
	"strconv"
)

const alertNotificationAPI = "/api/alert-notifications"

type AlertNotificationInterface interface {
	Create(*types.CreateAlertNotification) (*types.ResponseAlertNotification, error)
	CreateTest(*types.CreateTestAlertNotification) error
	Update(int64, *types.UpdateAlertNotification) (*types.ResponseAlertNotification, error)
	Get() ([]*types.ResponseAlertNotification, error)
	GetByID(int64) (*types.ResponseAlertNotification, error)
	Delete(int64) error
}

func newAlertNotification(client *grafanahttp.RESTClient) AlertNotificationInterface {
	return &alertNotification{
		client: client,
	}
}

type alertNotification struct {
	AlertNotificationInterface
	client *grafanahttp.RESTClient
}

func (c *alertNotification) Create(body *types.CreateAlertNotification) (*types.ResponseAlertNotification, error) {
	response := &types.ResponseAlertNotification{}
	err := c.client.Post(alertNotificationAPI).
		Body(body).
		Do().
		SaveAsObj(response)
	return response, err
}

func (c *alertNotification) CreateTest(body *types.CreateTestAlertNotification) error {
	return c.client.Post(alertNotificationAPI).
		SetSubPath("/test").
		Body(body).
		Do().
		Error()
}

func (c *alertNotification) Update(notificationID int64, body *types.UpdateAlertNotification) (*types.ResponseAlertNotification, error) {
	response := &types.ResponseAlertNotification{}
	err := c.client.Put(alertNotificationAPI).
		SetSubPath("/:notificationId").
		SetPathParam("notificationId", strconv.FormatInt(notificationID, 10)).
		Body(body).
		Do().
		SaveAsObj(response)
	return response, err
}

func (c *alertNotification) Get() ([]*types.ResponseAlertNotification, error) {
	var result []*types.ResponseAlertNotification
	err := c.client.Get(alertNotificationAPI).
		Do().
		SaveAsObj(&result)

	return result, err
}

func (c *alertNotification) GetByID(notificationID int64) (*types.ResponseAlertNotification, error) {
	response := &types.ResponseAlertNotification{}
	err := c.client.Get(alertNotificationAPI).
		SetSubPath("/:notificationId").
		SetPathParam("notificationId", strconv.FormatInt(notificationID, 10)).
		Do().
		SaveAsObj(response)
	return response, err
}

func (c *alertNotification) Delete(notificationID int64) error {
	return c.client.Delete(alertNotificationAPI).
		SetSubPath("/:notificationId").
		SetPathParam("notificationId", strconv.FormatInt(notificationID, 10)).
		Do().
		Error()
}
