// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/pkg/agent/catalog (interfaces: Catalog)

// Package mock_catalog is a generated GoMock package.
package mock_catalog

import (
	gomock "github.com/golang/mock/gomock"
	catalog "github.com/spiffe/spire/pkg/agent/catalog"
	reflect "reflect"
)

// MockCatalog is a mock of Catalog interface
type MockCatalog struct {
	ctrl     *gomock.Controller
	recorder *MockCatalogMockRecorder
}

// MockCatalogMockRecorder is the mock recorder for MockCatalog
type MockCatalogMockRecorder struct {
	mock *MockCatalog
}

// NewMockCatalog creates a new mock instance
func NewMockCatalog(ctrl *gomock.Controller) *MockCatalog {
	mock := &MockCatalog{ctrl: ctrl}
	mock.recorder = &MockCatalogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCatalog) EXPECT() *MockCatalogMockRecorder {
	return m.recorder
}

// KeyManagers mocks base method
func (m *MockCatalog) KeyManagers() []*catalog.ManagedKeyManager {
	ret := m.ctrl.Call(m, "KeyManagers")
	ret0, _ := ret[0].([]*catalog.ManagedKeyManager)
	return ret0
}

// KeyManagers indicates an expected call of KeyManagers
func (mr *MockCatalogMockRecorder) KeyManagers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeyManagers", reflect.TypeOf((*MockCatalog)(nil).KeyManagers))
}

// NodeAttestors mocks base method
func (m *MockCatalog) NodeAttestors() []*catalog.ManagedNodeAttestor {
	ret := m.ctrl.Call(m, "NodeAttestors")
	ret0, _ := ret[0].([]*catalog.ManagedNodeAttestor)
	return ret0
}

// NodeAttestors indicates an expected call of NodeAttestors
func (mr *MockCatalogMockRecorder) NodeAttestors() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeAttestors", reflect.TypeOf((*MockCatalog)(nil).NodeAttestors))
}

// WorkloadAttestors mocks base method
func (m *MockCatalog) WorkloadAttestors() []*catalog.ManagedWorkloadAttestor {
	ret := m.ctrl.Call(m, "WorkloadAttestors")
	ret0, _ := ret[0].([]*catalog.ManagedWorkloadAttestor)
	return ret0
}

// WorkloadAttestors indicates an expected call of WorkloadAttestors
func (mr *MockCatalogMockRecorder) WorkloadAttestors() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkloadAttestors", reflect.TypeOf((*MockCatalog)(nil).WorkloadAttestors))
}
