// Code generated by MockGen. DO NOT EDIT.
// Source: ./modules/bookstore/service/repository.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(ctx context.Context, email string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, email)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), ctx, email)
}

// MockBookRepository is a mock of BookRepository interface.
type MockBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryMockRecorder
}

// MockBookRepositoryMockRecorder is the mock recorder for MockBookRepository.
type MockBookRepositoryMockRecorder struct {
	mock *MockBookRepository
}

// NewMockBookRepository creates a new mock instance.
func NewMockBookRepository(ctrl *gomock.Controller) *MockBookRepository {
	mock := &MockBookRepository{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepository) EXPECT() *MockBookRepositoryMockRecorder {
	return m.recorder
}

// GetBooks mocks base method.
func (m *MockBookRepository) GetBooks(ctx context.Context, arg entity.GetBooksParams) ([]entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBooks", ctx, arg)
	ret0, _ := ret[0].([]entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBooks indicates an expected call of GetBooks.
func (mr *MockBookRepositoryMockRecorder) GetBooks(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBooks", reflect.TypeOf((*MockBookRepository)(nil).GetBooks), ctx, arg)
}

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderRepository) CreateOrder(ctx context.Context, arg entity.CreateOrderParams) (*entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, arg)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepositoryMockRecorder) CreateOrder(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepository)(nil).CreateOrder), ctx, arg)
}

// CreateOrderItem mocks base method.
func (m *MockOrderRepository) CreateOrderItem(ctx context.Context, params entity.CreateOrderItemParams) (*entity.OrderItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrderItem", ctx, params)
	ret0, _ := ret[0].(*entity.OrderItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrderItem indicates an expected call of CreateOrderItem.
func (mr *MockOrderRepositoryMockRecorder) CreateOrderItem(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrderItem", reflect.TypeOf((*MockOrderRepository)(nil).CreateOrderItem), ctx, params)
}

// FindBook mocks base method.
func (m *MockOrderRepository) FindBook(ctx context.Context, id int64) (*entity.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBook", ctx, id)
	ret0, _ := ret[0].(*entity.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBook indicates an expected call of FindBook.
func (mr *MockOrderRepositoryMockRecorder) FindBook(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBook", reflect.TypeOf((*MockOrderRepository)(nil).FindBook), ctx, id)
}

// GetMyOrders mocks base method.
func (m *MockOrderRepository) GetMyOrders(ctx context.Context, arg entity.GetMyOrdersParams) ([]entity.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMyOrders", ctx, arg)
	ret0, _ := ret[0].([]entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMyOrders indicates an expected call of GetMyOrders.
func (mr *MockOrderRepositoryMockRecorder) GetMyOrders(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMyOrders", reflect.TypeOf((*MockOrderRepository)(nil).GetMyOrders), ctx, arg)
}
