NAME=cronohub

.PHONY: build
build:
	go build -ldflags="-s -w" -i -o cmd/${NAME}

.PHONY: test
test:
	go test ./...
