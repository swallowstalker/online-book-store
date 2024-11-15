package entity

type Book struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetBooksParams struct {
	Limit  int64 `validate:"gt=0"`
	Offset int64 `validate:"gte=0"`
}
