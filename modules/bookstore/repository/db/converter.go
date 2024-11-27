package db

import (
	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

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
