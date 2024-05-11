package application

import (
	"github.com/thinhlh/go-market/internal/app/product/domain"
	"github.com/thinhlh/go-market/internal/app/product/infrastructure"
)

type ProductService interface {
	GetAllProducts() []domain.Product
	GetProductById(id string) domain.Product
}

type ProductServiceImpl struct {
	repository infrastructure.ProductRepository
}

func NewProductService(repository infrastructure.ProductRepository) ProductServiceImpl {
	return ProductServiceImpl{repository: repository}
}

func (s ProductServiceImpl) GetAllProducts() []domain.Product {
	return s.repository.GetAllProducts()
}

func (s ProductServiceImpl) GetProductById(id string) domain.Product {
	return s.repository.GetProductById(id)
}
