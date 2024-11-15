package service

import (
	"context"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string) (*entity.User, error)
}

type BookRepository interface {
	GetBooks(ctx context.Context, arg entity.GetBooksParams) ([]entity.Book, error)
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, arg entity.CreateOrderParams) (*entity.Order, error)
	GetMyOrders(ctx context.Context, arg entity.GetMyOrdersParams) ([]entity.Order, error)
	FindUser(ctx context.Context, email string) (*entity.User, error)
}
