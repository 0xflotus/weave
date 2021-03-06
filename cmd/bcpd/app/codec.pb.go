// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/codec.proto

/*
	Package app is a generated protocol buffer package.

	It is generated from these files:
		app/codec.proto

	It has these top-level messages:
		Tx
*/
package app

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cash "github.com/iov-one/weave/x/cash"
import escrow "github.com/iov-one/weave/x/escrow"
import namecoin "github.com/iov-one/weave/x/namecoin"
import sigs "github.com/iov-one/weave/x/sigs"
import multisig "github.com/iov-one/weave/x/multisig"
import validators "github.com/iov-one/weave/x/validators"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Tx contains the message
type Tx struct {
	// msg is a sum type over all allowed messages on this chain.
	//
	// Types that are valid to be assigned to Sum:
	//	*Tx_SendMsg
	//	*Tx_NewTokenMsg
	//	*Tx_SetNameMsg
	//	*Tx_CreateEscrowMsg
	//	*Tx_ReleaseEscrowMsg
	//	*Tx_ReturnEscrowMsg
	//	*Tx_UpdateEscrowMsg
	//	*Tx_CreateContractMsg
	//	*Tx_UpdateContractMsg
	//	*Tx_SetValidatorsMsg
	Sum isTx_Sum `protobuf_oneof:"sum"`
	// fee info, autogenerates GetFees()
	Fees *cash.FeeInfo `protobuf:"bytes,20,opt,name=fees" json:"fees,omitempty"`
	// signatures, autogenerates GetSignatures()
	Signatures []*sigs.StdSignature `protobuf:"bytes,21,rep,name=signatures" json:"signatures,omitempty"`
	// preimage for hashlock, autogenerates GetPreimage
	Preimage []byte `protobuf:"bytes,22,opt,name=preimage,proto3" json:"preimage,omitempty"`
	// id of multisig contract, autogenerates GetMultisig
	Multisig [][]byte `protobuf:"bytes,23,rep,name=multisig" json:"multisig,omitempty"`
}

func (m *Tx) Reset()                    { *m = Tx{} }
func (m *Tx) String() string            { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()               {}
func (*Tx) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{0} }

type isTx_Sum interface {
	isTx_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Tx_SendMsg struct {
	SendMsg *cash.SendMsg `protobuf:"bytes,1,opt,name=send_msg,json=sendMsg,oneof"`
}
type Tx_NewTokenMsg struct {
	NewTokenMsg *namecoin.NewTokenMsg `protobuf:"bytes,2,opt,name=new_token_msg,json=newTokenMsg,oneof"`
}
type Tx_SetNameMsg struct {
	SetNameMsg *namecoin.SetWalletNameMsg `protobuf:"bytes,3,opt,name=set_name_msg,json=setNameMsg,oneof"`
}
type Tx_CreateEscrowMsg struct {
	CreateEscrowMsg *escrow.CreateEscrowMsg `protobuf:"bytes,4,opt,name=create_escrow_msg,json=createEscrowMsg,oneof"`
}
type Tx_ReleaseEscrowMsg struct {
	ReleaseEscrowMsg *escrow.ReleaseEscrowMsg `protobuf:"bytes,5,opt,name=release_escrow_msg,json=releaseEscrowMsg,oneof"`
}
type Tx_ReturnEscrowMsg struct {
	ReturnEscrowMsg *escrow.ReturnEscrowMsg `protobuf:"bytes,6,opt,name=return_escrow_msg,json=returnEscrowMsg,oneof"`
}
type Tx_UpdateEscrowMsg struct {
	UpdateEscrowMsg *escrow.UpdateEscrowPartiesMsg `protobuf:"bytes,7,opt,name=update_escrow_msg,json=updateEscrowMsg,oneof"`
}
type Tx_CreateContractMsg struct {
	CreateContractMsg *multisig.CreateContractMsg `protobuf:"bytes,8,opt,name=create_contract_msg,json=createContractMsg,oneof"`
}
type Tx_UpdateContractMsg struct {
	UpdateContractMsg *multisig.UpdateContractMsg `protobuf:"bytes,9,opt,name=update_contract_msg,json=updateContractMsg,oneof"`
}
type Tx_SetValidatorsMsg struct {
	SetValidatorsMsg *validators.SetValidatorsMsg `protobuf:"bytes,10,opt,name=set_validators_msg,json=setValidatorsMsg,oneof"`
}

func (*Tx_SendMsg) isTx_Sum()           {}
func (*Tx_NewTokenMsg) isTx_Sum()       {}
func (*Tx_SetNameMsg) isTx_Sum()        {}
func (*Tx_CreateEscrowMsg) isTx_Sum()   {}
func (*Tx_ReleaseEscrowMsg) isTx_Sum()  {}
func (*Tx_ReturnEscrowMsg) isTx_Sum()   {}
func (*Tx_UpdateEscrowMsg) isTx_Sum()   {}
func (*Tx_CreateContractMsg) isTx_Sum() {}
func (*Tx_UpdateContractMsg) isTx_Sum() {}
func (*Tx_SetValidatorsMsg) isTx_Sum()  {}

func (m *Tx) GetSum() isTx_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Tx) GetSendMsg() *cash.SendMsg {
	if x, ok := m.GetSum().(*Tx_SendMsg); ok {
		return x.SendMsg
	}
	return nil
}

func (m *Tx) GetNewTokenMsg() *namecoin.NewTokenMsg {
	if x, ok := m.GetSum().(*Tx_NewTokenMsg); ok {
		return x.NewTokenMsg
	}
	return nil
}

func (m *Tx) GetSetNameMsg() *namecoin.SetWalletNameMsg {
	if x, ok := m.GetSum().(*Tx_SetNameMsg); ok {
		return x.SetNameMsg
	}
	return nil
}

func (m *Tx) GetCreateEscrowMsg() *escrow.CreateEscrowMsg {
	if x, ok := m.GetSum().(*Tx_CreateEscrowMsg); ok {
		return x.CreateEscrowMsg
	}
	return nil
}

func (m *Tx) GetReleaseEscrowMsg() *escrow.ReleaseEscrowMsg {
	if x, ok := m.GetSum().(*Tx_ReleaseEscrowMsg); ok {
		return x.ReleaseEscrowMsg
	}
	return nil
}

func (m *Tx) GetReturnEscrowMsg() *escrow.ReturnEscrowMsg {
	if x, ok := m.GetSum().(*Tx_ReturnEscrowMsg); ok {
		return x.ReturnEscrowMsg
	}
	return nil
}

func (m *Tx) GetUpdateEscrowMsg() *escrow.UpdateEscrowPartiesMsg {
	if x, ok := m.GetSum().(*Tx_UpdateEscrowMsg); ok {
		return x.UpdateEscrowMsg
	}
	return nil
}

func (m *Tx) GetCreateContractMsg() *multisig.CreateContractMsg {
	if x, ok := m.GetSum().(*Tx_CreateContractMsg); ok {
		return x.CreateContractMsg
	}
	return nil
}

func (m *Tx) GetUpdateContractMsg() *multisig.UpdateContractMsg {
	if x, ok := m.GetSum().(*Tx_UpdateContractMsg); ok {
		return x.UpdateContractMsg
	}
	return nil
}

func (m *Tx) GetSetValidatorsMsg() *validators.SetValidatorsMsg {
	if x, ok := m.GetSum().(*Tx_SetValidatorsMsg); ok {
		return x.SetValidatorsMsg
	}
	return nil
}

func (m *Tx) GetFees() *cash.FeeInfo {
	if m != nil {
		return m.Fees
	}
	return nil
}

func (m *Tx) GetSignatures() []*sigs.StdSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Tx) GetPreimage() []byte {
	if m != nil {
		return m.Preimage
	}
	return nil
}

func (m *Tx) GetMultisig() [][]byte {
	if m != nil {
		return m.Multisig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Tx) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Tx_OneofMarshaler, _Tx_OneofUnmarshaler, _Tx_OneofSizer, []interface{}{
		(*Tx_SendMsg)(nil),
		(*Tx_NewTokenMsg)(nil),
		(*Tx_SetNameMsg)(nil),
		(*Tx_CreateEscrowMsg)(nil),
		(*Tx_ReleaseEscrowMsg)(nil),
		(*Tx_ReturnEscrowMsg)(nil),
		(*Tx_UpdateEscrowMsg)(nil),
		(*Tx_CreateContractMsg)(nil),
		(*Tx_UpdateContractMsg)(nil),
		(*Tx_SetValidatorsMsg)(nil),
	}
}

func _Tx_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Tx)
	// sum
	switch x := m.Sum.(type) {
	case *Tx_SendMsg:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SendMsg); err != nil {
			return err
		}
	case *Tx_NewTokenMsg:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NewTokenMsg); err != nil {
			return err
		}
	case *Tx_SetNameMsg:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SetNameMsg); err != nil {
			return err
		}
	case *Tx_CreateEscrowMsg:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateEscrowMsg); err != nil {
			return err
		}
	case *Tx_ReleaseEscrowMsg:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReleaseEscrowMsg); err != nil {
			return err
		}
	case *Tx_ReturnEscrowMsg:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ReturnEscrowMsg); err != nil {
			return err
		}
	case *Tx_UpdateEscrowMsg:
		_ = b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateEscrowMsg); err != nil {
			return err
		}
	case *Tx_CreateContractMsg:
		_ = b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreateContractMsg); err != nil {
			return err
		}
	case *Tx_UpdateContractMsg:
		_ = b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateContractMsg); err != nil {
			return err
		}
	case *Tx_SetValidatorsMsg:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SetValidatorsMsg); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Tx.Sum has unexpected type %T", x)
	}
	return nil
}

func _Tx_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Tx)
	switch tag {
	case 1: // sum.send_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(cash.SendMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_SendMsg{msg}
		return true, err
	case 2: // sum.new_token_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(namecoin.NewTokenMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_NewTokenMsg{msg}
		return true, err
	case 3: // sum.set_name_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(namecoin.SetWalletNameMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_SetNameMsg{msg}
		return true, err
	case 4: // sum.create_escrow_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(escrow.CreateEscrowMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_CreateEscrowMsg{msg}
		return true, err
	case 5: // sum.release_escrow_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(escrow.ReleaseEscrowMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_ReleaseEscrowMsg{msg}
		return true, err
	case 6: // sum.return_escrow_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(escrow.ReturnEscrowMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_ReturnEscrowMsg{msg}
		return true, err
	case 7: // sum.update_escrow_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(escrow.UpdateEscrowPartiesMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_UpdateEscrowMsg{msg}
		return true, err
	case 8: // sum.create_contract_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(multisig.CreateContractMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_CreateContractMsg{msg}
		return true, err
	case 9: // sum.update_contract_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(multisig.UpdateContractMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_UpdateContractMsg{msg}
		return true, err
	case 10: // sum.set_validators_msg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(validators.SetValidatorsMsg)
		err := b.DecodeMessage(msg)
		m.Sum = &Tx_SetValidatorsMsg{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Tx_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Tx)
	// sum
	switch x := m.Sum.(type) {
	case *Tx_SendMsg:
		s := proto.Size(x.SendMsg)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_NewTokenMsg:
		s := proto.Size(x.NewTokenMsg)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_SetNameMsg:
		s := proto.Size(x.SetNameMsg)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_CreateEscrowMsg:
		s := proto.Size(x.CreateEscrowMsg)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_ReleaseEscrowMsg:
		s := proto.Size(x.ReleaseEscrowMsg)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_ReturnEscrowMsg:
		s := proto.Size(x.ReturnEscrowMsg)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_UpdateEscrowMsg:
		s := proto.Size(x.UpdateEscrowMsg)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_CreateContractMsg:
		s := proto.Size(x.CreateContractMsg)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_UpdateContractMsg:
		s := proto.Size(x.UpdateContractMsg)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Tx_SetValidatorsMsg:
		s := proto.Size(x.SetValidatorsMsg)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Tx)(nil), "app.Tx")
}
func (m *Tx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tx) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		nn1, err := m.Sum.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.Fees != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Fees.Size()))
		n2, err := m.Fees.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Signatures) > 0 {
		for _, msg := range m.Signatures {
			dAtA[i] = 0xaa
			i++
			dAtA[i] = 0x1
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Preimage) > 0 {
		dAtA[i] = 0xb2
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Preimage)))
		i += copy(dAtA[i:], m.Preimage)
	}
	if len(m.Multisig) > 0 {
		for _, b := range m.Multisig {
			dAtA[i] = 0xba
			i++
			dAtA[i] = 0x1
			i++
			i = encodeVarintCodec(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	return i, nil
}

func (m *Tx_SendMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.SendMsg != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.SendMsg.Size()))
		n3, err := m.SendMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *Tx_NewTokenMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.NewTokenMsg != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.NewTokenMsg.Size()))
		n4, err := m.NewTokenMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}
func (m *Tx_SetNameMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.SetNameMsg != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.SetNameMsg.Size()))
		n5, err := m.SetNameMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *Tx_CreateEscrowMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.CreateEscrowMsg != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.CreateEscrowMsg.Size()))
		n6, err := m.CreateEscrowMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}
func (m *Tx_ReleaseEscrowMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ReleaseEscrowMsg != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.ReleaseEscrowMsg.Size()))
		n7, err := m.ReleaseEscrowMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}
func (m *Tx_ReturnEscrowMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.ReturnEscrowMsg != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.ReturnEscrowMsg.Size()))
		n8, err := m.ReturnEscrowMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	return i, nil
}
func (m *Tx_UpdateEscrowMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.UpdateEscrowMsg != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.UpdateEscrowMsg.Size()))
		n9, err := m.UpdateEscrowMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	return i, nil
}
func (m *Tx_CreateContractMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.CreateContractMsg != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.CreateContractMsg.Size()))
		n10, err := m.CreateContractMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n10
	}
	return i, nil
}
func (m *Tx_UpdateContractMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.UpdateContractMsg != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.UpdateContractMsg.Size()))
		n11, err := m.UpdateContractMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n11
	}
	return i, nil
}
func (m *Tx_SetValidatorsMsg) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.SetValidatorsMsg != nil {
		dAtA[i] = 0x52
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.SetValidatorsMsg.Size()))
		n12, err := m.SetValidatorsMsg.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n12
	}
	return i, nil
}
func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Tx) Size() (n int) {
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	if m.Fees != nil {
		l = m.Fees.Size()
		n += 2 + l + sovCodec(uint64(l))
	}
	if len(m.Signatures) > 0 {
		for _, e := range m.Signatures {
			l = e.Size()
			n += 2 + l + sovCodec(uint64(l))
		}
	}
	l = len(m.Preimage)
	if l > 0 {
		n += 2 + l + sovCodec(uint64(l))
	}
	if len(m.Multisig) > 0 {
		for _, b := range m.Multisig {
			l = len(b)
			n += 2 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *Tx_SendMsg) Size() (n int) {
	var l int
	_ = l
	if m.SendMsg != nil {
		l = m.SendMsg.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_NewTokenMsg) Size() (n int) {
	var l int
	_ = l
	if m.NewTokenMsg != nil {
		l = m.NewTokenMsg.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_SetNameMsg) Size() (n int) {
	var l int
	_ = l
	if m.SetNameMsg != nil {
		l = m.SetNameMsg.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_CreateEscrowMsg) Size() (n int) {
	var l int
	_ = l
	if m.CreateEscrowMsg != nil {
		l = m.CreateEscrowMsg.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_ReleaseEscrowMsg) Size() (n int) {
	var l int
	_ = l
	if m.ReleaseEscrowMsg != nil {
		l = m.ReleaseEscrowMsg.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_ReturnEscrowMsg) Size() (n int) {
	var l int
	_ = l
	if m.ReturnEscrowMsg != nil {
		l = m.ReturnEscrowMsg.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_UpdateEscrowMsg) Size() (n int) {
	var l int
	_ = l
	if m.UpdateEscrowMsg != nil {
		l = m.UpdateEscrowMsg.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_CreateContractMsg) Size() (n int) {
	var l int
	_ = l
	if m.CreateContractMsg != nil {
		l = m.CreateContractMsg.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_UpdateContractMsg) Size() (n int) {
	var l int
	_ = l
	if m.UpdateContractMsg != nil {
		l = m.UpdateContractMsg.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}
func (m *Tx_SetValidatorsMsg) Size() (n int) {
	var l int
	_ = l
	if m.SetValidatorsMsg != nil {
		l = m.SetValidatorsMsg.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Tx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &cash.SendMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_SendMsg{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewTokenMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &namecoin.NewTokenMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_NewTokenMsg{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetNameMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &namecoin.SetWalletNameMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_SetNameMsg{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateEscrowMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &escrow.CreateEscrowMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_CreateEscrowMsg{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReleaseEscrowMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &escrow.ReleaseEscrowMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_ReleaseEscrowMsg{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnEscrowMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &escrow.ReturnEscrowMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_ReturnEscrowMsg{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateEscrowMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &escrow.UpdateEscrowPartiesMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_UpdateEscrowMsg{v}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateContractMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &multisig.CreateContractMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_CreateContractMsg{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateContractMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &multisig.UpdateContractMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_UpdateContractMsg{v}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetValidatorsMsg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &validators.SetValidatorsMsg{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Tx_SetValidatorsMsg{v}
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fees == nil {
				m.Fees = &cash.FeeInfo{}
			}
			if err := m.Fees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signatures = append(m.Signatures, &sigs.StdSignature{})
			if err := m.Signatures[len(m.Signatures)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 22:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Preimage", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Preimage = append(m.Preimage[:0], dAtA[iNdEx:postIndex]...)
			if m.Preimage == nil {
				m.Preimage = []byte{}
			}
			iNdEx = postIndex
		case 23:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multisig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Multisig = append(m.Multisig, make([]byte, postIndex-iNdEx))
			copy(m.Multisig[len(m.Multisig)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCodec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("app/codec.proto", fileDescriptorCodec) }

var fileDescriptorCodec = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x75, 0xeb, 0x8a, 0xbb, 0x69, 0x9b, 0x61, 0x2c, 0x2a, 0xa8, 0x2a, 0x9c, 0x2a,
	0xd0, 0x1c, 0x51, 0x8e, 0x48, 0x1c, 0x36, 0x0d, 0x0d, 0x69, 0x9b, 0x50, 0x3a, 0xe0, 0x58, 0x79,
	0xc9, 0x5b, 0x16, 0xd1, 0xd8, 0x91, 0xed, 0xb4, 0xfb, 0x18, 0x7c, 0x2c, 0x8e, 0x5c, 0xb8, 0xa3,
	0xf2, 0x45, 0x90, 0xed, 0xa4, 0x8d, 0x73, 0xa8, 0x76, 0xeb, 0xfb, 0xbf, 0xff, 0xff, 0xd7, 0x67,
	0xe7, 0x19, 0xed, 0xd3, 0x3c, 0x0f, 0x22, 0x1e, 0x43, 0x44, 0x72, 0xc1, 0x15, 0xc7, 0x2d, 0x9a,
	0xe7, 0xbd, 0xb7, 0x49, 0xaa, 0xee, 0x8b, 0x5b, 0x12, 0xf1, 0x2c, 0x48, 0xf9, 0xec, 0x84, 0x33,
	0x08, 0xe6, 0x40, 0x67, 0x10, 0x3c, 0x04, 0x11, 0x95, 0xf7, 0xf5, 0x44, 0xef, 0x64, 0x8d, 0x19,
	0x64, 0x24, 0xf8, 0xdc, 0xb1, 0x07, 0x6b, 0xec, 0x8c, 0x66, 0x10, 0xf1, 0x94, 0x39, 0x81, 0x75,
	0xc3, 0xc8, 0x34, 0x91, 0x8f, 0xa6, 0x67, 0xc5, 0x54, 0xa5, 0x32, 0x4d, 0x9c, 0xc0, 0xbb, 0x35,
	0x81, 0x19, 0x9d, 0xa6, 0x31, 0x55, 0x5c, 0x38, 0xff, 0xf1, 0xfa, 0x4f, 0x1b, 0x6d, 0xde, 0x3c,
	0xe0, 0x37, 0xa8, 0x23, 0x81, 0xc5, 0x93, 0x4c, 0x26, 0xbe, 0x37, 0xf0, 0x86, 0xdd, 0xd1, 0x1e,
	0xd1, 0x97, 0x43, 0xc6, 0xc0, 0xe2, 0x2b, 0x99, 0x5c, 0x6c, 0x84, 0x3b, 0xd2, 0xfe, 0xc4, 0x1f,
	0xd0, 0x1e, 0x83, 0xf9, 0x44, 0xf1, 0x1f, 0xc0, 0x4c, 0x60, 0xd3, 0x04, 0x8e, 0x48, 0x75, 0x62,
	0x72, 0x0d, 0xf3, 0x1b, 0xdd, 0xb5, 0xc1, 0x2e, 0x5b, 0x95, 0xf8, 0x23, 0xda, 0x95, 0xa0, 0x26,
	0xda, 0x6a, 0xb2, 0x2d, 0x93, 0xed, 0xad, 0xb2, 0x63, 0x50, 0xdf, 0xe9, 0x74, 0x0a, 0xea, 0x9a,
	0x66, 0x60, 0x01, 0x48, 0x2e, 0x2b, 0x7c, 0x8e, 0x0e, 0x23, 0x01, 0x54, 0xc1, 0xc4, 0x7e, 0x0e,
	0x03, 0xd9, 0x32, 0x90, 0x63, 0x62, 0x25, 0x72, 0x66, 0x0c, 0xe7, 0xa6, 0xb0, 0x84, 0xfd, 0xc8,
	0x95, 0xf0, 0x05, 0xc2, 0x02, 0xa6, 0x40, 0xa5, 0xc3, 0xd9, 0x36, 0x1c, 0xbf, 0xe2, 0x84, 0xd6,
	0x51, 0x07, 0x1d, 0x88, 0x86, 0xa6, 0x07, 0x12, 0xa0, 0x0a, 0xc1, 0xea, 0xa0, 0xb6, 0x3b, 0x50,
	0x68, 0x0c, 0xce, 0x40, 0xc2, 0x95, 0xf0, 0x25, 0x3a, 0x2c, 0xf2, 0xb8, 0x71, 0xae, 0x1d, 0x83,
	0xe9, 0x57, 0x98, 0xaf, 0xc6, 0x60, 0x33, 0x5f, 0xa8, 0x50, 0x29, 0xc8, 0x92, 0x56, 0xd4, 0x3a,
	0x9a, 0x76, 0x85, 0x9e, 0x96, 0xb7, 0x14, 0x71, 0xa6, 0x04, 0x8d, 0x94, 0xe1, 0x75, 0x0c, 0xef,
	0x05, 0xa9, 0x96, 0xa7, 0xbc, 0xa9, 0xb3, 0xd2, 0x63, 0x61, 0xe5, 0xfd, 0xd6, 0x44, 0x8d, 0x2b,
	0x87, 0x73, 0x70, 0x4f, 0x9a, 0x38, 0x3b, 0x60, 0x03, 0x57, 0x34, 0x45, 0x7c, 0x89, 0xb0, 0xde,
	0x81, 0xd5, 0x46, 0x1a, 0x1a, 0x32, 0xb4, 0x97, 0x64, 0x25, 0xeb, 0x5d, 0xf8, 0xb6, 0xac, 0xca,
	0x0f, 0x20, 0x1b, 0x1a, 0x7e, 0x85, 0xb6, 0xee, 0x00, 0xa4, 0xff, 0xac, 0xbe, 0xb6, 0x9f, 0x00,
	0x3e, 0xb3, 0x3b, 0x1e, 0x9a, 0x16, 0x1e, 0x21, 0x24, 0xd3, 0x84, 0x51, 0x55, 0x08, 0x90, 0xfe,
	0xd1, 0xa0, 0x35, 0xec, 0x8e, 0x30, 0xd1, 0xef, 0x8d, 0x8c, 0x55, 0x3c, 0xae, 0x5a, 0x61, 0xcd,
	0x85, 0x7b, 0xa8, 0x93, 0x0b, 0x48, 0x33, 0x9a, 0x80, 0xff, 0x7c, 0xe0, 0x0d, 0x77, 0xc3, 0x65,
	0xad, 0x7b, 0xd5, 0x99, 0xfd, 0xe3, 0x41, 0x4b, 0xf7, 0xaa, 0xfa, 0x74, 0x1b, 0xb5, 0x64, 0x91,
	0x9d, 0x1e, 0xfc, 0x5a, 0xf4, 0xbd, 0xdf, 0x8b, 0xbe, 0xf7, 0x77, 0xd1, 0xf7, 0x7e, 0xfe, 0xeb,
	0x6f, 0xdc, 0xb6, 0xcd, 0x83, 0x7b, 0xff, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x56, 0x07, 0xc6,
	0xa6, 0x04, 0x00, 0x00,
}
