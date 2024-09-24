PROTOC_GEN_GO := protoc \
	--proto_path=./proto \
	--proto_path=./googleapis \
	--go_out=./proto \
	--go_opt=paths=source_relative \
	--go-grpc_out=./proto \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=./proto \
	--grpc-gateway_opt=paths=source_relative

PROTOS := proto/assessments/assessments.proto \
          proto/certificate/certificate.proto \
          proto/enrollment/enrollment.proto \
          proto/meta/meta.proto \
          proto/payments/payments.proto

all: generate

generate:
	$(PROTOC_GEN_GO) $(PROTOS)

clean:
	rm -rf proto/**/*.pb.go

.PHONY: all generate clean
