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
	"fmt"
	"net/url"
	"testing"

	"github.com/nexucis/grafana-go-client/api/types"
	"github.com/stretchr/testify/assert"
)

func TestQueryParamAnnotation_GetValues(t *testing.T) {
	testSuites := []struct {
		title  string
		query  *QueryParamAnnotation
		result url.Values
	}{
		{
			title: "test with all parameter",
			query: &QueryParamAnnotation{
				AlertID:        4532185,
				AnnotationType: types.AnnSearchByAlert,
				DashboardID:    48653231,
				From:           1558256013,
				Limit:          10,
				PanelID:        456321,
				Tags: []string{
					"my",
					"test",
					"tag",
				},
				To:     1558256013,
				UserID: 45328,
			},
			result: url.Values{
				"alertId":     []string{"4532185"},
				"type":        []string{"alert"},
				"dashboardId": []string{"48653231"},
				"from":        []string{"1558256013"},
				"limit":       []string{"10"},
				"panelId":     []string{"456321"},
				"tags":        []string{"my", "test", "tag"},
				"to":          []string{"1558256013"},
				"userId":      []string{"45328"},
			},
		},
	}

	for _, testSuite := range testSuites {
		assert.Equal(t, testSuite.result, testSuite.query.GetValues(), fmt.Sprintf("error in test %s", testSuite.title))
	}
}

func TestQueryParamAlert_GetValues(t *testing.T) {
	testSuites := []struct {
		title  string
		query  *QueryParamAlert
		result url.Values
	}{
		{
			title: "test with all parameter",
			query: &QueryParamAlert{
				PanelID:        456321,
				Limit:          10,
				Query:          "my-query",
				DashboardQuery: "dashboard-name",
				States: []types.AlertState{
					types.AlertStateOk,
					types.AlertStateAlerting,
				},
				DashboardIDs: []int64{56432, 156584},
				FolderIDs:    []int64{79665, 112321, 887543},
				DashboardTags: []string{
					"my",
					"test",
					"tag",
				},
			},
			result: url.Values{
				"panelId":        []string{"456321"},
				"limit":          []string{"10"},
				"query":          []string{"my-query"},
				"dashboardQuery": []string{"dashboard-name"},
				"state":          []string{"ok", "alerting"},
				"dashboardId":    []string{"56432", "156584"},
				"folderId":       []string{"79665", "112321", "887543"},
				"dashboardTag":   []string{"my", "test", "tag"},
			},
		},
	}

	for _, testSuite := range testSuites {
		assert.Equal(t, testSuite.result, testSuite.query.GetValues(), fmt.Sprintf("error in test %s", testSuite.title))
	}
}

func TestQueryParameterUsers_GetValues(t *testing.T) {
	testSuites := []struct {
		title  string
		query  *QueryParameterUsers
		result url.Values
	}{
		{
			title: "test with all parameter",
			query: &QueryParameterUsers{
				PerPage: 50,
				Page:    10,
				Query:   "user-name",
			},
			result: url.Values{
				"perpage": []string{"50"},
				"page":    []string{"10"},
				"query":   []string{"user-name"},
			},
		},
	}

	for _, testSuite := range testSuites {
		assert.Equal(t, testSuite.result, testSuite.query.GetValues(), fmt.Sprintf("error in test %s", testSuite.title))
	}
}

func TestQueryParameterTeams_GetValues(t *testing.T) {
	testSuites := []struct {
		title  string
		query  *QueryParameterTeams
		result url.Values
	}{
		{
			title: "test with all parameter",
			query: &QueryParameterTeams{
				PerPage: 50,
				Page:    10,
				Query:   "contains-team-name",
				Name:    "exact-team-name",
			},
			result: url.Values{
				"perpage": []string{"50"},
				"page":    []string{"10"},
				"query":   []string{"contains-team-name"},
				"name":    []string{"exact-team-name"},
			},
		},
	}

	for _, testSuite := range testSuites {
		assert.Equal(t, testSuite.result, testSuite.query.GetValues(), fmt.Sprintf("error in test %s", testSuite.title))
	}
}

func TestQueryParameterOrgs_GetValues(t *testing.T) {
	testSuites := []struct {
		title  string
		query  *QueryParameterOrgs
		result url.Values
	}{
		{
			title: "test with all parameter",
			query: &QueryParameterOrgs{
				Query: "contains-org-name",
				Name:  "exact-org-name",
			},
			result: url.Values{
				"query": []string{"contains-org-name"},
				"name":  []string{"exact-org-name"},
			},
		},
	}

	for _, testSuite := range testSuites {
		assert.Equal(t, testSuite.result, testSuite.query.GetValues(), fmt.Sprintf("error in test %s", testSuite.title))
	}
}

func TestQueryParameterPlaylist_GetValues(t *testing.T) {
	testSuites := []struct {
		title  string
		query  *QueryParameterPlaylist
		result url.Values
	}{
		{
			title: "test with all parameter",
			query: &QueryParameterPlaylist{
				Query: "playlist-name",
				Limit: 45,
			},
			result: url.Values{
				"query": []string{"playlist-name"},
				"limit": []string{"45"},
			},
		},
	}

	for _, testSuite := range testSuites {
		assert.Equal(t, testSuite.result, testSuite.query.GetValues(), fmt.Sprintf("error in test %s", testSuite.title))
	}
}

func TestQueryParameterSearch_GetValues(t *testing.T) {
	testSuites := []struct {
		title  string
		query  *QueryParameterSearch
		result url.Values
	}{
		{
			title: "test with all parameter",
			query: &QueryParameterSearch{
				Limit: 10,
				Query: "my-query",
				Tags: []string{
					"my",
					"test",
					"tag",
				},
				SearchType:   types.SearchDashboardType,
				DashboardIDs: []int64{56432, 156584},
				FolderIDs:    []int64{79665, 112321, 887543},
				Starred:      false,
				Permission:   types.PermissionEditAsString,
			},
			result: url.Values{
				"limit":        []string{"10"},
				"query":        []string{"my-query"},
				"tag":          []string{"my", "test", "tag"},
				"type":         []string{"dash-db"},
				"dashboardIds": []string{"56432", "156584"},
				"folderIds":    []string{"79665", "112321", "887543"},
				"starred":      []string{"false"},
				"permission":   []string{"Edit"},
			},
		},
	}

	for _, testSuite := range testSuites {
		assert.Equal(t, testSuite.result, testSuite.query.GetValues(), fmt.Sprintf("error in test %s", testSuite.title))
	}
}
