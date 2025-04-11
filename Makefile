ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
TOOLS_DIR := .tools

DB=tenant_api
DEV_DB=${DB}_dev
DEV_URI="postgresql://root@crdb:26257/${DEV_DB}?sslmode=disable"

# Determine OS and ARCH for some tool versions.
OS := linux
ARCH := amd64

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	OS = darwin
endif

UNAME_P := $(shell uname -p)
ifneq ($(filter arm%,$(UNAME_P)),)
	ARCH = arm64
endif

# Tool Versions
COCKROACH_VERSION = v22.2.8

OS_VERSION = $(OS)
ifeq ($(OS),darwin)
OS_VERSION = darwin-10.9
ifeq ($(ARCH),arm64)
OS_VERSION = darwin-11.0
endif
endif

COCKROACH_VERSION_FILE = cockroach-$(COCKROACH_VERSION).$(OS_VERSION)-$(ARCH)
COCKROACH_RELEASE_URL = https://binaries.cockroachdb.com/$(COCKROACH_VERSION_FILE).tgz

# go files to be checked
GO_FILES=$(shell git ls-files '*.go')

# Targets

.PHONY: help
help: Makefile ## Print help.
	@grep -h "##" $(MAKEFILE_LIST) | grep -v grep | sed -e 's/:.*##/#/' | column -c 2 -t -s#

.PHONY: all
all: lint test  ## Lints and tests.

.PHONY: ci
ci: | dev-database golint test coverage  ## Setup dev database and run tests.

.PHONY: dev-database
dev-database: | vendor $(TOOLS_DIR)/cockroach  ## Initializes dev database "${DEV_DB}"
	@$(TOOLS_DIR)/cockroach sql -e "drop database if exists ${DEV_DB}"
	@$(TOOLS_DIR)/cockroach sql -e "create database ${DEV_DB}"
	@TENANTAPI_CRDB_URI="${DEV_URI}" go run main.go migrate up

.PHONY: generate
generate: | dev-database  ## Regenerate files.
	@echo Regenerating files...
	@PATH="$(ROOT_DIR)/$(TOOLS_DIR):$$PATH" \
		go generate ./...

.PHONY: test
test: | generate unit-test  ## Regenerate files and run unit tests.

.PHONY: unit-test
unit-test: | $(TOOLS_DIR)/cockroach  ## Runs unit tests.
	@echo Running unit tests...
	@PATH="$(ROOT_DIR)/$(TOOLS_DIR):$$PATH" \
		go test -timeout 30s -cover -short ./...

.PHONY: coverage
coverage: | $(TOOLS_DIR)/cockroach  ## Generates a test coverage report.
	@echo Generating coverage report...
	@PATH="$(ROOT_DIR)/$(TOOLS_DIR):$$PATH" \
		go test -timeout 30s ./... -coverprofile=coverage.out -covermode=atomic
	@PATH="$(ROOT_DIR)/$(TOOLS_DIR):$$PATH" \
		go tool cover -func=coverage.out
	@PATH="$(ROOT_DIR)/$(TOOLS_DIR):$$PATH" \
		go tool cover -html=coverage.out

.PHONY: lint
lint: golint  ## Runs all lint checks.

golint: | vendor  ## Runs Go lint checks.
	@echo Linting Go files...
	@go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint run --fix --timeout=5m

fixlint:
	@echo Fixing go imports
	@find . -type f -iname '*.go' | xargs go run golang.org/x/tools/cmd/goimports -w -local go.infratographer.com/tenant-api

vendor:  ## Downloads and tidies go modules.
	@go mod download
	@go mod tidy

# Tools setup
$(TOOLS_DIR):
	mkdir -p $(TOOLS_DIR)

$(TOOLS_DIR)/cockroach: | $(TOOLS_DIR)
	@echo "Downloading cockroach: $(COCKROACH_RELEASE_URL)"
	@curl --silent --fail "$(COCKROACH_RELEASE_URL)" \
		| tar -xz --strip-components 1 -C $(TOOLS_DIR) $(COCKROACH_VERSION_FILE)/cockroach

	$@ version

