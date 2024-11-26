// Code generated by MockGen. DO NOT EDIT.
// Source: C:\var\tmp\Roost\RoostGPT\MiniProjects\1732655001\source\awesome-go\generated\my_products_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	generated "github.com/avelino/awesome-go/generated"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockProductServiceClient is a mock of ProductServiceClient interface.
type MockProductServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceClientMockRecorder
}

// MockProductServiceClientMockRecorder is the mock recorder for MockProductServiceClient.
type MockProductServiceClientMockRecorder struct {
	mock *MockProductServiceClient
}

// NewMockProductServiceClient creates a new mock instance.
func NewMockProductServiceClient(ctrl *gomock.Controller) *MockProductServiceClient {
	mock := &MockProductServiceClient{ctrl: ctrl}
	mock.recorder = &MockProductServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductServiceClient) EXPECT() *MockProductServiceClientMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductServiceClient) CreateProduct(ctx context.Context, in *generated.Product, opts ...grpc.CallOption) (*generated.Product, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProduct", varargs...)
	ret0, _ := ret[0].(*generated.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductServiceClientMockRecorder) CreateProduct(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductServiceClient)(nil).CreateProduct), varargs...)
}

// DeleteProduct mocks base method.
func (m *MockProductServiceClient) DeleteProduct(ctx context.Context, in *generated.ProductId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteProduct", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductServiceClientMockRecorder) DeleteProduct(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProductServiceClient)(nil).DeleteProduct), varargs...)
}

// GetAllProducts mocks base method.
func (m *MockProductServiceClient) GetAllProducts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*generated.ProductList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllProducts", varargs...)
	ret0, _ := ret[0].(*generated.ProductList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockProductServiceClientMockRecorder) GetAllProducts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockProductServiceClient)(nil).GetAllProducts), varargs...)
}

// GetProduct mocks base method.
func (m *MockProductServiceClient) GetProduct(ctx context.Context, in *generated.ProductId, opts ...grpc.CallOption) (*generated.Product, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProduct", varargs...)
	ret0, _ := ret[0].(*generated.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockProductServiceClientMockRecorder) GetProduct(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductServiceClient)(nil).GetProduct), varargs...)
}

// UpdateProduct mocks base method.
func (m *MockProductServiceClient) UpdateProduct(ctx context.Context, in *generated.UpdateProductRequest, opts ...grpc.CallOption) (*generated.Product, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProduct", varargs...)
	ret0, _ := ret[0].(*generated.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductServiceClientMockRecorder) UpdateProduct(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductServiceClient)(nil).UpdateProduct), varargs...)
}

// MockProductServiceServer is a mock of ProductServiceServer interface.
type MockProductServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceServerMockRecorder
}

// MockProductServiceServerMockRecorder is the mock recorder for MockProductServiceServer.
type MockProductServiceServerMockRecorder struct {
	mock *MockProductServiceServer
}

// NewMockProductServiceServer creates a new mock instance.
func NewMockProductServiceServer(ctrl *gomock.Controller) *MockProductServiceServer {
	mock := &MockProductServiceServer{ctrl: ctrl}
	mock.recorder = &MockProductServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductServiceServer) EXPECT() *MockProductServiceServerMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductServiceServer) CreateProduct(arg0 context.Context, arg1 *generated.Product) (*generated.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(*generated.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductServiceServerMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductServiceServer)(nil).CreateProduct), arg0, arg1)
}

// DeleteProduct mocks base method.
func (m *MockProductServiceServer) DeleteProduct(arg0 context.Context, arg1 *generated.ProductId) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductServiceServerMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProductServiceServer)(nil).DeleteProduct), arg0, arg1)
}

// GetAllProducts mocks base method.
func (m *MockProductServiceServer) GetAllProducts(arg0 context.Context, arg1 *emptypb.Empty) (*generated.ProductList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts", arg0, arg1)
	ret0, _ := ret[0].(*generated.ProductList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockProductServiceServerMockRecorder) GetAllProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockProductServiceServer)(nil).GetAllProducts), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockProductServiceServer) GetProduct(arg0 context.Context, arg1 *generated.ProductId) (*generated.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(*generated.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockProductServiceServerMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductServiceServer)(nil).GetProduct), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockProductServiceServer) UpdateProduct(arg0 context.Context, arg1 *generated.UpdateProductRequest) (*generated.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(*generated.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductServiceServerMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductServiceServer)(nil).UpdateProduct), arg0, arg1)
}

// mustEmbedUnimplementedProductServiceServer mocks base method.
func (m *MockProductServiceServer) mustEmbedUnimplementedProductServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProductServiceServer")
}

// mustEmbedUnimplementedProductServiceServer indicates an expected call of mustEmbedUnimplementedProductServiceServer.
func (mr *MockProductServiceServerMockRecorder) mustEmbedUnimplementedProductServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProductServiceServer", reflect.TypeOf((*MockProductServiceServer)(nil).mustEmbedUnimplementedProductServiceServer))
}

// MockUnsafeProductServiceServer is a mock of UnsafeProductServiceServer interface.
type MockUnsafeProductServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeProductServiceServerMockRecorder
}

// MockUnsafeProductServiceServerMockRecorder is the mock recorder for MockUnsafeProductServiceServer.
type MockUnsafeProductServiceServerMockRecorder struct {
	mock *MockUnsafeProductServiceServer
}

// NewMockUnsafeProductServiceServer creates a new mock instance.
func NewMockUnsafeProductServiceServer(ctrl *gomock.Controller) *MockUnsafeProductServiceServer {
	mock := &MockUnsafeProductServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeProductServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeProductServiceServer) EXPECT() *MockUnsafeProductServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedProductServiceServer mocks base method.
func (m *MockUnsafeProductServiceServer) mustEmbedUnimplementedProductServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProductServiceServer")
}

// mustEmbedUnimplementedProductServiceServer indicates an expected call of mustEmbedUnimplementedProductServiceServer.
func (mr *MockUnsafeProductServiceServerMockRecorder) mustEmbedUnimplementedProductServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProductServiceServer", reflect.TypeOf((*MockUnsafeProductServiceServer)(nil).mustEmbedUnimplementedProductServiceServer))
}
