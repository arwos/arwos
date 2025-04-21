SHELL:=/bin/bash

run_setup:
	DEBUG=dev go run ./main.go setup --sock=/tmp/app.sock

run_app:
	DEBUG=dev go run ./main.go --sock=/tmp/app.sock

run_cli:
	go clean -testcache && go test -v -race ./...