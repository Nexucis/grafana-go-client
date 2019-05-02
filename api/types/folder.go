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

type Folder struct {
	ID        int64     `json:"id"`
	UID       string    `json:"uid"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	HasACL    bool      `json:"hasAcl"`
	CanSave   bool      `json:"canSave"`
	CanEdit   bool      `json:"canEdit"`
	CanAdmin  bool      `json:"canAdmin"`
	CreatedBy string    `json:"createdBy"`
	Created   time.Time `json:"created"`
	UpdatedBy string    `json:"updatedBy"`
	Updated   time.Time `json:"updated"`
	Version   int       `json:"version"`
}

type SimpleFolder struct {
	ID    int64  `json:"id"`
	UID   string `json:"uid"`
	Title string `json:"title"`
}

type UpdateFolder struct {
	// Provide another unique identifier to change the unique identifier stored.
	UID       string `json:"uid"`
	Title     string `json:"title"`
	Version   int    `json:"version"`
	Overwrite bool   `json:"overwrite"`
}

type DashboardACLUpdateItem struct {
	UserID     int64          `json:"userId"`
	TeamID     int64          `json:"teamId"`
	Role       *RoleType      `json:"role,omitempty"`
	Permission PermissionType `json:"permission"`
}
