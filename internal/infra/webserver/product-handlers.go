package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/dto"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/entity"
	"github.com/vitorpcruz/goexpert/9-APIS/internal/infra/database"

	pkg "github.com/vitorpcruz/goexpert/9-APIS/pkg/entity"
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

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductRepository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if product.ID, err = pkg.ParseID(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := h.ProductRepository.FindByID(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductRepository.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
