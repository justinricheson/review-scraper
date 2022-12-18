ROOT_DIR := $(abspath ./)
BUILD_DIR := $(ROOT_DIR)/build
COVERAGE_DIR := $(BUILD_DIR)/coverage

.PHONY: *

build:
	@echo "--> building project"
	go build -v -o $(BUILD_DIR)/review-scraper ./...

run:
	@echo "--> running project"
	go run $(ROOT_DIR)/cmd/. -company-id=$(COMPANY_ID)

clean:
	@echo "--> cleaning project"
	go clean -v -i -r
	rm -rf build/

coverage:
	@echo "--> generating code coverage report"
	mkdir -p $(COVERAGE_DIR)
	go test -coverprofile $(COVERAGE_DIR)/coverage.out ./...
	go tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html
	open $(COVERAGE_DIR)/coverage.html

format:
	@echo "--> formatting code"
	golangci-lint run -v --disable-all --enable=goimports --fix

test:
	@echo "--> running unit tests"
	go test ./...

tidy:
	@echo "--> tidying project"
	GOFLAGS=-mod=vendor go get -u ./...
	GOFLAGS=-mod=vendor go mod vendor
	GOFLAGS=-mod=vendor go mod tidy