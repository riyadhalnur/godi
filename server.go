package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	// PORT the tcp port the server will listen on. Defaults to 8080
	PORT = "3001"
	// STATIC the folder to server static files from. Defaults to /static
	STATIC = "/static"

	defaultTimeout = 30 * time.Second
)

func init() {
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}

	if os.Getenv("STATIC") != "" {
		STATIC = os.Getenv("STATIC")
	}
}

func main() {
	srv := createServer()

	go func() {
		log.Printf("Listening on port %s\n", PORT)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// listen for terminate signals from the OS
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	<-c

	// wait for active connections to finish their jobs
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

func createServer() *http.Server {
	router := NewRouter()
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", PORT),
		WriteTimeout: defaultTimeout,
		ReadTimeout:  defaultTimeout,
		IdleTimeout:  defaultTimeout,
		Handler:      router,
	}
}
