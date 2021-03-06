EXECUTABLE ?= github-webhooks-exporter
IMAGE ?= quay.io/cbrgm/$(EXECUTABLE)
GO := CGO_ENABLED=0 go
DATE := $(shell date -u '+%FT%T%z')

LDFLAGS += -X main.Version=$(shell git describe --tags --abbrev=0)
LDFLAGS += -X main.Revision=$(shell git rev-parse --short=7 HEAD)
LDFLAGS += -X "main.BuildDate=$(DATE)"
LDFLAGS += -extldflags '-static'

PACKAGES = $(shell go list ./...)

.PHONY: all
all: build

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -rf ./bin/

.PHONY: format
format: go/fmt

.PHONY: go/fmt
go/fmt:
	$(GO) fmt $(PACKAGES)

.PHONY: go/lint
go/lint:
	golangci-lint run

.PHONY: test
test:
	@for PKG in $(PACKAGES); do $(GO) test -cover $$PKG || exit 1; done;

.PHONY: build
build: \
	cmd/github-webhooks-exporter

.PHONY: cmd/github-webhooks-exporter
cmd/github-webhooks-exporter:
	mkdir -p bin
	$(GO) build -v -ldflags '-w $(LDFLAGS)' -o ./bin/$(EXECUTABLE)

.PHONY: release
release:
	GOOS=windows GOARCH=amd64 go build -v -ldflags '-w $(LDFLAGS)' -o ./bin/$(EXECUTABLE)_windows_amd64
	GOOS=linux GOARCH=amd64 go build -v -ldflags '-w $(LDFLAGS)' -o ./bin/$(EXECUTABLE)_linux_amd64
	GOOS=darwin GOARCH=amd64 go build -v -ldflags '-w $(LDFLAGS)' -o ./bin/$(EXECUTABLE)_darwin_amd64

.PHONY: container
container:
	podman build -t $(IMAGE):$(shell git describe --tags --abbrev=0) .
	podman build -t $(IMAGE):latest .

.PHONY: container-push
container-push: container
	podman push $(IMAGE):$(shell git describe --tags --abbrev=0)
    podman push $(IMAGE):latest
