package cli

import (
    "flag"

    qrcode "github.com/skip2/go-qrcode"
)

type Config struct {
    URL   string
    Out   string
    Size  int
    Level qrcode.RecoveryLevel
}

func Parse() Config {
    url := flag.String("url", "", "URL to generate QR code for")
    out := flag.String("o", "qr.png", "output file (default: qr.png)")
    size := flag.Int("size", 256, "pixel size of output PNG")
    levelStr := flag.String("level", "M", "error correction level: L, M, Q, H")
    flag.Parse()

    var level qrcode.RecoveryLevel
    switch *levelStr {
    case "L":
        level = qrcode.Low
    case "Q":
        level = qrcode.High
    case "H":
        level = qrcode.High
    default:
        level = qrcode.Medium
    }

    return Config{URL: *url, Out: *out, Size: *size, Level: level}
}
