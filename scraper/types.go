// Package scraper contains all types used in the scraper package.
package scraper

// WhiskyConfig defines the scraping configuration of a whisky site.
type WhiskyConfig struct {
	URL           string `yaml:"url"`
	UserAgent     string `yaml:"userAgent"`
	PriceSelector string `yaml:"priceSelector"`
}

// Config represents the overall application configuration.
type Config struct {
	ScrapeConfig WhiskyConfig `yaml:"scrapeConfig"`
}

// WhiskyScraper defines the interface for scraping whisky data
// and returning a price.
type WhiskyScraper interface {
	Scrape(config WhiskyConfig) (string, error)
}
