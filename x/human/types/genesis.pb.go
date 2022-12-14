// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: human/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the human module's genesis state.
type GenesisState struct {
	Params              Params            `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FeeBalanceList      []FeeBalance      `protobuf:"bytes,2,rep,name=feeBalanceList,proto3" json:"feeBalanceList"`
	KeysignVoteDataList []KeysignVoteData `protobuf:"bytes,3,rep,name=keysignVoteDataList,proto3" json:"keysignVoteDataList"`
	ObserveVoteList     []ObserveVote     `protobuf:"bytes,4,rep,name=observeVoteList,proto3" json:"observeVoteList"`
	PoolBalanceList     []PoolBalance     `protobuf:"bytes,5,rep,name=poolBalanceList,proto3" json:"poolBalanceList"`
	PoolBalanapList     []PoolBalanap     `protobuf:"bytes,6,rep,name=poolBalanapList,proto3" json:"poolBalanapList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_76e17e7acb62ce69, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFeeBalanceList() []FeeBalance {
	if m != nil {
		return m.FeeBalanceList
	}
	return nil
}

func (m *GenesisState) GetKeysignVoteDataList() []KeysignVoteData {
	if m != nil {
		return m.KeysignVoteDataList
	}
	return nil
}

func (m *GenesisState) GetObserveVoteList() []ObserveVote {
	if m != nil {
		return m.ObserveVoteList
	}
	return nil
}

func (m *GenesisState) GetPoolBalanceList() []PoolBalance {
	if m != nil {
		return m.PoolBalanceList
	}
	return nil
}

func (m *GenesisState) GetPoolBalanapList() []PoolBalanap {
	if m != nil {
		return m.PoolBalanapList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "human.human.GenesisState")
}

func init() { proto.RegisterFile("human/genesis.proto", fileDescriptor_76e17e7acb62ce69) }

var fileDescriptor_76e17e7acb62ce69 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x86, 0xdb, 0x0f, 0x3e, 0x16, 0x83, 0xd1, 0x64, 0x30, 0xa1, 0x21, 0x3a, 0x12, 0x57, 0x6c,
	0x2c, 0x11, 0xef, 0x80, 0xf8, 0x97, 0x68, 0x22, 0x51, 0xe3, 0xc2, 0x0d, 0x39, 0xe8, 0xa1, 0x12,
	0xa1, 0x33, 0x61, 0x46, 0x22, 0x77, 0xe1, 0x05, 0x79, 0x01, 0x2c, 0x59, 0xba, 0x32, 0xa6, 0xbd,
	0x11, 0xd3, 0x99, 0x51, 0x98, 0x42, 0xdc, 0x40, 0x73, 0xde, 0xf7, 0x79, 0xda, 0x33, 0x2d, 0xa9,
	0x3c, 0xbd, 0x8c, 0x20, 0x6e, 0x46, 0x18, 0xa3, 0x1c, 0xc8, 0x50, 0x8c, 0xb9, 0xe2, 0xb4, 0xac,
	0x87, 0xa1, 0xfe, 0xad, 0x6d, 0x47, 0x3c, 0xe2, 0x7a, 0xde, 0xcc, 0xae, 0x4c, 0xa5, 0x46, 0x0d,
	0x27, 0x60, 0x0c, 0x23, 0x8b, 0xd5, 0xaa, 0x66, 0xd6, 0x47, 0xec, 0xf6, 0x60, 0x08, 0xf1, 0x03,
	0xda, 0x60, 0xd7, 0x04, 0xcf, 0x38, 0x95, 0x83, 0x28, 0xee, 0x4e, 0xb8, 0xc2, 0xee, 0x23, 0x28,
	0xb0, 0x71, 0x60, 0x62, 0xde, 0x93, 0x38, 0x9e, 0xa0, 0x8e, 0xdd, 0x44, 0x70, 0x3e, 0xcc, 0x29,
	0x57, 0x12, 0x10, 0x26, 0xd9, 0x7f, 0x2f, 0x90, 0x8d, 0x33, 0xb3, 0xce, 0x8d, 0x02, 0x85, 0xf4,
	0x90, 0x94, 0xcc, 0x63, 0x06, 0x7e, 0xdd, 0x6f, 0x94, 0x5b, 0x95, 0x70, 0x69, 0xbd, 0xb0, 0xa3,
	0xa3, 0x76, 0x71, 0xf6, 0xb9, 0xe7, 0x5d, 0xdb, 0x22, 0x3d, 0x21, 0x9b, 0x7d, 0xc4, 0xb6, 0xb9,
	0xe3, 0xe5, 0x40, 0xaa, 0xe0, 0x5f, 0xbd, 0xd0, 0x28, 0xb7, 0xaa, 0x0e, 0x7a, 0xfa, 0x5b, 0xb1,
	0x78, 0x0e, 0xa2, 0xb7, 0xa4, 0x62, 0x77, 0xbe, 0xe3, 0x0a, 0x8f, 0x41, 0x81, 0x76, 0x15, 0xb4,
	0x6b, 0xc7, 0x71, 0x5d, 0xb8, 0x3d, 0x2b, 0x5c, 0x87, 0xd3, 0x73, 0xb2, 0x65, 0x8f, 0x2a, 0x1b,
	0x6b, 0x63, 0x51, 0x1b, 0x03, 0xc7, 0x78, 0xb5, 0xe8, 0x58, 0x5b, 0x1e, 0xcb, 0x4c, 0xd9, 0x01,
	0x2e, 0xef, 0xf9, 0x7f, 0x8d, 0xa9, 0xb3, 0xe8, 0xfc, 0x98, 0x72, 0x98, 0x63, 0x02, 0xa1, 0x4d,
	0xa5, 0xbf, 0x4c, 0x20, 0x56, 0x4c, 0x06, 0x6b, 0x1f, 0xcc, 0x12, 0xe6, 0xcf, 0x13, 0xe6, 0x7f,
	0x25, 0xcc, 0x7f, 0x4b, 0x99, 0x37, 0x4f, 0x99, 0xf7, 0x91, 0x32, 0xef, 0xde, 0x7e, 0xaa, 0xaf,
	0x4d, 0xf3, 0xaf, 0xa6, 0x02, 0x65, 0xaf, 0xa4, 0x5f, 0xfa, 0xd1, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0x33, 0x18, 0xc9, 0xc8, 0x02, 0x00, 0x00,
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
	if len(m.PoolBalanapList) > 0 {
		for iNdEx := len(m.PoolBalanapList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolBalanapList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.PoolBalanceList) > 0 {
		for iNdEx := len(m.PoolBalanceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolBalanceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ObserveVoteList) > 0 {
		for iNdEx := len(m.ObserveVoteList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ObserveVoteList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.KeysignVoteDataList) > 0 {
		for iNdEx := len(m.KeysignVoteDataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KeysignVoteDataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FeeBalanceList) > 0 {
		for iNdEx := len(m.FeeBalanceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeBalanceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FeeBalanceList) > 0 {
		for _, e := range m.FeeBalanceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.KeysignVoteDataList) > 0 {
		for _, e := range m.KeysignVoteDataList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ObserveVoteList) > 0 {
		for _, e := range m.ObserveVoteList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PoolBalanceList) > 0 {
		for _, e := range m.PoolBalanceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PoolBalanapList) > 0 {
		for _, e := range m.PoolBalanapList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeBalanceList", wireType)
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
			m.FeeBalanceList = append(m.FeeBalanceList, FeeBalance{})
			if err := m.FeeBalanceList[len(m.FeeBalanceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeysignVoteDataList", wireType)
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
			m.KeysignVoteDataList = append(m.KeysignVoteDataList, KeysignVoteData{})
			if err := m.KeysignVoteDataList[len(m.KeysignVoteDataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserveVoteList", wireType)
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
			m.ObserveVoteList = append(m.ObserveVoteList, ObserveVote{})
			if err := m.ObserveVoteList[len(m.ObserveVoteList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolBalanceList", wireType)
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
			m.PoolBalanceList = append(m.PoolBalanceList, PoolBalance{})
			if err := m.PoolBalanceList[len(m.PoolBalanceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolBalanapList", wireType)
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
			m.PoolBalanapList = append(m.PoolBalanapList, PoolBalanap{})
			if err := m.PoolBalanapList[len(m.PoolBalanapList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
