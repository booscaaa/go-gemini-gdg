package domain

import "time"

type Product struct {
	ID         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Price      float64   `json:"price" db:"price"`
	InsertedAt time.Time `json:"inserted_at" db:"inserted_at"`
	Company    string    `json:"company" db:"company"`
}
