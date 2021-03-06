// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package google_security_meshca_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MeshCertificateServiceClient is the client API for MeshCertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeshCertificateServiceClient interface {
	// Using provided CSR, returns a signed certificate that represents a GCP
	// service account identity.
	CreateCertificate(ctx context.Context, in *MeshCertificateRequest, opts ...grpc.CallOption) (*MeshCertificateResponse, error)
}

type meshCertificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMeshCertificateServiceClient(cc grpc.ClientConnInterface) MeshCertificateServiceClient {
	return &meshCertificateServiceClient{cc}
}

func (c *meshCertificateServiceClient) CreateCertificate(ctx context.Context, in *MeshCertificateRequest, opts ...grpc.CallOption) (*MeshCertificateResponse, error) {
	out := new(MeshCertificateResponse)
	err := c.cc.Invoke(ctx, "/google.security.meshca.v1.MeshCertificateService/CreateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeshCertificateServiceServer is the server API for MeshCertificateService service.
// All implementations should embed UnimplementedMeshCertificateServiceServer
// for forward compatibility
type MeshCertificateServiceServer interface {
	// Using provided CSR, returns a signed certificate that represents a GCP
	// service account identity.
	CreateCertificate(context.Context, *MeshCertificateRequest) (*MeshCertificateResponse, error)
}

// UnimplementedMeshCertificateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedMeshCertificateServiceServer struct {
}

func (*UnimplementedMeshCertificateServiceServer) CreateCertificate(context.Context, *MeshCertificateRequest) (*MeshCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCertificate not implemented")
}

func RegisterMeshCertificateServiceServer(s *grpc.Server, srv MeshCertificateServiceServer) {
	s.RegisterService(&_MeshCertificateService_serviceDesc, srv)
}

func _MeshCertificateService_CreateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeshCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshCertificateServiceServer).CreateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.security.meshca.v1.MeshCertificateService/CreateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshCertificateServiceServer).CreateCertificate(ctx, req.(*MeshCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MeshCertificateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.security.meshca.v1.MeshCertificateService",
	HandlerType: (*MeshCertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCertificate",
			Handler:    _MeshCertificateService_CreateCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "istio/google/security/meshca/v1/meshca.proto",
}
