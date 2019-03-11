GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test

all: build

build:
	$(GOBUILD)

deps:
	$(GOGET) ./..

test:
	$(GOTEST) -v --tags mock

test-api:
	$(GOTEST) -v