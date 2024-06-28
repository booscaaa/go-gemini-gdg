package contract

import (
	"context"
	"net/http"

	"github.com/booscaaa/go-gemini-gdg/api/internals/core/domain"
	"github.com/booscaaa/go-gemini-gdg/api/internals/core/dto"
)

type ProductScraperRepository interface {
	FindProducts(context.Context) ([]domain.Product, error)
}

type ProductDataBaseRepository interface {
	Fetch(context.Context) ([]domain.Product, error)
	Create(context.Context, domain.Product) (*domain.Product, error)
}

type ProductScraperUseCase interface {
	SeedProducts(context.Context) ([]domain.Product, error)
}

type ProductUseCase interface {
	GetMenu(context.Context) (*string, error)
	SearchForTips(context.Context) (*dto.AlexaResponse, error)
}

type ProductLLMRepository interface {
	GetMenu(context.Context, []domain.Product) (*string, error)
}

type ProductController interface {
	SearchForTips(http.ResponseWriter, *http.Request)
}

type ProductCLI interface {
	SeedProducts(context.Context)
}
