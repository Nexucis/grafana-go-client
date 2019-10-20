Grafana-go-client
=========================
[![CircleCI](https://circleci.com/gh/Nexucis/grafana-go-client.svg?style=shield)](https://circleci.com/gh/Nexucis/grafana-go-client) [![GitHub license](https://img.shields.io/badge/license-Apache2-blue.svg)](./LICENSE) [![codecov](https://codecov.io/gh/Nexucis/grafana-go-client/branch/master/graph/badge.svg)](https://codecov.io/gh/Nexucis/grafana-go-client)

A go client for the [Grafana API](http://docs.grafana.org/http_api/)


1. [Current Status](#overview) 
2. [Installation](#installation)
3. [Quickstart](#quickstart)
4. [Contributions](#contributions)
5. [Development](#development)
6. [License](#license)

## Overview
For the moment, the client only covers the Grafana API in the version 5. There is no guarantee that the client is compatible with previous or latest Grafana version.

### Versioning
This project uses the following version rules: 

```
X.Y.Z
```

Where : 
* X is the major version of Grafana supported by this project
* Y is the major version of this client. Be careful with this rule, you can have some breaking changes between two **Y** number. (even if usually it won't happen)
* Z is the minor version of this client. It will be increased when there are some bug fixes.

### Supported Version

| Grafana Version | Support Branch  |
| --------------- | --------------- |
| >= 5.0, < 6.0   | master          |


### Roadmap
Here is an overview of what is done and what is going to be done

#### API
*currently not well tested*

- [x] Alerting
- [x] Admin
- [x] Annotations
- [x] Authentication (key API)
- [ ] Dashboard ( not yet fully implemented, need to specify the dashboard struct)
   - [x] Dashboard Versions
   - [x] Dashboard Permissions
- [x] Data Source
- [x] Folder
   - [x] Folder Permissions
- [x] Folder/dashboard search
- [x] Organisation
   - [x] Current Org
   - [x] Manipulate Org as admin
- [x] Playlist
- [x] Snapshot
- [x] User
   - [x] Current User
   - [x] Manipulate User as admin
- [x] Team

## Installation
If you use [dep](https://golang.github.io/dep/) as dependency manager, you fire the following command:

```bash
dep ensure -add github.com/nexucis/grafana-go-client
```

If you are using go mod instead, you can perform the following command:

```bash
go get github.com/nexucis/grafana-go-client
```

## Quickstart

```go
package main

import (

	"github.com/golang/glog"
	"github.com/nexucis/grafana-go-client/grafanahttp"
	"github.com/nexucis/grafana-go-client/api"
)

func main() {
	rest, err := grafanahttp.NewWithURL("http://admin:admin@localhost:3000")
	if err != nil {
		glog.Fatal(err)
	}
	
	client := api.NewWithClient(rest)
	
	user,err := client.CurrentUser().Get()
	
	// do something with the information get from the api
}
```

## Contributions
Any contribution or suggestion would be really appreciated. Feel free to use the Issue section or to send a pull request.

## Development
All following tools are running by [circleci](https://circleci.com/gh/Nexucis/grafana-go-client), so in order to help you to improve your code and make easier your life, here it is how you can launch the tools with the correct parameter.

### Run unit test
You can run the unit test using make :

```bash
make test
```

This command only run the unit test which basically only the test in the http package. All test written in the package /api/v1 needs grafana. See [Run integration test](#run-integration-test) section. 

### Run integration test
If you want to launch the unit test, you need to have a local grafana instance which must be accessible through the url http://localhost:3000. A simply way to launch it, is to start the [corresponding container](https://hub.docker.com/r/grafana/grafana/) : 

```bash
docker run -d --name=grafana -p 3000:3000 grafana/grafana:5.4.3
```

Once Grafana is up, you can run the following command :

```bash
make integration-test
```


## License
This library is distributed under the [Apache 2.0](./LICENSE) license

