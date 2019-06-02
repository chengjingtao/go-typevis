OS = Linux

VERSION = 0.0.1
COMMIT=$(shell git rev-parse --short HEAD)
CURDIR = $(shell pwd)
SOURCEDIR = $(CURDIR)
COVER = $($3)
export GO111MODULE=on

ECHO = echo
RM = rm -rf
MKDIR = mkdir

# If the first argument is "cover"...
ifeq (cover,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

BUILD_PATH = $(shell if [ "$(ALAUDACI_DEST_DIR)" != "" ]; then echo "$(ALAUDACI_DEST_DIR)" ; else echo "$(PWD)"; fi)
PACKAGES = $(shell go list ./... | grep -v './vendor/\|./tests\|./mock')
COMMIT=$(shell git rev-parse --short HEAD)
DATE=$(shell date +%m%d%H%M)
BUILDDATE=$(shell date +%Y%m%d%H%M)

.PHONY: all

default: test lint vet


test:
	go test -cover=true $(PACKAGES)

race:
	go test -cover=true -race $(PACKAGES)

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt $(PACKAGES)

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint .

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	go vet $(PACKAGES)

cover: collect-cover-data test-cover-html open-cover-html


collect-cover-data:
	echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES),\
		go test -v -coverprofile=coverage.out -covermode=count $(pkg);\
		if [ -f coverage.out ]; then\
			tail -n +2 coverage.out >> coverage-all.out;\
		fi;)

test-cover-html:
	go tool cover -html=coverage-all.out -o coverage.html

test-cover-func:
	go tool cover -func=coverage-all.out

open-cover-html:
	open coverage.html

build-pkg:
	@$(ECHO) "Will build on "$(BUILD_PATH)
	go build -ldflags "-w -s" -v ./

build: build-darwin build-linux build-win

build-darwin:
	go build -o ./bin/darwin/go-typevis \
		-ldflags '-w -s -X main.version=${COMMIT} -X main.buildDate=${BUILDDATE}' \
		./
	cp ./bin/darwin/go-typevis ${GOPATH}/bin/

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/linux-amd64/go-typevis \
		-ldflags '-w -s -X main.version=${COMMIT} -X main.buildDate=${BUILDDATE}' \
		./

build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/windows-amd64/go-typevis.exe \
		-ldflags '-w -s -X main.version=${COMMIT} -X main.buildDate=${BUILDDATE}' \
		./

clean:
		rm -r ./bin/*

demo: build
	go-typevis types -p github.com/chengjingtao/go-typevis > demo.dot