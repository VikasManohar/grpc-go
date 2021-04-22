all: setup build

setup:
	@echo "inside setup"
	@mkdir -p /go/src/server
	@cd /go/src/server
	@cp -r . /go/src/server
	@go mod download

build:
	@echo "inside build"
	@cd /go/src/server/calServer
	@go build server.go