GIT_COMMIT = $(shell git rev-parse --short HEAD)
VERSION_FLAG = -X pixie_adapter/internal/config.Version=$(GIT_COMMIT)
GC = go build -ldflags="-s -w $(VERSION_FLAG)" -trimpath

.PHONY: configure
configure:
	go mod tidy

.PHONY: build
build:
	@echo "building..."
	@$(GC) .

.PHONY: install
install: build
	go install .

.PHONY: lint
lint:
	golangci-lint run ./...
	revive -config ./revive.toml  ./...

.PHONY: test
test:
	go test ./...
