package main

import (
	"flag"
	"fmt"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	url := flag.String("url", "", "URL to generate QR code for")
	out := flag.String("o", "qr.png", "output file (default: qr.png)")
	flag.Parse()

	if *url == "" {
		fmt.Println("Usage: makeqr --url https://example.com [-o file]")
		os.Exit(1)
	}

	err := qrcode.WriteFile(*url, qrcode.Medium, 256, *out)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("QR saved at:", *out)
}
