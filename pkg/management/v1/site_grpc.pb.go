// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: site.proto

package management

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

// SiteServiceClient is the client API for SiteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SiteServiceClient interface {
	// CreateSite - implements (this is a comment for public API)
	CreateSite(ctx context.Context, in *CreateSideRequest, opts ...grpc.CallOption) (*CreateSideResponse, error)
	GetSite(ctx context.Context, in *GetSiteRequest, opts ...grpc.CallOption) (*GetSiteResponse, error)
	DeleteSite(ctx context.Context, in *DeleteSiteRequest, opts ...grpc.CallOption) (*DeleteSiteResponse, error)
	ListSites(ctx context.Context, in *ListSitesRequest, opts ...grpc.CallOption) (*ListSitesResponse, error)
}

type siteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSiteServiceClient(cc grpc.ClientConnInterface) SiteServiceClient {
	return &siteServiceClient{cc}
}

func (c *siteServiceClient) CreateSite(ctx context.Context, in *CreateSideRequest, opts ...grpc.CallOption) (*CreateSideResponse, error) {
	out := new(CreateSideResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.SiteService/CreateSite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteServiceClient) GetSite(ctx context.Context, in *GetSiteRequest, opts ...grpc.CallOption) (*GetSiteResponse, error) {
	out := new(GetSiteResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.SiteService/GetSite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteServiceClient) DeleteSite(ctx context.Context, in *DeleteSiteRequest, opts ...grpc.CallOption) (*DeleteSiteResponse, error) {
	out := new(DeleteSiteResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.SiteService/DeleteSite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *siteServiceClient) ListSites(ctx context.Context, in *ListSitesRequest, opts ...grpc.CallOption) (*ListSitesResponse, error) {
	out := new(ListSitesResponse)
	err := c.cc.Invoke(ctx, "/online_shop.management.v1.SiteService/ListSites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SiteServiceServer is the server API for SiteService service.
// All implementations must embed UnimplementedSiteServiceServer
// for forward compatibility
type SiteServiceServer interface {
	// CreateSite - implements (this is a comment for public API)
	CreateSite(context.Context, *CreateSideRequest) (*CreateSideResponse, error)
	GetSite(context.Context, *GetSiteRequest) (*GetSiteResponse, error)
	DeleteSite(context.Context, *DeleteSiteRequest) (*DeleteSiteResponse, error)
	ListSites(context.Context, *ListSitesRequest) (*ListSitesResponse, error)
	mustEmbedUnimplementedSiteServiceServer()
}

// UnimplementedSiteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSiteServiceServer struct {
}

func (UnimplementedSiteServiceServer) CreateSite(context.Context, *CreateSideRequest) (*CreateSideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSite not implemented")
}
func (UnimplementedSiteServiceServer) GetSite(context.Context, *GetSiteRequest) (*GetSiteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSite not implemented")
}
func (UnimplementedSiteServiceServer) DeleteSite(context.Context, *DeleteSiteRequest) (*DeleteSiteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSite not implemented")
}
func (UnimplementedSiteServiceServer) ListSites(context.Context, *ListSitesRequest) (*ListSitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSites not implemented")
}
func (UnimplementedSiteServiceServer) mustEmbedUnimplementedSiteServiceServer() {}

// UnsafeSiteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SiteServiceServer will
// result in compilation errors.
type UnsafeSiteServiceServer interface {
	mustEmbedUnimplementedSiteServiceServer()
}

func RegisterSiteServiceServer(s grpc.ServiceRegistrar, srv SiteServiceServer) {
	s.RegisterService(&SiteService_ServiceDesc, srv)
}

func _SiteService_CreateSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteServiceServer).CreateSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.SiteService/CreateSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteServiceServer).CreateSite(ctx, req.(*CreateSideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteService_GetSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteServiceServer).GetSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.SiteService/GetSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteServiceServer).GetSite(ctx, req.(*GetSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteService_DeleteSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteServiceServer).DeleteSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.SiteService/DeleteSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteServiceServer).DeleteSite(ctx, req.(*DeleteSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SiteService_ListSites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiteServiceServer).ListSites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/online_shop.management.v1.SiteService/ListSites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiteServiceServer).ListSites(ctx, req.(*ListSitesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SiteService_ServiceDesc is the grpc.ServiceDesc for SiteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SiteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "online_shop.management.v1.SiteService",
	HandlerType: (*SiteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSite",
			Handler:    _SiteService_CreateSite_Handler,
		},
		{
			MethodName: "GetSite",
			Handler:    _SiteService_GetSite_Handler,
		},
		{
			MethodName: "DeleteSite",
			Handler:    _SiteService_DeleteSite_Handler,
		},
		{
			MethodName: "ListSites",
			Handler:    _SiteService_ListSites_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "site.proto",
}
