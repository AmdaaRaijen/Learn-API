package products

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/amdaaraijen/Learn-API/internal/json"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{service: s}
}

func(h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	textName := pgtype.Text{String: name, Valid: name != ""}

	products, err := h.service.GetListOfProducts(r.Context(), textName)

	if err != nil {
		slog.Error("Error getting list product", "error", err)
		http.Error(w, "Failed to get product", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func(h *handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")

	productIdInt, err := strconv.ParseInt(productId, 10, 64)

	if err != nil {
		http.Error(w, "Invalid product id", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductById(r.Context(), productIdInt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		slog.Error("Error getting product by id", "error", err)
		http.Error(w, "Failed to get product", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}