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

type SimpleDashboard struct {
	Version int    `json:"version"`
	ID      int64  `json:"id"`
	UID     string `json:"uid"`
	Status  string `json:"status"`
	Slug    string `json:"slug"`
	Url     string `json:"url"`
}

type Dashboard struct {
	Version int    `json:"version"`
	GnetID  int64  `json:"gnetId"`
	ID      int64  `json:"id"`
	UID     string `json:"uid"`
	Title   string `json:"title"`
}

type DashboardTags struct {
	Term  string `json:"term"`
	Count int    `json:"count"`
}

type DashboardVersion struct {
	ID            int         `json:"id"`
	DashboardId   int64       `json:"dashboardId"`
	ParentVersion int         `json:"parentVersion"`
	RestoredFrom  int         `json:"restoredFrom"`
	Version       int         `json:"version"`
	Created       time.Time   `json:"created"`
	CreatedBy     string      `json:"createdBy"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data"`
}

type DashboardVersionMeta struct {
	DashboardVersion
	CreatedBy string `json:"createdBy"`
}
