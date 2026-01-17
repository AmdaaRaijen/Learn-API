package orders

import (
	"net/http"

	"github.com/amdaaraijen/Learn-API/internal/authctx"
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
	Userid := r.Context().Value(authctx.UserIDKey).(int64)

	var tempOrder createOrderParams

	err := json.Read(r, &tempOrder)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := h.service.PlaceOrder(r.Context(), tempOrder, Userid)

	if err != nil {
		if err == ErrProductNotFound || err == ErrProductOutOfStock || err == ErrCustomerNotFound {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusCreated, createdOrder)
}
