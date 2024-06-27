package contract

import (
	"context"

	"github.com/booscaaa/go-gemini-gdg/internals/core/domain"
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
}

type ProductLLMRepository interface {
	GetMenu(context.Context, []domain.Product) (*string, error)
}
