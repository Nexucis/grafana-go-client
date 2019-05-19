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

type PlaylistItemType string

const (
	DashboardByTag = "dashboard_by_tag"
	DashboardByID  = "dashboard_by_id"
)

type SimplePlaylist struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Interval string `json:"interval"`
}

type Playlist struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Interval string         `json:"interval"`
	Items    []PlaylistItem `json:"items"`
}

type PlaylistItem struct {
	ID         int64            `json:"id"`
	PlaylistID int64            `json:"playlistid"`
	Type       PlaylistItemType `json:"type"`
	Title      string           `json:"title"`
	Value      string           `json:"value"`
	Order      int              `json:"order"`
}

type PlaylistDashboard struct {
	ID    int64  `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}
