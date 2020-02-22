TAG := $(shell git describe --always --long)
REV := $(shell git rev-parse HEAD)
DATE := $(shell date +"%F %X")


default: build

test: export GO111MODULE = on
test:
	@cd .. && go install gotest.tools/gotestsum
	@gotestsum --junitfile test-results.xml -- -covermode=count -coverprofile=count.out ./...
	@go tool cover -func=count.out
	@go tool cover -html=count.out -o coverage.html
	@go clean -i gotest.tools/gotestsum

build:
	go build

release:
	go build -ldflags="-X 'github.com/tlkamp/mockbob/cmd.versionInfo=$(TAG)' -X 'github.com/tlkamp/mockbob/cmd.gitRev=$(REV)' -X 'github.com/tlkamp/mockbob/cmd.buildDate=$(DATE)'"

safe-clean:
	@git clean -Xi

clean:
	@git clean -Xf
