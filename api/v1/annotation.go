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

const annotationAPI = "/api/annotations"

type AnnotationInterface interface {
	Create(*PostAnnotations) (*ResponseCreateAnnotation, error)
	CreateGraphite(*PostGraphiteAnnotations) (*ResponseCreateGraphiteAnnotation, error)
	Update(*UpdateAnnotations) error
	Delete(id int64) error
	MassiveDelete(*DeleteAnnotations) error
	Get(query *QueryParamAnnotation) (*ResponseGetAnnotation, error)
}

func newAnnotation(client *http.RESTClient) AnnotationInterface {
	return &annotation{
		client: client,
	}
}

type annotation struct {
	AnnotationInterface
	client *http.RESTClient
}

func (c *annotation) Create(annotations *PostAnnotations) (*ResponseCreateAnnotation, error) {
	response := &ResponseCreateAnnotation{}
	err := c.client.Post(annotationAPI).
		Body(annotations).
		Do().
		SaveAsObj(response)

	return response, err
}

func (c *annotation) CreateGraphite(annotations *PostGraphiteAnnotations) (*ResponseCreateGraphiteAnnotation, error) {
	response := &ResponseCreateGraphiteAnnotation{}
	err := c.client.Post(annotationAPI).
		Body(annotations).
		SetSubPath("/graphite").
		Do().
		SaveAsObj(response)

	return response, err
}

func (c *annotation) Update(annotations *UpdateAnnotations) error {
	return c.client.Put(annotationAPI).
		SetSubPath("/:id").
		SetPathParam("id", strconv.FormatInt(annotations.Id, 10)).
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

func (c *annotation) MassiveDelete(annotations DeleteAnnotations) error {
	return c.client.Post(annotationAPI).
		SetSubPath("/mass-delete").
		Body(annotations).
		Do().
		Error()
}

func (c *annotation) Get(queryParam *QueryParamAnnotation) (*ResponseGetAnnotation, error) {
	response := &ResponseGetAnnotation{}
	request := c.client.Post(annotationAPI).
		SetSubPath("/graphite")

	setQueryParamAnnotation(request, queryParam)

	err := request.Do().
		SaveAsObj(response)

	return response, err
}

func setQueryParamAnnotation(request *http.Request, query *QueryParamAnnotation) {
	if query.from > 0 {
		request.AddQueryParam("from", strconv.FormatInt(query.from, 10))
	}

	if query.to > 0 {
		request.AddQueryParam("to", strconv.FormatInt(query.from, 10))
	}

	if query.userId > 0 {
		request.AddQueryParam("userId", strconv.FormatInt(query.userId, 10))
	}

	if query.alertId > 0 {
		request.AddQueryParam("alertId", strconv.FormatInt(query.alertId, 10))
	}

	if query.dashboardId > 0 {
		request.AddQueryParam("dashboardId", strconv.FormatInt(query.dashboardId, 10))
	}

	if query.panelId > 0 {
		request.AddQueryParam("panelId", strconv.FormatInt(query.panelId, 10))
	}

	if query.limit > 0 {
		request.AddQueryParam("limit", strconv.FormatInt(query.limit, 10))
	}

	if len(query._type) > 0 {
		request.AddQueryParam("type", query._type)
	}

	if query.tags != nil {
		for _, tag := range query.tags {
			request.AddQueryParam("tags", tag)
		}
	}

}
