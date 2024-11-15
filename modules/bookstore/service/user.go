package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/raymondwongso/gogox/errorx"
	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

type UserService struct {
	repo      UserRepository
	validator *validator.Validate
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *UserService) CreateUser(ctx context.Context, params entity.CreateUserParam) (*entity.User, error) {
	if err := s.validator.Struct(params); err != nil {
		return nil, errorx.ErrInvalidParameter("Email is invalid")
	}

	return s.repo.CreateUser(ctx, params.Email)
}
