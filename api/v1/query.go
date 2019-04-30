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

import "github.com/nexucis/grafana-go-client/api/v1/types"

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
	annotationType types.AnnotationQueryType
	// number. Optional - default is 100. Max limit for results returned.
	limit int64
}

func (q *QueryParamAnnotation) From(from int64) *QueryParamAnnotation {
	q.from = from
	return q
}

func (q *QueryParamAnnotation) To(to int64) *QueryParamAnnotation {
	q.to = to
	return q
}

func (q *QueryParamAnnotation) AlertID(alertID int64) *QueryParamAnnotation {
	q.alertId = alertID
	return q
}

func (q *QueryParamAnnotation) UserID(userID int64) *QueryParamAnnotation {
	q.userId = userID
	return q
}

func (q *QueryParamAnnotation) DashboardID(dashboardID int64) *QueryParamAnnotation {
	q.dashboardId = dashboardID
	return q
}

func (q *QueryParamAnnotation) PanelID(panelID int64) *QueryParamAnnotation {
	q.panelId = panelID
	return q
}

func (q *QueryParamAnnotation) AddTag(tag string) *QueryParamAnnotation {
	q.tags = append(q.tags, tag)
	return q
}

func (q *QueryParamAnnotation) Type(queryType types.AnnotationQueryType) *QueryParamAnnotation {
	q.annotationType = queryType
	return q
}

func (q *QueryParamAnnotation) Limit(limit int64) *QueryParamAnnotation {
	q.limit = limit
	return q
}

type QueryParamAlert struct {
	// Limit response to alert for a specified panel on a dashboard.
	panelId int64
	// Limit response to X number of alerts.
	limit int64
	// Limit response to alerts having a name like this value
	query string
	// Limit response to alerts having a dashboard name like this value.
	dashboardQuery string
	// Return alerts with one or more of the alert states. You can specify multiple state
	states []types.AlertState
	// Limit response to alerts in specified dashboard(s). You can specify multiple dashboards
	dashboardIds []int64
	// Limit response to alerts of dashboards in specified folder(s).You can specify multiple folders
	folderIds []int64
	// Limit response to alerts of dashboards with specified tags. To do an “AND” filtering with multiple tags, specify the tags parameter multiple times.
	dashboardTags []string
}

func (q *QueryParamAlert) Query(query string) *QueryParamAlert {
	q.query = query
	return q
}

func (q *QueryParamAlert) PanelID(panelID int64) *QueryParamAlert {
	q.panelId = panelID
	return q
}

func (q *QueryParamAlert) Limit(limit int64) *QueryParamAlert {
	q.limit = limit
	return q
}

func (q *QueryParamAlert) AddAlertState(state types.AlertState) *QueryParamAlert {
	q.states = append(q.states, state)
	return q
}

func (q *QueryParamAlert) AddDashboardTag(dashboardTag string) *QueryParamAlert {
	q.dashboardTags = append(q.dashboardTags, dashboardTag)
	return q
}

func (q *QueryParamAlert) AddDashboardID(dashboardID int64) *QueryParamAlert {
	q.dashboardIds = append(q.dashboardIds, dashboardID)
	return q
}

func (q *QueryParamAlert) AddFolderID(folderId int64) *QueryParamAlert {
	q.folderIds = append(q.folderIds, folderId)
	return q
}

type QueryParameterUsers struct {
	// The number of user per page
	perPage int64
	// The number of the page querying
	page int64
	// Limit response to user having a similar name, login or email like this value.
	query string
}

func (q *QueryParameterUsers) Query(query string) *QueryParameterUsers {
	q.query = query
	return q
}

func (q *QueryParameterUsers) PerPage(perPage int64) *QueryParameterUsers {
	q.perPage = perPage
	return q
}

func (q *QueryParameterUsers) Page(page int64) *QueryParameterUsers {
	q.page = page
	return q
}

type QueryParameterTeams struct {
	perPage int64
	page    int64
	query   string
	name    string
}

func (q *QueryParameterTeams) Query(query string) *QueryParameterTeams {
	q.query = query
	return q
}

func (q *QueryParameterTeams) Name(name string) *QueryParameterTeams {
	q.name = name
	return q
}

func (q *QueryParameterTeams) PerPage(perPage int64) *QueryParameterTeams {
	q.perPage = perPage
	return q
}

func (q *QueryParameterTeams) Page(page int64) *QueryParameterTeams {
	q.page = page
	return q
}

type QueryParameterOrgs struct {
	name  string
	query string
}

func (q *QueryParameterOrgs) Query(query string) *QueryParameterOrgs {
	q.query = query
	return q
}

func (q *QueryParameterOrgs) Name(name string) *QueryParameterOrgs {
	q.name = name
	return q
}

type QueryParameterPlaylist struct {
	limit int64
	query string
}

func (q *QueryParameterPlaylist) Query(query string) *QueryParameterPlaylist {
	q.query = query
	return q
}

func (q *QueryParameterPlaylist) Limit(limit int64) *QueryParameterPlaylist {
	q.limit = limit
	return q
}

type QueryParameterSearch struct {
	// search by title
	query string
	// List of tags to search
	tags []string
	//  Type to search for, dash-folder or dash-db
	searchType types.SearchType
	// List of dashboard id’s to search
	dashboardIds []int64
	// List of folder id’s to search in for dashboards
	folderIds []int64
	// Flag indicating if only starred Dashboards should be returned
	starred    bool
	limit      int
	permission types.PermissionTypeAsString
}

func (q *QueryParameterSearch) Query(query string) *QueryParameterSearch {
	q.query = query
	return q
}

func (q *QueryParameterSearch) Tag(tag string) *QueryParameterSearch {
	q.tags = append(q.tags, tag)
	return q
}

func (q *QueryParameterSearch) Type(queryType types.SearchType) *QueryParameterSearch {
	q.searchType = queryType
	return q
}

func (q *QueryParameterSearch) DashboardIDS(id int64) *QueryParameterSearch {
	q.dashboardIds = append(q.dashboardIds, id)
	return q
}

func (q *QueryParameterSearch) FolderIDS(id int64) *QueryParameterSearch {
	q.folderIds = append(q.folderIds, id)
	return q
}

func (q *QueryParameterSearch) Starred(starred bool) *QueryParameterSearch {
	q.starred = starred
	return q
}

func (q *QueryParameterSearch) Limit(limit int) *QueryParameterSearch {
	q.limit = limit
	return q
}

func (q *QueryParameterSearch) Permission(permission types.PermissionTypeAsString) *QueryParameterSearch {
	q.permission = permission
	return q
}
