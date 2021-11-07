// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package t2

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

// LiderServiceClient is the client API for LiderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiderServiceClient interface {
	// Simple rpc
	Unirse(ctx context.Context, in *Solicitud, opts ...grpc.CallOption) (*RespuestaSolicitud, error)
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	// etapas
	ProcesarJugada(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Message, error)
	ProcesarJugadaDos(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Message, error)
	ProcesarJugadaTres(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Message, error)
	GetResultadosRonda(ctx context.Context, in *RespuestaSolicitud, opts ...grpc.CallOption) (*ResultadoJugada, error)
	EmpezarEtapa(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type liderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiderServiceClient(cc grpc.ClientConnInterface) LiderServiceClient {
	return &liderServiceClient{cc}
}

func (c *liderServiceClient) Unirse(ctx context.Context, in *Solicitud, opts ...grpc.CallOption) (*RespuestaSolicitud, error) {
	out := new(RespuestaSolicitud)
	err := c.cc.Invoke(ctx, "/t2.LiderService/Unirse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/t2.LiderService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServiceClient) ProcesarJugada(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/t2.LiderService/ProcesarJugada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServiceClient) ProcesarJugadaDos(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/t2.LiderService/ProcesarJugadaDos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServiceClient) ProcesarJugadaTres(ctx context.Context, in *Jugada, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/t2.LiderService/ProcesarJugadaTres", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServiceClient) GetResultadosRonda(ctx context.Context, in *RespuestaSolicitud, opts ...grpc.CallOption) (*ResultadoJugada, error) {
	out := new(ResultadoJugada)
	err := c.cc.Invoke(ctx, "/t2.LiderService/GetResultadosRonda", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liderServiceClient) EmpezarEtapa(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/t2.LiderService/EmpezarEtapa", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiderServiceServer is the server API for LiderService service.
// All implementations must embed UnimplementedLiderServiceServer
// for forward compatibility
type LiderServiceServer interface {
	// Simple rpc
	Unirse(context.Context, *Solicitud) (*RespuestaSolicitud, error)
	SayHello(context.Context, *Message) (*Message, error)
	// etapas
	ProcesarJugada(context.Context, *Jugada) (*Message, error)
	ProcesarJugadaDos(context.Context, *Jugada) (*Message, error)
	ProcesarJugadaTres(context.Context, *Jugada) (*Message, error)
	GetResultadosRonda(context.Context, *RespuestaSolicitud) (*ResultadoJugada, error)
	EmpezarEtapa(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedLiderServiceServer()
}

// UnimplementedLiderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLiderServiceServer struct {
}

func (UnimplementedLiderServiceServer) Unirse(context.Context, *Solicitud) (*RespuestaSolicitud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unirse not implemented")
}
func (UnimplementedLiderServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedLiderServiceServer) ProcesarJugada(context.Context, *Jugada) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcesarJugada not implemented")
}
func (UnimplementedLiderServiceServer) ProcesarJugadaDos(context.Context, *Jugada) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcesarJugadaDos not implemented")
}
func (UnimplementedLiderServiceServer) ProcesarJugadaTres(context.Context, *Jugada) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcesarJugadaTres not implemented")
}
func (UnimplementedLiderServiceServer) GetResultadosRonda(context.Context, *RespuestaSolicitud) (*ResultadoJugada, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResultadosRonda not implemented")
}
func (UnimplementedLiderServiceServer) EmpezarEtapa(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmpezarEtapa not implemented")
}
func (UnimplementedLiderServiceServer) mustEmbedUnimplementedLiderServiceServer() {}

// UnsafeLiderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiderServiceServer will
// result in compilation errors.
type UnsafeLiderServiceServer interface {
	mustEmbedUnimplementedLiderServiceServer()
}

func RegisterLiderServiceServer(s grpc.ServiceRegistrar, srv LiderServiceServer) {
	s.RegisterService(&LiderService_ServiceDesc, srv)
}

func _LiderService_Unirse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Solicitud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).Unirse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t2.LiderService/Unirse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).Unirse(ctx, req.(*Solicitud))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t2.LiderService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderService_ProcesarJugada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).ProcesarJugada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t2.LiderService/ProcesarJugada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).ProcesarJugada(ctx, req.(*Jugada))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderService_ProcesarJugadaDos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).ProcesarJugadaDos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t2.LiderService/ProcesarJugadaDos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).ProcesarJugadaDos(ctx, req.(*Jugada))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderService_ProcesarJugadaTres_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Jugada)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).ProcesarJugadaTres(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t2.LiderService/ProcesarJugadaTres",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).ProcesarJugadaTres(ctx, req.(*Jugada))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderService_GetResultadosRonda_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RespuestaSolicitud)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).GetResultadosRonda(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t2.LiderService/GetResultadosRonda",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).GetResultadosRonda(ctx, req.(*RespuestaSolicitud))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiderService_EmpezarEtapa_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiderServiceServer).EmpezarEtapa(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/t2.LiderService/EmpezarEtapa",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiderServiceServer).EmpezarEtapa(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// LiderService_ServiceDesc is the grpc.ServiceDesc for LiderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "t2.LiderService",
	HandlerType: (*LiderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unirse",
			Handler:    _LiderService_Unirse_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _LiderService_SayHello_Handler,
		},
		{
			MethodName: "ProcesarJugada",
			Handler:    _LiderService_ProcesarJugada_Handler,
		},
		{
			MethodName: "ProcesarJugadaDos",
			Handler:    _LiderService_ProcesarJugadaDos_Handler,
		},
		{
			MethodName: "ProcesarJugadaTres",
			Handler:    _LiderService_ProcesarJugadaTres_Handler,
		},
		{
			MethodName: "GetResultadosRonda",
			Handler:    _LiderService_GetResultadosRonda_Handler,
		},
		{
			MethodName: "EmpezarEtapa",
			Handler:    _LiderService_EmpezarEtapa_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "my_proto.proto",
}