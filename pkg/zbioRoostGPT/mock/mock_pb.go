// Code generated by MockGen. DO NOT EDIT.
// Source: C:\var\tmp\Roost\RoostGPT\MiniProjects\1733318482\source\awesome-go\pkg\zbioRoostGPT\generated\zbioRoostGPT_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	generated "github.com/zbioRoostGPT/zbioRoostGPT/generated"
	grpc "google.golang.org/grpc"
)

// MockRoostGPTClient is a mock of RoostGPTClient interface.
type MockRoostGPTClient struct {
	ctrl     *gomock.Controller
	recorder *MockRoostGPTClientMockRecorder
}

// MockRoostGPTClientMockRecorder is the mock recorder for MockRoostGPTClient.
type MockRoostGPTClientMockRecorder struct {
	mock *MockRoostGPTClient
}

// NewMockRoostGPTClient creates a new mock instance.
func NewMockRoostGPTClient(ctrl *gomock.Controller) *MockRoostGPTClient {
	mock := &MockRoostGPTClient{ctrl: ctrl}
	mock.recorder = &MockRoostGPTClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoostGPTClient) EXPECT() *MockRoostGPTClientMockRecorder {
	return m.recorder
}

// AbortTestExecute mocks base method.
func (m *MockRoostGPTClient) AbortTestExecute(ctx context.Context, in *generated.AbortTestExecuteRequest, opts ...grpc.CallOption) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AbortTestExecute", varargs...)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbortTestExecute indicates an expected call of AbortTestExecute.
func (mr *MockRoostGPTClientMockRecorder) AbortTestExecute(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortTestExecute", reflect.TypeOf((*MockRoostGPTClient)(nil).AbortTestExecute), varargs...)
}

// AbortTrigger mocks base method.
func (m *MockRoostGPTClient) AbortTrigger(ctx context.Context, in *generated.AbortTriggerRequest, opts ...grpc.CallOption) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AbortTrigger", varargs...)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbortTrigger indicates an expected call of AbortTrigger.
func (mr *MockRoostGPTClientMockRecorder) AbortTrigger(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortTrigger", reflect.TypeOf((*MockRoostGPTClient)(nil).AbortTrigger), varargs...)
}

// AddTest mocks base method.
func (m *MockRoostGPTClient) AddTest(ctx context.Context, in *generated.AddTestRequest, opts ...grpc.CallOption) (*generated.TestGptEntity, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddTest", varargs...)
	ret0, _ := ret[0].(*generated.TestGptEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTest indicates an expected call of AddTest.
func (mr *MockRoostGPTClientMockRecorder) AddTest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTest", reflect.TypeOf((*MockRoostGPTClient)(nil).AddTest), varargs...)
}

// DeleteTest mocks base method.
func (m *MockRoostGPTClient) DeleteTest(ctx context.Context, in *generated.DeleteTestRequest, opts ...grpc.CallOption) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTest", varargs...)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTest indicates an expected call of DeleteTest.
func (mr *MockRoostGPTClientMockRecorder) DeleteTest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTest", reflect.TypeOf((*MockRoostGPTClient)(nil).DeleteTest), varargs...)
}

// EditTest mocks base method.
func (m *MockRoostGPTClient) EditTest(ctx context.Context, in *generated.EditTestRequest, opts ...grpc.CallOption) (*generated.TestGptEntity, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditTest", varargs...)
	ret0, _ := ret[0].(*generated.TestGptEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditTest indicates an expected call of EditTest.
func (mr *MockRoostGPTClientMockRecorder) EditTest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditTest", reflect.TypeOf((*MockRoostGPTClient)(nil).EditTest), varargs...)
}

// EditTriggerEvent mocks base method.
func (m *MockRoostGPTClient) EditTriggerEvent(ctx context.Context, in *generated.EditTriggerEventRequest, opts ...grpc.CallOption) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditTriggerEvent", varargs...)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditTriggerEvent indicates an expected call of EditTriggerEvent.
func (mr *MockRoostGPTClientMockRecorder) EditTriggerEvent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditTriggerEvent", reflect.TypeOf((*MockRoostGPTClient)(nil).EditTriggerEvent), varargs...)
}

// ExecuteTest mocks base method.
func (m *MockRoostGPTClient) ExecuteTest(ctx context.Context, in *generated.ExecuteTestRequest, opts ...grpc.CallOption) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecuteTest", varargs...)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteTest indicates an expected call of ExecuteTest.
func (mr *MockRoostGPTClientMockRecorder) ExecuteTest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteTest", reflect.TypeOf((*MockRoostGPTClient)(nil).ExecuteTest), varargs...)
}

// GetAllEvents mocks base method.
func (m *MockRoostGPTClient) GetAllEvents(ctx context.Context, in *generated.GetAllEventsRequest, opts ...grpc.CallOption) (*generated.GetAllEventsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllEvents", varargs...)
	ret0, _ := ret[0].(*generated.GetAllEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEvents indicates an expected call of GetAllEvents.
func (mr *MockRoostGPTClientMockRecorder) GetAllEvents(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEvents", reflect.TypeOf((*MockRoostGPTClient)(nil).GetAllEvents), varargs...)
}

// GetAllTests mocks base method.
func (m *MockRoostGPTClient) GetAllTests(ctx context.Context, in *generated.GetAllTestsRequest, opts ...grpc.CallOption) (*generated.GetAllTestsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllTests", varargs...)
	ret0, _ := ret[0].(*generated.GetAllTestsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTests indicates an expected call of GetAllTests.
func (mr *MockRoostGPTClientMockRecorder) GetAllTests(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTests", reflect.TypeOf((*MockRoostGPTClient)(nil).GetAllTests), varargs...)
}

// GetLogs mocks base method.
func (m *MockRoostGPTClient) GetLogs(ctx context.Context, in *generated.GetLogsRequest, opts ...grpc.CallOption) (*generated.GetLogsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLogs", varargs...)
	ret0, _ := ret[0].(*generated.GetLogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockRoostGPTClientMockRecorder) GetLogs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockRoostGPTClient)(nil).GetLogs), varargs...)
}

// GetOneEvent mocks base method.
func (m *MockRoostGPTClient) GetOneEvent(ctx context.Context, in *generated.GetOneEventRequest, opts ...grpc.CallOption) (*generated.Event, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOneEvent", varargs...)
	ret0, _ := ret[0].(*generated.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneEvent indicates an expected call of GetOneEvent.
func (mr *MockRoostGPTClientMockRecorder) GetOneEvent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneEvent", reflect.TypeOf((*MockRoostGPTClient)(nil).GetOneEvent), varargs...)
}

// GetOneTest mocks base method.
func (m *MockRoostGPTClient) GetOneTest(ctx context.Context, in *generated.GetOneTestRequest, opts ...grpc.CallOption) (*generated.TestGptEntity, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOneTest", varargs...)
	ret0, _ := ret[0].(*generated.TestGptEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneTest indicates an expected call of GetOneTest.
func (mr *MockRoostGPTClientMockRecorder) GetOneTest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneTest", reflect.TypeOf((*MockRoostGPTClient)(nil).GetOneTest), varargs...)
}

// GetTestExecutionFileStatus mocks base method.
func (m *MockRoostGPTClient) GetTestExecutionFileStatus(ctx context.Context, in *generated.GetTestExecutionFileStatusRequest, opts ...grpc.CallOption) (*generated.TestExecutionFileStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTestExecutionFileStatus", varargs...)
	ret0, _ := ret[0].(*generated.TestExecutionFileStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestExecutionFileStatus indicates an expected call of GetTestExecutionFileStatus.
func (mr *MockRoostGPTClientMockRecorder) GetTestExecutionFileStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestExecutionFileStatus", reflect.TypeOf((*MockRoostGPTClient)(nil).GetTestExecutionFileStatus), varargs...)
}

// GetTestExecutionReport mocks base method.
func (m *MockRoostGPTClient) GetTestExecutionReport(ctx context.Context, in *generated.GetTestExecutionReportRequest, opts ...grpc.CallOption) (*generated.TestExecutionReport, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTestExecutionReport", varargs...)
	ret0, _ := ret[0].(*generated.TestExecutionReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestExecutionReport indicates an expected call of GetTestExecutionReport.
func (mr *MockRoostGPTClientMockRecorder) GetTestExecutionReport(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestExecutionReport", reflect.TypeOf((*MockRoostGPTClient)(nil).GetTestExecutionReport), varargs...)
}

// RetriggerTest mocks base method.
func (m *MockRoostGPTClient) RetriggerTest(ctx context.Context, in *generated.RetriggerTestRequest, opts ...grpc.CallOption) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RetriggerTest", varargs...)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetriggerTest indicates an expected call of RetriggerTest.
func (mr *MockRoostGPTClientMockRecorder) RetriggerTest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetriggerTest", reflect.TypeOf((*MockRoostGPTClient)(nil).RetriggerTest), varargs...)
}

// TriggerTest mocks base method.
func (m *MockRoostGPTClient) TriggerTest(ctx context.Context, in *generated.TriggerTestRequest, opts ...grpc.CallOption) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TriggerTest", varargs...)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerTest indicates an expected call of TriggerTest.
func (mr *MockRoostGPTClientMockRecorder) TriggerTest(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerTest", reflect.TypeOf((*MockRoostGPTClient)(nil).TriggerTest), varargs...)
}

// MockRoostGPTServer is a mock of RoostGPTServer interface.
type MockRoostGPTServer struct {
	ctrl     *gomock.Controller
	recorder *MockRoostGPTServerMockRecorder
}

// MockRoostGPTServerMockRecorder is the mock recorder for MockRoostGPTServer.
type MockRoostGPTServerMockRecorder struct {
	mock *MockRoostGPTServer
}

// NewMockRoostGPTServer creates a new mock instance.
func NewMockRoostGPTServer(ctrl *gomock.Controller) *MockRoostGPTServer {
	mock := &MockRoostGPTServer{ctrl: ctrl}
	mock.recorder = &MockRoostGPTServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoostGPTServer) EXPECT() *MockRoostGPTServerMockRecorder {
	return m.recorder
}

// AbortTestExecute mocks base method.
func (m *MockRoostGPTServer) AbortTestExecute(arg0 context.Context, arg1 *generated.AbortTestExecuteRequest) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortTestExecute", arg0, arg1)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbortTestExecute indicates an expected call of AbortTestExecute.
func (mr *MockRoostGPTServerMockRecorder) AbortTestExecute(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortTestExecute", reflect.TypeOf((*MockRoostGPTServer)(nil).AbortTestExecute), arg0, arg1)
}

// AbortTrigger mocks base method.
func (m *MockRoostGPTServer) AbortTrigger(arg0 context.Context, arg1 *generated.AbortTriggerRequest) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbortTrigger", arg0, arg1)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbortTrigger indicates an expected call of AbortTrigger.
func (mr *MockRoostGPTServerMockRecorder) AbortTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortTrigger", reflect.TypeOf((*MockRoostGPTServer)(nil).AbortTrigger), arg0, arg1)
}

// AddTest mocks base method.
func (m *MockRoostGPTServer) AddTest(arg0 context.Context, arg1 *generated.AddTestRequest) (*generated.TestGptEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTest", arg0, arg1)
	ret0, _ := ret[0].(*generated.TestGptEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTest indicates an expected call of AddTest.
func (mr *MockRoostGPTServerMockRecorder) AddTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTest", reflect.TypeOf((*MockRoostGPTServer)(nil).AddTest), arg0, arg1)
}

// DeleteTest mocks base method.
func (m *MockRoostGPTServer) DeleteTest(arg0 context.Context, arg1 *generated.DeleteTestRequest) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTest", arg0, arg1)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTest indicates an expected call of DeleteTest.
func (mr *MockRoostGPTServerMockRecorder) DeleteTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTest", reflect.TypeOf((*MockRoostGPTServer)(nil).DeleteTest), arg0, arg1)
}

// EditTest mocks base method.
func (m *MockRoostGPTServer) EditTest(arg0 context.Context, arg1 *generated.EditTestRequest) (*generated.TestGptEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditTest", arg0, arg1)
	ret0, _ := ret[0].(*generated.TestGptEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditTest indicates an expected call of EditTest.
func (mr *MockRoostGPTServerMockRecorder) EditTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditTest", reflect.TypeOf((*MockRoostGPTServer)(nil).EditTest), arg0, arg1)
}

// EditTriggerEvent mocks base method.
func (m *MockRoostGPTServer) EditTriggerEvent(arg0 context.Context, arg1 *generated.EditTriggerEventRequest) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditTriggerEvent", arg0, arg1)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditTriggerEvent indicates an expected call of EditTriggerEvent.
func (mr *MockRoostGPTServerMockRecorder) EditTriggerEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditTriggerEvent", reflect.TypeOf((*MockRoostGPTServer)(nil).EditTriggerEvent), arg0, arg1)
}

// ExecuteTest mocks base method.
func (m *MockRoostGPTServer) ExecuteTest(arg0 context.Context, arg1 *generated.ExecuteTestRequest) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteTest", arg0, arg1)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteTest indicates an expected call of ExecuteTest.
func (mr *MockRoostGPTServerMockRecorder) ExecuteTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteTest", reflect.TypeOf((*MockRoostGPTServer)(nil).ExecuteTest), arg0, arg1)
}

// GetAllEvents mocks base method.
func (m *MockRoostGPTServer) GetAllEvents(arg0 context.Context, arg1 *generated.GetAllEventsRequest) (*generated.GetAllEventsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEvents", arg0, arg1)
	ret0, _ := ret[0].(*generated.GetAllEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEvents indicates an expected call of GetAllEvents.
func (mr *MockRoostGPTServerMockRecorder) GetAllEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEvents", reflect.TypeOf((*MockRoostGPTServer)(nil).GetAllEvents), arg0, arg1)
}

// GetAllTests mocks base method.
func (m *MockRoostGPTServer) GetAllTests(arg0 context.Context, arg1 *generated.GetAllTestsRequest) (*generated.GetAllTestsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTests", arg0, arg1)
	ret0, _ := ret[0].(*generated.GetAllTestsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTests indicates an expected call of GetAllTests.
func (mr *MockRoostGPTServerMockRecorder) GetAllTests(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTests", reflect.TypeOf((*MockRoostGPTServer)(nil).GetAllTests), arg0, arg1)
}

// GetLogs mocks base method.
func (m *MockRoostGPTServer) GetLogs(arg0 context.Context, arg1 *generated.GetLogsRequest) (*generated.GetLogsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", arg0, arg1)
	ret0, _ := ret[0].(*generated.GetLogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockRoostGPTServerMockRecorder) GetLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockRoostGPTServer)(nil).GetLogs), arg0, arg1)
}

// GetOneEvent mocks base method.
func (m *MockRoostGPTServer) GetOneEvent(arg0 context.Context, arg1 *generated.GetOneEventRequest) (*generated.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneEvent", arg0, arg1)
	ret0, _ := ret[0].(*generated.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneEvent indicates an expected call of GetOneEvent.
func (mr *MockRoostGPTServerMockRecorder) GetOneEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneEvent", reflect.TypeOf((*MockRoostGPTServer)(nil).GetOneEvent), arg0, arg1)
}

// GetOneTest mocks base method.
func (m *MockRoostGPTServer) GetOneTest(arg0 context.Context, arg1 *generated.GetOneTestRequest) (*generated.TestGptEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOneTest", arg0, arg1)
	ret0, _ := ret[0].(*generated.TestGptEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOneTest indicates an expected call of GetOneTest.
func (mr *MockRoostGPTServerMockRecorder) GetOneTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOneTest", reflect.TypeOf((*MockRoostGPTServer)(nil).GetOneTest), arg0, arg1)
}

// GetTestExecutionFileStatus mocks base method.
func (m *MockRoostGPTServer) GetTestExecutionFileStatus(arg0 context.Context, arg1 *generated.GetTestExecutionFileStatusRequest) (*generated.TestExecutionFileStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestExecutionFileStatus", arg0, arg1)
	ret0, _ := ret[0].(*generated.TestExecutionFileStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestExecutionFileStatus indicates an expected call of GetTestExecutionFileStatus.
func (mr *MockRoostGPTServerMockRecorder) GetTestExecutionFileStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestExecutionFileStatus", reflect.TypeOf((*MockRoostGPTServer)(nil).GetTestExecutionFileStatus), arg0, arg1)
}

// GetTestExecutionReport mocks base method.
func (m *MockRoostGPTServer) GetTestExecutionReport(arg0 context.Context, arg1 *generated.GetTestExecutionReportRequest) (*generated.TestExecutionReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestExecutionReport", arg0, arg1)
	ret0, _ := ret[0].(*generated.TestExecutionReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestExecutionReport indicates an expected call of GetTestExecutionReport.
func (mr *MockRoostGPTServerMockRecorder) GetTestExecutionReport(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestExecutionReport", reflect.TypeOf((*MockRoostGPTServer)(nil).GetTestExecutionReport), arg0, arg1)
}

// RetriggerTest mocks base method.
func (m *MockRoostGPTServer) RetriggerTest(arg0 context.Context, arg1 *generated.RetriggerTestRequest) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetriggerTest", arg0, arg1)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetriggerTest indicates an expected call of RetriggerTest.
func (mr *MockRoostGPTServerMockRecorder) RetriggerTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetriggerTest", reflect.TypeOf((*MockRoostGPTServer)(nil).RetriggerTest), arg0, arg1)
}

// TriggerTest mocks base method.
func (m *MockRoostGPTServer) TriggerTest(arg0 context.Context, arg1 *generated.TriggerTestRequest) (*generated.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerTest", arg0, arg1)
	ret0, _ := ret[0].(*generated.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerTest indicates an expected call of TriggerTest.
func (mr *MockRoostGPTServerMockRecorder) TriggerTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerTest", reflect.TypeOf((*MockRoostGPTServer)(nil).TriggerTest), arg0, arg1)
}

// mustEmbedUnimplementedRoostGPTServer mocks base method.
func (m *MockRoostGPTServer) mustEmbedUnimplementedRoostGPTServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedRoostGPTServer")
}

// mustEmbedUnimplementedRoostGPTServer indicates an expected call of mustEmbedUnimplementedRoostGPTServer.
func (mr *MockRoostGPTServerMockRecorder) mustEmbedUnimplementedRoostGPTServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedRoostGPTServer", reflect.TypeOf((*MockRoostGPTServer)(nil).mustEmbedUnimplementedRoostGPTServer))
}

// MockUnsafeRoostGPTServer is a mock of UnsafeRoostGPTServer interface.
type MockUnsafeRoostGPTServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeRoostGPTServerMockRecorder
}

// MockUnsafeRoostGPTServerMockRecorder is the mock recorder for MockUnsafeRoostGPTServer.
type MockUnsafeRoostGPTServerMockRecorder struct {
	mock *MockUnsafeRoostGPTServer
}

// NewMockUnsafeRoostGPTServer creates a new mock instance.
func NewMockUnsafeRoostGPTServer(ctrl *gomock.Controller) *MockUnsafeRoostGPTServer {
	mock := &MockUnsafeRoostGPTServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeRoostGPTServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeRoostGPTServer) EXPECT() *MockUnsafeRoostGPTServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedRoostGPTServer mocks base method.
func (m *MockUnsafeRoostGPTServer) mustEmbedUnimplementedRoostGPTServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedRoostGPTServer")
}

// mustEmbedUnimplementedRoostGPTServer indicates an expected call of mustEmbedUnimplementedRoostGPTServer.
func (mr *MockUnsafeRoostGPTServerMockRecorder) mustEmbedUnimplementedRoostGPTServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedRoostGPTServer", reflect.TypeOf((*MockUnsafeRoostGPTServer)(nil).mustEmbedUnimplementedRoostGPTServer))
}
