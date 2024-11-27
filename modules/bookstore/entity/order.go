package entity

import "time"

type Order struct {
	ID        int64       `json:"id"`
	UserID    int64       `json:"user_id"`
	Email     string      `json:"email,omitempty"`
	Items     []OrderItem `json:"items"`
	CreatedAt time.Time   `json:"created_at"`
}

type OrderItem struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"order_id"`
	BookID    int64     `json:"book_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type BookAmount struct {
	BookID int64 `json:"book_id" validate:"required,gt=0"`
	Amount int64 `json:"amount" validate:"required,gt=0"`
}

type CreateOrderParams struct {
	UserID int64                   `validate:"required,gt=0"`
	Items  []CreateOrderItemParams `json:"items"`
}

type CreateOrderItemParams struct {
	OrderID int64
	BookID  int64 `json:"book_id" validate:"required,gt=0"`
	Amount  int64 `json:"amount" validate:"required,gt=0"`
}

type GetMyOrdersParams struct {
	UserID int64 `validate:"required,gt=0"`
	Limit  int64 `validate:"gt=0"`
	Offset int64 `validate:"gte=0"`
}
