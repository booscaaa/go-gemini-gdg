package di

import (
	"github.com/booscaaa/go-gemini-gdg/internals/core/contract"
	"github.com/booscaaa/go-gemini-gdg/internals/core/usecase"
	"github.com/booscaaa/go-gemini-gdg/internals/infra/database"
	"github.com/booscaaa/go-gemini-gdg/internals/infra/repository/didios"
	"github.com/booscaaa/go-gemini-gdg/internals/infra/repository/gemini"
	"github.com/booscaaa/go-gemini-gdg/internals/infra/repository/potatos"
	"github.com/google/generative-ai-go/genai"
	"github.com/jmoiron/sqlx"
	"github.com/playwright-community/playwright-go"
)

func InitializeProductScraperDI(scraper *playwright.Playwright, db *sqlx.DB) contract.ProductScraperUseCase {
	didiosRepository := didios.NewRepository(scraper)
	potatosRepository := potatos.NewRepository(scraper)
	productDatabaseRepository := database.NewProductDatabase(db)
	productScraperUseCase := usecase.NewProductScraperUsecase(
		didiosRepository,
		potatosRepository,
		productDatabaseRepository,
	)
	return productScraperUseCase
}

func InitializeProductDI(gm *genai.Client, db *sqlx.DB) contract.ProductUseCase {
	productDatabaseRepository := database.NewProductDatabase(db)
	productLLMRepository := gemini.NewGeminiRepository(gm)
	productUseCase := usecase.NewProductUsecase(
		productDatabaseRepository,
		productLLMRepository,
	)
	return productUseCase
}
