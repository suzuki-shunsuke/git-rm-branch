// Code generated by MockGen. DO NOT EDIT.
// Source: domain/interface.go

// Package usecase is a generated GoMock package.
package usecase

import (
	os "os"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockOSIntf is a mock of OSIntf interface
type MockOSIntf struct {
	ctrl     *gomock.Controller
	recorder *MockOSIntfMockRecorder
}

// MockOSIntfMockRecorder is the mock recorder for MockOSIntf
type MockOSIntfMockRecorder struct {
	mock *MockOSIntf
}

// NewMockOSIntf creates a new mock instance
func NewMockOSIntf(ctrl *gomock.Controller) *MockOSIntf {
	mock := &MockOSIntf{ctrl: ctrl}
	mock.recorder = &MockOSIntfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOSIntf) EXPECT() *MockOSIntfMockRecorder {
	return m.recorder
}

// Getwd mocks base method
func (m *MockOSIntf) Getwd() (string, error) {
	ret := m.ctrl.Call(m, "Getwd")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Getwd indicates an expected call of Getwd
func (mr *MockOSIntfMockRecorder) Getwd() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getwd", reflect.TypeOf((*MockOSIntf)(nil).Getwd))
}

// Stat mocks base method
func (m *MockOSIntf) Stat(arg0 string) (os.FileInfo, error) {
	ret := m.ctrl.Call(m, "Stat", arg0)
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat
func (mr *MockOSIntfMockRecorder) Stat(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockOSIntf)(nil).Stat), arg0)
}

// MockIOUtilIntf is a mock of IOUtilIntf interface
type MockIOUtilIntf struct {
	ctrl     *gomock.Controller
	recorder *MockIOUtilIntfMockRecorder
}

// MockIOUtilIntfMockRecorder is the mock recorder for MockIOUtilIntf
type MockIOUtilIntfMockRecorder struct {
	mock *MockIOUtilIntf
}

// NewMockIOUtilIntf creates a new mock instance
func NewMockIOUtilIntf(ctrl *gomock.Controller) *MockIOUtilIntf {
	mock := &MockIOUtilIntf{ctrl: ctrl}
	mock.recorder = &MockIOUtilIntfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIOUtilIntf) EXPECT() *MockIOUtilIntfMockRecorder {
	return m.recorder
}

// WriteFile mocks base method
func (m *MockIOUtilIntf) WriteFile(filename string, data []byte, perm os.FileMode) error {
	ret := m.ctrl.Call(m, "WriteFile", filename, data, perm)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile
func (mr *MockIOUtilIntfMockRecorder) WriteFile(filename, data, perm interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockIOUtilIntf)(nil).WriteFile), filename, data, perm)
}

// ReadFile mocks base method
func (m *MockIOUtilIntf) ReadFile(filename string) ([]byte, error) {
	ret := m.ctrl.Call(m, "ReadFile", filename)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile
func (mr *MockIOUtilIntfMockRecorder) ReadFile(filename interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockIOUtilIntf)(nil).ReadFile), filename)
}

// MockExecIntf is a mock of ExecIntf interface
type MockExecIntf struct {
	ctrl     *gomock.Controller
	recorder *MockExecIntfMockRecorder
}

// MockExecIntfMockRecorder is the mock recorder for MockExecIntf
type MockExecIntfMockRecorder struct {
	mock *MockExecIntf
}

// NewMockExecIntf creates a new mock instance
func NewMockExecIntf(ctrl *gomock.Controller) *MockExecIntf {
	mock := &MockExecIntf{ctrl: ctrl}
	mock.recorder = &MockExecIntfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecIntf) EXPECT() *MockExecIntfMockRecorder {
	return m.recorder
}

// CommandCombinedOutput mocks base method
func (m *MockExecIntf) CommandCombinedOutput(name string, arg ...string) ([]byte, error) {
	varargs := []interface{}{name}
	for _, a := range arg {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommandCombinedOutput", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandCombinedOutput indicates an expected call of CommandCombinedOutput
func (mr *MockExecIntfMockRecorder) CommandCombinedOutput(name interface{}, arg ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name}, arg...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandCombinedOutput", reflect.TypeOf((*MockExecIntf)(nil).CommandCombinedOutput), varargs...)
}

// MockOSFileInfo is a mock of OSFileInfo interface
type MockOSFileInfo struct {
	ctrl     *gomock.Controller
	recorder *MockOSFileInfoMockRecorder
}

// MockOSFileInfoMockRecorder is the mock recorder for MockOSFileInfo
type MockOSFileInfoMockRecorder struct {
	mock *MockOSFileInfo
}

// NewMockOSFileInfo creates a new mock instance
func NewMockOSFileInfo(ctrl *gomock.Controller) *MockOSFileInfo {
	mock := &MockOSFileInfo{ctrl: ctrl}
	mock.recorder = &MockOSFileInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOSFileInfo) EXPECT() *MockOSFileInfoMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockOSFileInfo) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockOSFileInfoMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockOSFileInfo)(nil).Name))
}

// Size mocks base method
func (m *MockOSFileInfo) Size() int64 {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockOSFileInfoMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockOSFileInfo)(nil).Size))
}

// Mode mocks base method
func (m *MockOSFileInfo) Mode() os.FileMode {
	ret := m.ctrl.Call(m, "Mode")
	ret0, _ := ret[0].(os.FileMode)
	return ret0
}

// Mode indicates an expected call of Mode
func (mr *MockOSFileInfoMockRecorder) Mode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mode", reflect.TypeOf((*MockOSFileInfo)(nil).Mode))
}

// ModTime mocks base method
func (m *MockOSFileInfo) ModTime() time.Time {
	ret := m.ctrl.Call(m, "ModTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// ModTime indicates an expected call of ModTime
func (mr *MockOSFileInfoMockRecorder) ModTime() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModTime", reflect.TypeOf((*MockOSFileInfo)(nil).ModTime))
}

// IsDir mocks base method
func (m *MockOSFileInfo) IsDir() bool {
	ret := m.ctrl.Call(m, "IsDir")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDir indicates an expected call of IsDir
func (mr *MockOSFileInfoMockRecorder) IsDir() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDir", reflect.TypeOf((*MockOSFileInfo)(nil).IsDir))
}

// Sys mocks base method
func (m *MockOSFileInfo) Sys() interface{} {
	ret := m.ctrl.Call(m, "Sys")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Sys indicates an expected call of Sys
func (mr *MockOSFileInfoMockRecorder) Sys() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sys", reflect.TypeOf((*MockOSFileInfo)(nil).Sys))
}
