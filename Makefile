.PHONY: test benchmark

test:
	go test -race ./...

benchmark:
	go test -bench=. ./...
