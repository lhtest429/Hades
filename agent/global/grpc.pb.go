// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package global

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RawData struct {
	IntranetIPv4         []string  `protobuf:"bytes,1,rep,name=IntranetIPv4,proto3" json:"IntranetIPv4,omitempty"`
	IntranetIPv6         []string  `protobuf:"bytes,2,rep,name=IntranetIPv6,proto3" json:"IntranetIPv6,omitempty"`
	Hostname             string    `protobuf:"bytes,3,opt,name=Hostname,proto3" json:"Hostname,omitempty"`
	AgentID              string    `protobuf:"bytes,4,opt,name=AgentID,proto3" json:"AgentID,omitempty"`
	Timestamp            int64     `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Version              string    `protobuf:"bytes,6,opt,name=Version,proto3" json:"Version,omitempty"`
	Pkg                  []*Record `protobuf:"bytes,7,rep,name=Pkg,proto3" json:"Pkg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RawData) Reset()         { *m = RawData{} }
func (m *RawData) String() string { return proto.CompactTextString(m) }
func (*RawData) ProtoMessage()    {}
func (*RawData) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *RawData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawData.Unmarshal(m, b)
}
func (m *RawData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawData.Marshal(b, m, deterministic)
}
func (m *RawData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawData.Merge(m, src)
}
func (m *RawData) XXX_Size() int {
	return xxx_messageInfo_RawData.Size(m)
}
func (m *RawData) XXX_DiscardUnknown() {
	xxx_messageInfo_RawData.DiscardUnknown(m)
}

var xxx_messageInfo_RawData proto.InternalMessageInfo

func (m *RawData) GetIntranetIPv4() []string {
	if m != nil {
		return m.IntranetIPv4
	}
	return nil
}

func (m *RawData) GetIntranetIPv6() []string {
	if m != nil {
		return m.IntranetIPv6
	}
	return nil
}

func (m *RawData) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *RawData) GetAgentID() string {
	if m != nil {
		return m.AgentID
	}
	return ""
}

func (m *RawData) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *RawData) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RawData) GetPkg() []*Record {
	if m != nil {
		return m.Pkg
	}
	return nil
}

type Record struct {
	Message              map[string]string `protobuf:"bytes,1,rep,name=message,proto3" json:"message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetMessage() map[string]string {
	if m != nil {
		return m.Message
	}
	return nil
}

type Command struct {
	Message              map[string]string `protobuf:"bytes,1,rep,name=message,proto3" json:"message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{2}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetMessage() map[string]string {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*RawData)(nil), "proto.RawData")
	proto.RegisterType((*Record)(nil), "proto.Record")
	proto.RegisterMapType((map[string]string)(nil), "proto.Record.MessageEntry")
	proto.RegisterType((*Command)(nil), "proto.Command")
	proto.RegisterMapType((map[string]string)(nil), "proto.Command.MessageEntry")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x51, 0xcd, 0x4e, 0x32, 0x31,
	0x14, 0xfd, 0xca, 0x7c, 0xcc, 0xc0, 0x05, 0x89, 0x69, 0x5c, 0x34, 0xa3, 0x89, 0x93, 0x59, 0xcd,
	0x8a, 0x10, 0x44, 0x62, 0x88, 0x1b, 0x15, 0x13, 0x59, 0x98, 0x90, 0x86, 0xb8, 0x70, 0x57, 0xa0,
	0x4e, 0x08, 0xb4, 0x25, 0x6d, 0xc5, 0x60, 0x7c, 0x58, 0x1f, 0xc5, 0xd0, 0x82, 0x32, 0xb8, 0x75,
	0xd5, 0x7b, 0x7e, 0x6e, 0xce, 0x3d, 0x29, 0x40, 0xae, 0x97, 0x93, 0xe6, 0x52, 0x2b, 0xab, 0x70,
	0xd9, 0x3d, 0xe9, 0x27, 0x82, 0x88, 0xb2, 0xb7, 0x3e, 0xb3, 0x0c, 0xa7, 0x50, 0x1f, 0x48, 0xab,
	0x99, 0xe4, 0x76, 0x30, 0x5c, 0x75, 0x08, 0x4a, 0x82, 0xac, 0x4a, 0x0b, 0xdc, 0x81, 0xa7, 0x4b,
	0x4a, 0xbf, 0x3c, 0x5d, 0x1c, 0x43, 0xe5, 0x41, 0x19, 0x2b, 0x99, 0xe0, 0x24, 0x48, 0x50, 0x56,
	0xa5, 0xdf, 0x18, 0x13, 0x88, 0x6e, 0x72, 0x2e, 0xed, 0xa0, 0x4f, 0xfe, 0x3b, 0x69, 0x07, 0xf1,
	0x19, 0x54, 0x47, 0x33, 0xc1, 0x8d, 0x65, 0x62, 0x49, 0xca, 0x09, 0xca, 0x02, 0xfa, 0x43, 0x6c,
	0xf6, 0x9e, 0xb8, 0x36, 0x33, 0x25, 0x49, 0xe8, 0xf7, 0xb6, 0x10, 0x9f, 0x43, 0x30, 0x9c, 0xe7,
	0x24, 0x4a, 0x82, 0xac, 0xd6, 0x3e, 0xf2, 0xed, 0x9a, 0x94, 0x4f, 0x94, 0x9e, 0xd2, 0x8d, 0x92,
	0xbe, 0x43, 0xe8, 0x21, 0xee, 0x40, 0x24, 0xb8, 0x31, 0x2c, 0xe7, 0xae, 0x5b, 0xad, 0x1d, 0x17,
	0xec, 0xcd, 0x47, 0x2f, 0xde, 0x4b, 0xab, 0xd7, 0x74, 0x67, 0x8d, 0x7b, 0x50, 0xdf, 0x17, 0xf0,
	0x31, 0x04, 0x73, 0xbe, 0x26, 0xc8, 0x9d, 0xb1, 0x19, 0xf1, 0x09, 0x94, 0x57, 0x6c, 0xf1, 0xca,
	0x49, 0xc9, 0x71, 0x1e, 0xf4, 0x4a, 0x57, 0x28, 0xfd, 0x80, 0xe8, 0x4e, 0x09, 0xc1, 0xe4, 0x14,
	0x5f, 0x1e, 0x86, 0x9f, 0x6e, 0xc3, 0xb7, 0x86, 0xbf, 0x4f, 0x6f, 0x5f, 0x43, 0x65, 0xa4, 0x99,
	0x34, 0x2f, 0x5c, 0xe3, 0xd6, 0xde, 0xdc, 0xd8, 0xd5, 0xf6, 0x1f, 0x1f, 0x37, 0x8a, 0x97, 0xa4,
	0xff, 0x32, 0xd4, 0x42, 0xb7, 0x95, 0xe7, 0x30, 0x5f, 0xa8, 0x31, 0x5b, 0x8c, 0x43, 0x27, 0x5f,
	0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x31, 0x41, 0xdc, 0x88, 0x40, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransferClient is the client API for Transfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransferClient interface {
	Transfer(ctx context.Context, opts ...grpc.CallOption) (Transfer_TransferClient, error)
}

type transferClient struct {
	cc *grpc.ClientConn
}

func NewTransferClient(cc *grpc.ClientConn) TransferClient {
	return &transferClient{cc}
}

func (c *transferClient) Transfer(ctx context.Context, opts ...grpc.CallOption) (Transfer_TransferClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Transfer_serviceDesc.Streams[0], "/proto.Transfer/Transfer", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferTransferClient{stream}
	return x, nil
}

type Transfer_TransferClient interface {
	Send(*RawData) error
	Recv() (*Command, error)
	grpc.ClientStream
}

type transferTransferClient struct {
	grpc.ClientStream
}

func (x *transferTransferClient) Send(m *RawData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transferTransferClient) Recv() (*Command, error) {
	m := new(Command)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransferServer is the server API for Transfer service.
type TransferServer interface {
	Transfer(Transfer_TransferServer) error
}

// UnimplementedTransferServer can be embedded to have forward compatible implementations.
type UnimplementedTransferServer struct {
}

func (*UnimplementedTransferServer) Transfer(srv Transfer_TransferServer) error {
	return status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}

func RegisterTransferServer(s *grpc.Server, srv TransferServer) {
	s.RegisterService(&_Transfer_serviceDesc, srv)
}

func _Transfer_Transfer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransferServer).Transfer(&transferTransferServer{stream})
}

type Transfer_TransferServer interface {
	Send(*Command) error
	Recv() (*RawData, error)
	grpc.ServerStream
}

type transferTransferServer struct {
	grpc.ServerStream
}

func (x *transferTransferServer) Send(m *Command) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transferTransferServer) Recv() (*RawData, error) {
	m := new(RawData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Transfer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Transfer",
	HandlerType: (*TransferServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Transfer",
			Handler:       _Transfer_Transfer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}
