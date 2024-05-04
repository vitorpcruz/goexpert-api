package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/vitorpcruz/goexpert/9-APIS/internal/dto"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/entity"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/infra/database"
)

type ProductHandler struct {
	ProductRepository database.ProductRepositoryInterface
}

func NewProductHandler(db database.ProductRepositoryInterface) *ProductHandler {
	return &ProductHandler{ProductRepository: db}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productDto dto.CreateProductInput

	err := json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := entity.NewProduct(productDto.Name, productDto.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductRepository.Create(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
