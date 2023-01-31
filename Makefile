GO ?= go
OS = $(shell uname -s | tr A-Z a-z)
SERVICE_NAME = $(shell basename "$(PWD)")
ROOT = $(shell pwd)
export GOBIN = ${ROOT}/bin

PATH := $(PATH):$(GOBIN)

ENV_FILE = .env
EXPORT_ENV = export $(shell test -e ./$(ENV_FILE) || cp ./.env.example ./$(ENV_FILE) && grep -v '^#' ./$(ENV_FILE) | xargs -d '\n')

MIGRATE_VERSION = v4.15.2
MIGRATE = ${GOBIN}/migrate
MIGRATE_DOWNLOAD = (curl --progress-bar -fL -o $(MIGRATE).tar.gz https://github.com/golang-migrate/migrate/releases/download/$(MIGRATE_VERSION)/migrate.$(OS)-amd64.tar.gz; tar -xzf $(MIGRATE).tar.gz -C $(GOBIN); rm $(MIGRATE).tar.gz ; rm bin/LICENSE bin/README.md )
MIGRATE_CONFIG = -source file://migrations -database "${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=${DB_SSLMODE}"

TPARSE = $(GOBIN)/tparse
TPARSE_DOWNLOAD = $(GO) install github.com/mfridman/tparse@latest

LINT_VERSION = v1.50.1
LINT = ${GOBIN}/golangci-lint
LINT_DOWNLOAD = curl --progress-bar -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) $(LINT_VERSION)

VERSION_TAG = $(shell git describe --tags --abbrev=0 --always)
VERSION_COMMIT = $(shell git rev-parse --short HEAD)
VERSION_DATE = $(shell git show -s --format=%cI HEAD)
VERSION = -X main.versionTag=$(VERSION_TAG) -X main.versionCommit=$(VERSION_COMMIT)  -X main.versionDate=$(VERSION_DATE) -X main.serviceName=$(SERVICE_NAME)

BUILD_OUTPUT = ./bin/${SERVICE_NAME}

COMPILE_DAEMON = $(GOBIN)/CompileDaemon
COMPILE_DAEMON_DOWNLOAD = $(GO) install github.com/githubnemo/CompileDaemon@latest

.PHONY: help
help: ## Display this help message
	@ cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: env
env: ## Check and create env file from .env.example
	@ $(EXPORT_ENV) && env

.PHONY: migrate
migrate: ## Check and downlowe migrate file
	@ test -e $(MIGRATE) || $(MIGRATE_DOWNLOAD)
	@ $(MIGRATE) --version

.PHONY: migrate-create
migrate-create:migrate ## Create new migration file
	@ $(MIGRATE) create -ext sql -seq -dir ./migrations $(NAME)

.PHONY: migrate-up
migrate-up:migrate ## Apply all up migrations
	@ $(MIGRATE) $(MIGRATE_CONFIG) up

.PHONY:	migrate-down
migrate-down:migrate ## Apply all down migrations
	@ $(MIGRATE) $(MIGRATE_CONFIG) down $(STEP)

.PHONY: migrate-recreate
migrate-recreate:migrate ## Apply down and up migrations
	@ test -e $(MIGRATE) || $(MIGRATE_DOWNLOAD)
	@ $(MIGRATE) $(MIGRATE_CONFIG) down -all && $(MIGRATE) $(MIGRATE_CONFIG) up

.PHONY: mod
mod: ## Get dependency packages
	@ $(GO) mod tidy

.PHONY: test-base
test-base: ## Check and downlowe tparse file
	@ test -e $(TPARSE) || $(TPARSE_DOWNLOAD)
	@ $(TPARSE) --version

.PHONY: test
test:test-base ## Run data race detector
	@ $(GO) test -timeout 10m -short ./... -json -cover | $(TPARSE) -all -smallscreen

.PHONY: coverage
coverage: ## Check coverage test code of sample https://penkovski.com/post/gitlab-golang-test-coverage/
	@ $(GO) test -timeout 10m ./internal/... -coverprofile=coverage.out
	@ $(GO) tool cover -func=coverage.out
	@ $(GO) tool cover -html=coverage.out -o coverage.html;

.PHONY: lint
lint: ## Lint the files
	@ test -e $(LINT) || $(LINT_DOWNLOAD)
	@ $(LINT) version
	@ $(LINT) --timeout 10m run

.PHONY: build
build: ## Build the project
	@ $(GO) env -w GO111MODULE=on
	@ $(GO) env -w CGO_ENABLED=0
	@ $(GO) build -o ${BUILD_OUTPUT} ./cmd/*.go

.PHONY: run
run: ## Run as development reload if code changes
	@ test -e $(COMPILE_DAEMON) || $(COMPILE_DAEMON_DOWNLOAD)
	@ $(COMPILE_DAEMON) --build="make build"  --command="$(GOBIN)/$(SERVICE_NAME)"
