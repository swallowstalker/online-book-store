package repository_test

import (
	"context"
	"database/sql"
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

	querierParams := db.CreateOrderParams{
		UserID: 9919,
		BookID: 123,
		Amount: 10,
	}
	wrapperParams := entity.CreateOrderParams{
		UserID: 9919,
		BookID: 123,
		Amount: 10,
	}

	expectedOrder := &entity.Order{
		ID:        90,
		UserID:    9919,
		Email:     "someone@test.com",
		BookID:    123,
		BookName:  "Book A",
		Amount:    10,
		CreatedAt: now,
	}

	rowFromDB := &db.Order{
		ID:     90,
		UserID: 9919,
		BookID: 123,
		Amount: 10,
		CreatedAt: pgtype.Timestamptz{
			Time:  now,
			Valid: true,
		},
	}

	latestCreatedOrder := &db.FindOrderRow{
		ID:        rowFromDB.ID,
		UserID:    rowFromDB.UserID,
		BookID:    rowFromDB.BookID,
		Amount:    rowFromDB.Amount,
		CreatedAt: rowFromDB.CreatedAt,
		Email:     "someone@test.com",
		BookName:  "Book A",
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

	s.Run("create order but no book found", func() {
		s.querierRepo.EXPECT().CreateOrder(ctx, querierParams).
			Return(nil, errors.New("violates foreign key constraint \"fk_order_books\"")).Times(1)

		result, err := wrapper.CreateOrder(ctx, wrapperParams)
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().EqualError(goxErr, "book cannot be found")
		s.Assert().Contains(goxErr.LogError(), "[common.not_found] book cannot be found: violates foreign key constraint \"fk_order_books\"")
	})

	s.Run("create order error when getting last created order", func() {
		s.querierRepo.EXPECT().CreateOrder(ctx, querierParams).
			Return(rowFromDB, nil).Times(1)
		s.querierRepo.EXPECT().FindOrder(ctx, rowFromDB.ID).
			Return(nil, errors.New("find order error")).Times(1)

		result, err := wrapper.CreateOrder(ctx, wrapperParams)
		s.Assert().Nil(result)

		goxErr, ok := errorx.Parse(err)
		s.Require().True(ok)
		s.Assert().EqualError(goxErr, "internal server error")
		s.Assert().Contains(goxErr.LogError(), "[common.internal] internal server error: find order error")
	})

	s.Run("create order successful", func() {
		s.querierRepo.EXPECT().CreateOrder(ctx, querierParams).
			Return(rowFromDB, nil).Times(1)
		s.querierRepo.EXPECT().FindOrder(ctx, rowFromDB.ID).
			Return(latestCreatedOrder, nil).Times(1)

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
			ID:        123,
			UserID:    9919,
			Email:     "someone@test.com",
			BookID:    920,
			BookName:  "Book A",
			Amount:    10,
			CreatedAt: now,
		},
		{
			ID:        124,
			UserID:    9919,
			Email:     "someone@test.com",
			BookID:    921,
			BookName:  "Book B",
			Amount:    20,
			CreatedAt: now,
		},
	}

	rowsFromDB := []*db.GetMyOrdersRow{
		{
			ID:     123,
			UserID: 9919,
			BookID: 920,
			Amount: 10,
			CreatedAt: pgtype.Timestamptz{
				Time:  now,
				Valid: true,
			},
			Email:    "someone@test.com",
			BookName: "Book A",
		},
		{
			ID:     124,
			UserID: 9919,
			BookID: 921,
			Amount: 20,
			CreatedAt: pgtype.Timestamptz{
				Time:  now,
				Valid: true,
			},
			Email:    "someone@test.com",
			BookName: "Book B",
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
