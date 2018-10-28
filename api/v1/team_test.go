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
	teamID, err := team.Create(teamToCreate)

	assert.Nil(t, err)
	assert.True(t, teamID > 0)

	//clean test
	removeTeam(t, teamID)
}

func TestTeam_Update(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	team := initTeamTest(t)
	teamToCreate := &types.CreateOrUpdateTeam{Name: "my_team", Email: "john.doe@my-company.com"}
	teamID, _ := team.Create(teamToCreate)

	teamToUpdate := &types.CreateOrUpdateTeam{Name: "my_other_team", Email: "john.doe2@my-company.com"}

	err := team.Update(teamID, teamToUpdate)
	assert.Nil(t, err)

	//clean test
	removeTeam(t, teamID)
}

func TestTeam_Get(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	//clean test
	team := initTeamTest(t)
	teamToCreate := &types.CreateOrUpdateTeam{Name: "my_team", Email: "john.doe@my-company.com"}
	teamID, _ := team.Create(teamToCreate)

	//Get team
	response, err := team.Get((&QueryParameterTeams{}).PerPage(int64(10)))
	assert.Nil(t, err)
	assert.Equal(t, int64(1), response.TotalCount)
	assert.Equal(t, teamID, response.Teams[0].Id)
	assert.Equal(t, "my_team", response.Teams[0].Name)
	assert.Equal(t, "john.doe@my-company.com", response.Teams[0].Email)
	assert.Equal(t, int64(0), response.Teams[0].MemberCount)

	//clean test
	removeTeam(t, teamID)
}

func TestTeam_GetByID(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	team := initTeamTest(t)
	teamToCreate := &types.CreateOrUpdateTeam{Name: "my_team", Email: "john.doe@my-company.com"}
	teamID, _ := team.Create(teamToCreate)

	teamInfo, err := team.GetByID(teamID)
	assert.Nil(t, err)
	assert.Equal(t, "my_team", teamInfo.Name)
	assert.Equal(t, teamID, teamInfo.Id)
	assert.Equal(t, "john.doe@my-company.com", teamInfo.Email)

	// clean test
	removeTeam(t, teamID)
}

func TestTeam_Delete(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}



	team := initTeamTest(t)
	teamToCreate := &types.CreateOrUpdateTeam{Name: "my_team", Email: "john.doe@my-company.com"}
	teamID, _ := team.Create(teamToCreate)

	//trying to delete the user now
	err := team.Delete(teamID)
	assert.Nil(t, err)

	//clean test
	removeTeam(t, teamID)
}

func TestTeam_AddMembers(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	team := initTeamTest(t)
	teamToCreate := &types.CreateOrUpdateTeam{Name: "my_team", Email: "john.doe@my-company.com"}
	teamID, _ := team.Create(teamToCreate)

	//create user before binding it to the team
	admin := initAdminTest(t)
	user := &types.AdminCreateUserForm{Email: "jdoe@compagny.com", Login: "jdoe", Password: "jdoe", Name: "John Doe"}
	userResponse, _ := admin.CreateUser(user)
	userID := userResponse.ID

	//binding
	err := team.AddMembers(teamID, userID)
	assert.Nil(t, err)

	// clean test
	removeTeam(t, teamID)
	removeGlobalUser(t, userID)
}

func TestTeam_DeleteMembers(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	team := initTeamTest(t)
	teamToCreate := &types.CreateOrUpdateTeam{Name: "my_team", Email: "john.doe@my-company.com"}
	teamID, _ := team.Create(teamToCreate)

	//create user before binding it to the team
	admin := initAdminTest(t)
	user := &types.AdminCreateUserForm{Email: "jdoe@compagny.com", Login: "jdoe", Password: "jdoe", Name: "John Doe"}
	userResponse, _ := admin.CreateUser(user)
	userID := userResponse.ID

	//binding
	team.AddMembers(teamID, userResponse.ID)

	//delete the user in the team
	err := team.DeleteMembers(teamID, userResponse.ID)
	assert.Nil(t, err)

	//clean test
	removeTeam(t, teamID)
	removeGlobalUser(t, userID)
}

func TestTeam_GetMembers(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	team := initTeamTest(t)
	teamToCreate := &types.CreateOrUpdateTeam{Name: "my_team", Email: "john.doe@my-company.com"}
	teamID, _ := team.Create(teamToCreate)

	//create user before binding it to the team
	admin := initAdminTest(t)
	user := &types.AdminCreateUserForm{Email: "jdoe@compagny.com", Login: "jdoe", Password: "jdoe", Name: "John Doe"}
	userResponse, _ := admin.CreateUser(user)
	userID := userResponse.ID

	//binding
	team.AddMembers(teamID, userResponse.ID)

	// get team member
	teamMembers, err := team.GetMembers(teamID)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(teamMembers))
	assert.Equal(t, "jdoe@compagny.com", teamMembers[0].Email)
	assert.Equal(t, "jdoe", teamMembers[0].Login)
	assert.Equal(t, userResponse.ID, teamMembers[0].UserId)
	assert.Equal(t, teamID, teamMembers[0].TeamId)

	//clean test
	removeTeam(t, teamID)
	removeGlobalUser(t, userID)
}

func initTeamTest(t *testing.T) TeamInterface {
	httpClient, err := getRestClientWithBasicAuth()
	assert.Nil(t, err)
	return newTeam(httpClient)
}

func removeTeam(t *testing.T, ids ...int64) {
	teamClient := initTeamTest(t)
	for _, id := range ids {
		teamClient.Delete(id)
	}
}
