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

type AlertState string

const (
	AlertStateAll      AlertState = "ALL"
	AlertStateNoData   AlertState = "no_data"
	AlertStatePaused   AlertState = "paused"
	AlertStateAlerting AlertState = "alerting"
	AlertStateOk       AlertState = "ok"
	AlertStatePending  AlertState = "pending"
)

type ResponseGetAlert struct {
	ID             int64       `json:"id"`
	DashboardID    int64       `json:"dashboardId"`
	DashboardUID   string      `json:"dashboardUid"`
	DashboardSlug  string      `json:"dashboardSlug"`
	PanelID        int64       `json:"panelId"`
	Name           string      `json:"name"`
	State          AlertState  `json:"state"`
	NewStateDate   time.Time   `json:"newStateDate"`
	EvalDate       time.Time   `json:"evalDate"`
	EvalData       interface{} `json:"evalData"`
	ExecutionError string      `json:"executionError"`
	URL            string      `json:"url"`
}

type ResponseGetStatesForDashboard struct {
	ID           int64      `json:"id"`
	DashboardID  int64      `json:"dashboardId"`
	PanelID      int64      `json:"panelId"`
	State        AlertState `json:"state"`
	NewStateDate time.Time  `json:"newStateDate"`
}

type PostAlertTest struct {
	// Dashboard shall at least contains a key id with type int64
	Dashboard interface{} `json:"dashboard" binding:"Required"`
	PanelID   int64       `json:"panelId" binding:"Required"`
}

type ResponsePostAlertTest struct {
	Firing         bool                  `json:"firing"`
	State          AlertState            `json:"state"`
	ConditionEvals string                `json:"conditionEvals"`
	TimeMs         string                `json:"timeMs"`
	Error          string                `json:"error,omitempty"`
	EvalMatches    []*EvalMatch          `json:"matches,omitempty"`
	Logs           []*AlertTestResultLog `json:"logs,omitempty"`
}

type AlertTestResultLog struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EvalMatch struct {
	Tags   map[string]string `json:"tags,omitempty"`
	Metric string            `json:"metric"`
	Value  NullFloat64       `json:"value"`
}

// NullFloat64 represents a float64 that may be null.
type NullFloat64 struct {
	Float64 float64
	Valid   bool // Valid is true if Float64 is not NULL
}

type PostPauseAlert struct {
	AlertID int64 `json:"alertId"`
	Paused  bool  `json:"paused"`
}
