// Code generated by MockGen. DO NOT EDIT.
// Source: ports.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/tweet/usecases/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCases is a mock of UseCases interface.
type MockUseCases struct {
	ctrl     *gomock.Controller
	recorder *MockUseCasesMockRecorder
}

// MockUseCasesMockRecorder is the mock recorder for MockUseCases.
type MockUseCasesMockRecorder struct {
	mock *MockUseCases
}

// NewMockUseCases creates a new mock instance.
func NewMockUseCases(ctrl *gomock.Controller) *MockUseCases {
	mock := &MockUseCases{ctrl: ctrl}
	mock.recorder = &MockUseCasesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCases) EXPECT() *MockUseCasesMockRecorder {
	return m.recorder
}

// CreateTweet mocks base method.
func (m *MockUseCases) CreateTweet(arg0 context.Context, arg1 *domain.Tweet) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTweet", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTweet indicates an expected call of CreateTweet.
func (mr *MockUseCasesMockRecorder) CreateTweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTweet", reflect.TypeOf((*MockUseCases)(nil).CreateTweet), arg0, arg1)
}

// GetTimeline mocks base method.
func (m *MockUseCases) GetTimeline(arg0 context.Context, arg1 string) ([]domain.Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeline", arg0, arg1)
	ret0, _ := ret[0].([]domain.Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimeline indicates an expected call of GetTimeline.
func (mr *MockUseCasesMockRecorder) GetTimeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeline", reflect.TypeOf((*MockUseCases)(nil).GetTimeline), arg0, arg1)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetTweetsByUserIDs mocks base method.
func (m *MockRepository) GetTweetsByUserIDs(arg0 context.Context, arg1 []string, arg2, arg3 int) ([]domain.Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTweetsByUserIDs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]domain.Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTweetsByUserIDs indicates an expected call of GetTweetsByUserIDs.
func (mr *MockRepositoryMockRecorder) GetTweetsByUserIDs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTweetsByUserIDs", reflect.TypeOf((*MockRepository)(nil).GetTweetsByUserIDs), arg0, arg1, arg2, arg3)
}

// InsertTweetIntoTimeline mocks base method.
func (m *MockRepository) InsertTweetIntoTimeline(arg0 context.Context, arg1 string, arg2 *domain.Tweet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTweetIntoTimeline", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTweetIntoTimeline indicates an expected call of InsertTweetIntoTimeline.
func (mr *MockRepositoryMockRecorder) InsertTweetIntoTimeline(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTweetIntoTimeline", reflect.TypeOf((*MockRepository)(nil).InsertTweetIntoTimeline), arg0, arg1, arg2)
}

// SaveTweet mocks base method.
func (m *MockRepository) SaveTweet(arg0 context.Context, arg1 *domain.Tweet) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTweet", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveTweet indicates an expected call of SaveTweet.
func (mr *MockRepositoryMockRecorder) SaveTweet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTweet", reflect.TypeOf((*MockRepository)(nil).SaveTweet), arg0, arg1)
}

// MockCache is a mock of Cache interface.
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache.
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance.
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCache) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockCacheMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCache)(nil).Close))
}

// GetTimeline mocks base method.
func (m *MockCache) GetTimeline(arg0 context.Context, arg1 string) ([]domain.Tweet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeline", arg0, arg1)
	ret0, _ := ret[0].([]domain.Tweet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimeline indicates an expected call of GetTimeline.
func (mr *MockCacheMockRecorder) GetTimeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeline", reflect.TypeOf((*MockCache)(nil).GetTimeline), arg0, arg1)
}

// InvalidateUserTimeline mocks base method.
func (m *MockCache) InvalidateUserTimeline(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvalidateUserTimeline", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateUserTimeline indicates an expected call of InvalidateUserTimeline.
func (mr *MockCacheMockRecorder) InvalidateUserTimeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateUserTimeline", reflect.TypeOf((*MockCache)(nil).InvalidateUserTimeline), arg0, arg1)
}

// PushTweetToTimeline mocks base method.
func (m *MockCache) PushTweetToTimeline(arg0 context.Context, arg1 string, arg2 *domain.Tweet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushTweetToTimeline", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushTweetToTimeline indicates an expected call of PushTweetToTimeline.
func (mr *MockCacheMockRecorder) PushTweetToTimeline(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushTweetToTimeline", reflect.TypeOf((*MockCache)(nil).PushTweetToTimeline), arg0, arg1, arg2)
}

// SetTimeline mocks base method.
func (m *MockCache) SetTimeline(arg0 context.Context, arg1 string, arg2 []domain.Tweet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTimeline", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTimeline indicates an expected call of SetTimeline.
func (mr *MockCacheMockRecorder) SetTimeline(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeline", reflect.TypeOf((*MockCache)(nil).SetTimeline), arg0, arg1, arg2)
}

// MockBroker is a mock of Broker interface.
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker.
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance.
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// PublishTweetCreated mocks base method.
func (m *MockBroker) PublishTweetCreated(arg0 context.Context, arg1 *domain.Tweet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishTweetCreated", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishTweetCreated indicates an expected call of PublishTweetCreated.
func (mr *MockBrokerMockRecorder) PublishTweetCreated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishTweetCreated", reflect.TypeOf((*MockBroker)(nil).PublishTweetCreated), arg0, arg1)
}
