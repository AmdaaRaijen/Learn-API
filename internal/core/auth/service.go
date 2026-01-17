package auth

import (
	"context"
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
	return req, nil
}
