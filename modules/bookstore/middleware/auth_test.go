package middleware_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/raymondwongso/gogox/errorx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
	"github.com/swallowstalker/online-book-store/modules/bookstore/middleware"
	mock_middleware "github.com/swallowstalker/online-book-store/test/mock/modules/bookstore/middleware"
)

type MiddlewareTestSuite struct {
	suite.Suite

	userRepo *mock_middleware.MockUserEmailChecker
}

func (s *MiddlewareTestSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.userRepo = mock_middleware.NewMockUserEmailChecker(ctrl)
}

func TestWrapperRepo(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}

func (s *MiddlewareTestSuite) TestCreateUser() {
	middleware := middleware.NewAuthMiddleware(s.userRepo)
	expectedUser := &entity.User{
		ID:    123,
		Email: "someone@test.com",
	}

	s.Run("auth header empty", func() {
		r := httptest.NewRequest(http.MethodGet, "http://localhost/test-middleware", nil)
		r.Header.Set("Authorization", "")
		w := httptest.NewRecorder()

		router := httprouter.New()

		handlerFunc := func(_ http.ResponseWriter, _ *http.Request) {}

		router.HandlerFunc(http.MethodGet, "/test-middleware", middleware.CheckEmailMiddleware(handlerFunc))
		router.ServeHTTP(w, r)
		resp := w.Result()

		assert.Equal(s.T(), http.StatusUnauthorized, resp.StatusCode)
		rawRespBody, err := io.ReadAll(resp.Body)
		require.NoError(s.T(), err)
		expected, err := json.Marshal(map[string]string{"message": "Unauthorized"})
		require.NoError(s.T(), err)

		assert.JSONEq(s.T(), string(expected), string(rawRespBody))

		resp.Body.Close()
	})

	s.Run("auth header is not email", func() {
		r := httptest.NewRequest(http.MethodGet, "http://localhost/test-middleware", nil)
		r.Header.Set("Authorization", "ahoy ahoy")
		w := httptest.NewRecorder()

		router := httprouter.New()

		handlerFunc := func(_ http.ResponseWriter, _ *http.Request) {}

		router.HandlerFunc(http.MethodGet, "/test-middleware", middleware.CheckEmailMiddleware(handlerFunc))
		router.ServeHTTP(w, r)
		resp := w.Result()

		assert.Equal(s.T(), http.StatusUnauthorized, resp.StatusCode)
		rawRespBody, err := io.ReadAll(resp.Body)
		require.NoError(s.T(), err)
		expected, err := json.Marshal(map[string]string{"message": "Unauthorized"})
		require.NoError(s.T(), err)

		assert.JSONEq(s.T(), string(expected), string(rawRespBody))
	})

	s.Run("no user for given email", func() {
		r := httptest.NewRequest(http.MethodGet, "http://localhost/test-middleware", nil)
		r.Header.Set("Authorization", "someone@test.com")
		w := httptest.NewRecorder()

		s.userRepo.EXPECT().FindUser(context.Background(), "someone@test.com").
			Return(nil, errorx.Wrap(sql.ErrNoRows, errorx.CodeNotFound, "user not found")).Times(1)

		router := httprouter.New()

		handlerFunc := func(_ http.ResponseWriter, _ *http.Request) {}

		router.HandlerFunc(http.MethodGet, "/test-middleware", middleware.CheckEmailMiddleware(handlerFunc))
		router.ServeHTTP(w, r)
		resp := w.Result()

		assert.Equal(s.T(), http.StatusUnauthorized, resp.StatusCode)
		rawRespBody, err := io.ReadAll(resp.Body)
		require.NoError(s.T(), err)
		expected, err := json.Marshal(map[string]string{"message": "Unauthorized"})
		require.NoError(s.T(), err)

		assert.JSONEq(s.T(), string(expected), string(rawRespBody))
	})

	s.Run("user repo error", func() {
		r := httptest.NewRequest(http.MethodGet, "http://localhost/test-middleware", nil)
		r.Header.Set("Authorization", "someone@test.com")
		w := httptest.NewRecorder()

		s.userRepo.EXPECT().FindUser(context.Background(), "someone@test.com").
			Return(nil, errors.New("something happened")).Times(1)

		router := httprouter.New()

		handlerFunc := func(_ http.ResponseWriter, _ *http.Request) {}

		router.HandlerFunc(http.MethodGet, "/test-middleware", middleware.CheckEmailMiddleware(handlerFunc))
		router.ServeHTTP(w, r)
		resp := w.Result()

		assert.Equal(s.T(), http.StatusInternalServerError, resp.StatusCode)
		rawRespBody, err := io.ReadAll(resp.Body)
		require.NoError(s.T(), err)
		expected, err := json.Marshal(map[string]string{"message": "Internal server error"})
		require.NoError(s.T(), err)

		assert.JSONEq(s.T(), string(expected), string(rawRespBody))
	})

	s.Run("success", func() {
		r := httptest.NewRequest(http.MethodGet, "http://localhost/test-middleware", nil)
		r.Header.Set("Authorization", "someone@test.com")
		w := httptest.NewRecorder()

		s.userRepo.EXPECT().FindUser(context.Background(), "someone@test.com").
			Return(expectedUser, nil).Times(1)

		router := httprouter.New()

		handlerFunc := func(w http.ResponseWriter, r *http.Request) {
			userID, ok := r.Context().Value(entity.UserContextKey{}).(int64)
			require.True(s.T(), ok)
			assert.Equal(s.T(), int64(123), userID)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"status": "ok"}`))
			require.NoError(s.T(), err)
		}

		router.HandlerFunc(http.MethodGet, "/test-middleware", middleware.CheckEmailMiddleware(handlerFunc))
		router.ServeHTTP(w, r)
		resp := w.Result()

		assert.Equal(s.T(), http.StatusOK, resp.StatusCode)
		rawRespBody, err := io.ReadAll(resp.Body)
		require.NoError(s.T(), err)
		expected, err := json.Marshal(map[string]string{"status": "ok"})
		require.NoError(s.T(), err)

		assert.JSONEq(s.T(), string(expected), string(rawRespBody))
	})
}
