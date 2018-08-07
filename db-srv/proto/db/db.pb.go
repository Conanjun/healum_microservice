// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/db-srv/proto/db/db.proto

/*
Package go_micro_srv_db is a generated protocol buffer package.

It is generated from these files:
	server/db-srv/proto/db/db.proto

It has these top-level messages:
	Database
	Record
	InitDbRequest
	InitDbResponse
	RemoveDbRequest
	RemoveDbResponse
	ReadRequest
	ReadResponse
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	SearchRequest
	SearchResponse
	RunQueryRequest
	RunQueryResponse
	CreateDatabaseRequest
	CreateDatabaseResponse
	DeleteDatabaseRequest
	DeleteDatabaseResponse
*/
package go_micro_srv_db

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Database struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Table    string            `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
	Driver   string            `protobuf:"bytes,3,opt,name=driver" json:"driver,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Database) Reset()                    { *m = Database{} }
func (m *Database) String() string            { return proto.CompactTextString(m) }
func (*Database) ProtoMessage()               {}
func (*Database) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Database) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Database) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *Database) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Database) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Record struct {
	Id         string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Created    int64             `protobuf:"varint,2,opt,name=created" json:"created,omitempty"`
	Updated    int64             `protobuf:"varint,3,opt,name=updated" json:"updated,omitempty"`
	Name       string            `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Parameter1 string            `protobuf:"bytes,5,opt,name=parameter1" json:"parameter1,omitempty"`
	Parameter2 string            `protobuf:"bytes,6,opt,name=parameter2" json:"parameter2,omitempty"`
	Parameter3 string            `protobuf:"bytes,7,opt,name=parameter3" json:"parameter3,omitempty"`
	Lat        float64           `protobuf:"fixed64,8,opt,name=lat" json:"lat,omitempty"`
	Lng        float64           `protobuf:"fixed64,9,opt,name=lng" json:"lng,omitempty"`
	Metadata   map[string]string `protobuf:"bytes,10,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Record) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Record) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Record) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *Record) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Record) GetParameter1() string {
	if m != nil {
		return m.Parameter1
	}
	return ""
}

func (m *Record) GetParameter2() string {
	if m != nil {
		return m.Parameter2
	}
	return ""
}

func (m *Record) GetParameter3() string {
	if m != nil {
		return m.Parameter3
	}
	return ""
}

func (m *Record) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Record) GetLng() float64 {
	if m != nil {
		return m.Lng
	}
	return 0
}

func (m *Record) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type InitDbRequest struct {
}

func (m *InitDbRequest) Reset()                    { *m = InitDbRequest{} }
func (m *InitDbRequest) String() string            { return proto.CompactTextString(m) }
func (*InitDbRequest) ProtoMessage()               {}
func (*InitDbRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type InitDbResponse struct {
}

func (m *InitDbResponse) Reset()                    { *m = InitDbResponse{} }
func (m *InitDbResponse) String() string            { return proto.CompactTextString(m) }
func (*InitDbResponse) ProtoMessage()               {}
func (*InitDbResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RemoveDbRequest struct {
}

func (m *RemoveDbRequest) Reset()                    { *m = RemoveDbRequest{} }
func (m *RemoveDbRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveDbRequest) ProtoMessage()               {}
func (*RemoveDbRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type RemoveDbResponse struct {
}

func (m *RemoveDbResponse) Reset()                    { *m = RemoveDbResponse{} }
func (m *RemoveDbResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveDbResponse) ProtoMessage()               {}
func (*RemoveDbResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ReadRequest struct {
	Database   *Database `protobuf:"bytes,1,opt,name=database" json:"database,omitempty"`
	Id         string    `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Parameter3 string    `protobuf:"bytes,3,opt,name=parameter3" json:"parameter3,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadRequest) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadRequest) GetParameter3() string {
	if m != nil {
		return m.Parameter3
	}
	return ""
}

type ReadResponse struct {
	Record *Record `protobuf:"bytes,1,opt,name=record" json:"record,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadResponse) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type CreateRequest struct {
	Database *Database `protobuf:"bytes,1,opt,name=database" json:"database,omitempty"`
	Record   *Record   `protobuf:"bytes,2,opt,name=record" json:"record,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateRequest) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *CreateRequest) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type UpdateRequest struct {
	Database *Database `protobuf:"bytes,1,opt,name=database" json:"database,omitempty"`
	Record   *Record   `protobuf:"bytes,2,opt,name=record" json:"record,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateRequest) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *UpdateRequest) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type DeleteRequest struct {
	Database   *Database `protobuf:"bytes,1,opt,name=database" json:"database,omitempty"`
	Id         string    `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Parameter3 string    `protobuf:"bytes,3,opt,name=parameter3" json:"parameter3,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *DeleteRequest) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteRequest) GetParameter3() string {
	if m != nil {
		return m.Parameter3
	}
	return ""
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type SearchRequest struct {
	Database *Database         `protobuf:"bytes,1,opt,name=database" json:"database,omitempty"`
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// timeseries bounds
	From int64 `protobuf:"varint,3,opt,name=from" json:"from,omitempty"`
	To   int64 `protobuf:"varint,4,opt,name=to" json:"to,omitempty"`
	// range bounds
	Limit  int64 `protobuf:"varint,5,opt,name=limit" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
	// order
	Reverse bool `protobuf:"varint,7,opt,name=reverse" json:"reverse,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *SearchRequest) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *SearchRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SearchRequest) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *SearchRequest) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *SearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchRequest) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

type SearchResponse struct {
	Records []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *SearchResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type RunQueryRequest struct {
	Database *Database `protobuf:"bytes,1,opt,name=database" json:"database,omitempty"`
	Query    string    `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
}

func (m *RunQueryRequest) Reset()                    { *m = RunQueryRequest{} }
func (m *RunQueryRequest) String() string            { return proto.CompactTextString(m) }
func (*RunQueryRequest) ProtoMessage()               {}
func (*RunQueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *RunQueryRequest) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

func (m *RunQueryRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type RunQueryResponse struct {
	Records []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *RunQueryResponse) Reset()                    { *m = RunQueryResponse{} }
func (m *RunQueryResponse) String() string            { return proto.CompactTextString(m) }
func (*RunQueryResponse) ProtoMessage()               {}
func (*RunQueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *RunQueryResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type CreateDatabaseRequest struct {
	Database *Database `protobuf:"bytes,1,opt,name=database" json:"database,omitempty"`
}

func (m *CreateDatabaseRequest) Reset()                    { *m = CreateDatabaseRequest{} }
func (m *CreateDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDatabaseRequest) ProtoMessage()               {}
func (*CreateDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *CreateDatabaseRequest) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

type CreateDatabaseResponse struct {
}

func (m *CreateDatabaseResponse) Reset()                    { *m = CreateDatabaseResponse{} }
func (m *CreateDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateDatabaseResponse) ProtoMessage()               {}
func (*CreateDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

type DeleteDatabaseRequest struct {
	Database *Database `protobuf:"bytes,1,opt,name=database" json:"database,omitempty"`
}

func (m *DeleteDatabaseRequest) Reset()                    { *m = DeleteDatabaseRequest{} }
func (m *DeleteDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDatabaseRequest) ProtoMessage()               {}
func (*DeleteDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *DeleteDatabaseRequest) GetDatabase() *Database {
	if m != nil {
		return m.Database
	}
	return nil
}

type DeleteDatabaseResponse struct {
}

func (m *DeleteDatabaseResponse) Reset()                    { *m = DeleteDatabaseResponse{} }
func (m *DeleteDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteDatabaseResponse) ProtoMessage()               {}
func (*DeleteDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func init() {
	proto.RegisterType((*Database)(nil), "go.micro.srv.db.Database")
	proto.RegisterType((*Record)(nil), "go.micro.srv.db.Record")
	proto.RegisterType((*InitDbRequest)(nil), "go.micro.srv.db.InitDbRequest")
	proto.RegisterType((*InitDbResponse)(nil), "go.micro.srv.db.InitDbResponse")
	proto.RegisterType((*RemoveDbRequest)(nil), "go.micro.srv.db.RemoveDbRequest")
	proto.RegisterType((*RemoveDbResponse)(nil), "go.micro.srv.db.RemoveDbResponse")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.srv.db.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.srv.db.ReadResponse")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.srv.db.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.srv.db.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.srv.db.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.srv.db.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.srv.db.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.srv.db.DeleteResponse")
	proto.RegisterType((*SearchRequest)(nil), "go.micro.srv.db.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "go.micro.srv.db.SearchResponse")
	proto.RegisterType((*RunQueryRequest)(nil), "go.micro.srv.db.RunQueryRequest")
	proto.RegisterType((*RunQueryResponse)(nil), "go.micro.srv.db.RunQueryResponse")
	proto.RegisterType((*CreateDatabaseRequest)(nil), "go.micro.srv.db.CreateDatabaseRequest")
	proto.RegisterType((*CreateDatabaseResponse)(nil), "go.micro.srv.db.CreateDatabaseResponse")
	proto.RegisterType((*DeleteDatabaseRequest)(nil), "go.micro.srv.db.DeleteDatabaseRequest")
	proto.RegisterType((*DeleteDatabaseResponse)(nil), "go.micro.srv.db.DeleteDatabaseResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DB service

type DBClient interface {
	InitDb(ctx context.Context, in *InitDbRequest, opts ...client.CallOption) (*InitDbResponse, error)
	RemoveDb(ctx context.Context, in *RemoveDbRequest, opts ...client.CallOption) (*RemoveDbResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	RunQuery(ctx context.Context, in *RunQueryRequest, opts ...client.CallOption) (*RunQueryResponse, error)
	CreateDatabase(ctx context.Context, in *CreateDatabaseRequest, opts ...client.CallOption) (*CreateDatabaseResponse, error)
	DeleteDatabase(ctx context.Context, in *DeleteDatabaseRequest, opts ...client.CallOption) (*DeleteDatabaseResponse, error)
}

type dBClient struct {
	c           client.Client
	serviceName string
}

func NewDBClient(serviceName string, c client.Client) DBClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.db"
	}
	return &dBClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *dBClient) InitDb(ctx context.Context, in *InitDbRequest, opts ...client.CallOption) (*InitDbResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.InitDb", in)
	out := new(InitDbResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) RemoveDb(ctx context.Context, in *RemoveDbRequest, opts ...client.CallOption) (*RemoveDbResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.RemoveDb", in)
	out := new(RemoveDbResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) RunQuery(ctx context.Context, in *RunQueryRequest, opts ...client.CallOption) (*RunQueryResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.RunQuery", in)
	out := new(RunQueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) CreateDatabase(ctx context.Context, in *CreateDatabaseRequest, opts ...client.CallOption) (*CreateDatabaseResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.CreateDatabase", in)
	out := new(CreateDatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBClient) DeleteDatabase(ctx context.Context, in *DeleteDatabaseRequest, opts ...client.CallOption) (*DeleteDatabaseResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DB.DeleteDatabase", in)
	out := new(DeleteDatabaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DB service

type DBHandler interface {
	InitDb(context.Context, *InitDbRequest, *InitDbResponse) error
	RemoveDb(context.Context, *RemoveDbRequest, *RemoveDbResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	RunQuery(context.Context, *RunQueryRequest, *RunQueryResponse) error
	CreateDatabase(context.Context, *CreateDatabaseRequest, *CreateDatabaseResponse) error
	DeleteDatabase(context.Context, *DeleteDatabaseRequest, *DeleteDatabaseResponse) error
}

func RegisterDBHandler(s server.Server, hdlr DBHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&DB{hdlr}, opts...))
}

type DB struct {
	DBHandler
}

func (h *DB) InitDb(ctx context.Context, in *InitDbRequest, out *InitDbResponse) error {
	return h.DBHandler.InitDb(ctx, in, out)
}

func (h *DB) RemoveDb(ctx context.Context, in *RemoveDbRequest, out *RemoveDbResponse) error {
	return h.DBHandler.RemoveDb(ctx, in, out)
}

func (h *DB) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.DBHandler.Read(ctx, in, out)
}

func (h *DB) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.DBHandler.Create(ctx, in, out)
}

func (h *DB) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.DBHandler.Update(ctx, in, out)
}

func (h *DB) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.DBHandler.Delete(ctx, in, out)
}

func (h *DB) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.DBHandler.Search(ctx, in, out)
}

func (h *DB) RunQuery(ctx context.Context, in *RunQueryRequest, out *RunQueryResponse) error {
	return h.DBHandler.RunQuery(ctx, in, out)
}

func (h *DB) CreateDatabase(ctx context.Context, in *CreateDatabaseRequest, out *CreateDatabaseResponse) error {
	return h.DBHandler.CreateDatabase(ctx, in, out)
}

func (h *DB) DeleteDatabase(ctx context.Context, in *DeleteDatabaseRequest, out *DeleteDatabaseResponse) error {
	return h.DBHandler.DeleteDatabase(ctx, in, out)
}

func init() { proto.RegisterFile("server/db-srv/proto/db/db.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 798 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0x6e, 0xec, 0xd4, 0x4d, 0xa7, 0x7f, 0x0e, 0xff, 0xaa, 0x2d, 0x26, 0x82, 0x36, 0x58, 0x82,
	0xf6, 0x02, 0x12, 0x35, 0x15, 0x12, 0x82, 0x0b, 0x04, 0x4d, 0x25, 0x50, 0x45, 0xa5, 0x2e, 0xe2,
	0x16, 0x69, 0x1d, 0x6f, 0x4b, 0x44, 0x12, 0xa7, 0xeb, 0x8d, 0xab, 0x3e, 0x1a, 0x12, 0xcf, 0xc0,
	0x53, 0xf0, 0x20, 0x68, 0x0f, 0x76, 0xb2, 0xee, 0x86, 0x53, 0x40, 0xdc, 0xed, 0xec, 0x8c, 0xe7,
	0x9b, 0xfd, 0x66, 0xe6, 0x4b, 0x60, 0x37, 0xa1, 0x2c, 0xa5, 0xac, 0x13, 0x85, 0x8f, 0x12, 0x96,
	0x76, 0x26, 0x2c, 0xe6, 0x71, 0x27, 0x0a, 0x3b, 0x51, 0xd8, 0x96, 0x67, 0x54, 0xbf, 0x88, 0xdb,
	0xa3, 0x41, 0x9f, 0xc5, 0xed, 0x84, 0xa5, 0xed, 0x28, 0x0c, 0xbe, 0x94, 0xa0, 0xd2, 0x23, 0x9c,
	0x84, 0x24, 0xa1, 0x08, 0x41, 0x79, 0x4c, 0x46, 0xd4, 0x2f, 0xb5, 0x4a, 0xfb, 0xeb, 0x58, 0x9e,
	0xd1, 0x26, 0xac, 0x72, 0x12, 0x0e, 0xa9, 0xef, 0xc8, 0x4b, 0x65, 0xa0, 0x6d, 0xf0, 0x22, 0x36,
	0x48, 0x29, 0xf3, 0x5d, 0x79, 0xad, 0x2d, 0x74, 0x04, 0x95, 0x11, 0xe5, 0x24, 0x22, 0x9c, 0xf8,
	0xe5, 0x96, 0xbb, 0xbf, 0xd1, 0xdd, 0x6b, 0x17, 0x20, 0xdb, 0x19, 0x5c, 0xfb, 0x8d, 0x8e, 0x3c,
	0x1e, 0x73, 0x76, 0x8d, 0xf3, 0x0f, 0x9b, 0xcf, 0xa0, 0x6a, 0xb8, 0x50, 0x03, 0xdc, 0x8f, 0xf4,
	0x5a, 0x97, 0x25, 0x8e, 0xa2, 0xaa, 0x94, 0x0c, 0xa7, 0x79, 0x55, 0xd2, 0x78, 0xea, 0x3c, 0x29,
	0x05, 0x5f, 0x1d, 0xf0, 0x30, 0xed, 0xc7, 0x2c, 0x42, 0x35, 0x70, 0x06, 0x91, 0xfe, 0xca, 0x19,
	0x44, 0xc8, 0x87, 0xb5, 0x3e, 0xa3, 0x84, 0xd3, 0x48, 0x7e, 0xe6, 0xe2, 0xcc, 0x14, 0x9e, 0xe9,
	0x24, 0x92, 0x1e, 0x57, 0x79, 0xb4, 0x99, 0x53, 0x52, 0x9e, 0xa3, 0x64, 0x07, 0x60, 0x42, 0x18,
	0x19, 0x51, 0x4e, 0xd9, 0x81, 0xbf, 0x2a, 0x3d, 0x73, 0x37, 0x86, 0xbf, 0xeb, 0x7b, 0x05, 0x7f,
	0xd7, 0xf0, 0x1f, 0xfa, 0x6b, 0x05, 0xff, 0xa1, 0x78, 0xee, 0x90, 0x70, 0xbf, 0xd2, 0x2a, 0xed,
	0x97, 0xb0, 0x38, 0xca, 0x9b, 0xf1, 0x85, 0xbf, 0xae, 0x6f, 0xc6, 0x17, 0xe8, 0xc5, 0x1c, 0xd1,
	0x20, 0x89, 0xbe, 0x7f, 0x83, 0x68, 0x45, 0xc3, 0xdf, 0xa1, 0xb9, 0x0e, 0xd5, 0xd7, 0xe3, 0x01,
	0xef, 0x85, 0x98, 0x5e, 0x4e, 0x69, 0xc2, 0x83, 0x06, 0xd4, 0xb2, 0x8b, 0x64, 0x12, 0x8f, 0x13,
	0x1a, 0xfc, 0x0f, 0x75, 0x4c, 0x47, 0x71, 0x4a, 0x67, 0x41, 0x08, 0x1a, 0xb3, 0x2b, 0x1d, 0xc6,
	0x61, 0x03, 0x53, 0x12, 0xe9, 0x10, 0xf4, 0x18, 0x2a, 0x91, 0x1e, 0x10, 0x59, 0xc9, 0x46, 0xf7,
	0xf6, 0xc2, 0x09, 0xc2, 0x79, 0xa8, 0xee, 0xb5, 0x93, 0xf7, 0xda, 0xe4, 0xd8, 0x2d, 0x72, 0x1c,
	0x3c, 0x87, 0xff, 0x14, 0xaa, 0xaa, 0x02, 0x75, 0xc0, 0x63, 0x92, 0x2e, 0x0d, 0x7a, 0x6b, 0x01,
	0x9b, 0x58, 0x87, 0x05, 0x57, 0x50, 0x3d, 0x92, 0xd3, 0xb3, 0x64, 0xe1, 0x33, 0x60, 0xe7, 0xe7,
	0x80, 0x1b, 0x50, 0xcb, 0x80, 0x35, 0x83, 0x57, 0x50, 0x7d, 0x27, 0xc7, 0xf5, 0x1f, 0x94, 0x92,
	0x01, 0xeb, 0x52, 0x52, 0xa8, 0xf6, 0xe8, 0x90, 0x2e, 0x5d, 0xca, 0xaf, 0xb6, 0xb3, 0x01, 0xb5,
	0x0c, 0x57, 0x57, 0xf2, 0xc9, 0x81, 0xea, 0x5b, 0x4a, 0x58, 0xff, 0xc3, 0x92, 0xa5, 0xbc, 0x9a,
	0xdb, 0x34, 0x47, 0x6e, 0xda, 0xc3, 0x1b, 0x9f, 0x19, 0x40, 0x8b, 0x16, 0x4e, 0x68, 0xc9, 0x39,
	0x8b, 0x47, 0x5a, 0x62, 0xe4, 0x59, 0x3c, 0x94, 0xc7, 0x52, 0x5d, 0x5c, 0xec, 0xf0, 0x58, 0x6c,
	0xdc, 0x70, 0x30, 0x1a, 0x70, 0x29, 0x2b, 0x2e, 0x56, 0x86, 0x90, 0xdb, 0xf8, 0xfc, 0x3c, 0xa1,
	0x5c, 0xaa, 0x89, 0x8b, 0xb5, 0x25, 0x74, 0x8b, 0xd1, 0x94, 0xb2, 0x84, 0x4a, 0x19, 0xa9, 0xe0,
	0xcc, 0x5c, 0x6e, 0xb9, 0x8f, 0xa0, 0x96, 0xbd, 0x48, 0xaf, 0xc7, 0x81, 0x00, 0x12, 0x3d, 0x4f,
	0xfc, 0x92, 0xe4, 0x60, 0xe1, 0x6c, 0x64, 0x71, 0xc1, 0x7b, 0xa8, 0xe3, 0xe9, 0xf8, 0x6c, 0x4a,
	0xd9, 0xf5, 0x92, 0x1d, 0xd8, 0x84, 0xd5, 0x4b, 0x91, 0x26, 0x2b, 0x54, 0x1a, 0xc1, 0x31, 0x34,
	0x66, 0xf9, 0x7f, 0xbf, 0xcc, 0x53, 0xd8, 0x52, 0xeb, 0x94, 0x03, 0x2f, 0x55, 0x6c, 0xe0, 0xc3,
	0x76, 0x31, 0x9f, 0x9e, 0xc8, 0x53, 0xd8, 0x52, 0x33, 0xfa, 0xe7, 0x90, 0x8a, 0xf9, 0x14, 0x52,
	0xf7, 0xb3, 0x07, 0x4e, 0xef, 0x25, 0x3a, 0x01, 0x4f, 0x49, 0x32, 0xda, 0xb9, 0x91, 0xcf, 0x10,
	0xef, 0xe6, 0xee, 0x42, 0xbf, 0xae, 0x7d, 0x05, 0x9d, 0x41, 0x25, 0x93, 0x6e, 0xd4, 0xb2, 0xb0,
	0x6a, 0x08, 0x7d, 0xf3, 0xde, 0x77, 0x22, 0xf2, 0x94, 0xc7, 0x50, 0x16, 0x1a, 0x8c, 0xee, 0x58,
	0x82, 0xf3, 0x1f, 0x84, 0xe6, 0xdd, 0x05, 0xde, 0x3c, 0xcd, 0x09, 0x78, 0x8a, 0x71, 0xcb, 0x33,
	0x0d, 0x89, 0xb6, 0x3c, 0xb3, 0xa0, 0xa4, 0x32, 0x99, 0x92, 0x34, 0x4b, 0x32, 0x43, 0x64, 0x2d,
	0xc9, 0x0a, 0x5a, 0x28, 0x93, 0xa9, 0x0e, 0x59, 0x92, 0x19, 0x32, 0x69, 0x49, 0x56, 0x90, 0x33,
	0x99, 0x4c, 0x2d, 0xa5, 0x25, 0x99, 0xa1, 0x3f, 0x96, 0x64, 0xe6, 0x36, 0xeb, 0x6e, 0xea, 0xe5,
	0xb1, 0x75, 0xd3, 0xdc, 0x5b, 0x5b, 0x37, 0x0b, 0x9b, 0x17, 0xac, 0xa0, 0x7e, 0xf6, 0xbb, 0x94,
	0xff, 0x9d, 0x7c, 0xb0, 0x80, 0xee, 0xc2, 0xfc, 0x37, 0xf7, 0x7e, 0x18, 0x37, 0x0f, 0x62, 0xce,
	0xbc, 0x05, 0xc4, 0xba, 0x64, 0x16, 0x10, 0xfb, 0xf2, 0x04, 0x2b, 0xa1, 0x27, 0xff, 0x2b, 0x1f,
	0x7e, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x09, 0x91, 0xda, 0x4e, 0x0b, 0x00, 0x00,
}
