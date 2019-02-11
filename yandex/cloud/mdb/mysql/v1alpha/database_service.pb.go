// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1alpha/database_service.proto

package mysql // import "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mysql/v1alpha"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/yandex-cloud/go-genproto/yandex/api"
import operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
import _ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetDatabaseRequest struct {
	// Required. ID of the MySQL cluster.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required.
	DatabaseName         string   `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDatabaseRequest) Reset()         { *m = GetDatabaseRequest{} }
func (m *GetDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*GetDatabaseRequest) ProtoMessage()    {}
func (*GetDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_service_da98f3645adba7f1, []int{0}
}
func (m *GetDatabaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDatabaseRequest.Unmarshal(m, b)
}
func (m *GetDatabaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDatabaseRequest.Marshal(b, m, deterministic)
}
func (dst *GetDatabaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDatabaseRequest.Merge(dst, src)
}
func (m *GetDatabaseRequest) XXX_Size() int {
	return xxx_messageInfo_GetDatabaseRequest.Size(m)
}
func (m *GetDatabaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDatabaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDatabaseRequest proto.InternalMessageInfo

func (m *GetDatabaseRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *GetDatabaseRequest) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

type ListDatabasesRequest struct {
	// Required. ID of the MySQL cluster.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The maximum number of results per page that should be returned. If the number of available
	// results is larger than `page_size`, the service returns a `next_page_token` that can be used
	// to get the next page of results in subsequent ListDatabases requests.
	// Acceptable values are 0 to 1000, inclusive. Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set `page_token` to the `next_page_token` returned by a previous ListDatabases
	// request to get the next page of results.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDatabasesRequest) Reset()         { *m = ListDatabasesRequest{} }
func (m *ListDatabasesRequest) String() string { return proto.CompactTextString(m) }
func (*ListDatabasesRequest) ProtoMessage()    {}
func (*ListDatabasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_service_da98f3645adba7f1, []int{1}
}
func (m *ListDatabasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDatabasesRequest.Unmarshal(m, b)
}
func (m *ListDatabasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDatabasesRequest.Marshal(b, m, deterministic)
}
func (dst *ListDatabasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDatabasesRequest.Merge(dst, src)
}
func (m *ListDatabasesRequest) XXX_Size() int {
	return xxx_messageInfo_ListDatabasesRequest.Size(m)
}
func (m *ListDatabasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDatabasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDatabasesRequest proto.InternalMessageInfo

func (m *ListDatabasesRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ListDatabasesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListDatabasesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListDatabasesResponse struct {
	// Requested list of MySQL clusters.
	Databases []*Database `protobuf:"bytes,1,rep,name=databases,proto3" json:"databases,omitempty"`
	// This token allows you to get the next page of results for ListDatabases requests,
	// if the number of results is larger than `page_size` specified in the request.
	// To get the next page, specify the value of `next_page_token` as a value for
	// the `page_token` parameter in the next ListDatabases request. Subsequent ListDatabases
	// requests will have their own `next_page_token` to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDatabasesResponse) Reset()         { *m = ListDatabasesResponse{} }
func (m *ListDatabasesResponse) String() string { return proto.CompactTextString(m) }
func (*ListDatabasesResponse) ProtoMessage()    {}
func (*ListDatabasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_service_da98f3645adba7f1, []int{2}
}
func (m *ListDatabasesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDatabasesResponse.Unmarshal(m, b)
}
func (m *ListDatabasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDatabasesResponse.Marshal(b, m, deterministic)
}
func (dst *ListDatabasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDatabasesResponse.Merge(dst, src)
}
func (m *ListDatabasesResponse) XXX_Size() int {
	return xxx_messageInfo_ListDatabasesResponse.Size(m)
}
func (m *ListDatabasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDatabasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDatabasesResponse proto.InternalMessageInfo

func (m *ListDatabasesResponse) GetDatabases() []*Database {
	if m != nil {
		return m.Databases
	}
	return nil
}

func (m *ListDatabasesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type CreateDatabaseRequest struct {
	// Required. ID of the MySQL cluster.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required.
	DatabaseSpec         *DatabaseSpec `protobuf:"bytes,2,opt,name=database_spec,json=databaseSpec,proto3" json:"database_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateDatabaseRequest) Reset()         { *m = CreateDatabaseRequest{} }
func (m *CreateDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDatabaseRequest) ProtoMessage()    {}
func (*CreateDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_service_da98f3645adba7f1, []int{3}
}
func (m *CreateDatabaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDatabaseRequest.Unmarshal(m, b)
}
func (m *CreateDatabaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDatabaseRequest.Marshal(b, m, deterministic)
}
func (dst *CreateDatabaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDatabaseRequest.Merge(dst, src)
}
func (m *CreateDatabaseRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDatabaseRequest.Size(m)
}
func (m *CreateDatabaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDatabaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDatabaseRequest proto.InternalMessageInfo

func (m *CreateDatabaseRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateDatabaseRequest) GetDatabaseSpec() *DatabaseSpec {
	if m != nil {
		return m.DatabaseSpec
	}
	return nil
}

type CreateDatabaseMetadata struct {
	// Required. ID of the MySQL cluster.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required. Name of the creating database.
	DatabaseName         string   `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDatabaseMetadata) Reset()         { *m = CreateDatabaseMetadata{} }
func (m *CreateDatabaseMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateDatabaseMetadata) ProtoMessage()    {}
func (*CreateDatabaseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_service_da98f3645adba7f1, []int{4}
}
func (m *CreateDatabaseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDatabaseMetadata.Unmarshal(m, b)
}
func (m *CreateDatabaseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDatabaseMetadata.Marshal(b, m, deterministic)
}
func (dst *CreateDatabaseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDatabaseMetadata.Merge(dst, src)
}
func (m *CreateDatabaseMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateDatabaseMetadata.Size(m)
}
func (m *CreateDatabaseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDatabaseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDatabaseMetadata proto.InternalMessageInfo

func (m *CreateDatabaseMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateDatabaseMetadata) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

type UpdateDatabaseRequest struct {
	// Required. ID of the MySQL cluster.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required. Name of the database to modify.
	DatabaseName         string                `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateDatabaseRequest) Reset()         { *m = UpdateDatabaseRequest{} }
func (m *UpdateDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDatabaseRequest) ProtoMessage()    {}
func (*UpdateDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_service_da98f3645adba7f1, []int{5}
}
func (m *UpdateDatabaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDatabaseRequest.Unmarshal(m, b)
}
func (m *UpdateDatabaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDatabaseRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateDatabaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDatabaseRequest.Merge(dst, src)
}
func (m *UpdateDatabaseRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDatabaseRequest.Size(m)
}
func (m *UpdateDatabaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDatabaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDatabaseRequest proto.InternalMessageInfo

func (m *UpdateDatabaseRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *UpdateDatabaseRequest) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

func (m *UpdateDatabaseRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type UpdateDatabaseMetadata struct {
	// Required. ID of the MySQL cluster.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required. Name of the database.
	DatabaseName         string   `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateDatabaseMetadata) Reset()         { *m = UpdateDatabaseMetadata{} }
func (m *UpdateDatabaseMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateDatabaseMetadata) ProtoMessage()    {}
func (*UpdateDatabaseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_service_da98f3645adba7f1, []int{6}
}
func (m *UpdateDatabaseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDatabaseMetadata.Unmarshal(m, b)
}
func (m *UpdateDatabaseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDatabaseMetadata.Marshal(b, m, deterministic)
}
func (dst *UpdateDatabaseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDatabaseMetadata.Merge(dst, src)
}
func (m *UpdateDatabaseMetadata) XXX_Size() int {
	return xxx_messageInfo_UpdateDatabaseMetadata.Size(m)
}
func (m *UpdateDatabaseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDatabaseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDatabaseMetadata proto.InternalMessageInfo

func (m *UpdateDatabaseMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *UpdateDatabaseMetadata) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

type DeleteDatabaseRequest struct {
	// Required. ID of the MySQL cluster.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required. Name of the database to delete.
	DatabaseName         string   `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDatabaseRequest) Reset()         { *m = DeleteDatabaseRequest{} }
func (m *DeleteDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDatabaseRequest) ProtoMessage()    {}
func (*DeleteDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_service_da98f3645adba7f1, []int{7}
}
func (m *DeleteDatabaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDatabaseRequest.Unmarshal(m, b)
}
func (m *DeleteDatabaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDatabaseRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteDatabaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDatabaseRequest.Merge(dst, src)
}
func (m *DeleteDatabaseRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDatabaseRequest.Size(m)
}
func (m *DeleteDatabaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDatabaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDatabaseRequest proto.InternalMessageInfo

func (m *DeleteDatabaseRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *DeleteDatabaseRequest) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

type DeleteDatabaseMetadata struct {
	// Required. ID of the MySQL cluster.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required. Name of the deleting database.
	DatabaseName         string   `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDatabaseMetadata) Reset()         { *m = DeleteDatabaseMetadata{} }
func (m *DeleteDatabaseMetadata) String() string { return proto.CompactTextString(m) }
func (*DeleteDatabaseMetadata) ProtoMessage()    {}
func (*DeleteDatabaseMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_database_service_da98f3645adba7f1, []int{8}
}
func (m *DeleteDatabaseMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDatabaseMetadata.Unmarshal(m, b)
}
func (m *DeleteDatabaseMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDatabaseMetadata.Marshal(b, m, deterministic)
}
func (dst *DeleteDatabaseMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDatabaseMetadata.Merge(dst, src)
}
func (m *DeleteDatabaseMetadata) XXX_Size() int {
	return xxx_messageInfo_DeleteDatabaseMetadata.Size(m)
}
func (m *DeleteDatabaseMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDatabaseMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDatabaseMetadata proto.InternalMessageInfo

func (m *DeleteDatabaseMetadata) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *DeleteDatabaseMetadata) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDatabaseRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.GetDatabaseRequest")
	proto.RegisterType((*ListDatabasesRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.ListDatabasesRequest")
	proto.RegisterType((*ListDatabasesResponse)(nil), "yandex.cloud.mdb.mysql.v1alpha.ListDatabasesResponse")
	proto.RegisterType((*CreateDatabaseRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.CreateDatabaseRequest")
	proto.RegisterType((*CreateDatabaseMetadata)(nil), "yandex.cloud.mdb.mysql.v1alpha.CreateDatabaseMetadata")
	proto.RegisterType((*UpdateDatabaseRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.UpdateDatabaseRequest")
	proto.RegisterType((*UpdateDatabaseMetadata)(nil), "yandex.cloud.mdb.mysql.v1alpha.UpdateDatabaseMetadata")
	proto.RegisterType((*DeleteDatabaseRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.DeleteDatabaseRequest")
	proto.RegisterType((*DeleteDatabaseMetadata)(nil), "yandex.cloud.mdb.mysql.v1alpha.DeleteDatabaseMetadata")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	// Returns the specified MySQL database.
	Get(ctx context.Context, in *GetDatabaseRequest, opts ...grpc.CallOption) (*Database, error)
	// Retrieves a list of MySQL databases.
	List(ctx context.Context, in *ListDatabasesRequest, opts ...grpc.CallOption) (*ListDatabasesResponse, error)
	// Creates a new MySQL database.
	Create(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Modifies the specified MySQL database.
	Update(ctx context.Context, in *UpdateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified MySQL database.
	Delete(ctx context.Context, in *DeleteDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type databaseServiceClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseServiceClient(cc *grpc.ClientConn) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) Get(ctx context.Context, in *GetDatabaseRequest, opts ...grpc.CallOption) (*Database, error) {
	out := new(Database)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) List(ctx context.Context, in *ListDatabasesRequest, opts ...grpc.CallOption) (*ListDatabasesResponse, error) {
	out := new(ListDatabasesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Create(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Update(ctx context.Context, in *UpdateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Delete(ctx context.Context, in *DeleteDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
type DatabaseServiceServer interface {
	// Returns the specified MySQL database.
	Get(context.Context, *GetDatabaseRequest) (*Database, error)
	// Retrieves a list of MySQL databases.
	List(context.Context, *ListDatabasesRequest) (*ListDatabasesResponse, error)
	// Creates a new MySQL database.
	Create(context.Context, *CreateDatabaseRequest) (*operation.Operation, error)
	// Modifies the specified MySQL database.
	Update(context.Context, *UpdateDatabaseRequest) (*operation.Operation, error)
	// Deletes the specified MySQL database.
	Delete(context.Context, *DeleteDatabaseRequest) (*operation.Operation, error)
}

func RegisterDatabaseServiceServer(s *grpc.Server, srv DatabaseServiceServer) {
	s.RegisterService(&_DatabaseService_serviceDesc, srv)
}

func _DatabaseService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Get(ctx, req.(*GetDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatabasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).List(ctx, req.(*ListDatabasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Create(ctx, req.(*CreateDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Update(ctx, req.(*UpdateDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.DatabaseService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Delete(ctx, req.(*DeleteDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatabaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.mysql.v1alpha.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DatabaseService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DatabaseService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DatabaseService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DatabaseService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DatabaseService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/mysql/v1alpha/database_service.proto",
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1alpha/database_service.proto", fileDescriptor_database_service_da98f3645adba7f1)
}

var fileDescriptor_database_service_da98f3645adba7f1 = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcb, 0x4f, 0x13, 0x5b,
	0x1c, 0xce, 0xa1, 0xdc, 0x86, 0x9e, 0xc2, 0x25, 0x39, 0xb9, 0xbd, 0x69, 0x9a, 0x0b, 0xe1, 0xce,
	0x4d, 0xb8, 0x4d, 0x75, 0x5e, 0x45, 0x88, 0x0a, 0x98, 0x58, 0x10, 0xf0, 0x01, 0x9a, 0xa2, 0x31,
	0x41, 0x4d, 0x73, 0xda, 0x39, 0x0c, 0x13, 0xe6, 0x45, 0xcf, 0x14, 0x79, 0x84, 0x85, 0x2e, 0x34,
	0xb2, 0xd4, 0xc4, 0x9d, 0xff, 0x04, 0xfe, 0x0f, 0x42, 0xe2, 0x0e, 0x97, 0x6e, 0x8d, 0x71, 0xed,
	0xd2, 0x95, 0x99, 0x73, 0xa6, 0x8f, 0x81, 0x42, 0x2b, 0x74, 0xe1, 0xae, 0x73, 0x7e, 0xbf, 0xef,
	0xcc, 0xf7, 0x9d, 0xf3, 0xfb, 0xbe, 0x0e, 0x1c, 0xdd, 0xc4, 0xb6, 0x46, 0x36, 0xe4, 0x92, 0xe9,
	0x54, 0x34, 0xd9, 0xd2, 0x8a, 0xb2, 0xb5, 0x49, 0xd7, 0x4c, 0x79, 0x5d, 0xc5, 0xa6, 0xbb, 0x82,
	0x65, 0x0d, 0x7b, 0xb8, 0x88, 0x29, 0x29, 0x50, 0x52, 0x5e, 0x37, 0x4a, 0x44, 0x72, 0xcb, 0x8e,
	0xe7, 0xa0, 0x41, 0x0e, 0x93, 0x18, 0x4c, 0xb2, 0xb4, 0xa2, 0xc4, 0x60, 0x52, 0x00, 0x4b, 0xfd,
	0xa3, 0x3b, 0x8e, 0x6e, 0x12, 0x19, 0xbb, 0x86, 0x8c, 0x6d, 0xdb, 0xf1, 0xb0, 0x67, 0x38, 0x36,
	0xe5, 0xe8, 0xd4, 0x50, 0x50, 0x65, 0x4f, 0xc5, 0xca, 0xb2, 0xbc, 0x6c, 0x10, 0x53, 0x2b, 0x58,
	0x98, 0xae, 0x06, 0x1d, 0xa9, 0x80, 0x96, 0x8f, 0x77, 0x5c, 0x52, 0x66, 0xf0, 0xa0, 0x36, 0x1c,
	0xa2, 0x5c, 0xab, 0x1e, 0xeb, 0x1b, 0x08, 0xf5, 0xad, 0x63, 0xd3, 0xd0, 0x1a, 0xcb, 0x62, 0x9b,
	0xca, 0x79, 0xbb, 0xf0, 0x02, 0x40, 0x34, 0x4b, 0xbc, 0xe9, 0x60, 0x35, 0x4f, 0xd6, 0x2a, 0x84,
	0x7a, 0xe8, 0x02, 0x84, 0x25, 0xb3, 0x42, 0x3d, 0x52, 0x2e, 0x18, 0x5a, 0x12, 0x0c, 0x81, 0x74,
	0x2c, 0xd7, 0xfb, 0x6d, 0x5f, 0x05, 0xbb, 0x07, 0x6a, 0xf7, 0xc4, 0xe4, 0xa8, 0x92, 0x8f, 0x05,
	0xf5, 0x9b, 0x1a, 0x9a, 0x82, 0x7d, 0xb5, 0xf3, 0xb4, 0xb1, 0x45, 0x92, 0x5d, 0xac, 0x7f, 0xd0,
	0xef, 0xff, 0xbe, 0xaf, 0xfe, 0xf9, 0x08, 0x8b, 0x5b, 0xd7, 0xc5, 0x25, 0x45, 0xbc, 0x52, 0x10,
	0x9f, 0x64, 0xf8, 0x0e, 0x63, 0x23, 0xf9, 0xde, 0x2a, 0x68, 0x01, 0x5b, 0x44, 0x78, 0x0b, 0xe0,
	0x5f, 0x77, 0x0c, 0x5a, 0x63, 0x42, 0xcf, 0x44, 0xe5, 0x7f, 0x18, 0x73, 0xb1, 0x4e, 0x0a, 0xd4,
	0xd8, 0xe2, 0x34, 0x22, 0x39, 0xf8, 0x63, 0x5f, 0x8d, 0x2a, 0xa2, 0xaa, 0x28, 0x4a, 0xbe, 0xc7,
	0x2f, 0x2e, 0x1a, 0x5b, 0x04, 0xa5, 0x21, 0x64, 0x8d, 0x9e, 0xb3, 0x4a, 0xec, 0x64, 0x84, 0xed,
	0x1a, 0xdb, 0x3d, 0x50, 0xff, 0x98, 0x98, 0x54, 0x15, 0x25, 0xcf, 0x76, 0xb9, 0xef, 0xd7, 0x84,
	0x97, 0x00, 0x26, 0x8e, 0x10, 0xa3, 0xae, 0x63, 0x53, 0x82, 0x66, 0x60, 0xac, 0x2a, 0x81, 0x26,
	0xc1, 0x50, 0x24, 0x1d, 0xcf, 0xa6, 0xa5, 0xd3, 0x27, 0x48, 0xaa, 0x1d, 0x74, 0x1d, 0x8a, 0x86,
	0x61, 0xbf, 0x4d, 0x36, 0xbc, 0x42, 0x03, 0x21, 0x76, 0x82, 0xf9, 0x3e, 0x7f, 0xf9, 0x5e, 0x8d,
	0xc9, 0x3b, 0x00, 0x13, 0x53, 0x65, 0x82, 0x3d, 0x72, 0xae, 0xeb, 0x7a, 0xd8, 0x70, 0x5d, 0xd4,
	0x25, 0x25, 0xf6, 0xb2, 0x78, 0xf6, 0x62, 0xbb, 0xd4, 0x17, 0x5d, 0x52, 0xca, 0x75, 0xfb, 0xbb,
	0xd7, 0xaf, 0xd0, 0x5f, 0x13, 0x1e, 0xc3, 0xbf, 0xc3, 0xf4, 0xe6, 0x89, 0x87, 0xfd, 0x0e, 0x34,
	0x70, 0x9c, 0x5f, 0x23, 0xa3, 0xff, 0x9a, 0x0e, 0xd0, 0x91, 0x01, 0xf9, 0x00, 0x60, 0xe2, 0x81,
	0xab, 0x9d, 0x57, 0x7d, 0x27, 0x86, 0x15, 0x8d, 0xc3, 0x78, 0x85, 0x51, 0x61, 0xe6, 0x66, 0xe3,
	0x13, 0xcf, 0xa6, 0x24, 0xee, 0x7f, 0xa9, 0xea, 0x7f, 0x69, 0xc6, 0xf7, 0xff, 0x3c, 0xa6, 0xab,
	0x79, 0xc8, 0xdb, 0xfd, 0xdf, 0xfe, 0x31, 0x85, 0x75, 0x74, 0xf4, 0x98, 0x5e, 0x01, 0x98, 0x98,
	0x26, 0x26, 0xf9, 0x0d, 0x8e, 0xc9, 0x57, 0x1a, 0xa6, 0xd2, 0x49, 0xa5, 0xd9, 0xd7, 0x3d, 0xb0,
	0xbf, 0x36, 0x93, 0x3c, 0xc6, 0xd1, 0x7b, 0x00, 0x23, 0xb3, 0xc4, 0x43, 0xd9, 0x56, 0xc3, 0x7c,
	0x3c, 0xf3, 0x52, 0x6d, 0x7b, 0x57, 0x58, 0x78, 0xfe, 0xe9, 0xcb, 0x9b, 0xae, 0x39, 0x34, 0x23,
	0x5b, 0xd8, 0xc6, 0x3a, 0xd1, 0xc4, 0x70, 0xc6, 0x06, 0x42, 0xa8, 0xbc, 0x5d, 0x17, 0xb9, 0x53,
	0x4b, 0x5e, 0x2a, 0x6f, 0x87, 0xc4, 0xed, 0xf8, 0xac, 0xbb, 0xfd, 0x88, 0x41, 0x97, 0x5a, 0x51,
	0x68, 0x96, 0x90, 0xa9, 0xd1, 0x5f, 0x44, 0xf1, 0xf8, 0x12, 0xae, 0x31, 0x15, 0x97, 0xd1, 0xd8,
	0xd9, 0x54, 0xa0, 0x8f, 0x00, 0x46, 0xb9, 0xdf, 0x51, 0x4b, 0x06, 0x4d, 0x63, 0x2b, 0xf5, 0x6f,
	0x18, 0x56, 0xff, 0xa7, 0xbb, 0x5b, 0xfd, 0x25, 0xe8, 0x7b, 0x87, 0x19, 0xe1, 0xc4, 0x5c, 0xe9,
	0xa9, 0xae, 0x30, 0x29, 0xe3, 0xc2, 0x19, 0xa5, 0x5c, 0x05, 0x19, 0xf4, 0x19, 0xc0, 0x28, 0xb7,
	0x65, 0x6b, 0x35, 0x4d, 0x63, 0xa8, 0x1d, 0x35, 0x4f, 0xb9, 0x9a, 0x13, 0xec, 0x1f, 0x56, 0x73,
	0x3b, 0xdb, 0xa1, 0xf1, 0xf2, 0xd5, 0x7d, 0x05, 0x30, 0xca, 0xad, 0xd8, 0x5a, 0x5d, 0xd3, 0xf4,
	0x68, 0x47, 0xdd, 0x33, 0xb0, 0x77, 0x98, 0x91, 0x4f, 0xf4, 0x7c, 0xe2, 0x68, 0x3c, 0xde, 0xb0,
	0x5c, 0x6f, 0x93, 0x5b, 0x29, 0xd3, 0x21, 0xad, 0xb9, 0x5b, 0x4b, 0x73, 0xba, 0xe1, 0xad, 0x54,
	0x8a, 0x52, 0xc9, 0xb1, 0x64, 0x4e, 0x59, 0xe4, 0xdf, 0x42, 0xba, 0x23, 0xea, 0xc4, 0x66, 0x6f,
	0x97, 0x4f, 0xff, 0x48, 0x1a, 0x67, 0x4f, 0xc5, 0x28, 0xeb, 0x1d, 0xf9, 0x19, 0x00, 0x00, 0xff,
	0xff, 0x03, 0xc6, 0xed, 0xeb, 0x4d, 0x0a, 0x00, 0x00,
}
