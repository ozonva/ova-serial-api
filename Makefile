GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.16","$(shell printf "$(GO_VERSION_SHORT)\n1.16" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.16. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org|direct

PGV_VERSION:="v0.6.1"
GOOGLEAPIS_VERSION="master"
BUF_VERSION:="v0.51.0"

all: generate build

.PHONY: vendor-proto
vendor-proto:
	$(eval THIRD_PARTY:=$(CURDIR)/third_party)
	@[ -f $(THIRD_PARTY)/validate/validate.proto ] || (mkdir -p $(THIRD_PARTY)/validate/ && curl -sSL0 https://raw.githubusercontent.com/envoyproxy/protoc-gen-validate/$(PGV_VERSION)/validate/validate.proto -o $(THIRD_PARTY)/validate/validate.proto)
	@[ -f $(THIRD_PARTY)/google/api/http.proto ] || (mkdir -p $(THIRD_PARTY)/google/api/ && curl -sSL0 https://raw.githubusercontent.com/googleapis/googleapis/$(GOOGLEAPIS_VERSION)/google/api/http.proto -o $(THIRD_PARTY)/google/api/http.proto)
	@[ -f $(THIRD_PARTY)/google/api/annotations.proto ] || (mkdir -p $(THIRD_PARTY)/google/api/ && curl -sSL0 https://raw.githubusercontent.com/googleapis/googleapis/$(GOOGLEAPIS_VERSION)/google/api/annotations.proto -o $(THIRD_PARTY)/google/api/annotations.proto)

GOBIN?=$(GOPATH)/bin
buf.work.yaml:
	@echo "version: v1\ndirectories:\n  - protos\n  - third_party\n" > $(CURDIR)/buf.work.yaml
buf.gen.yaml:
	@echo "version: v1\nplugins:\n  - name: go\n    out: .\n    opt: module=github.com/ozonva/ova-serial-api\n  - name: go\n    out: .\n    opt: module=github.com/ozonva/ova-serial-api\n  - name: go-grpc\n    out: .\n    opt: module=github.com/ozonva/ova-serial-api\n  - name: grpc-gateway\n    out: .\n    opt: logtostderr=true,module=github.com/ozonva/ova-serial-api\n  - name: validate\n    out: .\n    opt: lang=go,module=github.com/ozonva/ova-serial-api\n  - name: openapiv2\n    out: swagger\n    opt: allow_merge=true,merge_file_name=api" > $(CURDIR)/buf.gen.yaml

.PHONY: generate
generate: vendor-proto buf.work.yaml buf.gen.yaml
	@command -v buf 2>&1 > /dev/null || (mkdir -p $(GOBIN) && curl -sSL0 https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$(shell uname -s)-$(shell uname -m) -o $(GOBIN)/buf && chmod +x $(GOBIN)/buf)
	$(eval PROTOS:=$(CURDIR)/api)
	@[ -f $(PROTOS)/buf.yaml ] || PATH=$(GOBIN):$(PATH) buf mod init --doc -o $(PROTOS)
	PATH=$(GOBIN):$(PATH) buf generate $(PROTOS)

.PHONY: deps
deps:
	@[ -f go.mod ] || go mod init github.com/ozonva/ova-serial-api
	find . -name go.mod -exec bash -c 'pushd "$${1%go.mod}" && go mod tidy && popd' _ {} \;

.PHONY: build
build: deps
	go build -o $(CURDIR)/bin/project $(CURDIR)/cmd/ova-serial-api/main.go

.PHONY: run
run:
	go run ./cmd/ova-serial-api/main.go

go-deps:
	ls go.mod
		GOBIN=$(LOCAL_BIN) go get -u github.com/pkg/errors
		GOBIN=$(LOCAL_BIN) go get -u github.com/onsi/ginkgo
		GOBIN=$(LOCAL_BIN) go get -u github.com/onsi/gomega
		GOBIN=$(LOCAL_BIN) go get -u github.com/golang/mock
		GOBIN=$(LOCAL_BIN) go get -u github.com/rs/zerolog/log
		GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto
		GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/protoc-gen-go
		GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc
		GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		GOBIN=$(LOCAL_BIN) go get -u google.golang.org/protobuf/reflect/protoreflect
		GOBIN=$(LOCAL_BIN) go get -u google.golang.org/protobuf/runtime/protoimpl
		GOBIN=$(LOCAL_BIN) go get -u github.com/jmoiron/sqlx
		GOBIN=$(LOCAL_BIN) go get -u github.com/lib/pq
		GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto
		GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/protoc-gen-go
		GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc
		GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
		GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate
		GOBIN=$(LOCAL_BIN) go install github.com/opentracing/opentracing-go/log
		GOBIN=$(LOCAL_BIN) go install github.com/segmentio/kafka-go
		GOBIN=$(LOCAL_BIN) go install github.com/prometheus/client_golang/prometheus
		GOBIN=$(LOCAL_BIN) go install github.com/prometheus/client_golang/prometheus/promauto
