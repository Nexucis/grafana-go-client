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

type ResponseCreateAnnotation struct {
	Id      int64  `json:"id"`
	EndId   int64  `json:"endId"`
	Message string `json:"message"`
}

type ResponseCreateGraphiteAnnotation struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type ResponseGetAnnotation struct {
	Id          int64       `json:"id"`
	AlertId     int64       `json:"alertId"`
	AlertName   string      `json:"alertName"`
	DashboardId int64       `json:"dashboardId"`
	PanelId     int64       `json:"panelId"`
	UserId      int64       `json:"userId"`
	NewState    string      `json:"newState"`
	PrevState   string      `json:"prevState"`
	Created     int64       `json:"created"`
	Updated     int64       `json:"updated"`
	Time        int64       `json:"time"`
	Text        string      `json:"text"`
	RegionId    int64       `json:"regionId"`
	Tags        []string    `json:"tags"`
	Login       string      `json:"login"`
	Email       string      `json:"email"`
	AvatarUrl   string      `json:"avatarUrl"`
	Data        interface{} `json:"data"`
}

type QueryParamAnnotation struct {
	// epoch datetime in milliseconds. Optional.
	from int64
	// epoch datetime in milliseconds. Optional.
	to int64
	// number. Optional. Find annotations created by a specific user
	userId int64
	// number. Optional. Find annotations for a specified alert.
	alertId int64
	// number. Optional. Find annotations that are scoped to a specific dashboard
	dashboardId int64
	// number. Optional. Find annotations that are scoped to a specific panel
	panelId int64
	// string. Optional. Use this to filter global annotations. Global annotations are annotations from an annotation data source that are not connected specifically to a dashboard or panel
	tags []string
	// string. Optional. alert|annotation Return alerts or user created annotations
	_type string
	// number. Optional - default is 100. Max limit for results returned.
	limit int64
}

func (query *QueryParamAnnotation) From(from int64) *QueryParamAnnotation {
	query.from = from
	return query
}

func (query *QueryParamAnnotation) To(to int64) *QueryParamAnnotation {
	query.to = to
	return query
}

func (query *QueryParamAnnotation) AlertID(alertID int64) *QueryParamAnnotation {
	query.alertId = alertID
	return query
}

func (query *QueryParamAnnotation) UserID(userID int64) *QueryParamAnnotation {
	query.userId = userID
	return query
}

func (query *QueryParamAnnotation) DashboardID(dashboardID int64) *QueryParamAnnotation {
	query.dashboardId = dashboardID
	return query
}

func (query *QueryParamAnnotation) PanelID(panelID int64) *QueryParamAnnotation {
	query.panelId = panelID
	return query
}

func (query *QueryParamAnnotation) AddTag(tag string) *QueryParamAnnotation {
	query.tags = append(query.tags, tag)
	return query
}

func (query *QueryParamAnnotation) Type(t string) *QueryParamAnnotation {
	query._type = t
	return query
}

func (query *QueryParamAnnotation) Limit(limit int64) *QueryParamAnnotation {
	query.limit = limit
	return query
}
