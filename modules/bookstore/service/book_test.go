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

type BookServiceTestSuite struct {
	suite.Suite

	repo *mock_service.MockBookRepository
}

func (s *BookServiceTestSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.repo = mock_service.NewMockBookRepository(ctrl)
}

func TestBookServiceRepo(t *testing.T) {
	suite.Run(t, new(BookServiceTestSuite))
}

func (s *BookServiceTestSuite) TestGetBooks() {
	ctx := context.Background()
	svc := service.NewBookService(s.repo)

	svcParams := entity.GetBooksParams{
		Limit:  10,
		Offset: 0,
	}

	s.Run("get books validation error", func() {
		svcParams := entity.GetBooksParams{
			Limit: 0,
		}

		result, err := svc.GetBooks(ctx, svcParams)
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().EqualError(goxErr, "Input is invalid")
	})

	s.Run("get books repo error", func() {
		s.repo.EXPECT().GetBooks(ctx, svcParams).
			Return(nil, errors.New("repo error")).Times(1)

		result, err := svc.GetBooks(ctx, svcParams)
		s.Assert().Nil(result)
		s.Assert().Contains(err.Error(), "repo error")
	})

	s.Run("get books success", func() {
		s.repo.EXPECT().GetBooks(ctx, svcParams).
			Return([]entity.Book{}, nil).Times(1)

		result, err := svc.GetBooks(ctx, svcParams)
		s.Assert().Nil(err)
		s.Assert().NotNil(result)
	})
}
