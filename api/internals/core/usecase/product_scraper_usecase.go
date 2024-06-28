package usecase

import (
	"context"
	"slices"
	"time"

	"github.com/booscaaa/go-gemini-gdg/api/internals/core/contract"
	"github.com/booscaaa/go-gemini-gdg/api/internals/core/domain"
)

type scraperUsecase struct {
	didiosRepository          contract.ProductScraperRepository
	potatosRepository         contract.ProductScraperRepository
	productDatabaseRepository contract.ProductDataBaseRepository
}

// SeedProducts implements contract.ProductScraperUseCase.
func (usecase *scraperUsecase) SeedProducts(ctx context.Context) ([]domain.Product, error) {
	currenteDate := time.Now()
	productsCreated := []domain.Product{}

	potatosProducts, err := usecase.potatosRepository.FindProducts(ctx)
	if err != nil {
		return nil, err
	}

	didiosProducts, err := usecase.didiosRepository.FindProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := slices.Concat(potatosProducts, didiosProducts)

	for _, product := range products {
		product.InsertedAt = currenteDate
		productCreated, err := usecase.productDatabaseRepository.Create(ctx, product)
		if err != nil {
			return nil, err
		}

		productsCreated = append(productsCreated, *productCreated)
	}

	return productsCreated, nil
}

func NewProductScraperUsecase(
	didiosRepository contract.ProductScraperRepository,
	potatosRepository contract.ProductScraperRepository,
	productDatabaseRepository contract.ProductDataBaseRepository,
) contract.ProductScraperUseCase {
	return &scraperUsecase{
		didiosRepository:          didiosRepository,
		potatosRepository:         potatosRepository,
		productDatabaseRepository: productDatabaseRepository,
	}
}
