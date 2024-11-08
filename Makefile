GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date +%FT%T%Z)
PROTO_FILES=$(shell find . -name *.proto)

.PHONY: api
api:
	protoc --proto_path=. \
           --go_out=paths=source_relative:. \
           --go-grpc_out=paths=source_relative:. \
           --validate-go_out=paths=source_relative:. \
           --openapi_out=. $(PROTO_FILES)