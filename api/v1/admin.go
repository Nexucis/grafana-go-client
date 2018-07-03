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

const adminAPI = "/api/admin"

type AdminInterface interface {
	GetSettings() (*AdminSettings, error)
	CreateUser(*AdminCreateUserForm) (*AdminCreateUserResponse, error)
	UpdateUserPassword(int64, string) error
	UpdateUserPermissions(int64, bool) error
	DeleteUser(int64) error
	GetUserQuotas(int64) (*UserQuotaResponse, error)
	UpdateUserQuotas(int64, string, *UpdateUserQuotaForm) error
	GetStats() (*AdminStats, error)
	PauseAllAlerts(bool) (*PauseAllAlertsResponse, error)
}

func newAdmin(client *http.RESTClient) AdminInterface {
	return &admin{
		client: client,
	}
}

type admin struct {
	AdminInterface
	client *http.RESTClient
}

func (c *admin) GetSettings() (*AdminSettings, error) {
	result := &AdminSettings{}
	err := c.client.Get(adminAPI).
		SetSubPath("/settings").
		Do().
		SaveAsObj(result)

	return result, err
}

func (c *admin) CreateUser(user *AdminCreateUserForm) (*AdminCreateUserResponse, error) {
	result := &AdminCreateUserResponse{}
	err := c.client.Post(adminAPI).
		SetSubPath("/users").
		Body(user).
		Do().
		SaveAsObj(result)

	return result, err
}

func (c *admin) UpdateUserPassword(id int64, password string) error {
	form := &AdminUpdateUserPasswordForm{Password: password}
	return c.client.Put(adminAPI).
		SetSubPath("/users/:id/password").
		SetPathParam("id", strconv.FormatInt(id, 10)).
		Body(form).
		Do().
		Error()
}

func (c *admin) UpdateUserPermissions(id int64, permission bool) error {
	perm := &AdminUpdateUserPermissionsForm{permission}
	return c.client.Put(adminAPI).
		SetSubPath("/users/:id/permissions").
		SetPathParam("id", strconv.FormatInt(id, 10)).
		Body(perm).
		Do().
		Error()
}

func (c *admin) DeleteUser(id int64) error {
	return c.client.Delete(adminAPI).
		SetSubPath("/users/:id").
		SetPathParam("id", strconv.FormatInt(id, 10)).
		Do().
		Error()
}

func (c *admin) GetUserQuotas(id int64) (*UserQuotaResponse, error) {
	result := &UserQuotaResponse{}
	err := c.client.Get(adminAPI).
		SetSubPath("/users/:id/quotas").
		SetPathParam("id", strconv.FormatInt(id, 10)).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *admin) UpdateUserQuotas(id int64, target string, quotas *UpdateUserQuotaForm) error {
	return c.client.Put(adminAPI).
		SetSubPath("/users/:id/quotas/:target").
		SetPathParam("id", strconv.FormatInt(id, 10)).
		SetPathParam("target", target).
		Body(quotas).
		Do().
		Error()
}

func (c *admin) GetStats() (*AdminStats, error) {
	result := &AdminStats{}
	err := c.client.Get(adminAPI).
		SetSubPath("/stats").
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *admin) PauseAllAlerts(paused bool) (*PauseAllAlertsResponse, error) {
	body := &PauseAllAlertsForm{Paused: paused}
	result := &PauseAllAlertsResponse{}
	err := c.client.Post(adminAPI).
		SetSubPath("/pause-all-alerts").
		Body(body).
		Do().
		SaveAsObj(result)
	return result, err
}
