package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/raymondwongso/gogox/errorx"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
	"github.com/swallowstalker/online-book-store/modules/bookstore/repository"
)

type OrderService struct {
	repo      OrderRepository
	validator *validator.Validate
	txStarter repository.TxStarter
}

func NewOrderService(repo OrderRepository, txStarter repository.TxStarter) *OrderService {
	return &OrderService{
		repo:      repo,
		validator: validator.New(),
		txStarter: txStarter,
	}
}

func (s *OrderService) GetOrders(ctx context.Context, params entity.GetMyOrdersParams) ([]entity.Order, error) {
	if err := s.validator.Struct(params); err != nil {
		return nil, errorx.ErrInvalidParameter("Input is invalid")
	}

	return s.repo.GetMyOrders(ctx, params)
}

func (s *OrderService) CreateOrder(ctx context.Context, params entity.CreateOrderParams) (*entity.Order, error) {
	var err error
	if err = s.validator.Struct(params); err != nil {
		return nil, errorx.ErrInvalidParameter("Input is invalid")
	}

	var tx repository.Transactionable
	tx, err = s.txStarter(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	for _, d := range params.Items {
		if err = s.validator.Struct(d); err != nil {
			return nil, errorx.ErrInvalidParameter("Input is invalid")
		}

		if _, err = s.repo.FindBook(ctx, tx, d.BookID); err != nil {
			return nil, err
		}
	}

	var order *entity.Order
	order, err = s.repo.CreateOrder(ctx, tx, params)
	if err != nil {
		return nil, err
	}

	for _, itemInParam := range params.Items {
		var item *entity.OrderItem
		item, err = s.repo.CreateOrderItem(ctx, tx, entity.CreateOrderItemParams{
			OrderID: order.ID,
			BookID:  itemInParam.BookID,
			Amount:  itemInParam.Amount,
		})
		if err != nil {
			return nil, err
		}

		order.Items = append(order.Items, *item)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return order, nil
}
