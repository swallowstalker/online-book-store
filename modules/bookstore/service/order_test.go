package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/raymondwongso/gogox/errorx"
	"github.com/stretchr/testify/suite"
	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
	"github.com/swallowstalker/online-book-store/modules/bookstore/service"
	mock_service "github.com/swallowstalker/online-book-store/test/mock/modules/bookstore/service"
)

type OrderServiceTestSuite struct {
	suite.Suite

	repo *mock_service.MockOrderRepository
}

func (s *OrderServiceTestSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.repo = mock_service.NewMockOrderRepository(ctrl)
}

func TestOrderServiceRepo(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}

func (s *OrderServiceTestSuite) TestGetOrders() {
	ctx := context.Background()
	svc := service.NewOrderService(s.repo)
	now := time.Now()

	svcParams := entity.GetMyOrdersParams{
		UserID: 123,
		Limit:  10,
		Offset: 0,
	}

	rowFromDB := []entity.Order{
		{
			ID:        1,
			UserID:    123,
			Email:     "someone@test.com",
			BookID:    99,
			BookName:  "Book A",
			Amount:    1,
			CreatedAt: now,
		},
	}

	s.Run("get order validation error", func() {
		svcParams := entity.GetMyOrdersParams{
			UserID: 0,
			Limit:  0,
			Offset: 0,
		}

		result, err := svc.GetOrders(ctx, svcParams)
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().EqualError(goxErr, "Input is invalid")
	})

	s.Run("get order repo error", func() {
		s.repo.EXPECT().GetMyOrders(ctx, svcParams).
			Return(nil, errors.New("repo error")).Times(1)

		result, err := svc.GetOrders(ctx, svcParams)
		s.Assert().Nil(result)
		s.Assert().Contains(err.Error(), "repo error")
	})

	s.Run("get order success", func() {
		s.repo.EXPECT().GetMyOrders(ctx, svcParams).
			Return(rowFromDB, nil).Times(1)

		result, err := svc.GetOrders(ctx, svcParams)
		s.Assert().Nil(err)
		s.Assert().Equal(int64(99), result[0].BookID)
	})
}

func (s *OrderServiceTestSuite) TestCreateOrder() {
	ctx := context.Background()
	svc := service.NewOrderService(s.repo)
	now := time.Now()

	svcParams := entity.CreateOrderParams{
		UserID: 123,
		BookID: 99,
		Amount: 3,
	}

	rowFromDB := &entity.Order{
		ID:        1,
		UserID:    123,
		Email:     "someone@test.com",
		BookID:    99,
		BookName:  "Book A",
		Amount:    1,
		CreatedAt: now,
	}

	s.Run("create order validation error", func() {
		svcParams := entity.CreateOrderParams{
			UserID: 0,
			BookID: 0,
			Amount: 0,
		}

		result, err := svc.CreateOrder(ctx, svcParams)
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().EqualError(goxErr, "Input is invalid")
	})

	s.Run("create order repo error", func() {
		s.repo.EXPECT().CreateOrder(ctx, svcParams).
			Return(nil, errors.New("repo error")).Times(1)

		result, err := svc.CreateOrder(ctx, svcParams)
		s.Assert().Nil(result)
		s.Assert().Contains(err.Error(), "repo error")
	})

	s.Run("create order success", func() {
		s.repo.EXPECT().CreateOrder(ctx, svcParams).
			Return(rowFromDB, nil).Times(1)

		result, err := svc.CreateOrder(ctx, svcParams)
		s.Assert().Nil(err)
		s.Assert().Equal(int64(99), result.BookID)
	})
}
