install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/blockchain.proto

main:
	go run main.go

docker:
	docker build -t 'in-memory-blockchain' .
	docker run -p 9000:9000 in-memory-blockchain

.PHONY: test
test:
	go clean -testcache
	go test ./...