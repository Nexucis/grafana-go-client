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
	"github.com/nexucis/grafana-go-client/http"
)

type ClientInterface interface {
	RESTClient() *http.RESTClient
	Admin() AdminInterface
	Alerts() AlertInterface
	AlertNotifications() AlertNotificationInterface
	Annotations() AnnotationInterface
	CurrentUser() CurrentUserInterface
	CurrentOrganisation() CurrentOrgInterface
	Dashboards() DashboardInterface
	DataSources() DataSourceInterface
	Folders() FolderInterface
	Keys() KeyInterface
	Organisations() OrganisationsInterface
	Playlist() PlaylistInterface
	Search() SearchInterface
	Snapshots() SnapshotInterface
	Teams() TeamInterface
	Users() UsersInterface
}

type client struct {
	restClient *http.RESTClient
}

func NewWithClient(restClient *http.RESTClient) ClientInterface {
	return &client{
		restClient: restClient,
	}
}

func (c *client) RESTClient() *http.RESTClient {
	return c.restClient
}

func (c *client) Admin() AdminInterface {
	return newAdmin(c.restClient)
}

func (c *client) Alerts() AlertInterface {
	return newAlert(c.restClient)
}

func (c *client) AlertNotifications() AlertNotificationInterface {
	return newAlertNotification(c.restClient)
}

func (c *client) Annotations() AnnotationInterface {
	return newAnnotation(c.restClient)
}

func (c *client) CurrentUser() CurrentUserInterface {
	return newCurrentUser(c.restClient)
}

func (c *client) CurrentOrganisation() CurrentOrgInterface {
	return newCurrentOrg(c.restClient)
}

func (c *client) Dashboards() DashboardInterface {
	return newDashboard(c.restClient)
}

func (c *client) DataSources() DataSourceInterface {
	return newDataSource(c.restClient)
}

func (c *client) Folders() FolderInterface {
	return newFolder(c.restClient)
}

func (c *client) Keys() KeyInterface {
	return newKey(c.restClient)
}

func (c *client) Playlist() PlaylistInterface {
	return newPlaylist(c.restClient)
}

func (c *client) Search() SearchInterface {
	return newSearch(c.restClient)
}

func (c *client) Snapshots() SnapshotInterface {
	return newSnapshot(c.restClient)
}

func (c *client) Organisations() OrganisationsInterface {
	return newOrgs(c.restClient)
}

func (c *client) Teams() TeamInterface {
	return newTeam(c.restClient)
}

func (c *client) Users() UsersInterface {
	return newUsers(c.restClient)
}
