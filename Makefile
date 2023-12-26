GOHOSTOS:=$(shell go env GOHOSTOS)
# The short Git commit hash
SHORT_COMMIT := $(shell git rev-parse --short HEAD)
BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD | tr '/' '-')
# The Git commit hash
COMMIT := $(shell git rev-parse HEAD)
# The tag of the current commit, otherwise empty
VERSION := $(shell git describe --tags --abbrev=2 --match "v*" 2>/dev/null)
# Image tag: if image tag is not set, set it with version (or short commit if empty)
ifeq (${IMAGE_TAG},)
IMAGE_TAG := ${VERSION}
endif

ifeq (${IMAGE_TAG},)
IMAGE_TAG := ${SHORT_COMMIT}
endif
CONTAINER_REGISTRY=
# Name of the cover profile
COVER_PROFILE := coverage.txt
# Disable go sum database lookup for private repos
GOPRIVATE=
# OS
UNAME := $(shell uname)

# Used when building within docker
GOARCH := $(shell go env GOARCH)

API_PROTO_FILES=$(shell find api -name *.proto)
INTERNAL_PROTO_FILES=$(shell find app -name *.proto)
ifeq ($(GOHOSTOS), darwin)
	wireCmd=xargs -I F sh -c 'cd F && echo && wire'
else
	wireCmd=xargs -i sh -c 'cd {} && echo && wire'
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# generate app proto
config:
	for var in $(INTERNAL_PROTO_FILES); do \
        protoc --proto_path=./app \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./app \
	       $$var; \
    done

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
		   --go-errors_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:./api \
	       $(API_PROTO_FILES)
#    		--openapiv2_out=./api 
#     		--openapiv2_opt logtostderr=true 

.PHONY: build
# build
build:
	mkdir -p deploy/bin/ && go build -ldflags "-X cmd.Version=$(VERSION)" -o ./deploy/bin/ ./...

.PHONY: generate
# generate wire_gen.go
wire:
	find app  -mindepth 2 -maxdepth 2 | grep cmd | $(wireCmd)

export DOCKER_BUILDKIT := 1

names=$(shell find app -name main.go|xargs -I X dirname X)
.PHONY: build-docker-production
build-docker-production:
	for x in $(names); do \
		docker build -f Dockerfile  \
			--build-arg TARGET=./$$x \
			--build-arg COMMIT=$(COMMIT)  \
			--build-arg VERSION=$(IMAGE_TAG) \
			--build-arg GOARCH=$(GOARCH) \
			--target production \
			--secret id=git_creds,env=GITHUB_CREDS --build-arg GOPRIVATE=$(GOPRIVATE) \
			--label "git_commit=$(COMMIT)" --label "git_tag=${IMAGE_TAG}" \
			-t "$(CONTAINER_REGISTRY)$$x:latest" \
			-t "$(CONTAINER_REGISTRY)$$x:$(IMAGE_TAG)" . ;\
	done

.PHONY: build-docker-debug
build-docker-debug:
	for x in $(names); do \
		docker build -f Dockerfile  \
			--build-arg TARGET=./app/$$x \
			--build-arg COMMIT=$(COMMIT)  \
			--build-arg VERSION=$(IMAGE_TAG) \
			--build-arg GOARCH=$(GOARCH) \
			--target debug \
			--secret id=git_creds,env=GITHUB_CREDS --build-arg GOPRIVATE=$(GOPRIVATE) \
			--label "git_commit=$(COMMIT)" --label "git_tag=${IMAGE_TAG}" \
			-t "$(CONTAINER_REGISTRY)$$x-debug:latest" \
			-t "$(CONTAINER_REGISTRY)$$x-debug:$(IMAGE_TAG)" . ;\
	done

.PHONY: all
all: init api wire build build-docker-production

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
