#!/usr/bin/env bash

set -e
echo "" > coverage.txt

GO111MODULE=on go test -race -coverprofile=profile.out -covermode=atomic ./grafanahttp/

if [ -f profile.out ]; then
   cat profile.out >> coverage.txt
   rm profile.out
fi

GO111MODULE=on go test -race -coverprofile=profile.out -covermode=atomic ./api -integration

if [ -f profile.out ]; then
   cat profile.out >> coverage.txt
   rm profile.out
fi

echo "Publishing go code coverage"
bash <(curl -s https://codecov.io/bash) -cF go
