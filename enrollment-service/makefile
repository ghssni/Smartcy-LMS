# Command for generating Go code from .proto files
PROTOC_GEN_GO := protoc \
	--proto_path=./proto \
	--proto_path=./googleapis \
	--go_out=./proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=./proto \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=./proto \
	--grpc-gateway_opt=paths=source_relative

# List of .proto files to generate code for
PROTOS := proto/assessments/assessments.proto \
          proto/certificate/certificate.proto \
          proto/enrollment/enrollment.proto \
          proto/payments/payments.proto

# Default task
all: generate

# Task for generating Go code from .proto files
generate:
	$(PROTOC_GEN_GO) $(PROTOS)

# Task for cleaning up generated files
clean:
	rm -rf proto/**/*.pb.go

# Phony targets (to avoid conflicts with files named 'all', 'generate', or 'clean')
.PHONY: all generate clean
