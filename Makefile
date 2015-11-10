NAME = platform-users

GO ?= go
PROTOC ?= protoc

PROTOCFLAGS = --go_out=plugins=grpc:.
PROTOSPATH = protos
PLATFORMPROTOPATH = $(GOPATH)/src
PROTOS = protos/*.proto
#GOPATH := $(CURDIR):$(GOPATH)

all: build

build:
	@echo "Building $(NAME) protocol buffers..."
	$(PROTOC) -I $(PROTOSPATH) -I $(PLATFORMPROTOPATH) $(PROTOS) $(PROTOCFLAGS)

test:
	$(GO) test
