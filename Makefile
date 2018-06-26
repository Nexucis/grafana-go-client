GO                         ?= go
GOFMT                      ?= $(GO)fmt
pkgs                        = $$($(GO) list ./... | grep -v vendor)


all: install-dep test

.PHONY: install-dep
install-dep:
	@echo ">> install dependency"
	dep ensure

build:
	@echo ">> build all package"
	$(GO) build github.com/nexucis/grafana-go-client/http/...
	$(GO) build github.com/nexucis/grafana-go-client/api/...

.PHONY: verify
verify: checkformat checkstyle

.PHONY: checkstyle
checkstyle:
	@echo ">> checking code style"
	$(GOMETA) ./... --deadline=120s --vendor

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
