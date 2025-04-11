package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fist-it/sonda/internal/config"
	// "github.com/fist-it/sonda/internal/checker"
	// "github.com/fist-it/sonda/internal/metrics"
)

func main() {
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Logging the loaded configuration {{{
	log.Printf("Loaded config: %+v", cfg)
	log.Printf("Number of services: %d", len(cfg.Services))
	for _, service := range cfg.Services {
		log.Printf("Service: %s, URL: %s, Interval: %d, Timeout: %d",
			service.Name, service.URL, service.Interval, service.Timeout)
	}
	log.Printf("Port: %d", cfg.Port)
	// }}}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)


	// Checkers/metrics server

	<-sigChan
	log.Println("Received shutdown signal, shutting down...")
}

