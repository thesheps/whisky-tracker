package scraper

import (
	"strings"

	"github.com/playwright-community/playwright-go"
)

// Scraper implements the WhiskyScraper interface.
type Scraper struct {
}

// Scrape scrapes the whisky price from the given URL using the provided configuration.
func (s *Scraper) Scrape(config WhiskyConfig) (string, error) {
	err := playwright.Install()
	if err != nil {
		return "", err
	}

	pw, err := playwright.Run()
	if err != nil {
		return "", err
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		return "", err
	}
	defer browser.Close()

	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent: playwright.String(config.UserAgent),
	})
	if err != nil {
		return "", err
	}

	page, err := context.NewPage()
	if err != nil {
		return "", err
	}

	_, err = page.Goto(
		config.URL,
		playwright.PageGotoOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		},
	)
	if err != nil {
		return "", err
	}

	priceLocator := page.Locator(config.PriceSelector)
	text, err := priceLocator.First().TextContent()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(text), nil
}
