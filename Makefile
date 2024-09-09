VENV_PATH = ./venv
VENV = . $(VENV_PATH)/bin/activate;
GC = go build -ldflags="-s -w" -trimpath

.PHONY: configure
configure:
	rm -rf "$(VENV_PATH)"
	python3.11 -m venv "$(VENV_PATH)"
	$(VENV) pip install -r requirements.txt

.PHONY: build
build:
	$(GC) -o ./build/pixie-adapter pixie-adapter.go

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
