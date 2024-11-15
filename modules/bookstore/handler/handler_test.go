package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
	"github.com/swallowstalker/online-book-store/modules/bookstore/handler"
	mock_handler "github.com/swallowstalker/online-book-store/test/mock/modules/bookstore/handler"
)

type HandlerTestSuite struct {
	suite.Suite

	userSvc  *mock_handler.MockUserService
	orderSvc *mock_handler.MockOrderService
	bookSvc  *mock_handler.MockBookService
}

func (s *HandlerTestSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())

	s.userSvc = mock_handler.NewMockUserService(ctrl)
	s.orderSvc = mock_handler.NewMockOrderService(ctrl)
	s.bookSvc = mock_handler.NewMockBookService(ctrl)
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (s *HandlerTestSuite) TestCreateUser() {
	s.Run("error while decoding json request body", func() {
		ctx := context.Background()
		requestBody := `{"email":"someone@test.com"`

		r := httptest.NewRequestWithContext(ctx, http.MethodPost, "http://localhost/users", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.CreateUser(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusBadRequest, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "Input is invalid"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("service error", func() {
		ctx := context.Background()
		requestBody := `{"email":"someone@test.com"}`

		s.userSvc.EXPECT().CreateUser(ctx, entity.CreateUserParam{Email: "someone@test.com"}).
			Return(nil, errors.New("service error")).Times(1)

		r := httptest.NewRequestWithContext(ctx, http.MethodPost, "http://localhost/users", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.CreateUser(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusInternalServerError, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "Internal server error"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("successful", func() {
		ctx := context.Background()
		requestBody := `{"email":"  someone@test.com  "}`

		expectedUser := entity.User{
			ID:    1,
			Email: "someone@test.com",
		}

		params := entity.CreateUserParam{Email: "someone@test.com"}

		s.userSvc.EXPECT().CreateUser(ctx, params).
			Return(&expectedUser, nil).Times(1)

		r := httptest.NewRequestWithContext(ctx, http.MethodPost, "http://localhost/users", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.CreateUser(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusCreated, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(params)
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})
}

func (s *HandlerTestSuite) TestGetBooks() {
	s.Run("invalid limit", func() {
		ctx := context.Background()

		r := httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/books?limit=somenumbers", nil)
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.GetBooks(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusBadRequest, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "limit invalid"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("invalid offset", func() {
		ctx := context.Background()

		r := httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/books?limit=10&offset=somenumbers", nil)
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.GetBooks(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusBadRequest, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "offset invalid"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("service error", func() {
		ctx := context.Background()

		r := httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/books", nil)
		w := httptest.NewRecorder()

		params := entity.GetBooksParams{
			Limit:  10,
			Offset: 0,
		}

		s.bookSvc.EXPECT().GetBooks(ctx, params).Return(nil, errors.New("service error")).Times(1)

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.GetBooks(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusInternalServerError, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "Internal server error"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("successful", func() {
		ctx := context.Background()

		r := httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/books", nil)
		w := httptest.NewRecorder()

		params := entity.GetBooksParams{
			Limit:  10,
			Offset: 0,
		}

		expectedBooks := []entity.Book{
			{
				ID:   1,
				Name: "Book A",
			},
			{
				ID:   2,
				Name: "Book B",
			},
			{
				ID:   3,
				Name: "Book C",
			},
		}

		s.bookSvc.EXPECT().GetBooks(ctx, params).
			Return(expectedBooks, nil).Times(1)

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.GetBooks(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusOK, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(expectedBooks)
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})
}

func (s *HandlerTestSuite) TestCreateOrder() {
	now := time.Now()

	s.Run("error while decoding json request body", func() {
		ctx := context.Background()
		requestBody := `{"user_id":0,"book_id":0,"amount":0`

		r := httptest.NewRequestWithContext(ctx, http.MethodPost, "http://localhost/orders", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.CreateOrder(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusBadRequest, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "Input is invalid"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("context has no user id", func() {
		ctx := context.WithValue(context.Background(), entity.UserContextKey{}, 0)
		requestBody := `{"book_id":99,"amount":10}`

		r := httptest.NewRequestWithContext(ctx, http.MethodPost, "http://localhost/orders", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.CreateOrder(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusUnauthorized, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "Unauthorized"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("service error", func() {
		ctx := context.WithValue(context.Background(), entity.UserContextKey{}, int64(123))
		requestBody := `{"book_id":99,"amount":10}`
		params := entity.CreateOrderParams{
			UserID: 123,
			BookID: 99,
			Amount: 10,
		}

		s.orderSvc.EXPECT().CreateOrder(ctx, params).
			Return(nil, errors.New("service error")).Times(1)

		r := httptest.NewRequestWithContext(ctx, http.MethodPost, "http://localhost/orders", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.CreateOrder(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusInternalServerError, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "Internal server error"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("successful", func() {
		ctx := context.WithValue(context.Background(), entity.UserContextKey{}, int64(123))
		requestBody := `{"book_id":99,"amount":10}`

		expectedOrder := entity.Order{
			ID:        1,
			UserID:    123,
			Email:     "",
			BookID:    99,
			BookName:  "",
			Amount:    10,
			CreatedAt: now,
		}
		params := entity.CreateOrderParams{
			UserID: 123,
			BookID: 99,
			Amount: 10,
		}

		s.orderSvc.EXPECT().CreateOrder(ctx, params).
			Return(&expectedOrder, nil).Times(1)

		r := httptest.NewRequestWithContext(ctx, http.MethodPost, "http://localhost/orders", strings.NewReader(requestBody))
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.CreateOrder(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusCreated, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(expectedOrder)
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})
}

func (s *HandlerTestSuite) TestGetMyOrders() {
	now := time.Now()

	s.Run("invalid limit", func() {
		ctx := context.WithValue(context.Background(), entity.UserContextKey{}, int64(99))

		r := httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/orders?limit=somenumbers", nil)
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.GetMyOrders(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusBadRequest, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "limit invalid"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("invalid offset", func() {
		ctx := context.WithValue(context.Background(), entity.UserContextKey{}, int64(99))

		r := httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/orders?limit=10&offset=somenumbers", nil)
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.GetMyOrders(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusBadRequest, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "offset invalid"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("context has no user id", func() {
		ctx := context.WithValue(context.Background(), entity.UserContextKey{}, int64(0))

		r := httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/orders", nil)
		w := httptest.NewRecorder()

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.GetMyOrders(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusUnauthorized, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "Unauthorized"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("service error", func() {
		ctx := context.WithValue(context.Background(), entity.UserContextKey{}, int64(99))

		r := httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/orders", nil)
		w := httptest.NewRecorder()

		params := entity.GetMyOrdersParams{
			Limit:  10,
			Offset: 0,
			UserID: 99,
		}

		s.orderSvc.EXPECT().GetOrders(ctx, params).
			Return(nil, errors.New("service error")).Times(1)

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.GetMyOrders(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusInternalServerError, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(entity.ErrorHandleResponse{Message: "Internal server error"})
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})

	s.Run("successful", func() {
		ctx := context.WithValue(context.Background(), entity.UserContextKey{}, int64(99))

		r := httptest.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/orders", nil)
		w := httptest.NewRecorder()

		params := entity.GetMyOrdersParams{
			Limit:  10,
			Offset: 0,
			UserID: 99,
		}

		expectedBooks := []entity.Order{
			{
				ID:        1,
				UserID:    99,
				Email:     "someone@test.com",
				BookID:    1,
				BookName:  "Book A",
				Amount:    20,
				CreatedAt: now,
			},
			{
				ID:        2,
				UserID:    99,
				Email:     "someone@test.com",
				BookID:    2,
				BookName:  "Book B",
				Amount:    10,
				CreatedAt: now,
			},
		}

		s.orderSvc.EXPECT().GetOrders(ctx, params).
			Return(expectedBooks, nil).Times(1)

		h := handler.NewHandler(s.userSvc, s.bookSvc, s.orderSvc)
		h.GetMyOrders(w, r)
		resp := w.Result()

		s.Assert().Equal(http.StatusOK, resp.StatusCode)

		rawRespBody, err := io.ReadAll(resp.Body)
		s.Require().NoError(err)
		expected, err := json.Marshal(expectedBooks)
		s.Require().NoError(err)

		s.JSONEq(string(expected), string(rawRespBody))
	})
}
