package middleware

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

type TokenCheckerRepo interface {
	FindUserByToken(ctx context.Context, token string) (*entity.User, error)
}

type Auth struct {
	userRepo TokenCheckerRepo
}

func NewAuthMiddleware(userRepo TokenCheckerRepo) *Auth {
	return &Auth{
		userRepo: userRepo,
	}
}

func (m *Auth) CheckTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		auth := r.Header.Get("Authorization")
		token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))
		if token == "" || strings.Contains(token, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(entity.ErrorHandleResponse{Message: "Unauthorized"})
			return
		}

		user, err := m.userRepo.FindUserByToken(r.Context(), token)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusUnauthorized)
				_ = json.NewEncoder(w).Encode(entity.ErrorHandleResponse{Message: "Unauthorized"})
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(entity.ErrorHandleResponse{Message: "Internal server error"})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, entity.UserContextKey{}, user.ID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
