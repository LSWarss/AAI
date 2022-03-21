.DEFAULT_GOAL := test

fmt:
	go fmt ./...
.PHONY:fmt

vet: fmt
	go vet ./...
.PHONY:vet

test: vet
	go test -v ./... -cover
.PHONY:test

benchmark:
	go test ./... -bench=.
.PHONY:benchmark

errors:
	errcheck ./...
.PHONY:errors