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
	"net/http"
	"net/url"
)

type RESTClient struct {
	// the token used to be authenticated, not mandatory if it's used the basic authentication
	token string
	// base is the root URL for all invocations of the client
	baseURL *url.URL
	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	client *http.Client
}

func NewWithUrl(rawURL string) (*RESTClient, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	return &RESTClient{
		baseURL: u,
	}, nil
}

func (c *RESTClient) Get(pathPrefix string) *Request {
	return c.newRequest(http.MethodGet, pathPrefix)
}

func (c *RESTClient) Post(pathPrefix string) *Request {
	return c.newRequest(http.MethodPost, pathPrefix)
}

func (c *RESTClient) Put(pathPrefix string) *Request {
	return c.newRequest(http.MethodPut, pathPrefix)
}

func (c *RESTClient) Delete(pathPrefix string) *Request {
	return c.newRequest(http.MethodDelete, pathPrefix)
}

func (c *RESTClient) newRequest(method string, pathPrefix string) *Request {
	return NewRequest(c.client, method, c.baseURL, pathPrefix, c.token)
}
