# Credit: https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects

BINARY_NAME=federation-cli
API_VERSION=v1alpha
PROTO_DIR=api/protobuf/$(API_VERSION)
PROTOBUF_MESSAGES_INPUT=$(PROTO_DIR)/FederationService.proto
PROTOBUF_GRPC_INPUT=$(PROTO_DIR)/FederationService.proto
GO_OUT_DIR=pkg/api
PB_GO_FILE=$(GO_OUT_DIR)/$(API_VERSION)/FederationService.pb.go
GRPC_PB_GO_FILE=$(GO_OUT_DIR)/$(API_VERSION)/FederationService_grpc.pb.go

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} #| column -t -s ':' |  sed -e 's/^/ /'

.PHONY: no-dirty
no-dirty:
	git diff --exit-code

.PHONY: clean
clean:
	rm -f $(BINARY_NAME) $(PB_GO_FILE) $(GRPC_PB_GO_FILE)

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## build: build the application binary
.PHONY: build
build: protobuf
	go build -o $(BINARY_NAME) main.go

## protobuf: generate GRPC and protobuf files
.PHONY: protobuf
protobuf: $(GRPC_PB_GO_FILE) $(PB_GO_FILE)

$(PB_GO_FILE): $(PROTOBUF_MESSAGES_INPUT)
	mkdir -p $(GO_OUT_DIR)/$(API_VERSION)
	protoc --proto_path=$(PROTO_DIR) --proto_path=third_party --go_out=$(GO_OUT_DIR) $(PROTOBUF_MESSAGES_INPUT)

$(GRPC_PB_GO_FILE): $(PROTOBUF_MESSAGES_INPUT) $(PROTOBUF_GRPC_INPUT)
	mkdir -p $(GO_OUT_DIR)/$(API_VERSION)
	protoc --proto_path=$(PROTO_DIR) --proto_path=third_party --go-grpc_out=$(GO_OUT_DIR) $(PROTOBUF_GRPC_INPUT)
