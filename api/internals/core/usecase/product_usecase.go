package usecase

import (
	"context"

	"github.com/booscaaa/go-gemini-gdg/api/internals/core/contract"
	"github.com/booscaaa/go-gemini-gdg/api/internals/core/dto"
)

type productUsecase struct {
	productDatabaseRepository contract.ProductDataBaseRepository
	productLLMRepository      contract.ProductLLMRepository
}

// SearchForTips implements contract.ProductUseCase.
func (usecase *productUsecase) SearchForTips(ctx context.Context) (*dto.AlexaResponse, error) {
	products, err := usecase.productDatabaseRepository.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	output, err := usecase.productLLMRepository.GetMenu(ctx, products)
	if err != nil {
		return nil, err
	}

	return &dto.AlexaResponse{
		Version: "1.0",
		Response: dto.AlexaResponseContent{
			OutputSpeech: dto.AlexaOutputSpeech{
				Type: "PlainText",
				Text: *output,
			},
			ShouldEndSession: true,
		},
	}, nil
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
