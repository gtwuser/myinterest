# Copy this file to your project and customize it.

#.DEFAULT_GOAL := help
#
#ACTUAL_PWD := $(PWD)
#BIN_DIR := $(ACTUAL_PWD)/bin
#BUILD_DIR := $(ACTUAL_PWD)/build
#CACHE_DIR := $(ACTUAL_PWD)/.cache
#
##### FOR GRPC CODE GENERATION ####
#API_PROTO_FILES_IN_SRC := $(wildcard qrcode/cmd/grpcapps/pob/*.proto)
## API_PROTO_FILES_W_MODULE_VERSION := $(foreach f,$(API_PROTO_FILES_IN_SRC),$(basename $(shell echo $(f) | cut -d'/' -f3-)))
#API_PROTO_FILES_BASENAMES := $(foreach f,$(API_PROTO_FILES_IN_SRC),$(basename $(shell echo $(f) | cut -d'/' -f5-)))
## SWAGGER_JSONS := $(foreach f,$(API_PROTO_FILES_BASENAMES),$(shell echo $(f) | awk '{ print "api/pb/" $$1 ".swagger.json" }'))
## WEB_PROTOS    := $(foreach f,$(API_PROTO_FILES_W_MODULE_VERSION),$(shell echo $(f) | awk '{ print "api/web/" $$1 ".proto" }'))
## WEB_TS        := $(foreach f,$(API_PROTO_FILES_W_MODULE_VERSION),$(shell echo $(f) | awk '{ print "api/web/" $$1 "_pb.d.ts" }'))
#GRPC_PB       := $(foreach f,$(API_PROTO_FILES_BASENAMES),$(shell echo $(f) | awk '{ print "qrcode/cmd/grpcapps/pob/" $$1 ".pb.go" }'))
#GRPC_PB_GW    := $(foreach f,$(API_PROTO_FILES_BASENAMES),$(shell echo $(f) | awk '{ print "qrcode/cmd/grpcapps/pob/" $$1 ".pb.gw.go" }'))
#
#$(shell mkdir -p $(BUILD_DIR))
#$(shell mkdir -p $(CACHE_DIR))
#$(shell mkdir -p $(BIN_DIR))
#
#GOFLAGS     := GOFLAGS="-mod=vendor"  # For vendored deps.
#CGO_ENABLED := 0
#
#
#UNAME_S?= $(shell uname -s)
#ifeq ("${UNAME_S}", "Darwin")
#    GOOS = darwin
#else
#    GOOS = linux
#endif
#GOARCH   := amd64
#
#GOFILES := $(shell find . -type d \( -name .git -o -name vendor -o -name .submodules -o -name .cache \) -prune -o -type f -name "*.go" -print)
#
#protoc := ./bin/protoc
#protoc += -I qrcode/cmd/grpcapps/pob
#protoc += --proto_path=./qrcode/cmd/grpcapps/pob --go_out=./qrcode/cmd/grpcapps/pob --go_opt=paths=source_relative --go-grpc_out=./qrcode/cmd/grpcapps/pob --go-grpc_opt=paths=source_relative --grpc-gateway_out ./qrcode/cmd/grpcapps/pob  --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative
#
#
#.PHONY: gen-grpc
#gen-grpc:  # Generate gRPC artifacts.
#gen-grpc: gen-grpc-pb
#
#.PHONY: gen-grpc-pb
## Generate gRPC Protocol Buffers artifacts.
#gen-grpc-pb: $(GRPC_PB)
#qrcode/cmd/grpcapps/pob/%.pb.go: qrcode/cmd/grpcapps/pob/%.proto
#	@echo "==> Generating .proto ($<) to GRPC backend ($@)"
#	$(protoc) --go_out="$$(dirname $<)" --go_opt=paths=source_relative "$<"
#

#
# .PHONY: gen-grpc-pb-gw
# gen-grpc-pb-gw: $(GRPC_PB_GW)
# qrcode/cmd/grpcapps/pob/%.pb.gw.go: qrcode/cmd/grpcapps/pob/%.proto
# 	@echo "==> Generating .proto ($<) to GRPC Gateway ($@) JSON/REST proxy"
# 	$(protoc) -I "$$(dirname $<)" --grpc-gateway_out=logtostderr=true:api/pb \
# 	  --plugin=protoc-gen-grpc-gateway=./bin/protoc-gen-grpc-gateway "$<"
#
# .PHONY: gen-grpc-swagger
# gen-grpc-swagger: $(SWAGGER_JSONS)
# api/pb/%.swagger.json: api/src/*/*/%.proto
# 	@echo "==> Generating .proto ($<) to OpenAPI Swagger ($@)"
# 	$(protoc) -I "$$(dirname $<)" --swagger_out=logtostderr=true:api/pb "$<"

LIST = 1 2 3 4
DIRS = builds builds/bins

.PHONY: all
all:
all: clean touch loop loop1 loop2 loop3

.PHONY: clean
clean:
	@echo "Removing"
	@rm -rf *.00

.PHONY: touch
touch:
	@echo "creating directory"
    $(shell mkdir -p $(DIRS))

.PHONY: loop
loop:
	@echo "Inside loop"
	@for num in $(LIST); do	\
		echo $$num; \
	done

.PHONY: loop1
loop1:
	@echo "Inside loop1"

.PHONY: loop2
loop2:
	@echo "Inside loop2"
	@curl http://stash.compciv.org/ssa_baby_names/names.zip  \
       --output babynames.zip
	@unzip babynames.zip -d temp

.PHONY: loop3
loop3:
	@echo "Inside loop3"
