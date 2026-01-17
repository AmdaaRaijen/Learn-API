package auth

import (
	"context"

	repo "github.com/amdaaraijen/Learn-API/internal/adapters/pgsql/sqlc"
	"github.com/amdaaraijen/Learn-API/internal/encrypt"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	ResgisterUser(ctx context.Context, req registerRequestParams) (repo.Customer, error)
}

type service struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) ResgisterUser(ctx context.Context, req registerRequestParams) (repo.Customer, error) {
	hashed, err := encrypt.HashPassword(req.Password)

	if err != nil {
		return repo.Customer{}, err
	}

	user, err := s.repo.CreateUser(ctx, repo.CreateUserParams{
		Name:        req.Name,
		Email:       pgtype.Text{String: req.Email, Valid: req.Email != ""},
		PhoneNumber: pgtype.Text{String: req.PhoneNumber, Valid: req.PhoneNumber != ""},
		Password:    hashed,
	})

	if err != nil {
		return repo.Customer{}, err
	}

	return user, nil
}
