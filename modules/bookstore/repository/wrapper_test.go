package repository_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/raymondwongso/gogox/errorx"
	"github.com/stretchr/testify/suite"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
	"github.com/swallowstalker/online-book-store/modules/bookstore/repository"
	"github.com/swallowstalker/online-book-store/modules/bookstore/repository/db"
	mock_repo "github.com/swallowstalker/online-book-store/test/mock/modules/bookstore/repository/db"
)

type WrapperTestSuite struct {
	suite.Suite

	querierRepo *mock_repo.MockQuerier
}

func (s *WrapperTestSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.querierRepo = mock_repo.NewMockQuerier(ctrl)
}

func TestWrapperRepo(t *testing.T) {
	suite.Run(t, new(WrapperTestSuite))
}

func (s *WrapperTestSuite) TestCreateUser() {
	ctx := context.Background()
	now := time.Now()
	wrapper := repository.NewDbWrapperRepo(s.querierRepo)
	expectedUser := &entity.User{
		ID:    123,
		Email: "someone@test.com",
	}

	rowFromDB := &db.User{
		ID:    123,
		Email: "someone@test.com",
		CreatedAt: pgtype.Timestamptz{
			Time:  now,
			Valid: true,
		},
	}

	s.Run("create user got querier error", func() {
		s.querierRepo.EXPECT().CreateUser(ctx, "someone@test.com").
			Return(nil, errors.New("querier error")).Times(1)

		result, err := wrapper.CreateUser(ctx, "someone@test.com")
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().EqualError(goxErr, "internal server error")
		s.Assert().Contains(goxErr.LogError(), "querier error")
	})

	s.Run("no row after create user should find the user", func() {
		s.querierRepo.EXPECT().CreateUser(ctx, "someone@test.com").
			Return(nil, sql.ErrNoRows).Times(1)
		s.querierRepo.EXPECT().FindUser(ctx, "someone@test.com").
			Return(rowFromDB, nil).Times(1)

		result, err := wrapper.CreateUser(ctx, "someone@test.com")
		s.Assert().Equal(expectedUser, result)
		s.Assert().Nil(err)
	})

	s.Run("no row after create user should find the user but user still not found", func() {
		s.querierRepo.EXPECT().CreateUser(ctx, "someone@test.com").
			Return(nil, sql.ErrNoRows).Times(1)
		s.querierRepo.EXPECT().FindUser(ctx, "someone@test.com").
			Return(nil, sql.ErrNoRows).Times(1)

		result, err := wrapper.CreateUser(ctx, "someone@test.com")
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().Equal(errorx.CodeInternal, goxErr.Code)
		s.Assert().EqualError(goxErr, "internal server error")
		s.Assert().Contains(goxErr.LogError(), "insert returns no row but existing user not found: [common.not_found] user not found: sql: no rows in result set")
	})

	s.Run("create user successful", func() {
		s.querierRepo.EXPECT().CreateUser(ctx, "someone@test.com").
			Return(rowFromDB, nil).Times(1)

		result, err := wrapper.CreateUser(ctx, "someone@test.com")
		s.Assert().Equal(expectedUser, result)
		s.Assert().Nil(err)
	})
}

func (s *WrapperTestSuite) TestFindUser() {
	ctx := context.Background()
	now := time.Now()
	wrapper := repository.NewDbWrapperRepo(s.querierRepo)
	expectedUser := &entity.User{
		ID:    123,
		Email: "someone@test.com",
	}

	rowFromDB := &db.User{
		ID:    123,
		Email: "someone@test.com",
		CreatedAt: pgtype.Timestamptz{
			Time:  now,
			Valid: true,
		},
	}

	s.Run("find user got querier error", func() {
		s.querierRepo.EXPECT().FindUser(ctx, "someone@test.com").
			Return(nil, errors.New("querier error")).Times(1)

		result, err := wrapper.FindUser(ctx, "someone@test.com")
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().Equal(errorx.CodeInternal, goxErr.Code)
		s.Assert().EqualError(goxErr, "internal server error")
		s.Assert().Contains(goxErr.LogError(), "[common.internal] internal server error: querier error")
	})

	s.Run("no row should return sql no error", func() {
		s.querierRepo.EXPECT().FindUser(ctx, "someone@test.com").
			Return(nil, sql.ErrNoRows).Times(1)

		result, err := wrapper.FindUser(ctx, "someone@test.com")
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().Equal(errorx.CodeNotFound, goxErr.Code)
		s.Assert().EqualError(goxErr, "user not found")
		s.Assert().Contains(goxErr.LogError(), "[common.not_found] user not found: sql: no rows in result set")
	})

	s.Run("find user successful", func() {
		s.querierRepo.EXPECT().FindUser(ctx, "someone@test.com").
			Return(rowFromDB, nil).Times(1)

		result, err := wrapper.FindUser(ctx, "someone@test.com")
		s.Assert().Equal(expectedUser, result)
		s.Assert().Nil(err)
	})
}

func (s *WrapperTestSuite) TestGetBooks() {
	ctx := context.Background()
	now := time.Now()
	wrapper := repository.NewDbWrapperRepo(s.querierRepo)
	expectedBooks := []entity.Book{
		{
			ID:   123,
			Name: "Book A",
		},
		{
			ID:   124,
			Name: "Book B",
		},
	}
	rowsFromDB := []*db.Book{
		{
			ID:   123,
			Name: "Book A",
			CreatedAt: pgtype.Timestamptz{
				Time:  now,
				Valid: true,
			},
		},
		{
			ID:   124,
			Name: "Book B",
			CreatedAt: pgtype.Timestamptz{
				Time:  now,
				Valid: true,
			},
		},
	}

	querierParams := db.GetBooksParams{
		Limit:  100,
		Offset: 2,
	}
	wrapperParams := entity.GetBooksParams{
		Limit:  100,
		Offset: 2,
	}

	s.Run("get books got querier error", func() {
		s.querierRepo.EXPECT().GetBooks(ctx, querierParams).
			Return(nil, errors.New("querier error")).Times(1)

		result, err := wrapper.GetBooks(ctx, wrapperParams)
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().Equal(errorx.CodeInternal, goxErr.Code)
		s.Assert().EqualError(goxErr, "internal server error")
		s.Assert().Contains(goxErr.LogError(), "[common.internal] internal server error: querier error")
	})

	s.Run("get books successful", func() {
		s.querierRepo.EXPECT().GetBooks(ctx, querierParams).
			Return(rowsFromDB, nil).Times(1)

		result, err := wrapper.GetBooks(ctx, wrapperParams)
		s.Assert().Equal(expectedBooks, result)
		s.Assert().Nil(err)
	})
}

func (s *WrapperTestSuite) TestCreateOrder() {
	ctx := context.Background()
	now := time.Now()
	wrapper := repository.NewDbWrapperRepo(s.querierRepo)

	wrapperParams := entity.CreateOrderParams{
		UserID: 9919,
		Details: []entity.BookAmount{
			{
				BookID: 123,
				Amount: 10,
			},
		},
	}

	b, err := json.Marshal(wrapperParams.Details)
	s.Require().Nil(err)

	querierParams := db.CreateOrderParams{
		UserID:  9919,
		Details: b,
	}

	expectedOrder := &entity.Order{
		ID:     90,
		UserID: 9919,
		Details: []entity.BookAmount{
			{
				BookID: 123,
				Amount: 10,
			},
		},
		CreatedAt: now,
	}

	rowFromDB := &db.CreateOrderRow{
		ID:      90,
		UserID:  9919,
		Details: b,
		CreatedAt: pgtype.Timestamptz{
			Time:  now,
			Valid: true,
		},
	}

	s.Run("create order got querier error", func() {
		s.querierRepo.EXPECT().CreateOrder(ctx, querierParams).
			Return(nil, errors.New("querier error")).Times(1)

		result, err := wrapper.CreateOrder(ctx, wrapperParams)
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().EqualError(goxErr, "internal server error")
		s.Assert().Contains(goxErr.LogError(), "querier error")
	})

	s.Run("create order successful", func() {
		s.querierRepo.EXPECT().CreateOrder(ctx, querierParams).
			Return(rowFromDB, nil).Times(1)

		result, err := wrapper.CreateOrder(ctx, wrapperParams)
		s.Assert().Equal(expectedOrder, result)
		s.Assert().Nil(err)
	})
}

func (s *WrapperTestSuite) TestGetMyOrders() {
	ctx := context.WithValue(context.Background(), entity.UserContextKey{}, int64(9919))
	now := time.Now()
	wrapper := repository.NewDbWrapperRepo(s.querierRepo)
	expectedOrders := []entity.Order{
		{
			ID:     123,
			UserID: 9919,
			Email:  "someone@test.com",
			Details: []entity.BookAmount{
				{
					BookID: 920,
					Amount: 10,
				},
			},
			CreatedAt: now,
		},
		{
			ID:     124,
			UserID: 9919,
			Email:  "someone@test.com",
			Details: []entity.BookAmount{
				{
					BookID: 920,
					Amount: 20,
				},
			},
			CreatedAt: now,
		},
	}

	b1, err := json.Marshal(expectedOrders[0].Details)
	s.Require().Nil(err)
	b2, err := json.Marshal(expectedOrders[1].Details)
	s.Require().Nil(err)

	rowsFromDB := []*db.GetMyOrdersRow{
		{
			ID:      123,
			UserID:  9919,
			Details: b1,
			CreatedAt: pgtype.Timestamptz{
				Time:  now,
				Valid: true,
			},
			Email: "someone@test.com",
		},
		{
			ID:      124,
			UserID:  9919,
			Details: b2,
			CreatedAt: pgtype.Timestamptz{
				Time:  now,
				Valid: true,
			},
			Email: "someone@test.com",
		},
	}

	querierParams := db.GetMyOrdersParams{
		UserID: 9919,
		Limit:  100,
		Offset: 2,
	}
	wrapperParams := entity.GetMyOrdersParams{
		UserID: 9919,
		Limit:  100,
		Offset: 2,
	}

	s.Run("get my orders got querier error", func() {
		s.querierRepo.EXPECT().GetMyOrders(ctx, querierParams).
			Return(nil, errors.New("querier error")).Times(1)

		result, err := wrapper.GetMyOrders(ctx, wrapperParams)
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().Equal(errorx.CodeInternal, goxErr.Code)
		s.Assert().EqualError(goxErr, "internal server error")
		s.Assert().Contains(goxErr.LogError(), "[common.internal] internal server error: querier error")
	})

	s.Run("get my orders successful", func() {
		s.querierRepo.EXPECT().GetMyOrders(ctx, querierParams).
			Return(rowsFromDB, nil).Times(1)

		result, err := wrapper.GetMyOrders(ctx, wrapperParams)
		s.Assert().Equal(expectedOrders, result)
		s.Assert().Nil(err)
	})
}
