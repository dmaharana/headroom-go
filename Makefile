# Headroom Go Makefile

# Binary names
BINARY_NAME=headroom-mcp
BUILD_DIR=bin

# Build flags
# -s: Omit the symbol table and debug information.
# -w: Omit the DWARF symbol table.
LDFLAGS=-ldflags="-s -w"

.PHONY: all build clean test run-mcp help

all: build

## build: Build all binaries
build:
	@echo "Building optimized binaries..."
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/mcp/main.go
	@echo "Built $(BUILD_DIR)/$(BINARY_NAME)"

## clean: Remove build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

## test: Run all tests
test:
	@echo "Running tests..."
	go test ./... -v

## run-mcp: Run the MCP server directly
run-mcp: build
	./$(BUILD_DIR)/$(BINARY_NAME)

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'
