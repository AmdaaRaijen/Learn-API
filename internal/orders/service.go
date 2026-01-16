package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/amdaaraijen/Learn-API/internal/adapters/pgsql/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound = errors.New("Product not found")
	ErrProductOutOfStock = errors.New("Product out of stock")
	ErrCustomerNotFound = errors.New("Customer not found")
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

	tx, err := s.db.Begin(ctx)
	
	if err != nil {
		return repo.Order{}, err
	}

	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	
	customer, err := qtx.GetCustomerById(ctx, tempOrder.CustomerId)

	if err != nil {
		return repo.Order{}, ErrCustomerNotFound
	}

	order, err := qtx.CreateOrder(ctx, customer.ID)



	for _, item := range tempOrder.Items {
		product, err := qtx.FindProductByID(ctx, item.ProductID)

		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductOutOfStock
		}

		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID: order.ID,
			ProductID: product.ID,
			Quantity: item.Quantity,
			Price: product.Price,
		})

		if err != nil {
			return repo.Order{}, err
		}

		qtx.UpdateProduct(ctx, repo.UpdateProductParams{
			ID: item.ProductID,
			Name: product.Name,
			Price: product.Price,
			Quantity: product.Quantity - item.Quantity,
		})
	}

	tx.Commit(ctx)

	return order, nil
}

