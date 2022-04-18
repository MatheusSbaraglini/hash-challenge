package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/matheussbaraglini/hash-challenge/internal/domain/product"
	"github.com/matheussbaraglini/hash-challenge/internal/infrastructure/client"
	"github.com/matheussbaraglini/hash-challenge/internal/infrastructure/server/http"
	"github.com/matheussbaraglini/hash-challenge/internal/infrastructure/storage/memory"
	"github.com/matheussbaraglini/hash-challenge/pkg/env"
)

const (
	dateLayout                 = "02/01/2006 15:04:05"
	envVarServerPort           = "SERVER_PORT"
	envVarDiscountServiceURL   = "DISCOUNT_SERVICE_URL"
	envVarStartDateBlackFriday = "START_DATE_BLACK_FRIDAY"
	envVarEndDateBlackFriday   = "END_DATE_BLACK_FRIDAY"

	defaultPort = "4040"
)

func main() {
	log := log.New(os.Stderr, "", log.LstdFlags)

	if err := env.CheckRequired(*log, envVarStartDateBlackFriday, envVarEndDateBlackFriday); err != nil {
		log.Fatal(err)
	}

	// Storage
	productStorage, err := memory.NewMemoryProductStorage(getProductsFile())
	if err != nil {
		log.Fatal(err)
	}

	// Clients
	discountClient := client.NewDiscountClient(getDiscountServiceURL())

	// Service
	funcNow := func() time.Time {
		return time.Now()
	}

	blackFridayStart, err := time.ParseInLocation(dateLayout, getStartDateBlackFriday(), time.Local)
	if err != nil {
		log.Fatalf("invalid start date format: %v", err)
	}

	blackFridayEnd, err := time.ParseInLocation(dateLayout, getEndDateBlackFriday(), time.Local)
	if err != nil {
		log.Fatalf("invalid end date format: %v", err)
	}

	if blackFridayStart.After(blackFridayEnd) {
		log.Fatal("black friday start date must be before of the end date")
	}

	checkoutService := product.NewCheckoutService(discountClient, productStorage, funcNow, blackFridayStart, blackFridayEnd, log)

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

func getStartDateBlackFriday() string {
	return env.GetString(envVarStartDateBlackFriday)
}

func getEndDateBlackFriday() string {
	return env.GetString(envVarEndDateBlackFriday)
}

func getDiscountServiceURL() string {
	return env.GetString(envVarDiscountServiceURL)
}
