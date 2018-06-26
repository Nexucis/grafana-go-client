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

type PostAnnotations struct {
	DashboardId int64       `json:"dashboardId"`
	PanelId     int64       `json:"panelId"`
	Time        int64       `json:"time"`
	Text        string      `json:"text"`
	Tags        []string    `json:"tags"`
	Data        interface{} `json:"data"`
	IsRegion    bool        `json:"isRegion"`
	TimeEnd     int64       `json:"timeEnd"`
}

type UpdateAnnotations struct {
	Id       int64    `json:"id"`
	Time     int64    `json:"time"`
	Text     string   `json:"text"`
	Tags     []string `json:"tags"`
	IsRegion bool     `json:"isRegion"`
	TimeEnd  int64    `json:"timeEnd"`
}

type DeleteAnnotations struct {
	AlertId      int64 `json:"alertId"`
	DashboardId  int64 `json:"dashboardId"`
	PanelId      int64 `json:"panelId"`
	AnnotationId int64 `json:"annotationId"`
	RegionId     int64 `json:"regionId"`
}

type PostGraphiteAnnotations struct {
	When int64    `json:"when"`
	What string   `json:"what"`
	Data string   `json:"data"`
	Tags []string `json:"tags"`
}
