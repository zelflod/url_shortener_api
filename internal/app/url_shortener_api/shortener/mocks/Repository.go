// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_shortener is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	models "url_shortener_api/internal/app/url_shortener_api/models"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// InsertUrl mocks base method
func (m *MockRepository) InsertUrl(ctx context.Context, url string, ttlSeconds int) (*models.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUrl", ctx, url, ttlSeconds)
	ret0, _ := ret[0].(*models.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUrl indicates an expected call of InsertUrl
func (mr *MockRepositoryMockRecorder) InsertUrl(ctx, url, ttlSeconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUrl", reflect.TypeOf((*MockRepository)(nil).InsertUrl), ctx, url, ttlSeconds)
}

// GetByUrl mocks base method
func (m *MockRepository) GetByUrl(ctx context.Context, url string) (*models.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUrl", ctx, url)
	ret0, _ := ret[0].(*models.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUrl indicates an expected call of GetByUrl
func (mr *MockRepositoryMockRecorder) GetByUrl(ctx, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUrl", reflect.TypeOf((*MockRepository)(nil).GetByUrl), ctx, url)
}

// UpdateLink mocks base method
func (m *MockRepository) UpdateLink(ctx context.Context, link *models.Link) (*models.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLink", ctx, link)
	ret0, _ := ret[0].(*models.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLink indicates an expected call of UpdateLink
func (mr *MockRepositoryMockRecorder) UpdateLink(ctx, link interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLink", reflect.TypeOf((*MockRepository)(nil).UpdateLink), ctx, link)
}

// GetById mocks base method
func (m *MockRepository) GetById(ctx context.Context, id int64, ttlSeconds int) (*models.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id, ttlSeconds)
	ret0, _ := ret[0].(*models.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockRepositoryMockRecorder) GetById(ctx, id, ttlSeconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRepository)(nil).GetById), ctx, id, ttlSeconds)
}

// GetAllLinks mocks base method
func (m *MockRepository) GetAllLinks(ctx context.Context) (*models.AllLinks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLinks", ctx)
	ret0, _ := ret[0].(*models.AllLinks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLinks indicates an expected call of GetAllLinks
func (mr *MockRepositoryMockRecorder) GetAllLinks(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLinks", reflect.TypeOf((*MockRepository)(nil).GetAllLinks), ctx)
}

// DeleteExpired mocks base method
func (m *MockRepository) DeleteExpired() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpired")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExpired indicates an expected call of DeleteExpired
func (mr *MockRepositoryMockRecorder) DeleteExpired() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpired", reflect.TypeOf((*MockRepository)(nil).DeleteExpired))
}
