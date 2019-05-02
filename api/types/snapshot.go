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

type DashboardSnapshot struct {
	Dashboard interface{} `json:"dashboard" binding:"Required"`
	Name      string      `json:"name"`
	Expires   int64       `json:"expires"`
	// these are passed when storing an external snapshot ref
	External  bool   `json:"external"`
	Key       string `json:"key"`
	DeleteKey string `json:"deleteKey"`
}

type CreateSnaphostResponse struct {
	Key       string `json:"key"`
	DeleteKey string `json:"deleteKey"`
	URL       string `json:"url"`
	DeleteURL string `json:"deleteUrl"`
}

type SharedOptionSnaphost struct {
	ExternalSnapshotURL  string `json:"externalSnapshotURL"`
	ExternalSnapshotName string `json:"externalSnapshotName"`
	ExternalEnabled      string `json:"externalEnabled"`
}
