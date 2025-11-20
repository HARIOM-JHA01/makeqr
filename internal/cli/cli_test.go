package cli

import (
    "flag"
    "os"
    "testing"

    qrcode "github.com/skip2/go-qrcode"
)

func TestParseDefaults(t *testing.T) {
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()

    os.Args = []string{"makeqr"}
    flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
    cfg := Parse()
    if cfg.Out != "qr.png" {
        t.Fatalf("expected default out qr.png, got %s", cfg.Out)
    }
    if cfg.Size != 256 {
        t.Fatalf("expected default size 256, got %d", cfg.Size)
    }
    if cfg.Level != qrcode.Medium {
        t.Fatalf("expected default level Medium, got %v", cfg.Level)
    }
}
