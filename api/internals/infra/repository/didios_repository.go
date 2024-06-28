package repository

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/booscaaa/go-gemini-gdg/api/internals/core/contract"
	"github.com/booscaaa/go-gemini-gdg/api/internals/core/domain"
	"github.com/playwright-community/playwright-go"
)

const (
	DIDIOS_SITE = "https://dbmarau.menudino.com/"
)

type didiosRepository struct {
	scrapper *playwright.Playwright
}

// FindProducts implements contract.ProductScraperRepository.
func (repository *didiosRepository) FindProducts(context.Context) ([]domain.Product, error) {
	products := []domain.Product{}

	browser, err := repository.scrapper.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		return nil, err
	}

	context, err := browser.NewContext()
	if err != nil {
		return nil, err
	}

	page, err := context.NewPage()
	if err != nil {
		return nil, err
	}

	_, err = page.Goto(DIDIOS_SITE)
	if err != nil {
		return nil, err
	}

	categories, err := page.Locator("#cardapio > section.cardapio-body > div > div.categories > div").All()
	if err != nil {
		return nil, err
	}

	for _, category := range categories {
		category.ScrollIntoViewIfNeeded()

		time.Sleep(time.Millisecond * 500)

		cards, err := category.Locator("div:nth-child(2) > div > div").All()
		if err != nil {
			return nil, err
		}

		for _, card := range cards {
			card.ScrollIntoViewIfNeeded()

			productName, err := card.Locator("a > div > div.media-body > div.name > span").TextContent()
			if err != nil {
				return nil, err
			}

			productPrice, err := card.Locator("a > div > div.media-body > div.priceDescription > div").TextContent()
			if err != nil {
				return nil, err
			}

			productPrice = strings.ReplaceAll(productPrice, "R$ ", "")
			productPrice = strings.ReplaceAll(productPrice, ",", ".")

			product := domain.Product{
				Name:    productName,
				Company: "DIDIOS",
			}

			if price, err := strconv.ParseFloat(productPrice, 64); err == nil {
				product.Price = price
			}

			products = append(products, product)
		}
	}

	return products, nil
}

func NewDidiosRepository(scraper *playwright.Playwright) contract.ProductScraperRepository {
	return &didiosRepository{
		scrapper: scraper,
	}
}
