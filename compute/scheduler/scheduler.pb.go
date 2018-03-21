// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheduler.proto

/*
Package scheduler is a generated protocol buffer package.

It is generated from these files:
	scheduler.proto

It has these top-level messages:
	Resources
	Node
	GetNodeRequest
	ListNodesRequest
	ListNodesResponse
	PutNodeResponse
	DeleteNodeResponse
*/
package scheduler

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type NodeState int32

const (
	NodeState_UNINITIALIZED NodeState = 0
	NodeState_ALIVE         NodeState = 1
	NodeState_DEAD          NodeState = 2
	NodeState_GONE          NodeState = 3
	NodeState_INITIALIZING  NodeState = 4
)

var NodeState_name = map[int32]string{
	0: "UNINITIALIZED",
	1: "ALIVE",
	2: "DEAD",
	3: "GONE",
	4: "INITIALIZING",
}
var NodeState_value = map[string]int32{
	"UNINITIALIZED": 0,
	"ALIVE":         1,
	"DEAD":          2,
	"GONE":          3,
	"INITIALIZING":  4,
}

func (x NodeState) String() string {
	return proto.EnumName(NodeState_name, int32(x))
}
func (NodeState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Resources struct {
	Cpus uint32 `protobuf:"varint,1,opt,name=cpus" json:"cpus,omitempty"`
	// In GB
	RamGb float64 `protobuf:"fixed64,2,opt,name=ram_gb,json=ramGb" json:"ram_gb,omitempty"`
	// In GB
	DiskGb float64 `protobuf:"fixed64,3,opt,name=disk_gb,json=diskGb" json:"disk_gb,omitempty"`
}

func (m *Resources) Reset()                    { *m = Resources{} }
func (m *Resources) String() string            { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()               {}
func (*Resources) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Resources) GetCpus() uint32 {
	if m != nil {
		return m.Cpus
	}
	return 0
}

func (m *Resources) GetRamGb() float64 {
	if m != nil {
		return m.RamGb
	}
	return 0
}

func (m *Resources) GetDiskGb() float64 {
	if m != nil {
		return m.DiskGb
	}
	return 0
}

type Node struct {
	Id          string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Resources   *Resources `protobuf:"bytes,5,opt,name=resources" json:"resources,omitempty"`
	Available   *Resources `protobuf:"bytes,6,opt,name=available" json:"available,omitempty"`
	State       NodeState  `protobuf:"varint,8,opt,name=state,enum=scheduler.NodeState" json:"state,omitempty"`
	Preemptible bool       `protobuf:"varint,9,opt,name=preemptible" json:"preemptible,omitempty"`
	Zone        string     `protobuf:"bytes,11,opt,name=zone" json:"zone,omitempty"`
	Hostname    string     `protobuf:"bytes,13,opt,name=hostname" json:"hostname,omitempty"`
	// Timestamp version of the record in the database. Used to prevent write conflicts and as the last ping time.
	Version  int64             `protobuf:"varint,14,opt,name=version" json:"version,omitempty"`
	Metadata map[string]string `protobuf:"bytes,15,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	TaskIds  []string          `protobuf:"bytes,16,rep,name=task_ids,json=taskIds" json:"task_ids,omitempty"`
	LastPing int64             `protobuf:"varint,17,opt,name=last_ping,json=lastPing" json:"last_ping,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Node) GetAvailable() *Resources {
	if m != nil {
		return m.Available
	}
	return nil
}

func (m *Node) GetState() NodeState {
	if m != nil {
		return m.State
	}
	return NodeState_UNINITIALIZED
}

func (m *Node) GetPreemptible() bool {
	if m != nil {
		return m.Preemptible
	}
	return false
}

func (m *Node) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Node) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Node) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetTaskIds() []string {
	if m != nil {
		return m.TaskIds
	}
	return nil
}

func (m *Node) GetLastPing() int64 {
	if m != nil {
		return m.LastPing
	}
	return 0
}

type GetNodeRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetNodeRequest) Reset()                    { *m = GetNodeRequest{} }
func (m *GetNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeRequest) ProtoMessage()               {}
func (*GetNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetNodeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListNodesRequest struct {
}

func (m *ListNodesRequest) Reset()                    { *m = ListNodesRequest{} }
func (m *ListNodesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListNodesRequest) ProtoMessage()               {}
func (*ListNodesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ListNodesResponse struct {
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *ListNodesResponse) Reset()                    { *m = ListNodesResponse{} }
func (m *ListNodesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListNodesResponse) ProtoMessage()               {}
func (*ListNodesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type PutNodeResponse struct {
}

func (m *PutNodeResponse) Reset()                    { *m = PutNodeResponse{} }
func (m *PutNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*PutNodeResponse) ProtoMessage()               {}
func (*PutNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteNodeResponse struct {
}

func (m *DeleteNodeResponse) Reset()                    { *m = DeleteNodeResponse{} }
func (m *DeleteNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeResponse) ProtoMessage()               {}
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Resources)(nil), "scheduler.Resources")
	proto.RegisterType((*Node)(nil), "scheduler.Node")
	proto.RegisterType((*GetNodeRequest)(nil), "scheduler.GetNodeRequest")
	proto.RegisterType((*ListNodesRequest)(nil), "scheduler.ListNodesRequest")
	proto.RegisterType((*ListNodesResponse)(nil), "scheduler.ListNodesResponse")
	proto.RegisterType((*PutNodeResponse)(nil), "scheduler.PutNodeResponse")
	proto.RegisterType((*DeleteNodeResponse)(nil), "scheduler.DeleteNodeResponse")
	proto.RegisterEnum("scheduler.NodeState", NodeState_name, NodeState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SchedulerService service

type SchedulerServiceClient interface {
	PutNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*PutNodeResponse, error)
	DeleteNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*DeleteNodeResponse, error)
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error)
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error)
}

type schedulerServiceClient struct {
	cc *grpc.ClientConn
}

func NewSchedulerServiceClient(cc *grpc.ClientConn) SchedulerServiceClient {
	return &schedulerServiceClient{cc}
}

func (c *schedulerServiceClient) PutNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*PutNodeResponse, error) {
	out := new(PutNodeResponse)
	err := grpc.Invoke(ctx, "/scheduler.SchedulerService/PutNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) DeleteNode(ctx context.Context, in *Node, opts ...grpc.CallOption) (*DeleteNodeResponse, error) {
	out := new(DeleteNodeResponse)
	err := grpc.Invoke(ctx, "/scheduler.SchedulerService/DeleteNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := grpc.Invoke(ctx, "/scheduler.SchedulerService/ListNodes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := grpc.Invoke(ctx, "/scheduler.SchedulerService/GetNode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SchedulerService service

type SchedulerServiceServer interface {
	PutNode(context.Context, *Node) (*PutNodeResponse, error)
	DeleteNode(context.Context, *Node) (*DeleteNodeResponse, error)
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error)
	GetNode(context.Context, *GetNodeRequest) (*Node, error)
}

func RegisterSchedulerServiceServer(s *grpc.Server, srv SchedulerServiceServer) {
	s.RegisterService(&_SchedulerService_serviceDesc, srv)
}

func _SchedulerService_PutNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).PutNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/PutNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).PutNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_DeleteNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).DeleteNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/DeleteNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).DeleteNode(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.SchedulerService/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SchedulerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.SchedulerService",
	HandlerType: (*SchedulerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutNode",
			Handler:    _SchedulerService_PutNode_Handler,
		},
		{
			MethodName: "DeleteNode",
			Handler:    _SchedulerService_DeleteNode_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _SchedulerService_ListNodes_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _SchedulerService_GetNode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduler.proto",
}

func init() { proto.RegisterFile("scheduler.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6a, 0xdb, 0x4a,
	0x10, 0x8e, 0x24, 0xff, 0x48, 0xe3, 0x63, 0x5b, 0x1e, 0x72, 0xce, 0x51, 0x94, 0x1c, 0x10, 0x82,
	0x03, 0x22, 0x17, 0x31, 0x75, 0x6f, 0xd2, 0x14, 0x0a, 0x01, 0x1b, 0x23, 0x48, 0x9c, 0xa0, 0xb4,
	0x85, 0xf6, 0x26, 0xac, 0xad, 0xc5, 0x59, 0x22, 0x4b, 0xaa, 0x76, 0x6d, 0x48, 0x4b, 0x6f, 0xfa,
	0x0a, 0x7d, 0x97, 0x3e, 0x42, 0x5f, 0xa0, 0xaf, 0xd0, 0x07, 0x29, 0x5a, 0x2b, 0x8a, 0xa2, 0x94,
	0xde, 0xed, 0x7c, 0xf3, 0xcd, 0xcc, 0x37, 0xf3, 0xc1, 0x42, 0x9f, 0x2f, 0x6e, 0x68, 0xb8, 0x8e,
	0x68, 0x76, 0x94, 0x66, 0x89, 0x48, 0xd0, 0x28, 0x01, 0xfb, 0x60, 0x99, 0x24, 0xcb, 0x88, 0x0e,
	0x49, 0xca, 0x86, 0x24, 0x8e, 0x13, 0x41, 0x04, 0x4b, 0x62, 0xbe, 0x25, 0xba, 0x17, 0x60, 0x04,
	0x94, 0x27, 0xeb, 0x6c, 0x41, 0x39, 0x22, 0x34, 0x16, 0xe9, 0x9a, 0x5b, 0x8a, 0xa3, 0x78, 0xdd,
	0x40, 0xbe, 0xf1, 0x6f, 0x68, 0x65, 0x64, 0x75, 0xbd, 0x9c, 0x5b, 0xaa, 0xa3, 0x78, 0x4a, 0xd0,
	0xcc, 0xc8, 0x6a, 0x3a, 0xc7, 0x7f, 0xa1, 0x1d, 0x32, 0x7e, 0x9b, 0xe3, 0x9a, 0xc4, 0x5b, 0x79,
	0x38, 0x9d, 0xbb, 0xdf, 0x35, 0x68, 0xcc, 0x92, 0x90, 0x62, 0x0f, 0x54, 0x16, 0xca, 0x56, 0x46,
	0xa0, 0xb2, 0x10, 0x47, 0x60, 0x64, 0xf7, 0x93, 0xac, 0xa6, 0xa3, 0x78, 0x9d, 0xd1, 0xee, 0xd1,
	0x83, 0xee, 0x52, 0x45, 0xf0, 0x40, 0xcb, 0x6b, 0xc8, 0x86, 0xb0, 0x88, 0xcc, 0x23, 0x6a, 0xb5,
	0xfe, 0x54, 0x53, 0xd2, 0xf0, 0x10, 0x9a, 0x5c, 0x10, 0x41, 0x2d, 0xdd, 0x51, 0xbc, 0xde, 0x23,
	0x7e, 0xae, 0xeb, 0x2a, 0xcf, 0x05, 0x5b, 0x0a, 0x3a, 0xd0, 0x49, 0x33, 0x4a, 0x57, 0xa9, 0x60,
	0xf9, 0x04, 0xc3, 0x51, 0x3c, 0x3d, 0xa8, 0x42, 0xf9, 0x49, 0x3e, 0x26, 0x31, 0xb5, 0x3a, 0x72,
	0x0f, 0xf9, 0x46, 0x1b, 0xf4, 0x9b, 0x84, 0x8b, 0x98, 0xac, 0xa8, 0xd5, 0x95, 0x78, 0x19, 0xa3,
	0x05, 0xed, 0x0d, 0xcd, 0x38, 0x4b, 0x62, 0xab, 0xe7, 0x28, 0x9e, 0x16, 0xdc, 0x87, 0xf8, 0x02,
	0xf4, 0x15, 0x15, 0x24, 0x24, 0x82, 0x58, 0x7d, 0x47, 0xf3, 0x3a, 0xa3, 0xff, 0x6a, 0xd2, 0x8e,
	0xce, 0x8b, 0xfc, 0x24, 0x16, 0xd9, 0x5d, 0x50, 0xd2, 0x71, 0x0f, 0x74, 0x41, 0xf8, 0xed, 0x35,
	0x0b, 0xb9, 0x65, 0x3a, 0x9a, 0x67, 0x04, 0xed, 0x3c, 0xf6, 0x43, 0x8e, 0xfb, 0x60, 0x44, 0x84,
	0x8b, 0xeb, 0x94, 0xc5, 0x4b, 0x6b, 0x20, 0x27, 0xea, 0x39, 0x70, 0xc9, 0xe2, 0xa5, 0xfd, 0x12,
	0xba, 0x8f, 0x5a, 0xa2, 0x09, 0xda, 0x2d, 0xbd, 0x2b, 0x4c, 0xc9, 0x9f, 0xb8, 0x0b, 0xcd, 0x0d,
	0x89, 0xd6, 0x54, 0xba, 0x6b, 0x04, 0xdb, 0xe0, 0x44, 0x3d, 0x56, 0x5c, 0x07, 0x7a, 0x53, 0x2a,
	0x72, 0x5d, 0x01, 0xfd, 0xb0, 0xa6, 0x5c, 0xd4, 0x1d, 0x75, 0x11, 0xcc, 0x33, 0xc6, 0x25, 0x85,
	0x17, 0x1c, 0xf7, 0x04, 0x06, 0x15, 0x8c, 0xa7, 0x49, 0xcc, 0x29, 0xfe, 0x0f, 0xcd, 0x38, 0x07,
	0x2c, 0x45, 0xee, 0xdd, 0xaf, 0xed, 0x1d, 0x6c, 0xb3, 0xee, 0x00, 0xfa, 0x97, 0xeb, 0x62, 0xe2,
	0xb6, 0xd2, 0xdd, 0x05, 0x1c, 0xd3, 0x88, 0x0a, 0x5a, 0x45, 0x0f, 0x2f, 0xc0, 0x28, 0xad, 0xc4,
	0x01, 0x74, 0xdf, 0xcc, 0xfc, 0x99, 0xff, 0xda, 0x3f, 0x3d, 0xf3, 0xdf, 0x4f, 0xc6, 0xe6, 0x0e,
	0x1a, 0xd0, 0x3c, 0x3d, 0xf3, 0xdf, 0x4e, 0x4c, 0x05, 0x75, 0x68, 0x8c, 0x27, 0xa7, 0x63, 0x53,
	0xcd, 0x5f, 0xd3, 0x8b, 0xd9, 0xc4, 0xd4, 0xd0, 0x84, 0xbf, 0x4a, 0xbe, 0x3f, 0x9b, 0x9a, 0x8d,
	0xd1, 0x37, 0x15, 0xcc, 0xab, 0x7b, 0x4d, 0x57, 0x34, 0xdb, 0xb0, 0x05, 0xc5, 0x63, 0x68, 0x17,
	0x72, 0xb0, 0xae, 0xd8, 0xb6, 0x2b, 0x40, 0x5d, 0xf3, 0x0e, 0xbe, 0x02, 0x78, 0x50, 0xfd, 0xb4,
	0xb8, 0xea, 0xfb, 0xd3, 0xed, 0xdc, 0x1d, 0x7c, 0x07, 0x46, 0x79, 0x44, 0xdc, 0xaf, 0xb0, 0xeb,
	0xe7, 0xb6, 0x0f, 0x7e, 0x9f, 0x2c, 0x3a, 0x0d, 0xbe, 0xfc, 0xf8, 0xf9, 0x55, 0xed, 0xa0, 0x31,
	0xdc, 0x3c, 0x1b, 0xca, 0x1b, 0xe3, 0x39, 0xb4, 0x0b, 0x57, 0x71, 0xaf, 0x52, 0xfb, 0xd8, 0x69,
	0xbb, 0x2e, 0xd9, 0xfd, 0x47, 0x76, 0x32, 0xb1, 0x57, 0x76, 0x1a, 0x7e, 0x62, 0xe1, 0xe7, 0x79,
	0x4b, 0xfe, 0x22, 0xcf, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x13, 0xf9, 0x0e, 0x81, 0x04,
	0x00, 0x00,
}
