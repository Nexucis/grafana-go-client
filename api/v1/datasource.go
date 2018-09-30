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

const datasourceAPI = "/api/datasources"

type DataSourceInterface interface {
	Get() ([]*types.DataSource, error)
	Create(*types.AddDataSource) (*types.WriteDataSourceResponse, error)
	Update(int64, types.UpdateDataSource) (*types.WriteDataSourceResponse, error)
	Delete(int64) error
	DeleteByName(string) error
	GetByID(int64) (*types.DataSource, error)
	GetByName(string) (*types.DataSource, error)
	GetIDByName(string) (int64, error)
}

func newDataSource(client *http.RESTClient) DataSourceInterface {
	return &dataSource{
		client: client,
	}
}

type dataSource struct {
	DataSourceInterface
	client *http.RESTClient
}

func (c *dataSource) Get() ([]*types.DataSource, error) {
	var result []*types.DataSource
	err := c.client.Get(datasourceAPI).
		Do().
		SaveAsObj(&result)
	return result, err
}

func (c *dataSource) Create(source *types.AddDataSource) (*types.WriteDataSourceResponse, error) {
	result := &types.WriteDataSourceResponse{}
	err := c.client.Post(datasourceAPI).
		Body(source).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *dataSource) Update(sourceID int64, source types.UpdateDataSource) (*types.WriteDataSourceResponse, error) {
	result := &types.WriteDataSourceResponse{}
	err := c.client.Put(datasourceAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(sourceID, 10)).
		Body(source).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *dataSource) Delete(sourceID int64) error {
	return c.client.Delete(datasourceAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(sourceID, 10)).
		Do().
		Error()
}

func (c *dataSource) DeleteByName(sourceName string) error {
	return c.client.Delete(datasourceAPI).
		SetSubPath("/name/:name").
		SetPathParam("name", sourceName).
		Do().
		Error()
}

func (c *dataSource) GetByID(sourceID int64) (*types.DataSource, error) {
	result := &types.DataSource{}
	err := c.client.Get(datasourceAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(sourceID, 10)).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *dataSource) GetByName(sourceName string) (*types.DataSource, error) {
	result := &types.DataSource{}
	err := c.client.Get(datasourceAPI).
		SetSubPath("/name/:name").
		SetPathParam("name", sourceName).
		Do().
		SaveAsObj(result)
	return result, err
}

func (c *dataSource) GetIDByName(sourceName string) (int64, error) {
	result := &struct {
		Id int64 `json:"id"`
	}{}
	err := c.client.Get(datasourceAPI).
		SetSubPath("/id/:name").
		SetPathParam("name", sourceName).
		Do().
		SaveAsObj(result)
	return result.Id, err
}
