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

type DataSource struct {
	BasicAuth   bool        `json:"basicAuth"`
	IsDefault   bool        `json:"isDefault"`
	ReadOnly    bool        `json:"readOnly"`
	ID          int64       `json:"id"`
	OrgID       int64       `json:"orgId"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	TypeLogoURL string      `json:"typeLogoUrl"`
	Access      string      `json:"access"`
	URL         string      `json:"url"`
	Password    string      `json:"password"`
	User        string      `json:"user"`
	Database    string      `json:"database"`
	JSONData    interface{} `json:"jsonData,omitempty"`
}

// Also acts as api DTO
type AddDataSource struct {
	BasicAuth         bool              `json:"basicAuth"`
	WithCredentials   bool              `json:"withCredentials"`
	IsDefault         bool              `json:"isDefault"`
	ReadOnly          bool              `json:"readOnly"`
	Name              string            `json:"name" binding:"Required"`
	Type              string            `json:"type" binding:"Required"`
	Access            string            `json:"access" binding:"Required"`
	URL               string            `json:"url"`
	Password          string            `json:"password"`
	Database          string            `json:"database"`
	User              string            `json:"user"`
	BasicAuthUser     string            `json:"basicAuthUser"`
	BasicAuthPassword string            `json:"basicAuthPassword"`
	JSONData          interface{}       `json:"jsonData"`
	SecureJSONData    map[string]string `json:"secureJsonData"`
}

type WriteDataSourceResponse struct {
	ID         int64      `json:"id"`
	Message    string     `json:"message"`
	Name       string     `json:"name"`
	Datasource DataSource `json:"datasource"`
}

// Also acts as api DTO
type UpdateDataSource struct {
	ReadOnly          bool              `json:"readOnly"`
	WithCredentials   bool              `json:"withCredentials"`
	IsDefault         bool              `json:"isDefault"`
	BasicAuth         bool              `json:"basicAuth"`
	Version           int               `json:"version"`
	Name              string            `json:"name" binding:"Required"`
	Type              string            `json:"type" binding:"Required"`
	Access            string            `json:"access" binding:"Required"`
	URL               string            `json:"url"`
	Password          string            `json:"password"`
	User              string            `json:"user"`
	Database          string            `json:"database"`
	BasicAuthUser     string            `json:"basicAuthUser"`
	BasicAuthPassword string            `json:"basicAuthPassword"`
	JSONData          interface{}       `json:"jsonData"`
	SecureJSONData    map[string]string `json:"secureJsonData"`
}
