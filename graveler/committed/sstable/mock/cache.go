// Code generated by MockGen. DO NOT EDIT.
// Source: graveler/committed/sstable/cache.go

// Package mock is a generated GoMock package.
package mock

import (
	sstable "github.com/cockroachdb/pebble/sstable"
	gomock "github.com/golang/mock/gomock"
	committed "github.com/treeverse/lakefs/graveler/committed"
	reflect "reflect"
)

// Mockcache is a mock of cache interface
type Mockcache struct {
	ctrl     *gomock.Controller
	recorder *MockcacheMockRecorder
}

// MockcacheMockRecorder is the mock recorder for Mockcache
type MockcacheMockRecorder struct {
	mock *Mockcache
}

// NewMockcache creates a new mock instance
func NewMockcache(ctrl *gomock.Controller) *Mockcache {
	mock := &Mockcache{ctrl: ctrl}
	mock.recorder = &MockcacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockcache) EXPECT() *MockcacheMockRecorder {
	return m.recorder
}

// GetOrOpen mocks base method
func (m *Mockcache) GetOrOpen(namespace string, id committed.ID) (*sstable.Reader, func() error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrOpen", namespace, id)
	ret0, _ := ret[0].(*sstable.Reader)
	ret1, _ := ret[1].(func() error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrOpen indicates an expected call of GetOrOpen
func (mr *MockcacheMockRecorder) GetOrOpen(namespace, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrOpen", reflect.TypeOf((*Mockcache)(nil).GetOrOpen), namespace, id)
}
