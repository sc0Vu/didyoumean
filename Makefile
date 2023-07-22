GOPATH= $(shell go env GOPATH)

.PHONY: default
default: lint test fuzz

.PHONY: test
test:
	go test -race ./...

.PHONY: fuzz
fuzz:
	go test -run=FuzzFindEditDistance -fuzztime=1s ./...

.PHONY: lint
lint:
	# go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.37.0
	# $(GOPATH)/bin/golangci-lint run -e gosec ./...
	go mod tidy
	go vet ./...
	go fmt ./...
