#!/usr/bin/env bash

set -e
echo "" > coverage.txt

go test -race -coverprofile=profile.out -covermode=atomic http/

if [ -f profile.out ]; then
   cat profile.out >> coverage.txt
   rm profile.out
fi

go test -race -coverprofile=profile.out -covermode=atomic ./api/v1 -integration

if [ -f profile.out ]; then
   cat profile.out >> coverage.txt
   rm profile.out
fi

echo "Publishing go code coverage"
bash <(curl -s https://codecov.io/bash) -cF go