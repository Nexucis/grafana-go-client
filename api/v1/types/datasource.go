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
	Id          int64       `json:"id"`
	OrgId       int64       `json:"orgId"`
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	TypeLogoUrl string      `json:"typeLogoUrl"`
	Access      string      `json:"access"`
	Url         string      `json:"url"`
	Password    string      `json:"password"`
	User        string      `json:"user"`
	Database    string      `json:"database"`
	BasicAuth   bool        `json:"basicAuth"`
	IsDefault   bool        `json:"isDefault"`
	JsonData    interface{} `json:"jsonData,omitempty"`
	ReadOnly    bool        `json:"readOnly"`
}

// Also acts as api DTO
type AddDataSource struct {
	Name              string            `json:"name" binding:"Required"`
	Type              string            `json:"type" binding:"Required"`
	Access            string            `json:"access" binding:"Required"`
	Url               string            `json:"url"`
	Password          string            `json:"password"`
	Database          string            `json:"database"`
	User              string            `json:"user"`
	BasicAuth         bool              `json:"basicAuth"`
	BasicAuthUser     string            `json:"basicAuthUser"`
	BasicAuthPassword string            `json:"basicAuthPassword"`
	WithCredentials   bool              `json:"withCredentials"`
	IsDefault         bool              `json:"isDefault"`
	JsonData          interface{}       `json:"jsonData"`
	SecureJsonData    map[string]string `json:"secureJsonData"`
	ReadOnly          bool              `json:"readOnly"`
}

type WriteDataSourceResponse struct {
	ID         int64      `json:"id"`
	Message    string     `json:"message"`
	Name       string     `json:"name"`
	Datasource DataSource `json:"datasource"`
}

// Also acts as api DTO
type UpdateDataSource struct {
	Name              string            `json:"name" binding:"Required"`
	Type              string            `json:"type" binding:"Required"`
	Access            string            `json:"access" binding:"Required"`
	Url               string            `json:"url"`
	Password          string            `json:"password"`
	User              string            `json:"user"`
	Database          string            `json:"database"`
	BasicAuth         bool              `json:"basicAuth"`
	BasicAuthUser     string            `json:"basicAuthUser"`
	BasicAuthPassword string            `json:"basicAuthPassword"`
	WithCredentials   bool              `json:"withCredentials"`
	IsDefault         bool              `json:"isDefault"`
	JsonData          interface{}       `json:"jsonData"`
	SecureJsonData    map[string]string `json:"secureJsonData"`
	Version           int               `json:"version"`
	ReadOnly          bool              `json:"readOnly"`
}
