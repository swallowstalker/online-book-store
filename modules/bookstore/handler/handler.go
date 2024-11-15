package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/raymondwongso/gogox/errorx"

	"github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

const (
	DefaultLimit  = 10
	DefaultOffset = 0
)

var HTTPErrorCodeMapping = map[string]int{
	errorx.CodeInvalidParameter: http.StatusBadRequest,
	errorx.CodeUnauthorized:     http.StatusUnauthorized,
	errorx.CodeNotFound:         http.StatusNotFound,
}

type UserService interface {
	CreateUser(ctx context.Context, params entity.CreateUserParam) (*entity.User, error)
}

type BookService interface {
	GetBooks(ctx context.Context, params entity.GetBooksParams) ([]entity.Book, error)
}

type OrderService interface {
	GetOrders(ctx context.Context, params entity.GetMyOrdersParams) ([]entity.Order, error)
	CreateOrder(ctx context.Context, params entity.CreateOrderParams) (*entity.Order, error)
}

type RestHandler struct {
	userService  UserService
	bookService  BookService
	orderService OrderService
}

func NewHandler(userService UserService, bookService BookService, orderService OrderService) *RestHandler {
	return &RestHandler{
		userService:  userService,
		bookService:  bookService,
		orderService: orderService,
	}
}

func (h *RestHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params entity.CreateUserParam
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		err = errorx.Wrap(err, errorx.CodeInvalidParameter, "Input is invalid")
		handleError(err, w)
		return
	}

	params.Email = strings.TrimSpace(params.Email)

	ctx := r.Context()
	user, err := h.userService.CreateUser(ctx, params)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(entity.CreateUserParam{Email: user.Email})
}

func (h *RestHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	limit, offset, err := parseLimitOffset(r)
	if err != nil {
		handleError(err, w)
		return
	}

	params := entity.GetBooksParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	ctx := r.Context()
	books, err := h.bookService.GetBooks(ctx, params)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(books)
}

func (h *RestHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params entity.CreateOrderParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		err = errorx.Wrap(err, errorx.CodeInvalidParameter, "Input is invalid")
		handleError(err, w)
		return
	}

	ctx := r.Context()
	params.UserID, err = getUserIDFromContext(ctx)
	if err != nil {
		handleError(err, w)
		return
	}

	order, err := h.orderService.CreateOrder(ctx, params)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(order)
}

func (h *RestHandler) GetMyOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	limit, offset, err := parseLimitOffset(r)
	if err != nil {
		handleError(err, w)
		return
	}

	params := entity.GetMyOrdersParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	ctx := r.Context()
	params.UserID, err = getUserIDFromContext(ctx)
	if err != nil {
		handleError(err, w)
		return
	}

	books, err := h.orderService.GetOrders(ctx, params)
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(books)
}

func handleError(err error, w http.ResponseWriter) {
	errx := errorx.ParseAndWrap(err, "server error")

	message := errx.Error()
	status, exist := HTTPErrorCodeMapping[errx.Code]
	if !exist {
		fmt.Println(errx.LogError())
		status = http.StatusInternalServerError
		message = "Internal server error"
	}

	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(entity.ErrorHandleResponse{Message: message})
}

func parseLimitOffset(r *http.Request) (limit, offset int, err error) {
	limitRaw := r.URL.Query().Get("limit")
	offsetRaw := r.URL.Query().Get("offset")
	if strings.TrimSpace(limitRaw) == "" {
		limit = DefaultLimit
	} else {
		limit, err = strconv.Atoi(limitRaw)
		if err != nil {
			return 0, 0, errorx.ErrInvalidParameter("limit invalid")
		}
	}

	if strings.TrimSpace(offsetRaw) == "" {
		offset = DefaultOffset
	} else {
		offset, err = strconv.Atoi(offsetRaw)
		if err != nil {
			return 0, 0, errorx.ErrInvalidParameter("offset invalid")
		}
	}

	return limit, offset, nil
}

func getUserIDFromContext(ctx context.Context) (int64, error) {
	userID, ok := ctx.Value(entity.UserContextKey{}).(int64)
	if !ok || userID == 0 {
		return 0, errorx.ErrUnauthorized("Unauthorized")
	}

	return userID, nil
}
