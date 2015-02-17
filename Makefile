all: dependencies test

test:
	go test ./...

dependencies:
	which go

.PHONY: all dependencies test
