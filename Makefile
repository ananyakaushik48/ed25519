 build:
	go build -o bin/blocker

run:
	./bin/blocker

test:
	go test -v ./...