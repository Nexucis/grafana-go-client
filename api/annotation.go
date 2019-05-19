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
	"strconv"

	"github.com/nexucis/grafana-go-client/api/types"
	"github.com/nexucis/grafana-go-client/grafanahttp"
)

const annotationAPI = "/api/annotations"

type AnnotationInterface interface {
	Create(*types.PostAnnotations) (*types.ResponseCreateAnnotation, error)
	CreateGraphite(*types.PostGraphiteAnnotations) (*types.ResponseCreateGraphiteAnnotation, error)
	Update(*types.UpdateAnnotations) error
	Delete(id int64) error
	MassiveDelete(*types.DeleteAnnotations) error
	Get(query QueryParamAnnotation) (*types.ResponseGetAnnotation, error)
}

func newAnnotation(client *grafanahttp.RESTClient) AnnotationInterface {
	return &annotation{
		client: client,
	}
}

type annotation struct {
	AnnotationInterface
	client *grafanahttp.RESTClient
}

func (c *annotation) Create(annotations *types.PostAnnotations) (*types.ResponseCreateAnnotation, error) {
	response := &types.ResponseCreateAnnotation{}
	err := c.client.Post(annotationAPI).
		Body(annotations).
		Do().
		SaveAsObj(response)

	return response, err
}

func (c *annotation) CreateGraphite(annotations *types.PostGraphiteAnnotations) (*types.ResponseCreateGraphiteAnnotation, error) {
	response := &types.ResponseCreateGraphiteAnnotation{}
	err := c.client.Post(annotationAPI).
		Body(annotations).
		SetSubPath("/graphite").
		Do().
		SaveAsObj(response)

	return response, err
}

func (c *annotation) Update(annotations *types.UpdateAnnotations) error {
	return c.client.Put(annotationAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(annotations.ID, 10)).
		Body(annotations).
		Do().
		Error()
}

func (c *annotation) Delete(id int64) error {
	return c.client.Delete(annotationAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(id, 10)).
		Do().
		Error()
}

func (c *annotation) MassiveDelete(annotations *types.DeleteAnnotations) error {
	return c.client.Post(annotationAPI).
		SetSubPath("/mass-delete").
		Body(annotations).
		Do().
		Error()
}

func (c *annotation) Get(queryParam QueryParamAnnotation) (*types.ResponseGetAnnotation, error) {
	response := &types.ResponseGetAnnotation{}
	err := c.client.Get(annotationAPI).
		Query(&queryParam).
		Do().
		SaveAsObj(response)

	return response, err
}
