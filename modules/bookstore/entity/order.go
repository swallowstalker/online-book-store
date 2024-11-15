package entity

import "time"

type Order struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Email     string    `json:"email"`
	BookID    int64     `json:"book_id"`
	BookName  string    `json:"book_name"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateOrderParams struct {
	UserID int64 `validate:"required,gt=0"`
	BookID int64 `json:"book_id"           validate:"required,gt=0"`
	Amount int64 `json:"amount"            validate:"required,gt=0"`
}

type GetMyOrdersParams struct {
	UserID int64 `validate:"required,gt=0"`
	Limit  int64 `validate:"gt=0"`
	Offset int64 `validate:"gte=0"`
}
