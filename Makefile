GOCMD=go
BINARY_NAME=makeqr

.PHONY: build test lint clean

build:
	$(GOCMD) build -o $(BINARY_NAME) ./cmd/makeqr

test:
	$(GOCMD) test ./...

lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint not found; install with 'brew install golangci-lint'"; exit 1; }
	golangci-lint run

clean:
	rm -f $(BINARY_NAME)
