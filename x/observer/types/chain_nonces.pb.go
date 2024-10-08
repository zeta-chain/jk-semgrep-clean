// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/observer/chain_nonces.proto

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

type ChainNonces struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// deprecated(v19): index has been replaced by chain_id for unique identifier
	Index           string   `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"` // Deprecated: Do not use.
	ChainId         int64    `protobuf:"varint,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Nonce           uint64   `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Signers         []string `protobuf:"bytes,5,rep,name=signers,proto3" json:"signers,omitempty"`
	FinalizedHeight uint64   `protobuf:"varint,6,opt,name=finalizedHeight,proto3" json:"finalizedHeight,omitempty"`
}

func (m *ChainNonces) Reset()         { *m = ChainNonces{} }
func (m *ChainNonces) String() string { return proto.CompactTextString(m) }
func (*ChainNonces) ProtoMessage()    {}
func (*ChainNonces) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8bc11608907ed56, []int{0}
}
func (m *ChainNonces) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainNonces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainNonces.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainNonces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainNonces.Merge(m, src)
}
func (m *ChainNonces) XXX_Size() int {
	return m.Size()
}
func (m *ChainNonces) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainNonces.DiscardUnknown(m)
}

var xxx_messageInfo_ChainNonces proto.InternalMessageInfo

func (m *ChainNonces) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// Deprecated: Do not use.
func (m *ChainNonces) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ChainNonces) GetChainId() int64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *ChainNonces) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *ChainNonces) GetSigners() []string {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *ChainNonces) GetFinalizedHeight() uint64 {
	if m != nil {
		return m.FinalizedHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*ChainNonces)(nil), "zetachain.zetacore.observer.ChainNonces")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/observer/chain_nonces.proto", fileDescriptor_d8bc11608907ed56)
}

var fileDescriptor_d8bc11608907ed56 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x37, 0xdb, 0xed, 0xae, 0xc6, 0x83, 0x10, 0xf6, 0x10, 0x15, 0x42, 0xf1, 0x54, 0x10,
	0xdb, 0x83, 0x6f, 0xb0, 0x22, 0xe8, 0xc5, 0x43, 0x8f, 0x5e, 0xa4, 0x7f, 0xc6, 0x36, 0xa0, 0x99,
	0x25, 0x89, 0xb2, 0xee, 0x53, 0xf8, 0x34, 0x3e, 0x83, 0xc7, 0x3d, 0x7a, 0x94, 0xf6, 0x45, 0x24,
	0x89, 0x55, 0xf0, 0x36, 0xdf, 0xc7, 0xf7, 0xcd, 0x0c, 0x3f, 0x9a, 0x6d, 0xc1, 0x96, 0x75, 0x57,
	0x4a, 0x95, 0xfb, 0x09, 0x35, 0xe4, 0x58, 0x19, 0xd0, 0x2f, 0xa0, 0x73, 0xef, 0xdf, 0x2b, 0x54,
	0x35, 0x98, 0x6c, 0xad, 0xd1, 0x22, 0x3b, 0xf9, 0xcd, 0x67, 0x63, 0x3e, 0x1b, 0xf3, 0xc7, 0xcb,
	0x16, 0x5b, 0xf4, 0xb9, 0xdc, 0x4d, 0xa1, 0x72, 0xfa, 0x4e, 0xe8, 0xc1, 0xa5, 0x6b, 0xdc, 0xfa,
	0x45, 0x8c, 0xd3, 0x45, 0xad, 0xa1, 0xb4, 0xa8, 0x39, 0x49, 0x48, 0xba, 0x5f, 0x8c, 0x92, 0x71,
	0x1a, 0x4b, 0xd5, 0xc0, 0x86, 0x4f, 0x9d, 0xbf, 0x9a, 0x72, 0x52, 0x04, 0x83, 0x1d, 0xd1, 0xbd,
	0xf0, 0x8c, 0x6c, 0x78, 0x94, 0x90, 0x34, 0x2a, 0x16, 0x5e, 0xdf, 0x34, 0x6c, 0x49, 0x63, 0xff,
	0x21, 0x9f, 0x25, 0x24, 0x9d, 0x15, 0x41, 0xb8, 0x23, 0x46, 0xb6, 0x0a, 0xb4, 0xe1, 0x71, 0x12,
	0xb9, 0x23, 0x3f, 0x92, 0xa5, 0xf4, 0xf0, 0x41, 0xaa, 0xf2, 0x51, 0x6e, 0xa1, 0xb9, 0x06, 0xd9,
	0x76, 0x96, 0xcf, 0x7d, 0xf3, 0xbf, 0xbd, 0xba, 0xfa, 0xe8, 0x05, 0xd9, 0xf5, 0x82, 0x7c, 0xf5,
	0x82, 0xbc, 0x0d, 0x62, 0xb2, 0x1b, 0xc4, 0xe4, 0x73, 0x10, 0x93, 0xbb, 0xb3, 0x56, 0xda, 0xee,
	0xb9, 0xca, 0x6a, 0x7c, 0xf2, 0xd8, 0xce, 0x03, 0x41, 0x85, 0x0d, 0xe4, 0x9b, 0x3f, 0x7e, 0xf6,
	0x75, 0x0d, 0xa6, 0x9a, 0x7b, 0x0c, 0x17, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xa1, 0x0d,
	0xd7, 0x6b, 0x01, 0x00, 0x00,
}

func (m *ChainNonces) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainNonces) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainNonces) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalizedHeight != 0 {
		i = encodeVarintChainNonces(dAtA, i, uint64(m.FinalizedHeight))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintChainNonces(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Nonce != 0 {
		i = encodeVarintChainNonces(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x20
	}
	if m.ChainId != 0 {
		i = encodeVarintChainNonces(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintChainNonces(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintChainNonces(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintChainNonces(dAtA []byte, offset int, v uint64) int {
	offset -= sovChainNonces(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainNonces) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovChainNonces(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovChainNonces(uint64(l))
	}
	if m.ChainId != 0 {
		n += 1 + sovChainNonces(uint64(m.ChainId))
	}
	if m.Nonce != 0 {
		n += 1 + sovChainNonces(uint64(m.Nonce))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovChainNonces(uint64(l))
		}
	}
	if m.FinalizedHeight != 0 {
		n += 1 + sovChainNonces(uint64(m.FinalizedHeight))
	}
	return n
}

func sovChainNonces(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChainNonces(x uint64) (n int) {
	return sovChainNonces(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainNonces) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChainNonces
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
			return fmt.Errorf("proto: ChainNonces: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainNonces: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainNonces
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
				return ErrInvalidLengthChainNonces
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainNonces
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainNonces
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
				return ErrInvalidLengthChainNonces
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainNonces
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainNonces
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainNonces
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainNonces
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
				return ErrInvalidLengthChainNonces
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChainNonces
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalizedHeight", wireType)
			}
			m.FinalizedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChainNonces
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalizedHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChainNonces(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChainNonces
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
func skipChainNonces(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChainNonces
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
					return 0, ErrIntOverflowChainNonces
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
					return 0, ErrIntOverflowChainNonces
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
				return 0, ErrInvalidLengthChainNonces
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChainNonces
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChainNonces
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChainNonces        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChainNonces          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChainNonces = fmt.Errorf("proto: unexpected end of group")
)
