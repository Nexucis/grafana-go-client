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
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/nexucis/grafana-go-client/http"
)

func TestKey_CreateError(t *testing.T) {
	httpClient, err := getRestClientWithBasicAuth()
	assert.Nil(t, err)

	key := newKey(httpClient)

	_, err = key.Create(&APIKeyForm{"test_key", "admin"})

	assert.Equal(t, 400, err.(*http.RequestError).StatusCode)
	assert.Equal(t, "JSON validation error: invalid role value: admin", err.(*http.RequestError).Message)
}

func TestKey_Create(t *testing.T) {
	httpClient, err := getRestClientWithBasicAuth()
	assert.Nil(t, err)

	key := newKey(httpClient)

	response, err := key.Create(&APIKeyForm{"test_key", RoleAdmin})

	assert.Nil(t, err)
	assert.Equal(t, "test_key", response.Name)
	assert.True(t, len(response.Key) > 0)
}
