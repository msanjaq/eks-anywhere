// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/eks-anywhere/pkg/govmomi (interfaces: VSphereClient,VMOMIAuthorizationManager,VMOMIFinder,VMOMISessionBuilder,VMOMIFinderBuilder,VMOMIAuthorizationManagerBuilder)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	url "net/url"
	reflect "reflect"

	govmomi "github.com/aws/eks-anywhere/pkg/govmomi"
	gomock "github.com/golang/mock/gomock"
	govmomi0 "github.com/vmware/govmomi"
	find "github.com/vmware/govmomi/find"
	object "github.com/vmware/govmomi/object"
	vim25 "github.com/vmware/govmomi/vim25"
	types "github.com/vmware/govmomi/vim25/types"
)

// MockVSphereClient is a mock of VSphereClient interface.
type MockVSphereClient struct {
	ctrl     *gomock.Controller
	recorder *MockVSphereClientMockRecorder
}

// MockVSphereClientMockRecorder is the mock recorder for MockVSphereClient.
type MockVSphereClientMockRecorder struct {
	mock *MockVSphereClient
}

// NewMockVSphereClient creates a new mock instance.
func NewMockVSphereClient(ctrl *gomock.Controller) *MockVSphereClient {
	mock := &MockVSphereClient{ctrl: ctrl}
	mock.recorder = &MockVSphereClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVSphereClient) EXPECT() *MockVSphereClientMockRecorder {
	return m.recorder
}

// GetPrivsOnEntity mocks base method.
func (m *MockVSphereClient) GetPrivsOnEntity(arg0 context.Context, arg1, arg2, arg3 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivsOnEntity", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivsOnEntity indicates an expected call of GetPrivsOnEntity.
func (mr *MockVSphereClientMockRecorder) GetPrivsOnEntity(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivsOnEntity", reflect.TypeOf((*MockVSphereClient)(nil).GetPrivsOnEntity), arg0, arg1, arg2, arg3)
}

// Username mocks base method.
func (m *MockVSphereClient) Username() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Username")
	ret0, _ := ret[0].(string)
	return ret0
}

// Username indicates an expected call of Username.
func (mr *MockVSphereClientMockRecorder) Username() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Username", reflect.TypeOf((*MockVSphereClient)(nil).Username))
}

// MockVMOMIAuthorizationManager is a mock of VMOMIAuthorizationManager interface.
type MockVMOMIAuthorizationManager struct {
	ctrl     *gomock.Controller
	recorder *MockVMOMIAuthorizationManagerMockRecorder
}

// MockVMOMIAuthorizationManagerMockRecorder is the mock recorder for MockVMOMIAuthorizationManager.
type MockVMOMIAuthorizationManagerMockRecorder struct {
	mock *MockVMOMIAuthorizationManager
}

// NewMockVMOMIAuthorizationManager creates a new mock instance.
func NewMockVMOMIAuthorizationManager(ctrl *gomock.Controller) *MockVMOMIAuthorizationManager {
	mock := &MockVMOMIAuthorizationManager{ctrl: ctrl}
	mock.recorder = &MockVMOMIAuthorizationManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVMOMIAuthorizationManager) EXPECT() *MockVMOMIAuthorizationManagerMockRecorder {
	return m.recorder
}

// FetchUserPrivilegeOnEntities mocks base method.
func (m *MockVMOMIAuthorizationManager) FetchUserPrivilegeOnEntities(arg0 context.Context, arg1 []types.ManagedObjectReference, arg2 string) ([]types.UserPrivilegeResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserPrivilegeOnEntities", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types.UserPrivilegeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUserPrivilegeOnEntities indicates an expected call of FetchUserPrivilegeOnEntities.
func (mr *MockVMOMIAuthorizationManagerMockRecorder) FetchUserPrivilegeOnEntities(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserPrivilegeOnEntities", reflect.TypeOf((*MockVMOMIAuthorizationManager)(nil).FetchUserPrivilegeOnEntities), arg0, arg1, arg2)
}

// MockVMOMIFinder is a mock of VMOMIFinder interface.
type MockVMOMIFinder struct {
	ctrl     *gomock.Controller
	recorder *MockVMOMIFinderMockRecorder
}

// MockVMOMIFinderMockRecorder is the mock recorder for MockVMOMIFinder.
type MockVMOMIFinderMockRecorder struct {
	mock *MockVMOMIFinder
}

// NewMockVMOMIFinder creates a new mock instance.
func NewMockVMOMIFinder(ctrl *gomock.Controller) *MockVMOMIFinder {
	mock := &MockVMOMIFinder{ctrl: ctrl}
	mock.recorder = &MockVMOMIFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVMOMIFinder) EXPECT() *MockVMOMIFinderMockRecorder {
	return m.recorder
}

// Datacenter mocks base method.
func (m *MockVMOMIFinder) Datacenter(arg0 context.Context, arg1 string) (*object.Datacenter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Datacenter", arg0, arg1)
	ret0, _ := ret[0].(*object.Datacenter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Datacenter indicates an expected call of Datacenter.
func (mr *MockVMOMIFinderMockRecorder) Datacenter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Datacenter", reflect.TypeOf((*MockVMOMIFinder)(nil).Datacenter), arg0, arg1)
}

// Datastore mocks base method.
func (m *MockVMOMIFinder) Datastore(arg0 context.Context, arg1 string) (*object.Datastore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Datastore", arg0, arg1)
	ret0, _ := ret[0].(*object.Datastore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Datastore indicates an expected call of Datastore.
func (mr *MockVMOMIFinderMockRecorder) Datastore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Datastore", reflect.TypeOf((*MockVMOMIFinder)(nil).Datastore), arg0, arg1)
}

// Folder mocks base method.
func (m *MockVMOMIFinder) Folder(arg0 context.Context, arg1 string) (*object.Folder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Folder", arg0, arg1)
	ret0, _ := ret[0].(*object.Folder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Folder indicates an expected call of Folder.
func (mr *MockVMOMIFinderMockRecorder) Folder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Folder", reflect.TypeOf((*MockVMOMIFinder)(nil).Folder), arg0, arg1)
}

// Network mocks base method.
func (m *MockVMOMIFinder) Network(arg0 context.Context, arg1 string) (object.NetworkReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Network", arg0, arg1)
	ret0, _ := ret[0].(object.NetworkReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Network indicates an expected call of Network.
func (mr *MockVMOMIFinderMockRecorder) Network(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*MockVMOMIFinder)(nil).Network), arg0, arg1)
}

// ResourcePool mocks base method.
func (m *MockVMOMIFinder) ResourcePool(arg0 context.Context, arg1 string) (*object.ResourcePool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourcePool", arg0, arg1)
	ret0, _ := ret[0].(*object.ResourcePool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResourcePool indicates an expected call of ResourcePool.
func (mr *MockVMOMIFinderMockRecorder) ResourcePool(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourcePool", reflect.TypeOf((*MockVMOMIFinder)(nil).ResourcePool), arg0, arg1)
}

// SetDatacenter mocks base method.
func (m *MockVMOMIFinder) SetDatacenter(arg0 *object.Datacenter) *find.Finder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDatacenter", arg0)
	ret0, _ := ret[0].(*find.Finder)
	return ret0
}

// SetDatacenter indicates an expected call of SetDatacenter.
func (mr *MockVMOMIFinderMockRecorder) SetDatacenter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDatacenter", reflect.TypeOf((*MockVMOMIFinder)(nil).SetDatacenter), arg0)
}

// VirtualMachine mocks base method.
func (m *MockVMOMIFinder) VirtualMachine(arg0 context.Context, arg1 string) (*object.VirtualMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VirtualMachine", arg0, arg1)
	ret0, _ := ret[0].(*object.VirtualMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VirtualMachine indicates an expected call of VirtualMachine.
func (mr *MockVMOMIFinderMockRecorder) VirtualMachine(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VirtualMachine", reflect.TypeOf((*MockVMOMIFinder)(nil).VirtualMachine), arg0, arg1)
}

// MockVMOMISessionBuilder is a mock of VMOMISessionBuilder interface.
type MockVMOMISessionBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockVMOMISessionBuilderMockRecorder
}

// MockVMOMISessionBuilderMockRecorder is the mock recorder for MockVMOMISessionBuilder.
type MockVMOMISessionBuilderMockRecorder struct {
	mock *MockVMOMISessionBuilder
}

// NewMockVMOMISessionBuilder creates a new mock instance.
func NewMockVMOMISessionBuilder(ctrl *gomock.Controller) *MockVMOMISessionBuilder {
	mock := &MockVMOMISessionBuilder{ctrl: ctrl}
	mock.recorder = &MockVMOMISessionBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVMOMISessionBuilder) EXPECT() *MockVMOMISessionBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockVMOMISessionBuilder) Build(arg0 context.Context, arg1 *url.URL, arg2 bool) (*govmomi0.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0, arg1, arg2)
	ret0, _ := ret[0].(*govmomi0.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Build indicates an expected call of Build.
func (mr *MockVMOMISessionBuilderMockRecorder) Build(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockVMOMISessionBuilder)(nil).Build), arg0, arg1, arg2)
}

// MockVMOMIFinderBuilder is a mock of VMOMIFinderBuilder interface.
type MockVMOMIFinderBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockVMOMIFinderBuilderMockRecorder
}

// MockVMOMIFinderBuilderMockRecorder is the mock recorder for MockVMOMIFinderBuilder.
type MockVMOMIFinderBuilderMockRecorder struct {
	mock *MockVMOMIFinderBuilder
}

// NewMockVMOMIFinderBuilder creates a new mock instance.
func NewMockVMOMIFinderBuilder(ctrl *gomock.Controller) *MockVMOMIFinderBuilder {
	mock := &MockVMOMIFinderBuilder{ctrl: ctrl}
	mock.recorder = &MockVMOMIFinderBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVMOMIFinderBuilder) EXPECT() *MockVMOMIFinderBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockVMOMIFinderBuilder) Build(arg0 *vim25.Client, arg1 ...bool) govmomi.VMOMIFinder {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Build", varargs...)
	ret0, _ := ret[0].(govmomi.VMOMIFinder)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockVMOMIFinderBuilderMockRecorder) Build(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockVMOMIFinderBuilder)(nil).Build), varargs...)
}

// MockVMOMIAuthorizationManagerBuilder is a mock of VMOMIAuthorizationManagerBuilder interface.
type MockVMOMIAuthorizationManagerBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockVMOMIAuthorizationManagerBuilderMockRecorder
}

// MockVMOMIAuthorizationManagerBuilderMockRecorder is the mock recorder for MockVMOMIAuthorizationManagerBuilder.
type MockVMOMIAuthorizationManagerBuilderMockRecorder struct {
	mock *MockVMOMIAuthorizationManagerBuilder
}

// NewMockVMOMIAuthorizationManagerBuilder creates a new mock instance.
func NewMockVMOMIAuthorizationManagerBuilder(ctrl *gomock.Controller) *MockVMOMIAuthorizationManagerBuilder {
	mock := &MockVMOMIAuthorizationManagerBuilder{ctrl: ctrl}
	mock.recorder = &MockVMOMIAuthorizationManagerBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVMOMIAuthorizationManagerBuilder) EXPECT() *MockVMOMIAuthorizationManagerBuilderMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockVMOMIAuthorizationManagerBuilder) Build(arg0 *vim25.Client) *object.AuthorizationManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0)
	ret0, _ := ret[0].(*object.AuthorizationManager)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockVMOMIAuthorizationManagerBuilderMockRecorder) Build(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockVMOMIAuthorizationManagerBuilder)(nil).Build), arg0)
}
