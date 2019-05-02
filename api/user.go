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

const currentUserAPI = "/api/user"

type CurrentUserInterface interface {
	Get() (*types.UserProfile, error)
	Update(*types.UpdateCurrentUser) error
	ChangeActiveOrganization(int64) error
	GetOrg() (*types.UserOrgList, error)
	StarDashboard(int64) error
	UnstarDashboard(int64) error
	UpdatePassword(string, string) error
	GetQuotas() (*types.UserQuota, error)
	AddHelpFlags(int64) error
	ClearHelpFlags() error
	GetPreference() (*types.UserPreference, error)
	UpdatePreference(*types.UserPreference) error
}

func newCurrentUser(client *grafanahttp.RESTClient) CurrentUserInterface {
	return &currentUser{
		client: client,
	}
}

type currentUser struct {
	CurrentUserInterface
	client *grafanahttp.RESTClient
}

func (c *currentUser) Get() (*types.UserProfile, error) {
	result := &types.UserProfile{}
	err := c.client.Get(currentUserAPI).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *currentUser) Update(user *types.UpdateCurrentUser) error {
	return c.client.Put(currentUserAPI).
		Body(user).
		Do().
		Error()
}

func (c *currentUser) ChangeActiveOrganization(organisationID int64) error {
	return c.client.Post(currentUserAPI).
		SetSubPath("/using/:organizationID").
		SetPathParam("organizationID", strconv.FormatInt(organisationID, 10)).
		Do().
		Error()
}

func (c *currentUser) GetOrg() (*types.UserOrgList, error) {
	result := &types.UserOrgList{}
	err := c.client.Get(currentUserAPI).
		SetSubPath("/orgs").
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *currentUser) StarDashboard(dashboardID int64) error {
	return c.client.Post(currentUserAPI).
		SetSubPath("/stars/dashboard/:dashboardID").
		SetPathParam("dashboardID", strconv.FormatInt(dashboardID, 10)).
		Do().
		Error()
}

func (c *currentUser) UnstarDashboard(dashboardID int64) error {
	return c.client.Delete(currentUserAPI).
		SetSubPath("/stars/dashboard/:dashboardID").
		SetPathParam("dashboardID", strconv.FormatInt(dashboardID, 10)).
		Do().
		Error()
}

func (c *currentUser) UpdatePassword(oldPassword string, newPassword string) error {
	body := &types.UpdatePassword{OldPassword: oldPassword, NewPassword: newPassword}
	return c.client.Put(currentUserAPI).
		SetSubPath("/password").
		Body(body).
		Do().
		Error()
}

func (c *currentUser) GetQuotas() (*types.UserQuota, error) {
	result := &types.UserQuota{}
	err := c.client.Get(currentUserAPI).
		SetSubPath("/quotas").
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *currentUser) AddHelpFlags(flagID int64) error {
	return c.client.Put(currentUserAPI).
		SetSubPath("/helpflags/:flagID").
		SetPathParam("flagID", strconv.FormatInt(flagID, 10)).
		Do().
		Error()
}

func (c *currentUser) ClearHelpFlags() error {
	return c.client.Get(currentUserAPI).
		SetSubPath("/helpflags/clear").
		Do().
		Error()
}

func (c *currentUser) GetPreference() (*types.UserPreference, error) {
	result := &types.UserPreference{}
	err := c.client.Get(currentUserAPI).
		SetSubPath("/preferences").
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *currentUser) UpdatePreference(preference *types.UserPreference) error {
	return c.client.Put(currentUserAPI).
		SetSubPath("/preferences").
		Body(preference).
		Do().
		Error()
}
