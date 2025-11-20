package main

import (
    "fmt"
    "os"

    "makeqr/internal/cli"
    "makeqr/internal/qr"
)

func main() {
    cfg := cli.Parse()
    if cfg.URL == "" {
        fmt.Println("Usage: makeqr --url https://example.com [-o file]")
        os.Exit(1)
    }
    if err := qr.GenerateFile(cfg.URL, cfg.Level, cfg.Size, cfg.Out); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    fmt.Println("QR saved at:", cfg.Out)
}
