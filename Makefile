BINARY=ttt
VERSION=1.0.0
BUILD=`git rev-parse HEAD`
PLATFORMS=windows linux darwin  
ARCH=amd64
K := $(foreach exec,$(EXECUTABLES),\
        $(if $(shell which $(exec)),some string,$(error "No $(exec) in PATH")))

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

FLAGS= -a -tags netgo -ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD} -w -extldflags -static" 

all: gomod clean build_all

build:
	go mod download
	(cd cmd/ttt/;go build ${FLAGS} )

build_all:
	go mod download
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCH), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH);cd cmd/ttt/;go build ${FLAGS} -o $(BINARY)-$(GOOS)-$(GOARCH)))) 
	@echo "build_all done"

clean:
	@find cmd/ttt/ -name '${BINARY}[-?][a-zA-Z0-9]*[-?][a-zA-Z0-9]*' -delete
	@echo "cleanup done"

tests:
	go test -race -v ./... 

gomod:
	@go mod tidy
	@go mod download
	@echo "go mod tidy & download done"

.PHONY: gomod clean build_all all tests