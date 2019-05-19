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

package grafanahttp

import (
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"net/url"
	"time"
)

const connectionTimeout = 30 * time.Second

// RestConfigClient defines all parameter that can be set to customize the RESTClient
type RestConfigClient struct {
	InsecureTLS bool   `yaml:"insecure-tls"`
	BaseURL     string `yaml:"baseURL"`
	Token       string `yaml:"token"`
}

func NewWithURL(rawURL string) (*RESTClient, error) {
	return NewFromConfig(&RestConfigClient{
		BaseURL: rawURL,
	})
}

// NewFromConfig create an instance of RESTClient using the config passed as parameter
func NewFromConfig(config *RestConfigClient) (*RESTClient, error) {
	if config == nil {
		return nil, errors.New("configuration cannot be empty")
	}
	roundTripper := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   connectionTimeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: config.InsecureTLS}, // nolint: gas, gosec
	}

	httpClient := &http.Client{
		Transport: roundTripper,
		Timeout:   connectionTimeout,
	}

	u, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, err
	}

	return &RESTClient{
		Token:   config.Token,
		BaseURL: u,
		Client:  httpClient,
	}, nil

}

// RESTClient defines an HTTP client designed for the HTTP request to a REST API.
type RESTClient struct {
	// the token used to be authenticated, not mandatory if it's used the basic authentication
	Token string
	// base is the root URL for all invocations of the client
	BaseURL *url.URL
	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	Client *http.Client
}

// Get begins a GET request. Short for c.newRequest("GET")
func (c *RESTClient) Get(pathPrefix string) *Request {
	return c.newRequest(http.MethodGet, pathPrefix)
}

// Post begins a Post request. Short for c.newRequest("POST")
func (c *RESTClient) Post(pathPrefix string) *Request {
	return c.newRequest(http.MethodPost, pathPrefix)
}

// Put begins a Put request. Short for c.newRequest("PUT")
func (c *RESTClient) Put(pathPrefix string) *Request {
	return c.newRequest(http.MethodPut, pathPrefix)
}

// Patch begins a Patch request. Short for c.newRequest("PATCH")
func (c *RESTClient) Patch(pathPrefix string) *Request {
	return c.newRequest(http.MethodPatch, pathPrefix)
}

// Delete begins a Delete request. Short for c.newRequest("DELETE")
func (c *RESTClient) Delete(pathPrefix string) *Request {
	return c.newRequest(http.MethodDelete, pathPrefix)
}

func (c *RESTClient) newRequest(method string, pathPrefix string) *Request {
	return NewRequest(c.Client, method, c.BaseURL, pathPrefix, c.Token)
}
