package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/matheussbaraglini/hash-challenge/internal/domain/product"
	"github.com/matheussbaraglini/hash-challenge/internal/infrastructure/server/http"
	"github.com/matheussbaraglini/hash-challenge/internal/infrastructure/storage/memory"
	"github.com/matheussbaraglini/hash-challenge/pkg/env"
)

const (
	envVarServerPort = "SERVER_PORT"

	defaultPort = "5001"
)

func main() {
	log := log.New(os.Stderr, "", log.LstdFlags)

	// Storage
	productStorage, err := memory.NewMemoryProductStorage(getProductsFile())
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Service
	checkoutService := product.NewCheckoutService(productStorage, log)

	// HTTP server
	handler := http.NewHandler(checkoutService)

	server := http.New(handler, getServerPort(), *log)
	server.ListenAndServe()

	// graceful shutdown
	forever := make(chan os.Signal, 1)
	<-forever
	server.Shutdown()
}

func getProductsFile() *os.File {
	filePath, err := filepath.Abs("products.json")
	if err != nil {
		log.Fatalf("failed to load products file: %v", err)
	}

	productsFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to load products file: %v", err)
	}

	return productsFile
}

func getServerPort() string {
	return env.GetString(envVarServerPort, defaultPort)
}
