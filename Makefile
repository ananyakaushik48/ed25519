 build:
	go build -o bin/blocker

run:
	./bin/blocker

test:
	go test -v ./...

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
    proto/types.proto


.PHONY: proto