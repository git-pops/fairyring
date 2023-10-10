// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/keyshare/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	KeyExpiry                  uint64                                 `protobuf:"varint,1,opt,name=key_expiry,json=keyExpiry,proto3" json:"key_expiry,omitempty"`
	TrustedAddresses           []string                               `protobuf:"bytes,2,rep,name=trusted_addresses,json=trustedAddresses,proto3" json:"trusted_addresses,omitempty"`
	SlashFractionNoKeyshare    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=slash_fraction_no_keyshare,json=slashFractionNoKeyshare,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_no_keyshare"`
	SlashFractionWrongKeyshare github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=slash_fraction_wrong_keyshare,json=slashFractionWrongKeyshare,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_wrong_keyshare"`
	MinimumBonded              uint64                                 `protobuf:"varint,5,opt,name=minimum_bonded,json=minimumBonded,proto3" json:"minimum_bonded,omitempty"`
	MaxIdledBlock              uint64                                 `protobuf:"varint,6,opt,name=max_idled_block,json=maxIdledBlock,proto3" json:"max_idled_block,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ef7bd565425b36, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetKeyExpiry() uint64 {
	if m != nil {
		return m.KeyExpiry
	}
	return 0
}

func (m *Params) GetTrustedAddresses() []string {
	if m != nil {
		return m.TrustedAddresses
	}
	return nil
}

func (m *Params) GetMinimumBonded() uint64 {
	if m != nil {
		return m.MinimumBonded
	}
	return 0
}

func (m *Params) GetMaxIdledBlock() uint64 {
	if m != nil {
		return m.MaxIdledBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "fairyring.keyshare.Params")
}

func init() { proto.RegisterFile("fairyring/keyshare/params.proto", fileDescriptor_09ef7bd565425b36) }

var fileDescriptor_09ef7bd565425b36 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x4b, 0xfb, 0x30,
	0x1c, 0xc6, 0xdb, 0xdf, 0xf6, 0x1b, 0x2c, 0xfc, 0xfe, 0x68, 0x10, 0x2c, 0x85, 0xb5, 0x43, 0x70,
	0x0c, 0xc4, 0xf6, 0xa0, 0x27, 0x6f, 0x16, 0x15, 0x44, 0x10, 0xd9, 0x45, 0xf0, 0x12, 0xd2, 0x26,
	0xeb, 0x42, 0x97, 0xa6, 0x26, 0x1d, 0xb6, 0x2f, 0xc1, 0x9b, 0x47, 0x8f, 0xbe, 0x9c, 0x1d, 0x77,
	0x14, 0x0f, 0x43, 0xb6, 0x37, 0x22, 0x8d, 0xdd, 0x86, 0x1e, 0x3d, 0x25, 0x7c, 0xbe, 0x4f, 0x9e,
	0xe7, 0x4b, 0x1e, 0xe0, 0x0e, 0x31, 0x93, 0xa5, 0x64, 0x69, 0xec, 0x27, 0xb4, 0x54, 0x23, 0x2c,
	0xa9, 0x9f, 0x61, 0x89, 0xb9, 0xf2, 0x32, 0x29, 0x72, 0x01, 0xe1, 0x5a, 0xe0, 0xad, 0x04, 0xf6,
	0x4e, 0x2c, 0x62, 0xa1, 0xc7, 0x7e, 0x75, 0xfb, 0x54, 0xee, 0x3d, 0x36, 0x40, 0xeb, 0x46, 0x3f,
	0x85, 0x1d, 0x00, 0x12, 0x5a, 0x22, 0x5a, 0x64, 0x4c, 0x96, 0x96, 0xd9, 0x35, 0xfb, 0xcd, 0x41,
	0x3b, 0xa1, 0xe5, 0xb9, 0x06, 0xf0, 0x00, 0x6c, 0xe7, 0x72, 0xa2, 0x72, 0x4a, 0x10, 0x26, 0x44,
	0x52, 0xa5, 0xa8, 0xb2, 0x7e, 0x75, 0x1b, 0xfd, 0xf6, 0x60, 0xab, 0x1e, 0x9c, 0xae, 0x38, 0x4c,
	0x80, 0xad, 0xc6, 0x58, 0x8d, 0xd0, 0x50, 0xe2, 0x28, 0x67, 0x22, 0x45, 0xa9, 0x40, 0xab, 0x55,
	0xac, 0x46, 0xd7, 0xec, 0xff, 0x09, 0xbc, 0xe9, 0xdc, 0x35, 0xde, 0xe6, 0x6e, 0x2f, 0x66, 0xf9,
	0x68, 0x12, 0x7a, 0x91, 0xe0, 0x7e, 0x24, 0x14, 0x17, 0xaa, 0x3e, 0x0e, 0x15, 0x49, 0xfc, 0xbc,
	0xcc, 0xa8, 0xf2, 0xce, 0x68, 0x34, 0xd8, 0xd5, 0x8e, 0x17, 0xb5, 0xe1, 0xb5, 0xb8, 0xaa, 0xed,
	0xe0, 0x3d, 0xe8, 0x7c, 0x0b, 0x7b, 0x90, 0x22, 0x8d, 0x37, 0x79, 0xcd, 0x1f, 0xe5, 0xd9, 0x5f,
	0xf2, 0x6e, 0x2b, 0xcb, 0x75, 0xe4, 0x3e, 0xf8, 0xc7, 0x59, 0xca, 0xf8, 0x84, 0xa3, 0x50, 0xa4,
	0x84, 0x12, 0xeb, 0xb7, 0xfe, 0xaf, 0xbf, 0x35, 0x0d, 0x34, 0x84, 0x3d, 0xf0, 0x9f, 0xe3, 0x02,
	0x31, 0x32, 0xa6, 0x04, 0x85, 0x63, 0x11, 0x25, 0x56, 0xab, 0xd6, 0xe1, 0xe2, 0xb2, 0xa2, 0x41,
	0x05, 0x4f, 0x9a, 0xcf, 0x2f, 0xae, 0x11, 0x1c, 0x4f, 0x17, 0x8e, 0x39, 0x5b, 0x38, 0xe6, 0xfb,
	0xc2, 0x31, 0x9f, 0x96, 0x8e, 0x31, 0x5b, 0x3a, 0xc6, 0xeb, 0xd2, 0x31, 0xee, 0xec, 0x4d, 0xe1,
	0xc5, 0xa6, 0x72, 0xbd, 0x6a, 0xd8, 0xd2, 0x45, 0x1e, 0x7d, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6b,
	0xfc, 0xe2, 0xef, 0x15, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxIdledBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxIdledBlock))
		i--
		dAtA[i] = 0x30
	}
	if m.MinimumBonded != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinimumBonded))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.SlashFractionWrongKeyshare.Size()
		i -= size
		if _, err := m.SlashFractionWrongKeyshare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.SlashFractionNoKeyshare.Size()
		i -= size
		if _, err := m.SlashFractionNoKeyshare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.TrustedAddresses) > 0 {
		for iNdEx := len(m.TrustedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TrustedAddresses[iNdEx])
			copy(dAtA[i:], m.TrustedAddresses[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.TrustedAddresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.KeyExpiry != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.KeyExpiry))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.KeyExpiry != 0 {
		n += 1 + sovParams(uint64(m.KeyExpiry))
	}
	if len(m.TrustedAddresses) > 0 {
		for _, s := range m.TrustedAddresses {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = m.SlashFractionNoKeyshare.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashFractionWrongKeyshare.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MinimumBonded != 0 {
		n += 1 + sovParams(uint64(m.MinimumBonded))
	}
	if m.MaxIdledBlock != 0 {
		n += 1 + sovParams(uint64(m.MaxIdledBlock))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyExpiry", wireType)
			}
			m.KeyExpiry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyExpiry |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustedAddresses = append(m.TrustedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionNoKeyshare", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionNoKeyshare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionWrongKeyshare", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFractionWrongKeyshare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumBonded", wireType)
			}
			m.MinimumBonded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinimumBonded |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxIdledBlock", wireType)
			}
			m.MaxIdledBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxIdledBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)