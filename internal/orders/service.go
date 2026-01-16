package orders

import (
	"context"
	"fmt"

	repo "github.com/amdaaraijen/Learn-API/internal/adapters/pgsql/sqlc"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error)
}

type service struct {
	repo *repo.Queries
	db *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return  &service{
		repo: repo,
		db: db,
	}
}

func (s *service) PlaceOrder(ctx context.Context, tempOrder createOrderParams) (repo.Order, error) {
	if tempOrder.CustomerId == 0 {
		return repo.Order{}, fmt.Errorf("CustomerId is required!")
	}

	if len(tempOrder.Items) == 0 {
		return repo.Order{}, fmt.Errorf("Need at least 1 item to order!")
	}

	return repo.Order{}, nil
}

