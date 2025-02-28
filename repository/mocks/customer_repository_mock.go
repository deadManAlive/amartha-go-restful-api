// Code generated by MockGen. DO NOT EDIT.
// Source: repository/customer_repository.go
//
// Generated by this command:
//
//	mockgen -source=repository/customer_repository.go -destination=repository/mocks/customer_repository_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/aronipurwanto/go-restful-api/model/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockCustomerRepository is a mock of CustomerRepository interface.
type MockCustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepositoryMockRecorder
	isgomock struct{}
}

// MockCustomerRepositoryMockRecorder is the mock recorder for MockCustomerRepository.
type MockCustomerRepositoryMockRecorder struct {
	mock *MockCustomerRepository
}

// NewMockCustomerRepository creates a new mock instance.
func NewMockCustomerRepository(ctrl *gomock.Controller) *MockCustomerRepository {
	mock := &MockCustomerRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerRepository) EXPECT() *MockCustomerRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCustomerRepository) Delete(ctx context.Context, customer domain.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, customer)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomerRepositoryMockRecorder) Delete(ctx, customer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomerRepository)(nil).Delete), ctx, customer)
}

// FindAll mocks base method.
func (m *MockCustomerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].([]domain.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCustomerRepositoryMockRecorder) FindAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCustomerRepository)(nil).FindAll), ctx)
}

// FindById mocks base method.
func (m *MockCustomerRepository) FindById(ctx context.Context, customerId uint64) (domain.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, customerId)
	ret0, _ := ret[0].(domain.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockCustomerRepositoryMockRecorder) FindById(ctx, customerId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockCustomerRepository)(nil).FindById), ctx, customerId)
}

// Save mocks base method.
func (m *MockCustomerRepository) Save(ctx context.Context, Customer domain.Customer) (domain.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, Customer)
	ret0, _ := ret[0].(domain.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockCustomerRepositoryMockRecorder) Save(ctx, Customer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCustomerRepository)(nil).Save), ctx, Customer)
}

// Update mocks base method.
func (m *MockCustomerRepository) Update(ctx context.Context, Customer domain.Customer) (domain.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, Customer)
	ret0, _ := ret[0].(domain.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCustomerRepositoryMockRecorder) Update(ctx, Customer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomerRepository)(nil).Update), ctx, Customer)
}
