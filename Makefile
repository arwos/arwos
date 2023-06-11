
.PHONY: install
install:
	go install github.com/dewep-online/devtool@latest

.PHONY: setup
setup:
	devtool setup-lib

.PHONY: lint
lint:
	devtool lint

.PHONY: build
build:
	devtool build --arch=amd64

.PHONY: tests
tests:
	devtool test

.PHONY: pre-commite
pre-commite: setup lint build tests

.PHONY: ci
ci: install setup lint build tests

run_server:
	go run -race cmd/arvos-server/main.go --config config/server.dev.yaml

run_agent:
	go run -race cmd/arwos-agent/main.go --config config/agent.dev.yaml
