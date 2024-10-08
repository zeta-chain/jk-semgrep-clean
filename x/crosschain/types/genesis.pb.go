// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/crosschain/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the crosschain module's genesis state.
type GenesisState struct {
	OutboundTrackerList   []OutboundTracker   `protobuf:"bytes,2,rep,name=outboundTrackerList,proto3" json:"outboundTrackerList"`
	GasPriceList          []*GasPrice         `protobuf:"bytes,5,rep,name=gasPriceList,proto3" json:"gasPriceList,omitempty"`
	CrossChainTxs         []*CrossChainTx     `protobuf:"bytes,7,rep,name=CrossChainTxs,proto3" json:"CrossChainTxs,omitempty"`
	LastBlockHeightList   []*LastBlockHeight  `protobuf:"bytes,8,rep,name=lastBlockHeightList,proto3" json:"lastBlockHeightList,omitempty"`
	InboundHashToCctxList []InboundHashToCctx `protobuf:"bytes,9,rep,name=inboundHashToCctxList,proto3" json:"inboundHashToCctxList"`
	InboundTrackerList    []InboundTracker    `protobuf:"bytes,11,rep,name=inbound_tracker_list,json=inboundTrackerList,proto3" json:"inbound_tracker_list"`
	ZetaAccounting        ZetaAccounting      `protobuf:"bytes,12,opt,name=zeta_accounting,json=zetaAccounting,proto3" json:"zeta_accounting"`
	FinalizedInbounds     []string            `protobuf:"bytes,16,rep,name=FinalizedInbounds,proto3" json:"FinalizedInbounds,omitempty"`
	RateLimiterFlags      RateLimiterFlags    `protobuf:"bytes,17,opt,name=rate_limiter_flags,json=rateLimiterFlags,proto3" json:"rate_limiter_flags"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_547615497292ea23, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetOutboundTrackerList() []OutboundTracker {
	if m != nil {
		return m.OutboundTrackerList
	}
	return nil
}

func (m *GenesisState) GetGasPriceList() []*GasPrice {
	if m != nil {
		return m.GasPriceList
	}
	return nil
}

func (m *GenesisState) GetCrossChainTxs() []*CrossChainTx {
	if m != nil {
		return m.CrossChainTxs
	}
	return nil
}

func (m *GenesisState) GetLastBlockHeightList() []*LastBlockHeight {
	if m != nil {
		return m.LastBlockHeightList
	}
	return nil
}

func (m *GenesisState) GetInboundHashToCctxList() []InboundHashToCctx {
	if m != nil {
		return m.InboundHashToCctxList
	}
	return nil
}

func (m *GenesisState) GetInboundTrackerList() []InboundTracker {
	if m != nil {
		return m.InboundTrackerList
	}
	return nil
}

func (m *GenesisState) GetZetaAccounting() ZetaAccounting {
	if m != nil {
		return m.ZetaAccounting
	}
	return ZetaAccounting{}
}

func (m *GenesisState) GetFinalizedInbounds() []string {
	if m != nil {
		return m.FinalizedInbounds
	}
	return nil
}

func (m *GenesisState) GetRateLimiterFlags() RateLimiterFlags {
	if m != nil {
		return m.RateLimiterFlags
	}
	return RateLimiterFlags{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "zetachain.zetacore.crosschain.GenesisState")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/crosschain/genesis.proto", fileDescriptor_547615497292ea23)
}

var fileDescriptor_547615497292ea23 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x5b, 0x06, 0x83, 0x79, 0x05, 0x36, 0x6f, 0x48, 0x51, 0x25, 0x42, 0xc5, 0x85, 0x49,
	0xa3, 0x09, 0xda, 0x00, 0x71, 0x65, 0x95, 0xd6, 0x21, 0x2a, 0x01, 0xa1, 0xa7, 0x09, 0xc9, 0xb8,
	0xae, 0x97, 0x58, 0xcb, 0xe2, 0x2a, 0x7e, 0x95, 0x4a, 0x3f, 0x05, 0x9f, 0x88, 0xf3, 0x8e, 0x3b,
	0x72, 0x42, 0xa8, 0xfd, 0x22, 0xc8, 0x8e, 0x57, 0x9a, 0xb6, 0x4a, 0x7a, 0x7b, 0x7a, 0x79, 0xff,
	0xff, 0xef, 0x29, 0x7f, 0x3f, 0x74, 0x38, 0xe6, 0x40, 0x59, 0x44, 0x45, 0xe2, 0x9b, 0x4a, 0xa6,
	0xdc, 0x67, 0xa9, 0x54, 0x2a, 0xeb, 0x85, 0x3c, 0xe1, 0x4a, 0x28, 0x6f, 0x90, 0x4a, 0x90, 0xf8,
	0xe9, 0x6c, 0xd8, 0xbb, 0x1d, 0xf6, 0xfe, 0x0f, 0xd7, 0x8f, 0x8a, 0xbd, 0x4c, 0x49, 0x4c, 0x4d,
	0x60, 0x94, 0x59, 0xd6, 0x9b, 0x25, 0x7c, 0xaa, 0xc8, 0x20, 0x15, 0x8c, 0xdb, 0xf1, 0x77, 0xc5,
	0xe3, 0x22, 0xe9, 0xc9, 0x61, 0xd2, 0x27, 0x11, 0x55, 0x11, 0x01, 0x49, 0x18, 0x9b, 0x81, 0x8e,
	0xd7, 0x53, 0x42, 0x4a, 0xd9, 0x25, 0x4f, 0xad, 0xe8, 0x4d, 0xb1, 0x28, 0xa6, 0x0a, 0x48, 0x2f,
	0x96, 0xec, 0x92, 0x44, 0x5c, 0x84, 0x11, 0x58, 0xd9, 0xeb, 0x62, 0x99, 0x1c, 0xc2, 0x2a, 0xd8,
	0xdb, 0x62, 0x55, 0x4a, 0x81, 0x93, 0x58, 0x5c, 0x09, 0xe0, 0x29, 0xb9, 0x88, 0x69, 0x68, 0x53,
	0xa9, 0xef, 0x87, 0x32, 0x94, 0xa6, 0xf4, 0x75, 0x95, 0x75, 0x9f, 0xff, 0xda, 0x44, 0xb5, 0x76,
	0x96, 0xde, 0x57, 0xa0, 0xc0, 0xf1, 0x05, 0xda, 0xbb, 0x05, 0x77, 0x33, 0x6e, 0x47, 0x28, 0x70,
	0xee, 0x34, 0x36, 0x0e, 0xb6, 0x8f, 0x3c, 0xaf, 0x30, 0x5a, 0xef, 0x53, 0x5e, 0x79, 0x72, 0xf7,
	0xfa, 0xcf, 0xb3, 0x4a, 0xb0, 0xca, 0x10, 0x7f, 0x44, 0xb5, 0x90, 0xaa, 0xcf, 0x3a, 0x34, 0x03,
	0xb8, 0x67, 0x00, 0x2f, 0x4a, 0x00, 0x6d, 0x2b, 0x09, 0x72, 0x62, 0xfc, 0x05, 0x3d, 0x6c, 0xe9,
	0xa1, 0x96, 0x1e, 0xea, 0x8e, 0x94, 0x73, 0xdf, 0xb8, 0x1d, 0x96, 0xb8, 0xcd, 0x6b, 0x82, 0xbc,
	0x03, 0xfe, 0x8e, 0xf6, 0x74, 0x6e, 0x27, 0x3a, 0xb6, 0x33, 0x93, 0x9a, 0x59, 0xf3, 0xc1, 0x5a,
	0xff, 0xa1, 0x93, 0x57, 0x06, 0xab, 0xac, 0x70, 0x8c, 0x9e, 0xd8, 0xe7, 0x74, 0x46, 0x55, 0xd4,
	0x95, 0x2d, 0x06, 0x23, 0xc3, 0xd8, 0x32, 0x8c, 0x57, 0x25, 0x8c, 0x0f, 0x8b, 0x5a, 0xfb, 0xb7,
	0x57, 0x9b, 0x62, 0x8e, 0xf6, 0x17, 0x1e, 0x2f, 0x89, 0x35, 0x6c, 0xdb, 0xc0, 0x9a, 0xeb, 0xc1,
	0xf2, 0xb9, 0x62, 0x91, 0x2c, 0xc5, 0xfa, 0x0d, 0x3d, 0xd6, 0x7a, 0x42, 0x19, 0x93, 0xc3, 0x04,
	0x44, 0x12, 0x3a, 0xb5, 0x46, 0x75, 0x0d, 0xc2, 0x39, 0x07, 0xfa, 0x7e, 0x26, 0xb2, 0x84, 0x47,
	0xe3, 0x5c, 0x17, 0xbf, 0x44, 0xbb, 0xa7, 0x22, 0xa1, 0xb1, 0x18, 0xf3, 0xbe, 0x5d, 0x49, 0x39,
	0x3b, 0x8d, 0x8d, 0x83, 0xad, 0x60, 0xf9, 0x03, 0x66, 0x08, 0x2f, 0x5f, 0x83, 0xb3, 0x6b, 0xd6,
	0xf1, 0x4b, 0xd6, 0x09, 0x28, 0xf0, 0x4e, 0xa6, 0x3b, 0xd5, 0x32, 0xbb, 0xd0, 0x4e, 0xba, 0xd8,
	0x6f, 0x5f, 0x4f, 0xdc, 0xea, 0xcd, 0xc4, 0xad, 0xfe, 0x9d, 0xb8, 0xd5, 0x9f, 0x53, 0xb7, 0x72,
	0x33, 0x75, 0x2b, 0xbf, 0xa7, 0x6e, 0xe5, 0xbc, 0x19, 0x0a, 0x88, 0x86, 0x3d, 0x8f, 0xc9, 0x2b,
	0x73, 0xa9, 0xcd, 0xec, 0x40, 0x13, 0xd9, 0xe7, 0xfe, 0x68, 0xfe, 0x64, 0xe1, 0xc7, 0x80, 0xab,
	0xde, 0xa6, 0x39, 0xc8, 0xe3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xde, 0x67, 0x86, 0x4a, 0x6b,
	0x05, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RateLimiterFlags.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x8a
	if len(m.FinalizedInbounds) > 0 {
		for iNdEx := len(m.FinalizedInbounds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FinalizedInbounds[iNdEx])
			copy(dAtA[i:], m.FinalizedInbounds[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.FinalizedInbounds[iNdEx])))
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	{
		size, err := m.ZetaAccounting.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if len(m.InboundTrackerList) > 0 {
		for iNdEx := len(m.InboundTrackerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InboundTrackerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.InboundHashToCctxList) > 0 {
		for iNdEx := len(m.InboundHashToCctxList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InboundHashToCctxList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.LastBlockHeightList) > 0 {
		for iNdEx := len(m.LastBlockHeightList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LastBlockHeightList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CrossChainTxs) > 0 {
		for iNdEx := len(m.CrossChainTxs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CrossChainTxs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.GasPriceList) > 0 {
		for iNdEx := len(m.GasPriceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasPriceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.OutboundTrackerList) > 0 {
		for iNdEx := len(m.OutboundTrackerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OutboundTrackerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.OutboundTrackerList) > 0 {
		for _, e := range m.OutboundTrackerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GasPriceList) > 0 {
		for _, e := range m.GasPriceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CrossChainTxs) > 0 {
		for _, e := range m.CrossChainTxs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LastBlockHeightList) > 0 {
		for _, e := range m.LastBlockHeightList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.InboundHashToCctxList) > 0 {
		for _, e := range m.InboundHashToCctxList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.InboundTrackerList) > 0 {
		for _, e := range m.InboundTrackerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.ZetaAccounting.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FinalizedInbounds) > 0 {
		for _, s := range m.FinalizedInbounds {
			l = len(s)
			n += 2 + l + sovGenesis(uint64(l))
		}
	}
	l = m.RateLimiterFlags.Size()
	n += 2 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundTrackerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutboundTrackerList = append(m.OutboundTrackerList, OutboundTracker{})
			if err := m.OutboundTrackerList[len(m.OutboundTrackerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPriceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPriceList = append(m.GasPriceList, &GasPrice{})
			if err := m.GasPriceList[len(m.GasPriceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrossChainTxs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CrossChainTxs = append(m.CrossChainTxs, &CrossChainTx{})
			if err := m.CrossChainTxs[len(m.CrossChainTxs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockHeightList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastBlockHeightList = append(m.LastBlockHeightList, &LastBlockHeight{})
			if err := m.LastBlockHeightList[len(m.LastBlockHeightList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InboundHashToCctxList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InboundHashToCctxList = append(m.InboundHashToCctxList, InboundHashToCctx{})
			if err := m.InboundHashToCctxList[len(m.InboundHashToCctxList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InboundTrackerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InboundTrackerList = append(m.InboundTrackerList, InboundTracker{})
			if err := m.InboundTrackerList[len(m.InboundTrackerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZetaAccounting", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ZetaAccounting.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizedInbounds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalizedInbounds = append(m.FinalizedInbounds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RateLimiterFlags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RateLimiterFlags.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
