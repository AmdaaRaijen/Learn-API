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

	res := registerResponse{
		Message: "User registered successfully",
		Data: userDTO{
			ID:          registeredUser.ID,
			Name:        registeredUser.Name,
			Email:       registeredUser.Email,
			PhoneNumber: &registeredUser.PhoneNumber.String,
			CreatedAt:   registeredUser.CreatedAt.Time,
		},
	}

	json.Write(w, http.StatusCreated, res)
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequestParams

	err := json.Read(r, &req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.service.LoginUser(r.Context(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.Write(w, http.StatusOK, "login success")
}
