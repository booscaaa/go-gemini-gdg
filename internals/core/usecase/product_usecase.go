package usecase

import (
	"context"

	"github.com/booscaaa/go-gemini-gdg/internals/core/contract"
)

type productUsecase struct {
	productDatabaseRepository contract.ProductDataBaseRepository
	productLLMRepository      contract.ProductLLMRepository
}

// SeedProducts implements contract.ProductScraperUseCase.
func (usecase *productUsecase) GetMenu(ctx context.Context) (*string, error) {
	products, err := usecase.productDatabaseRepository.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return usecase.productLLMRepository.GetMenu(ctx, products)
}

func NewProductUsecase(
	productDatabaseRepository contract.ProductDataBaseRepository,
	productLLMRepository contract.ProductLLMRepository,
) contract.ProductUseCase {
	return &productUsecase{
		productDatabaseRepository: productDatabaseRepository,
		productLLMRepository:      productLLMRepository,
	}
}
