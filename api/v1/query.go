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
	states [] types.AlertState
	// Limit response to alerts in specified dashboard(s). You can specify multiple dashboards
	dashboardIds []int64
	// Limit response to alerts of dashboards in specified folder(s).You can specify multiple folders
	folderIds []int64
	// Limit response to alerts of dashboards with specified tags. To do an “AND” filtering with multiple tags, specify the tags parameter multiple times.
	dashboardTags []string
}

func (query *QueryParamAlert) Query(q string) *QueryParamAlert {
	query.query = q
	return query
}

func (query *QueryParamAlert) PanelID(panelID int64) *QueryParamAlert {
	query.panelId = panelID
	return query
}

func (query *QueryParamAlert) Limit(limit int64) *QueryParamAlert {
	query.limit = limit
	return query
}

func (query *QueryParamAlert) AddAlertState(state types.AlertState) *QueryParamAlert {
	query.states = append(query.states, state)
	return query
}

func (query *QueryParamAlert) AddDashboardTag(dashboardTag string) *QueryParamAlert {
	query.dashboardTags = append(query.dashboardTags, dashboardTag)
	return query
}

func (query *QueryParamAlert) AddDashboardID(dashboardID int64) *QueryParamAlert {
	query.dashboardIds = append(query.dashboardIds, dashboardID)
	return query
}

func (query *QueryParamAlert) AddFolderID(folderId int64) *QueryParamAlert {
	query.folderIds = append(query.folderIds, folderId)
	return query
}
