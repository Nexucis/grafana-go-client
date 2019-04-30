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
	"testing"

	"github.com/nexucis/grafana-go-client/api/types"
	"github.com/nexucis/grafana-go-client/http"
	"github.com/stretchr/testify/assert"
)

func TestKey_CreateError(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	key := initKeyTest(t)

	_, err := key.Create(&types.APIKeyForm{Name: "test_key", Role: "admin"})

	assert.Equal(t, 400, err.(*http.RequestError).StatusCode)
	assert.Equal(t, "JSON validation error: invalid role value: admin", err.(*http.RequestError).Message)
	teardownKey(t)
}

func TestKey_Create(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	key := initKeyTest(t)

	response, err := key.Create(&types.APIKeyForm{Name: "test_key", Role: types.RoleAdmin})

	assert.Nil(t, err)
	assert.Equal(t, "test_key", response.Name)
	assert.True(t, len(response.Key) > 0)

	teardownKey(t)
}

func TestKey_Get(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	key := initKeyTest(t)

	_, err := key.Create(&types.APIKeyForm{Name: "test_key", Role: types.RoleAdmin})
	assert.Nil(t, err)

	response, err := key.Get()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(response))
	assert.Equal(t, "test_key", response[0].Name)
	assert.Equal(t, types.RoleAdmin, response[0].Role)

	teardownKey(t)
}

func initKeyTest(t *testing.T) KeyInterface {
	httpClient, err := getRestClientWithBasicAuth()
	assert.Nil(t, err)
	return newKey(httpClient)
}

func teardownKey(t *testing.T) {
	keyClient := initKeyTest(t)
	keys, _ := keyClient.Get()
	for _, key := range keys {
		keyClient.Delete(key.Id)
	}
}
