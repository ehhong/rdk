// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RobotServiceClient is the client API for RobotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RobotServiceClient interface {
	// Status returns the robot's underlying status.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	// StatusStream periodically sends the robot's status.
	StatusStream(ctx context.Context, in *StatusStreamRequest, opts ...grpc.CallOption) (RobotService_StatusStreamClient, error)
	// Config gets the config from a server
	// It is only partial a config, including the pieces relevant to remote robots,
	// And not the pieces relevant to local configuration (pins, security keys, etc...)
	Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	// DoAction runs an action on the underlying robot.
	DoAction(ctx context.Context, in *DoActionRequest, opts ...grpc.CallOption) (*DoActionResponse, error)
	// TODO(https://github.com/viamrobotics/rdk/issues/407): refactor to functions service
	ExecuteFunction(ctx context.Context, in *ExecuteFunctionRequest, opts ...grpc.CallOption) (*ExecuteFunctionResponse, error)
	ExecuteSource(ctx context.Context, in *ExecuteSourceRequest, opts ...grpc.CallOption) (*ExecuteSourceResponse, error)
	// ResourceRunCommand runs an arbitrary command on a resource if it supports it.
	ResourceRunCommand(ctx context.Context, in *ResourceRunCommandRequest, opts ...grpc.CallOption) (*ResourceRunCommandResponse, error)
}

type robotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRobotServiceClient(cc grpc.ClientConnInterface) RobotServiceClient {
	return &robotServiceClient{cc}
}

func (c *robotServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/proto.api.robot.v1.RobotService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) StatusStream(ctx context.Context, in *StatusStreamRequest, opts ...grpc.CallOption) (RobotService_StatusStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &RobotService_ServiceDesc.Streams[0], "/proto.api.robot.v1.RobotService/StatusStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &robotServiceStatusStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RobotService_StatusStreamClient interface {
	Recv() (*StatusStreamResponse, error)
	grpc.ClientStream
}

type robotServiceStatusStreamClient struct {
	grpc.ClientStream
}

func (x *robotServiceStatusStreamClient) Recv() (*StatusStreamResponse, error) {
	m := new(StatusStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *robotServiceClient) Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/proto.api.robot.v1.RobotService/Config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) DoAction(ctx context.Context, in *DoActionRequest, opts ...grpc.CallOption) (*DoActionResponse, error) {
	out := new(DoActionResponse)
	err := c.cc.Invoke(ctx, "/proto.api.robot.v1.RobotService/DoAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ExecuteFunction(ctx context.Context, in *ExecuteFunctionRequest, opts ...grpc.CallOption) (*ExecuteFunctionResponse, error) {
	out := new(ExecuteFunctionResponse)
	err := c.cc.Invoke(ctx, "/proto.api.robot.v1.RobotService/ExecuteFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ExecuteSource(ctx context.Context, in *ExecuteSourceRequest, opts ...grpc.CallOption) (*ExecuteSourceResponse, error) {
	out := new(ExecuteSourceResponse)
	err := c.cc.Invoke(ctx, "/proto.api.robot.v1.RobotService/ExecuteSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotServiceClient) ResourceRunCommand(ctx context.Context, in *ResourceRunCommandRequest, opts ...grpc.CallOption) (*ResourceRunCommandResponse, error) {
	out := new(ResourceRunCommandResponse)
	err := c.cc.Invoke(ctx, "/proto.api.robot.v1.RobotService/ResourceRunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RobotServiceServer is the server API for RobotService service.
// All implementations must embed UnimplementedRobotServiceServer
// for forward compatibility
type RobotServiceServer interface {
	// Status returns the robot's underlying status.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	// StatusStream periodically sends the robot's status.
	StatusStream(*StatusStreamRequest, RobotService_StatusStreamServer) error
	// Config gets the config from a server
	// It is only partial a config, including the pieces relevant to remote robots,
	// And not the pieces relevant to local configuration (pins, security keys, etc...)
	Config(context.Context, *ConfigRequest) (*ConfigResponse, error)
	// DoAction runs an action on the underlying robot.
	DoAction(context.Context, *DoActionRequest) (*DoActionResponse, error)
	// TODO(https://github.com/viamrobotics/rdk/issues/407): refactor to functions service
	ExecuteFunction(context.Context, *ExecuteFunctionRequest) (*ExecuteFunctionResponse, error)
	ExecuteSource(context.Context, *ExecuteSourceRequest) (*ExecuteSourceResponse, error)
	// ResourceRunCommand runs an arbitrary command on a resource if it supports it.
	ResourceRunCommand(context.Context, *ResourceRunCommandRequest) (*ResourceRunCommandResponse, error)
	mustEmbedUnimplementedRobotServiceServer()
}

// UnimplementedRobotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRobotServiceServer struct {
}

func (UnimplementedRobotServiceServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedRobotServiceServer) StatusStream(*StatusStreamRequest, RobotService_StatusStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method StatusStream not implemented")
}
func (UnimplementedRobotServiceServer) Config(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Config not implemented")
}
func (UnimplementedRobotServiceServer) DoAction(context.Context, *DoActionRequest) (*DoActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoAction not implemented")
}
func (UnimplementedRobotServiceServer) ExecuteFunction(context.Context, *ExecuteFunctionRequest) (*ExecuteFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteFunction not implemented")
}
func (UnimplementedRobotServiceServer) ExecuteSource(context.Context, *ExecuteSourceRequest) (*ExecuteSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteSource not implemented")
}
func (UnimplementedRobotServiceServer) ResourceRunCommand(context.Context, *ResourceRunCommandRequest) (*ResourceRunCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResourceRunCommand not implemented")
}
func (UnimplementedRobotServiceServer) mustEmbedUnimplementedRobotServiceServer() {}

// UnsafeRobotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RobotServiceServer will
// result in compilation errors.
type UnsafeRobotServiceServer interface {
	mustEmbedUnimplementedRobotServiceServer()
}

func RegisterRobotServiceServer(s grpc.ServiceRegistrar, srv RobotServiceServer) {
	s.RegisterService(&RobotService_ServiceDesc, srv)
}

func _RobotService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.robot.v1.RobotService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_StatusStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StatusStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RobotServiceServer).StatusStream(m, &robotServiceStatusStreamServer{stream})
}

type RobotService_StatusStreamServer interface {
	Send(*StatusStreamResponse) error
	grpc.ServerStream
}

type robotServiceStatusStreamServer struct {
	grpc.ServerStream
}

func (x *robotServiceStatusStreamServer) Send(m *StatusStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _RobotService_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.robot.v1.RobotService/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).Config(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_DoAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).DoAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.robot.v1.RobotService/DoAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).DoAction(ctx, req.(*DoActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ExecuteFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ExecuteFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.robot.v1.RobotService/ExecuteFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ExecuteFunction(ctx, req.(*ExecuteFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ExecuteSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ExecuteSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.robot.v1.RobotService/ExecuteSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ExecuteSource(ctx, req.(*ExecuteSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotService_ResourceRunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceRunCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotServiceServer).ResourceRunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.api.robot.v1.RobotService/ResourceRunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotServiceServer).ResourceRunCommand(ctx, req.(*ResourceRunCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RobotService_ServiceDesc is the grpc.ServiceDesc for RobotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RobotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.api.robot.v1.RobotService",
	HandlerType: (*RobotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _RobotService_Status_Handler,
		},
		{
			MethodName: "Config",
			Handler:    _RobotService_Config_Handler,
		},
		{
			MethodName: "DoAction",
			Handler:    _RobotService_DoAction_Handler,
		},
		{
			MethodName: "ExecuteFunction",
			Handler:    _RobotService_ExecuteFunction_Handler,
		},
		{
			MethodName: "ExecuteSource",
			Handler:    _RobotService_ExecuteSource_Handler,
		},
		{
			MethodName: "ResourceRunCommand",
			Handler:    _RobotService_ResourceRunCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StatusStream",
			Handler:       _RobotService_StatusStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/api/robot/v1/robot.proto",
}
