package service

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string) (*entity.User, error)
}

type BookRepository interface {
	GetBooks(ctx context.Context, arg entity.GetBooksParams) ([]entity.Book, error)
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, tx pgx.Tx, arg entity.CreateOrderParams) (*entity.Order, error)
	CreateOrderItem(ctx context.Context, tx pgx.Tx, params entity.CreateOrderItemParams) (*entity.OrderItem, error)
	GetMyOrders(ctx context.Context, arg entity.GetMyOrdersParams) ([]entity.Order, error)
	FindBook(ctx context.Context, tx pgx.Tx, id int64) (*entity.Book, error)
}
