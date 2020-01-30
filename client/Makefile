# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=client_binary
BINARY_UNIX=$(BINARY_NAME)_unix

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
	docker build -t guanw0826/example-client .

docker-push:
	docker push guanw0826/example-client

docker-run:
	docker run -d -p 8082:8081 guanw0826/example-client