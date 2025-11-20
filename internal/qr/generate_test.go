package qr

import (
    "os"
    "testing"
    "io"

    qrcode "github.com/skip2/go-qrcode"
)

func TestGenerateFile(t *testing.T) {
    tmp, err := os.CreateTemp("", "qr_test_*.png")
    if err != nil {
        t.Fatalf("failed to create temp file: %v", err)
    }
    path := tmp.Name()
    tmp.Close()
    defer os.Remove(path)

    if err := GenerateFile("https://example.com", qrcode.Medium, 256, path); err != nil {
        t.Fatalf("GenerateFile returned error: %v", err)
    }

    fi, err := os.Stat(path)
    if err != nil {
        t.Fatalf("stat file: %v", err)
    }
    if fi.Size() == 0 {
        t.Fatalf("generated file is empty")
    }
}

func TestGeneratedIsPNG(t *testing.T) {
    tmp, err := os.CreateTemp("", "qr_test_*.png")
    if err != nil {
        t.Fatalf("failed to create temp file: %v", err)
    }
    path := tmp.Name()
    tmp.Close()
    defer os.Remove(path)

    if err := GenerateFile("https://example.com/golden", qrcode.Medium, 256, path); err != nil {
        t.Fatalf("GenerateFile returned error: %v", err)
    }

    f, err := os.Open(path)
    if err != nil {
        t.Fatalf("open generated file: %v", err)
    }
    defer f.Close()

    sig := make([]byte, 8)
    if _, err := io.ReadFull(f, sig); err != nil {
        t.Fatalf("read signature: %v", err)
    }
    pngSig := []byte{137, 80, 78, 71, 13, 10, 26, 10}
    for i := range pngSig {
        if sig[i] != pngSig[i] {
            t.Fatalf("generated file is not a PNG (bad signature)")
        }
    }
}

func TestExampleIsPNG(t *testing.T) {
    f, err := os.Open("../../testdata/example.png")
    if err != nil {
        t.Fatalf("open example png: %v", err)
    }
    defer f.Close()

    sig := make([]byte, 8)
    if _, err := io.ReadFull(f, sig); err != nil {
        t.Fatalf("read signature: %v", err)
    }
    pngSig := []byte{137, 80, 78, 71, 13, 10, 26, 10}
    for i := range pngSig {
        if sig[i] != pngSig[i] {
            t.Fatalf("testdata/example.png is not a PNG (bad signature)")
        }
    }
}
