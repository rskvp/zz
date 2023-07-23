// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: station/station.proto

package station

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

const (
	StationAPI_GetManyStations_FullMethodName       = "/station.StationAPI/GetManyStations"
	StationAPI_GetOneStation_FullMethodName         = "/station.StationAPI/GetOneStation"
	StationAPI_SetFavoriteOneStation_FullMethodName = "/station.StationAPI/SetFavoriteOneStation"
)

// StationAPIClient is the client API for StationAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StationAPIClient interface {
	// GetManyStations fetch a paginated list of station.
	GetManyStations(ctx context.Context, in *GetManyStationsRequest, opts ...grpc.CallOption) (*GetManyStationsResponse, error)
	// GetOneStation fetches the details of a station.
	GetOneStation(ctx context.Context, in *GetOneStationRequest, opts ...grpc.CallOption) (*GetOneStationResponse, error)
	// SetFavoriteOneStation set a station to favorite for a user.
	SetFavoriteOneStation(ctx context.Context, in *SetFavoriteOneStationRequest, opts ...grpc.CallOption) (*SetFavoriteOneStationResponse, error)
}

type stationAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewStationAPIClient(cc grpc.ClientConnInterface) StationAPIClient {
	return &stationAPIClient{cc}
}

func (c *stationAPIClient) GetManyStations(ctx context.Context, in *GetManyStationsRequest, opts ...grpc.CallOption) (*GetManyStationsResponse, error) {
	out := new(GetManyStationsResponse)
	err := c.cc.Invoke(ctx, StationAPI_GetManyStations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stationAPIClient) GetOneStation(ctx context.Context, in *GetOneStationRequest, opts ...grpc.CallOption) (*GetOneStationResponse, error) {
	out := new(GetOneStationResponse)
	err := c.cc.Invoke(ctx, StationAPI_GetOneStation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stationAPIClient) SetFavoriteOneStation(ctx context.Context, in *SetFavoriteOneStationRequest, opts ...grpc.CallOption) (*SetFavoriteOneStationResponse, error) {
	out := new(SetFavoriteOneStationResponse)
	err := c.cc.Invoke(ctx, StationAPI_SetFavoriteOneStation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StationAPIServer is the server API for StationAPI service.
// All implementations must embed UnimplementedStationAPIServer
// for forward compatibility
type StationAPIServer interface {
	// GetManyStations fetch a paginated list of station.
	GetManyStations(context.Context, *GetManyStationsRequest) (*GetManyStationsResponse, error)
	// GetOneStation fetches the details of a station.
	GetOneStation(context.Context, *GetOneStationRequest) (*GetOneStationResponse, error)
	// SetFavoriteOneStation set a station to favorite for a user.
	SetFavoriteOneStation(context.Context, *SetFavoriteOneStationRequest) (*SetFavoriteOneStationResponse, error)
	mustEmbedUnimplementedStationAPIServer()
}

// UnimplementedStationAPIServer must be embedded to have forward compatible implementations.
type UnimplementedStationAPIServer struct {
}

func (UnimplementedStationAPIServer) GetManyStations(context.Context, *GetManyStationsRequest) (*GetManyStationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyStations not implemented")
}
func (UnimplementedStationAPIServer) GetOneStation(context.Context, *GetOneStationRequest) (*GetOneStationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneStation not implemented")
}
func (UnimplementedStationAPIServer) SetFavoriteOneStation(context.Context, *SetFavoriteOneStationRequest) (*SetFavoriteOneStationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFavoriteOneStation not implemented")
}
func (UnimplementedStationAPIServer) mustEmbedUnimplementedStationAPIServer() {}

// UnsafeStationAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StationAPIServer will
// result in compilation errors.
type UnsafeStationAPIServer interface {
	mustEmbedUnimplementedStationAPIServer()
}

func RegisterStationAPIServer(s grpc.ServiceRegistrar, srv StationAPIServer) {
	s.RegisterService(&StationAPI_ServiceDesc, srv)
}

func _StationAPI_GetManyStations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyStationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StationAPIServer).GetManyStations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StationAPI_GetManyStations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StationAPIServer).GetManyStations(ctx, req.(*GetManyStationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StationAPI_GetOneStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneStationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StationAPIServer).GetOneStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StationAPI_GetOneStation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StationAPIServer).GetOneStation(ctx, req.(*GetOneStationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StationAPI_SetFavoriteOneStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFavoriteOneStationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StationAPIServer).SetFavoriteOneStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StationAPI_SetFavoriteOneStation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StationAPIServer).SetFavoriteOneStation(ctx, req.(*SetFavoriteOneStationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StationAPI_ServiceDesc is the grpc.ServiceDesc for StationAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StationAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "station.StationAPI",
	HandlerType: (*StationAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManyStations",
			Handler:    _StationAPI_GetManyStations_Handler,
		},
		{
			MethodName: "GetOneStation",
			Handler:    _StationAPI_GetOneStation_Handler,
		},
		{
			MethodName: "SetFavoriteOneStation",
			Handler:    _StationAPI_SetFavoriteOneStation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "station/station.proto",
}
