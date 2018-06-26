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
