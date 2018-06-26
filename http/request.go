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

package http

import (
	"context"
	"net/url"
	"io"
	"encoding/json"
	"bytes"
	"net/http"
	"fmt"
	"regexp"
	"strings"
)

// Client is an interface for testing a request object.
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type Request struct {
	client Client
	method string
	token  string
	ctx    context.Context

	// all component relative to the url
	baseURL    *url.URL
	pathPrefix string // it's the api endpoint such as /api/annotations
	subpath    string // the sub path of the endpoint like /:id, it can be empty
	queryParam url.Values
	pathParam  map[string]string

	body io.Reader
	err  error
}

func NewRequest(client Client, method string, baseURL *url.URL, pathPrefix string, token string) *Request {
	return &Request{
		client:     client,
		method:     method,
		token:      token,
		baseURL:    baseURL,
		pathPrefix: pathPrefix,
	}
}

func (r *Request) AddQueryParam(queryName string, value string) *Request {
	if r.queryParam == nil {
		r.queryParam = make(url.Values)
	}
	r.queryParam[queryName] = append(r.queryParam[queryName], value)
	return r
}

func (r *Request) SetPathParam(pathName string, value string) *Request {
	if r.pathParam == nil {
		r.pathParam = make(map[string]string)
	}
	r.pathParam[pathName] = value
	return r
}

func (r *Request) SetSubPath(subPath string) *Request {
	r.subpath = subPath
	return r
}

func (r *Request) Context(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

func (r *Request) Body(obj interface{}) *Request {
	data, err := json.Marshal(obj)

	if err != nil {
		r.err = err
	} else {
		r.body = bytes.NewBuffer(data)
	}
	return r
}

func (r *Request) Do(objResponse interface{}) error {
	if r.err != nil {
		return r.err
	}

	httpClient := r.client
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	httpRequest, err := r.prepareRequest()

	if err != nil {
		return err
	}

	resp, err := httpClient.Do(httpRequest)

	if err != nil {
		ctx := httpRequest.Context()
		if ctx != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
		}

		return err
	}

	defer func() {
		resp.Body.Close()
	}()

	// check code result
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusPartialContent {
		return fmt.Errorf("grafana respond with the status code %d for the request %s", resp.StatusCode, resp.Request.URL.String())
	}

	// Deserialize the json response
	err = json.NewDecoder(resp.Body).Decode(objResponse)
	if err != nil {
		if err == io.EOF {
			err = nil // ignore EOF errors caused by empty response body
		}
		return err
	}

	return nil
}

func (r *Request) prepareRequest() (*http.Request, error) {
	finalUrl, err := r.url()
	if err != nil {
		return nil, err
	}
	httpRequest, err := http.NewRequest(r.method, finalUrl, r.body)

	if err != nil {
		return nil, err
	}

	// set the context if exists
	if r.ctx != nil {
		httpRequest.WithContext(r.ctx)
	}

	// set the default content type
	if r.body != nil {
		httpRequest.Header.Set("Content-Type", "application/json")
	}

	// set the accept content type
	httpRequest.Header.Set("Accept", "application/json")

	// set the token
	if len(r.token) > 0 {
		httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.token))
	}

	return httpRequest, nil
}

func (r *Request) url() (string, error) {
	subPath, err := r.buildSubpath()

	if err != nil {
		return "", nil
	}

	finalURL := &url.URL{}
	if r.baseURL != nil {
		finalURL = r.baseURL
	}
	finalURL.Path = r.pathPrefix + subPath

	if r.queryParam == nil {
		finalURL.RawQuery = r.queryParam.Encode()
	}

	return finalURL.String(), nil
}

func (r *Request) buildSubpath() (string, error) {
	subPath := r.subpath
	if len(subPath) <= 0 {
		return "", nil
	}

	// a subpath exist, try to figure if we had to replace some path parameter
	pathParamRegexp := regexp.MustCompile(`/:(?P<PathParam>[[:alpha:]]+)*`)

	if !pathParamRegexp.Match([]byte(subPath)) {
		return subPath, nil
	}

	// the regexp match some pathParam, so we have to check if we have some pathParam
	if r.pathParam == nil {
		return "", fmt.Errorf("unable to replace the path parameter because it's empty")
	}

	matchGroups := pathParamRegexp.FindAllStringSubmatch(subPath, -1)

	for _, matchGroup := range matchGroups {
		if len(matchGroup) == 2 {
			pathParam := matchGroup[1]
			paramValue, hasParam := r.pathParam[pathParam]
			if !hasParam {
				return "", fmt.Errorf("unable to find the value of the path parameter %s", pathParam)
			}
			subPath = strings.Replace(subPath, ":"+pathParam, paramValue, 1)
		}
	}
	return subPath, nil
}
