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

import (
	"time"
)

type ResponseAlertNotification struct {
	Id        int64       `json:"id"`
	OrgId     int64       `json:"-"`
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	IsDefault bool        `json:"isDefault"`
	Settings  interface{} `json:"settings"`
	Created   time.Time   `json:"created"`
	Updated   time.Time   `json:"updated"`
}

type CreateAlertNotification struct {
	Name      string      `json:"name"  binding:"Required"`
	Type      string      `json:"type"  binding:"Required"`
	IsDefault bool        `json:"isDefault"`
	Settings  interface{} `json:"settings"`
}

type CreateTestAlertNotification struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	Settings interface{} `json:"settings"`
}

type UpdateAlertNotification struct {
	Id        int64       `json:"id"  binding:"Required"`
	Name      string      `json:"name"  binding:"Required"`
	Type      string      `json:"type"  binding:"Required"`
	IsDefault bool        `json:"isDefault"`
	Settings  interface{} `json:"settings"  binding:"Required"`
}
