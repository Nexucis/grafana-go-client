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

const teamAPI = "/api/teams"

type TeamInterface interface {
	Get(query QueryParameterTeams) (*types.SearchTeam, error)
	GetByID(teamID int64) (*types.Team, error)
	GetMembers(teamID int64) ([]*types.TeamMember, error)
	Create(team *types.CreateOrUpdateTeam) (int64, error)
	Update(teamID int64, team *types.CreateOrUpdateTeam) error
	AddMembers(teamID int64, userID int64) error
	Delete(teamID int64) error
	DeleteMembers(teamID int64, userID int64) error
}

func newTeam(client *grafanahttp.RESTClient) TeamInterface {
	return &team{
		client: client,
	}
}

type team struct {
	TeamInterface
	client *grafanahttp.RESTClient
}

func (c *team) Get(query QueryParameterTeams) (*types.SearchTeam, error) {
	response := &types.SearchTeam{}
	err := c.client.Get(teamAPI).
		SetSubPath("/search").
		Query(&query).
		Do().
		SaveAsObj(response)

	return response, err
}

func (c *team) GetByID(teamID int64) (*types.Team, error) {
	response := &types.Team{}
	err := c.client.Get(teamAPI).
		SetSubPath("/:teamId").
		SetPathParam("teamId", strconv.FormatInt(teamID, 10)).
		Do().
		SaveAsObj(response)
	return response, err
}

func (c *team) GetMembers(teamID int64) ([]*types.TeamMember, error) {
	var response []*types.TeamMember
	err := c.client.Get(teamAPI).
		SetSubPath("/:teamId/members").
		SetPathParam("teamId", strconv.FormatInt(teamID, 10)).
		Do().
		SaveAsObj(&response)
	return response, err
}

func (c *team) Create(team *types.CreateOrUpdateTeam) (int64, error) {
	response := &struct {
		TeamID int64 `json:"teamId"`
	}{}
	err := c.client.Post(teamAPI).
		Body(team).
		Do().
		SaveAsObj(response)
	return response.TeamID, err
}

func (c *team) Update(teamID int64, team *types.CreateOrUpdateTeam) error {
	return c.client.Put(teamAPI).
		SetSubPath("/:teamId").
		SetPathParam("teamId", strconv.FormatInt(teamID, 10)).
		Body(team).
		Do().
		Error()
}

func (c *team) AddMembers(teamID int64, userID int64) error {
	teamMember := &struct {
		UserId int64 `json:"userId" binding:"Required"`
	}{
		UserId: userID,
	}
	return c.client.Post(teamAPI).
		SetSubPath("/:teamId/members").
		SetPathParam("teamId", strconv.FormatInt(teamID, 10)).
		Body(teamMember).
		Do().
		Error()
}

func (c *team) Delete(teamID int64) error {
	return c.client.Delete(teamAPI).
		SetSubPath("/:teamId").
		SetPathParam("teamId", strconv.FormatInt(teamID, 10)).
		Do().
		Error()
}

func (c *team) DeleteMembers(teamID int64, userID int64) error {
	return c.client.Delete(teamAPI).
		SetSubPath("/:teamId/members/:userId").
		SetPathParam("teamId", strconv.FormatInt(teamID, 10)).
		SetPathParam("userId", strconv.FormatInt(userID, 10)).
		Do().
		Error()
}
