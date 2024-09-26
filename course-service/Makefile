.PHONY: example dev protoc

# Example task
run_example:
	export APP_ENV=example && go run cmd/main.go

# Development task
run_dev:
	export APP_ENV=dev && air

# Protobuf generation task
protoc:
	protoc proto/*.proto \
		--go_out=pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=require_unimplemented_servers=false:pb \
		--go-grpc_opt=paths=source_relative \
		-I=proto \
		--experimental_allow_proto3_optional
