// Code generated by MockGen. DO NOT EDIT.
// Source: ./modules/bookstore/middleware/auth.go

// Package mock_middleware is a generated GoMock package.
package mock_middleware

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/swallowstalker/online-book-store/modules/bookstore/entity"
)

// MockTokenCheckerRepo is a mock of TokenCheckerRepo interface.
type MockTokenCheckerRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTokenCheckerRepoMockRecorder
}

// MockTokenCheckerRepoMockRecorder is the mock recorder for MockTokenCheckerRepo.
type MockTokenCheckerRepoMockRecorder struct {
	mock *MockTokenCheckerRepo
}

// NewMockTokenCheckerRepo creates a new mock instance.
func NewMockTokenCheckerRepo(ctrl *gomock.Controller) *MockTokenCheckerRepo {
	mock := &MockTokenCheckerRepo{ctrl: ctrl}
	mock.recorder = &MockTokenCheckerRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenCheckerRepo) EXPECT() *MockTokenCheckerRepoMockRecorder {
	return m.recorder
}

// FindUserByToken mocks base method.
func (m *MockTokenCheckerRepo) FindUserByToken(ctx context.Context, token string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByToken", ctx, token)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByToken indicates an expected call of FindUserByToken.
func (mr *MockTokenCheckerRepoMockRecorder) FindUserByToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByToken", reflect.TypeOf((*MockTokenCheckerRepo)(nil).FindUserByToken), ctx, token)
}
