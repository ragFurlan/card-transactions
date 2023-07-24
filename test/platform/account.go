// Code generated by MockGen. DO NOT EDIT.
// Source: internal/platform/repositories/account.go

// Package repository is a generated GoMock package.
package repository

import (
	entity "card-transactions/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccounts is a mock of Accounts interface.
type MockAccounts struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsMockRecorder
}

// MockAccountsMockRecorder is the mock recorder for MockAccounts.
type MockAccountsMockRecorder struct {
	mock *MockAccounts
}

// NewMockAccounts creates a new mock instance.
func NewMockAccounts(ctrl *gomock.Controller) *MockAccounts {
	mock := &MockAccounts{ctrl: ctrl}
	mock.recorder = &MockAccountsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccounts) EXPECT() *MockAccountsMockRecorder {
	return m.recorder
}

// GetByDocumentNumber mocks base method.
func (m *MockAccounts) GetByDocumentNumber(documentNumber string) (entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDocumentNumber", documentNumber)
	ret0, _ := ret[0].(entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDocumentNumber indicates an expected call of GetByDocumentNumber.
func (mr *MockAccountsMockRecorder) GetByDocumentNumber(documentNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDocumentNumber", reflect.TypeOf((*MockAccounts)(nil).GetByDocumentNumber), documentNumber)
}

// GetByID mocks base method.
func (m *MockAccounts) GetByID(id string) (entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAccountsMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAccounts)(nil).GetByID), id)
}

// Save mocks base method.
func (m *MockAccounts) Save(account entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockAccountsMockRecorder) Save(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockAccounts)(nil).Save), account)
}