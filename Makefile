GOPATH= $(shell go env GOPATH)

.PHONY: default
default: lint fmt test

.PHONY: test
test:
	go test -race ./...

.PHONY: lint
lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.37.0
	$(GOPATH)/bin/golangci-lint run -e gosec ./...
	go mod tidy

.PHONY: fmt
fmt:
	go fmt ./...

