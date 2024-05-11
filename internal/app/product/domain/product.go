package domain

import (
	"errors"

	"github.com/google/uuid"
)

// Root Entity & Aggregate
type Product struct {
	ID    uuid.UUID `json:"id"`
	Name  *string   `json:"name"`
	Price *float64  `json:"price"`
}

var ProductTableName string = "product"

var (
	ErrInvalidProductName = errors.New("invalid product name")
)

func New(name string, price float64, properties ProductProperty) (*Product, error) {

	if name == "" {
		return nil, ErrInvalidProductName
	}
	product := &Product{
		Name:  &name,
		Price: &price,
	}

	return product, nil
}
