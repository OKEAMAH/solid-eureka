// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wormhole/genesis.proto

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

// GenesisState defines the wormhole module's genesis state.
type GenesisState struct {
	GuardianSetList            []GuardianSet                          `protobuf:"bytes,1,rep,name=guardianSetList,proto3" json:"guardianSetList"`
	Config                     *Config                                `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	ReplayProtectionList       []ReplayProtection                     `protobuf:"bytes,3,rep,name=replayProtectionList,proto3" json:"replayProtectionList"`
	SequenceCounterList        []SequenceCounter                      `protobuf:"bytes,4,rep,name=sequenceCounterList,proto3" json:"sequenceCounterList"`
	ConsensusGuardianSetIndex  *ConsensusGuardianSetIndex             `protobuf:"bytes,5,opt,name=consensusGuardianSetIndex,proto3" json:"consensusGuardianSetIndex,omitempty"`
	GuardianValidatorList      []GuardianValidator                    `protobuf:"bytes,6,rep,name=guardianValidatorList,proto3" json:"guardianValidatorList"`
	AllowedAddresses           []ValidatorAllowedAddress              `protobuf:"bytes,7,rep,name=allowedAddresses,proto3" json:"allowedAddresses"`
	WasmInstantiateAllowlist   []WasmInstantiateAllowedContractCodeId `protobuf:"bytes,8,rep,name=wasmInstantiateAllowlist,proto3" json:"wasmInstantiateAllowlist"`
	IbcComposabilityMwContract IbcComposabilityMwContract             `protobuf:"bytes,9,opt,name=ibcComposabilityMwContract,proto3" json:"ibcComposabilityMwContract"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a7ced3fe0304831, []int{0}
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

func (m *GenesisState) GetGuardianSetList() []GuardianSet {
	if m != nil {
		return m.GuardianSetList
	}
	return nil
}

func (m *GenesisState) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *GenesisState) GetReplayProtectionList() []ReplayProtection {
	if m != nil {
		return m.ReplayProtectionList
	}
	return nil
}

func (m *GenesisState) GetSequenceCounterList() []SequenceCounter {
	if m != nil {
		return m.SequenceCounterList
	}
	return nil
}

func (m *GenesisState) GetConsensusGuardianSetIndex() *ConsensusGuardianSetIndex {
	if m != nil {
		return m.ConsensusGuardianSetIndex
	}
	return nil
}

func (m *GenesisState) GetGuardianValidatorList() []GuardianValidator {
	if m != nil {
		return m.GuardianValidatorList
	}
	return nil
}

func (m *GenesisState) GetAllowedAddresses() []ValidatorAllowedAddress {
	if m != nil {
		return m.AllowedAddresses
	}
	return nil
}

func (m *GenesisState) GetWasmInstantiateAllowlist() []WasmInstantiateAllowedContractCodeId {
	if m != nil {
		return m.WasmInstantiateAllowlist
	}
	return nil
}

func (m *GenesisState) GetIbcComposabilityMwContract() IbcComposabilityMwContract {
	if m != nil {
		return m.IbcComposabilityMwContract
	}
	return IbcComposabilityMwContract{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "wormhole_foundation.wormchain.wormhole.GenesisState")
}

func init() { proto.RegisterFile("wormhole/genesis.proto", fileDescriptor_9a7ced3fe0304831) }

var fileDescriptor_9a7ced3fe0304831 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0xb3, 0xb6, 0xc6, 0x3a, 0x15, 0x94, 0xb1, 0xd5, 0x35, 0x87, 0x6d, 0xf0, 0x20, 0x05,
	0x71, 0x17, 0xda, 0x83, 0xf6, 0x24, 0x49, 0xc0, 0x12, 0xa8, 0x20, 0x09, 0x28, 0x78, 0x59, 0x26,
	0x3b, 0xaf, 0x9b, 0x81, 0xcd, 0x4c, 0xba, 0x33, 0x4b, 0x9a, 0x93, 0x57, 0x4f, 0xe2, 0xb7, 0xf0,
	0xab, 0xf4, 0xd8, 0xa3, 0x27, 0x91, 0xe4, 0x8b, 0xc8, 0xce, 0xce, 0x4e, 0xda, 0xba, 0x29, 0xdb,
	0xdb, 0xf0, 0x66, 0xde, 0xef, 0xff, 0x7f, 0xff, 0xb7, 0x2c, 0x7a, 0x36, 0x13, 0xe9, 0x64, 0x2c,
	0x12, 0x08, 0x62, 0xe0, 0x20, 0x99, 0xf4, 0xa7, 0xa9, 0x50, 0x02, 0xbf, 0x2a, 0xeb, 0xe1, 0xa9,
	0xc8, 0x38, 0x25, 0x8a, 0x09, 0xee, 0xe7, 0xb5, 0x68, 0x4c, 0x58, 0x71, 0xca, 0x6f, 0x5b, 0xcf,
	0x57, 0xfd, 0x19, 0x49, 0x29, 0x23, 0xbc, 0x00, 0xb4, 0x76, 0xed, 0x45, 0x24, 0xf8, 0x29, 0x8b,
	0x4d, 0xb9, 0x6d, 0xcb, 0x29, 0x4c, 0x13, 0x32, 0x0f, 0xf3, 0x32, 0x44, 0x1a, 0x5f, 0xbc, 0xd8,
	0xb3, 0x2f, 0x24, 0x9c, 0x65, 0xc0, 0x23, 0x08, 0x23, 0x91, 0x71, 0x05, 0xa9, 0x79, 0xf0, 0xfa,
	0x2a, 0x59, 0x02, 0x97, 0x99, 0x0c, 0x4b, 0xf1, 0x50, 0x82, 0x0a, 0x19, 0xa7, 0x70, 0x6e, 0x1e,
	0xef, 0xc4, 0x22, 0x16, 0xfa, 0x18, 0xe4, 0xa7, 0xa2, 0xfa, 0xf2, 0xd7, 0x16, 0x7a, 0x74, 0x5c,
	0xcc, 0x3b, 0x54, 0x44, 0x01, 0x8e, 0xd0, 0xe3, 0x12, 0x31, 0x04, 0x75, 0xc2, 0xa4, 0x72, 0x9d,
	0xf6, 0xc6, 0xfe, 0xf6, 0xc1, 0xa1, 0x5f, 0x2f, 0x08, 0xff, 0x78, 0xd5, 0xde, 0xdd, 0xbc, 0xf8,
	0xb3, 0xd7, 0x18, 0xdc, 0x24, 0xe2, 0x0f, 0xa8, 0x59, 0x64, 0xe1, 0xde, 0x6b, 0x3b, 0xfb, 0xdb,
	0x07, 0x7e, 0x5d, 0x76, 0x4f, 0x77, 0x0d, 0x4c, 0x37, 0x4e, 0xd1, 0x4e, 0x11, 0xde, 0x27, 0x9b,
	0x9d, 0x76, 0xbc, 0xa1, 0x1d, 0xbf, 0xab, 0x4b, 0x1d, 0xdc, 0x60, 0x18, 0xdb, 0x95, 0x6c, 0x2c,
	0xd0, 0xd3, 0x72, 0x1d, 0xbd, 0x62, 0x1b, 0x5a, 0x72, 0x53, 0x4b, 0xbe, 0xad, 0x2b, 0x39, 0xbc,
	0x8e, 0x30, 0x8a, 0x55, 0x64, 0xfc, 0x0d, 0xbd, 0xb0, 0xeb, 0xbd, 0x92, 0x6d, 0x3f, 0xdf, 0xad,
	0x7b, 0x5f, 0xe7, 0xd7, 0xb9, 0x43, 0x7e, 0xd5, 0xa0, 0xc1, 0x7a, 0x0d, 0x9c, 0xa1, 0xdd, 0x72,
	0x81, 0x9f, 0x49, 0xc2, 0x28, 0x51, 0xa2, 0x98, 0xb9, 0xa9, 0x67, 0x3e, 0xba, 0xeb, 0x87, 0x61,
	0x21, 0x66, 0xea, 0x6a, 0x3a, 0x3e, 0x43, 0x4f, 0x48, 0x92, 0x88, 0x19, 0xd0, 0x0e, 0xa5, 0x29,
	0x48, 0x09, 0xd2, 0x7d, 0xa0, 0x15, 0xdf, 0xd7, 0x55, 0xb4, 0xc0, 0xce, 0x35, 0x90, 0xd1, 0xfd,
	0x0f, 0x8f, 0x7f, 0x38, 0xc8, 0x9d, 0x11, 0x39, 0xe9, 0x73, 0xa9, 0x08, 0x57, 0x8c, 0x28, 0xd0,
	0x9d, 0x49, 0x3e, 0xed, 0x96, 0xd6, 0x3e, 0xa9, 0xab, 0xfd, 0xa5, 0x82, 0x03, 0xb4, 0x27, 0xb8,
	0x4a, 0x49, 0xa4, 0x7a, 0x82, 0x42, 0x9f, 0x1a, 0x23, 0x6b, 0x35, 0xf1, 0x77, 0x07, 0xb5, 0xd8,
	0x28, 0xea, 0x89, 0xc9, 0x54, 0x48, 0x32, 0x62, 0x09, 0x53, 0xf3, 0x8f, 0xb3, 0x12, 0xe2, 0x3e,
	0xd4, 0xdb, 0xef, 0xd6, 0xb5, 0xd4, 0x5f, 0x4b, 0x32, 0x46, 0x6e, 0xd1, 0xea, 0x0e, 0x2f, 0x16,
	0x9e, 0x73, 0xb9, 0xf0, 0x9c, 0xbf, 0x0b, 0xcf, 0xf9, 0xb9, 0xf4, 0x1a, 0x97, 0x4b, 0xaf, 0xf1,
	0x7b, 0xe9, 0x35, 0xbe, 0x1e, 0xc5, 0x4c, 0x8d, 0xb3, 0x91, 0x1f, 0x89, 0x49, 0x50, 0x6a, 0xbd,
	0x59, 0x39, 0x09, 0xac, 0x93, 0xe0, 0xdc, 0xde, 0x07, 0x6a, 0x3e, 0x05, 0x39, 0x6a, 0xea, 0xbf,
	0xd0, 0xe1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xa6, 0x7e, 0xa5, 0x7d, 0x05, 0x00, 0x00,
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
		size, err := m.IbcComposabilityMwContract.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.WasmInstantiateAllowlist) > 0 {
		for iNdEx := len(m.WasmInstantiateAllowlist) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WasmInstantiateAllowlist[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.AllowedAddresses) > 0 {
		for iNdEx := len(m.AllowedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllowedAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.GuardianValidatorList) > 0 {
		for iNdEx := len(m.GuardianValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GuardianValidatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.ConsensusGuardianSetIndex != nil {
		{
			size, err := m.ConsensusGuardianSetIndex.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SequenceCounterList) > 0 {
		for iNdEx := len(m.SequenceCounterList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SequenceCounterList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ReplayProtectionList) > 0 {
		for iNdEx := len(m.ReplayProtectionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReplayProtectionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.GuardianSetList) > 0 {
		for iNdEx := len(m.GuardianSetList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GuardianSetList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
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
	if len(m.GuardianSetList) > 0 {
		for _, e := range m.GuardianSetList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ReplayProtectionList) > 0 {
		for _, e := range m.ReplayProtectionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SequenceCounterList) > 0 {
		for _, e := range m.SequenceCounterList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ConsensusGuardianSetIndex != nil {
		l = m.ConsensusGuardianSetIndex.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.GuardianValidatorList) > 0 {
		for _, e := range m.GuardianValidatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AllowedAddresses) > 0 {
		for _, e := range m.AllowedAddresses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.WasmInstantiateAllowlist) > 0 {
		for _, e := range m.WasmInstantiateAllowlist {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.IbcComposabilityMwContract.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field GuardianSetList", wireType)
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
			m.GuardianSetList = append(m.GuardianSetList, GuardianSet{})
			if err := m.GuardianSetList[len(m.GuardianSetList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
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
			if m.Config == nil {
				m.Config = &Config{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplayProtectionList", wireType)
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
			m.ReplayProtectionList = append(m.ReplayProtectionList, ReplayProtection{})
			if err := m.ReplayProtectionList[len(m.ReplayProtectionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceCounterList", wireType)
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
			m.SequenceCounterList = append(m.SequenceCounterList, SequenceCounter{})
			if err := m.SequenceCounterList[len(m.SequenceCounterList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusGuardianSetIndex", wireType)
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
			if m.ConsensusGuardianSetIndex == nil {
				m.ConsensusGuardianSetIndex = &ConsensusGuardianSetIndex{}
			}
			if err := m.ConsensusGuardianSetIndex.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuardianValidatorList", wireType)
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
			m.GuardianValidatorList = append(m.GuardianValidatorList, GuardianValidator{})
			if err := m.GuardianValidatorList[len(m.GuardianValidatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedAddresses", wireType)
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
			m.AllowedAddresses = append(m.AllowedAddresses, ValidatorAllowedAddress{})
			if err := m.AllowedAddresses[len(m.AllowedAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WasmInstantiateAllowlist", wireType)
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
			m.WasmInstantiateAllowlist = append(m.WasmInstantiateAllowlist, WasmInstantiateAllowedContractCodeId{})
			if err := m.WasmInstantiateAllowlist[len(m.WasmInstantiateAllowlist)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcComposabilityMwContract", wireType)
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
			if err := m.IbcComposabilityMwContract.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
