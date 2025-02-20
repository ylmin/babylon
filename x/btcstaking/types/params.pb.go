// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/btcstaking/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_babylonchain_babylon_types "github.com/babylonchain/babylon/types"
	_ "github.com/cosmos/cosmos-proto"
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
	// covenant_pks is the list of public keys held by the covenant committee
	// each PK follows encoding in BIP-340 spec on Bitcoin
	CovenantPks []github_com_babylonchain_babylon_types.BIP340PubKey `protobuf:"bytes,1,rep,name=covenant_pks,json=covenantPks,proto3,customtype=github.com/babylonchain/babylon/types.BIP340PubKey" json:"covenant_pks,omitempty"`
	// covenant_quorum is the minimum number of signatures needed for the covenant
	// multisignature
	CovenantQuorum uint32 `protobuf:"varint,2,opt,name=covenant_quorum,json=covenantQuorum,proto3" json:"covenant_quorum,omitempty"`
	// slashing address is the address that the slashed BTC goes to
	// the address is in string on Bitcoin
	SlashingAddress string `protobuf:"bytes,3,opt,name=slashing_address,json=slashingAddress,proto3" json:"slashing_address,omitempty"`
	// min_slashing_tx_fee_sat is the minimum amount of tx fee (quantified
	// in Satoshi) needed for the pre-signed slashing tx
	// TODO: change to satoshi per byte?
	MinSlashingTxFeeSat int64 `protobuf:"varint,4,opt,name=min_slashing_tx_fee_sat,json=minSlashingTxFeeSat,proto3" json:"min_slashing_tx_fee_sat,omitempty"`
	// min_commission_rate is the chain-wide minimum commission rate that a finality provider can charge their delegators
	MinCommissionRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=min_commission_rate,json=minCommissionRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min_commission_rate"`
	// slashing_rate determines the portion of the staked amount to be slashed,
	// expressed as a decimal (e.g., 0.5 for 50%).
	SlashingRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,6,opt,name=slashing_rate,json=slashingRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slashing_rate"`
	// max_active_finality_providers is the maximum number of active finality providers in the BTC staking protocol
	MaxActiveFinalityProviders uint32 `protobuf:"varint,7,opt,name=max_active_finality_providers,json=maxActiveFinalityProviders,proto3" json:"max_active_finality_providers,omitempty"`
	// min_unbonding_time is the minimum time for unbonding transaction timelock in BTC blocks
	MinUnbondingTime uint32 `protobuf:"varint,8,opt,name=min_unbonding_time,json=minUnbondingTime,proto3" json:"min_unbonding_time,omitempty"`
	// min_unbonding_rate is the minimum amount of BTC that are required in unbonding
	// output, expressed as a fraction of staking output
	// example: if min_unbonding_rate=0.9, then the unbonding output value
	// must be at least 90% of staking output, for staking request to be considered
	// valid
	MinUnbondingRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,9,opt,name=min_unbonding_rate,json=minUnbondingRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min_unbonding_rate"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1392776a3e15b9, []int{0}
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

func (m *Params) GetCovenantQuorum() uint32 {
	if m != nil {
		return m.CovenantQuorum
	}
	return 0
}

func (m *Params) GetSlashingAddress() string {
	if m != nil {
		return m.SlashingAddress
	}
	return ""
}

func (m *Params) GetMinSlashingTxFeeSat() int64 {
	if m != nil {
		return m.MinSlashingTxFeeSat
	}
	return 0
}

func (m *Params) GetMaxActiveFinalityProviders() uint32 {
	if m != nil {
		return m.MaxActiveFinalityProviders
	}
	return 0
}

func (m *Params) GetMinUnbondingTime() uint32 {
	if m != nil {
		return m.MinUnbondingTime
	}
	return 0
}

// StoredParams attach information about the version of stored parameters
type StoredParams struct {
	// version of the stored parameters. Each parameters update
	// increments version number by 1
	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// NOTE: Parameters must always be provided
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *StoredParams) Reset()         { *m = StoredParams{} }
func (m *StoredParams) String() string { return proto.CompactTextString(m) }
func (*StoredParams) ProtoMessage()    {}
func (*StoredParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d1392776a3e15b9, []int{1}
}
func (m *StoredParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredParams.Merge(m, src)
}
func (m *StoredParams) XXX_Size() int {
	return m.Size()
}
func (m *StoredParams) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredParams.DiscardUnknown(m)
}

var xxx_messageInfo_StoredParams proto.InternalMessageInfo

func (m *StoredParams) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StoredParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*Params)(nil), "babylon.btcstaking.v1.Params")
	proto.RegisterType((*StoredParams)(nil), "babylon.btcstaking.v1.StoredParams")
}

func init() {
	proto.RegisterFile("babylon/btcstaking/v1/params.proto", fileDescriptor_8d1392776a3e15b9)
}

var fileDescriptor_8d1392776a3e15b9 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0xbf, 0xe4, 0x4b, 0xe9, 0x34, 0xa5, 0x65, 0x00, 0x61, 0x82, 0xea, 0x44, 0x61, 0x41,
	0x90, 0xc0, 0x26, 0x6d, 0xc5, 0x02, 0x56, 0x09, 0xa8, 0x12, 0xa2, 0x8b, 0xe0, 0x14, 0x24, 0xd8,
	0x8c, 0xc6, 0xf6, 0xd4, 0x19, 0x25, 0x33, 0x13, 0x3c, 0x13, 0xcb, 0x7e, 0x0b, 0x96, 0x2c, 0x79,
	0x08, 0x1e, 0xa2, 0xcb, 0x8a, 0x15, 0xea, 0x22, 0x42, 0xc9, 0x1b, 0xf0, 0x04, 0xc8, 0x63, 0x3b,
	0xfc, 0x08, 0x09, 0xc4, 0xce, 0xf7, 0xdc, 0x73, 0xcf, 0xfd, 0xf1, 0x19, 0xd0, 0xf1, 0xb0, 0x97,
	0x4e, 0x05, 0x77, 0x3c, 0xe5, 0x4b, 0x85, 0x27, 0x94, 0x87, 0x4e, 0xdc, 0x73, 0x66, 0x38, 0xc2,
	0x4c, 0xda, 0xb3, 0x48, 0x28, 0x01, 0xaf, 0x17, 0x1c, 0xfb, 0x3b, 0xc7, 0x8e, 0x7b, 0xcd, 0x6b,
	0xa1, 0x08, 0x85, 0x66, 0x38, 0xd9, 0x57, 0x4e, 0x6e, 0xde, 0xf4, 0x85, 0x64, 0x42, 0xa2, 0x3c,
	0x91, 0x07, 0x79, 0xaa, 0xf3, 0xb5, 0x06, 0xea, 0x43, 0x2d, 0x0c, 0x5f, 0x83, 0x86, 0x2f, 0x62,
	0xc2, 0x31, 0x57, 0x68, 0x36, 0x91, 0xa6, 0xd1, 0xae, 0x76, 0x1b, 0x83, 0x87, 0x17, 0x8b, 0xd6,
	0x7e, 0x48, 0xd5, 0x78, 0xee, 0xd9, 0xbe, 0x60, 0x4e, 0xd1, 0xd7, 0x1f, 0x63, 0xca, 0xcb, 0xc0,
	0x51, 0xe9, 0x8c, 0x48, 0x7b, 0xf0, 0x6c, 0x78, 0x70, 0xf8, 0x60, 0x38, 0xf7, 0x9e, 0x93, 0xd4,
	0xdd, 0x2a, 0xb5, 0x86, 0x13, 0x09, 0xef, 0x80, 0x9d, 0xb5, 0xf4, 0xdb, 0xb9, 0x88, 0xe6, 0xcc,
	0xfc, 0xaf, 0x6d, 0x74, 0xb7, 0xdd, 0xcb, 0x25, 0xfc, 0x42, 0xa3, 0xf0, 0x2e, 0xd8, 0x95, 0x53,
	0x2c, 0xc7, 0x94, 0x87, 0x08, 0x07, 0x41, 0x44, 0xa4, 0x34, 0xab, 0x6d, 0xa3, 0xbb, 0xe9, 0xee,
	0x94, 0x78, 0x3f, 0x87, 0xe1, 0x21, 0xb8, 0xc1, 0x28, 0x47, 0x6b, 0xba, 0x4a, 0xd0, 0x29, 0x21,
	0x48, 0x62, 0x65, 0xd6, 0xda, 0x46, 0xb7, 0xea, 0x5e, 0x65, 0x94, 0x8f, 0x8a, 0xec, 0x49, 0x72,
	0x44, 0xc8, 0x08, 0x2b, 0x38, 0x02, 0x19, 0x8c, 0x7c, 0xc1, 0x18, 0x95, 0x92, 0x0a, 0x8e, 0x22,
	0xac, 0x88, 0xf9, 0x7f, 0xd6, 0x63, 0x70, 0xfb, 0x6c, 0xd1, 0xaa, 0x5c, 0x2c, 0x5a, 0xb7, 0xf2,
	0x13, 0xc9, 0x60, 0x62, 0x53, 0xe1, 0x30, 0xac, 0xc6, 0xf6, 0x31, 0x09, 0xb1, 0x9f, 0x3e, 0x25,
	0xbe, 0x7b, 0x85, 0x51, 0xfe, 0x64, 0x5d, 0xee, 0x62, 0x45, 0xe0, 0x2b, 0xb0, 0xbd, 0x1e, 0x43,
	0xcb, 0xd5, 0xb5, 0x5c, 0xef, 0x2f, 0xe4, 0x3e, 0x7d, 0xbc, 0x0f, 0x8a, 0x1f, 0x92, 0x89, 0x37,
	0x4a, 0x1d, 0xad, 0xdb, 0x07, 0x7b, 0x0c, 0x27, 0x08, 0xfb, 0x8a, 0xc6, 0x04, 0x9d, 0x52, 0x8e,
	0xa7, 0x54, 0xa5, 0xd9, 0x6f, 0x8c, 0x69, 0x40, 0x22, 0x69, 0x6e, 0xe8, 0x23, 0x36, 0x19, 0x4e,
	0xfa, 0x9a, 0x73, 0x54, 0x50, 0x86, 0x25, 0x03, 0xde, 0x03, 0x30, 0xdb, 0x77, 0xce, 0x3d, 0xc1,
	0x03, 0x7d, 0x26, 0xca, 0x88, 0x79, 0x49, 0xd7, 0xed, 0x32, 0xca, 0x5f, 0x96, 0x89, 0x13, 0xca,
	0x08, 0x44, 0xbf, 0xb2, 0xf5, 0x36, 0x9b, 0xff, 0xba, 0xcd, 0x4f, 0x0d, 0xb2, 0x8d, 0x1e, 0xd5,
	0xde, 0x7f, 0x68, 0x55, 0x3a, 0x04, 0x34, 0x46, 0x4a, 0x44, 0x24, 0x28, 0x9c, 0x67, 0x82, 0x8d,
	0x98, 0x44, 0xd9, 0x39, 0x4d, 0x43, 0x4f, 0x56, 0x86, 0xf0, 0x31, 0xa8, 0xe7, 0xb6, 0xd7, 0x7e,
	0xd9, 0xda, 0xdf, 0xb3, 0x7f, 0xeb, 0x7b, 0x3b, 0x17, 0x1a, 0xd4, 0xb2, 0x19, 0xdd, 0xa2, 0x64,
	0x70, 0x7c, 0xb6, 0xb4, 0x8c, 0xf3, 0xa5, 0x65, 0x7c, 0x59, 0x5a, 0xc6, 0xbb, 0x95, 0x55, 0x39,
	0x5f, 0x59, 0x95, 0xcf, 0x2b, 0xab, 0xf2, 0xe6, 0x8f, 0x86, 0x4e, 0x7e, 0x7c, 0x7b, 0xda, 0xdd,
	0x5e, 0x5d, 0x3f, 0x98, 0x83, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x91, 0xdb, 0xb2, 0x9e,
	0x03, 0x00, 0x00,
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
	{
		size := m.MinUnbondingRate.Size()
		i -= size
		if _, err := m.MinUnbondingRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.MinUnbondingTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinUnbondingTime))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxActiveFinalityProviders != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxActiveFinalityProviders))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.SlashingRate.Size()
		i -= size
		if _, err := m.SlashingRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.MinCommissionRate.Size()
		i -= size
		if _, err := m.MinCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.MinSlashingTxFeeSat != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinSlashingTxFeeSat))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SlashingAddress) > 0 {
		i -= len(m.SlashingAddress)
		copy(dAtA[i:], m.SlashingAddress)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SlashingAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.CovenantQuorum != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CovenantQuorum))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CovenantPks) > 0 {
		for iNdEx := len(m.CovenantPks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.CovenantPks[iNdEx].Size()
				i -= size
				if _, err := m.CovenantPks[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *StoredParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Version != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Version))
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
	if len(m.CovenantPks) > 0 {
		for _, e := range m.CovenantPks {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.CovenantQuorum != 0 {
		n += 1 + sovParams(uint64(m.CovenantQuorum))
	}
	l = len(m.SlashingAddress)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.MinSlashingTxFeeSat != 0 {
		n += 1 + sovParams(uint64(m.MinSlashingTxFeeSat))
	}
	l = m.MinCommissionRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.SlashingRate.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.MaxActiveFinalityProviders != 0 {
		n += 1 + sovParams(uint64(m.MaxActiveFinalityProviders))
	}
	if m.MinUnbondingTime != 0 {
		n += 1 + sovParams(uint64(m.MinUnbondingTime))
	}
	l = m.MinUnbondingRate.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *StoredParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovParams(uint64(m.Version))
	}
	l = m.Params.Size()
	n += 1 + l + sovParams(uint64(l))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CovenantPks", wireType)
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
			var v github_com_babylonchain_babylon_types.BIP340PubKey
			m.CovenantPks = append(m.CovenantPks, v)
			if err := m.CovenantPks[len(m.CovenantPks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CovenantQuorum", wireType)
			}
			m.CovenantQuorum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CovenantQuorum |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashingAddress", wireType)
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
			m.SlashingAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSlashingTxFeeSat", wireType)
			}
			m.MinSlashingTxFeeSat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinSlashingTxFeeSat |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCommissionRate", wireType)
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
			if err := m.MinCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashingRate", wireType)
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
			if err := m.SlashingRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxActiveFinalityProviders", wireType)
			}
			m.MaxActiveFinalityProviders = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxActiveFinalityProviders |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUnbondingTime", wireType)
			}
			m.MinUnbondingTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinUnbondingTime |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinUnbondingRate", wireType)
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
			if err := m.MinUnbondingRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *StoredParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StoredParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
