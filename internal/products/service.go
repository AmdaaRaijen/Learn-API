package products

import (
	"context"

	repo "github.com/amdaaraijen/Learn-API/internal/adapters/pgsql/sqlc"
)

type Service interface {
	GetListOfProducts(ctx context.Context) ([]repo.Product, error)
}

type service struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) *service {
	return  &service{
		repo: repo,
	}
}

func (s *service) GetListOfProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}