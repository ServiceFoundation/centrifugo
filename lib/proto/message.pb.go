// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		message.proto

	It has these top-level messages:
		Message
		ClientInfo
		Publication
		Join
		Leave
		Unsubscribe
*/
package proto

import proto1 "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MessageTypePublication MessageType = 0
	MessageTypeJoin        MessageType = 1
	MessageTypeLeave       MessageType = 2
	MessageTypeUnsubscribe MessageType = 3
)

var MessageType_name = map[int32]string{
	0: "PUBLICATION",
	1: "JOIN",
	2: "LEAVE",
	3: "UNSUBSCRIBE",
}
var MessageType_value = map[string]int32{
	"PUBLICATION": 0,
	"JOIN":        1,
	"LEAVE":       2,
	"UNSUBSCRIBE": 3,
}

func (x MessageType) String() string {
	return proto1.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

type Message struct {
	Type    MessageType `protobuf:"varint,1,opt,name=Type,proto3,enum=proto.MessageType" json:"t,omitempty"`
	Channel string      `protobuf:"bytes,3,opt,name=Channel,proto3" json:"c,omitempty"`
	Data    Raw         `protobuf:"bytes,4,opt,name=Data,proto3,customtype=Raw" json:"d"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto1.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

func (m *Message) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageTypePublication
}

func (m *Message) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type ClientInfo struct {
	User     string `protobuf:"bytes,1,opt,name=User,proto3" json:"user"`
	Client   string `protobuf:"bytes,2,opt,name=Client,proto3" json:"client"`
	ConnInfo Raw    `protobuf:"bytes,3,opt,name=ConnInfo,proto3,customtype=Raw" json:"conn_info,omitempty"`
	ChanInfo Raw    `protobuf:"bytes,4,opt,name=ChanInfo,proto3,customtype=Raw" json:"chan_info,omitempty"`
}

func (m *ClientInfo) Reset()                    { *m = ClientInfo{} }
func (m *ClientInfo) String() string            { return proto1.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()               {}
func (*ClientInfo) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{1} }

func (m *ClientInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ClientInfo) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

type Publication struct {
	UID  string      `protobuf:"bytes,1,opt,name=UID,proto3" json:"uid,omitempty"`
	Data Raw         `protobuf:"bytes,2,opt,name=Data,proto3,customtype=Raw" json:"data"`
	Info *ClientInfo `protobuf:"bytes,3,opt,name=Info" json:"info,omitempty"`
}

func (m *Publication) Reset()                    { *m = Publication{} }
func (m *Publication) String() string            { return proto1.CompactTextString(m) }
func (*Publication) ProtoMessage()               {}
func (*Publication) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{2} }

func (m *Publication) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Publication) GetInfo() *ClientInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type Join struct {
	Info ClientInfo `protobuf:"bytes,1,opt,name=Info" json:"data"`
}

func (m *Join) Reset()                    { *m = Join{} }
func (m *Join) String() string            { return proto1.CompactTextString(m) }
func (*Join) ProtoMessage()               {}
func (*Join) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{3} }

func (m *Join) GetInfo() ClientInfo {
	if m != nil {
		return m.Info
	}
	return ClientInfo{}
}

type Leave struct {
	Info ClientInfo `protobuf:"bytes,1,opt,name=Info" json:"data"`
}

func (m *Leave) Reset()                    { *m = Leave{} }
func (m *Leave) String() string            { return proto1.CompactTextString(m) }
func (*Leave) ProtoMessage()               {}
func (*Leave) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{4} }

func (m *Leave) GetInfo() ClientInfo {
	if m != nil {
		return m.Info
	}
	return ClientInfo{}
}

type Unsubscribe struct {
}

func (m *Unsubscribe) Reset()                    { *m = Unsubscribe{} }
func (m *Unsubscribe) String() string            { return proto1.CompactTextString(m) }
func (*Unsubscribe) ProtoMessage()               {}
func (*Unsubscribe) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{5} }

func init() {
	proto1.RegisterType((*Message)(nil), "proto.Message")
	proto1.RegisterType((*ClientInfo)(nil), "proto.ClientInfo")
	proto1.RegisterType((*Publication)(nil), "proto.Publication")
	proto1.RegisterType((*Join)(nil), "proto.Join")
	proto1.RegisterType((*Leave)(nil), "proto.Leave")
	proto1.RegisterType((*Unsubscribe)(nil), "proto.Unsubscribe")
	proto1.RegisterEnum("proto.MessageType", MessageType_name, MessageType_value)
}
func (this *Message) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Channel != that1.Channel {
		return false
	}
	if !this.Data.Equal(that1.Data) {
		return false
	}
	return true
}
func (this *ClientInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClientInfo)
	if !ok {
		that2, ok := that.(ClientInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.User != that1.User {
		return false
	}
	if this.Client != that1.Client {
		return false
	}
	if !this.ConnInfo.Equal(that1.ConnInfo) {
		return false
	}
	if !this.ChanInfo.Equal(that1.ChanInfo) {
		return false
	}
	return true
}
func (this *Publication) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Publication)
	if !ok {
		that2, ok := that.(Publication)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.UID != that1.UID {
		return false
	}
	if !this.Data.Equal(that1.Data) {
		return false
	}
	if !this.Info.Equal(that1.Info) {
		return false
	}
	return true
}
func (this *Join) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Join)
	if !ok {
		that2, ok := that.(Join)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Info.Equal(&that1.Info) {
		return false
	}
	return true
}
func (this *Leave) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Leave)
	if !ok {
		that2, ok := that.(Leave)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Info.Equal(&that1.Info) {
		return false
	}
	return true
}
func (this *Unsubscribe) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Unsubscribe)
	if !ok {
		that2, ok := that.(Unsubscribe)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Type))
	}
	if len(m.Channel) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Channel)))
		i += copy(dAtA[i:], m.Channel)
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintMessage(dAtA, i, uint64(m.Data.Size()))
	n1, err := m.Data.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *ClientInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.User)))
		i += copy(dAtA[i:], m.User)
	}
	if len(m.Client) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Client)))
		i += copy(dAtA[i:], m.Client)
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintMessage(dAtA, i, uint64(m.ConnInfo.Size()))
	n2, err := m.ConnInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x22
	i++
	i = encodeVarintMessage(dAtA, i, uint64(m.ChanInfo.Size()))
	n3, err := m.ChanInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *Publication) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Publication) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.UID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.UID)))
		i += copy(dAtA[i:], m.UID)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintMessage(dAtA, i, uint64(m.Data.Size()))
	n4, err := m.Data.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.Info != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Info.Size()))
		n5, err := m.Info.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *Join) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Join) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMessage(dAtA, i, uint64(m.Info.Size()))
	n6, err := m.Info.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func (m *Leave) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Leave) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMessage(dAtA, i, uint64(m.Info.Size()))
	n7, err := m.Info.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	return i, nil
}

func (m *Unsubscribe) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Unsubscribe) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedMessage(r randyMessage, easy bool) *Message {
	this := &Message{}
	this.Type = MessageType([]int32{0, 1, 2, 3}[r.Intn(4)])
	this.Channel = string(randStringMessage(r))
	v1 := NewPopulatedRaw(r)
	this.Data = *v1
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedClientInfo(r randyMessage, easy bool) *ClientInfo {
	this := &ClientInfo{}
	this.User = string(randStringMessage(r))
	this.Client = string(randStringMessage(r))
	v2 := NewPopulatedRaw(r)
	this.ConnInfo = *v2
	v3 := NewPopulatedRaw(r)
	this.ChanInfo = *v3
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedPublication(r randyMessage, easy bool) *Publication {
	this := &Publication{}
	this.UID = string(randStringMessage(r))
	v4 := NewPopulatedRaw(r)
	this.Data = *v4
	if r.Intn(10) != 0 {
		this.Info = NewPopulatedClientInfo(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedJoin(r randyMessage, easy bool) *Join {
	this := &Join{}
	v5 := NewPopulatedClientInfo(r, easy)
	this.Info = *v5
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedLeave(r randyMessage, easy bool) *Leave {
	this := &Leave{}
	v6 := NewPopulatedClientInfo(r, easy)
	this.Info = *v6
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedUnsubscribe(r randyMessage, easy bool) *Unsubscribe {
	this := &Unsubscribe{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyMessage interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMessage(r randyMessage) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMessage(r randyMessage) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RuneMessage(r)
	}
	return string(tmps)
}
func randUnrecognizedMessage(r randyMessage, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMessage(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMessage(dAtA []byte, r randyMessage, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMessage(dAtA, uint64(key))
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulateMessage(dAtA, uint64(v8))
	case 1:
		dAtA = encodeVarintPopulateMessage(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMessage(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMessage(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMessage(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMessage(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Message) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMessage(uint64(m.Type))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = m.Data.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func (m *ClientInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Client)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = m.ConnInfo.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.ChanInfo.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func (m *Publication) Size() (n int) {
	var l int
	_ = l
	l = len(m.UID)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = m.Data.Size()
	n += 1 + l + sovMessage(uint64(l))
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *Join) Size() (n int) {
	var l int
	_ = l
	l = m.Info.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func (m *Leave) Size() (n int) {
	var l int
	_ = l
	l = m.Info.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func (m *Unsubscribe) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (MessageType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *ClientInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: ClientInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Client", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Client = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnInfo", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConnInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChanInfo", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChanInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Publication) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Publication: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Publication: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &ClientInfo{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Join) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Join: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Join: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Leave) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Leave: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Leave: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func (m *Unsubscribe) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Unsubscribe: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Unsubscribe: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
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
				next, err := skipMessage(dAtA[start:])
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
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("message.proto", fileDescriptorMessage) }

var fileDescriptorMessage = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x6f, 0xd3, 0x4e,
	0x18, 0xc6, 0x73, 0x8d, 0x9b, 0xb6, 0xaf, 0x93, 0xd6, 0xbd, 0xfc, 0xf5, 0x57, 0x64, 0x15, 0xdb,
	0x32, 0x48, 0x04, 0x0a, 0xa9, 0xd4, 0x0e, 0x08, 0x81, 0x90, 0xea, 0x34, 0x83, 0xab, 0x90, 0x56,
	0x6e, 0xcd, 0x8a, 0x6c, 0xe7, 0x92, 0x58, 0x4a, 0xee, 0xa2, 0xd8, 0x06, 0xf5, 0x1b, 0xa0, 0x4c,
	0x48, 0x48, 0x6c, 0x99, 0x58, 0x58, 0xd8, 0xf9, 0x08, 0x9d, 0x10, 0x33, 0x83, 0x05, 0x61, 0xf3,
	0x27, 0x60, 0x44, 0x39, 0x9b, 0xd6, 0x85, 0xb2, 0x30, 0xd9, 0x77, 0xcf, 0xf3, 0x7b, 0x9f, 0xf7,
	0x7d, 0x6d, 0xa8, 0x8c, 0x48, 0x10, 0x38, 0x7d, 0xd2, 0x18, 0x4f, 0x58, 0xc8, 0xf0, 0x32, 0x7f,
	0xc8, 0xf7, 0xfb, 0x7e, 0x38, 0x88, 0xdc, 0x86, 0xc7, 0x46, 0x3b, 0x7d, 0xd6, 0x67, 0x3b, 0xfc,
	0xda, 0x8d, 0x7a, 0xfc, 0xc4, 0x0f, 0xfc, 0x2d, 0xa5, 0xf4, 0x37, 0x08, 0x56, 0x9e, 0xa6, 0x75,
	0xf0, 0x03, 0x10, 0x4e, 0xcf, 0xc6, 0xa4, 0x86, 0x34, 0x54, 0x5f, 0xdf, 0xc5, 0xa9, 0xa3, 0x91,
	0xa9, 0x0b, 0xc5, 0xd8, 0x48, 0x62, 0x55, 0x0c, 0xef, 0xb1, 0x91, 0x1f, 0x92, 0xd1, 0x38, 0x3c,
	0xb3, 0x38, 0x80, 0xef, 0xc0, 0x4a, 0x73, 0xe0, 0x50, 0x4a, 0x86, 0xb5, 0xa2, 0x86, 0xea, 0x6b,
	0xa9, 0xcf, 0xcb, 0xf9, 0x7e, 0xe9, 0xf8, 0x16, 0x08, 0x07, 0x4e, 0xe8, 0xd4, 0x04, 0x0d, 0xd5,
	0xcb, 0x86, 0x74, 0x1e, 0xab, 0x85, 0x2f, 0xb1, 0x5a, 0xb4, 0x9c, 0x97, 0x49, 0xac, 0xa2, 0xae,
	0xc5, 0x55, 0xfd, 0x13, 0x02, 0x68, 0x0e, 0x7d, 0x42, 0x43, 0x93, 0xf6, 0x18, 0xde, 0x02, 0xc1,
	0x0e, 0xc8, 0x84, 0x37, 0xb6, 0x66, 0xac, 0x26, 0xb1, 0x2a, 0x44, 0x01, 0x99, 0x58, 0xfc, 0x16,
	0xeb, 0x50, 0x4a, 0xbd, 0xb5, 0x25, 0xae, 0x43, 0x12, 0xab, 0x25, 0x8f, 0xdf, 0x58, 0x99, 0x82,
	0x9f, 0xc0, 0x6a, 0x93, 0x51, 0xba, 0xa8, 0xc6, 0x5b, 0x2c, 0x1b, 0xfa, 0xd5, 0xe8, 0xaa, 0xc7,
	0x28, 0x7d, 0xee, 0xd3, 0x1e, 0xcb, 0x75, 0x7d, 0xc1, 0x70, 0x7e, 0xe0, 0xa4, 0xbc, 0x70, 0x3d,
	0x3f, 0x70, 0xae, 0xe1, 0x33, 0x46, 0x7f, 0x8b, 0x40, 0x3c, 0x8e, 0xdc, 0xa1, 0xef, 0x39, 0xa1,
	0xcf, 0x28, 0xbe, 0x09, 0x45, 0xdb, 0x3c, 0xc8, 0x06, 0xda, 0x4c, 0x62, 0xb5, 0x12, 0xf9, 0xdd,
	0x1c, 0xb9, 0x50, 0xf1, 0xed, 0x6c, 0x57, 0x4b, 0x3c, 0xb0, 0x7a, 0x35, 0x50, 0xe8, 0x3a, 0xa1,
	0x93, 0xae, 0x0b, 0x3f, 0x04, 0xe1, 0x62, 0x32, 0x71, 0x77, 0x33, 0xfb, 0x70, 0x97, 0x0b, 0x34,
	0x70, 0x12, 0xab, 0xeb, 0xbf, 0x35, 0xc7, 0x11, 0xfd, 0x11, 0x08, 0x87, 0xcc, 0xa7, 0x78, 0x2f,
	0x2b, 0x81, 0xfe, 0x56, 0xa2, 0xbc, 0x88, 0xbf, 0xcc, 0xe5, 0xf0, 0x63, 0x58, 0x6e, 0x13, 0xe7,
	0x05, 0xf9, 0x37, 0xba, 0x02, 0xa2, 0x4d, 0x83, 0xc8, 0x0d, 0xbc, 0x89, 0xef, 0x92, 0xbb, 0x1f,
	0x10, 0x88, 0xb9, 0x7f, 0x0d, 0x6f, 0x83, 0x78, 0x6c, 0x1b, 0x6d, 0xb3, 0xb9, 0x7f, 0x6a, 0x1e,
	0x75, 0xa4, 0x82, 0x2c, 0x4f, 0x67, 0xda, 0xff, 0x39, 0x47, 0x7e, 0x9f, 0x37, 0x40, 0x38, 0x3c,
	0x32, 0x3b, 0x12, 0x92, 0xab, 0xd3, 0x99, 0xb6, 0x91, 0x73, 0xf1, 0xe9, 0x54, 0x58, 0x6e, 0xb7,
	0xf6, 0x9f, 0xb5, 0xa4, 0x25, 0xf9, 0xbf, 0xe9, 0x4c, 0x93, 0x72, 0x7a, 0x3a, 0xc0, 0x36, 0x88,
	0x76, 0xe7, 0xc4, 0x36, 0x4e, 0x9a, 0x96, 0x69, 0xb4, 0xa4, 0xe2, 0x1f, 0x61, 0xb9, 0x4e, 0x65,
	0xe1, 0xd5, 0x3b, 0xa5, 0x60, 0x6c, 0xfd, 0xf8, 0xa6, 0xa0, 0xf7, 0x73, 0x05, 0x7d, 0x9c, 0x2b,
	0xe8, 0x7c, 0xae, 0xa0, 0xcf, 0x73, 0x05, 0x7d, 0x9d, 0x2b, 0xe8, 0xf5, 0x77, 0xa5, 0xe0, 0x96,
	0xf8, 0x0a, 0xf6, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x9c, 0x01, 0xf9, 0xa5, 0x03, 0x00,
	0x00,
}
