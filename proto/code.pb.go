// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/code.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_129000e125143313, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type CodeFile struct {
	Meta                 string   `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeFile) Reset()         { *m = CodeFile{} }
func (m *CodeFile) String() string { return proto.CompactTextString(m) }
func (*CodeFile) ProtoMessage()    {}
func (*CodeFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_129000e125143313, []int{1}
}
func (m *CodeFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeFile.Unmarshal(m, b)
}
func (m *CodeFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeFile.Marshal(b, m, deterministic)
}
func (dst *CodeFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeFile.Merge(dst, src)
}
func (m *CodeFile) XXX_Size() int {
	return xxx_messageInfo_CodeFile.Size(m)
}
func (m *CodeFile) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeFile.DiscardUnknown(m)
}

var xxx_messageInfo_CodeFile proto.InternalMessageInfo

func (m *CodeFile) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *CodeFile) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type VmissRequest struct {
	Files                []*CodeFile `protobuf:"bytes,1,rep,name=files" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VmissRequest) Reset()         { *m = VmissRequest{} }
func (m *VmissRequest) String() string { return proto.CompactTextString(m) }
func (*VmissRequest) ProtoMessage()    {}
func (*VmissRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_code_129000e125143313, []int{2}
}
func (m *VmissRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmissRequest.Unmarshal(m, b)
}
func (m *VmissRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmissRequest.Marshal(b, m, deterministic)
}
func (dst *VmissRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmissRequest.Merge(dst, src)
}
func (m *VmissRequest) XXX_Size() int {
	return xxx_messageInfo_VmissRequest.Size(m)
}
func (m *VmissRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VmissRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VmissRequest proto.InternalMessageInfo

func (m *VmissRequest) GetFiles() []*CodeFile {
	if m != nil {
		return m.Files
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*CodeFile)(nil), "proto.CodeFile")
	proto.RegisterType((*VmissRequest)(nil), "proto.VmissRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VmissRPCClient is the client API for VmissRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VmissRPCClient interface {
	StoreFiles(ctx context.Context, in *VmissRequest, opts ...grpc.CallOption) (*Empty, error)
}

type vmissRPCClient struct {
	cc *grpc.ClientConn
}

func NewVmissRPCClient(cc *grpc.ClientConn) VmissRPCClient {
	return &vmissRPCClient{cc}
}

func (c *vmissRPCClient) StoreFiles(ctx context.Context, in *VmissRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.VmissRPC/StoreFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VmissRPC service

type VmissRPCServer interface {
	StoreFiles(context.Context, *VmissRequest) (*Empty, error)
}

func RegisterVmissRPCServer(s *grpc.Server, srv VmissRPCServer) {
	s.RegisterService(&_VmissRPC_serviceDesc, srv)
}

func _VmissRPC_StoreFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VmissRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VmissRPCServer).StoreFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VmissRPC/StoreFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VmissRPCServer).StoreFiles(ctx, req.(*VmissRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VmissRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VmissRPC",
	HandlerType: (*VmissRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreFiles",
			Handler:    _VmissRPC_StoreFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/code.proto",
}

func init() { proto.RegisterFile("proto/code.proto", fileDescriptor_code_129000e125143313) }

var fileDescriptor_code_129000e125143313 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0x4f, 0x49, 0xd5, 0x03, 0x33, 0x85, 0x58, 0xc1, 0x94, 0x12, 0x3b, 0x17,
	0xab, 0x6b, 0x6e, 0x41, 0x49, 0xa5, 0x92, 0x11, 0x17, 0x87, 0x73, 0x7e, 0x4a, 0xaa, 0x5b, 0x66,
	0x4e, 0xaa, 0x90, 0x10, 0x17, 0x4b, 0x6e, 0x6a, 0x49, 0xa2, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x98, 0x0d, 0x12, 0x4b, 0xca, 0x4f, 0xa9, 0x94, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x09, 0x02,
	0xb3, 0x95, 0x4c, 0xb9, 0x78, 0xc2, 0x72, 0x33, 0x8b, 0x8b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0x54, 0xb9, 0x58, 0xd3, 0x32, 0x73, 0x52, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8,
	0x8d, 0xf8, 0x21, 0x56, 0xe9, 0xc1, 0xcc, 0x0d, 0x82, 0xc8, 0x1a, 0x59, 0x73, 0x71, 0x40, 0xb4,
	0x05, 0x38, 0x0b, 0xe9, 0x73, 0x71, 0x05, 0x97, 0xe4, 0x17, 0x81, 0xe5, 0x8b, 0x85, 0x84, 0xa1,
	0x3a, 0x90, 0x4d, 0x95, 0xe2, 0x81, 0x0a, 0x82, 0xdd, 0x99, 0xc4, 0x06, 0xe6, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x98, 0xed, 0xaf, 0xfb, 0xd2, 0x00, 0x00, 0x00,
}
