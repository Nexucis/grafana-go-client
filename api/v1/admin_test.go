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

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAdmin_GetSettings(t *testing.T) {
	if !*integration {
		// test is ignored
		t.Log("test is ignored")
		return
	}

	admin := initAdminTest(t)
	settings, err := admin.GetSettings()

	assert.Nil(t, err)
	assert.NotNil(t, settings)
}

func initAdminTest(t *testing.T) AdminInterface {
	httpClient, err := getRestClientWithBasicAuth()
	assert.Nil(t, err)
	return newAdmin(httpClient)
}
