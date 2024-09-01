.PHONY: test benchmark

test:
	go test -race ./...

bench:
	go test -bench=. ./...
