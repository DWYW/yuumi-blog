RELATIVE_PATH=api
PROTO_FILES=$(shell cd $(RELATIVE_PATH) && find . -name *.proto)
GW_FILES=$(shell cd $(RELATIVE_PATH) && find . -name *.pb.gw.go)
PB_FILES=$(shell cd $(RELATIVE_PATH) && find . -name *.pb.go)

.PHONY: init
init:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go

# generate api proto
.PHONY: api
api:
	cd $(RELATIVE_PATH) && rm -rf $(PB_FILES)
	cd $(RELATIVE_PATH) && rm -rf $(GW_FILES)
	cd $(RELATIVE_PATH) && protoc \
		--proto_path=. \
		--proto_path=../third_party \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--grpc-gateway_out . \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt allow_delete_body=true \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt generate_unbound_methods=true \
		$(PROTO_FILES)
