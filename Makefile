GOCMD=go
BINARY_NAME=makeqr

.PHONY: build test lint clean

build:
	$(GOCMD) build -o $(BINARY_NAME) ./cmd/makeqr

test:
	$(GOCMD) test ./...

lint:
	-golangci-lint run

clean:
	rm -f $(BINARY_NAME)
