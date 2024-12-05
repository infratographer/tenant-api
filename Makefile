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

GCI_REPO = github.com/daixiang0/gci
GCI_VERSION = v0.10.1

GOLANGCI_LINT_REPO = github.com/golangci/golangci-lint
GOLANGCI_LINT_VERSION = v1.62.2

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
lint: golint gci-diff  ## Runs all lint checks.

golint: | vendor $(TOOLS_DIR)/golangci-lint  ## Runs Go lint checks.
	@echo Linting Go files...
	@$(TOOLS_DIR)/golangci-lint run --timeout=5m

vendor:  ## Downloads and tidies go modules.
	@go mod download
	@go mod tidy

.PHONY: gci-diff gci-write gci
gci-diff: $(GO_FILES) | $(TOOLS_DIR)/gci  ## Outputs improper go import ordering.
	@results=`$(TOOLS_DIR)/gci diff -s standard -s default -s 'prefix(github.com/infratographer)' $^` \
		&& echo "$$results" \
		&& [ -n "$$results" ] \
			&& [ "$(IGNORE_DIFF_ERROR)" != "true" ] \
			&& echo "Run make gci" \
			&& exit 1 || true

gci-write: $(GO_FILES) | $(TOOLS_DIR)/gci  ## Checks and updates all go files for proper import ordering.
	@$(TOOLS_DIR)/gci write -s standard -s default -s 'prefix(github.com/infratographer)' $^

gci: IGNORE_DIFF_ERROR=true
gci: | gci-diff gci-write  ## Outputs and corrects all improper go import ordering.

# Tools setup
$(TOOLS_DIR):
	mkdir -p $(TOOLS_DIR)

$(TOOLS_DIR)/cockroach: | $(TOOLS_DIR)
	@echo "Downloading cockroach: $(COCKROACH_RELEASE_URL)"
	@curl --silent --fail "$(COCKROACH_RELEASE_URL)" \
		| tar -xz --strip-components 1 -C $(TOOLS_DIR) $(COCKROACH_VERSION_FILE)/cockroach

	$@ version

$(TOOLS_DIR)/gci: | $(TOOLS_DIR)
	@echo "Installing $(GCI_REPO)@$(GCI_VERSION)"
	@GOBIN=$(ROOT_DIR)/$(TOOLS_DIR) go install $(GCI_REPO)@$(GCI_VERSION)
	$@ --version

$(TOOLS_DIR)/golangci-lint: | $(TOOLS_DIR)
	@echo "Installing $(GOLANGCI_LINT_REPO)/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)"
	@GOBIN=$(ROOT_DIR)/$(TOOLS_DIR) go install $(GOLANGCI_LINT_REPO)/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	$@ version
	$@ linters
