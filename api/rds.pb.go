// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rds.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RDSNode struct {
	Region string `protobuf:"bytes,3,opt,name=region" json:"region,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
}

func (m *RDSNode) Reset()                    { *m = RDSNode{} }
func (m *RDSNode) String() string            { return proto.CompactTextString(m) }
func (*RDSNode) ProtoMessage()               {}
func (*RDSNode) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *RDSNode) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *RDSNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RDSService struct {
	Address       string `protobuf:"bytes,4,opt,name=address" json:"address,omitempty"`
	Port          uint32 `protobuf:"varint,5,opt,name=port" json:"port,omitempty"`
	Engine        string `protobuf:"bytes,6,opt,name=engine" json:"engine,omitempty"`
	EngineVersion string `protobuf:"bytes,7,opt,name=engine_version,json=engineVersion" json:"engine_version,omitempty"`
}

func (m *RDSService) Reset()                    { *m = RDSService{} }
func (m *RDSService) String() string            { return proto.CompactTextString(m) }
func (*RDSService) ProtoMessage()               {}
func (*RDSService) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *RDSService) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RDSService) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RDSService) GetEngine() string {
	if m != nil {
		return m.Engine
	}
	return ""
}

func (m *RDSService) GetEngineVersion() string {
	if m != nil {
		return m.EngineVersion
	}
	return ""
}

type RDSInstanceID struct {
	Region string `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *RDSInstanceID) Reset()                    { *m = RDSInstanceID{} }
func (m *RDSInstanceID) String() string            { return proto.CompactTextString(m) }
func (*RDSInstanceID) ProtoMessage()               {}
func (*RDSInstanceID) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *RDSInstanceID) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *RDSInstanceID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RDSInstance struct {
	Node    *RDSNode    `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Service *RDSService `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
}

func (m *RDSInstance) Reset()                    { *m = RDSInstance{} }
func (m *RDSInstance) String() string            { return proto.CompactTextString(m) }
func (*RDSInstance) ProtoMessage()               {}
func (*RDSInstance) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *RDSInstance) GetNode() *RDSNode {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *RDSInstance) GetService() *RDSService {
	if m != nil {
		return m.Service
	}
	return nil
}

type RDSDiscoverRequest struct {
	AwsAccessKeyId     string `protobuf:"bytes,1,opt,name=aws_access_key_id,json=awsAccessKeyId" json:"aws_access_key_id,omitempty"`
	AwsSecretAccessKey string `protobuf:"bytes,2,opt,name=aws_secret_access_key,json=awsSecretAccessKey" json:"aws_secret_access_key,omitempty"`
}

func (m *RDSDiscoverRequest) Reset()                    { *m = RDSDiscoverRequest{} }
func (m *RDSDiscoverRequest) String() string            { return proto.CompactTextString(m) }
func (*RDSDiscoverRequest) ProtoMessage()               {}
func (*RDSDiscoverRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *RDSDiscoverRequest) GetAwsAccessKeyId() string {
	if m != nil {
		return m.AwsAccessKeyId
	}
	return ""
}

func (m *RDSDiscoverRequest) GetAwsSecretAccessKey() string {
	if m != nil {
		return m.AwsSecretAccessKey
	}
	return ""
}

type RDSDiscoverResponse struct {
	Instances []*RDSInstance `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
}

func (m *RDSDiscoverResponse) Reset()                    { *m = RDSDiscoverResponse{} }
func (m *RDSDiscoverResponse) String() string            { return proto.CompactTextString(m) }
func (*RDSDiscoverResponse) ProtoMessage()               {}
func (*RDSDiscoverResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *RDSDiscoverResponse) GetInstances() []*RDSInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type RDSListRequest struct {
}

func (m *RDSListRequest) Reset()                    { *m = RDSListRequest{} }
func (m *RDSListRequest) String() string            { return proto.CompactTextString(m) }
func (*RDSListRequest) ProtoMessage()               {}
func (*RDSListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

type RDSListResponse struct {
	Instances []*RDSInstance `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
}

func (m *RDSListResponse) Reset()                    { *m = RDSListResponse{} }
func (m *RDSListResponse) String() string            { return proto.CompactTextString(m) }
func (*RDSListResponse) ProtoMessage()               {}
func (*RDSListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *RDSListResponse) GetInstances() []*RDSInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type RDSAddRequest struct {
	AwsAccessKeyId     string           `protobuf:"bytes,1,opt,name=aws_access_key_id,json=awsAccessKeyId" json:"aws_access_key_id,omitempty"`
	AwsSecretAccessKey string           `protobuf:"bytes,2,opt,name=aws_secret_access_key,json=awsSecretAccessKey" json:"aws_secret_access_key,omitempty"`
	Ids                []*RDSInstanceID `protobuf:"bytes,3,rep,name=ids" json:"ids,omitempty"`
}

func (m *RDSAddRequest) Reset()                    { *m = RDSAddRequest{} }
func (m *RDSAddRequest) String() string            { return proto.CompactTextString(m) }
func (*RDSAddRequest) ProtoMessage()               {}
func (*RDSAddRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *RDSAddRequest) GetAwsAccessKeyId() string {
	if m != nil {
		return m.AwsAccessKeyId
	}
	return ""
}

func (m *RDSAddRequest) GetAwsSecretAccessKey() string {
	if m != nil {
		return m.AwsSecretAccessKey
	}
	return ""
}

func (m *RDSAddRequest) GetIds() []*RDSInstanceID {
	if m != nil {
		return m.Ids
	}
	return nil
}

type RDSAddResponse struct {
}

func (m *RDSAddResponse) Reset()                    { *m = RDSAddResponse{} }
func (m *RDSAddResponse) String() string            { return proto.CompactTextString(m) }
func (*RDSAddResponse) ProtoMessage()               {}
func (*RDSAddResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

type RDSRemoveRequest struct {
	Ids []*RDSInstanceID `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *RDSRemoveRequest) Reset()                    { *m = RDSRemoveRequest{} }
func (m *RDSRemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RDSRemoveRequest) ProtoMessage()               {}
func (*RDSRemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *RDSRemoveRequest) GetIds() []*RDSInstanceID {
	if m != nil {
		return m.Ids
	}
	return nil
}

type RDSRemoveResponse struct {
}

func (m *RDSRemoveResponse) Reset()                    { *m = RDSRemoveResponse{} }
func (m *RDSRemoveResponse) String() string            { return proto.CompactTextString(m) }
func (*RDSRemoveResponse) ProtoMessage()               {}
func (*RDSRemoveResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func init() {
	proto.RegisterType((*RDSNode)(nil), "api.RDSNode")
	proto.RegisterType((*RDSService)(nil), "api.RDSService")
	proto.RegisterType((*RDSInstanceID)(nil), "api.RDSInstanceID")
	proto.RegisterType((*RDSInstance)(nil), "api.RDSInstance")
	proto.RegisterType((*RDSDiscoverRequest)(nil), "api.RDSDiscoverRequest")
	proto.RegisterType((*RDSDiscoverResponse)(nil), "api.RDSDiscoverResponse")
	proto.RegisterType((*RDSListRequest)(nil), "api.RDSListRequest")
	proto.RegisterType((*RDSListResponse)(nil), "api.RDSListResponse")
	proto.RegisterType((*RDSAddRequest)(nil), "api.RDSAddRequest")
	proto.RegisterType((*RDSAddResponse)(nil), "api.RDSAddResponse")
	proto.RegisterType((*RDSRemoveRequest)(nil), "api.RDSRemoveRequest")
	proto.RegisterType((*RDSRemoveResponse)(nil), "api.RDSRemoveResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RDS service

type RDSClient interface {
	Discover(ctx context.Context, in *RDSDiscoverRequest, opts ...grpc.CallOption) (*RDSDiscoverResponse, error)
	List(ctx context.Context, in *RDSListRequest, opts ...grpc.CallOption) (*RDSListResponse, error)
	Add(ctx context.Context, in *RDSAddRequest, opts ...grpc.CallOption) (*RDSAddResponse, error)
	Remove(ctx context.Context, in *RDSRemoveRequest, opts ...grpc.CallOption) (*RDSRemoveResponse, error)
}

type rDSClient struct {
	cc *grpc.ClientConn
}

func NewRDSClient(cc *grpc.ClientConn) RDSClient {
	return &rDSClient{cc}
}

func (c *rDSClient) Discover(ctx context.Context, in *RDSDiscoverRequest, opts ...grpc.CallOption) (*RDSDiscoverResponse, error) {
	out := new(RDSDiscoverResponse)
	err := grpc.Invoke(ctx, "/api.RDS/Discover", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rDSClient) List(ctx context.Context, in *RDSListRequest, opts ...grpc.CallOption) (*RDSListResponse, error) {
	out := new(RDSListResponse)
	err := grpc.Invoke(ctx, "/api.RDS/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rDSClient) Add(ctx context.Context, in *RDSAddRequest, opts ...grpc.CallOption) (*RDSAddResponse, error) {
	out := new(RDSAddResponse)
	err := grpc.Invoke(ctx, "/api.RDS/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rDSClient) Remove(ctx context.Context, in *RDSRemoveRequest, opts ...grpc.CallOption) (*RDSRemoveResponse, error) {
	out := new(RDSRemoveResponse)
	err := grpc.Invoke(ctx, "/api.RDS/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RDS service

type RDSServer interface {
	Discover(context.Context, *RDSDiscoverRequest) (*RDSDiscoverResponse, error)
	List(context.Context, *RDSListRequest) (*RDSListResponse, error)
	Add(context.Context, *RDSAddRequest) (*RDSAddResponse, error)
	Remove(context.Context, *RDSRemoveRequest) (*RDSRemoveResponse, error)
}

func RegisterRDSServer(s *grpc.Server, srv RDSServer) {
	s.RegisterService(&_RDS_serviceDesc, srv)
}

func _RDS_Discover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RDSDiscoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDSServer).Discover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RDS/Discover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDSServer).Discover(ctx, req.(*RDSDiscoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RDS_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RDSListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDSServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RDS/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDSServer).List(ctx, req.(*RDSListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RDS_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RDSAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDSServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RDS/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDSServer).Add(ctx, req.(*RDSAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RDS_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RDSRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDSServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RDS/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDSServer).Remove(ctx, req.(*RDSRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RDS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.RDS",
	HandlerType: (*RDSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Discover",
			Handler:    _RDS_Discover_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RDS_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _RDS_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _RDS_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rds.proto",
}

func init() { proto.RegisterFile("rds.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x95, 0x26, 0xf4, 0xcf, 0x74, 0x77, 0x9b, 0x9d, 0xb2, 0x4b, 0x54, 0x71, 0xa8, 0x22,
	0x90, 0xba, 0x7b, 0x68, 0xa1, 0x5c, 0x10, 0x9c, 0xba, 0x0a, 0x87, 0x76, 0x11, 0x07, 0x47, 0x02,
	0x89, 0x4b, 0x65, 0x62, 0xab, 0xb2, 0x60, 0xe3, 0x10, 0x87, 0x56, 0x7b, 0xe5, 0xca, 0x71, 0x1f,
	0x85, 0x47, 0xe1, 0x15, 0x78, 0x10, 0x14, 0xc7, 0x6e, 0xd3, 0x82, 0x38, 0x70, 0xe0, 0x66, 0xcf,
	0xf8, 0xfb, 0xf9, 0xfb, 0x3c, 0x51, 0xa0, 0x93, 0x33, 0x35, 0xce, 0x72, 0x59, 0x48, 0x74, 0x69,
	0x26, 0x06, 0x0f, 0x57, 0x52, 0xae, 0x3e, 0xf1, 0x09, 0xcd, 0xc4, 0x84, 0xa6, 0xa9, 0x2c, 0x68,
	0x21, 0x64, 0x6a, 0x8e, 0x84, 0x33, 0x68, 0x91, 0x28, 0x7e, 0x23, 0x19, 0xc7, 0x73, 0x68, 0xe6,
	0x7c, 0x25, 0x64, 0x1a, 0xb8, 0x43, 0x67, 0xd4, 0x21, 0x66, 0x87, 0x08, 0x5e, 0x4a, 0x6f, 0x78,
	0xe0, 0xe9, 0xaa, 0x5e, 0x2f, 0xbc, 0xb6, 0xe3, 0x37, 0x16, 0x5e, 0xbb, 0xe1, 0xbb, 0xe1, 0x37,
	0x07, 0x80, 0x44, 0x71, 0xcc, 0xf3, 0xb5, 0x48, 0x38, 0x06, 0xd0, 0xa2, 0x8c, 0xe5, 0x5c, 0x29,
	0xa3, 0xb0, 0xdb, 0x12, 0x94, 0xc9, 0xbc, 0x08, 0xee, 0x0d, 0x9d, 0xd1, 0x31, 0xd1, 0xeb, 0xf2,
	0x52, 0x9e, 0xae, 0x44, 0xca, 0x83, 0x66, 0x75, 0x69, 0xb5, 0xc3, 0xc7, 0x70, 0x52, 0xad, 0x96,
	0x6b, 0x9e, 0xab, 0xd2, 0x54, 0x4b, 0xf7, 0x8f, 0xab, 0xea, 0xdb, 0xaa, 0x58, 0xf7, 0xb1, 0xf0,
	0xda, 0xae, 0xef, 0x85, 0x2f, 0xe1, 0x98, 0x44, 0xf1, 0x3c, 0x55, 0x05, 0x4d, 0x13, 0x3e, 0x8f,
	0x6a, 0xb1, 0x9c, 0x3f, 0xc6, 0x6a, 0xec, 0x62, 0x85, 0xef, 0xa1, 0x5b, 0x13, 0xe3, 0x10, 0xbc,
	0x54, 0x32, 0xae, 0x85, 0xdd, 0xe9, 0xd1, 0x98, 0x66, 0x62, 0x6c, 0x5e, 0x8b, 0xe8, 0x0e, 0x5e,
	0x40, 0x4b, 0x55, 0xb9, 0x35, 0xa7, 0x3b, 0xed, 0xd9, 0x43, 0xe6, 0x39, 0x88, 0xed, 0x87, 0x39,
	0x20, 0x89, 0xe2, 0x48, 0xa8, 0x44, 0xae, 0x79, 0x4e, 0xf8, 0xe7, 0x2f, 0x5c, 0x15, 0x78, 0x01,
	0xa7, 0x74, 0xa3, 0x96, 0x34, 0x49, 0xb8, 0x52, 0xcb, 0x8f, 0xfc, 0x76, 0x29, 0x98, 0x31, 0x7a,
	0x42, 0x37, 0x6a, 0xa6, 0xeb, 0xd7, 0xfc, 0x76, 0xce, 0xf0, 0x29, 0x9c, 0x95, 0x47, 0x15, 0x4f,
	0x72, 0x5e, 0xd4, 0x14, 0x26, 0x01, 0xd2, 0x8d, 0x8a, 0x75, 0x6f, 0x2b, 0x0a, 0x5f, 0x41, 0x7f,
	0xef, 0x4e, 0x95, 0xc9, 0x54, 0x71, 0x1c, 0x43, 0x47, 0x98, 0x8c, 0x2a, 0x70, 0x86, 0xee, 0xa8,
	0x3b, 0xf5, 0xad, 0x6f, 0x1b, 0x9e, 0xec, 0x8e, 0x84, 0x3e, 0x9c, 0x90, 0x28, 0x7e, 0x2d, 0x54,
	0x61, 0x6c, 0x87, 0x33, 0xe8, 0x6d, 0x2b, 0xff, 0x08, 0xbd, 0x73, 0xf4, 0xa4, 0x66, 0x8c, 0xfd,
	0x97, 0xb7, 0xc0, 0x47, 0xe0, 0x0a, 0xa6, 0x02, 0x57, 0x3b, 0xc3, 0x43, 0x67, 0xf3, 0x88, 0x94,
	0x6d, 0x13, 0x55, 0x9b, 0xaa, 0x72, 0x85, 0xcf, 0xc1, 0x27, 0x51, 0x4c, 0xf8, 0x8d, 0x5c, 0x73,
	0xeb, 0xd4, 0xb0, 0x9c, 0xbf, 0xb3, 0xfa, 0x70, 0x5a, 0x53, 0x56, 0xb8, 0xe9, 0xf7, 0x06, 0xb8,
	0x24, 0x8a, 0xf1, 0x1d, 0xb4, 0xed, 0x5c, 0xf0, 0x81, 0x25, 0x1c, 0x7c, 0x1d, 0x83, 0xe0, 0xf7,
	0x86, 0x71, 0x15, 0x7c, 0xfd, 0xf1, 0xf3, 0xae, 0x81, 0xe8, 0x4f, 0xd6, 0x4f, 0x26, 0x39, 0x53,
	0x13, 0x66, 0x61, 0x57, 0xe0, 0x95, 0x73, 0xc1, 0xbe, 0xd5, 0xd6, 0xe6, 0x36, 0xb8, 0xbf, 0x5f,
	0x34, 0xb0, 0x9e, 0x86, 0x75, 0xb0, 0x65, 0x60, 0x78, 0x05, 0xee, 0x8c, 0x31, 0xdc, 0x26, 0xdb,
	0x0d, 0x69, 0xd0, 0xdf, 0xab, 0x19, 0x00, 0x6a, 0xc0, 0x51, 0x68, 0x01, 0x2f, 0x9c, 0x4b, 0xbc,
	0x86, 0x66, 0x15, 0x1d, 0xcf, 0xac, 0x64, 0xef, 0x11, 0x07, 0xe7, 0x87, 0xe5, 0x7d, 0xd8, 0x65,
	0x0d, 0xf6, 0xa1, 0xa9, 0xff, 0x56, 0xcf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x5a, 0x76,
	0x73, 0xdd, 0x04, 0x00, 0x00,
}