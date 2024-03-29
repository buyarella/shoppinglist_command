# kernel-style V=1 build verbosity
ifeq ("$(origin V)", "command line")
       BUILD_VERBOSE = $(V)
endif

ifeq ($(BUILD_VERBOSE),1)
       Q =
else
       Q = @
endif

VERSION = $(shell git describe --dirty --tags --always)
SERVICE_NAME = shoppinglist_command

export CGO_ENABLED:=0

all: verify build

lint: fmt
		$(Q)echo "linting...."
		$(Q)GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
		$(Q)golangci-lint run -E gofmt -E golint -E goconst -E gocritic -E golint -E gosec -E maligned -E nakedret -E prealloc -E unconvert -E gocyclo -E scopelint -E goimports
		$(Q)echo linting OK

test:
		$(Q)echo "unit testing...."
		$(Q)go test ./pkg/...

verify: lint test

api:
		$(Q)protoc -I protos --go_out=plugins=grpc:. protos/models/v1/models.proto	
		$(Q)protoc -I protos --go_out=plugins=grpc:. protos/shoppinglist_command/v1/shoppinglist_command.proto	

prepare: fmt generate verify

clean:
		$(Q)rm -rf build

fmt:
		$(Q)echo "fixing imports and format...."
		$(Q)goimports -w .

generate:
		$(Q)go get github.com/golang/mock/gomock
		$(Q)go install github.com/golang/mock/mockgen
		$(Q)go generate ./...

build:
		$(Q)$(GOARGS) go build -gcflags "all=-trimpath=${GOPATH}" -asmflags "all=-trimpath=${GOPATH}" -o ./artifacts/$(SERVICE_NAME) ./$(SERVICE_NAME).go
