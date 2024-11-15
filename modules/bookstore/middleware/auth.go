package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/swallowstalker/online-book-store/modules/bookstore/customerror"
	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

type UserEmailChecker interface {
	FindUser(ctx context.Context, email string) (*entity.User, error)
}

type Auth struct {
	userRepo  UserEmailChecker
	validator *validator.Validate
}

func NewAuthMiddleware(userRepo UserEmailChecker) *Auth {
	return &Auth{
		userRepo:  userRepo,
		validator: validator.New(),
	}
}

func (m *Auth) CheckEmailMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		auth := r.Header.Get("Authorization")
		email := strings.TrimSpace(auth)
		if err := m.validator.Var(email, "required,email"); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(entity.ErrorHandleResponse{Message: "Unauthorized"})
			return
		}

		ctx := r.Context()
		user, err := m.userRepo.FindUser(ctx, email)
		if err != nil {
			if customerror.IsErrNotFound(err) {
				w.WriteHeader(http.StatusUnauthorized)
				_ = json.NewEncoder(w).Encode(entity.ErrorHandleResponse{Message: "Unauthorized"})
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(entity.ErrorHandleResponse{Message: "Internal server error"})
			return
		}

		ctx = context.WithValue(ctx, entity.UserContextKey{}, user.ID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
