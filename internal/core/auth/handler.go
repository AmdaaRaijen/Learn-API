package auth

import (
	"net/http"

	"github.com/amdaaraijen/Learn-API/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(s *Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var tempRegister registerRequestParams

	err := json.Read(r, &tempRegister)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Write(w, http.StatusCreated, tempRegister)
}
