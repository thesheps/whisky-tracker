// Package main provides functionality for tracking whisky collections and ratings.
// This application tracks prices for various whisky bottles and alerts users to price drops.
package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"thesheps.dev/whisky-tracker/scraper"
)

// main initializes the whisky tracker module.
func main() {
	// Load configuration from config.yml
	bytes, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg scraper.Config
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create scraper and scrape price
	scraper := &scraper.Scraper{}
	price, err := scraper.Scrape(cfg.ScrapeConfig)
	if err != nil {
		log.Fatalf("Failed to scrape price: %v", err)
	}

	fmt.Printf("Current price: %s\n", price)
}
