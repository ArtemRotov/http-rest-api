.PHONY: run
run:
	go run ./cmd/apiserver/main.go
 
.PHONY: build
build:
	go build -v -o http-rest-api-server ./cmd/apiserver/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build