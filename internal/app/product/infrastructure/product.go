package infrastructure

import (
	"github.com/thinhlh/go-market/internal/app/product/domain"
	"github.com/thinhlh/go-market/internal/core/database"
)

type ProductRepository interface {
	GetAllProducts() []domain.Product
	GetProductById(id string) domain.Product
}

func NewProductRepository(db *database.Database) ProductRepository {
	return ProductRepositoryImpl{DB: db}
}

type ProductRepositoryImpl struct {
	DB *database.Database
}

func (r ProductRepositoryImpl) GetAllProducts() []domain.Product {
	var products []domain.Product

	r.DB.Table(domain.ProductTableName).Find(&products)

	return products

}

func (r ProductRepositoryImpl) GetProductById(id string) domain.Product {

	var product domain.Product
	r.DB.Table(domain.ProductTableName).First(&product, "id=?", id)

	return product
}
