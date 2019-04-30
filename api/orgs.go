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

const orgsAPI = "/api/orgs"

type OrganisationsInterface interface {
	Create(string) (int64, error)
	Search(*QueryParameterOrgs) ([]*types.SimpleOrg, error)
	GetByID(int64) (*types.Org, error)
	GetByName(string) (*types.Org, error)
	Delete(int64) error
	Update(int64, string) error
	UpdateAddress(int64, *types.Address) error
	GetUsers(int64) ([]*types.OrgUser, error)
	AddUser(int64, *types.AddOrgUser) error
	UpdateUser(int64, int64, *types.UpdateOrgUser) error
	DeleteUser(int64, int64) error
	GetQuotas(int64) (*types.OrgQuota, error)
	UpdateQuotas(int64, string, int64) error
}

func newOrgs(client *http.RESTClient) OrganisationsInterface {
	return &orgs{
		client: client,
	}
}

type orgs struct {
	OrganisationsInterface
	client *http.RESTClient
}

func (c *orgs) Create(name string) (int64, error) {
	body := struct {
		Name string `json:"name" binding:"Required"`
	}{
		Name: name,
	}

	result := &struct {
		OrgID int64 `json:"orgId"`
	}{}

	err := c.client.Post(orgsAPI).
		Body(body).
		Do().
		SaveAsObj(result)

	return result.OrgID, err
}

func (c *orgs) Search(query *QueryParameterOrgs) ([]*types.SimpleOrg, error) {
	var response []*types.SimpleOrg
	request := c.client.Get(orgsAPI)

	setQueryParamOrgs(request, query)

	err := request.Do().
		SaveAsObj(&response)

	return response, err
}

func (c *orgs) GetByID(orgID int64) (*types.Org, error) {
	result := &types.Org{}
	err := c.client.Get(orgsAPI).
		SetSubPath("/:orgId").
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *orgs) GetByName(name string) (*types.Org, error) {
	result := &types.Org{}
	err := c.client.Get(orgsAPI).
		SetSubPath("/name/:name").
		SetPathParam("name", name).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *orgs) Delete(orgID int64) error {
	return c.client.Delete(orgsAPI).
		SetSubPath("/:orgId").
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		Do().
		Error()
}

func (c *orgs) Update(orgID int64, name string) error {
	entity := &struct {
		Name string `json:"name" binding:"Required"`
	}{Name: name}

	return c.client.Put(orgsAPI).
		SetSubPath("/:orgId").
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		Body(entity).
		Do().
		Error()
}

func (c *orgs) UpdateAddress(orgID int64, address *types.Address) error {
	return c.client.Put(orgsAPI).
		SetSubPath("/address").
		SetSubPath("/:orgId").
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		Body(address).
		Do().
		Error()
}

func (c *orgs) GetUsers(orgID int64) ([]*types.OrgUser, error) {
	var result []*types.OrgUser
	err := c.client.Get(orgsAPI).
		SetSubPath("/:orgId/users").
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		Do().
		SaveAsObj(&result)

	return result, err
}

func (c *orgs) AddUser(orgID int64, user *types.AddOrgUser) error {
	return c.client.Put(orgsAPI).
		SetSubPath("/:orgId/users").
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		Body(user).
		Do().
		Error()
}

func (c *orgs) UpdateUser(orgID int64, userID int64, user *types.UpdateOrgUser) error {
	return c.client.Patch(orgsAPI).
		SetSubPath("/:orgId/users/:userId").
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		SetPathParam("userId", strconv.FormatInt(userID, 10)).
		Body(user).
		Do().
		Error()
}

func (c *orgs) DeleteUser(orgID int64, userID int64) error {
	return c.client.Delete(orgsAPI).
		SetSubPath("/:orgId/users/:userId").
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		SetPathParam("userId", strconv.FormatInt(userID, 10)).
		Do().
		Error()
}

func (c *orgs) GetQuotas(orgID int64) (*types.OrgQuota, error) {
	result := &types.OrgQuota{}
	err := c.client.Get(orgsAPI).
		SetSubPath("/:orgId/quotas").
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *orgs) UpdateQuotas(orgID int64, target string, limit int64) error {
	entity := &struct {
		Limit int64 `json:"limit"`
	}{Limit: limit}

	return c.client.Put(orgsAPI).
		SetSubPath("/:orgId/quotas/:target").
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		SetPathParam("target", target).
		Body(entity).
		Do().
		Error()
}

func setQueryParamOrgs(request *http.Request, query *QueryParameterOrgs) {

	if len(query.query) > 0 {
		request.AddQueryParam("query", query.query)
	}

	if len(query.name) > 0 {
		request.AddQueryParam("name", query.name)
	}
}
