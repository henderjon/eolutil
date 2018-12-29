BIN=eolutil
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S %z %Z')
DEPLOYMENT_PATH=bin/release/$(BIN)/$(BIN)-$(HEAD)

LDFLAGS="-X main.buildVersion=$(HEAD) -X 'main.buildTimestamp=$(TIMESTAMP)'"

all: clean build

tar: clean build build-tar

.PHONY: build
build: darwin64 linux64

.PHONY: build-tar
build-tar: darwin64-tar linux64-tar

.PHONY: dep
dep:
	go mod vendor

clean:
	rm -f $(BIN) $(BIN)-*

.PHONY: local
local:
	go build -ldflags $(LDFLAGS) -o $(BIN)

darwin64:
	env GOOS=darwin GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN)-darwin64-$(HEAD)

linux64:
	env GOOS=linux GOARCH=amd64 go build -ldflags $(LDFLAGS) -o $(BIN)-linux64-$(HEAD)

darwin64-tar:
	tar czvf $(BIN)-darwin64-$(HEAD).tgz $(BIN)-darwin64-$(HEAD)

linux64-tar:
	tar czvf $(BIN)-linux64-$(HEAD).tgz $(BIN)-linux64-$(HEAD)

.PHONY: test
test:
	go test -coverprofile=coverage.out -covermode=count

.PHONY: test-report
test-report: test
	go tool cover -html=coverage.out
