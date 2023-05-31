.PHONY: proto

proto:
	protoc --proto_path=proto --go_out=pb --go-grpc_out=pb \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*.proto