// Code generated by protoc-gen-go. DO NOT EDIT.
// source: processor_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type CPU struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NumberCores          uint32   `protobuf:"varint,3,opt,name=number_cores,json=numberCores,proto3" json:"number_cores,omitempty"`
	NumberThreads        uint32   `protobuf:"varint,4,opt,name=number_threads,json=numberThreads,proto3" json:"number_threads,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,5,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,6,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPU) Reset()         { *m = CPU{} }
func (m *CPU) String() string { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()    {}
func (*CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_466578cecc6db379, []int{0}
}

func (m *CPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPU.Unmarshal(m, b)
}
func (m *CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPU.Marshal(b, m, deterministic)
}
func (m *CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPU.Merge(m, src)
}
func (m *CPU) XXX_Size() int {
	return xxx_messageInfo_CPU.Size(m)
}
func (m *CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CPU proto.InternalMessageInfo

func (m *CPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CPU) GetNumberCores() uint32 {
	if m != nil {
		return m.NumberCores
	}
	return 0
}

func (m *CPU) GetNumberThreads() uint32 {
	if m != nil {
		return m.NumberThreads
	}
	return 0
}

func (m *CPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *CPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

type GPU struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,3,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,4,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	Memory               *Memory  `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GPU) Reset()         { *m = GPU{} }
func (m *GPU) String() string { return proto.CompactTextString(m) }
func (*GPU) ProtoMessage()    {}
func (*GPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_466578cecc6db379, []int{1}
}

func (m *GPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GPU.Unmarshal(m, b)
}
func (m *GPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GPU.Marshal(b, m, deterministic)
}
func (m *GPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPU.Merge(m, src)
}
func (m *GPU) XXX_Size() int {
	return xxx_messageInfo_GPU.Size(m)
}
func (m *GPU) XXX_DiscardUnknown() {
	xxx_messageInfo_GPU.DiscardUnknown(m)
}

var xxx_messageInfo_GPU proto.InternalMessageInfo

func (m *GPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *GPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *GPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

func (m *GPU) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterType((*CPU)(nil), "CPU")
	proto.RegisterType((*GPU)(nil), "GPU")
}

func init() { proto.RegisterFile("processor_message.proto", fileDescriptor_466578cecc6db379) }

var fileDescriptor_466578cecc6db379 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x14, 0x84, 0xe5, 0xa6, 0x4d, 0xc5, 0x2b, 0x65, 0xb0, 0x2a, 0x35, 0x62, 0x69, 0xa8, 0x84, 0xc8,
	0xe4, 0x01, 0xfe, 0x01, 0x1d, 0x32, 0x21, 0x55, 0x11, 0x2c, 0x2c, 0x91, 0xed, 0x5a, 0x71, 0x44,
	0x9d, 0x67, 0xd9, 0x41, 0x2a, 0x19, 0xf9, 0x3b, 0xfc, 0x49, 0x84, 0x9d, 0xa1, 0x0c, 0x0c, 0x6c,
	0xef, 0x7d, 0x77, 0xd2, 0x9d, 0x0e, 0xd6, 0xd6, 0xa1, 0x54, 0xde, 0xa3, 0xab, 0x8d, 0xf2, 0x9e,
	0x37, 0x8a, 0x59, 0x87, 0x3d, 0x5e, 0xaf, 0x8c, 0x32, 0xe8, 0x3e, 0x7e, 0xd3, 0xed, 0x17, 0x81,
	0x64, 0xb7, 0x7f, 0xa1, 0x2b, 0x98, 0x09, 0xc7, 0xbb, 0x43, 0x46, 0x72, 0x52, 0x5c, 0x54, 0xf1,
	0xa1, 0x14, 0xa6, 0x1d, 0x37, 0x2a, 0x9b, 0x04, 0x18, 0x6e, 0x7a, 0x03, 0x97, 0xdd, 0xbb, 0x11,
	0xca, 0xd5, 0x12, 0x9d, 0xf2, 0x59, 0x92, 0x93, 0x62, 0x59, 0x2d, 0x22, 0xdb, 0xfd, 0x20, 0x7a,
	0x0b, 0x57, 0xa3, 0xa5, 0xd7, 0x4e, 0xf1, 0x83, 0xcf, 0xa6, 0xc1, 0xb4, 0x8c, 0xf4, 0x39, 0x42,
	0xba, 0x86, 0xb9, 0x69, 0xbb, 0xba, 0xd1, 0x43, 0x36, 0xcb, 0x49, 0x41, 0xaa, 0xd4, 0xb4, 0x5d,
	0xa9, 0x87, 0x20, 0xf0, 0x53, 0x10, 0xd2, 0x51, 0xe0, 0xa7, 0x52, 0x0f, 0xdb, 0x4f, 0x02, 0x49,
	0xf9, 0xaf, 0xb6, 0x67, 0x19, 0xc9, 0x5f, 0x19, 0xd3, 0xf3, 0x0c, 0xba, 0x81, 0x34, 0x2e, 0x15,
	0x4a, 0x2d, 0xee, 0xe7, 0xec, 0x29, 0xbc, 0xd5, 0x88, 0x1f, 0xef, 0x60, 0x23, 0xd1, 0xb0, 0xa6,
	0xed, 0x8f, 0x5c, 0xb0, 0x5e, 0x49, 0xed, 0xa5, 0x46, 0x3c, 0x32, 0x2b, 0x05, 0xe2, 0x1b, 0xb3,
	0x62, 0x4f, 0x5e, 0x27, 0x56, 0x88, 0x34, 0x4c, 0xfc, 0xf0, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x09,
	0x33, 0xd3, 0xf0, 0x93, 0x01, 0x00, 0x00,
}