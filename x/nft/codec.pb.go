// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/nft/codec.proto

/*
	Package nft is a generated protocol buffer package.

	It is generated from these files:
		x/nft/codec.proto

	It has these top-level messages:
		NonFungibleToken
		ActionApprovals
		Approval
		ApprovalOptions
		AddApprovalMsg
		RemoveApprovalMsg
*/
package nft

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

type NonFungibleToken struct {
	Id              []byte             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner           []byte             `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	ActionApprovals []*ActionApprovals `protobuf:"bytes,3,rep,name=actionApprovals" json:"actionApprovals,omitempty"`
}

func (m *NonFungibleToken) Reset()                    { *m = NonFungibleToken{} }
func (m *NonFungibleToken) String() string            { return proto.CompactTextString(m) }
func (*NonFungibleToken) ProtoMessage()               {}
func (*NonFungibleToken) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{0} }

func (m *NonFungibleToken) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *NonFungibleToken) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *NonFungibleToken) GetActionApprovals() []*ActionApprovals {
	if m != nil {
		return m.ActionApprovals
	}
	return nil
}

type ActionApprovals struct {
	Action    string      `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Approvals []*Approval `protobuf:"bytes,2,rep,name=approvals" json:"approvals,omitempty"`
}

func (m *ActionApprovals) Reset()                    { *m = ActionApprovals{} }
func (m *ActionApprovals) String() string            { return proto.CompactTextString(m) }
func (*ActionApprovals) ProtoMessage()               {}
func (*ActionApprovals) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{1} }

func (m *ActionApprovals) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *ActionApprovals) GetApprovals() []*Approval {
	if m != nil {
		return m.Approvals
	}
	return nil
}

type Approval struct {
	Address []byte          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Options ApprovalOptions `protobuf:"bytes,2,opt,name=options" json:"options"`
}

func (m *Approval) Reset()                    { *m = Approval{} }
func (m *Approval) String() string            { return proto.CompactTextString(m) }
func (*Approval) ProtoMessage()               {}
func (*Approval) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{2} }

func (m *Approval) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Approval) GetOptions() ApprovalOptions {
	if m != nil {
		return m.Options
	}
	return ApprovalOptions{}
}

type ApprovalOptions struct {
	UntilBlockHeight int64 `protobuf:"varint,1,opt,name=untilBlockHeight,proto3" json:"untilBlockHeight,omitempty"`
	Count            int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Immutable        bool  `protobuf:"varint,3,opt,name=immutable,proto3" json:"immutable,omitempty"`
}

func (m *ApprovalOptions) Reset()                    { *m = ApprovalOptions{} }
func (m *ApprovalOptions) String() string            { return proto.CompactTextString(m) }
func (*ApprovalOptions) ProtoMessage()               {}
func (*ApprovalOptions) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{3} }

func (m *ApprovalOptions) GetUntilBlockHeight() int64 {
	if m != nil {
		return m.UntilBlockHeight
	}
	return 0
}

func (m *ApprovalOptions) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ApprovalOptions) GetImmutable() bool {
	if m != nil {
		return m.Immutable
	}
	return false
}

type AddApprovalMsg struct {
	Id      []byte          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address []byte          `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Action  string          `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Options ApprovalOptions `protobuf:"bytes,4,opt,name=options" json:"options"`
}

func (m *AddApprovalMsg) Reset()                    { *m = AddApprovalMsg{} }
func (m *AddApprovalMsg) String() string            { return proto.CompactTextString(m) }
func (*AddApprovalMsg) ProtoMessage()               {}
func (*AddApprovalMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{4} }

func (m *AddApprovalMsg) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *AddApprovalMsg) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AddApprovalMsg) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AddApprovalMsg) GetOptions() ApprovalOptions {
	if m != nil {
		return m.Options
	}
	return ApprovalOptions{}
}

type RemoveApprovalMsg struct {
	Id      []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Action  string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (m *RemoveApprovalMsg) Reset()                    { *m = RemoveApprovalMsg{} }
func (m *RemoveApprovalMsg) String() string            { return proto.CompactTextString(m) }
func (*RemoveApprovalMsg) ProtoMessage()               {}
func (*RemoveApprovalMsg) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{5} }

func (m *RemoveApprovalMsg) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RemoveApprovalMsg) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *RemoveApprovalMsg) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func init() {
	proto.RegisterType((*NonFungibleToken)(nil), "nft.NonFungibleToken")
	proto.RegisterType((*ActionApprovals)(nil), "nft.ActionApprovals")
	proto.RegisterType((*Approval)(nil), "nft.Approval")
	proto.RegisterType((*ApprovalOptions)(nil), "nft.ApprovalOptions")
	proto.RegisterType((*AddApprovalMsg)(nil), "nft.AddApprovalMsg")
	proto.RegisterType((*RemoveApprovalMsg)(nil), "nft.RemoveApprovalMsg")
}
func (m *NonFungibleToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NonFungibleToken) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Owner) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Owner)))
		i += copy(dAtA[i:], m.Owner)
	}
	if len(m.ActionApprovals) > 0 {
		for _, msg := range m.ActionApprovals {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ActionApprovals) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActionApprovals) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Action) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	if len(m.Approvals) > 0 {
		for _, msg := range m.Approvals {
			dAtA[i] = 0x12
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Approval) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Approval) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintCodec(dAtA, i, uint64(m.Options.Size()))
	n1, err := m.Options.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *ApprovalOptions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApprovalOptions) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UntilBlockHeight != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.UntilBlockHeight))
	}
	if m.Count != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Count))
	}
	if m.Immutable {
		dAtA[i] = 0x18
		i++
		if m.Immutable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *AddApprovalMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddApprovalMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Action) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintCodec(dAtA, i, uint64(m.Options.Size()))
	n2, err := m.Options.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *RemoveApprovalMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveApprovalMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Address) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Address)))
		i += copy(dAtA[i:], m.Address)
	}
	if len(m.Action) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Action)))
		i += copy(dAtA[i:], m.Action)
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
func (m *NonFungibleToken) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.ActionApprovals) > 0 {
		for _, e := range m.ActionApprovals {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *ActionApprovals) Size() (n int) {
	var l int
	_ = l
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *Approval) Size() (n int) {
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = m.Options.Size()
	n += 1 + l + sovCodec(uint64(l))
	return n
}

func (m *ApprovalOptions) Size() (n int) {
	var l int
	_ = l
	if m.UntilBlockHeight != 0 {
		n += 1 + sovCodec(uint64(m.UntilBlockHeight))
	}
	if m.Count != 0 {
		n += 1 + sovCodec(uint64(m.Count))
	}
	if m.Immutable {
		n += 2
	}
	return n
}

func (m *AddApprovalMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Action)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = m.Options.Size()
	n += 1 + l + sovCodec(uint64(l))
	return n
}

func (m *RemoveApprovalMsg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Action)
	if l > 0 {
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
func (m *NonFungibleToken) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NonFungibleToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NonFungibleToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionApprovals", wireType)
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
			m.ActionApprovals = append(m.ActionApprovals, &ActionApprovals{})
			if err := m.ActionApprovals[len(m.ActionApprovals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *ActionApprovals) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ActionApprovals: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActionApprovals: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
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
			m.Approvals = append(m.Approvals, &Approval{})
			if err := m.Approvals[len(m.Approvals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *Approval) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Approval: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Approval: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
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
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *ApprovalOptions) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ApprovalOptions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApprovalOptions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UntilBlockHeight", wireType)
			}
			m.UntilBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UntilBlockHeight |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Immutable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Immutable = bool(v != 0)
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
func (m *AddApprovalMsg) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: AddApprovalMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddApprovalMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
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
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *RemoveApprovalMsg) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: RemoveApprovalMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveApprovalMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = string(dAtA[iNdEx:postIndex])
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

func init() { proto.RegisterFile("x/nft/codec.proto", fileDescriptorCodec) }

var fileDescriptorCodec = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcb, 0x6e, 0x9b, 0x40,
	0x14, 0x35, 0xe0, 0xfa, 0x71, 0xdb, 0xfa, 0x31, 0xb2, 0x2a, 0x54, 0x55, 0xd4, 0x62, 0x65, 0xb5,
	0x2a, 0x48, 0x6d, 0xd7, 0x91, 0xec, 0x45, 0x94, 0x4d, 0x12, 0x69, 0x94, 0x64, 0x91, 0x1d, 0x8f,
	0x01, 0x8f, 0x0c, 0x33, 0x04, 0x06, 0xc7, 0x9f, 0x90, 0x65, 0x3e, 0xcb, 0xcb, 0x7c, 0x41, 0x14,
	0x39, 0x3f, 0x12, 0x79, 0x00, 0x3f, 0x37, 0x59, 0x64, 0xc7, 0x79, 0x5c, 0xce, 0x3d, 0x17, 0xa0,
	0xbf, 0xb0, 0x59, 0x20, 0x6c, 0x8f, 0xfb, 0xc4, 0xb3, 0x92, 0x94, 0x0b, 0x8e, 0x34, 0x16, 0x88,
	0xef, 0x7f, 0x42, 0x2a, 0xa6, 0xb9, 0x6b, 0x79, 0x3c, 0xb6, 0x43, 0x1e, 0x72, 0x5b, 0x6a, 0x6e,
	0x1e, 0x48, 0x24, 0x81, 0x7c, 0x2a, 0x66, 0xcc, 0x05, 0xf4, 0x2e, 0x38, 0x3b, 0xcd, 0x59, 0x48,
	0xdd, 0x88, 0x5c, 0xf1, 0x19, 0x61, 0xa8, 0x03, 0x2a, 0xf5, 0x75, 0x65, 0xa8, 0x8c, 0xbe, 0x60,
	0x95, 0xfa, 0x68, 0x00, 0x9f, 0xf8, 0x3d, 0x23, 0xa9, 0xae, 0x4a, 0xaa, 0x00, 0xe8, 0x04, 0xba,
	0x8e, 0x27, 0x28, 0x67, 0xe3, 0x24, 0x49, 0xf9, 0xdc, 0x89, 0x32, 0x5d, 0x1b, 0x6a, 0xa3, 0xcf,
	0x7f, 0x07, 0x16, 0x0b, 0x84, 0x35, 0xde, 0xd7, 0xf0, 0xa1, 0xd9, 0xbc, 0x81, 0xee, 0x81, 0x07,
	0x7d, 0x83, 0x46, 0xe1, 0x92, 0xe1, 0x6d, 0x5c, 0x22, 0xf4, 0x1b, 0xda, 0xce, 0x26, 0x44, 0x95,
	0x21, 0x5f, 0x8b, 0x90, 0x92, 0xc5, 0x5b, 0xdd, 0xbc, 0x85, 0x56, 0x45, 0x23, 0x1d, 0x9a, 0x8e,
	0xef, 0xa7, 0x24, 0xcb, 0xca, 0x3a, 0x15, 0x44, 0xff, 0xa1, 0xc9, 0x93, 0xf5, 0xcb, 0x33, 0xd9,
	0x6a, 0xb3, 0x75, 0x39, 0x79, 0x59, 0x68, 0x93, 0xfa, 0xf2, 0xf9, 0x67, 0x0d, 0x57, 0x56, 0xf3,
	0x0e, 0xba, 0x07, 0x0e, 0xf4, 0x0b, 0x7a, 0x39, 0x13, 0x34, 0x9a, 0x44, 0xdc, 0x9b, 0x9d, 0x11,
	0x1a, 0x4e, 0x85, 0xcc, 0xd2, 0xf0, 0x11, 0xbf, 0x3e, 0xa4, 0xc7, 0x73, 0x26, 0x64, 0xa4, 0x86,
	0x0b, 0x80, 0x7e, 0x40, 0x9b, 0xc6, 0x71, 0x2e, 0x1c, 0x37, 0x22, 0xba, 0x36, 0x54, 0x46, 0x2d,
	0xbc, 0x25, 0xcc, 0x07, 0x05, 0x3a, 0x63, 0xdf, 0xaf, 0x62, 0xcf, 0xb3, 0xf0, 0xe8, 0xfb, 0xec,
	0xb4, 0x54, 0xf7, 0x5b, 0x6e, 0x0f, 0xaa, 0xed, 0x1d, 0x74, 0xa7, 0x7d, 0xfd, 0xfd, 0xed, 0xaf,
	0xa1, 0x8f, 0x49, 0xcc, 0xe7, 0xe4, 0x43, 0x97, 0x99, 0xf4, 0x96, 0x2b, 0x43, 0x79, 0x5a, 0x19,
	0xca, 0xcb, 0xca, 0x50, 0x1e, 0x5f, 0x8d, 0x9a, 0xdb, 0x90, 0xff, 0xe6, 0xbf, 0xb7, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x16, 0x34, 0x15, 0x81, 0xe4, 0x02, 0x00, 0x00,
}
