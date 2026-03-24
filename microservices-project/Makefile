# ============================================================
#  Protobuf Code Generation
# ============================================================

# Install the two protoc plugins needed for Go code generation.
# Run this once on a new machine.
.PHONY: proto-tools
proto-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate Go code from account.proto → account/pb/
.PHONY: proto-account
proto-account:
	cd account && protoc \
		--go_out=./pb --go_opt=paths=source_relative \
		--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
		account.proto

# Generate Go code from catalog.proto → catalog/pb/
.PHONY: proto-catalog
proto-catalog:
	cd catalog && protoc \
		--go_out=./pb --go_opt=paths=source_relative \
		--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
		catalog.proto

# Generate Go code from order.proto → order/pb/
.PHONY: proto-order
proto-order:
	cd order && protoc \
		--go_out=./pb --go_opt=paths=source_relative \
		--go-grpc_out=./pb --go-grpc_opt=paths=source_relative \
		order.proto

# Generate all protobuf files
.PHONY: proto
proto: proto-account proto-catalog proto-order

# Remove all generated protobuf files
.PHONY: proto-clean
proto-clean:
	rm -f account/pb/*.pb.go
	rm -f catalog/pb/*.pb.go
	rm -f order/pb/*.pb.go

# Reinstall tools + regenerate everything from scratch
.PHONY: proto-rebuild
proto-rebuild: proto-tools proto-clean proto
