all: setup build

setup:
	@echo "inside setup"
	@go mod download

build:
	@echo "inside build"
	@cd /go/src/client/
	@go build calClient/client.go