package http

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRequest_BuildSubPath(t *testing.T) {
	testSuites := []struct {
		subPath       string
		pathParam     map[string]string
		expectedPath  string
		expectedError bool
	}{
		{"", map[string]string{"test": "value"}, "", false},
		{"/dasboard/:dashboardID/snapshot", map[string]string{"dashboardID": "15"}, "/dasboard/15/snapshot", false},
		{"/:id/test/:otherId", map[string]string{"id": "5", "otherId": "45"}, "/5/test/45", false},
		{"/:id/test/:otherId", map[string]string{"id": "5"}, "", true},
		{"/test", map[string]string{"id": "5", "otherId": "45"}, "/test", false},
	}

	for _, testSuite := range testSuites {
		request := Request{subpath: testSuite.subPath, pathParam: testSuite.pathParam}
		result, err := request.buildSubpath()

		if !testSuite.expectedError {
			assert.Nil(t, err)
			assert.Equal(t, testSuite.expectedPath, result)
		} else {
			assert.NotNil(t, err)
		}
	}
}

func TestRequest_AddQueryParam(t *testing.T) {
	queryParams := map[string][]string{"tags": {"tag1", "tag2"}, "limit": {"100"}}

	request := Request{}

	for queryName, v := range queryParams {
		for _, queryValue := range v {
			request.AddQueryParam(queryName, queryValue)
		}
	}

	assert.Equal(t, queryParams, map[string][]string(request.queryParam))
}

func TestRequest_SetPathParam(t *testing.T) {
	testSuites := []struct {
		pathParams map[string][]string
		expected   map[string]string
	}{
		{map[string][]string{"id": {"1"}, "dasboardID": {"45"}},
			map[string]string{"id": "1", "dasboardID": "45"}},
		{map[string][]string{"id": {"1"}, "dasboardID": {"45", "46", "47"}},
			map[string]string{"id": "1", "dasboardID": "47"}},
	}

	for _, testSuite := range testSuites {
		request := Request{}
		for pathName, v := range testSuite.pathParams {
			for _, pathValue := range v {
				request.SetPathParam(pathName, pathValue)
			}
		}

		assert.Equal(t, testSuite.expected, request.pathParam)
	}
}

func TestRequest_SetSubPath(t *testing.T) {
	subPath := "/:id/metrics"
	request := Request{}
	request.SetSubPath(subPath)
	assert.Equal(t, subPath, request.subpath)
}
