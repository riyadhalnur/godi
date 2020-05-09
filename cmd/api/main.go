package main

import (
	"github.com/simple-go-server/pkg/server"
)

func main() {
	cfg := &server.Config{
		Port:      "3001",
		Timeout:   30,
		StaticDir: "./../../static",
	}

	srv := server.NewServer(cfg)
	srv.Listen()
}
