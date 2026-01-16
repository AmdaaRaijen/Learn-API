package products

import (
	"context"

	repo "github.com/amdaaraijen/Learn-API/internal/adapters/pgsql/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	GetListOfProducts(ctx context.Context, name pgtype.Text) ([]repo.Product, error)
	GetProductById(ctx context.Context, id int64) (repo.Product, error)
}

type service struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetListOfProducts(ctx context.Context, name pgtype.Text) ([]repo.Product, error) {
	if name.Valid {
		return s.repo.ListProducts(ctx, name)
	}

	return s.repo.ListProducts(ctx, "")
}

func (s *service) GetProductById(ctx context.Context, id int64) (repo.Product, error) {
	return s.repo.FindProductByID(ctx, id)
}
