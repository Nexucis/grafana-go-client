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

	"github.com/nexucis/grafana-go-client/api/v1/types"
	"github.com/stretchr/testify/assert"
)

func TestTeam_Create(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}
	team := initTeamTest(t)
	teamToCreate := &types.CreateOrUpdateTeam{Name: "my_team", Email: "john.doe@my-company.com"}
	response, err := team.Create(teamToCreate)

	assert.Nil(t, err)
	assert.Equal(t, "my_team", response.Name)
	assert.Equal(t, "john.doe@my-company.com", response.Email)
	assert.Equal(t, 0, response.MemberCount)
	assert.True(t, response.Id > 0)
}

func initTeamTest(t *testing.T) TeamInterface {
	httpClient, err := getRestClientWithBasicAuth()
	assert.Nil(t, err)
	return newTeam(httpClient)
}
