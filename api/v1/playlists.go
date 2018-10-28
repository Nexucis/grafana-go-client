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
	"github.com/nexucis/grafana-go-client/api/v1/types"
	"github.com/nexucis/grafana-go-client/http"

	"strconv"
)

const playlistAPI = "/api/playlists"

type PlaylistInterface interface {
	Search(*QueryParameterPlaylist) ([]*types.SimplePlaylist, error)
	Create(*types.Playlist) (*types.SimplePlaylist, error)
	GetByID(int64) (*types.Playlist, error)
	GetItems(int64) ([]*types.PlaylistItem, error)
	GetDashboards(int64) ([]*types.PlaylistDashboard, error)
	Delete(int64) error
	Update(int64, *types.Playlist) (*types.Playlist, error)
}

func newPlaylist(client *http.RESTClient) PlaylistInterface {
	return &playlist{
		client: client,
	}
}

type playlist struct {
	PlaylistInterface
	client *http.RESTClient
}

func (c *playlist) Search(query *QueryParameterPlaylist) ([]*types.SimplePlaylist, error) {
	var result []*types.SimplePlaylist
	request := c.client.Get(playlistAPI)

	setQueryParamPlaylist(request, query)

	err := request.Do().
		SaveAsObj(&result)

	return result, err
}

func (c *playlist) Create(playlist *types.Playlist) (*types.SimplePlaylist, error) {
	result := &types.SimplePlaylist{}
	err := c.client.Post(playlistAPI).
		Body(playlist).
		Do().
		SaveAsObj(result)

	return result, err
}

func (c *playlist) GetByID(playlistID int64) (*types.Playlist, error) {
	result := &types.Playlist{}
	err := c.client.Get(playlistAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(playlistID, 10)).
		Do().
		SaveAsObj(result)

	return result, err
}

func (c *playlist) GetItems(playlistID int64) ([]*types.PlaylistItem, error) {
	var result []*types.PlaylistItem
	err := c.client.Get(playlistAPI).
		SetSubPath("/:id/items").
		SetPathParam("id", strconv.FormatInt(playlistID, 10)).
		Do().
		SaveAsObj(&result)

	return result, err
}

func (c *playlist) GetDashboards(playlistID int64) ([]*types.PlaylistDashboard, error) {
	var result []*types.PlaylistDashboard
	err := c.client.Get(playlistAPI).
		SetSubPath("/:id/dashboards").
		SetPathParam("id", strconv.FormatInt(playlistID, 10)).
		Do().
		SaveAsObj(&result)

	return result, err
}

func (c *playlist) Delete(playlistID int64) error {
	return c.client.Delete(playlistAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(playlistID, 10)).
		Do().
		Error()
}

func (c *playlist) Update(playlistID int64, playlist *types.Playlist) (*types.Playlist, error) {
	result := &types.Playlist{}
	err := c.client.Put(playlistAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(playlistID, 10)).
		Body(playlist).
		Do().
		SaveAsObj(result)

	return result, err
}

func setQueryParamPlaylist(request *http.Request, query *QueryParameterPlaylist) {

	if len(query.query) > 0 {
		request.AddQueryParam("query", query.query)
	}

	if query.limit > 0 {
		request.AddQueryParam("limit", strconv.FormatInt(query.limit, 10))
	}
}
