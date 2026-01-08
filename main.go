// Package main provides functionality for tracking whisky collections and ratings.
// This application tracks prices for various whisky bottles and alerts users to price drops.
package main

import (
	"fmt"
	"strings"

	"github.com/playwright-community/playwright-go"
)

// main initializes the whisky tracker module.
func main() {
	err := playwright.Install()
	if err != nil {
		panic(err)
	}

	pw, err := playwright.Run()
	if err != nil {
		panic(err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		panic(err)
	}
	defer browser.Close()

	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent: playwright.String(
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		),
	})
	if err != nil {
		panic(err)
	}

	page, err := context.NewPage()
	if err != nil {
		panic(err)
	}

	_, err = page.Goto(
		"https://www.thewhiskyexchange.com/p/85794/bunnahabhain-12-year-old-cask-strength-2025-release",
		playwright.PageGotoOptions{
			WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		},
	)
	if err != nil {
		panic(err)
	}

	priceLocator := page.Locator(".product-action__price")

	text, err := priceLocator.First().TextContent()
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.TrimSpace(text))
}
