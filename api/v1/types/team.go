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

package types

type Team struct {
	Id          int64  `json:"id"`
	OrgId       int64  `json:"orgId"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AvatarUrl   string `json:"avatarUrl"`
	MemberCount int64  `json:"memberCount"`
}

type SearchTeam struct {
	TotalCount int64   `json:"totalCount"`
	Teams      []*Team `json:"teams"`
	Page       int     `json:"page"`
	PerPage    int     `json:"perPage"`
}

type TeamMember struct {
	OrgId     int64  `json:"orgId"`
	TeamId    int64  `json:"teamId"`
	UserId    int64  `json:"userId"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	AvatarUrl string `json:"avatarUrl"`
}

type CreateOrUpdateTeam struct {
	Name  string `json:"name" binding:"Required"`
	Email string `json:"email"`
}
