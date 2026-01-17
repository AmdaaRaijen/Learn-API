package auth

import (
	"context"

	"github.com/amdaaraijen/Learn-API/internal/encrypt"
)

type Service interface {
	ResgisterUser(ctx context.Context, req registerRequestParams) (registerRequestParams, error)
}

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) ResgisterUser(ctx context.Context, req registerRequestParams) (registerRequestParams, error) {
	hashed, err := encrypt.HashPassword(req.Password)

	if err != nil {
		return registerRequestParams{}, err
	}

	req.Password = hashed

	return req, nil
}
