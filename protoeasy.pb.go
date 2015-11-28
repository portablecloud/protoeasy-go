// Code generated by protoc-gen-go.
// source: protoeasy.proto
// DO NOT EDIT!

package protoeasy

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

type Directives struct {
	Cpp                     bool     `protobuf:"varint,1,opt,name=cpp" json:"cpp,omitempty"`
	CppRelOutDirPath        string   `protobuf:"bytes,2,opt,name=cpp_rel_out_dir_path" json:"cpp_rel_out_dir_path,omitempty"`
	Csharp                  bool     `protobuf:"varint,3,opt,name=csharp" json:"csharp,omitempty"`
	CsharpRelOutDirPath     string   `protobuf:"bytes,4,opt,name=csharp_rel_out_dir_path" json:"csharp_rel_out_dir_path,omitempty"`
	Objectivec              bool     `protobuf:"varint,5,opt,name=objectivec" json:"objectivec,omitempty"`
	ObjectivecRelOutDirPath string   `protobuf:"bytes,6,opt,name=objectivec_rel_out_dir_path" json:"objectivec_rel_out_dir_path,omitempty"`
	Python                  bool     `protobuf:"varint,7,opt,name=python" json:"python,omitempty"`
	PythonRelOutDirPath     string   `protobuf:"bytes,8,opt,name=python_rel_out_dir_path" json:"python_rel_out_dir_path,omitempty"`
	Ruby                    bool     `protobuf:"varint,9,opt,name=ruby" json:"ruby,omitempty"`
	RubyRelOutDirPath       string   `protobuf:"bytes,10,opt,name=ruby_rel_out_dir_path" json:"ruby_rel_out_dir_path,omitempty"`
	Go                      bool     `protobuf:"varint,11,opt,name=go" json:"go,omitempty"`
	GoRelOutDirPath         string   `protobuf:"bytes,12,opt,name=go_rel_out_dir_path" json:"go_rel_out_dir_path,omitempty"`
	GoImportPath            string   `protobuf:"bytes,14,opt,name=go_import_path" json:"go_import_path,omitempty"`
	Grpc                    bool     `protobuf:"varint,15,opt,name=grpc" json:"grpc,omitempty"`
	GrpcGateway             bool     `protobuf:"varint,16,opt,name=grpc_gateway" json:"grpc_gateway,omitempty"`
	ExcludePattern          []string `protobuf:"bytes,18,rep,name=exclude_pattern" json:"exclude_pattern,omitempty"`
}

func (m *Directives) Reset()                    { *m = Directives{} }
func (m *Directives) String() string            { return proto.CompactTextString(m) }
func (*Directives) ProtoMessage()               {}
func (*Directives) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Args struct {
	Value []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *Args) Reset()                    { *m = Args{} }
func (m *Args) String() string            { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()               {}
func (*Args) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CompileRequest struct {
	Tar        []byte      `protobuf:"bytes,1,opt,name=tar,proto3" json:"tar,omitempty"`
	Directives *Directives `protobuf:"bytes,2,opt,name=directives" json:"directives,omitempty"`
}

func (m *CompileRequest) Reset()                    { *m = CompileRequest{} }
func (m *CompileRequest) String() string            { return proto.CompactTextString(m) }
func (*CompileRequest) ProtoMessage()               {}
func (*CompileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CompileRequest) GetDirectives() *Directives {
	if m != nil {
		return m.Directives
	}
	return nil
}

type CompileResponse struct {
	Args []*Args `protobuf:"bytes,1,rep,name=args" json:"args,omitempty"`
	Tar  []byte  `protobuf:"bytes,2,opt,name=tar,proto3" json:"tar,omitempty"`
}

func (m *CompileResponse) Reset()                    { *m = CompileResponse{} }
func (m *CompileResponse) String() string            { return proto.CompactTextString(m) }
func (*CompileResponse) ProtoMessage()               {}
func (*CompileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CompileResponse) GetArgs() []*Args {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*Directives)(nil), "protoeasy.Directives")
	proto.RegisterType((*Args)(nil), "protoeasy.Args")
	proto.RegisterType((*CompileRequest)(nil), "protoeasy.CompileRequest")
	proto.RegisterType((*CompileResponse)(nil), "protoeasy.CompileResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for API service

type APIClient interface {
	Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Compile(ctx context.Context, in *CompileRequest, opts ...grpc.CallOption) (*CompileResponse, error) {
	out := new(CompileResponse)
	err := grpc.Invoke(ctx, "/protoeasy.API/Compile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	Compile(context.Context, *CompileRequest) (*CompileResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Compile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CompileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(APIServer).Compile(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoeasy.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compile",
			Handler:    _API_Compile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x53, 0xe2, 0x40,
	0x14, 0x5c, 0x48, 0xf8, 0xc8, 0x4b, 0x36, 0xd9, 0x9a, 0x85, 0x65, 0x16, 0xb4, 0xa4, 0xe2, 0x05,
	0x2f, 0x1c, 0xf0, 0xec, 0x01, 0xf5, 0x20, 0x37, 0xcb, 0x3f, 0x90, 0x1a, 0xc2, 0x54, 0x88, 0x15,
	0x98, 0x71, 0x32, 0x41, 0xf3, 0xe3, 0xad, 0x32, 0x79, 0xe1, 0xcb, 0x8a, 0xb7, 0xe9, 0x7e, 0xdd,
	0xaf, 0xa7, 0x5e, 0x83, 0x27, 0x95, 0xd0, 0x82, 0xb3, 0x34, 0x9f, 0xe2, 0x8b, 0x58, 0x47, 0xc2,
	0xff, 0x6c, 0x02, 0x3c, 0xc6, 0x8a, 0x87, 0x3a, 0xde, 0xf1, 0x94, 0xd8, 0x60, 0x84, 0x52, 0xd2,
	0xc6, 0xb8, 0x31, 0xe9, 0x92, 0x0b, 0xe8, 0x15, 0x20, 0x50, 0x3c, 0x09, 0x44, 0xa6, 0x83, 0x55,
	0xac, 0x02, 0xc9, 0xf4, 0x9a, 0x36, 0x8b, 0xa9, 0x45, 0x5c, 0x68, 0x87, 0xe9, 0x9a, 0x29, 0x49,
	0x0d, 0x54, 0x5f, 0xc1, 0xa0, 0xc2, 0x75, 0x83, 0x89, 0x06, 0x02, 0x20, 0x96, 0xaf, 0x55, 0x52,
	0x48, 0x5b, 0x68, 0xba, 0x86, 0xd1, 0x89, 0xab, 0x1b, 0xdb, 0x87, 0x24, 0x99, 0xeb, 0xb5, 0xd8,
	0xd2, 0xce, 0x21, 0xa9, 0xc2, 0x75, 0x43, 0x17, 0x0d, 0x0e, 0x98, 0x2a, 0x5b, 0xe6, 0xd4, 0x42,
	0xf9, 0x25, 0xf4, 0x4b, 0x54, 0x17, 0x03, 0x8a, 0x01, 0x9a, 0x91, 0xa0, 0x36, 0x4a, 0x47, 0xf0,
	0x37, 0x12, 0x75, 0xa1, 0x83, 0xc2, 0x7f, 0xe0, 0x16, 0xc3, 0x78, 0x23, 0x85, 0xd2, 0x15, 0xef,
	0x1e, 0xd2, 0x22, 0x25, 0x43, 0xea, 0xe1, 0x8a, 0x1e, 0x38, 0x25, 0x0a, 0x22, 0xa6, 0xf9, 0x3b,
	0xcb, 0xe9, 0x1f, 0x64, 0x07, 0xe0, 0xf1, 0x8f, 0x30, 0xc9, 0x56, 0xbc, 0x74, 0x6a, 0xae, 0xb6,
	0x94, 0x8c, 0x8d, 0x89, 0xe5, 0xf7, 0xc1, 0x9c, 0xab, 0x28, 0x25, 0xbf, 0xa1, 0xb5, 0x63, 0x49,
	0xc6, 0x8b, 0xd3, 0x97, 0xf4, 0x13, 0xb8, 0x0f, 0x62, 0x23, 0xe3, 0x84, 0xbf, 0xf0, 0xb7, 0x8c,
	0xa7, 0xba, 0x6c, 0x46, 0x33, 0x85, 0xcd, 0x38, 0xe4, 0x06, 0x60, 0x75, 0x2c, 0x0d, 0xfb, 0xb0,
	0x67, 0xfd, 0xe9, 0xa9, 0xe6, 0x53, 0xa3, 0xfe, 0x1d, 0x78, 0xc7, 0x4d, 0xa9, 0x14, 0xdb, 0x94,
	0x17, 0x07, 0x31, 0x59, 0x91, 0x89, 0x51, 0xf6, 0xcc, 0x3b, 0xf3, 0xe1, 0x57, 0xf6, 0x49, 0xe5,
	0x56, 0x67, 0xb6, 0x00, 0x63, 0xfe, 0xbc, 0x20, 0xf7, 0xd0, 0xd9, 0x6f, 0x21, 0xff, 0xcf, 0xf4,
	0xdf, 0xff, 0x38, 0x1c, 0xfe, 0x34, 0xaa, 0x42, 0xfd, 0x5f, 0xcb, 0x36, 0x0e, 0x6f, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x53, 0x97, 0xac, 0x65, 0x8f, 0x02, 0x00, 0x00,
}
