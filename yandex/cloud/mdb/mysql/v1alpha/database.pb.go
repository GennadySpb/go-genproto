// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1alpha/database.proto

package mysql // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1alpha"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MySQL database.
type Database struct {
	// Required. Name of the database. 1-63 characters long.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. ID of the MySQL cluster.
	ClusterId            string   `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Database) Reset()         { *m = Database{} }
func (m *Database) String() string { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()    {}
func (*Database) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_736a63598b34ab9e, []int{0}
}
func (m *Database) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Database.Unmarshal(m, b)
}
func (m *Database) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Database.Marshal(b, m, deterministic)
}
func (dst *Database) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Database.Merge(dst, src)
}
func (m *Database) XXX_Size() int {
	return xxx_messageInfo_Database.Size(m)
}
func (m *Database) XXX_DiscardUnknown() {
	xxx_messageInfo_Database.DiscardUnknown(m)
}

var xxx_messageInfo_Database proto.InternalMessageInfo

func (m *Database) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Database) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

type DatabaseSpec struct {
	// Required. Name of the MySQL database.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DatabaseSpec) Reset()         { *m = DatabaseSpec{} }
func (m *DatabaseSpec) String() string { return proto.CompactTextString(m) }
func (*DatabaseSpec) ProtoMessage()    {}
func (*DatabaseSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_736a63598b34ab9e, []int{1}
}
func (m *DatabaseSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabaseSpec.Unmarshal(m, b)
}
func (m *DatabaseSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabaseSpec.Marshal(b, m, deterministic)
}
func (dst *DatabaseSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabaseSpec.Merge(dst, src)
}
func (m *DatabaseSpec) XXX_Size() int {
	return xxx_messageInfo_DatabaseSpec.Size(m)
}
func (m *DatabaseSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabaseSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DatabaseSpec proto.InternalMessageInfo

func (m *DatabaseSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Database)(nil), "yandex.cloud.mdb.mysql.v1alpha.Database")
	proto.RegisterType((*DatabaseSpec)(nil), "yandex.cloud.mdb.mysql.v1alpha.DatabaseSpec")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1alpha/database.proto", fileDescriptor_database_736a63598b34ab9e)
}

var fileDescriptor_database_736a63598b34ab9e = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xad, 0x4c, 0xcc, 0x4b,
	0x49, 0xad, 0xd0, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0xcf, 0x4d, 0x49, 0xd2, 0xcf, 0xad, 0x2c,
	0x2e, 0xcc, 0xd1, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0x4f, 0x49, 0x2c, 0x49, 0x4c,
	0x4a, 0x2c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x83, 0x28, 0xd7, 0x03, 0x2b,
	0xd7, 0xcb, 0x4d, 0x49, 0xd2, 0x03, 0x2b, 0xd7, 0x83, 0x2a, 0x97, 0x92, 0x45, 0x31, 0xae, 0x2c,
	0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0x33, 0x3f, 0x0f, 0xa2, 0x5d, 0xc9, 0x96, 0x8b, 0xc3, 0x05,
	0x6a, 0xa0, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x98, 0x2d, 0x24, 0xcb, 0xc5, 0x95, 0x9c, 0x53, 0x5a, 0x5c, 0x92, 0x5a, 0x14, 0x9f, 0x99,
	0x22, 0xc1, 0x04, 0x96, 0xe1, 0x84, 0x8a, 0x78, 0xa6, 0x28, 0x39, 0x71, 0xf1, 0xc0, 0xb4, 0x07,
	0x17, 0xa4, 0x26, 0x0b, 0x19, 0x21, 0x1b, 0xe1, 0x24, 0xf7, 0xe2, 0xb8, 0x21, 0xe3, 0xa7, 0xe3,
	0x86, 0x7c, 0xd1, 0x89, 0xba, 0x55, 0x8e, 0xba, 0x51, 0x06, 0xba, 0x96, 0xf1, 0xba, 0xb1, 0x5a,
	0x5d, 0x27, 0x0c, 0x59, 0x6c, 0x6c, 0xcd, 0x8c, 0x21, 0x56, 0x38, 0x79, 0x45, 0x79, 0xa4, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x9c, 0xab, 0x0b, 0x71, 0x6e, 0x7a,
	0xbe, 0x6e, 0x7a, 0x6a, 0x1e, 0xd8, 0xa5, 0xfa, 0xf8, 0x83, 0xc5, 0x1a, 0xcc, 0x4b, 0x62, 0x03,
	0xab, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x23, 0x12, 0xf3, 0x3e, 0x45, 0x01, 0x00, 0x00,
}
