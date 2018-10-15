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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type Request struct {
	client *http.Client
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

func NewRequest(client *http.Client, method string, baseURL *url.URL, pathPrefix string, token string) *Request {
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

func (r *Request) Do() *Response {
	if r.err != nil {
		return &Response{err: r.err}
	}

	httpClient := r.client
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	httpRequest, err := r.prepareRequest()

	if err != nil {
		return &Response{err: err}
	}

	resp, err := httpClient.Do(httpRequest)

	if err != nil {
		ctx := httpRequest.Context()
		if ctx != nil {
			select {
			case <-ctx.Done():
				return &Response{err: ctx.Err()}
			default:
			}
		}

		return &Response{err: err}
	}

	defer func() {
		resp.Body.Close()
	}()

	// Deserialize the json response
	if resp.Body != nil {
		data, err := ioutil.ReadAll(resp.Body)
		return &Response{body: data, err: err, statusCode: resp.StatusCode}
	}

	return &Response{statusCode: resp.StatusCode}
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
		httpRequest = httpRequest.WithContext(r.ctx)
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

	if r.queryParam != nil {
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

type GrafanaErrorResponse struct {
	Message string `json:"message,omitempty"`
	Status  string `json:"status,omitempty"`
}

type RequestError struct {
	Message    string
	StatusCode int
	Err        error
}

func (re *RequestError) Error() string {
	err := "something wrong happened with the request to Grafana."

	if re.Err != nil {
		err = err + " Error: " + re.Err.Error()
	}
	if len(re.Message) > 0 {
		err = err + " Message: " + re.Message
	}

	if re.StatusCode > 0 {
		err = err + " StatusCode: " + strconv.Itoa(re.StatusCode)
	}

	return err
}

type Response struct {
	body       []byte
	err        error
	statusCode int
}

func (r *Response) Error() error {

	e := &RequestError{Err: r.err}
	// check code result
	if r.statusCode < http.StatusOK || r.statusCode > http.StatusPartialContent {
		// check error message contains in the body
		if r.body != nil {
			g := &GrafanaErrorResponse{}
			err := json.Unmarshal(r.body, g)

			if err != nil {
				//trying to find the message in a more generic struct
				var genericMessage []map[string]interface{}
				err2 := json.Unmarshal(r.body, &genericMessage)
				if err2 != nil {
					// in this case something horrible append on client side
					e.Err = fmt.Errorf("initial error : %s. Something horrible append when the client tryed to decode the error message: %s", r.err, err2)
				} else {
					for _, j := range genericMessage {
						for k, v := range j {
							if k == "message" {
								g.Message = v.(string)
								break
							}
						}
						if len(g.Message) > 0 {
							break
						}
					}
				}

			}
			e.Message = g.Message
		}
		e.StatusCode = r.statusCode
	}

	if e.Err != nil || e.StatusCode > 0 || len(e.Message) > 0 {
		return e
	}

	return nil
}

func (r *Response) SaveAsObj(respObj interface{}) error {
	err := r.Error()

	if err != nil {
		return err
	}

	if r.body != nil {
		err = json.Unmarshal(r.body, respObj)
		if err != nil {
			return fmt.Errorf("unable to decode the response body. Error %s", err)
		}
	}
	return nil
}
