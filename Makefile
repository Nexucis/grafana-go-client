GO                         ?= go
GOCI                       ?= golangci-lint
GOFMT                      ?= $(GO)fmt
pkgs                        = $$($(GO) list ./... | grep -v vendor)


all: build test

build:
	@echo ">> build all package"
	$(GO) build github.com/nexucis/grafana-go-client/grafanahttp/...
	$(GO) build github.com/nexucis/grafana-go-client/api/...

.PHONY: verify
verify: checkformat checkstyle

.PHONY: checkstyle
checkstyle:
	@echo ">> checking code style"
	$(GOCI) run -E goconst -E unconvert -E gosec -E golint -E unparam -E maligned -E gocyclo

.PHONY: checkformat
checkformat:
	@echo ">> checking code format"
	! $(GOFMT) -d $$(find . -path ./vendor -prune -o -name '*.go' -print) | grep '^' ;\

.PHONY: fmt
fmt:
	@echo ">> format code"
	$(GO) fmt $(pkgs)

.PHONY: test
test:
	@echo ">> running all tests"
	$(GO) test -v $(pkgs)

.PHONY: integration-test
integration-test:
	@echo ">> running all tests"
	$(GO) test ./api/v1 -integration