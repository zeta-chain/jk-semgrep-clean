// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/lightclient/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgEnableHeaderVerification struct {
	Creator     string  `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ChainIdList []int64 `protobuf:"varint,2,rep,packed,name=chain_id_list,json=chainIdList,proto3" json:"chain_id_list,omitempty"`
}

func (m *MsgEnableHeaderVerification) Reset()         { *m = MsgEnableHeaderVerification{} }
func (m *MsgEnableHeaderVerification) String() string { return proto.CompactTextString(m) }
func (*MsgEnableHeaderVerification) ProtoMessage()    {}
func (*MsgEnableHeaderVerification) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fec9f445d2bf2d1, []int{0}
}
func (m *MsgEnableHeaderVerification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEnableHeaderVerification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEnableHeaderVerification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEnableHeaderVerification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEnableHeaderVerification.Merge(m, src)
}
func (m *MsgEnableHeaderVerification) XXX_Size() int {
	return m.Size()
}
func (m *MsgEnableHeaderVerification) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEnableHeaderVerification.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEnableHeaderVerification proto.InternalMessageInfo

func (m *MsgEnableHeaderVerification) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgEnableHeaderVerification) GetChainIdList() []int64 {
	if m != nil {
		return m.ChainIdList
	}
	return nil
}

type MsgEnableHeaderVerificationResponse struct {
}

func (m *MsgEnableHeaderVerificationResponse) Reset()         { *m = MsgEnableHeaderVerificationResponse{} }
func (m *MsgEnableHeaderVerificationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEnableHeaderVerificationResponse) ProtoMessage()    {}
func (*MsgEnableHeaderVerificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fec9f445d2bf2d1, []int{1}
}
func (m *MsgEnableHeaderVerificationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEnableHeaderVerificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEnableHeaderVerificationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEnableHeaderVerificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEnableHeaderVerificationResponse.Merge(m, src)
}
func (m *MsgEnableHeaderVerificationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgEnableHeaderVerificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEnableHeaderVerificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEnableHeaderVerificationResponse proto.InternalMessageInfo

type MsgDisableHeaderVerification struct {
	Creator     string  `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ChainIdList []int64 `protobuf:"varint,2,rep,packed,name=chain_id_list,json=chainIdList,proto3" json:"chain_id_list,omitempty"`
}

func (m *MsgDisableHeaderVerification) Reset()         { *m = MsgDisableHeaderVerification{} }
func (m *MsgDisableHeaderVerification) String() string { return proto.CompactTextString(m) }
func (*MsgDisableHeaderVerification) ProtoMessage()    {}
func (*MsgDisableHeaderVerification) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fec9f445d2bf2d1, []int{2}
}
func (m *MsgDisableHeaderVerification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDisableHeaderVerification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDisableHeaderVerification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDisableHeaderVerification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDisableHeaderVerification.Merge(m, src)
}
func (m *MsgDisableHeaderVerification) XXX_Size() int {
	return m.Size()
}
func (m *MsgDisableHeaderVerification) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDisableHeaderVerification.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDisableHeaderVerification proto.InternalMessageInfo

func (m *MsgDisableHeaderVerification) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDisableHeaderVerification) GetChainIdList() []int64 {
	if m != nil {
		return m.ChainIdList
	}
	return nil
}

type MsgDisableHeaderVerificationResponse struct {
}

func (m *MsgDisableHeaderVerificationResponse) Reset()         { *m = MsgDisableHeaderVerificationResponse{} }
func (m *MsgDisableHeaderVerificationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDisableHeaderVerificationResponse) ProtoMessage()    {}
func (*MsgDisableHeaderVerificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fec9f445d2bf2d1, []int{3}
}
func (m *MsgDisableHeaderVerificationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDisableHeaderVerificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDisableHeaderVerificationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDisableHeaderVerificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDisableHeaderVerificationResponse.Merge(m, src)
}
func (m *MsgDisableHeaderVerificationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDisableHeaderVerificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDisableHeaderVerificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDisableHeaderVerificationResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgEnableHeaderVerification)(nil), "zetachain.zetacore.lightclient.MsgEnableHeaderVerification")
	proto.RegisterType((*MsgEnableHeaderVerificationResponse)(nil), "zetachain.zetacore.lightclient.MsgEnableHeaderVerificationResponse")
	proto.RegisterType((*MsgDisableHeaderVerification)(nil), "zetachain.zetacore.lightclient.MsgDisableHeaderVerification")
	proto.RegisterType((*MsgDisableHeaderVerificationResponse)(nil), "zetachain.zetacore.lightclient.MsgDisableHeaderVerificationResponse")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/lightclient/tx.proto", fileDescriptor_6fec9f445d2bf2d1)
}

var fileDescriptor_6fec9f445d2bf2d1 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xaf, 0x4a, 0x2d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb3, 0xf2, 0x8b, 0x52, 0xf5, 0x73, 0x32, 0xd3, 0x33,
	0x4a, 0x92, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0xf4, 0x4b, 0x2a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xe4, 0xe0, 0x0a, 0xf5, 0x60, 0x0a, 0xf5, 0x90, 0x14, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7,
	0x83, 0x95, 0xea, 0x83, 0x58, 0x10, 0x5d, 0x52, 0x76, 0x04, 0x8c, 0x4f, 0xca, 0xc9, 0x4f, 0xce,
	0x8e, 0xcf, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x8a, 0x2f, 0x4b, 0x2d, 0xca, 0x4c, 0xcb, 0x4c, 0x4e,
	0x2c, 0xc9, 0xcc, 0xcf, 0x83, 0xe8, 0x57, 0x8a, 0xe6, 0x92, 0xf6, 0x2d, 0x4e, 0x77, 0xcd, 0x4b,
	0x4c, 0xca, 0x49, 0xf5, 0x00, 0xab, 0x0a, 0x43, 0x52, 0x24, 0x24, 0xc1, 0xc5, 0x9e, 0x5c, 0x94,
	0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0x29, 0x71,
	0xf1, 0x82, 0xad, 0x8d, 0xcf, 0x4c, 0x89, 0xcf, 0xc9, 0x2c, 0x2e, 0x91, 0x60, 0x52, 0x60, 0xd6,
	0x60, 0x0e, 0xe2, 0x06, 0x0b, 0x7a, 0xa6, 0xf8, 0x64, 0x16, 0x97, 0x28, 0xa9, 0x72, 0x29, 0xe3,
	0x31, 0x3c, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x29, 0x86, 0x4b, 0xc6, 0xb7, 0x38,
	0xdd, 0x25, 0xb3, 0x98, 0x26, 0x8e, 0x50, 0xe3, 0x52, 0xc1, 0x67, 0x3a, 0xcc, 0x15, 0x46, 0xc7,
	0x98, 0xb8, 0x98, 0x7d, 0x8b, 0xd3, 0x85, 0xe6, 0x30, 0x72, 0x49, 0xe0, 0x0c, 0x0f, 0x6b, 0x3d,
	0xfc, 0xb1, 0xa4, 0x87, 0xc7, 0xbf, 0x52, 0xce, 0x14, 0x68, 0x86, 0x39, 0x53, 0x68, 0x3e, 0x23,
	0x97, 0x24, 0xee, 0xa0, 0xb2, 0x21, 0xc2, 0x0a, 0x9c, 0xba, 0xa5, 0x5c, 0x28, 0xd1, 0x0d, 0x73,
	0xa1, 0x93, 0xc7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0xa5, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0x82, 0x53, 0xab, 0x2e, 0x24, 0xe1, 0xe6, 0xe5,
	0xa7, 0xa4, 0xea, 0x57, 0xa0, 0xe6, 0x8a, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x1a, 0x35,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xd2, 0xa9, 0x28, 0x44, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	EnableHeaderVerification(ctx context.Context, in *MsgEnableHeaderVerification, opts ...grpc.CallOption) (*MsgEnableHeaderVerificationResponse, error)
	DisableHeaderVerification(ctx context.Context, in *MsgDisableHeaderVerification, opts ...grpc.CallOption) (*MsgDisableHeaderVerificationResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) EnableHeaderVerification(ctx context.Context, in *MsgEnableHeaderVerification, opts ...grpc.CallOption) (*MsgEnableHeaderVerificationResponse, error) {
	out := new(MsgEnableHeaderVerificationResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.lightclient.Msg/EnableHeaderVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DisableHeaderVerification(ctx context.Context, in *MsgDisableHeaderVerification, opts ...grpc.CallOption) (*MsgDisableHeaderVerificationResponse, error) {
	out := new(MsgDisableHeaderVerificationResponse)
	err := c.cc.Invoke(ctx, "/zetachain.zetacore.lightclient.Msg/DisableHeaderVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	EnableHeaderVerification(context.Context, *MsgEnableHeaderVerification) (*MsgEnableHeaderVerificationResponse, error)
	DisableHeaderVerification(context.Context, *MsgDisableHeaderVerification) (*MsgDisableHeaderVerificationResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) EnableHeaderVerification(ctx context.Context, req *MsgEnableHeaderVerification) (*MsgEnableHeaderVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableHeaderVerification not implemented")
}
func (*UnimplementedMsgServer) DisableHeaderVerification(ctx context.Context, req *MsgDisableHeaderVerification) (*MsgDisableHeaderVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableHeaderVerification not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_EnableHeaderVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEnableHeaderVerification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EnableHeaderVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.lightclient.Msg/EnableHeaderVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EnableHeaderVerification(ctx, req.(*MsgEnableHeaderVerification))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DisableHeaderVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDisableHeaderVerification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DisableHeaderVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zetachain.zetacore.lightclient.Msg/DisableHeaderVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DisableHeaderVerification(ctx, req.(*MsgDisableHeaderVerification))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zetachain.zetacore.lightclient.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnableHeaderVerification",
			Handler:    _Msg_EnableHeaderVerification_Handler,
		},
		{
			MethodName: "DisableHeaderVerification",
			Handler:    _Msg_DisableHeaderVerification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zetachain/zetacore/lightclient/tx.proto",
}

func (m *MsgEnableHeaderVerification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEnableHeaderVerification) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEnableHeaderVerification) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainIdList) > 0 {
		dAtA2 := make([]byte, len(m.ChainIdList)*10)
		var j1 int
		for _, num1 := range m.ChainIdList {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTx(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgEnableHeaderVerificationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEnableHeaderVerificationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEnableHeaderVerificationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDisableHeaderVerification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDisableHeaderVerification) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDisableHeaderVerification) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainIdList) > 0 {
		dAtA4 := make([]byte, len(m.ChainIdList)*10)
		var j3 int
		for _, num1 := range m.ChainIdList {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintTx(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDisableHeaderVerificationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDisableHeaderVerificationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDisableHeaderVerificationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgEnableHeaderVerification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ChainIdList) > 0 {
		l = 0
		for _, e := range m.ChainIdList {
			l += sovTx(uint64(e))
		}
		n += 1 + sovTx(uint64(l)) + l
	}
	return n
}

func (m *MsgEnableHeaderVerificationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDisableHeaderVerification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.ChainIdList) > 0 {
		l = 0
		for _, e := range m.ChainIdList {
			l += sovTx(uint64(e))
		}
		n += 1 + sovTx(uint64(l)) + l
	}
	return n
}

func (m *MsgDisableHeaderVerificationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgEnableHeaderVerification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgEnableHeaderVerification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEnableHeaderVerification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ChainIdList = append(m.ChainIdList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ChainIdList) == 0 {
					m.ChainIdList = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ChainIdList = append(m.ChainIdList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainIdList", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgEnableHeaderVerificationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgEnableHeaderVerificationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEnableHeaderVerificationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDisableHeaderVerification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDisableHeaderVerification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDisableHeaderVerification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ChainIdList = append(m.ChainIdList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ChainIdList) == 0 {
					m.ChainIdList = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ChainIdList = append(m.ChainIdList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainIdList", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDisableHeaderVerificationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDisableHeaderVerificationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDisableHeaderVerificationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
