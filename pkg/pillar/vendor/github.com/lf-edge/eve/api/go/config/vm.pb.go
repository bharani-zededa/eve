// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/vm.proto

package config

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

// For now we need to tell the device which virtualization mode
// to use. Later we might use a single one for all VMs (on any particular
// ISA). If we end up keeping this we should make the names be less
// tied to a particular hypervisor.
type VmMode int32

const (
	VmMode_PV     VmMode = 0
	VmMode_HVM    VmMode = 1
	VmMode_Filler VmMode = 2
	VmMode_FML    VmMode = 3
)

var VmMode_name = map[int32]string{
	0: "PV",
	1: "HVM",
	2: "Filler",
	3: "FML",
}

var VmMode_value = map[string]int32{
	"PV":     0,
	"HVM":    1,
	"Filler": 2,
	"FML":    3,
}

func (x VmMode) String() string {
	return proto.EnumName(VmMode_name, int32(x))
}

func (VmMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_96baa922d70f35be, []int{0}
}

type VmConfig struct {
	Kernel               string   `protobuf:"bytes,1,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Ramdisk              string   `protobuf:"bytes,2,opt,name=ramdisk,proto3" json:"ramdisk,omitempty"`
	Memory               uint32   `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Maxmem               uint32   `protobuf:"varint,4,opt,name=maxmem,proto3" json:"maxmem,omitempty"`
	Vcpus                uint32   `protobuf:"varint,5,opt,name=vcpus,proto3" json:"vcpus,omitempty"`
	Maxcpus              uint32   `protobuf:"varint,6,opt,name=maxcpus,proto3" json:"maxcpus,omitempty"`
	Rootdev              string   `protobuf:"bytes,7,opt,name=rootdev,proto3" json:"rootdev,omitempty"`
	Extraargs            string   `protobuf:"bytes,8,opt,name=extraargs,proto3" json:"extraargs,omitempty"`
	Bootloader           string   `protobuf:"bytes,9,opt,name=bootloader,proto3" json:"bootloader,omitempty"`
	Cpus                 string   `protobuf:"bytes,10,opt,name=cpus,proto3" json:"cpus,omitempty"`
	Devicetree           string   `protobuf:"bytes,11,opt,name=devicetree,proto3" json:"devicetree,omitempty"`
	Dtdev                []string `protobuf:"bytes,12,rep,name=dtdev,proto3" json:"dtdev,omitempty"`
	Irqs                 []uint32 `protobuf:"varint,13,rep,packed,name=irqs,proto3" json:"irqs,omitempty"`
	Iomem                []string `protobuf:"bytes,14,rep,name=iomem,proto3" json:"iomem,omitempty"`
	VirtualizationMode   VmMode   `protobuf:"varint,15,opt,name=virtualizationMode,proto3,enum=VmMode" json:"virtualizationMode,omitempty"`
	EnableVnc            bool     `protobuf:"varint,16,opt,name=enableVnc,proto3" json:"enableVnc,omitempty"`
	VncDisplay           uint32   `protobuf:"varint,17,opt,name=vncDisplay,proto3" json:"vncDisplay,omitempty"`
	VncPasswd            string   `protobuf:"bytes,18,opt,name=vncPasswd,proto3" json:"vncPasswd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmConfig) Reset()         { *m = VmConfig{} }
func (m *VmConfig) String() string { return proto.CompactTextString(m) }
func (*VmConfig) ProtoMessage()    {}
func (*VmConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_96baa922d70f35be, []int{0}
}

func (m *VmConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmConfig.Unmarshal(m, b)
}
func (m *VmConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmConfig.Marshal(b, m, deterministic)
}
func (m *VmConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmConfig.Merge(m, src)
}
func (m *VmConfig) XXX_Size() int {
	return xxx_messageInfo_VmConfig.Size(m)
}
func (m *VmConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VmConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VmConfig proto.InternalMessageInfo

func (m *VmConfig) GetKernel() string {
	if m != nil {
		return m.Kernel
	}
	return ""
}

func (m *VmConfig) GetRamdisk() string {
	if m != nil {
		return m.Ramdisk
	}
	return ""
}

func (m *VmConfig) GetMemory() uint32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *VmConfig) GetMaxmem() uint32 {
	if m != nil {
		return m.Maxmem
	}
	return 0
}

func (m *VmConfig) GetVcpus() uint32 {
	if m != nil {
		return m.Vcpus
	}
	return 0
}

func (m *VmConfig) GetMaxcpus() uint32 {
	if m != nil {
		return m.Maxcpus
	}
	return 0
}

func (m *VmConfig) GetRootdev() string {
	if m != nil {
		return m.Rootdev
	}
	return ""
}

func (m *VmConfig) GetExtraargs() string {
	if m != nil {
		return m.Extraargs
	}
	return ""
}

func (m *VmConfig) GetBootloader() string {
	if m != nil {
		return m.Bootloader
	}
	return ""
}

func (m *VmConfig) GetCpus() string {
	if m != nil {
		return m.Cpus
	}
	return ""
}

func (m *VmConfig) GetDevicetree() string {
	if m != nil {
		return m.Devicetree
	}
	return ""
}

func (m *VmConfig) GetDtdev() []string {
	if m != nil {
		return m.Dtdev
	}
	return nil
}

func (m *VmConfig) GetIrqs() []uint32 {
	if m != nil {
		return m.Irqs
	}
	return nil
}

func (m *VmConfig) GetIomem() []string {
	if m != nil {
		return m.Iomem
	}
	return nil
}

func (m *VmConfig) GetVirtualizationMode() VmMode {
	if m != nil {
		return m.VirtualizationMode
	}
	return VmMode_PV
}

func (m *VmConfig) GetEnableVnc() bool {
	if m != nil {
		return m.EnableVnc
	}
	return false
}

func (m *VmConfig) GetVncDisplay() uint32 {
	if m != nil {
		return m.VncDisplay
	}
	return 0
}

func (m *VmConfig) GetVncPasswd() string {
	if m != nil {
		return m.VncPasswd
	}
	return ""
}

func init() {
	proto.RegisterEnum("VmMode", VmMode_name, VmMode_value)
	proto.RegisterType((*VmConfig)(nil), "VmConfig")
}

func init() {
	proto.RegisterFile("config/vm.proto", fileDescriptor_96baa922d70f35be)
}

var fileDescriptor_96baa922d70f35be = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x40, 0x49, 0xd3, 0x66, 0x77, 0x07, 0xb6, 0x0d, 0x16, 0x42, 0x3e, 0x20, 0x88, 0x10, 0x87,
	0x08, 0x89, 0x44, 0x82, 0x03, 0x77, 0x40, 0x85, 0x03, 0x2b, 0x55, 0x39, 0xe4, 0xc0, 0xcd, 0x6b,
	0x4f, 0x83, 0x55, 0x3b, 0x5e, 0x9c, 0x0f, 0xb6, 0xfd, 0xef, 0x48, 0xc8, 0xe3, 0x2d, 0xdb, 0x03,
	0x37, 0xbf, 0x37, 0x33, 0x9e, 0xf1, 0xc8, 0x70, 0x21, 0x5d, 0x7f, 0xad, 0xbb, 0x7a, 0xb6, 0xd5,
	0xce, 0xbb, 0xd1, 0xbd, 0xfe, 0x93, 0xc2, 0xb2, 0xb5, 0x9f, 0xc9, 0xb2, 0xe7, 0x90, 0xdd, 0xa0,
	0xef, 0xd1, 0xf0, 0xa4, 0x48, 0xca, 0x55, 0x73, 0x20, 0xc6, 0x61, 0xe1, 0x85, 0x55, 0x7a, 0xb8,
	0xe1, 0x27, 0x14, 0xb8, 0xc7, 0x50, 0x61, 0xd1, 0x3a, 0x7f, 0xcb, 0xd3, 0x22, 0x29, 0xd7, 0xcd,
	0x81, 0xc8, 0x8b, 0xbd, 0x45, 0xcb, 0x4f, 0x0f, 0x9e, 0x88, 0x3d, 0x83, 0xb3, 0x59, 0xee, 0xa6,
	0x81, 0x9f, 0x91, 0x8e, 0x10, 0xee, 0xb7, 0x62, 0x4f, 0x3e, 0x23, 0x7f, 0x8f, 0xd4, 0xd9, 0xb9,
	0x51, 0xe1, 0xcc, 0x17, 0x87, 0xce, 0x11, 0xd9, 0x0b, 0x58, 0xe1, 0x7e, 0xf4, 0x42, 0xf8, 0x6e,
	0xe0, 0x4b, 0x8a, 0x1d, 0x05, 0x7b, 0x09, 0xb0, 0x75, 0x6e, 0x34, 0x4e, 0x28, 0xf4, 0x7c, 0x45,
	0xe1, 0x07, 0x86, 0x31, 0x38, 0xa5, 0x76, 0x40, 0x11, 0x3a, 0x87, 0x1a, 0x85, 0xb3, 0x96, 0x38,
	0x7a, 0x44, 0xfe, 0x38, 0xd6, 0x1c, 0x4d, 0x98, 0x5d, 0xd1, 0x24, 0x4f, 0x8a, 0xb4, 0x5c, 0x35,
	0x11, 0xc2, 0x4d, 0xda, 0xff, 0x1a, 0xf8, 0xba, 0x48, 0xcb, 0x75, 0x43, 0xe7, 0x90, 0xa9, 0x5d,
	0x78, 0xfc, 0x79, 0xcc, 0x24, 0x60, 0x1f, 0x81, 0xcd, 0xda, 0x8f, 0x93, 0x30, 0xfa, 0x4e, 0x8c,
	0xda, 0xf5, 0x1b, 0xa7, 0x90, 0x5f, 0x14, 0x49, 0x79, 0xfe, 0x7e, 0x51, 0xb5, 0x36, 0x60, 0xf3,
	0x9f, 0x14, 0x7a, 0x6a, 0x2f, 0xb6, 0x06, 0xdb, 0x5e, 0xf2, 0xbc, 0x48, 0xca, 0x65, 0x73, 0x14,
	0x61, 0xec, 0xb9, 0x97, 0x5f, 0xf4, 0xb0, 0x33, 0xe2, 0x96, 0x3f, 0xa5, 0xfd, 0x3d, 0x30, 0xa1,
	0x7a, 0xee, 0xe5, 0x95, 0x18, 0x86, 0xdf, 0x8a, 0xb3, 0xb8, 0xa8, 0x7f, 0xe2, 0x6d, 0x05, 0x59,
	0xec, 0xcc, 0x32, 0x38, 0xb9, 0x6a, 0xf3, 0x47, 0x6c, 0x01, 0xe9, 0xb7, 0x76, 0x93, 0x27, 0x0c,
	0x20, 0xbb, 0xd4, 0xc6, 0xa0, 0xcf, 0x4f, 0x82, 0xbc, 0xdc, 0x7c, 0xcf, 0xd3, 0x4f, 0x5f, 0xe1,
	0x95, 0x74, 0xb6, 0xba, 0x43, 0x85, 0x4a, 0x54, 0xd2, 0xb8, 0x49, 0x55, 0xd3, 0x80, 0x3e, 0x6c,
	0x29, 0x7e, 0xa9, 0x1f, 0x6f, 0x3a, 0x3d, 0xfe, 0x9c, 0xb6, 0x95, 0x74, 0xb6, 0x36, 0xd7, 0xef,
	0x50, 0x75, 0x58, 0xe3, 0x8c, 0xb5, 0xd8, 0xe9, 0xba, 0x73, 0x75, 0xfc, 0x81, 0xdb, 0x8c, 0x92,
	0x3f, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x66, 0xae, 0x2d, 0xe5, 0x92, 0x02, 0x00, 0x00,
}
