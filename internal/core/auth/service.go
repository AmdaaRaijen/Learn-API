package auth

import (
	"context"
	"errors"
	"time"

	repo "github.com/amdaaraijen/Learn-API/internal/adapters/pgsql/sqlc"
	"github.com/amdaaraijen/Learn-API/internal/encrypt"
	"github.com/amdaaraijen/Learn-API/internal/pkg/token"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrInvalidCreds = errors.New("Incorrect email or password given")
)

type Service interface {
	ResgisterUser(ctx context.Context, req registerRequestParams) (repo.Customer, error)
	LoginUser(ctx context.Context, req loginRequestParams) (string, error)
}

type service struct {
	repo repo.Querier
	jwt  token.JWTMaker
}

func NewService(repo repo.Querier, jwt token.JWTMaker) *service {
	return &service{
		repo: repo,
		jwt:  jwt,
	}
}

func (s *service) ResgisterUser(ctx context.Context, req registerRequestParams) (repo.Customer, error) {
	hashed, err := encrypt.HashPassword(req.Password)

	if err != nil {
		return repo.Customer{}, err
	}

	user, err := s.repo.CreateUser(ctx, repo.CreateUserParams{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: pgtype.Text{String: req.PhoneNumber, Valid: req.PhoneNumber != ""},
		Password:    hashed,
	})

	if err != nil {
		return repo.Customer{}, err
	}

	return user, nil
}

func (s *service) LoginUser(ctx context.Context, req loginRequestParams) (string, error) {
	user, err := s.repo.GetCustomerByEmail(ctx, req.Email)

	if err != nil {
		return "", ErrInvalidCreds
	}

	err = encrypt.ComparePassword(user.Password, req.Password)

	if err != nil {
		return "", ErrInvalidCreds
	}

	token, err := s.jwt.GenerateToken(user.ID, time.Hour*1)

	return token, nil
}
