package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/raymondwongso/gogox/errorx"
	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

type BookService struct {
	repo      BookRepository
	validator *validator.Validate
}

func NewBookService(repo BookRepository) *BookService {
	return &BookService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *BookService) GetBooks(ctx context.Context, params entity.GetBooksParams) ([]entity.Book, error) {
	if err := s.validator.Struct(params); err != nil {
		return nil, errorx.ErrInvalidParameter("Input is invalid")
	}

	return s.repo.GetBooks(ctx, params)
}
