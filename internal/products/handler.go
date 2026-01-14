package products

import (
	"log/slog"
	"net/http"

	"github.com/amdaaraijen/Learn-API/internal/json"
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