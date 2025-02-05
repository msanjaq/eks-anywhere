// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/providers/vsphere/reconciler/reconciler.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cluster "github.com/aws/eks-anywhere/pkg/cluster"
	controller "github.com/aws/eks-anywhere/pkg/controller"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockCNIReconciler is a mock of CNIReconciler interface.
type MockCNIReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockCNIReconcilerMockRecorder
}

// MockCNIReconcilerMockRecorder is the mock recorder for MockCNIReconciler.
type MockCNIReconcilerMockRecorder struct {
	mock *MockCNIReconciler
}

// NewMockCNIReconciler creates a new mock instance.
func NewMockCNIReconciler(ctrl *gomock.Controller) *MockCNIReconciler {
	mock := &MockCNIReconciler{ctrl: ctrl}
	mock.recorder = &MockCNIReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCNIReconciler) EXPECT() *MockCNIReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method.
func (m *MockCNIReconciler) Reconcile(ctx context.Context, logger logr.Logger, client client.Client, spec *cluster.Spec) (controller.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", ctx, logger, client, spec)
	ret0, _ := ret[0].(controller.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reconcile indicates an expected call of Reconcile.
func (mr *MockCNIReconcilerMockRecorder) Reconcile(ctx, logger, client, spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockCNIReconciler)(nil).Reconcile), ctx, logger, client, spec)
}

// MockRemoteClientRegistry is a mock of RemoteClientRegistry interface.
type MockRemoteClientRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteClientRegistryMockRecorder
}

// MockRemoteClientRegistryMockRecorder is the mock recorder for MockRemoteClientRegistry.
type MockRemoteClientRegistryMockRecorder struct {
	mock *MockRemoteClientRegistry
}

// NewMockRemoteClientRegistry creates a new mock instance.
func NewMockRemoteClientRegistry(ctrl *gomock.Controller) *MockRemoteClientRegistry {
	mock := &MockRemoteClientRegistry{ctrl: ctrl}
	mock.recorder = &MockRemoteClientRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteClientRegistry) EXPECT() *MockRemoteClientRegistryMockRecorder {
	return m.recorder
}

// GetClient mocks base method.
func (m *MockRemoteClientRegistry) GetClient(ctx context.Context, cluster client.ObjectKey) (client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient", ctx, cluster)
	ret0, _ := ret[0].(client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClient indicates an expected call of GetClient.
func (mr *MockRemoteClientRegistryMockRecorder) GetClient(ctx, cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockRemoteClientRegistry)(nil).GetClient), ctx, cluster)
}
