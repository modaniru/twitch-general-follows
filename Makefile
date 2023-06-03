.PHONY: build
build: install
	go build -o twitch-general-follows-api src/main.go 

.PHONY: run
run: fmt
	go run src/main.go

.PHONY: fmt
fmt: install
	go fmt ./...

.PHONY: install
install:
	go mod download

.DEFAULT_GOAL := run