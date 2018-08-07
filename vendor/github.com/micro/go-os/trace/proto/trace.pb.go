// Code generated by protoc-gen-go.
// source: github.com/micro/go-os/trace/proto/trace.proto
// DO NOT EDIT!

/*
Package go_micro_os_trace is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-os/trace/proto/trace.proto

It has these top-level messages:
	Annotation
	Span
	Service
	Node
*/
package go_micro_os_trace

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

type Annotation_Type int32

const (
	Annotation_UNKNOWN             Annotation_Type = 0
	Annotation_START               Annotation_Type = 1
	Annotation_END                 Annotation_Type = 2
	Annotation_TIMEOUT             Annotation_Type = 3
	Annotation_CLIENT_REQUEST      Annotation_Type = 4
	Annotation_CLIENT_RESPONSE     Annotation_Type = 5
	Annotation_CLIENT_PUBLICATION  Annotation_Type = 6
	Annotation_SERVER_REQUEST      Annotation_Type = 7
	Annotation_SERVER_RESPONSE     Annotation_Type = 8
	Annotation_SERVER_SUBSCRIPTION Annotation_Type = 9
)

var Annotation_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "START",
	2: "END",
	3: "TIMEOUT",
	4: "CLIENT_REQUEST",
	5: "CLIENT_RESPONSE",
	6: "CLIENT_PUBLICATION",
	7: "SERVER_REQUEST",
	8: "SERVER_RESPONSE",
	9: "SERVER_SUBSCRIPTION",
}
var Annotation_Type_value = map[string]int32{
	"UNKNOWN":             0,
	"START":               1,
	"END":                 2,
	"TIMEOUT":             3,
	"CLIENT_REQUEST":      4,
	"CLIENT_RESPONSE":     5,
	"CLIENT_PUBLICATION":  6,
	"SERVER_REQUEST":      7,
	"SERVER_RESPONSE":     8,
	"SERVER_SUBSCRIPTION": 9,
}

func (x Annotation_Type) String() string {
	return proto.EnumName(Annotation_Type_name, int32(x))
}
func (Annotation_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Annotation struct {
	Timestamp int64             `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Type      Annotation_Type   `protobuf:"varint,2,opt,name=type,enum=go.micro.os.trace.Annotation_Type" json:"type,omitempty"`
	Key       string            `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Value     []byte            `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Debug     map[string]string `protobuf:"bytes,5,rep,name=debug" json:"debug,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Service   *Service          `protobuf:"bytes,6,opt,name=service" json:"service,omitempty"`
}

func (m *Annotation) Reset()                    { *m = Annotation{} }
func (m *Annotation) String() string            { return proto.CompactTextString(m) }
func (*Annotation) ProtoMessage()               {}
func (*Annotation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Annotation) GetDebug() map[string]string {
	if m != nil {
		return m.Debug
	}
	return nil
}

func (m *Annotation) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type Span struct {
	// name; topic, service method, etc
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// id of the span
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// trace root id
	TraceId string `protobuf:"bytes,3,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	// parent span id
	ParentId string `protobuf:"bytes,4,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	// microseconds from epoch. start of span
	Timestamp int64 `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	// microseconds. duration of span
	Duration int64 `protobuf:"varint,6,opt,name=duration" json:"duration,omitempty"`
	// should persist?
	Debug bool `protobuf:"varint,7,opt,name=debug" json:"debug,omitempty"`
	// source origin of the request
	Source *Service `protobuf:"bytes,8,opt,name=source" json:"source,omitempty"`
	// destination of the request
	Destination *Service `protobuf:"bytes,9,opt,name=destination" json:"destination,omitempty"`
	// annotations
	Annotations []*Annotation `protobuf:"bytes,10,rep,name=annotations" json:"annotations,omitempty"`
}

func (m *Span) Reset()                    { *m = Span{} }
func (m *Span) String() string            { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()               {}
func (*Span) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Span) GetSource() *Service {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Span) GetDestination() *Service {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *Span) GetAnnotations() []*Annotation {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type Service struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version  string            `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Nodes    []*Node           `protobuf:"bytes,4,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Service) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Node struct {
	Id       string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address  string            `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Port     int64             `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Annotation)(nil), "go.micro.os.trace.Annotation")
	proto.RegisterType((*Span)(nil), "go.micro.os.trace.Span")
	proto.RegisterType((*Service)(nil), "go.micro.os.trace.Service")
	proto.RegisterType((*Node)(nil), "go.micro.os.trace.Node")
	proto.RegisterEnum("go.micro.os.trace.Annotation_Type", Annotation_Type_name, Annotation_Type_value)
}

func init() { proto.RegisterFile("github.com/micro/go-os/trace/proto/trace.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x1c, 0xdb, 0xb1, 0x73, 0xf3, 0x51, 0xcc, 0x2d, 0xa2, 0x26, 0x80, 0x54, 0x45, 0x42,
	0xea, 0xa6, 0x8e, 0x14, 0x10, 0xaa, 0x00, 0x81, 0xd2, 0xd6, 0x8b, 0x88, 0xd6, 0x29, 0x63, 0x07,
	0x96, 0x95, 0x1b, 0x8f, 0x82, 0x05, 0xf1, 0x44, 0xf6, 0xa4, 0x52, 0x9f, 0x8c, 0x17, 0xe0, 0x55,
	0x58, 0xb0, 0xe7, 0x01, 0x18, 0xcf, 0x38, 0x6e, 0x0b, 0x49, 0x2b, 0xc4, 0xee, 0xfe, 0x9c, 0x7b,
	0x7c, 0x7f, 0xce, 0x18, 0xbc, 0x69, 0xca, 0x3f, 0x2d, 0xce, 0xbc, 0x09, 0x9b, 0xf5, 0x66, 0xe9,
	0x24, 0x67, 0xbd, 0x29, 0xdb, 0x65, 0x45, 0x8f, 0xe7, 0xf1, 0x84, 0xf6, 0xe6, 0x39, 0xe3, 0x4c,
	0xd9, 0x9e, 0xb4, 0xf1, 0xde, 0x94, 0x79, 0x12, 0xe7, 0xb1, 0xc2, 0x93, 0x89, 0xee, 0x4f, 0x1d,
	0x60, 0x90, 0x65, 0x8c, 0xc7, 0x3c, 0x65, 0x19, 0x3e, 0x86, 0x16, 0x4f, 0x67, 0xb4, 0xe0, 0xf1,
	0x6c, 0xee, 0x6a, 0xdb, 0xda, 0x8e, 0x4e, 0x2e, 0x03, 0xf8, 0x02, 0x0c, 0x7e, 0x31, 0xa7, 0x6e,
	0x43, 0x24, 0x36, 0xfa, 0x5d, 0xef, 0x0f, 0x3a, 0xef, 0x92, 0xca, 0x8b, 0x04, 0x92, 0x48, 0x3c,
	0x3a, 0xa0, 0x7f, 0xa6, 0x17, 0xae, 0x2e, 0xca, 0x5a, 0xa4, 0x34, 0xf1, 0x3e, 0x98, 0xe7, 0xf1,
	0x97, 0x05, 0x75, 0x0d, 0x11, 0xfb, 0x9f, 0x28, 0x07, 0xdf, 0x80, 0x99, 0xd0, 0xb3, 0xc5, 0xd4,
	0x35, 0xb7, 0xf5, 0x9d, 0x76, 0x7f, 0xe7, 0xe6, 0x0f, 0x1c, 0x96, 0x50, 0x3f, 0xe3, 0xf9, 0x05,
	0x51, 0x65, 0xf8, 0x1c, 0xac, 0x82, 0xe6, 0xe7, 0xe9, 0x84, 0xba, 0x4d, 0xc1, 0xdb, 0xee, 0x77,
	0x56, 0x30, 0x84, 0x0a, 0x41, 0x96, 0xd0, 0xce, 0x1e, 0xc0, 0x25, 0xd5, 0xb2, 0x57, 0x6d, 0x45,
	0xaf, 0x0d, 0x19, 0x53, 0xce, 0xcb, 0xc6, 0x9e, 0xd6, 0xfd, 0xaa, 0x81, 0x51, 0x8e, 0x89, 0x6d,
	0xb0, 0xc6, 0xc1, 0xbb, 0x60, 0xf4, 0x31, 0x70, 0xfe, 0xc3, 0x16, 0x98, 0x61, 0x34, 0x20, 0x91,
	0xa3, 0xa1, 0x05, 0xba, 0x1f, 0x1c, 0x3a, 0x8d, 0x12, 0x10, 0x0d, 0x8f, 0xfd, 0xd1, 0x38, 0x72,
	0x74, 0x44, 0xd8, 0x38, 0x38, 0x1a, 0xfa, 0x41, 0x74, 0x4a, 0xfc, 0xf7, 0x63, 0x3f, 0x8c, 0x1c,
	0x03, 0x37, 0xe1, 0x6e, 0x1d, 0x0b, 0x4f, 0x46, 0x41, 0xe8, 0x3b, 0x26, 0x3e, 0x00, 0xac, 0x82,
	0x27, 0xe3, 0xfd, 0xa3, 0xe1, 0xc1, 0x20, 0x1a, 0x8e, 0x02, 0xa7, 0x59, 0x12, 0x84, 0x3e, 0xf9,
	0xe0, 0x93, 0x9a, 0xc0, 0x2a, 0x09, 0xea, 0x58, 0x45, 0x60, 0xe3, 0x16, 0x6c, 0x56, 0xc1, 0x70,
	0xbc, 0x1f, 0x1e, 0x90, 0xe1, 0x89, 0x64, 0x68, 0x75, 0xbf, 0x37, 0xc0, 0x08, 0xe7, 0x71, 0x26,
	0xa8, 0x8c, 0x2c, 0x9e, 0xd1, 0x6a, 0x5e, 0x69, 0xe3, 0x06, 0x34, 0xd2, 0xa4, 0x9a, 0x56, 0x58,
	0xf8, 0x10, 0x6c, 0xb9, 0xba, 0x53, 0x11, 0x55, 0x37, 0xb4, 0xa4, 0x3f, 0x4c, 0xf0, 0x11, 0xb4,
	0xe6, 0x71, 0x4e, 0x33, 0x5e, 0xe6, 0x0c, 0x99, 0xb3, 0x55, 0x40, 0x24, 0xaf, 0x89, 0xc9, 0xfc,
	0x5d, 0x4c, 0x1d, 0xb0, 0x93, 0x45, 0x2e, 0x4f, 0x29, 0xaf, 0xa5, 0x93, 0xda, 0x2f, 0x57, 0xae,
	0x84, 0x60, 0x89, 0x84, 0xbd, 0x3c, 0x6f, 0x1f, 0x9a, 0x05, 0x5b, 0xe4, 0xe2, 0xba, 0xf6, 0xad,
	0xd7, 0xad, 0x90, 0xf8, 0x1a, 0xda, 0x89, 0xf8, 0x60, 0x9a, 0xa9, 0x0f, 0xb5, 0x6e, 0x2d, 0xbc,
	0x0a, 0xc7, 0xb7, 0xd0, 0x8e, 0x6b, 0xc1, 0x15, 0x2e, 0x48, 0x59, 0x3e, 0xb9, 0x51, 0x96, 0xe4,
	0x6a, 0x45, 0xf7, 0x87, 0x06, 0x56, 0xc5, 0xbc, 0x72, 0xd5, 0x2e, 0x58, 0xe7, 0x34, 0x2f, 0xca,
	0xd6, 0xd4, 0xbe, 0x97, 0x2e, 0x1e, 0x82, 0x3d, 0xa3, 0x3c, 0x4e, 0x62, 0x1e, 0x8b, 0xa5, 0xaf,
	0x7b, 0x0e, 0x15, 0xb7, 0x77, 0x5c, 0x41, 0xd5, 0x73, 0xa8, 0x2b, 0x71, 0x17, 0xcc, 0x8c, 0x89,
	0x89, 0xc4, 0x6d, 0x4a, 0x8a, 0xad, 0x15, 0x14, 0x81, 0xc8, 0x13, 0x85, 0xea, 0xbc, 0x82, 0x3b,
	0xd7, 0x98, 0xfe, 0xea, 0x35, 0x7c, 0x13, 0xaf, 0xa1, 0x24, 0xab, 0xf4, 0xa3, 0xd5, 0xfa, 0x11,
	0x43, 0xc6, 0x49, 0x92, 0xd3, 0xa2, 0x58, 0x0e, 0x59, 0xb9, 0xe5, 0x4a, 0xe6, 0x2c, 0xe7, 0x52,
	0x55, 0x3a, 0x91, 0x36, 0x0e, 0xae, 0x0c, 0xae, 0xba, 0x7e, 0xba, 0xa6, 0xeb, 0x75, 0x53, 0xff,
	0xd3, 0x18, 0x67, 0x4d, 0xf9, 0xaf, 0x7c, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xf4, 0x58,
	0x45, 0x5d, 0x05, 0x00, 0x00,
}
