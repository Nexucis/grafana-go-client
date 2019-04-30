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

type TempUserStatus string

const (
	TmpUserSignUpStarted TempUserStatus = "SignUpStarted"
	TmpUserInvitePending TempUserStatus = "InvitePending"
	TmpUserCompleted     TempUserStatus = "Completed"
	TmpUserRevoked       TempUserStatus = "Revoked"
)

type Org struct {
	Id      int64   `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type SimpleOrg struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type OrgQuota struct {
	OrgId  int64  `json:"org_id"`
	Target string `json:"target"`
	Limit  int64  `json:"limit"`
	Used   int64  `json:"used"`
}

type AddOrgUser struct {
	LoginOrEmail string   `json:"loginOrEmail" binding:"Required"`
	Role         RoleType `json:"role" binding:"Required"`
}

type UpdateOrgUser struct {
	Role RoleType `json:"role" binding:"Required"`
}

type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	ZipCode  string `json:"zipCode"`
	State    string `json:"state"`
	Country  string `json:"country"`
}

type OrgUser struct {
	OrgId         int64     `json:"orgId"`
	UserId        int64     `json:"userId"`
	Email         string    `json:"email"`
	AvatarUrl     string    `json:"avatarUrl"`
	Login         string    `json:"login"`
	Role          string    `json:"role"`
	LastSeenAt    time.Time `json:"lastSeenAt"`
	LastSeenAtAge string    `json:"lastSeenAtAge"`
}

type TempOrgUser struct {
	Id             int64          `json:"id"`
	OrgId          int64          `json:"orgId"`
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Role           RoleType       `json:"role"`
	InvitedByLogin string         `json:"invitedByLogin"`
	InvitedByEmail string         `json:"invitedByEmail"`
	InvitedByName  string         `json:"invitedByName"`
	Code           string         `json:"code"`
	Status         TempUserStatus `json:"status"`
	Url            string         `json:"url"`
	EmailSent      bool           `json:"emailSent"`
	EmailSentOn    time.Time      `json:"emailSentOn"`
	Created        time.Time      `json:"createdOn"`
}

type AddInvite struct {
	LoginOrEmail string   `json:"loginOrEmail" binding:"Required"`
	Name         string   `json:"name"`
	Role         RoleType `json:"role" binding:"Required"`
	SendEmail    bool     `json:"sendEmail"`
}

type OrgPrefs struct {
	Theme           string `json:"theme"`
	HomeDashboardID int64  `json:"homeDashboardId"`
	Timezone        string `json:"timezone"`
}
