

REPO    := mydata

GOOS    := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
GOARCH  := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))
GOENV   := GO111MODULE=on CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH)
GO      := $(GOENV) go
GOBUILD := $(GO) build -trimpath
GORUN   := $(GO) run
SHELL   := /usr/bin/env bash

COMMIT  := $(shell git describe --always --no-match --tags --dirty="-dev")
BUILDTS := $(shell date '+%Y-%m-%d %H:%M:%S')
GITHASH := $(shell git rev-parse HEAD)
GITREF  := $(shell git rev-parse --abbrev-ref HEAD)
GOVER   := $(shell go version)

LDFLAGS := -w -s
LDFLAGS += -X "$(REPO)/internal/version.ReleaseVersion=$(COMMIT)"
LDFLAGS += -X "$(REPO)/internal/version.BuildTS=$(BUILDTS)"
LDFLAGS += -X "$(REPO)/internal/version.GitHash=$(GITHASH)"
LDFLAGS += -X "$(REPO)/internal/version.GitBranch=$(GITREF)"
LDFLAGS += -X "$(REPO)/internal/version.GoVersion=$(GOVER)"

all: mydata

mydata:
	$(GOBUILD) -ldflags '$(LDFLAGS)' -o ./dist/mydata main.go
