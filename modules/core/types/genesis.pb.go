// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/types/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/cosmos/ibc-go/v10/modules/core/02-client/types"
	types3 "github.com/cosmos/ibc-go/v10/modules/core/02-client/v2/types"
	types1 "github.com/cosmos/ibc-go/v10/modules/core/03-connection/types"
	types2 "github.com/cosmos/ibc-go/v10/modules/core/04-channel/types"
	types4 "github.com/cosmos/ibc-go/v10/modules/core/04-channel/v2/types"
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

// GenesisState defines the ibc module's genesis state.
type GenesisState struct {
	// ICS002 - Clients genesis state
	ClientGenesis types.GenesisState `protobuf:"bytes,1,opt,name=client_genesis,json=clientGenesis,proto3" json:"client_genesis"`
	// ICS003 - Connections genesis state
	ConnectionGenesis types1.GenesisState `protobuf:"bytes,2,opt,name=connection_genesis,json=connectionGenesis,proto3" json:"connection_genesis"`
	// ICS004 - Channel genesis state
	ChannelGenesis types2.GenesisState `protobuf:"bytes,3,opt,name=channel_genesis,json=channelGenesis,proto3" json:"channel_genesis"`
	// ICS002 - Clients/v2 genesis state
	ClientV2Genesis types3.GenesisState `protobuf:"bytes,4,opt,name=client_v2_genesis,json=clientV2Genesis,proto3" json:"client_v2_genesis"`
	// ICS004 - Channel/v2 genesis state
	ChannelV2Genesis types4.GenesisState `protobuf:"bytes,5,opt,name=channel_v2_genesis,json=channelV2Genesis,proto3" json:"channel_v2_genesis"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a49c5663e6fc59, []int{0}
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

func (m *GenesisState) GetClientGenesis() types.GenesisState {
	if m != nil {
		return m.ClientGenesis
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetConnectionGenesis() types1.GenesisState {
	if m != nil {
		return m.ConnectionGenesis
	}
	return types1.GenesisState{}
}

func (m *GenesisState) GetChannelGenesis() types2.GenesisState {
	if m != nil {
		return m.ChannelGenesis
	}
	return types2.GenesisState{}
}

func (m *GenesisState) GetClientV2Genesis() types3.GenesisState {
	if m != nil {
		return m.ClientV2Genesis
	}
	return types3.GenesisState{}
}

func (m *GenesisState) GetChannelV2Genesis() types4.GenesisState {
	if m != nil {
		return m.ChannelV2Genesis
	}
	return types4.GenesisState{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.core.types.v1.GenesisState")
}

func init() { proto.RegisterFile("ibc/core/types/v1/genesis.proto", fileDescriptor_b9a49c5663e6fc59) }

var fileDescriptor_b9a49c5663e6fc59 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0xfa, 0x30,
	0x1c, 0xc7, 0xb7, 0x3f, 0xfc, 0x3d, 0x54, 0x05, 0x59, 0x3c, 0x18, 0x0e, 0x03, 0x0c, 0x07, 0x2f,
	0xb6, 0x6e, 0xbe, 0x01, 0x17, 0xe3, 0xc1, 0xc4, 0x60, 0x34, 0xd1, 0x8b, 0x61, 0xb5, 0x19, 0x4d,
	0xa0, 0x3f, 0x42, 0xcb, 0x12, 0xdf, 0xc2, 0x97, 0xf1, 0x1d, 0x38, 0x72, 0xf4, 0x64, 0x0c, 0xbc,
	0x88, 0x61, 0xed, 0xba, 0x05, 0xb6, 0xc4, 0xdb, 0xb2, 0xef, 0xa7, 0x9f, 0xfe, 0xbe, 0xcd, 0x0f,
	0x75, 0x78, 0x44, 0x09, 0x85, 0x39, 0x23, 0xea, 0x7d, 0xc6, 0x24, 0x49, 0x02, 0x12, 0x33, 0xc1,
	0x24, 0x97, 0x78, 0x36, 0x07, 0x05, 0x5e, 0x8b, 0x47, 0x14, 0x6f, 0x01, 0x9c, 0x02, 0x38, 0x09,
	0xda, 0xa7, 0x31, 0xc4, 0x90, 0xa6, 0x64, 0xfb, 0xa5, 0xc1, 0x76, 0xd7, 0x9a, 0xe8, 0x84, 0x33,
	0xa1, 0xf6, 0x54, 0x25, 0x44, 0xb8, 0x43, 0xf4, 0x73, 0x02, 0x84, 0x60, 0x54, 0x71, 0x10, 0xfb,
	0x9e, 0x5e, 0x4e, 0x8d, 0x47, 0x42, 0xb0, 0xc9, 0x9f, 0x90, 0x9d, 0xbb, 0xce, 0x3f, 0x6b, 0xe8,
	0xe8, 0x46, 0xff, 0x79, 0x50, 0x23, 0xc5, 0xbc, 0x3b, 0xd4, 0xd0, 0x73, 0xbd, 0x1a, 0xf0, 0xcc,
	0xed, 0xba, 0x17, 0x87, 0x61, 0x17, 0xdb, 0x27, 0xd0, 0x39, 0x4e, 0x02, 0x5c, 0x3c, 0x39, 0xa8,
	0x2f, 0xbf, 0x3b, 0xce, 0xf0, 0x58, 0xa7, 0x26, 0xf1, 0x9e, 0x91, 0x97, 0x97, 0xb0, 0xca, 0x7f,
	0xa9, 0xb2, 0x5f, 0x50, 0x5a, 0xa6, 0x42, 0xdb, 0xca, 0x89, 0x4c, 0x7d, 0x8f, 0x9a, 0xa6, 0x96,
	0xf5, 0xd6, 0x52, 0x6f, 0xaf, 0xe0, 0xd5, 0x40, 0x85, 0xb4, 0x61, 0xe2, 0xcc, 0x38, 0x44, 0x2d,
	0xd3, 0x3d, 0x09, 0xad, 0xb3, 0x5e, 0x55, 0x3f, 0x2c, 0x53, 0x36, 0x75, 0xfa, 0x14, 0x66, 0xce,
	0x47, 0xe4, 0x65, 0x53, 0x16, 0xa4, 0xff, 0x2b, 0x07, 0x2d, 0xb5, 0x9e, 0x98, 0xd8, 0x6a, 0x07,
	0xb7, 0xcb, 0xb5, 0xef, 0xae, 0xd6, 0xbe, 0xfb, 0xb3, 0xf6, 0xdd, 0x8f, 0x8d, 0xef, 0xac, 0x36,
	0xbe, 0xf3, 0xb5, 0xf1, 0x9d, 0x17, 0x12, 0x73, 0x35, 0x5e, 0x44, 0x98, 0xc2, 0x94, 0x50, 0x90,
	0x53, 0x90, 0x84, 0x47, 0xf4, 0x32, 0x06, 0x92, 0x04, 0x57, 0x64, 0x0a, 0x6f, 0x8b, 0x09, 0x93,
	0x85, 0x65, 0x8f, 0x0e, 0xd2, 0x4d, 0xb8, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x80, 0x8d, 0xfa,
	0x2d, 0x05, 0x03, 0x00, 0x00,
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
		size, err := m.ChannelV2Genesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.ClientV2Genesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.ChannelGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ConnectionGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ClientGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.ClientGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ConnectionGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ChannelGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ClientV2Genesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ChannelV2Genesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientGenesis", wireType)
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
			if err := m.ClientGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionGenesis", wireType)
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
			if err := m.ConnectionGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelGenesis", wireType)
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
			if err := m.ChannelGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientV2Genesis", wireType)
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
			if err := m.ClientV2Genesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelV2Genesis", wireType)
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
			if err := m.ChannelV2Genesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
