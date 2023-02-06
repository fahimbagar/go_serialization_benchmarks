// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: structdef-gogo.proto

package goserbench

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
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

type GogoProtoBufA struct {
	Name     string  `protobuf:"bytes,1,req,name=name" json:"name"`
	BirthDay int64   `protobuf:"varint,2,req,name=birthDay" json:"birthDay"`
	Phone    string  `protobuf:"bytes,3,req,name=phone" json:"phone"`
	Siblings int32   `protobuf:"varint,4,req,name=siblings" json:"siblings"`
	Spouse   bool    `protobuf:"varint,5,req,name=spouse" json:"spouse"`
	Money    float64 `protobuf:"fixed64,6,req,name=money" json:"money"`
}

func (m *GogoProtoBufA) Reset()         { *m = GogoProtoBufA{} }
func (m *GogoProtoBufA) String() string { return proto.CompactTextString(m) }
func (*GogoProtoBufA) ProtoMessage()    {}
func (*GogoProtoBufA) Descriptor() ([]byte, []int) {
	return fileDescriptor_445e4de7f4e06aee, []int{0}
}
func (m *GogoProtoBufA) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GogoProtoBufA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GogoProtoBufA.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GogoProtoBufA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GogoProtoBufA.Merge(m, src)
}
func (m *GogoProtoBufA) XXX_Size() int {
	return m.Size()
}
func (m *GogoProtoBufA) XXX_DiscardUnknown() {
	xxx_messageInfo_GogoProtoBufA.DiscardUnknown(m)
}

var xxx_messageInfo_GogoProtoBufA proto.InternalMessageInfo

func (m *GogoProtoBufA) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GogoProtoBufA) GetBirthDay() int64 {
	if m != nil {
		return m.BirthDay
	}
	return 0
}

func (m *GogoProtoBufA) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *GogoProtoBufA) GetSiblings() int32 {
	if m != nil {
		return m.Siblings
	}
	return 0
}

func (m *GogoProtoBufA) GetSpouse() bool {
	if m != nil {
		return m.Spouse
	}
	return false
}

func (m *GogoProtoBufA) GetMoney() float64 {
	if m != nil {
		return m.Money
	}
	return 0
}

func init() {
	proto.RegisterType((*GogoProtoBufA)(nil), "goserbench.GogoProtoBufA")
}

func init() { proto.RegisterFile("structdef-gogo.proto", fileDescriptor_445e4de7f4e06aee) }

var fileDescriptor_445e4de7f4e06aee = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x63, 0xf7, 0x47, 0xc5, 0x12, 0x4b, 0xc4, 0x60, 0x55, 0xc8, 0x58, 0x4c, 0x59, 0xda,
	0x3c, 0x00, 0x2c, 0x54, 0x48, 0xac, 0x88, 0x81, 0x81, 0xa5, 0x72, 0x82, 0xeb, 0x58, 0x34, 0xbe,
	0x91, 0xed, 0x0c, 0xe5, 0x29, 0x78, 0x27, 0x96, 0x8e, 0x1d, 0x99, 0x10, 0x4a, 0x5e, 0x04, 0xc5,
	0x24, 0xb4, 0xe3, 0xf9, 0xee, 0xfd, 0xce, 0x70, 0xc8, 0x85, 0xf3, 0xb6, 0xce, 0xfd, 0xab, 0xdc,
	0x2c, 0x14, 0x28, 0x58, 0x56, 0x16, 0x3c, 0xc4, 0x44, 0x81, 0x93, 0x36, 0x93, 0x26, 0x2f, 0xe6,
	0x0b, 0xa5, 0x7d, 0x51, 0x67, 0xcb, 0x1c, 0xca, 0xb4, 0x7b, 0x49, 0xc3, 0x4b, 0x56, 0x6f, 0x42,
	0x0a, 0x21, 0x3d, 0xaa, 0xd7, 0x9f, 0x88, 0x9c, 0x3f, 0x80, 0x82, 0xc7, 0x2e, 0xad, 0xea, 0xcd,
	0x5d, 0x4c, 0xc9, 0xd8, 0x88, 0x52, 0x52, 0xc4, 0x71, 0x72, 0xb6, 0x1a, 0xef, 0xbf, 0xaf, 0xa2,
	0xa7, 0x40, 0x62, 0x4e, 0x66, 0x99, 0xb6, 0xbe, 0xb8, 0x17, 0x3b, 0x8a, 0x39, 0x4e, 0x46, 0xfd,
	0xf5, 0x9f, 0xc6, 0x73, 0x32, 0xa9, 0x0a, 0x30, 0x92, 0x8e, 0x4e, 0xe4, 0x3f, 0xd4, 0xd9, 0x4e,
	0x67, 0x5b, 0x6d, 0x94, 0xa3, 0x63, 0x8e, 0x93, 0xc9, 0x60, 0x0f, 0x34, 0xbe, 0x24, 0x53, 0x57,
	0x41, 0xed, 0x24, 0x9d, 0x70, 0x9c, 0xcc, 0xfa, 0x7b, 0xcf, 0xba, 0xee, 0x12, 0x8c, 0xdc, 0xd1,
	0x29, 0xc7, 0x09, 0x1a, 0xba, 0x03, 0x5a, 0x3d, 0xef, 0x1b, 0x86, 0x0e, 0x0d, 0x43, 0x3f, 0x0d,
	0x43, 0x1f, 0x2d, 0x8b, 0x0e, 0x2d, 0x8b, 0xbe, 0x5a, 0x16, 0xbd, 0xdc, 0x9e, 0xcc, 0x21, 0xb6,
	0x32, 0xf7, 0x05, 0x94, 0xc2, 0xa5, 0x0a, 0xd6, 0x4e, 0x5a, 0x2d, 0xb6, 0xfa, 0x5d, 0x78, 0x0d,
	0x66, 0x1d, 0xc6, 0x2b, 0x85, 0x7d, 0x73, 0x37, 0xc7, 0x31, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x52, 0xf5, 0x1a, 0x3d, 0x6f, 0x01, 0x00, 0x00,
}

func (m *GogoProtoBufA) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GogoProtoBufA) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GogoProtoBufA) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= 8
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Money))))
	i--
	dAtA[i] = 0x31
	i--
	if m.Spouse {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x28
	i = encodeVarintStructdefGogo(dAtA, i, uint64(m.Siblings))
	i--
	dAtA[i] = 0x20
	i -= len(m.Phone)
	copy(dAtA[i:], m.Phone)
	i = encodeVarintStructdefGogo(dAtA, i, uint64(len(m.Phone)))
	i--
	dAtA[i] = 0x1a
	i = encodeVarintStructdefGogo(dAtA, i, uint64(m.BirthDay))
	i--
	dAtA[i] = 0x10
	i -= len(m.Name)
	copy(dAtA[i:], m.Name)
	i = encodeVarintStructdefGogo(dAtA, i, uint64(len(m.Name)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintStructdefGogo(dAtA []byte, offset int, v uint64) int {
	offset -= sovStructdefGogo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GogoProtoBufA) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovStructdefGogo(uint64(l))
	n += 1 + sovStructdefGogo(uint64(m.BirthDay))
	l = len(m.Phone)
	n += 1 + l + sovStructdefGogo(uint64(l))
	n += 1 + sovStructdefGogo(uint64(m.Siblings))
	n += 2
	n += 9
	return n
}

func sovStructdefGogo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStructdefGogo(x uint64) (n int) {
	return sovStructdefGogo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GogoProtoBufA) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStructdefGogo
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
			return fmt.Errorf("proto: GogoProtoBufA: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GogoProtoBufA: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructdefGogo
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
				return ErrInvalidLengthStructdefGogo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStructdefGogo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BirthDay", wireType)
			}
			m.BirthDay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructdefGogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BirthDay |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Phone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructdefGogo
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
				return ErrInvalidLengthStructdefGogo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStructdefGogo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Phone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Siblings", wireType)
			}
			m.Siblings = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructdefGogo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Siblings |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spouse", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructdefGogo
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
			m.Spouse = bool(v != 0)
			hasFields[0] |= uint64(0x00000010)
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Money", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Money = float64(math.Float64frombits(v))
			hasFields[0] |= uint64(0x00000020)
		default:
			iNdEx = preIndex
			skippy, err := skipStructdefGogo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStructdefGogo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("name")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("birthDay")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("phone")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("siblings")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("spouse")
	}
	if hasFields[0]&uint64(0x00000020) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("money")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStructdefGogo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStructdefGogo
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
					return 0, ErrIntOverflowStructdefGogo
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
					return 0, ErrIntOverflowStructdefGogo
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
				return 0, ErrInvalidLengthStructdefGogo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStructdefGogo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStructdefGogo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStructdefGogo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStructdefGogo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStructdefGogo = fmt.Errorf("proto: unexpected end of group")
)
