# Variables
GOOS=js
GOARCH=wasm
WASM_OUT_DIR=./static
WASM_BINARY=$(WASM_OUT_DIR)/main.wasm
WASM_EXEC_JS=$(WASM_OUT_DIR)/wasm_exec.js
GO_SOURCE=./wasm/client.go
GOROOT=$(shell go env GOROOT)

# Default target to build the Wasm and copy required files
all: build-wasm copy-wasm-exec

# Build the WebAssembly binary from the Go source code
build-wasm:
	@echo "Building WebAssembly binary..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(WASM_BINARY) $(GO_SOURCE)
	@echo "Wasm binary built at $(WASM_BINARY)"

# Copy the wasm_exec.js file needed for running the Wasm in the browser
copy-wasm-exec:
	@echo "Copying wasm_exec.js..."
	cp $(GOROOT)/misc/wasm/wasm_exec.js $(WASM_EXEC_JS)
	@echo "wasm_exec.js copied to $(WASM_EXEC_JS)"

# Clean up the generated files
clean:
	@echo "Cleaning up generated files..."
	rm -f $(WASM_BINARY) $(WASM_EXEC_JS)
	rm -f $(JS_OUT_DIR)/entities.js
	@echo "Cleaned up."

# Phony targets (not real files)
.PHONY: all build-wasm copy-wasm-exec generate-js clean