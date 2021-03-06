package main

import (
	"log"
	"os"
	"strconv"

	"github.com/riyadhalnur/godi/v2/pkg/server"
)

var (
	port      = "3001"
	timeout   = 30
	staticDir = "./../../static"
)

func init() {
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if os.Getenv("TIMEOUT") != "" {
		timeoutStr := os.Getenv("TIMEOUT")
		timeout, _ = strconv.Atoi(timeoutStr)
	}

	if os.Getenv("STATIC_DIR") != "" {
		staticDir = os.Getenv("STATIC_DIR")
	}
}

func main() {
	cfg := &server.Config{
		Port:      port,
		Timeout:   timeout,
		StaticDir: staticDir,
	}

	srv := server.NewServer(cfg)
	if err := srv.Listen(); err != nil {
		log.Fatalln(err)
	}
}
