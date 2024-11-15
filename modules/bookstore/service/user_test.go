package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/raymondwongso/gogox/errorx"
	"github.com/stretchr/testify/suite"
	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
	"github.com/swallowstalker/online-book-store/modules/bookstore/service"
	mock_service "github.com/swallowstalker/online-book-store/test/mock/modules/bookstore/service"
)

type UserServiceTestSuite struct {
	suite.Suite

	repo *mock_service.MockUserRepository
}

func (s *UserServiceTestSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.repo = mock_service.NewMockUserRepository(ctrl)
}

func TestUserServiceRepo(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (s *UserServiceTestSuite) TestCreateUser() {
	ctx := context.Background()
	svc := service.NewUserService(s.repo)

	svcParams := entity.CreateUserParam{
		Email: "someone@test.com",
	}

	rowFromDB := &entity.User{Email: "someone@test.com"}

	s.Run("create user validation error", func() {
		svcParams := entity.CreateUserParam{
			Email: "someone oi @test.com",
		}

		result, err := svc.CreateUser(ctx, svcParams)
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().EqualError(goxErr, "Email is invalid")
	})

	s.Run("create user repo error", func() {
		s.repo.EXPECT().CreateUser(ctx, svcParams.Email).
			Return(nil, errors.New("repo error")).Times(1)

		result, err := svc.CreateUser(ctx, svcParams)
		s.Assert().Nil(result)
		s.Assert().Contains(err.Error(), "repo error")
	})

	s.Run("create user success", func() {
		s.repo.EXPECT().CreateUser(ctx, svcParams.Email).
			Return(rowFromDB, nil).Times(1)

		result, err := svc.CreateUser(ctx, svcParams)
		s.Assert().Nil(err)
		s.Assert().Equal("someone@test.com", result.Email)
	})
}
