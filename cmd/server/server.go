package main

import (
	"log"
	"os"

	"github.com/hsmtkk/http2/pkg/server"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s cert key", os.Args[0])
	}
	cert := os.Args[1]
	key := os.Args[2]
	srv := server.New().MakeServer()
	srv.ListenAndServeTLS(cert, key)
}
