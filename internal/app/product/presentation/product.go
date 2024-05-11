package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thinhlh/go-market/internal/app/product/application"
	"github.com/thinhlh/go-market/internal/core/dto"
)

type ProductController struct {
	productService application.ProductService
}

func NewProductController(service application.ProductService) ProductController {
	return ProductController{service}
}

func (controller ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	result := controller.productService.GetAllProducts()

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(dto.NewSuccessResponse(result))
}

func (controller ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		panic("missing product id")
	}
	result := controller.productService.GetProductById(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
