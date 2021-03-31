GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMODTIDY=$(GOCMD) mod tidy
GOLINT=golint ./... | grep -v "have comment" | grep -v "comment on exported"
CURRENT := $(shell pwd)
BUILDDIR=./build
BINDIR=$(BUILDDIR)/bin
PKGDIR=$(BUILDDIR)/pkg
DISTDIR=$(BUILDDIR)/dist
NAME=api
GOPHER = 'ʕ◔ϖ◔ʔ'
DB_HOST_DEV=localhost
DB_PORT_DEV=35432
DB_USER_DEV=postgres
DB_PASS_DEV=postgres
DB_NAME_DEV=test
DB_URL_DEV=postgres://$(DB_USER_DEV):$(DB_PASS_DEV)@$(DB_HOST_DEV):$(DB_PORT_DEV)/$(DB_NAME_DEV)?sslmode=disable
MIGRATION_DIR=migrations

export GO111MODULES=on

gopher:
	@echo $(GOPHER)

## Set up Git Hooks
.PHONY: setup-githooks
setup-githooks:
	@./setup-hooks.sh

## Install Dependencies
.PHONY: deps
deps:
	@echo "install deps"

# 開発用の依存関係解決
## Setup Dev env
.PHONY: deps
dev-deps: deps
	@echo "install dev-deps"
	export GO111MODULE=off
	go get -u github.com/cosmtrek/air
	go get -u golang.org/x/tools/cmd/goimports
	go get -u golang.org/x/lint/golint
	go get -u github.com/fzipp/gocyclo/cmd/gocyclo
	go get github.com/stretchr/testify
	export GO111MODULE=on

.PHONY: prune
prune:
	@echo "check unused packages"
	@go mod tidy -v

.PHONY: migration-up
migration-up:
	@echo "Migration Up"
	migrate -path $(MIGRATION_DIR) -database $(DB_URL_DEV) up

.PHONY: migration-down
migration-down:
	@echo "Migration Down"
	migrate -path $(MIGRATION_DIR) -database $(DB_URL_DEV) down

.PHONY: migration-create
migration-create:
	@echo "Create New Migration File"
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq ${SEQ}

.PHONY: run
run:
	air -c .air.toml

.PHONY: build
## build: Compile the packages.
build: deps
	@go build -o $(NAME)

.PHONY: run-prod
## run-prod: Build and Run in production mode.
run-prod: build
	@./$(NAME) -e production

.PHONY: format
format:
	find . -print | grep --regex '.*\.go' | xargs goimports -w -l

.PHONY: lint
## lint: Run lint by go vet and golint
lint:
	go vet ./...
	gocyclo -over 15 .
	make golint

.PHONY: golint
golint:
	$(eval LINT_RESULT := $(shell $(GOLINT) | grep -Ev "^$$" | wc -l ))
	@if [ "$(LINT_RESULT)" -gt "0" ]; then \
		$(GOLINT); exit 1; else \
		echo "golint success"; \
	fi

.PHONY: test
## test: Run tests with verbose mode
test: dev-deps
	go test -v ./...
