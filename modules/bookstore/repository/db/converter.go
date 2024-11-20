package db

import (
	"encoding/json"

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
	details := []entity.BookAmount{}
	err := json.Unmarshal(o.Details, &details)
	if err != nil {
		details = []entity.BookAmount{}
	}
	return &entity.Order{
		ID:        o.ID,
		UserID:    o.UserID,
		Details:   details,
		CreatedAt: o.CreatedAt.Time,
	}
}

func (o *GetMyOrdersRow) ToEntity() *entity.Order {
	details := []entity.BookAmount{}
	err := json.Unmarshal(o.Details, &details)
	if err != nil {
		details = []entity.BookAmount{}
	}
	return &entity.Order{
		ID:        o.ID,
		UserID:    o.UserID,
		Details:   details,
		CreatedAt: o.CreatedAt.Time,
		Email:     o.Email,
	}
}
