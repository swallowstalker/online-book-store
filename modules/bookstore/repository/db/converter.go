package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

type QuerierWithTx interface {
	CreateOrder(ctx context.Context, userID int64) (*CreateOrderRow, error)
	CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (*OrderItem, error)
	CreateUser(ctx context.Context, email string) (*User, error)
	FindBook(ctx context.Context, id int64) (*Book, error)
	FindUser(ctx context.Context, email string) (*User, error)
	FindUserByToken(ctx context.Context, token pgtype.Text) (*User, error)
	GetBooks(ctx context.Context, arg GetBooksParams) ([]*Book, error)
	GetMyOrders(ctx context.Context, arg GetMyOrdersParams) ([]*GetMyOrdersRow, error)
	WrapTx(tx pgx.Tx) QuerierWithTx
}

func (q *Queries) WrapTx(tx pgx.Tx) QuerierWithTx {
	return q.WithTx(tx)
}

func (b *Book) ToEntity() *entity.Book {
	return &entity.Book{
		ID:   b.ID,
		Name: b.Name,
	}
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:    u.ID,
		Email: u.Email,
	}
}

func (o *CreateOrderRow) ToEntity() *entity.Order {
	return &entity.Order{
		ID:        o.ID,
		UserID:    o.UserID,
		CreatedAt: o.CreatedAt.Time,
	}
}

func (o *OrderItem) ToEntity() *entity.OrderItem {
	return &entity.OrderItem{
		ID:        o.ID,
		OrderID:   o.OrderID,
		BookID:    o.BookID,
		Amount:    o.Amount,
		CreatedAt: o.CreatedAt.Time,
	}
}
