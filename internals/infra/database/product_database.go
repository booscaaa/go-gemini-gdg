package database

import (
	"context"

	"github.com/booscaaa/go-gemini-gdg/internals/core/contract"
	"github.com/booscaaa/go-gemini-gdg/internals/core/domain"
	"github.com/jmoiron/sqlx"
)

type productDatabaseRepository struct {
	database *sqlx.DB
}

// Create implements contract.ProductDataBaseRepository.
func (repository *productDatabaseRepository) Create(ctx context.Context, input domain.Product) (*domain.Product, error) {
	output := domain.Product{}
	query := "INSERT INTO product (name, price, company, inserted_at) VALUES ($1, $2, $3, $4) RETURNING *;"

	err := repository.database.QueryRowxContext(
		ctx,
		query,
		input.Name,
		input.Price,
		input.Company,
		input.InsertedAt,
	).StructScan(&output)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

// Fetch implements contract.ProductDataBaseRepository.
func (repository *productDatabaseRepository) Fetch(ctx context.Context) ([]domain.Product, error) {
	output := []domain.Product{}
	query := `
		SELECT p.* FROM product p
		INNER JOIN (
			SELECT  MAX(inserted_at) AS MAXDATE
			FROM product
		) p2
		ON p.inserted_at = p2.MAXDATE
	`

	err := repository.database.SelectContext(ctx, &output, query)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func NewProductDatabase(database *sqlx.DB) contract.ProductDataBaseRepository {
	return &productDatabaseRepository{database: database}
}
