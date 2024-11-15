package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/raymondwongso/gogox/errorx"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
	"github.com/swallowstalker/online-book-store/modules/bookstore/repository/db"
)

type DbWrapperRepo struct {
	db db.Querier
}

func NewDbWrapperRepo(db db.Querier) *DbWrapperRepo {
	return &DbWrapperRepo{
		db: db,
	}
}

func (w *DbWrapperRepo) CreateUser(ctx context.Context, email string) (*entity.User, error) {
	result, err := w.db.CreateUser(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			existingUser, err := w.FindUser(ctx, email)
			if err != nil {
				return nil, errorx.WrapWithLog(err, errorx.CodeInternal, "internal server error",
					"insert returns no row but existing user not found")
			}

			return existingUser, nil
		}
		return nil, errorx.Wrap(err, errorx.CodeInternal, "internal server error")
	}

	return result.ToEntity(), nil
}

func (w *DbWrapperRepo) FindUser(ctx context.Context, email string) (*entity.User, error) {
	result, err := w.db.FindUser(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.Wrap(err, errorx.CodeNotFound, "user not found")
		}
		return nil, errorx.Wrap(err, errorx.CodeInternal, "internal server error")
	}

	return result.ToEntity(), nil
}

func (w *DbWrapperRepo) GetBooks(ctx context.Context, arg entity.GetBooksParams) ([]entity.Book, error) {
	result, err := w.db.GetBooks(ctx, db.GetBooksParams{
		Limit:  arg.Limit,
		Offset: arg.Offset,
	})
	if err != nil {
		return nil, errorx.Wrap(err, errorx.CodeInternal, "internal server error")
	}

	resp := []entity.Book{}
	for _, r := range result {
		b := r.ToEntity()
		resp = append(resp, *b)
	}

	return resp, nil
}

func (w *DbWrapperRepo) CreateOrder(ctx context.Context, arg entity.CreateOrderParams) (*entity.Order, error) {
	result, err := w.db.CreateOrder(ctx, db.CreateOrderParams{
		UserID: arg.UserID,
		BookID: arg.BookID,
		Amount: arg.Amount,
	})
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint \"fk_order_books\"") {
			return nil, errorx.Wrap(err, errorx.CodeNotFound, "book cannot be found")
		}
		return nil, errorx.Wrap(err, errorx.CodeInternal, "internal server error")
	}

	order, err := w.db.FindOrder(ctx, result.ID)
	if err != nil {
		return nil, errorx.Wrap(err, errorx.CodeInternal, "internal server error")
	}

	return order.ToEntity(), nil
}

func (w *DbWrapperRepo) GetMyOrders(ctx context.Context, arg entity.GetMyOrdersParams) ([]entity.Order, error) {
	result, err := w.db.GetMyOrders(ctx, db.GetMyOrdersParams{
		Limit:  arg.Limit,
		Offset: arg.Offset,
		UserID: arg.UserID,
	})
	if err != nil {
		return nil, errorx.Wrap(err, errorx.CodeInternal, "internal server error")
	}

	resp := []entity.Order{}
	for _, r := range result {
		b := entity.Order{
			ID:        r.ID,
			UserID:    r.UserID,
			Email:     r.Email,
			BookID:    r.BookID,
			BookName:  r.BookName,
			Amount:    r.Amount,
			CreatedAt: r.CreatedAt.Time,
		}
		resp = append(resp, b)
	}

	return resp, nil
}
