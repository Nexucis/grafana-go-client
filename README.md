Grafana-go-client
=========================
[![CircleCI](https://circleci.com/gh/Nexucis/grafana-go-client.svg?style=shield)](https://circleci.com/gh/Nexucis/grafana-go-client) [![GitHub license](https://img.shields.io/badge/license-Apache2-blue.svg)](./LICENSE) [![codecov](https://codecov.io/gh/Nexucis/grafana-go-client/branch/master/graph/badge.svg)](https://codecov.io/gh/Nexucis/grafana-go-client)

A go client for the [Grafana API](http://docs.grafana.org/http_api/)


1. [Current Status](#current-status) 
2. [Installation](#installation)
3. [Quickstart](#quickstart)
3. [Contributions](#contributions)
5. [License](#license)

# Current Status

:warning: this project is still under development. You can already use it, but you can encounter some bugs or problem s

For the moment, the client only covers the Grafana API in the version 5. There is no guarantee that the client is compatible with previous Grafana version

## Roadmap
Here is an overview of what is done and what is going to be done

### API

*currently not tested*

- [x] Alerting
- [x] Admin
- [x] Annotations
- [x] Authentication (key API)
- [ ] Dashboard
   - [ ] Dashboard Versions
   - [ ] Dashboard Permissions
- [x] Data Source
- [x] Folder
   - [x] Folder Permissions
- [ ] Folder/dashboard search
- [x] Organisation
   - [x] Current Org
   - [x] Manipulate Org as admin
- [ ] Plugin
- [x] Playlist
- [ ] Snapshot
- [x] User
   - [x] Current User
   - [x] Manipulate User as admin
- [x] Team

### Retro compatibility

- [ ] Grafana v4
- [ ] Grafana v3 (maybe)

# Installation
If you use [dep](https://golang.github.io/dep/) as dependency manager, you fire the following command:

```bash
dep ensure -add github.com/nexucis/grafana-go-client
```

# Quickstart

```go
package main

import (

	"github.com/golang/glog"
	"github.com/nexucis/grafana-go-client/http"
	"github.com/nexucis/grafana-go-client/api"
)

func main() {
	rest, err := http.NewWithUrl("http://admin:admin@localhost:3000")
	if err != nil {
		glog.Fatal(err)
	}
	
	client := api.NewWithClient(rest)
	
	user := client.V1().CurrentUser().Get()
	
	// do something with the information get from the api
}
```

# Contributions

Any contribution or suggestion would be really appreciated. Feel free to use the Issue section or to send a pull request.

# License

This library is distributed under the [Apache 2.0](./LICENSE) license

