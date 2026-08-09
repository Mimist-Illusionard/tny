// Code generated from api/v1/grpc/url.proto. DO NOT EDIT.

package urlgrpc

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	URLShortener_CreateShortLink_FullMethodName = "/shortener.v1.URLShortener/CreateShortLink"
	URLShortener_GetOriginalLink_FullMethodName = "/shortener.v1.URLShortener/GetOriginalLink"
)

type URLShortenerClient interface {
	CreateShortLink(ctx context.Context, in *CreateShortLinkRequest, opts ...grpc.CallOption) (*CreateShortLinkResponse, error)
	GetOriginalLink(ctx context.Context, in *GetOriginalLinkRequest, opts ...grpc.CallOption) (*GetOriginalLinkResponse, error)
}

type uRLShortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewURLShortenerClient(cc grpc.ClientConnInterface) URLShortenerClient {
	return &uRLShortenerClient{cc}
}

func (c *uRLShortenerClient) CreateShortLink(ctx context.Context, in *CreateShortLinkRequest, opts ...grpc.CallOption) (*CreateShortLinkResponse, error) {
	out := new(CreateShortLinkResponse)
	callOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, URLShortener_CreateShortLink_FullMethodName, in, out, callOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uRLShortenerClient) GetOriginalLink(ctx context.Context, in *GetOriginalLinkRequest, opts ...grpc.CallOption) (*GetOriginalLinkResponse, error) {
	out := new(GetOriginalLinkResponse)
	callOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	err := c.cc.Invoke(ctx, URLShortener_GetOriginalLink_FullMethodName, in, out, callOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type URLShortenerServer interface {
	CreateShortLink(context.Context, *CreateShortLinkRequest) (*CreateShortLinkResponse, error)
	GetOriginalLink(context.Context, *GetOriginalLinkRequest) (*GetOriginalLinkResponse, error)
	mustEmbedUnimplementedURLShortenerServer()
}

type UnimplementedURLShortenerServer struct{}

func (UnimplementedURLShortenerServer) CreateShortLink(context.Context, *CreateShortLinkRequest) (*CreateShortLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShortLink not implemented")
}

func (UnimplementedURLShortenerServer) GetOriginalLink(context.Context, *GetOriginalLinkRequest) (*GetOriginalLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOriginalLink not implemented")
}

func (UnimplementedURLShortenerServer) mustEmbedUnimplementedURLShortenerServer() {}
func (UnimplementedURLShortenerServer) testEmbeddedByValue()                      {}

type UnsafeURLShortenerServer interface {
	mustEmbedUnimplementedURLShortenerServer()
}

func RegisterURLShortenerServer(s grpc.ServiceRegistrar, srv URLShortenerServer) {
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&URLShortener_ServiceDesc, srv)
}

func _URLShortener_CreateShortLink_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(CreateShortLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServer).CreateShortLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortener_CreateShortLink_FullMethodName,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(URLShortenerServer).CreateShortLink(ctx, req.(*CreateShortLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _URLShortener_GetOriginalLink_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(GetOriginalLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShortenerServer).GetOriginalLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: URLShortener_GetOriginalLink_FullMethodName,
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(URLShortenerServer).GetOriginalLink(ctx, req.(*GetOriginalLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var URLShortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortener.v1.URLShortener",
	HandlerType: (*URLShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShortLink",
			Handler:    _URLShortener_CreateShortLink_Handler,
		},
		{
			MethodName: "GetOriginalLink",
			Handler:    _URLShortener_GetOriginalLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/grpc/url.proto",
}
