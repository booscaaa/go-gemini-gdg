package di

import (
	"github.com/booscaaa/go-gemini-gdg/api/internals/core/contract"
	"github.com/booscaaa/go-gemini-gdg/api/internals/core/usecase"
	"github.com/booscaaa/go-gemini-gdg/api/internals/infra/cli"
	"github.com/booscaaa/go-gemini-gdg/api/internals/infra/controller"
	"github.com/booscaaa/go-gemini-gdg/api/internals/infra/database"
	"github.com/booscaaa/go-gemini-gdg/api/internals/infra/repository"

	"github.com/google/generative-ai-go/genai"
	"github.com/jmoiron/sqlx"
	"github.com/playwright-community/playwright-go"
)

func InitializeProductCLIDI(scraper *playwright.Playwright, db *sqlx.DB) contract.ProductCLI {
	didiosRepository := repository.NewDidiosRepository(scraper)
	potatosRepository := repository.NewRepository(scraper)
	productDatabaseRepository := database.NewProductDatabase(db)
	productScraperUseCase := usecase.NewProductScraperUsecase(
		didiosRepository,
		potatosRepository,
		productDatabaseRepository,
	)
	productCLI := cli.NewProcuctCLI(productScraperUseCase)
	return productCLI
}

func InitializeProductDI(gm *genai.Client, db *sqlx.DB) contract.ProductUseCase {
	productDatabaseRepository := database.NewProductDatabase(db)
	productLLMRepository := repository.NewGeminiRepository(gm)
	productUseCase := usecase.NewProductUsecase(
		productDatabaseRepository,
		productLLMRepository,
	)
	return productUseCase
}

func InitializeProductControllerDI(gm *genai.Client, db *sqlx.DB) contract.ProductController {
	productDatabaseRepository := database.NewProductDatabase(db)
	productLLMRepository := repository.NewGeminiRepository(gm)
	productUseCase := usecase.NewProductUsecase(
		productDatabaseRepository,
		productLLMRepository,
	)
	productController := controller.NewProductController(productUseCase)
	return productController
}
