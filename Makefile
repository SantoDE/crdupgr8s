.PHONY: check build

TAG_NAME := $(shell git tag -l --contains HEAD)
SHA := $(shell git rev-parse --short HEAD)
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%R:%M:%S%p')

default: check build

build:
	@echo Version: $(VERSION) $(BUILD_DATE)
	go build -v -ldflags '-X "github.com/SantoDE/crdupgr8s/v3/meta.version=${VERSION}" -X "github.com/SantoDE/crdupgr8s/v3/meta.commit=${SHA}" -X "github.com/ldez/SantoDE/crdupgr8s/meta.date=${BUILD_DATE}"' -trimpath

check:
	golangci-lint run