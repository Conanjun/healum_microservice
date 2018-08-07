// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/todo-srv/proto/todo/todo.proto

/*
Package go_micro_srv_todo is a generated protocol buffer package.

It is generated from these files:
	server/todo-srv/proto/todo/todo.proto

It has these top-level messages:
	Data
	ArrData
	AllRequest
	AllResponse
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	ReadRequest
	ReadResponse
	DeleteRequest
	DeleteResponse
	SearchRequest
	SearchResponse
	ByCreatorRequest
	ByCreatorResponse
	Todo
	TodoItem
*/
package go_micro_srv_todo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_micro_srv_user "server/user-srv/proto/user"

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

type TodoStatus int32

const (
	TodoStatus_TodoStatus_NONE TodoStatus = 0
	TodoStatus_COMPLETE        TodoStatus = 1
	TodoStatus_INCOMPLETE      TodoStatus = 2
)

var TodoStatus_name = map[int32]string{
	0: "TodoStatus_NONE",
	1: "COMPLETE",
	2: "INCOMPLETE",
}
var TodoStatus_value = map[string]int32{
	"TodoStatus_NONE": 0,
	"COMPLETE":        1,
	"INCOMPLETE":      2,
}

func (x TodoStatus) String() string {
	return proto.EnumName(TodoStatus_name, int32(x))
}
func (TodoStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Data struct {
	Todo *Todo `protobuf:"bytes,1,opt,name=todo" json:"todo,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Data) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type ArrData struct {
	Todos []*Todo `protobuf:"bytes,1,rep,name=todos" json:"todos,omitempty"`
}

func (m *ArrData) Reset()                    { *m = ArrData{} }
func (m *ArrData) String() string            { return proto.CompactTextString(m) }
func (*ArrData) ProtoMessage()               {}
func (*ArrData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ArrData) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type AllRequest struct {
	OrgId         string `protobuf:"bytes,1,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
	TeamId        string `protobuf:"bytes,2,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	Limit         int64  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset        int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	SortParameter string `protobuf:"bytes,5,opt,name=sort_parameter,json=sortParameter" json:"sort_parameter,omitempty"`
	SortDirection string `protobuf:"bytes,6,opt,name=sort_direction,json=sortDirection" json:"sort_direction,omitempty"`
}

func (m *AllRequest) Reset()                    { *m = AllRequest{} }
func (m *AllRequest) String() string            { return proto.CompactTextString(m) }
func (*AllRequest) ProtoMessage()               {}
func (*AllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AllRequest) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *AllRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *AllRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *AllRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *AllRequest) GetSortParameter() string {
	if m != nil {
		return m.SortParameter
	}
	return ""
}

func (m *AllRequest) GetSortDirection() string {
	if m != nil {
		return m.SortDirection
	}
	return ""
}

type AllResponse struct {
	Data    *ArrData `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Code    int64    `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *AllResponse) Reset()                    { *m = AllResponse{} }
func (m *AllResponse) String() string            { return proto.CompactTextString(m) }
func (*AllResponse) ProtoMessage()               {}
func (*AllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AllResponse) GetData() *ArrData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *AllResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AllResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateRequest struct {
	Todo   *Todo  `protobuf:"bytes,1,opt,name=todo" json:"todo,omitempty"`
	OrgId  string `protobuf:"bytes,2,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
	TeamId string `protobuf:"bytes,3,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func (m *CreateRequest) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *CreateRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type CreateResponse struct {
	Data    *Data  `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Code    int64  `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateResponse) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CreateResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UpdateRequest struct {
	Todo   *Todo  `protobuf:"bytes,1,opt,name=todo" json:"todo,omitempty"`
	OrgId  string `protobuf:"bytes,2,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
	TeamId string `protobuf:"bytes,3,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func (m *UpdateRequest) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *UpdateRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type UpdateResponse struct {
	Data    *Data  `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Code    int64  `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateResponse) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UpdateResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReadRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrgId  string `protobuf:"bytes,2,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
	TeamId string `protobuf:"bytes,3,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadRequest) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *ReadRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type ReadResponse struct {
	Data    *Data  `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Code    int64  `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReadResponse) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ReadResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ReadResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	OrgId  string `protobuf:"bytes,2,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
	TeamId string `protobuf:"bytes,3,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteRequest) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *DeleteRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type DeleteResponse struct {
	Code    int64  `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DeleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SearchRequest struct {
	Name          string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	OrgId         string `protobuf:"bytes,2,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
	TeamId        string `protobuf:"bytes,3,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	Limit         int64  `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
	Offset        int64  `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
	From          int64  `protobuf:"varint,6,opt,name=from" json:"from,omitempty"`
	To            int64  `protobuf:"varint,7,opt,name=to" json:"to,omitempty"`
	SortParameter string `protobuf:"bytes,8,opt,name=sort_parameter,json=sortParameter" json:"sort_parameter,omitempty"`
	SortDirection string `protobuf:"bytes,9,opt,name=sort_direction,json=sortDirection" json:"sort_direction,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SearchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SearchRequest) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *SearchRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
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

func (m *SearchRequest) GetSortParameter() string {
	if m != nil {
		return m.SortParameter
	}
	return ""
}

func (m *SearchRequest) GetSortDirection() string {
	if m != nil {
		return m.SortDirection
	}
	return ""
}

type SearchResponse struct {
	Data    *ArrData `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Code    int64    `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SearchResponse) GetData() *ArrData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SearchResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SearchResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ByCreatorRequest struct {
	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	OrgId         string `protobuf:"bytes,11,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
	TeamId        string `protobuf:"bytes,12,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	Limit         int64  `protobuf:"varint,13,opt,name=limit" json:"limit,omitempty"`
	Offset        int64  `protobuf:"varint,14,opt,name=offset" json:"offset,omitempty"`
	SortParameter string `protobuf:"bytes,5,opt,name=sort_parameter,json=sortParameter" json:"sort_parameter,omitempty"`
	SortDirection string `protobuf:"bytes,6,opt,name=sort_direction,json=sortDirection" json:"sort_direction,omitempty"`
}

func (m *ByCreatorRequest) Reset()                    { *m = ByCreatorRequest{} }
func (m *ByCreatorRequest) String() string            { return proto.CompactTextString(m) }
func (*ByCreatorRequest) ProtoMessage()               {}
func (*ByCreatorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ByCreatorRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ByCreatorRequest) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *ByCreatorRequest) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *ByCreatorRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ByCreatorRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ByCreatorRequest) GetSortParameter() string {
	if m != nil {
		return m.SortParameter
	}
	return ""
}

func (m *ByCreatorRequest) GetSortDirection() string {
	if m != nil {
		return m.SortDirection
	}
	return ""
}

type ByCreatorResponse struct {
	Data    *ArrData `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Code    int64    `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *ByCreatorResponse) Reset()                    { *m = ByCreatorResponse{} }
func (m *ByCreatorResponse) String() string            { return proto.CompactTextString(m) }
func (*ByCreatorResponse) ProtoMessage()               {}
func (*ByCreatorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ByCreatorResponse) GetData() *ArrData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ByCreatorResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ByCreatorResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Todo struct {
	Id      string                  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string                  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Created int64                   `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
	Updated int64                   `protobuf:"varint,4,opt,name=updated" json:"updated,omitempty"`
	Creator *go_micro_srv_user.User `protobuf:"bytes,5,opt,name=creator" json:"creator,omitempty"`
	Items   []*TodoItem             `protobuf:"bytes,6,rep,name=items" json:"items,omitempty"`
	OrgId   string                  `protobuf:"bytes,7,opt,name=org_id,json=orgId" json:"org_id,omitempty"`
}

func (m *Todo) Reset()                    { *m = Todo{} }
func (m *Todo) String() string            { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()               {}
func (*Todo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Todo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Todo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *Todo) GetCreator() *go_micro_srv_user.User {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *Todo) GetItems() []*TodoItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Todo) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

type TodoItem struct {
	Id            string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title         string     `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Created       int64      `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
	Updated       int64      `protobuf:"varint,4,opt,name=updated" json:"updated,omitempty"`
	Status        TodoStatus `protobuf:"varint,5,opt,name=status,enum=go.micro.srv.todo.TodoStatus" json:"status,omitempty"`
	DateCompleted int64      `protobuf:"varint,6,opt,name=dateCompleted" json:"dateCompleted,omitempty"`
}

func (m *TodoItem) Reset()                    { *m = TodoItem{} }
func (m *TodoItem) String() string            { return proto.CompactTextString(m) }
func (*TodoItem) ProtoMessage()               {}
func (*TodoItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *TodoItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TodoItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TodoItem) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TodoItem) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *TodoItem) GetStatus() TodoStatus {
	if m != nil {
		return m.Status
	}
	return TodoStatus_TodoStatus_NONE
}

func (m *TodoItem) GetDateCompleted() int64 {
	if m != nil {
		return m.DateCompleted
	}
	return 0
}

func init() {
	proto.RegisterType((*Data)(nil), "go.micro.srv.todo.Data")
	proto.RegisterType((*ArrData)(nil), "go.micro.srv.todo.ArrData")
	proto.RegisterType((*AllRequest)(nil), "go.micro.srv.todo.AllRequest")
	proto.RegisterType((*AllResponse)(nil), "go.micro.srv.todo.AllResponse")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.srv.todo.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.srv.todo.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.srv.todo.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.srv.todo.UpdateResponse")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.srv.todo.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.srv.todo.ReadResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.srv.todo.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.srv.todo.DeleteResponse")
	proto.RegisterType((*SearchRequest)(nil), "go.micro.srv.todo.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "go.micro.srv.todo.SearchResponse")
	proto.RegisterType((*ByCreatorRequest)(nil), "go.micro.srv.todo.ByCreatorRequest")
	proto.RegisterType((*ByCreatorResponse)(nil), "go.micro.srv.todo.ByCreatorResponse")
	proto.RegisterType((*Todo)(nil), "go.micro.srv.todo.Todo")
	proto.RegisterType((*TodoItem)(nil), "go.micro.srv.todo.TodoItem")
	proto.RegisterEnum("go.micro.srv.todo.TodoStatus", TodoStatus_name, TodoStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TodoService service

type TodoServiceClient interface {
	All(ctx context.Context, in *AllRequest, opts ...client.CallOption) (*AllResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	ByCreator(ctx context.Context, in *ByCreatorRequest, opts ...client.CallOption) (*ByCreatorResponse, error)
}

type todoServiceClient struct {
	c           client.Client
	serviceName string
}

func NewTodoServiceClient(serviceName string, c client.Client) TodoServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.todo"
	}
	return &todoServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *todoServiceClient) All(ctx context.Context, in *AllRequest, opts ...client.CallOption) (*AllResponse, error) {
	req := c.c.NewRequest(c.serviceName, "TodoService.All", in)
	out := new(AllResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "TodoService.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "TodoService.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "TodoService.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "TodoService.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "TodoService.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ByCreator(ctx context.Context, in *ByCreatorRequest, opts ...client.CallOption) (*ByCreatorResponse, error) {
	req := c.c.NewRequest(c.serviceName, "TodoService.ByCreator", in)
	out := new(ByCreatorResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TodoService service

type TodoServiceHandler interface {
	All(context.Context, *AllRequest, *AllResponse) error
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	ByCreator(context.Context, *ByCreatorRequest, *ByCreatorResponse) error
}

func RegisterTodoServiceHandler(s server.Server, hdlr TodoServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&TodoService{hdlr}, opts...))
}

type TodoService struct {
	TodoServiceHandler
}

func (h *TodoService) All(ctx context.Context, in *AllRequest, out *AllResponse) error {
	return h.TodoServiceHandler.All(ctx, in, out)
}

func (h *TodoService) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.TodoServiceHandler.Create(ctx, in, out)
}

func (h *TodoService) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.TodoServiceHandler.Update(ctx, in, out)
}

func (h *TodoService) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.TodoServiceHandler.Read(ctx, in, out)
}

func (h *TodoService) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.TodoServiceHandler.Delete(ctx, in, out)
}

func (h *TodoService) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.TodoServiceHandler.Search(ctx, in, out)
}

func (h *TodoService) ByCreator(ctx context.Context, in *ByCreatorRequest, out *ByCreatorResponse) error {
	return h.TodoServiceHandler.ByCreator(ctx, in, out)
}

func init() { proto.RegisterFile("server/todo-srv/proto/todo/todo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xd1, 0x4e, 0xdb, 0x4a,
	0x10, 0x8d, 0x13, 0xc7, 0x21, 0x13, 0xe2, 0x0b, 0x7b, 0xb9, 0x17, 0x2b, 0x55, 0xdb, 0xd4, 0x05,
	0x09, 0x15, 0x11, 0x04, 0xa8, 0x52, 0x9f, 0x5a, 0x51, 0xa0, 0x52, 0xa4, 0x42, 0x90, 0x01, 0xa9,
	0x6f, 0xc8, 0x8d, 0x97, 0xd4, 0xc2, 0xce, 0x86, 0xdd, 0x4d, 0xa4, 0xfe, 0x59, 0x9f, 0x2a, 0xf5,
	0x2b, 0xfa, 0xd0, 0xcf, 0xe8, 0x0f, 0x54, 0xbb, 0xf6, 0x26, 0x76, 0xb3, 0x41, 0x4d, 0x11, 0xbc,
	0x58, 0x9e, 0xd9, 0x99, 0xd9, 0x39, 0xb3, 0xeb, 0x73, 0x0c, 0xeb, 0x0c, 0xd3, 0x11, 0xa6, 0xdb,
	0x9c, 0x04, 0x64, 0x8b, 0xd1, 0xd1, 0xf6, 0x80, 0x12, 0x4e, 0xa4, 0x29, 0x1f, 0x2d, 0x69, 0xa3,
	0xe5, 0x1e, 0x69, 0xc5, 0x61, 0x97, 0x92, 0x16, 0xa3, 0xa3, 0x96, 0x58, 0x68, 0xa8, 0xcc, 0x21,
	0xc3, 0x34, 0x93, 0x29, 0x4c, 0xf9, 0x48, 0x32, 0xdd, 0x3d, 0x30, 0x0f, 0x7d, 0xee, 0xa3, 0x4d,
	0x30, 0x45, 0x9a, 0x63, 0x34, 0x8d, 0x8d, 0xda, 0xee, 0x6a, 0x6b, 0xaa, 0x60, 0xeb, 0x9c, 0x04,
	0xc4, 0x93, 0x41, 0xee, 0x2b, 0xa8, 0xec, 0x53, 0x2a, 0xf3, 0xb6, 0xa0, 0x2c, 0x5c, 0xcc, 0x31,
	0x9a, 0xa5, 0xdb, 0x12, 0x93, 0x28, 0xf7, 0x8b, 0x01, 0xb0, 0x1f, 0x45, 0x1e, 0xbe, 0x19, 0x62,
	0xc6, 0xd1, 0x7f, 0x60, 0x11, 0xda, 0xbb, 0x0c, 0x03, 0xb9, 0x6f, 0xd5, 0x2b, 0x13, 0xda, 0x6b,
	0x07, 0x68, 0x15, 0x2a, 0x1c, 0xfb, 0xb1, 0xf0, 0x17, 0xa5, 0xdf, 0x12, 0x66, 0x3b, 0x40, 0x2b,
	0x50, 0x8e, 0xc2, 0x38, 0xe4, 0x4e, 0xa9, 0x69, 0x6c, 0x94, 0xbc, 0xc4, 0x40, 0xff, 0x83, 0x45,
	0xae, 0xae, 0x18, 0xe6, 0x8e, 0x29, 0xdd, 0xa9, 0x85, 0xd6, 0xc1, 0x66, 0x84, 0xf2, 0xcb, 0x81,
	0x4f, 0xfd, 0x18, 0x73, 0x4c, 0x9d, 0xb2, 0xac, 0x56, 0x17, 0xde, 0x53, 0xe5, 0x1c, 0x87, 0x05,
	0x21, 0xc5, 0x5d, 0x1e, 0x92, 0xbe, 0x63, 0x4d, 0xc2, 0x0e, 0x95, 0xd3, 0xbd, 0x86, 0x9a, 0xec,
	0x9c, 0x0d, 0x48, 0x9f, 0x61, 0xd4, 0x02, 0x33, 0xf0, 0xb9, 0x9f, 0x0e, 0xac, 0xa1, 0xc1, 0x9d,
	0x8e, 0xc8, 0x93, 0x71, 0x08, 0x81, 0xd9, 0x25, 0x01, 0x96, 0x80, 0x4a, 0x9e, 0x7c, 0x47, 0x0e,
	0x54, 0x62, 0xcc, 0x98, 0xdf, 0xc3, 0x12, 0x50, 0xd5, 0x53, 0xa6, 0x1b, 0x41, 0xfd, 0x80, 0x62,
	0x9f, 0x63, 0x35, 0xa9, 0x79, 0xce, 0x27, 0x33, 0xd6, 0xe2, 0x8c, 0xb1, 0x96, 0xb2, 0x63, 0x75,
	0xaf, 0xc1, 0x56, 0xbb, 0xa5, 0xe8, 0x36, 0x73, 0xe8, 0x74, 0xdb, 0xdd, 0x05, 0xda, 0xc5, 0x20,
	0x78, 0x40, 0x68, 0x6a, 0xb7, 0xfb, 0x87, 0x76, 0x0c, 0x35, 0x0f, 0xfb, 0x81, 0x02, 0x66, 0x43,
	0x71, 0x7c, 0xb3, 0x8b, 0x61, 0x30, 0x77, 0xef, 0x21, 0x2c, 0x26, 0xe5, 0xee, 0xbf, 0xf3, 0x0e,
	0xd4, 0x0f, 0x71, 0x84, 0x27, 0x87, 0x72, 0xd7, 0xde, 0x5f, 0x83, 0xad, 0x0a, 0xa6, 0xdd, 0xcf,
	0xd7, 0xd0, 0x4f, 0x03, 0xea, 0x67, 0xd8, 0xa7, 0xdd, 0x4f, 0xaa, 0x23, 0x04, 0x66, 0xdf, 0x8f,
	0x71, 0xda, 0x93, 0x7c, 0x9f, 0xb7, 0xab, 0x09, 0x7f, 0x98, 0x7a, 0xfe, 0x28, 0xe7, 0xf8, 0x03,
	0x81, 0x79, 0x45, 0x49, 0x2c, 0xe9, 0xa0, 0xe4, 0xc9, 0x77, 0x31, 0x17, 0x4e, 0x9c, 0x8a, 0xf4,
	0x14, 0x39, 0xd1, 0x70, 0xcc, 0xc2, 0x9f, 0x71, 0x4c, 0x55, 0xc7, 0x31, 0x7d, 0xb0, 0x15, 0xe8,
	0x07, 0xa1, 0x99, 0xef, 0x06, 0x2c, 0xbd, 0xfd, 0x2c, 0xbf, 0x7d, 0x42, 0xd5, 0xa0, 0x57, 0xa1,
	0x22, 0x04, 0x62, 0xc2, 0xca, 0x96, 0x30, 0xdb, 0xd9, 0x3b, 0x50, 0x9b, 0x31, 0xed, 0x45, 0xfd,
	0xb4, 0xeb, 0xfa, 0x69, 0xdb, 0xf7, 0xc8, 0xd6, 0x37, 0xb0, 0x9c, 0x01, 0xf6, 0x20, 0xc3, 0xfc,
	0x61, 0x80, 0x29, 0x98, 0x6a, 0xea, 0xdb, 0x51, 0x37, 0xb7, 0x98, 0xb9, 0xb9, 0x0e, 0x54, 0xba,
	0x92, 0x72, 0x83, 0x54, 0xcb, 0x94, 0x29, 0x56, 0x86, 0x92, 0xb1, 0x82, 0xf4, 0x96, 0x2a, 0x13,
	0xed, 0xa4, 0x39, 0x24, 0x19, 0xcd, 0x14, 0x05, 0x48, 0x59, 0xbf, 0x60, 0x98, 0x7a, 0x2a, 0x0e,
	0xed, 0x40, 0x39, 0xe4, 0x38, 0x66, 0x8e, 0x25, 0xe5, 0xf9, 0xd1, 0x0c, 0x72, 0x6d, 0x73, 0x1c,
	0x7b, 0x49, 0x64, 0xe6, 0x94, 0x2b, 0x99, 0x53, 0x76, 0xbf, 0x19, 0xb0, 0xa0, 0x42, 0xa7, 0x10,
	0xae, 0x40, 0x99, 0x87, 0x3c, 0x52, 0x10, 0x13, 0xe3, 0xaf, 0x30, 0xbe, 0x04, 0x8b, 0x71, 0x9f,
	0x0f, 0x99, 0x84, 0x68, 0xef, 0x3e, 0x9e, 0xd1, 0xf1, 0x99, 0x0c, 0xf2, 0xd2, 0x60, 0xb4, 0x06,
	0x75, 0x91, 0x7f, 0x40, 0xe2, 0x81, 0x20, 0x9d, 0x20, 0xfd, 0x66, 0xf3, 0xce, 0x17, 0x6f, 0x00,
	0x26, 0xb9, 0xe8, 0x5f, 0xf8, 0x67, 0x62, 0x5d, 0x9e, 0x74, 0x4e, 0x8e, 0x96, 0x0a, 0x68, 0x11,
	0x16, 0x0e, 0x3a, 0xc7, 0xa7, 0xef, 0x8f, 0xce, 0x8f, 0x96, 0x0c, 0x64, 0x03, 0xb4, 0x4f, 0xc6,
	0x76, 0x71, 0xf7, 0xab, 0x09, 0x35, 0x99, 0x83, 0xe9, 0x28, 0xec, 0x62, 0xf4, 0x0e, 0x4a, 0xfb,
	0x51, 0x84, 0x74, 0x4d, 0x4e, 0xfe, 0x72, 0x1a, 0x4f, 0x66, 0x2d, 0x27, 0xd7, 0xd2, 0x2d, 0xa0,
	0x0e, 0x58, 0x89, 0x00, 0xa3, 0xa6, 0x26, 0x36, 0xf7, 0x27, 0xd0, 0x78, 0x76, 0x4b, 0x44, 0xb6,
	0x60, 0x22, 0x7b, 0xda, 0x82, 0x39, 0xfd, 0xd5, 0x16, 0xcc, 0x6b, 0xa6, 0x5b, 0x40, 0x6d, 0x30,
	0x85, 0x16, 0x21, 0x1d, 0x96, 0x8c, 0xe6, 0x35, 0x9e, 0xce, 0x5c, 0xcf, 0xf6, 0x96, 0x48, 0x83,
	0xb6, 0xb7, 0x9c, 0x0c, 0x69, 0x7b, 0xcb, 0xeb, 0x4a, 0x52, 0x30, 0x61, 0x4d, 0x6d, 0xc1, 0x9c,
	0x8a, 0x68, 0x0b, 0xe6, 0x29, 0xd7, 0x2d, 0xa0, 0x0f, 0x50, 0x1d, 0x93, 0x07, 0x7a, 0xae, 0xc9,
	0xf8, 0x9d, 0x33, 0x1b, 0x6b, 0xb7, 0x07, 0xa9, 0xca, 0x1f, 0x2d, 0xf9, 0xd7, 0xbd, 0xf7, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0xac, 0x80, 0x9c, 0xb7, 0xd8, 0x0b, 0x00, 0x00,
}
