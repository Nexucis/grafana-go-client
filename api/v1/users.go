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

const usersAPI = "/api/users"

type UsersInterface interface {
	Get(*QueryParameterUsers) ([]*types.UserSearchHit, error)
	GetWithPaging(*QueryParameterUsers) (*types.UserSearchWithPaging, error)
	GetByID(int64) (*types.UserProfile, error)
	GetByLoginOrEmail(string) (*types.UserProfile, error)
	GetOrgs(int64) (*types.UserOrgList, error)
	Update(int64, *types.UpdateCurrentUser) error
	UpdateUserActiveOrg(int64, int64) error
}

func newUsers(client *http.RESTClient) UsersInterface {
	return &users{
		client: client,
	}
}

type users struct {
	UsersInterface
	client *http.RESTClient
}

func (c *users) Get(query *QueryParameterUsers) ([]*types.UserSearchHit, error) {
	var response []*types.UserSearchHit
	request := c.client.Get(usersAPI)

	setQueryParamUsers(request, query)

	err := request.Do().
		SaveAsObj(response)

	return response, err
}

func (c *users) GetWithPaging(query *QueryParameterUsers) (*types.UserSearchWithPaging, error) {
	var response *types.UserSearchWithPaging
	request := c.client.Get(usersAPI).SetSubPath("/search")

	setQueryParamUsers(request, query)

	err := request.Do().
		SaveAsObj(response)

	return response, err
}

func (c *users) GetByID(userID int64) (*types.UserProfile, error) {
	result := &types.UserProfile{}
	err := c.client.Get(usersAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(userID, 10)).
		Do().
		SaveAsObj(result)

	return result, err
}

func (c *users) GetByLoginOrEmail(loginOrEmail string) (*types.UserProfile, error) {
	result := &types.UserProfile{}
	err := c.client.Get(usersAPI).
		SetSubPath("/lookup").
		AddQueryParam("loginOrEmail", loginOrEmail).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *users) GetOrgs(userID int64) (*types.UserOrgList, error) {
	result := &types.UserOrgList{}
	err := c.client.Get(usersAPI).
		SetSubPath("/:d/orgs").
		SetPathParam("id", strconv.FormatInt(userID, 10)).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *users) Update(userID int64, user *types.UpdateCurrentUser) error {
	return c.client.Put(usersAPI).
		SetSubPath("/:d").
		SetPathParam("id", strconv.FormatInt(userID, 10)).
		Body(user).
		Do().
		Error()
}

func (c *users) UpdateUserActiveOrg(userID int64, orgID int64) error {
	return c.client.Post(usersAPI).
		SetSubPath("/:id/using/:orgId").
		SetPathParam("id", strconv.FormatInt(userID, 10)).
		SetPathParam("orgId", strconv.FormatInt(orgID, 10)).
		Do().
		Error()
}

func setQueryParamUsers(request *http.Request, query *QueryParameterUsers) {
	if query.page > 0 {
		request.AddQueryParam("page", strconv.FormatInt(query.page, 10))
	}

	if query.perPage > 0 {
		request.AddQueryParam("perpage", strconv.FormatInt(query.page, 10))
	}

	if len(query.query) > 0 {
		request.AddQueryParam("query", query.query)
	}
}
