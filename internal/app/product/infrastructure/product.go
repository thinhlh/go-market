package infrastructure

import (
	"fmt"

	"github.com/thinhlh/go-market/internal/app/product/domain"
	"github.com/thinhlh/go-market/internal/core/database"
)

type ProductRepository interface {
	GetAllProducts() []domain.Product
	GetProductById() domain.Product
}

func NewProductRepository(db *database.Database) ProductRepository {
	return ProductRepositoryImpl{DB: db}
}

type ProductRepositoryImpl struct {
	DB *database.Database
}

func (r ProductRepositoryImpl) GetAllProducts() []domain.Product {
	type result struct {
		Name string
	}

	var res result
	r.DB.Raw("SELECT 'Jamie' as name").Scan(&res)

	return make([]domain.Product, 0)
}

func (r ProductRepositoryImpl) GetProductById() domain.Product {
	type result struct {
		Name string
	}

	var res result
	r.DB.Exec("SELECT 'Jamie' as name").Scan(&res)

	fmt.Println(res)

	return make([]domain.Product, 0)[0]
}
