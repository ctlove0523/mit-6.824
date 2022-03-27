// Code generated by protoc-gen-go. DO NOT EDIT.
// source: coordinator.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type RegisterWorkerRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterWorkerRequest) Reset()         { *m = RegisterWorkerRequest{} }
func (m *RegisterWorkerRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterWorkerRequest) ProtoMessage()    {}
func (*RegisterWorkerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinator_a9b3f6f0ee6ec9cc, []int{0}
}
func (m *RegisterWorkerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterWorkerRequest.Unmarshal(m, b)
}
func (m *RegisterWorkerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterWorkerRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterWorkerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterWorkerRequest.Merge(dst, src)
}
func (m *RegisterWorkerRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterWorkerRequest.Size(m)
}
func (m *RegisterWorkerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterWorkerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterWorkerRequest proto.InternalMessageInfo

func (m *RegisterWorkerRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type RegisterWorkerResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CoordinatorAddress   string   `protobuf:"bytes,2,opt,name=coordinator_address,json=coordinatorAddress,proto3" json:"coordinator_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterWorkerResponse) Reset()         { *m = RegisterWorkerResponse{} }
func (m *RegisterWorkerResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterWorkerResponse) ProtoMessage()    {}
func (*RegisterWorkerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinator_a9b3f6f0ee6ec9cc, []int{1}
}
func (m *RegisterWorkerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterWorkerResponse.Unmarshal(m, b)
}
func (m *RegisterWorkerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterWorkerResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterWorkerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterWorkerResponse.Merge(dst, src)
}
func (m *RegisterWorkerResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterWorkerResponse.Size(m)
}
func (m *RegisterWorkerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterWorkerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterWorkerResponse proto.InternalMessageInfo

func (m *RegisterWorkerResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterWorkerResponse) GetCoordinatorAddress() string {
	if m != nil {
		return m.CoordinatorAddress
	}
	return ""
}

type CreateMapReduceTaskRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Inputs               []string `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty"`
	MapSize              int32    `protobuf:"varint,3,opt,name=map_size,json=mapSize,proto3" json:"map_size,omitempty"`
	ReduceSize           int32    `protobuf:"varint,4,opt,name=reduce_size,json=reduceSize,proto3" json:"reduce_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMapReduceTaskRequest) Reset()         { *m = CreateMapReduceTaskRequest{} }
func (m *CreateMapReduceTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateMapReduceTaskRequest) ProtoMessage()    {}
func (*CreateMapReduceTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinator_a9b3f6f0ee6ec9cc, []int{2}
}
func (m *CreateMapReduceTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMapReduceTaskRequest.Unmarshal(m, b)
}
func (m *CreateMapReduceTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMapReduceTaskRequest.Marshal(b, m, deterministic)
}
func (dst *CreateMapReduceTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMapReduceTaskRequest.Merge(dst, src)
}
func (m *CreateMapReduceTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMapReduceTaskRequest.Size(m)
}
func (m *CreateMapReduceTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMapReduceTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMapReduceTaskRequest proto.InternalMessageInfo

func (m *CreateMapReduceTaskRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateMapReduceTaskRequest) GetInputs() []string {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *CreateMapReduceTaskRequest) GetMapSize() int32 {
	if m != nil {
		return m.MapSize
	}
	return 0
}

func (m *CreateMapReduceTaskRequest) GetReduceSize() int32 {
	if m != nil {
		return m.ReduceSize
	}
	return 0
}

type CreateMapReduceTaskResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMapReduceTaskResponse) Reset()         { *m = CreateMapReduceTaskResponse{} }
func (m *CreateMapReduceTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMapReduceTaskResponse) ProtoMessage()    {}
func (*CreateMapReduceTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinator_a9b3f6f0ee6ec9cc, []int{3}
}
func (m *CreateMapReduceTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMapReduceTaskResponse.Unmarshal(m, b)
}
func (m *CreateMapReduceTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMapReduceTaskResponse.Marshal(b, m, deterministic)
}
func (dst *CreateMapReduceTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMapReduceTaskResponse.Merge(dst, src)
}
func (m *CreateMapReduceTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMapReduceTaskResponse.Size(m)
}
func (m *CreateMapReduceTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMapReduceTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMapReduceTaskResponse proto.InternalMessageInfo

func (m *CreateMapReduceTaskResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReportMapTaskProgressRequest struct {
	TaskId               string            `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Id                   uint32            `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	State                uint32            `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
	Output               map[uint32]string `protobuf:"bytes,4,rep,name=output,proto3" json:"output,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReportMapTaskProgressRequest) Reset()         { *m = ReportMapTaskProgressRequest{} }
func (m *ReportMapTaskProgressRequest) String() string { return proto.CompactTextString(m) }
func (*ReportMapTaskProgressRequest) ProtoMessage()    {}
func (*ReportMapTaskProgressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinator_a9b3f6f0ee6ec9cc, []int{4}
}
func (m *ReportMapTaskProgressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportMapTaskProgressRequest.Unmarshal(m, b)
}
func (m *ReportMapTaskProgressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportMapTaskProgressRequest.Marshal(b, m, deterministic)
}
func (dst *ReportMapTaskProgressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportMapTaskProgressRequest.Merge(dst, src)
}
func (m *ReportMapTaskProgressRequest) XXX_Size() int {
	return xxx_messageInfo_ReportMapTaskProgressRequest.Size(m)
}
func (m *ReportMapTaskProgressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportMapTaskProgressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportMapTaskProgressRequest proto.InternalMessageInfo

func (m *ReportMapTaskProgressRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *ReportMapTaskProgressRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReportMapTaskProgressRequest) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *ReportMapTaskProgressRequest) GetOutput() map[uint32]string {
	if m != nil {
		return m.Output
	}
	return nil
}

type ReportMapTaskProgressResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportMapTaskProgressResponse) Reset()         { *m = ReportMapTaskProgressResponse{} }
func (m *ReportMapTaskProgressResponse) String() string { return proto.CompactTextString(m) }
func (*ReportMapTaskProgressResponse) ProtoMessage()    {}
func (*ReportMapTaskProgressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinator_a9b3f6f0ee6ec9cc, []int{5}
}
func (m *ReportMapTaskProgressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportMapTaskProgressResponse.Unmarshal(m, b)
}
func (m *ReportMapTaskProgressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportMapTaskProgressResponse.Marshal(b, m, deterministic)
}
func (dst *ReportMapTaskProgressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportMapTaskProgressResponse.Merge(dst, src)
}
func (m *ReportMapTaskProgressResponse) XXX_Size() int {
	return xxx_messageInfo_ReportMapTaskProgressResponse.Size(m)
}
func (m *ReportMapTaskProgressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportMapTaskProgressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportMapTaskProgressResponse proto.InternalMessageInfo

func (m *ReportMapTaskProgressResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type ReportReduceTaskProgressRequest struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	State                uint32   `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
	Output               string   `protobuf:"bytes,4,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportReduceTaskProgressRequest) Reset()         { *m = ReportReduceTaskProgressRequest{} }
func (m *ReportReduceTaskProgressRequest) String() string { return proto.CompactTextString(m) }
func (*ReportReduceTaskProgressRequest) ProtoMessage()    {}
func (*ReportReduceTaskProgressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinator_a9b3f6f0ee6ec9cc, []int{6}
}
func (m *ReportReduceTaskProgressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportReduceTaskProgressRequest.Unmarshal(m, b)
}
func (m *ReportReduceTaskProgressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportReduceTaskProgressRequest.Marshal(b, m, deterministic)
}
func (dst *ReportReduceTaskProgressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportReduceTaskProgressRequest.Merge(dst, src)
}
func (m *ReportReduceTaskProgressRequest) XXX_Size() int {
	return xxx_messageInfo_ReportReduceTaskProgressRequest.Size(m)
}
func (m *ReportReduceTaskProgressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportReduceTaskProgressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportReduceTaskProgressRequest proto.InternalMessageInfo

func (m *ReportReduceTaskProgressRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *ReportReduceTaskProgressRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReportReduceTaskProgressRequest) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *ReportReduceTaskProgressRequest) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

type ReportReduceTaskProgressResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportReduceTaskProgressResponse) Reset()         { *m = ReportReduceTaskProgressResponse{} }
func (m *ReportReduceTaskProgressResponse) String() string { return proto.CompactTextString(m) }
func (*ReportReduceTaskProgressResponse) ProtoMessage()    {}
func (*ReportReduceTaskProgressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinator_a9b3f6f0ee6ec9cc, []int{7}
}
func (m *ReportReduceTaskProgressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportReduceTaskProgressResponse.Unmarshal(m, b)
}
func (m *ReportReduceTaskProgressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportReduceTaskProgressResponse.Marshal(b, m, deterministic)
}
func (dst *ReportReduceTaskProgressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportReduceTaskProgressResponse.Merge(dst, src)
}
func (m *ReportReduceTaskProgressResponse) XXX_Size() int {
	return xxx_messageInfo_ReportReduceTaskProgressResponse.Size(m)
}
func (m *ReportReduceTaskProgressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportReduceTaskProgressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportReduceTaskProgressResponse proto.InternalMessageInfo

func (m *ReportReduceTaskProgressResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*RegisterWorkerRequest)(nil), "RegisterWorkerRequest")
	proto.RegisterType((*RegisterWorkerResponse)(nil), "RegisterWorkerResponse")
	proto.RegisterType((*CreateMapReduceTaskRequest)(nil), "CreateMapReduceTaskRequest")
	proto.RegisterType((*CreateMapReduceTaskResponse)(nil), "CreateMapReduceTaskResponse")
	proto.RegisterType((*ReportMapTaskProgressRequest)(nil), "ReportMapTaskProgressRequest")
	proto.RegisterMapType((map[uint32]string)(nil), "ReportMapTaskProgressRequest.OutputEntry")
	proto.RegisterType((*ReportMapTaskProgressResponse)(nil), "ReportMapTaskProgressResponse")
	proto.RegisterType((*ReportReduceTaskProgressRequest)(nil), "ReportReduceTaskProgressRequest")
	proto.RegisterType((*ReportReduceTaskProgressResponse)(nil), "ReportReduceTaskProgressResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the samples package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoordinatorServerClient is the client API for CoordinatorServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoordinatorServerClient interface {
	RegisterWorker(ctx context.Context, in *RegisterWorkerRequest, opts ...grpc.CallOption) (*RegisterWorkerResponse, error)
	CreateTask(ctx context.Context, in *CreateMapReduceTaskRequest, opts ...grpc.CallOption) (*CreateMapReduceTaskResponse, error)
	ReportMapTaskProgress(ctx context.Context, in *ReportMapTaskProgressRequest, opts ...grpc.CallOption) (*ReportMapTaskProgressResponse, error)
	ReportReduceTaskProgress(ctx context.Context, in *ReportReduceTaskProgressRequest, opts ...grpc.CallOption) (*ReportReduceTaskProgressResponse, error)
}

type coordinatorServerClient struct {
	cc *grpc.ClientConn
}

func NewCoordinatorServerClient(cc *grpc.ClientConn) CoordinatorServerClient {
	return &coordinatorServerClient{cc}
}

func (c *coordinatorServerClient) RegisterWorker(ctx context.Context, in *RegisterWorkerRequest, opts ...grpc.CallOption) (*RegisterWorkerResponse, error) {
	out := new(RegisterWorkerResponse)
	err := c.cc.Invoke(ctx, "/CoordinatorServer/RegisterWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorServerClient) CreateTask(ctx context.Context, in *CreateMapReduceTaskRequest, opts ...grpc.CallOption) (*CreateMapReduceTaskResponse, error) {
	out := new(CreateMapReduceTaskResponse)
	err := c.cc.Invoke(ctx, "/CoordinatorServer/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorServerClient) ReportMapTaskProgress(ctx context.Context, in *ReportMapTaskProgressRequest, opts ...grpc.CallOption) (*ReportMapTaskProgressResponse, error) {
	out := new(ReportMapTaskProgressResponse)
	err := c.cc.Invoke(ctx, "/CoordinatorServer/ReportMapTaskProgress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coordinatorServerClient) ReportReduceTaskProgress(ctx context.Context, in *ReportReduceTaskProgressRequest, opts ...grpc.CallOption) (*ReportReduceTaskProgressResponse, error) {
	out := new(ReportReduceTaskProgressResponse)
	err := c.cc.Invoke(ctx, "/CoordinatorServer/ReportReduceTaskProgress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServerServer is the server API for CoordinatorServer service.
type CoordinatorServerServer interface {
	RegisterWorker(context.Context, *RegisterWorkerRequest) (*RegisterWorkerResponse, error)
	CreateTask(context.Context, *CreateMapReduceTaskRequest) (*CreateMapReduceTaskResponse, error)
	ReportMapTaskProgress(context.Context, *ReportMapTaskProgressRequest) (*ReportMapTaskProgressResponse, error)
	ReportReduceTaskProgress(context.Context, *ReportReduceTaskProgressRequest) (*ReportReduceTaskProgressResponse, error)
}

func RegisterCoordinatorServerServer(s *grpc.Server, srv CoordinatorServerServer) {
	s.RegisterService(&_CoordinatorServer_serviceDesc, srv)
}

func _CoordinatorServer_RegisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServerServer).RegisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoordinatorServer/RegisterWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServerServer).RegisterWorker(ctx, req.(*RegisterWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoordinatorServer_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMapReduceTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoordinatorServer/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServerServer).CreateTask(ctx, req.(*CreateMapReduceTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoordinatorServer_ReportMapTaskProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportMapTaskProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServerServer).ReportMapTaskProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoordinatorServer/ReportMapTaskProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServerServer).ReportMapTaskProgress(ctx, req.(*ReportMapTaskProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoordinatorServer_ReportReduceTaskProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportReduceTaskProgressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServerServer).ReportReduceTaskProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoordinatorServer/ReportReduceTaskProgress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServerServer).ReportReduceTaskProgress(ctx, req.(*ReportReduceTaskProgressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CoordinatorServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CoordinatorServer",
	HandlerType: (*CoordinatorServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterWorker",
			Handler:    _CoordinatorServer_RegisterWorker_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _CoordinatorServer_CreateTask_Handler,
		},
		{
			MethodName: "ReportMapTaskProgress",
			Handler:    _CoordinatorServer_ReportMapTaskProgress_Handler,
		},
		{
			MethodName: "ReportReduceTaskProgress",
			Handler:    _CoordinatorServer_ReportReduceTaskProgress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coordinator.proto",
}

func init() { proto.RegisterFile("coordinator.proto", fileDescriptor_coordinator_a9b3f6f0ee6ec9cc) }

var fileDescriptor_coordinator_a9b3f6f0ee6ec9cc = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4b, 0x6f, 0xd4, 0x3c,
	0x14, 0x9d, 0xcc, 0x23, 0xd3, 0xde, 0xd1, 0x54, 0x5f, 0xfd, 0x41, 0x1a, 0xd2, 0x96, 0x06, 0xaf,
	0x86, 0x05, 0x46, 0x94, 0x05, 0xd0, 0x5d, 0x19, 0xb1, 0xe8, 0xa2, 0x02, 0xb9, 0x48, 0x3c, 0x36,
	0x23, 0xd3, 0x5c, 0x55, 0xd1, 0x74, 0x62, 0x63, 0x3b, 0x15, 0xed, 0x9a, 0x7f, 0xc9, 0x86, 0x9f,
	0x82, 0xec, 0x64, 0xe8, 0x80, 0x32, 0xe9, 0x86, 0x5d, 0x8e, 0x7d, 0x1f, 0xc7, 0xe7, 0x9e, 0x1b,
	0xd8, 0x3e, 0x97, 0x52, 0x67, 0x79, 0x21, 0xac, 0xd4, 0x4c, 0x69, 0x69, 0x25, 0x7d, 0x06, 0xf7,
	0x39, 0x5e, 0xe4, 0xc6, 0xa2, 0xfe, 0x20, 0xf5, 0x1c, 0x35, 0xc7, 0xaf, 0x25, 0x1a, 0x4b, 0x62,
	0x18, 0x8a, 0x2c, 0xd3, 0x68, 0x4c, 0x1c, 0xa4, 0xc1, 0x64, 0x93, 0x2f, 0x21, 0xfd, 0x04, 0xd1,
	0xdf, 0x29, 0x46, 0xc9, 0xc2, 0x20, 0xd9, 0x82, 0x6e, 0x9e, 0xd5, 0xe1, 0xdd, 0x3c, 0x23, 0x4f,
	0xe1, 0xff, 0x95, 0x8e, 0xb3, 0x65, 0xbd, 0xae, 0x0f, 0x20, 0x2b, 0x57, 0xc7, 0x75, 0xe9, 0xef,
	0x01, 0x24, 0x53, 0x8d, 0xc2, 0xe2, 0xa9, 0x50, 0x1c, 0xb3, 0xf2, 0x1c, 0xdf, 0x0b, 0x33, 0x5f,
	0x72, 0x22, 0xd0, 0x2f, 0xc4, 0x02, 0xeb, 0x0e, 0xfe, 0x9b, 0x44, 0x10, 0xe6, 0x85, 0x2a, 0xad,
	0x2b, 0xdb, 0x9b, 0x6c, 0xf2, 0x1a, 0x91, 0x07, 0xb0, 0xb1, 0x10, 0x6a, 0x66, 0xf2, 0x1b, 0x8c,
	0x7b, 0x69, 0x30, 0x19, 0xf0, 0xe1, 0x42, 0xa8, 0xb3, 0xfc, 0x06, 0xc9, 0x01, 0x8c, 0xb4, 0xaf,
	0x5d, 0xdd, 0xf6, 0xfd, 0x2d, 0x54, 0x47, 0x2e, 0x80, 0x3e, 0x81, 0xdd, 0x46, 0x16, 0xcd, 0xcf,
	0xa4, 0x3f, 0x02, 0xd8, 0xe3, 0xa8, 0xa4, 0xb6, 0xa7, 0x42, 0xb9, 0xc8, 0x77, 0x5a, 0x5e, 0xb8,
	0xf7, 0x2c, 0x79, 0xef, 0xc0, 0xd0, 0x0a, 0x33, 0x9f, 0xfd, 0xce, 0x0a, 0x1d, 0x3c, 0xc9, 0xea,
	0x4a, 0x4e, 0x8f, 0xb1, 0x17, 0xec, 0x1e, 0x0c, 0x8c, 0x15, 0xb6, 0x62, 0x3c, 0xe6, 0x15, 0x20,
	0xc7, 0x10, 0xca, 0xd2, 0xaa, 0xd2, 0xc6, 0xfd, 0xb4, 0x37, 0x19, 0x1d, 0x3e, 0x66, 0x6d, 0xdd,
	0xd8, 0x5b, 0x1f, 0xfb, 0xa6, 0xb0, 0xfa, 0x9a, 0xd7, 0x89, 0xc9, 0x2b, 0x18, 0xad, 0x1c, 0x93,
	0xff, 0xa0, 0x37, 0xc7, 0x6b, 0x4f, 0x66, 0xcc, 0xdd, 0xa7, 0xeb, 0x7c, 0x25, 0x2e, 0x4b, 0xac,
	0x87, 0x53, 0x81, 0xa3, 0xee, 0xcb, 0x80, 0xbe, 0x80, 0xfd, 0x35, 0xed, 0x6a, 0x39, 0x22, 0x08,
	0x35, 0x9a, 0xf2, 0xd2, 0xfa, 0x7a, 0x1b, 0xbc, 0x46, 0xf4, 0x1b, 0x1c, 0x54, 0x89, 0xb7, 0x12,
	0xfe, 0x63, 0x61, 0xa2, 0x15, 0x61, 0x7c, 0x76, 0x85, 0xe8, 0x11, 0xa4, 0xeb, 0x3b, 0xb7, 0xb3,
	0x3e, 0xfc, 0xd9, 0x85, 0xed, 0xe9, 0xad, 0x33, 0xcf, 0x50, 0x5f, 0xa1, 0x26, 0x53, 0xd8, 0xfa,
	0xd3, 0xf3, 0x24, 0x62, 0x8d, 0x7b, 0x93, 0xec, 0xb0, 0xe6, 0xe5, 0xa0, 0x1d, 0x72, 0x02, 0x50,
	0xd9, 0xca, 0x11, 0x22, 0xbb, 0x6c, 0xbd, 0xd3, 0x93, 0x3d, 0xd6, 0x62, 0x40, 0xda, 0x21, 0x1f,
	0xdd, 0xda, 0x36, 0x0c, 0x85, 0xec, 0xb7, 0x7a, 0x23, 0x79, 0xc8, 0x5a, 0x67, 0x49, 0x3b, 0x44,
	0x40, 0xbc, 0x4e, 0x3b, 0x92, 0xb2, 0x3b, 0x06, 0x9a, 0x3c, 0x62, 0x77, 0x09, 0x4f, 0x3b, 0xaf,
	0x07, 0x9f, 0x7b, 0x42, 0xe5, 0x5f, 0x42, 0xff, 0x07, 0x7a, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x23, 0x35, 0xe5, 0x58, 0x96, 0x04, 0x00, 0x00,
}