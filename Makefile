APP_NAME := ftf
BUILD_DIR := bin
GO_FILES := $(shell find . -name '*.go')

.PHONY: all build clean run test lint vet fmt

all: vet fmt lint clean test build install

build: $(GO_FILES)
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) .

clean:
	rm -rf $(BUILD_DIR)
	rm -f $(APP_NAME)

run: build
	$(BUILD_DIR)/$(APP_NAME)

test:
	go test ./... -v

lint:
	staticcheck ./...

vet:
	go vet ./...

fmt:
	go fmt ./...

install:
	go install .
