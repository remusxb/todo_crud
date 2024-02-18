.PHONY: build test clean tools docker-compose docker git-hooks lint

BIN_NAME ?= todo_crud
DOCKERFILE ?= dockerfiles/Dockerfile

COMMIT = $(shell git rev-parse HEAD)
VERSION = $(shell git describe --tags --match=v* 2>/dev/null || git rev-parse HEAD)
BUILD_DATE = $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

build:
	@echo "Building $(BIN_NAME):$(VERSION)"
	@echo "GOPATH=$(GOPATH)"
	@echo "$(BUILD_DATE)"
	go mod tidy
	go build -ldflags "-X github.com/remusxb/todo_crud/version.Version=$(VERSION) -X github.com/remusxb/todo_crud/version.GitCommit=$(COMMIT) -X github.com/remusxb/todo_crud/version.BuildDate=$(BUILD_DATE)" -o $(BIN_NAME) ./cmd/todo_crud

test:
	@echo "Running unit, integration, and e2e tests..."
	go test -tags=unit,integration,e2e -v ./...

clean:
	rm -f todo_crud

tools:
	@echo "Getting tools"
	go generate tools/*.go

docker-compose: tools
	docker-compose up

docker: tools
	DOCKER_BUILDKIT=1 docker build --file $(DOCKERFILE) --tag "$(BIN_NAME):latest" .

git-hooks:
	python3 -m venv venv
	venv/bin/python -m pip install pre-commit
	venv/bin/python -m pip install gitlint
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	venv/bin/pre-commit install
	venv/bin/pre-commit install --hook-type commit-msg

lint:
	venv/bin/pre-commit run --all-files
