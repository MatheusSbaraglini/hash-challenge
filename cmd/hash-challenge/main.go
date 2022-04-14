package main

import (
	"log"
	"os"

	"github.com/matheussbaraglini/hash-challenge/internal/infrastructure/server/http"
	"github.com/matheussbaraglini/hash-challenge/pkg/env"
)

const (
	envVarServerPort = "SERVER_PORT"

	defaultPort = "5001"
)

func main() {
	log := log.New(os.Stderr, "", log.LstdFlags)

	// HTTP server
	handler := http.NewHandler()

	server := http.New(handler, getServerPort(), *log)
	server.ListenAndServe()

	// graceful shutdown
	forever := make(chan os.Signal, 1)
	<-forever
	server.Shutdown()
}

func getServerPort() string {
	return env.GetString(envVarServerPort, defaultPort)
}
