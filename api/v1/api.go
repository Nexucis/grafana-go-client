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

import "github.com/nexucis/grafana-go-client/http"

type APIInterface interface {
	Admin() AdminInterface
	Alerts() AlertInterface
	AlertNotifications() AlertNotificationInterface
	Annotations() AnnotationInterface
	CurrentUser() CurrentUserInterface
	Dashboards()
	Datasources()
	Folders()
	Keys() KeyInterface
	Organisation()
	Preferences()
	Search()
	Snapshots()
	Teams()
	Users()
}

type api struct {
	APIInterface
	client *http.RESTClient
}

func NewWithClient(client *http.RESTClient) APIInterface {
	return &api{
		client: client,
	}
}

func (a *api) Admin() AdminInterface {
	return newAdmin(a.client)
}

func (a *api) Alerts() AlertInterface {
	return newAlert(a.client)
}

func (a *api) AlertNotifications() AlertNotificationInterface {
	return newAlertNotification(a.client)
}

func (a *api) Annotations() AnnotationInterface {
	return newAnnotation(a.client)
}

func (a *api) CurrentUser() CurrentUserInterface {
	return newCurrentUser(a.client)
}

func (a *api) Keys() KeyInterface {
	return newKey(a.client)
}
