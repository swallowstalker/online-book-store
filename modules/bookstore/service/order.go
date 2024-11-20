package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/raymondwongso/gogox/errorx"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

type OrderService struct {
	repo      OrderRepository
	validator *validator.Validate
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *OrderService) GetOrders(ctx context.Context, params entity.GetMyOrdersParams) ([]entity.Order, error) {
	if err := s.validator.Struct(params); err != nil {
		return nil, errorx.ErrInvalidParameter("Input is invalid")
	}

	return s.repo.GetMyOrders(ctx, params)
}

func (s *OrderService) CreateOrder(ctx context.Context, params entity.CreateOrderParams) (*entity.Order, error) {
	if err := s.validator.Struct(params); err != nil {
		return nil, errorx.ErrInvalidParameter("Input is invalid")
	}

	for _, d := range params.Details {
		if err := s.validator.Struct(d); err != nil {
			return nil, errorx.ErrInvalidParameter("Input is invalid")
		}

		if _, err := s.repo.FindBook(ctx, d.BookID); err != nil {
			return nil, err
		}
	}

	return s.repo.CreateOrder(ctx, params)
}
