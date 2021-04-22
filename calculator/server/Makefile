all: setup build

setup:
	@echo "inside setup"
	@go mod download

build:
	@echo "inside build"
	@cd /go/src/server/
	@go build calServer/server.go