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

const currentOrgAPI = "/api/org"

type CurrentOrgInterface interface {
	Get() (*types.Org, error)
	Update(string) error
	GetQuotas() (*types.OrgQuota, error)
	UpdateAddress(*types.Address) error
	AddUser(*types.AddOrgUser) error
	GetUsers() ([]*types.OrgUser, error)
	UpdateUser(int64, *types.UpdateOrgUser) error
	DeleteUser(int64) error
	GetInvites() ([]*types.TempOrgUser, error)
	AddInvite(*types.AddInvite) error
	RevokeInvite(int64) error
	GetPreferences() (*types.OrgPrefs, error)
	UpdatePreferences(*types.OrgPrefs) error
}

func newCurrentOrg(client *http.RESTClient) CurrentOrgInterface {
	return &currentOrg{
		client: client,
	}
}

type currentOrg struct {
	CurrentOrgInterface
	client *http.RESTClient
}

func (c *currentOrg) Get() (*types.Org, error) {
	result := &types.Org{}
	err := c.client.Get(currentOrgAPI).
		Do().
		SaveAsObj(result)

	return result, err
}

func (c *currentOrg) Update(name string) error {
	entity := struct {
		Name string `json:"name" binding:"Required"`
	}{
		Name: name,
	}
	return c.client.Put(currentOrgAPI).
		Body(entity).
		Do().
		Error()
}

func (c *currentOrg) GetQuotas() (*types.OrgQuota, error) {
	result := &types.OrgQuota{}
	err := c.client.Get(currentOrgAPI).
		SetSubPath("/quotas").
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *currentOrg) UpdateAddress(address *types.Address) error {
	return c.client.Put(currentOrgAPI).
		SetSubPath("/address").
		Body(address).
		Do().
		Error()
}

func (c *currentOrg) AddUser(user *types.AddOrgUser) error {
	return c.client.Put(currentOrgAPI).
		SetSubPath("/users").
		Body(user).
		Do().
		Error()
}

func (c *currentOrg) GetUsers() ([]*types.OrgUser, error) {
	var result []*types.OrgUser
	err := c.client.Get(currentOrgAPI).
		SetSubPath("/users").
		Do().
		SaveAsObj(&result)

	return result, err
}

func (c *currentOrg) UpdateUser(userID int64, user *types.UpdateOrgUser) error {
	return c.client.Patch(currentOrgAPI).
		SetSubPath("/users/:userId").
		SetPathParam("userId", strconv.FormatInt(userID, 10)).
		Body(user).
		Do().
		Error()
}

func (c *currentOrg) DeleteUser(userID int64) error {
	return c.client.Delete(currentOrgAPI).
		SetSubPath("/users/:userId").
		SetPathParam("userId", strconv.FormatInt(userID, 10)).
		Do().
		Error()
}

func (c *currentOrg) GetInvites() ([]*types.TempOrgUser, error) {
	var result []*types.TempOrgUser
	err := c.client.Get(currentOrgAPI).
		SetSubPath("/invites").
		Do().
		SaveAsObj(&result)

	return result, err
}

func (c *currentOrg) AddInvite(invite *types.AddInvite) error {
	return c.client.Post(currentOrgAPI).
		SetSubPath("/invites").
		Body(invite).
		Do().
		Error()
}

func (c *currentOrg) RevokeInvite(code int64) error {
	return c.client.Patch(currentOrgAPI).
		SetSubPath("/invites/:code/revoke").
		SetPathParam("code", strconv.FormatInt(code, 10)).
		Do().
		Error()
}

func (c *currentOrg) GetPreferences() (*types.OrgPrefs, error) {
	result := &types.OrgPrefs{}
	err := c.client.Get(currentOrgAPI).
		SetSubPath("/preferences").
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *currentOrg) UpdatePreferences(pref *types.OrgPrefs) error {
	return c.client.Put(currentOrgAPI).
		SetSubPath("/preferences").
		Body(pref).
		Do().
		Error()
}
