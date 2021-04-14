// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: keyboard_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Keyboard_Layout int32

const (
	Keyboard_UNKNOWN Keyboard_Layout = 0
	Keyboard_QWERTY  Keyboard_Layout = 1
	Keyboard_QWERTZ  Keyboard_Layout = 2
	Keyboard_AZERTY  Keyboard_Layout = 3
)

var Keyboard_Layout_name = map[int32]string{
	0: "UNKNOWN",
	1: "QWERTY",
	2: "QWERTZ",
	3: "AZERTY",
}

var Keyboard_Layout_value = map[string]int32{
	"UNKNOWN": 0,
	"QWERTY":  1,
	"QWERTZ":  2,
	"AZERTY":  3,
}

func (x Keyboard_Layout) String() string {
	return proto.EnumName(Keyboard_Layout_name, int32(x))
}

func (Keyboard_Layout) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8fd4226b8635fe5d, []int{0, 0}
}

type Keyboard struct {
	Layout               Keyboard_Layout `protobuf:"varint,1,opt,name=layout,proto3,enum=study.proto.Keyboard_Layout" json:"layout,omitempty"`
	Blacklit             bool            `protobuf:"varint,2,opt,name=blacklit,proto3" json:"blacklit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Keyboard) Reset()         { *m = Keyboard{} }
func (m *Keyboard) String() string { return proto.CompactTextString(m) }
func (*Keyboard) ProtoMessage()    {}
func (*Keyboard) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fd4226b8635fe5d, []int{0}
}
func (m *Keyboard) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Keyboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Keyboard.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Keyboard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keyboard.Merge(m, src)
}
func (m *Keyboard) XXX_Size() int {
	return m.Size()
}
func (m *Keyboard) XXX_DiscardUnknown() {
	xxx_messageInfo_Keyboard.DiscardUnknown(m)
}

var xxx_messageInfo_Keyboard proto.InternalMessageInfo

func (m *Keyboard) GetLayout() Keyboard_Layout {
	if m != nil {
		return m.Layout
	}
	return Keyboard_UNKNOWN
}

func (m *Keyboard) GetBlacklit() bool {
	if m != nil {
		return m.Blacklit
	}
	return false
}

func init() {
	proto.RegisterEnum("study.proto.Keyboard_Layout", Keyboard_Layout_name, Keyboard_Layout_value)
	proto.RegisterType((*Keyboard)(nil), "study.proto.Keyboard")
}

func init() { proto.RegisterFile("keyboard_message.proto", fileDescriptor_8fd4226b8635fe5d) }

var fileDescriptor_8fd4226b8635fe5d = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x4e, 0xad, 0x4c,
	0xca, 0x4f, 0x2c, 0x4a, 0x89, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x2e, 0x29, 0x4d, 0xa9, 0x84, 0x70, 0x94, 0xa6, 0x33, 0x72, 0x71,
	0x78, 0x43, 0xd5, 0x09, 0x99, 0x70, 0xb1, 0xe5, 0x24, 0x56, 0xe6, 0x97, 0x96, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xe8, 0x21, 0x29, 0xd5, 0x83, 0x29, 0xd3, 0xf3, 0x01, 0xab, 0x09,
	0x82, 0xaa, 0x15, 0x92, 0xe2, 0xe2, 0x48, 0xca, 0x49, 0x4c, 0xce, 0xce, 0xc9, 0x2c, 0x91, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0xf3, 0x95, 0x2c, 0xb9, 0xd8, 0x20, 0xaa, 0x85, 0xb8, 0xb9,
	0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x84, 0xb8, 0xb8, 0xd8, 0x02, 0xc3,
	0x5d, 0x83, 0x42, 0x22, 0x05, 0x18, 0xe1, 0xec, 0x28, 0x01, 0x26, 0x10, 0xdb, 0x31, 0x0a, 0x2c,
	0xce, 0xec, 0x24, 0x72, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0xce, 0x78, 0x2c, 0xc7, 0x10, 0xc5, 0x54, 0x90, 0x94, 0xc4, 0x06, 0x76, 0x8b, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x89, 0xc9, 0x47, 0xab, 0xdd, 0x00, 0x00, 0x00,
}

func (m *Keyboard) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Keyboard) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Keyboard) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Blacklit {
		i--
		if m.Blacklit {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Layout != 0 {
		i = encodeVarintKeyboardMessage(dAtA, i, uint64(m.Layout))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeyboardMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeyboardMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Keyboard) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Layout != 0 {
		n += 1 + sovKeyboardMessage(uint64(m.Layout))
	}
	if m.Blacklit {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovKeyboardMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeyboardMessage(x uint64) (n int) {
	return sovKeyboardMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Keyboard) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeyboardMessage
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
			return fmt.Errorf("proto: Keyboard: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Keyboard: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Layout", wireType)
			}
			m.Layout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyboardMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Layout |= Keyboard_Layout(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blacklit", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyboardMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Blacklit = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipKeyboardMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeyboardMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipKeyboardMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeyboardMessage
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
					return 0, ErrIntOverflowKeyboardMessage
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
					return 0, ErrIntOverflowKeyboardMessage
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
				return 0, ErrInvalidLengthKeyboardMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeyboardMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeyboardMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeyboardMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeyboardMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeyboardMessage = fmt.Errorf("proto: unexpected end of group")
)