// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package results_go_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ResultsClient is the client API for Results service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResultsClient interface {
	CreateResult(ctx context.Context, in *CreateResultRequest, opts ...grpc.CallOption) (*Result, error)
	UpdateResult(ctx context.Context, in *UpdateResultRequest, opts ...grpc.CallOption) (*Result, error)
	GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*Result, error)
	DeleteResult(ctx context.Context, in *DeleteResultRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListResults(ctx context.Context, in *ListResultsRequest, opts ...grpc.CallOption) (*ListResultsResponse, error)
}

type resultsClient struct {
	cc grpc.ClientConnInterface
}

func NewResultsClient(cc grpc.ClientConnInterface) ResultsClient {
	return &resultsClient{cc}
}

func (c *resultsClient) CreateResult(ctx context.Context, in *CreateResultRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tekton.results.v1alpha1.Results/CreateResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultsClient) UpdateResult(ctx context.Context, in *UpdateResultRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tekton.results.v1alpha1.Results/UpdateResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultsClient) GetResult(ctx context.Context, in *GetResultRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tekton.results.v1alpha1.Results/GetResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultsClient) DeleteResult(ctx context.Context, in *DeleteResultRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/tekton.results.v1alpha1.Results/DeleteResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultsClient) ListResults(ctx context.Context, in *ListResultsRequest, opts ...grpc.CallOption) (*ListResultsResponse, error) {
	out := new(ListResultsResponse)
	err := c.cc.Invoke(ctx, "/tekton.results.v1alpha1.Results/ListResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultsServer is the server API for Results service.
// All implementations must embed UnimplementedResultsServer
// for forward compatibility
type ResultsServer interface {
	CreateResult(context.Context, *CreateResultRequest) (*Result, error)
	UpdateResult(context.Context, *UpdateResultRequest) (*Result, error)
	GetResult(context.Context, *GetResultRequest) (*Result, error)
	DeleteResult(context.Context, *DeleteResultRequest) (*emptypb.Empty, error)
	ListResults(context.Context, *ListResultsRequest) (*ListResultsResponse, error)
	mustEmbedUnimplementedResultsServer()
}

// UnimplementedResultsServer must be embedded to have forward compatible implementations.
type UnimplementedResultsServer struct {
}

func (UnimplementedResultsServer) CreateResult(context.Context, *CreateResultRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResult not implemented")
}
func (UnimplementedResultsServer) UpdateResult(context.Context, *UpdateResultRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResult not implemented")
}
func (UnimplementedResultsServer) GetResult(context.Context, *GetResultRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResult not implemented")
}
func (UnimplementedResultsServer) DeleteResult(context.Context, *DeleteResultRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResult not implemented")
}
func (UnimplementedResultsServer) ListResults(context.Context, *ListResultsRequest) (*ListResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListResults not implemented")
}
func (UnimplementedResultsServer) mustEmbedUnimplementedResultsServer() {}

// UnsafeResultsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResultsServer will
// result in compilation errors.
type UnsafeResultsServer interface {
	mustEmbedUnimplementedResultsServer()
}

func RegisterResultsServer(s grpc.ServiceRegistrar, srv ResultsServer) {
	s.RegisterService(&_Results_serviceDesc, srv)
}

func _Results_CreateResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultsServer).CreateResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tekton.results.v1alpha1.Results/CreateResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultsServer).CreateResult(ctx, req.(*CreateResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Results_UpdateResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultsServer).UpdateResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tekton.results.v1alpha1.Results/UpdateResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultsServer).UpdateResult(ctx, req.(*UpdateResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Results_GetResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultsServer).GetResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tekton.results.v1alpha1.Results/GetResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultsServer).GetResult(ctx, req.(*GetResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Results_DeleteResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultsServer).DeleteResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tekton.results.v1alpha1.Results/DeleteResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultsServer).DeleteResult(ctx, req.(*DeleteResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Results_ListResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultsServer).ListResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tekton.results.v1alpha1.Results/ListResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultsServer).ListResults(ctx, req.(*ListResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Results_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tekton.results.v1alpha1.Results",
	HandlerType: (*ResultsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateResult",
			Handler:    _Results_CreateResult_Handler,
		},
		{
			MethodName: "UpdateResult",
			Handler:    _Results_UpdateResult_Handler,
		},
		{
			MethodName: "GetResult",
			Handler:    _Results_GetResult_Handler,
		},
		{
			MethodName: "DeleteResult",
			Handler:    _Results_DeleteResult_Handler,
		},
		{
			MethodName: "ListResults",
			Handler:    _Results_ListResults_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
