// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reedsolomon/pb/readsolomon.proto

package recovery_pb

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type PBNode struct {
	Proto    []byte    `protobuf:"bytes,1,opt,name=proto,proto3" json:"proto,omitempty"`
	Recovery []*PBLink `protobuf:"bytes,2,rep,name=recovery,proto3" json:"recovery,omitempty"`
}

func (m *PBNode) Reset()      { *m = PBNode{} }
func (*PBNode) ProtoMessage() {}
func (*PBNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d7ca2a520a34643, []int{0}
}
func (m *PBNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PBNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PBNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PBNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBNode.Merge(m, src)
}
func (m *PBNode) XXX_Size() int {
	return m.Size()
}
func (m *PBNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PBNode.DiscardUnknown(m)
}

var xxx_messageInfo_PBNode proto.InternalMessageInfo

func (m *PBNode) GetProto() []byte {
	if m != nil {
		return m.Proto
	}
	return nil
}

func (m *PBNode) GetRecovery() []*PBLink {
	if m != nil {
		return m.Recovery
	}
	return nil
}

type PBLink struct {
	Hash  []byte `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Size_ uint64 `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (m *PBLink) Reset()      { *m = PBLink{} }
func (*PBLink) ProtoMessage() {}
func (*PBLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d7ca2a520a34643, []int{1}
}
func (m *PBLink) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PBLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PBLink.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PBLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBLink.Merge(m, src)
}
func (m *PBLink) XXX_Size() int {
	return m.Size()
}
func (m *PBLink) XXX_DiscardUnknown() {
	xxx_messageInfo_PBLink.DiscardUnknown(m)
}

var xxx_messageInfo_PBLink proto.InternalMessageInfo

func (m *PBLink) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PBLink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PBLink) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func init() {
	proto.RegisterType((*PBNode)(nil), "recovery.pb.PBNode")
	proto.RegisterType((*PBLink)(nil), "recovery.pb.PBLink")
}

func init() { proto.RegisterFile("reedsolomon/pb/readsolomon.proto", fileDescriptor_4d7ca2a520a34643) }

var fileDescriptor_4d7ca2a520a34643 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4a, 0x4d, 0x4d,
	0x29, 0xce, 0xcf, 0xc9, 0xcf, 0xcd, 0xcf, 0xd3, 0x2f, 0x48, 0xd2, 0x2f, 0x4a, 0x4d, 0x84, 0x71,
	0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb8, 0x8b, 0x52, 0x93, 0xf3, 0xcb, 0x52, 0x8b, 0x2a,
	0xf5, 0x0a, 0x92, 0x94, 0xfc, 0xb9, 0xd8, 0x02, 0x9c, 0xfc, 0xf2, 0x53, 0x52, 0x85, 0x44, 0xb8,
	0x58, 0xc1, 0xf2, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x10, 0x8e, 0x90, 0x3e, 0x17, 0x07,
	0x4c, 0xb9, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0xb0, 0x1e, 0x92, 0x7e, 0xbd, 0x00, 0x27,
	0x9f, 0xcc, 0xbc, 0xec, 0x20, 0xb8, 0x22, 0x25, 0x17, 0x90, 0x81, 0x20, 0x31, 0x21, 0x21, 0x2e,
	0x16, 0x8f, 0xc4, 0xe2, 0x0c, 0xa8, 0x79, 0x60, 0x36, 0x48, 0xcc, 0x2f, 0x31, 0x37, 0x55, 0x82,
	0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x06, 0x89, 0x05, 0x67, 0x56, 0xa5, 0x4a, 0x30, 0x2b,
	0x30, 0x6a, 0xb0, 0x04, 0x81, 0xd9, 0x4e, 0x26, 0x17, 0x1e, 0xca, 0x31, 0xdc, 0x78, 0x28, 0xc7,
	0xf0, 0xe1, 0xa1, 0x1c, 0x63, 0xc3, 0x23, 0x39, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x17, 0x8f, 0xe4, 0x18, 0x3e, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92,
	0xd8, 0xc0, 0x6e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xad, 0xa0, 0x5a, 0x12, 0x04, 0x01,
	0x00, 0x00,
}

func (this *PBNode) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PBNode)
	if !ok {
		that2, ok := that.(PBNode)
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
	if !bytes.Equal(this.Proto, that1.Proto) {
		return false
	}
	if len(this.Recovery) != len(that1.Recovery) {
		return false
	}
	for i := range this.Recovery {
		if !this.Recovery[i].Equal(that1.Recovery[i]) {
			return false
		}
	}
	return true
}
func (this *PBLink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PBLink)
	if !ok {
		that2, ok := that.(PBLink)
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
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Size_ != that1.Size_ {
		return false
	}
	return true
}
func (this *PBNode) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&recovery_pb.PBNode{")
	s = append(s, "Proto: "+fmt.Sprintf("%#v", this.Proto)+",\n")
	if this.Recovery != nil {
		s = append(s, "Recovery: "+fmt.Sprintf("%#v", this.Recovery)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PBLink) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&recovery_pb.PBLink{")
	s = append(s, "Hash: "+fmt.Sprintf("%#v", this.Hash)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Size_: "+fmt.Sprintf("%#v", this.Size_)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringReadsolomon(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *PBNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PBNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PBNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Recovery) > 0 {
		for iNdEx := len(m.Recovery) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Recovery[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReadsolomon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Proto) > 0 {
		i -= len(m.Proto)
		copy(dAtA[i:], m.Proto)
		i = encodeVarintReadsolomon(dAtA, i, uint64(len(m.Proto)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PBLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PBLink) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PBLink) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Size_ != 0 {
		i = encodeVarintReadsolomon(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintReadsolomon(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintReadsolomon(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReadsolomon(dAtA []byte, offset int, v uint64) int {
	offset -= sovReadsolomon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PBNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proto)
	if l > 0 {
		n += 1 + l + sovReadsolomon(uint64(l))
	}
	if len(m.Recovery) > 0 {
		for _, e := range m.Recovery {
			l = e.Size()
			n += 1 + l + sovReadsolomon(uint64(l))
		}
	}
	return n
}

func (m *PBLink) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovReadsolomon(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovReadsolomon(uint64(l))
	}
	if m.Size_ != 0 {
		n += 1 + sovReadsolomon(uint64(m.Size_))
	}
	return n
}

func sovReadsolomon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReadsolomon(x uint64) (n int) {
	return sovReadsolomon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PBNode) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForRecovery := "[]*PBLink{"
	for _, f := range this.Recovery {
		repeatedStringForRecovery += strings.Replace(f.String(), "PBLink", "PBLink", 1) + ","
	}
	repeatedStringForRecovery += "}"
	s := strings.Join([]string{`&PBNode{`,
		`Proto:` + fmt.Sprintf("%v", this.Proto) + `,`,
		`Recovery:` + repeatedStringForRecovery + `,`,
		`}`,
	}, "")
	return s
}
func (this *PBLink) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PBLink{`,
		`Hash:` + fmt.Sprintf("%v", this.Hash) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Size_:` + fmt.Sprintf("%v", this.Size_) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringReadsolomon(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PBNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReadsolomon
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
			return fmt.Errorf("proto: PBNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PBNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proto", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadsolomon
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
				return ErrInvalidLengthReadsolomon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReadsolomon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proto = append(m.Proto[:0], dAtA[iNdEx:postIndex]...)
			if m.Proto == nil {
				m.Proto = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recovery", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadsolomon
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
				return ErrInvalidLengthReadsolomon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReadsolomon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recovery = append(m.Recovery, &PBLink{})
			if err := m.Recovery[len(m.Recovery)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReadsolomon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReadsolomon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReadsolomon
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
func (m *PBLink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReadsolomon
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
			return fmt.Errorf("proto: PBLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PBLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadsolomon
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
				return ErrInvalidLengthReadsolomon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthReadsolomon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadsolomon
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
				return ErrInvalidLengthReadsolomon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReadsolomon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReadsolomon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReadsolomon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthReadsolomon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthReadsolomon
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
func skipReadsolomon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReadsolomon
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
					return 0, ErrIntOverflowReadsolomon
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
					return 0, ErrIntOverflowReadsolomon
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
				return 0, ErrInvalidLengthReadsolomon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReadsolomon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReadsolomon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReadsolomon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReadsolomon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReadsolomon = fmt.Errorf("proto: unexpected end of group")
)
