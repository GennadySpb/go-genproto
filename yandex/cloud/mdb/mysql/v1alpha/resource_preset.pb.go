// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1alpha/resource_preset.proto

package mysql // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1alpha"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A preset of resources for hardware configuration of MySQL hosts.
type ResourcePreset struct {
	// ID of the resource preset.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// IDs of availability zones where the resource preset is available.
	ZoneIds []string `protobuf:"bytes,2,rep,name=zone_ids,json=zoneIds,proto3" json:"zone_ids,omitempty"`
	// Number of CPU cores for a MySQL host created with the preset.
	Cores int64 `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
	// RAM volume for a MySQL host created with the preset, in bytes.
	Memory               int64    `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourcePreset) Reset()         { *m = ResourcePreset{} }
func (m *ResourcePreset) String() string { return proto.CompactTextString(m) }
func (*ResourcePreset) ProtoMessage()    {}
func (*ResourcePreset) Descriptor() ([]byte, []int) {
	return fileDescriptor_resource_preset_5892b46c76e77bf2, []int{0}
}
func (m *ResourcePreset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourcePreset.Unmarshal(m, b)
}
func (m *ResourcePreset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourcePreset.Marshal(b, m, deterministic)
}
func (dst *ResourcePreset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcePreset.Merge(dst, src)
}
func (m *ResourcePreset) XXX_Size() int {
	return xxx_messageInfo_ResourcePreset.Size(m)
}
func (m *ResourcePreset) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcePreset.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcePreset proto.InternalMessageInfo

func (m *ResourcePreset) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResourcePreset) GetZoneIds() []string {
	if m != nil {
		return m.ZoneIds
	}
	return nil
}

func (m *ResourcePreset) GetCores() int64 {
	if m != nil {
		return m.Cores
	}
	return 0
}

func (m *ResourcePreset) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func init() {
	proto.RegisterType((*ResourcePreset)(nil), "yandex.cloud.mdb.mysql.v1alpha.ResourcePreset")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1alpha/resource_preset.proto", fileDescriptor_resource_preset_5892b46c76e77bf2)
}

var fileDescriptor_resource_preset_5892b46c76e77bf2 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xcf, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x71, 0xda, 0xea, 0xe9, 0x65, 0xb8, 0x21, 0x88, 0xc4, 0x45, 0x8a, 0x53, 0x97, 0x4b,
	0x10, 0xdd, 0xdc, 0x9c, 0xd4, 0x49, 0x32, 0xba, 0x1c, 0x4d, 0xde, 0xa3, 0x0d, 0x34, 0x7d, 0x35,
	0x69, 0xc5, 0xfa, 0xd7, 0x8b, 0x49, 0x67, 0xc7, 0x6f, 0xc8, 0x07, 0xde, 0x8f, 0x3d, 0xae, 0xed,
	0x08, 0xf8, 0xad, 0xec, 0x40, 0x0b, 0x28, 0x0f, 0x46, 0xf9, 0x35, 0x7e, 0x0e, 0xea, 0xeb, 0xbe,
	0x1d, 0xa6, 0xbe, 0x55, 0x01, 0x23, 0x2d, 0xc1, 0xe2, 0x69, 0x0a, 0x18, 0x71, 0x96, 0x53, 0xa0,
	0x99, 0xf8, 0x6d, 0x56, 0x32, 0x29, 0xe9, 0xc1, 0xc8, 0xa4, 0xe4, 0xa6, 0xee, 0x1c, 0x3b, 0xe8,
	0x0d, 0xbe, 0x27, 0xc7, 0x0f, 0xac, 0x74, 0x20, 0x8a, 0xba, 0x68, 0xf6, 0xba, 0x74, 0xc0, 0x6f,
	0xd8, 0xe5, 0x0f, 0x8d, 0x78, 0x72, 0x10, 0x45, 0x59, 0x57, 0xcd, 0x5e, 0x5f, 0xfc, 0xf5, 0x2b,
	0x44, 0x7e, 0xc5, 0xce, 0x2d, 0x05, 0x8c, 0xa2, 0xaa, 0x8b, 0xa6, 0xd2, 0x39, 0xf8, 0x35, 0xdb,
	0x79, 0xf4, 0x14, 0x56, 0x71, 0x96, 0x9e, 0xb7, 0x7a, 0x7e, 0xfb, 0x78, 0xe9, 0xdc, 0xdc, 0x2f,
	0x46, 0x5a, 0xf2, 0x2a, 0xdf, 0x75, 0xcc, 0x6b, 0x3a, 0x3a, 0x76, 0x38, 0xa6, 0x8b, 0xd5, 0xff,
	0x33, 0x9f, 0x52, 0x99, 0x5d, 0xfa, 0xfb, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xee, 0xb5, 0x15,
	0x18, 0x15, 0x01, 0x00, 0x00,
}
