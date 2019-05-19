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

package api

import (
	"net/url"
	"strconv"

	"github.com/nexucis/grafana-go-client/api/types"
	"github.com/nexucis/grafana-go-client/grafanahttp"
)

type QueryParamAnnotation struct {
	grafanahttp.QueryInterface
	// epoch datetime in milliseconds. Optional.
	From int64
	// epoch datetime in milliseconds. Optional.
	To int64
	// number. Optional. Find annotations created by a specific user
	UserID int64
	// number. Optional. Find annotations for a specified alert.
	AlertID int64
	// number. Optional. Find annotations that are scoped to a specific dashboard
	DashboardID int64
	// number. Optional. Find annotations that are scoped to a specific panel
	PanelID int64
	// string. Optional. Use this to filter global annotations. Global annotations are annotations from an annotation data source that are not connected specifically to a dashboard or panel
	Tags []string
	// string. Optional. alert|annotation Return alerts or user created annotations
	AnnotationType types.AnnotationQueryType
	// number. Optional - default is 100. Max limit for results returned.
	Limit int64
}

func (q *QueryParamAnnotation) GetValues() url.Values {
	values := make(url.Values)

	if q.From > 0 {
		values["from"] = append(values["from"], strconv.FormatInt(q.From, 10))
	}

	if q.To > 0 {
		values["to"] = append(values["to"], strconv.FormatInt(q.To, 10))
	}

	if q.UserID > 0 {
		values["userId"] = append(values["userId"], strconv.FormatInt(q.UserID, 10))
	}

	if q.AlertID > 0 {
		values["alertId"] = append(values["alertId"], strconv.FormatInt(q.AlertID, 10))
	}

	if q.DashboardID > 0 {
		values["dashboardId"] = append(values["dashboardId"], strconv.FormatInt(q.DashboardID, 10))
	}

	if q.PanelID > 0 {
		values["panelId"] = append(values["panelId"], strconv.FormatInt(q.PanelID, 10))
	}

	if q.Limit > 0 {
		values["limit"] = append(values["limit"], strconv.FormatInt(q.Limit, 10))
	}

	if len(q.AnnotationType) > 0 {
		values["type"] = append(values["type"], string(q.AnnotationType))
	}

	if len(q.Tags) > 0 {
		values["tags"] = append(values["tags"], q.Tags...)
	}

	return values
}

type QueryParamAlert struct {
	grafanahttp.QueryInterface
	// Limit response to alert for a specified panel on a dashboard.
	PanelID int64
	// Limit response to X number of alerts.
	Limit int64
	// Limit response to alerts having a name like this value
	Query string
	// Limit response to alerts having a dashboard name like this value.
	DashboardQuery string
	// Return alerts with one or more of the alert states. You can specify multiple state
	States []types.AlertState
	// Limit response to alerts in specified dashboard(s). You can specify multiple dashboards
	DashboardIDs []int64
	// Limit response to alerts of dashboards in specified folder(s).You can specify multiple folders
	FolderIDs []int64
	// Limit response to alerts of dashboards with specified tags. To do an “AND” filtering with multiple tags, specify the tags parameter multiple times.
	DashboardTags []string
}

func (q *QueryParamAlert) GetValues() url.Values {
	values := make(url.Values)
	if q.PanelID > 0 {
		values["panelId"] = append(values["panelId"], strconv.FormatInt(q.PanelID, 10))
	}

	if len(q.DashboardQuery) > 0 {
		values["dashboardQuery"] = append(values["dashboardQuery"], q.DashboardQuery)
	}

	if len(q.Query) > 0 {
		values["query"] = append(values["query"], q.Query)
	}

	if q.Limit > 0 {
		values["limit"] = append(values["limit"], strconv.FormatInt(q.Limit, 10))
	}

	if len(q.States) > 0 {
		for _, state := range q.States {
			values["state"] = append(values["state"], string(state))
		}
	}

	if len(q.DashboardIDs) > 0 {
		for _, dashboardID := range q.DashboardIDs {
			values["dashboardId"] = append(values["dashboardId"], strconv.FormatInt(dashboardID, 10))
		}
	}

	if len(q.FolderIDs) > 0 {
		for _, folderID := range q.FolderIDs {
			values["folderId"] = append(values["folderId"], strconv.FormatInt(folderID, 10))
		}
	}

	if len(q.DashboardTags) > 0 {
		values["dashboardTag"] = append(values["dashboardTag"], q.DashboardTags...)
	}

	return values
}

type QueryParameterUsers struct {
	grafanahttp.QueryInterface
	// The number of user per page
	PerPage int64
	// The number of the page querying
	Page int64
	// Limit response to user having a similar name, login or email like this value.
	Query string
}

func (q *QueryParameterUsers) GetValues() url.Values {
	values := make(url.Values)

	if q.Page > 0 {
		values["page"] = append(values["page"], strconv.FormatInt(q.Page, 10))
	}

	if q.PerPage > 0 {
		values["perpage"] = append(values["perpage"], strconv.FormatInt(q.PerPage, 10))
	}

	if len(q.Query) > 0 {
		values["query"] = append(values["query"], q.Query)
	}

	return values
}

type QueryParameterTeams struct {
	grafanahttp.QueryInterface
	// The number of team per page
	PerPage int64
	// The number of the page querying
	Page int64
	// Limit response to team having a name like this value.
	Query string
	// Limit response to the team having a name that matches exactly this value
	Name string
}

func (q *QueryParameterTeams) GetValues() url.Values {
	values := make(url.Values)

	if q.PerPage > 0 {
		values["perpage"] = append(values["perpage"], strconv.FormatInt(q.PerPage, 10))
	}

	if len(q.Name) > 0 {
		values["name"] = append(values["name"], q.Name)
	}

	if len(q.Query) > 0 {
		values["query"] = append(values["query"], q.Query)
	}

	if q.Page > 0 {
		values["page"] = append(values["page"], strconv.FormatInt(q.Page, 10))
	}

	return values
}

type QueryParameterOrgs struct {
	grafanahttp.QueryInterface
	Name  string
	Query string
}

func (q *QueryParameterOrgs) GetValues() url.Values {
	values := make(url.Values)

	if len(q.Query) > 0 {
		values["query"] = append(values["query"], q.Query)
	}

	if len(q.Name) > 0 {
		values["name"] = append(values["name"], q.Name)
	}

	return values
}

type QueryParameterPlaylist struct {
	grafanahttp.QueryInterface
	// Limit response to X number of playlist.
	Limit int64
	// Limit response to playlist having a name like this value.
	Query string
}

func (q *QueryParameterPlaylist) GetValues() url.Values {
	values := make(url.Values)

	if len(q.Query) > 0 {
		values["query"] = append(values["query"], q.Query)
	}

	if q.Limit > 0 {
		values["limit"] = append(values["limit"], strconv.FormatInt(q.Limit, 10))
	}

	return values
}

type QueryParameterSearch struct {
	grafanahttp.QueryInterface
	// search by title
	Query string
	// List of tags to search
	Tags []string
	//  Type to search for, dash-folder or dash-db
	SearchType types.SearchType
	// List of dashboard id’s to search
	DashboardIDs []int64
	// List of folder id’s to search in for dashboards
	FolderIDs []int64
	// Flag indicating if only starred Dashboards should be returned
	Starred    bool
	Limit      int
	Permission types.PermissionTypeAsString
}

func (q *QueryParameterSearch) GetValues() url.Values {
	values := make(url.Values)

	if len(q.Query) > 0 {
		values["query"] = append(values["query"], q.Query)
	}

	if q.Tags != nil {
		values["tag"] = append(values["tag"], q.Tags...)
	}

	if len(q.SearchType) > 0 {
		values["type"] = append(values["type"], string(q.SearchType))
	}

	if len(q.DashboardIDs) > 0 {
		for _, id := range q.DashboardIDs {
			values["dashboardIds"] = append(values["dashboardIds"], strconv.FormatInt(id, 10))
		}
	}

	if len(q.FolderIDs) > 0 {
		for _, id := range q.FolderIDs {
			values["folderIds"] = append(values["folderIds"], strconv.FormatInt(id, 10))
		}
	}

	values["starred"] = append(values["starred"], strconv.FormatBool(q.Starred))

	if q.Limit > 0 {
		values["limit"] = append(values["limit"], strconv.Itoa(q.Limit))
	}

	if len(q.Permission) > 0 {
		values["permission"] = append(values["permission"], string(q.Permission))
	}

	return values
}
