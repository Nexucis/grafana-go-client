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

import "time"

type UserProfile struct {
	ID             int64  `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Login          string `json:"login"`
	Theme          string `json:"theme"`
	OrgID          int64  `json:"orgId"`
	IsGrafanaAdmin bool   `json:"isGrafanaAdmin"`
}

type UserSearchHit struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Login         string    `json:"login"`
	Email         string    `json:"email"`
	AvatarURL     string    `json:"avatarUrl"`
	IsAdmin       bool      `json:"isAdmin"`
	LastSeenAt    time.Time `json:"lastSeenAt"`
	LastSeenAtAge string    `json:"lastSeenAtAge"`
}

type UserSearchWithPaging struct {
	TotalCount int64
	Page       int64
	PerPage    int64
	Users      []*UserSearchHit
}

type UpdateCurrentUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Login string `json:"login"`
	Theme string `json:"theme"`
}

type UserQuota struct {
	UserID int64  `json:"user_id"`
	Target string `json:"target"`
	Limit  int64  `json:"limit"`
	Used   int64  `json:"used"`
}

type UpdateUserQuota struct {
	Target string `json:"target"`
	Limit  int64  `json:"limit"`
	UserID int64  `json:"-"`
}

type UserPreference struct {
	Theme           string `json:"theme"`
	HomeDashboardID int64  `json:"homeDashboardId"`
	Timezone        string `json:"timezone"`
}

type UserOrg struct {
	OrgID int64    `json:"orgId"`
	Name  string   `json:"name"`
	Role  RoleType `json:"role"`
}

type UserOrgList struct {
	UserID int64
	Result []*UserOrg
}

type UpdatePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
