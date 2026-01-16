package orders

import (
	"net/http"

	"github.com/amdaaraijen/Learn-API/internal/json"
)


type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}


func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	tempOrder := createOrderParams{}

	err := json.Read(r, &tempOrder)

	if err != nil {
		http.Error(w, "Failed to place order", http.StatusBadRequest)
		return
	}

	createdOrder, err := h.service.PlaceOrder(r.Context(), createOrderParams{})

	if err != nil {
		http.Error(w, "Failed to get product", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, createdOrder)
}