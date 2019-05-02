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
	"github.com/nexucis/grafana-go-client/grafanahttp"
	"testing"

	"github.com/nexucis/grafana-go-client/api/types"
	"github.com/stretchr/testify/assert"
)

func TestAdmin_GetSettings(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	admin := initAdminTest(t)
	settings, err := admin.GetSettings()

	assert.Nil(t, err)
	assert.NotNil(t, settings)
	assert.NotNil(t, settings.Default)
	assert.True(t, len(settings.Default.AppMode) > 0)
}

func TestAdmin_CreateUser(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	admin := initAdminTest(t)
	user := &types.AdminCreateUserForm{Email: "jdoe@compagny.com", Login: "jdoe", Password: "jdoe", Name: "John Doe"}
	response, err := admin.CreateUser(user)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.True(t, response.ID > 0)
	assert.Equal(t, "User created", response.Message)

	removeGlobalUser(t, response.ID)
}

func TestAdmin_CreateUserError(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	admin := initAdminTest(t)
	user := &types.AdminCreateUserForm{Email: "jdoe@compagny.com", Login: "jdoe", Password: "jdoe", Name: "John Doe"}
	response, err := admin.CreateUser(user)

	assert.Nil(t, err)
	assert.NotNil(t, response)

	// if we recreate the same user, it should return an error
	_, err = admin.CreateUser(user)
	assert.NotNil(t, err)
	assert.Equal(t, 500, err.(*grafanahttp.RequestError).StatusCode)
	assert.Equal(t, "failed to create user", err.(*grafanahttp.RequestError).Message)

	removeGlobalUser(t, response.ID)
}

func TestAdmin_UpdateUserPassword(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	admin := initAdminTest(t)
	user := &types.AdminCreateUserForm{Email: "jdoe@compagny.com", Login: "jdoe", Password: "jdoe", Name: "John Doe"}
	response, _ := admin.CreateUser(user)

	// trying to update the password now
	err := admin.UpdateUserPassword(response.ID, "anotherPassword")

	assert.Nil(t, err)
	removeGlobalUser(t, response.ID)
}

func TestAdmin_DeleteUser(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	admin := initAdminTest(t)
	user := &types.AdminCreateUserForm{Email: "jdoe@compagny.com", Login: "jdoe", Password: "jdoe", Name: "John Doe"}
	response, _ := admin.CreateUser(user)

	//trying to delete the user now
	err := admin.DeleteUser(response.ID)

	assert.Nil(t, err)
}

func TestAdmin_GetStats(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	admin := initAdminTest(t)
	stats, err := admin.GetStats()

	assert.Nil(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, 1, stats.Users)
	assert.Equal(t, 1, stats.Orgs)
}

func initAdminTest(t *testing.T) AdminInterface {
	httpClient, err := getRestClientWithBasicAuth()
	assert.Nil(t, err)
	return newAdmin(httpClient)
}

func removeGlobalUser(t *testing.T, ids ...int64) {
	adminClient := initAdminTest(t)
	for _, id := range ids {
		adminClient.DeleteUser(id)
	}
}
