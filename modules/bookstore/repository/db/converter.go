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

func (o *FindOrderRow) ToEntity() *entity.Order {
	return &entity.Order{
		ID:        o.ID,
		UserID:    o.UserID,
		Email:     o.Email,
		BookID:    o.BookID,
		BookName:  o.BookName,
		Amount:    o.Amount,
		CreatedAt: o.CreatedAt.Time,
	}
}
