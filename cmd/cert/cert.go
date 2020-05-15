package main

import (
	"log"
	"os"

	"github.com/hsmtkk/http2/pkg/cert"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s cert key", os.Args[0])
	}
	certPath := os.Args[1]
	keyPath := os.Args[2]
	if err := cert.New().MakeCert(certPath, keyPath); err != nil {
		log.Fatal(err)
	}
}
