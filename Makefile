BINARY := mockbob

.PHONY: fmt build test
fmt:
	@go fmt ./...

build:
	@go build -o $(BINARY) .

test:
	@go test -v ./... -cover -coverprofile=.coverage.txt
	@go tool cover -func=.coverage.txt

