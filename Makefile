GIT_COMMIT = $(shell git rev-parse --short HEAD)
VERSION_FLAG = -X pixie_adapter/internal/config.Version=$(GIT_COMMIT)
GC = go build -ldflags="-s -w $(VERSION_FLAG)" -trimpath

.PHONY: configure
configure:
	go mod tidy

.PHONY: build
build:
	@echo "building..."
	@$(GC) -o ./build/pixie-adapter pixie-adapter.go

.PHONY: install
install:
	sudo rm -f /usr/local/bin/pixie-adapter
	sudo cp ./build/pixie-adapter /usr/local/bin/pixie-adapter

.PHONY: lint
lint:
	golangci-lint run ./...
	revive -config ./revive.toml  ./...

.PHONY: test
test:
	go test ./...
