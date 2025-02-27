// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.2
// source: proto/payment.proto

package paymentpb

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

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	CreatePayment(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentRes, error)
	GetPayment(ctx context.Context, in *GetPaymentReq, opts ...grpc.CallOption) (*GetPaymentRes, error)
	CancelPayment(ctx context.Context, in *CancelPaymentReq, opts ...grpc.CallOption) (*CancelPaymentRes, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentRes, error) {
	out := new(CreatePaymentRes)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/createPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetPayment(ctx context.Context, in *GetPaymentReq, opts ...grpc.CallOption) (*GetPaymentRes, error) {
	out := new(GetPaymentRes)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/getPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CancelPayment(ctx context.Context, in *CancelPaymentReq, opts ...grpc.CallOption) (*CancelPaymentRes, error) {
	out := new(CancelPaymentRes)
	err := c.cc.Invoke(ctx, "/payment.PaymentService/cancelPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility
type PaymentServiceServer interface {
	CreatePayment(context.Context, *CreatePaymentReq) (*CreatePaymentRes, error)
	GetPayment(context.Context, *GetPaymentReq) (*GetPaymentRes, error)
	CancelPayment(context.Context, *CancelPaymentReq) (*CancelPaymentRes, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServiceServer struct {
}

func (UnimplementedPaymentServiceServer) CreatePayment(context.Context, *CreatePaymentReq) (*CreatePaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedPaymentServiceServer) GetPayment(context.Context, *GetPaymentReq) (*GetPaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedPaymentServiceServer) CancelPayment(context.Context, *CancelPaymentReq) (*CancelPaymentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelPayment not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/createPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CreatePayment(ctx, req.(*CreatePaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/getPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetPayment(ctx, req.(*GetPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CancelPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CancelPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.PaymentService/cancelPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CancelPayment(ctx, req.(*CancelPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createPayment",
			Handler:    _PaymentService_CreatePayment_Handler,
		},
		{
			MethodName: "getPayment",
			Handler:    _PaymentService_GetPayment_Handler,
		},
		{
			MethodName: "cancelPayment",
			Handler:    _PaymentService_CancelPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/payment.proto",
}
