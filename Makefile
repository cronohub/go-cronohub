NAME=cronohub

.PHONY: build
build:
	cd src && go build -ldflags="-s -w" -i -o ./../cmd/${NAME}

.PHONY: test
test:
	go test ./...
