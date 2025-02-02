.PHONY: gen
gen:
	@protoc \
		--proto_path=protobuf  \
		--go_out=services/common/genproto/orders \
		--go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/orders \
		--go-grpc_opt=paths=source_relative \
		"protobuf/orders.proto"

.PHONY: orders
orders: gen
	@go run services/orders/*.go

.PHONY: kitchen
kitchen: gen
	@go run services/kitchen/*.go
