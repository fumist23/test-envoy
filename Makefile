.PHONY: proto-generate

proto-generate:
	protoc -I/Users/satoubunnari/package-references/googleapis -I. --include_imports --include_source_info \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--descriptor_set_out=./test_service.pb \
	api/v1/*.proto


