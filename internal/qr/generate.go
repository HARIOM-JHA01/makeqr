package qr

import (
    qrcode "github.com/skip2/go-qrcode"
)

func GenerateFile(content string, level qrcode.RecoveryLevel, size int, out string) error {
    return qrcode.WriteFile(content, level, size, out)
}
