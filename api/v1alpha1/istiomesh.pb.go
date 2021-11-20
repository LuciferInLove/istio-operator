// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1alpha1/istiomesh.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/waynz0r/protobuf/gogoproto"
	proto "github.com/waynz0r/protobuf/proto"
	_ "github.com/waynz0r/protobuf/types"
	io "io"
	v1alpha1 "istio.io/api/mesh/v1alpha1"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	_ "k8s.io/api/core/v1"
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

// Mesh defines an Istio service mesh
//
// <!-- crd generation tags
// +cue-gen:IstioMesh:groupName:servicemesh.cisco.com
// +cue-gen:IstioMesh:version:v1alpha1
// +cue-gen:IstioMesh:storageVersion
// +cue-gen:IstioMesh:annotations:helm.sh/resource-policy=keep
// +cue-gen:IstioMesh:subresource:status
// +cue-gen:IstioMesh:scope:Namespaced
// +cue-gen:IstioMesh:resource:shortNames="im,imesh",plural="istiomeshes"
// +cue-gen:IstioMesh:printerColumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +cue-gen:IstioMesh:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type IstioMeshSpec struct {
	Config               *v1alpha1.MeshConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IstioMeshSpec) Reset()         { *m = IstioMeshSpec{} }
func (m *IstioMeshSpec) String() string { return proto.CompactTextString(m) }
func (*IstioMeshSpec) ProtoMessage()    {}
func (*IstioMeshSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b190aa132b1cfc9, []int{0}
}
func (m *IstioMeshSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioMeshSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioMeshSpec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioMeshSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioMeshSpec.Merge(m, src)
}
func (m *IstioMeshSpec) XXX_Size() int {
	return m.Size()
}
func (m *IstioMeshSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioMeshSpec.DiscardUnknown(m)
}

var xxx_messageInfo_IstioMeshSpec proto.InternalMessageInfo

func (m *IstioMeshSpec) GetConfig() *v1alpha1.MeshConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type IstioMeshStatus struct {
	// Reconciliation status of the Istio mesh
	Status ConfigState `protobuf:"varint,1,opt,name=status,proto3,enum=istio_operator.v2.api.v1alpha1.ConfigState" json:"status,omitempty"`
	// Reconciliation error message if any
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioMeshStatus) Reset()         { *m = IstioMeshStatus{} }
func (m *IstioMeshStatus) String() string { return proto.CompactTextString(m) }
func (*IstioMeshStatus) ProtoMessage()    {}
func (*IstioMeshStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b190aa132b1cfc9, []int{1}
}
func (m *IstioMeshStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IstioMeshStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IstioMeshStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IstioMeshStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioMeshStatus.Merge(m, src)
}
func (m *IstioMeshStatus) XXX_Size() int {
	return m.Size()
}
func (m *IstioMeshStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioMeshStatus.DiscardUnknown(m)
}

var xxx_messageInfo_IstioMeshStatus proto.InternalMessageInfo

func (m *IstioMeshStatus) GetStatus() ConfigState {
	if m != nil {
		return m.Status
	}
	return ConfigState_Unspecified
}

func (m *IstioMeshStatus) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*IstioMeshSpec)(nil), "istio_operator.v2.api.v1alpha1.IstioMeshSpec")
	proto.RegisterType((*IstioMeshStatus)(nil), "istio_operator.v2.api.v1alpha1.IstioMeshStatus")
}

func init() { proto.RegisterFile("api/v1alpha1/istiomesh.proto", fileDescriptor_3b190aa132b1cfc9) }

var fileDescriptor_3b190aa132b1cfc9 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3d, 0x6b, 0xf3, 0x30,
	0x10, 0xc7, 0xf1, 0x33, 0x04, 0x1e, 0xf5, 0x0d, 0x4c, 0x87, 0x34, 0x14, 0x27, 0x78, 0x0a, 0x94,
	0x4a, 0xc4, 0xa5, 0xb4, 0x73, 0xb3, 0xb4, 0x43, 0x96, 0x74, 0xeb, 0x12, 0x64, 0xe7, 0x22, 0x8b,
	0xda, 0x3e, 0x21, 0xc9, 0x2e, 0xe4, 0x13, 0x76, 0xec, 0x47, 0x28, 0xf9, 0x24, 0x45, 0xb2, 0x42,
	0x93, 0xa5, 0xdb, 0x71, 0xf7, 0xbb, 0xff, 0xff, 0x5e, 0xc8, 0x35, 0x57, 0x92, 0x75, 0x33, 0x5e,
	0xa9, 0x92, 0xcf, 0x98, 0x34, 0x56, 0x62, 0x0d, 0xa6, 0xa4, 0x4a, 0xa3, 0xc5, 0x38, 0xf1, 0x89,
	0x15, 0x2a, 0xd0, 0xdc, 0xa2, 0xa6, 0x5d, 0x46, 0xb9, 0x92, 0x74, 0xcf, 0x8f, 0x12, 0x81, 0x28,
	0x2a, 0x60, 0x9e, 0xce, 0xdb, 0x0d, 0xfb, 0xd0, 0x5c, 0x29, 0xd0, 0xa6, 0xef, 0x1f, 0x5d, 0x1d,
	0xa9, 0x17, 0x58, 0xd7, 0xd8, 0x84, 0xd2, 0xc8, 0xd9, 0x1c, 0xd6, 0x9a, 0x8d, 0x14, 0xa1, 0x76,
	0x29, 0x50, 0xa0, 0x0f, 0x99, 0x8b, 0x42, 0x76, 0x1c, 0xcc, 0x9c, 0xe6, 0x46, 0x42, 0xb5, 0x5e,
	0xe5, 0x50, 0xf2, 0x4e, 0xa2, 0x0e, 0x40, 0xfa, 0xfe, 0x68, 0xa8, 0x44, 0x0f, 0x14, 0xa8, 0x81,
	0x75, 0x33, 0x26, 0xa0, 0x71, 0xb3, 0xc3, 0xba, 0x67, 0xd2, 0x67, 0x72, 0xf6, 0xe2, 0x76, 0x5a,
	0x80, 0x29, 0x5f, 0x15, 0x14, 0xf1, 0x03, 0x19, 0xf4, 0xde, 0xc3, 0x68, 0x12, 0x4d, 0x4f, 0xb2,
	0x31, 0xf5, 0x3b, 0x53, 0x7f, 0x85, 0xfd, 0x78, 0xd4, 0xe1, 0x73, 0x8f, 0x2d, 0x03, 0x9e, 0x6e,
	0xc9, 0xc5, 0xaf, 0x92, 0xe5, 0xb6, 0x35, 0xf1, 0x9c, 0x0c, 0x8c, 0x8f, 0xbc, 0xd6, 0x79, 0x76,
	0x43, 0xff, 0xbe, 0x1f, 0xed, 0x25, 0x5d, 0x37, 0x2c, 0x43, 0x6b, 0x9c, 0x92, 0x53, 0xd0, 0x1a,
	0xf5, 0x02, 0x8c, 0xe1, 0x02, 0x86, 0xff, 0x26, 0xd1, 0xf4, 0xff, 0xf2, 0x28, 0xf7, 0x34, 0xff,
	0xdc, 0x25, 0xd1, 0xd7, 0x2e, 0x89, 0xbe, 0x77, 0x49, 0xf4, 0x76, 0x2f, 0xa4, 0x2d, 0xdb, 0x9c,
	0x16, 0x58, 0xb3, 0x9c, 0x37, 0x5b, 0x2e, 0x8b, 0x0a, 0xdb, 0x75, 0xff, 0xcd, 0xdb, 0xbd, 0x39,
	0xeb, 0x32, 0x76, 0xf8, 0x8e, 0x7c, 0xe0, 0x2f, 0x72, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0x91,
	0x00, 0xa0, 0x78, 0x03, 0x02, 0x00, 0x00,
}

func (m *IstioMeshSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioMeshSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioMeshSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIstiomesh(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IstioMeshStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IstioMeshStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IstioMeshStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ErrorMessage) > 0 {
		i -= len(m.ErrorMessage)
		copy(dAtA[i:], m.ErrorMessage)
		i = encodeVarintIstiomesh(dAtA, i, uint64(len(m.ErrorMessage)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintIstiomesh(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIstiomesh(dAtA []byte, offset int, v uint64) int {
	offset -= sovIstiomesh(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IstioMeshSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovIstiomesh(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IstioMeshStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovIstiomesh(uint64(m.Status))
	}
	l = len(m.ErrorMessage)
	if l > 0 {
		n += 1 + l + sovIstiomesh(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIstiomesh(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIstiomesh(x uint64) (n int) {
	return sovIstiomesh(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IstioMeshSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIstiomesh
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
			return fmt.Errorf("proto: IstioMeshSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioMeshSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiomesh
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
				return ErrInvalidLengthIstiomesh
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIstiomesh
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &v1alpha1.MeshConfig{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIstiomesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIstiomesh
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
func (m *IstioMeshStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIstiomesh
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
			return fmt.Errorf("proto: IstioMeshStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IstioMeshStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiomesh
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ConfigState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIstiomesh
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
				return ErrInvalidLengthIstiomesh
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIstiomesh
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIstiomesh(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIstiomesh
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
func skipIstiomesh(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIstiomesh
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
					return 0, ErrIntOverflowIstiomesh
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
					return 0, ErrIntOverflowIstiomesh
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
				return 0, ErrInvalidLengthIstiomesh
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIstiomesh
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIstiomesh
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIstiomesh        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIstiomesh          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIstiomesh = fmt.Errorf("proto: unexpected end of group")
)
