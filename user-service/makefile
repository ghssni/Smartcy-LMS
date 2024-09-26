# Command for generating Go code from .proto files
PROTOC := protoc

PROTOS := proto/user.proto proto/userLog.proto

all: generate

generate:
	$(PROTOC) \
    --proto_path=./proto \
    --go_out=./pb --go_opt=paths=source_relative \
    --go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
    $(PROTOS)

# Clean generated files
clean:
	rm -rf pb/*.pb.go

.PHONY: all generate clean
