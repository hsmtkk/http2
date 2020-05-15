package server

import (
	"net/http"

	"golang.org/x/net/http2"
)

type ServerMaker interface {
	MakeServer() *http.Server
}

func New() ServerMaker {
	return &serverMakerImpl{}
}

type serverMakerImpl struct{}

func (maker *serverMakerImpl) MakeServer() *http.Server {
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: &MyHandler{},
	}
	http2.ConfigureServer(&server, &http2.Server{})
	return &server
}
