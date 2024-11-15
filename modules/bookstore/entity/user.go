package entity

type User struct {
	ID    int64
	Email string
}

type CreateUserParam struct {
	Email string `json:"email" validate:"required,email"`
}
