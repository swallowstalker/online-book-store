package repository

import (
	"context"
	"database/sql"
	"errors"
	"slices"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/raymondwongso/gogox/errorx"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
	"github.com/swallowstalker/online-book-store/modules/bookstore/repository/db"
)

type DbWrapperRepo struct {
	db db.QuerierWithTx
}

func NewDbWrapperRepo(db db.QuerierWithTx) *DbWrapperRepo {
	return &DbWrapperRepo{
		db: db,
	}
}

func (w *DbWrapperRepo) WrapTx(tx pgx.Tx) db.QuerierWithTx {
	return w.db.WrapTx(tx)
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

func (w *DbWrapperRepo) CreateOrder(ctx context.Context, tx pgx.Tx, arg entity.CreateOrderParams) (*entity.Order, error) {
	result, err := w.db.WrapTx(tx).CreateOrder(ctx, arg.UserID)

	if err != nil {
		return nil, errorx.Wrap(err, errorx.CodeInternal, "internal server error")
	}

	return result.ToEntity(), nil
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
	orders := map[int64]entity.Order{}

	for _, r := range result {
		var order entity.Order
		var exist bool
		if order, exist = orders[r.OrderID]; !exist {
			order = entity.Order{
				ID:        r.OrderID,
				UserID:    r.UserID,
				Email:     r.Email,
				Items:     []entity.OrderItem{},
				CreatedAt: r.CreatedAt.Time,
			}
			orders[r.OrderID] = order
		}

		order.Items = append(order.Items, entity.OrderItem{
			ID:        r.ItemID,
			OrderID:   r.OrderID,
			BookID:    r.BookID,
			Amount:    r.Amount,
			CreatedAt: r.ItemCreatedAt.Time,
		})

		orders[r.OrderID] = order
	}

	for _, o := range orders {
		resp = append(resp, o)
	}

	slices.SortFunc(resp, func(a, b entity.Order) int {
		if a.CreatedAt.After(b.CreatedAt) {
			return -1
		}
		return 1
	})

	return resp, nil
}

func (w *DbWrapperRepo) FindBook(ctx context.Context, tx pgx.Tx, id int64) (*entity.Book, error) {
	result, err := w.db.WrapTx(tx).FindBook(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorx.Wrap(err, errorx.CodeNotFound, "book cannot be found")
		}
		return nil, errorx.Wrap(err, errorx.CodeInternal, "internal server error")
	}

	return result.ToEntity(), err
}

func (w *DbWrapperRepo) FindUserByToken(ctx context.Context, token string) (*entity.User, error) {
	result, err := w.db.FindUserByToken(ctx, pgtype.Text{
		String: token,
		Valid:  true,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errorx.Wrap(err, errorx.CodeInternal, "internal server error")
	}

	return result.ToEntity(), nil
}

func (w *DbWrapperRepo) CreateOrderItem(ctx context.Context, tx pgx.Tx, params entity.CreateOrderItemParams) (*entity.OrderItem, error) {
	result, err := w.db.WrapTx(tx).CreateOrderItem(ctx, db.CreateOrderItemParams{
		OrderID: params.OrderID,
		BookID:  params.BookID,
		Amount:  params.Amount,
	})
	if err != nil {
		return nil, errorx.Wrap(err, errorx.CodeInternal, "internal server error")
	}

	return result.ToEntity(), nil
}
