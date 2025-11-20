# makeqr

Small CLI utility to generate QR codes as PNG files.

## Project layout

- `cmd/makeqr` — CLI entrypoint
- `internal/cli` — CLI parsing and config
- `internal/qr` — QR generation logic
- `internal/logger` — logging helper
- `testdata` — golden files for tests

## Install / Build

Requires Go (1.20+ recommended).

Clone the repo and build the CLI:

```bash
git clone <repo-url>
cd makeqr
make build
```

The built binary is `./makeqr`.

## Usage

Generate a QR for a URL and save as `qr.png`:

```bash
./makeqr --url https://example.com
```

Specify output file and size:

```bash
./makeqr --url https://example.com -o myqrcode.png --size 512
```

## Library

Use the generator from other Go code:

```go
import "makeqr/internal/qr"

err := qr.GenerateFile("https://example.com", qrcode.Medium, 256, "out.png")
```

## Tests

Run all tests:

```bash
make test
```

There are golden-file checks under `testdata/` used by unit tests.

## CI

CI is provided in `.github/workflows/ci.yml` and runs `golangci-lint`, `go test` and `go vet`.

## License

This project is released under the MIT License — see `LICENSE`.
