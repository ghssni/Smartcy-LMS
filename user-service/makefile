# Command for generating Go code from .proto files
PROTOC := protoc

PROTOS := proto/user.proto proto/userLog.proto

# Output directories for generated Go files
PB_DIR := ./pb

all: generate

generate:
	$(PROTOC) \
		-I . \
		-I ./googleapis \
		--go_out=$(PB_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PB_DIR) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(PB_DIR) --grpc-gateway_opt=paths=source_relative \
		$(PROTOS)

# Clean generated files
clean:
	rm -rf $(PB_DIR)/*.pb.go

.PHONY: all generate clean
