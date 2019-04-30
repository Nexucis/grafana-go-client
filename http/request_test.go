package http

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestRequest_BuildSubPath(t *testing.T) {
	testSuites := []struct {
		title         string
		subPath       string
		pathParam     map[string]string
		expectedPath  string
		expectedError bool
	}{
		{
			title:         "Empty subpath",
			subPath:       "",
			pathParam:     map[string]string{"test": "value"},
			expectedPath:  "",
			expectedError: false,
		},
		{
			title:         "Unique path param to replace",
			subPath:       "/dasboard/:dashboardID/snapshot",
			pathParam:     map[string]string{"dashboardID": "15"},
			expectedPath:  "/dasboard/15/snapshot",
			expectedError: false,
		},
		{
			title:         "Multiple path param to replace",
			subPath:       "/:id/test/:otherId",
			pathParam:     map[string]string{"id": "5", "otherId": "45"},
			expectedPath:  "/5/test/45",
			expectedError: false,
		},
		{
			title:         "Not enough path param",
			subPath:       "/:id/test/:otherId",
			pathParam:     map[string]string{"id": "5"},
			expectedPath:  "",
			expectedError: true,
		},
		{
			title:         "Subpath without path param",
			subPath:       "/test",
			pathParam:     map[string]string{"id": "5", "otherId": "45"},
			expectedPath:  "/test",
			expectedError: false,
		},
	}

	for _, testSuite := range testSuites {
		info := fmt.Sprintf("test %s failed", testSuite.title)
		request := &Request{subpath: testSuite.subPath, pathParam: testSuite.pathParam}
		result, err := request.buildSubpath()

		assert.Equal(t, testSuite.expectedError, err != nil, info)
		assert.Equal(t, testSuite.expectedPath, result, info)
	}
}

func TestRequest_AddQueryParam(t *testing.T) {
	queryParams := map[string][]string{"tags": {"tag1", "tag2"}, "limit": {"100"}}

	request := &Request{}

	for queryName, v := range queryParams {
		for _, queryValue := range v {
			request.AddQueryParam(queryName, queryValue)
		}
	}

	assert.Equal(t, queryParams, map[string][]string(request.queryParam))
}

func TestRequest_SetPathParam(t *testing.T) {
	testSuites := []struct {
		title      string
		pathParams map[string][]string
		expected   map[string]string
	}{
		{
			title:      "unique value for each path param",
			pathParams: map[string][]string{"id": {"1"}, "dasboardID": {"45"}},
			expected:   map[string]string{"id": "1", "dasboardID": "45"},
		},
		{
			title:      "multiple value for each path param",
			pathParams: map[string][]string{"id": {"1"}, "dasboardID": {"45", "46", "47"}},
			expected:   map[string]string{"id": "1", "dasboardID": "47"},
		},
	}

	for _, testSuite := range testSuites {
		request := &Request{}
		for pathName, v := range testSuite.pathParams {
			for _, pathValue := range v {
				request.SetPathParam(pathName, pathValue)
			}
		}

		assert.Equal(t, testSuite.expected, request.pathParam, fmt.Sprintf("test %s failed", testSuite.title))
	}
}

func TestRequest_SetSubPath(t *testing.T) {
	subPath := "/:id/metrics"
	request := Request{}
	request.SetSubPath(subPath)
	assert.Equal(t, subPath, request.subpath)
}

func TestRequest_URL(t *testing.T) {
	testSuites := []struct {
		title         string
		baseUrl       url.URL
		pathPrefix    string
		subPath       string
		pathParam     map[string]string
		queryParams   map[string][]string
		expectedUrl   string
		expectedError bool
	}{
		{
			title:         "Url returns as is",
			baseUrl:       url.URL{Scheme: "http", Host: "localhost:8080"},
			expectedUrl:   "http://localhost:8080",
			expectedError: false,
		},
		{
			title:         "Url with subpath",
			baseUrl:       url.URL{Scheme: "http", Host: "localhost:8080"},
			pathPrefix:    "/api",
			subPath:       "/users",
			expectedUrl:   "http://localhost:8080/api/users",
			expectedError: false,
		},
		{
			title:         "Url with subpath and path param",
			baseUrl:       url.URL{Scheme: "http", Host: "localhost:8080"},
			subPath:       "/api/dasboard/:dashboardID/snapshot",
			pathParam:     map[string]string{"dashboardID": "15"},
			expectedUrl:   "http://localhost:8080/api/dasboard/15/snapshot",
			expectedError: false,
		},
		{
			title:   "url with ascii query param",
			baseUrl: url.URL{Scheme: "http", Host: "localhost:8080"},
			queryParams: map[string][]string{
				"tags":  {"tag1", "tag2"},
				"limit": {"100"},
			},
			expectedUrl:   "http://localhost:8080?limit=100&tags=tag1&tags=tag2",
			expectedError: false,
		},
		{
			title:   "url with encoding query param",
			baseUrl: url.URL{Scheme: "http", Host: "localhost:8080"},
			queryParams: map[string][]string{
				"version": {"v1&v2", "v3"},
				"filter":  {"表"},
			},
			expectedUrl:   "http://localhost:8080?filter=%E8%A1%A8&version=v1%26v2&version=v3",
			expectedError: false,
		},
		{
			title:      "complete test",
			baseUrl:    url.URL{Scheme: "http", Host: "localhost:8080"},
			pathPrefix: "/api",
			subPath:    "/dasboard/:dashboardID/snapshot",
			pathParam:  map[string]string{"dashboardID": "15"},
			queryParams: map[string][]string{
				"tags":    {"tag1", "tag2"},
				"limit":   {"100"},
				"version": {"v1&v2", "v3"},
				"filter":  {"表"},
			},
			expectedUrl:   "http://localhost:8080/api/dasboard/15/snapshot?filter=%E8%A1%A8&limit=100&tags=tag1&tags=tag2&version=v1%26v2&version=v3",
			expectedError: false,
		},
	}

	for _, testSuite := range testSuites {
		info := fmt.Sprintf("test %s failed", testSuite.title)
		request := &Request{
			baseURL:    &testSuite.baseUrl,
			pathPrefix: testSuite.pathPrefix,
			subpath:    testSuite.subPath,
			pathParam:  testSuite.pathParam,
			queryParam: testSuite.queryParams,
		}
		result, err := request.url()
		assert.Equal(t, testSuite.expectedError, err != nil, info)
		assert.Equal(t, testSuite.expectedUrl, result, info)
	}
}
