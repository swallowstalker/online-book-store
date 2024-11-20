package entity

import "time"

type Order struct {
	ID        int64        `json:"id"`
	UserID    int64        `json:"user_id"`
	Email     string       `json:"email,omitempty"`
	Details   []BookAmount `json:"details"`
	CreatedAt time.Time    `json:"created_at"`
}

type BookAmount struct {
	BookID int64 `json:"book_id" validate:"required,gt=0"`
	Amount int64 `json:"amount" validate:"required,gt=0"`
}

type CreateOrderParams struct {
	UserID  int64        `validate:"required,gt=0"`
	Details []BookAmount `json:"details"`
}

type GetMyOrdersParams struct {
	UserID int64 `validate:"required,gt=0"`
	Limit  int64 `validate:"gt=0"`
	Offset int64 `validate:"gte=0"`
}
