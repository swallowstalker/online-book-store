// Code generated by MockGen. DO NOT EDIT.
// Source: ./modules/bookstore/repository/db/querier.go

// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgtype "github.com/jackc/pgx/v5/pgtype"
	db "github.com/swallowstalker/online-book-store/modules/bookstore/repository/db"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockQuerier) CreateOrder(ctx context.Context, arg db.CreateOrderParams) (*db.CreateOrderRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, arg)
	ret0, _ := ret[0].(*db.CreateOrderRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockQuerierMockRecorder) CreateOrder(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockQuerier)(nil).CreateOrder), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockQuerier) CreateUser(ctx context.Context, email string) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, email)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockQuerierMockRecorder) CreateUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockQuerier)(nil).CreateUser), ctx, email)
}

// FindBook mocks base method.
func (m *MockQuerier) FindBook(ctx context.Context, id int64) (*db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBook", ctx, id)
	ret0, _ := ret[0].(*db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBook indicates an expected call of FindBook.
func (mr *MockQuerierMockRecorder) FindBook(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBook", reflect.TypeOf((*MockQuerier)(nil).FindBook), ctx, id)
}

// FindOrder mocks base method.
func (m *MockQuerier) FindOrder(ctx context.Context, id int64) (*db.FindOrderRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrder", ctx, id)
	ret0, _ := ret[0].(*db.FindOrderRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrder indicates an expected call of FindOrder.
func (mr *MockQuerierMockRecorder) FindOrder(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrder", reflect.TypeOf((*MockQuerier)(nil).FindOrder), ctx, id)
}

// FindUser mocks base method.
func (m *MockQuerier) FindUser(ctx context.Context, email string) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", ctx, email)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of FindUser.
func (mr *MockQuerierMockRecorder) FindUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockQuerier)(nil).FindUser), ctx, email)
}

// FindUserByToken mocks base method.
func (m *MockQuerier) FindUserByToken(ctx context.Context, token pgtype.Text) (*db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByToken", ctx, token)
	ret0, _ := ret[0].(*db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByToken indicates an expected call of FindUserByToken.
func (mr *MockQuerierMockRecorder) FindUserByToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByToken", reflect.TypeOf((*MockQuerier)(nil).FindUserByToken), ctx, token)
}

// GetBooks mocks base method.
func (m *MockQuerier) GetBooks(ctx context.Context, arg db.GetBooksParams) ([]*db.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooks", ctx, arg)
	ret0, _ := ret[0].([]*db.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooks indicates an expected call of GetBooks.
func (mr *MockQuerierMockRecorder) GetBooks(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooks", reflect.TypeOf((*MockQuerier)(nil).GetBooks), ctx, arg)
}

// GetMyOrders mocks base method.
func (m *MockQuerier) GetMyOrders(ctx context.Context, arg db.GetMyOrdersParams) ([]*db.GetMyOrdersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyOrders", ctx, arg)
	ret0, _ := ret[0].([]*db.GetMyOrdersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyOrders indicates an expected call of GetMyOrders.
func (mr *MockQuerierMockRecorder) GetMyOrders(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyOrders", reflect.TypeOf((*MockQuerier)(nil).GetMyOrders), ctx, arg)
}
