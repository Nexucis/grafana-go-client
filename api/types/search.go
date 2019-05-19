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

type SearchType string

const (
	SearchDashboardType SearchType = "dash-db"
	SearchDFolderType   SearchType = "dash-folder"
)

type SearchResult struct {
	ID          int64      `json:"id"`
	UID         string     `json:"uid"`
	Title       string     `json:"title"`
	URI         string     `json:"uri"`
	URL         string     `json:"url"`
	Type        SearchType `json:"type"`
	Tags        []string   `json:"tags"`
	IsStarred   bool       `json:"isStarred"`
	FolderID    int64      `json:"folderId,omitempty"`
	FolderUID   string     `json:"folderUid,omitempty"`
	FolderTitle string     `json:"folderTitle,omitempty"`
	FolderURL   string     `json:"folderUrl,omitempty"`
}
