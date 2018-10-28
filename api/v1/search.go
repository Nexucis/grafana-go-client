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
	"strconv"

	"github.com/nexucis/grafana-go-client/api/v1/types"
	"github.com/nexucis/grafana-go-client/http"
)

const searchAPI = "/api/search"

type SearchInterface interface {
	Query(*QueryParameterSearch) ([]*types.SearchResult, error)
}

func newSearch(client *http.RESTClient) SearchInterface {
	return &search{
		client: client,
	}
}

type search struct {
	SearchInterface
	client *http.RESTClient
}

func (c *search) Query(query *QueryParameterSearch) ([]*types.SearchResult, error) {
	var response []*types.SearchResult
	request := c.client.Get(searchAPI)

	setQueryParamSearch(request, query)

	err := request.Do().
		SaveAsObj(&response)

	return response, err
}

func setQueryParamSearch(request *http.Request, query *QueryParameterSearch) {

	if len(query.query) > 0 {
		request.AddQueryParam("query", query.query)
	}

	if query.tags != nil {
		for _, tag := range query.tags {
			request.AddQueryParam("tag", string(tag))
		}
	}

	if len(query.searchType) > 0 {
		request.AddQueryParam("type", string(query.searchType))
	}

	if query.dashboardIds != nil {
		for _, id := range query.dashboardIds {
			request.AddQueryParam("dashboardIds", strconv.FormatInt(id, 10))
		}
	}

	if query.folderIds != nil {
		for _, id := range query.folderIds {
			request.AddQueryParam("folderIds", strconv.FormatInt(id, 10))
		}
	}

	request.AddQueryParam("starred", strconv.FormatBool(query.starred))

	if query.limit > 0 {
		request.AddQueryParam("limit", strconv.Itoa(query.limit))
	}

	if len(query.permission) > 0 {
		request.AddQueryParam("permission", string(query.permission))
	}
}
