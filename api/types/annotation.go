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

type AnnotationQueryType string

const (
	AnnSearchByAnnotation AnnotationQueryType = "annotation"
	AnnSearchByAlert      AnnotationQueryType = "alert"
)

type PostAnnotations struct {
	DashboardID int64       `json:"dashboardId"`
	PanelID     int64       `json:"panelId"`
	Time        int64       `json:"time"`
	Text        string      `json:"text"`
	Tags        []string    `json:"tags"`
	Data        interface{} `json:"data"`
	IsRegion    bool        `json:"isRegion"`
	TimeEnd     int64       `json:"timeEnd"`
}

type UpdateAnnotations struct {
	ID       int64    `json:"id"`
	Time     int64    `json:"time"`
	Text     string   `json:"text"`
	Tags     []string `json:"tags"`
	IsRegion bool     `json:"isRegion"`
	TimeEnd  int64    `json:"timeEnd"`
}

type DeleteAnnotations struct {
	AlertID      int64 `json:"alertId"`
	DashboardID  int64 `json:"dashboardId"`
	PanelID      int64 `json:"panelId"`
	AnnotationID int64 `json:"annotationId"`
	RegionID     int64 `json:"regionId"`
}

type PostGraphiteAnnotations struct {
	When int64    `json:"when"`
	What string   `json:"what"`
	Data string   `json:"data"`
	Tags []string `json:"tags"`
}

type ResponseCreateAnnotation struct {
	ID      int64  `json:"id"`
	EndID   int64  `json:"endId"`
	Message string `json:"message"`
}

type ResponseCreateGraphiteAnnotation struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

type ResponseGetAnnotation struct {
	ID          int64       `json:"id"`
	AlertID     int64       `json:"alertId"`
	AlertName   string      `json:"alertName"`
	DashboardID int64       `json:"dashboardId"`
	PanelID     int64       `json:"panelId"`
	UserID      int64       `json:"userId"`
	NewState    string      `json:"newState"`
	PrevState   string      `json:"prevState"`
	Created     int64       `json:"created"`
	Updated     int64       `json:"updated"`
	Time        int64       `json:"time"`
	Text        string      `json:"text"`
	RegionID    int64       `json:"regionId"`
	Tags        []string    `json:"tags"`
	Login       string      `json:"login"`
	Email       string      `json:"email"`
	AvatarURL   string      `json:"avatarUrl"`
	Data        interface{} `json:"data"`
}
