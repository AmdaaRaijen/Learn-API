package auth

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

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequestParams

	err := json.Read(r, &req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	registeredUser, err := h.service.ResgisterUser(r.Context(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.Write(w, http.StatusCreated, registeredUser)
}
