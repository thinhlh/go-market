//go:build wireinject
// +build wireinject

package core

import (
	"github.com/google/wire"
	"github.com/thinhlh/go-market/internal/app/product/application"
	"github.com/thinhlh/go-market/internal/app/product/infrastructure"
	"github.com/thinhlh/go-market/internal/app/product/presentation"
)

type Initialization struct {
	productController presentation.ProductController
	productService    application.ProductService
	productRepository infrastructure.ProductRepository
}

// func NewInitialization(
// 	productController presentation.ProductController,
// 	productService application.ProductService,
// 	productRepository infrastructure.ProductRepository,

// ) *Initialization {
// 	return &Initialization{}
// }

func NewInitialization() Initialization {
	wire.Build(
		presentation.NewProductController,
		application.NewProductService,
		infrastructure.NewProductRepository,
	)
	return Initialization{}
}
