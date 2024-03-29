GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
API_PROTO_PATH=$(shell find api -name *.proto | xargs dirname | uniq)
BIN=./bin/
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest

.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=./$(API_PROTO_PATH) \
			--proto_path=./third_party \
			--go_out=paths=source_relative:./$(API_PROTO_PATH) \
			--go-errors_out=paths=source_relative:./$(API_PROTO_PATH) \
			$(API_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./$(API_PROTO_PATH) \
		--proto_path=./third_party \
		--go_out=paths=source_relative:./$(API_PROTO_PATH) \
		--go-http_out=paths=source_relative:./$(API_PROTO_PATH) \
		--go-grpc_out=paths=source_relative:./$(API_PROTO_PATH) \
		--go-errors_out=paths=source_relative:./$(API_PROTO_PATH) \
		--openapi_out==paths=source_relative:. \
		$(API_PROTO_FILES) 

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: build_linux
# build_linux
build_linux:
ifdef ONLINE
ifeq (${BRANCH},master)
	mkdir -p $(BIN) && env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o $(BIN) ./cmd/...
else
	$(error the current git branch:${BRANCH} not master)
endif
else
	mkdir -p $(BIN) && env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o $(BIN) ./cmd/...
endif

.PHONY: tool
# tool
tool:
	mkdir -p $(BIN) && env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o $(BIN) ./tools/...

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make config;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
