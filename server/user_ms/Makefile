binary_name=main

# Go build parameters
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
BUILD_FLAGS ?= -v

# Executable names
WINDOWS_EXEC := main.exe
LINUX_EXEC := main
MAC_EXEC := main

all: windows linux mac

windows:
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o ./bin/$(WINDOWS_EXEC) ./cmd/api/main.go
	
linux:
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o ./bin/$(LINUX_EXEC) ./cmd/api/main.go
	
mac:
	GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) -o ./bin/$(MAC_EXEC) ./cmd/api/main.go

run-windows:
	./bin/${WINDOWS_EXEC}

run-linux:
	./bin/${LINUX_EXEC}

run-mac:
	./bin/${MAC_EXEC}

clean:
	rm -rf ./bin/*

dev:
	go run ./cmd/api/main.go

testAll:
	go clean -testcache
	go test -v ./test/...

testConfig:
	go clean -testcache
	go test -v ./test/config

testHandlers:
	go clean -testcache
	go test -v ./test/handlers/...