// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comapi-service.proto

package protobuffer_def // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

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

type CMD int32

const (
	CMD_QUERY_USERS_BY_ROOM CMD = 0
)

var CMD_name = map[int32]string{
	0: "QUERY_USERS_BY_ROOM",
}
var CMD_value = map[string]int32{
	"QUERY_USERS_BY_ROOM": 0,
}

func (x CMD) String() string {
	return proto.EnumName(CMD_name, int32(x))
}
func (CMD) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_comapi_service_3a978815dda79e72, []int{0}
}

type ReturnCode int32

const (
	ReturnCode_SUCCESS               ReturnCode = 0
	ReturnCode_UNKOWN_CMD            ReturnCode = 1
	ReturnCode_BODY_IS_NULL          ReturnCode = 2
	ReturnCode_DESERIALIZATION_ERROR ReturnCode = 3
	ReturnCode_SERIALIZATION_ERROR   ReturnCode = 4
	ReturnCode_UNKOWN_ERROR          ReturnCode = 5
)

var ReturnCode_name = map[int32]string{
	0: "SUCCESS",
	1: "UNKOWN_CMD",
	2: "BODY_IS_NULL",
	3: "DESERIALIZATION_ERROR",
	4: "SERIALIZATION_ERROR",
	5: "UNKOWN_ERROR",
}
var ReturnCode_value = map[string]int32{
	"SUCCESS":               0,
	"UNKOWN_CMD":            1,
	"BODY_IS_NULL":          2,
	"DESERIALIZATION_ERROR": 3,
	"SERIALIZATION_ERROR":   4,
	"UNKOWN_ERROR":          5,
}

func (x ReturnCode) String() string {
	return proto.EnumName(ReturnCode_name, int32(x))
}
func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_comapi_service_3a978815dda79e72, []int{1}
}

type BaseRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	C                    CMD      `protobuf:"varint,2,opt,name=c,proto3,enum=CMD" json:"c,omitempty"`
	Body                 *any.Any `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseRequest) Reset()         { *m = BaseRequest{} }
func (m *BaseRequest) String() string { return proto.CompactTextString(m) }
func (*BaseRequest) ProtoMessage()    {}
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_comapi_service_3a978815dda79e72, []int{0}
}
func (m *BaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRequest.Unmarshal(m, b)
}
func (m *BaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRequest.Marshal(b, m, deterministic)
}
func (dst *BaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRequest.Merge(dst, src)
}
func (m *BaseRequest) XXX_Size() int {
	return xxx_messageInfo_BaseRequest.Size(m)
}
func (m *BaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRequest proto.InternalMessageInfo

func (m *BaseRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *BaseRequest) GetC() CMD {
	if m != nil {
		return m.C
	}
	return CMD_QUERY_USERS_BY_ROOM
}

func (m *BaseRequest) GetBody() *any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

type BaseResponse struct {
	RequestId            string     `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Code                 ReturnCode `protobuf:"varint,2,opt,name=code,proto3,enum=ReturnCode" json:"code,omitempty"`
	Desc                 string     `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	C                    CMD        `protobuf:"varint,4,opt,name=c,proto3,enum=CMD" json:"c,omitempty"`
	Body                 *any.Any   `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_comapi_service_3a978815dda79e72, []int{1}
}
func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (dst *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(dst, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *BaseResponse) GetCode() ReturnCode {
	if m != nil {
		return m.Code
	}
	return ReturnCode_SUCCESS
}

func (m *BaseResponse) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *BaseResponse) GetC() CMD {
	if m != nil {
		return m.C
	}
	return CMD_QUERY_USERS_BY_ROOM
}

func (m *BaseResponse) GetBody() *any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

// 根据roomid查询  请求
type QueryUsersByRoomReq struct {
	RoomId               string   `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUsersByRoomReq) Reset()         { *m = QueryUsersByRoomReq{} }
func (m *QueryUsersByRoomReq) String() string { return proto.CompactTextString(m) }
func (*QueryUsersByRoomReq) ProtoMessage()    {}
func (*QueryUsersByRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_comapi_service_3a978815dda79e72, []int{2}
}
func (m *QueryUsersByRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUsersByRoomReq.Unmarshal(m, b)
}
func (m *QueryUsersByRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUsersByRoomReq.Marshal(b, m, deterministic)
}
func (dst *QueryUsersByRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUsersByRoomReq.Merge(dst, src)
}
func (m *QueryUsersByRoomReq) XXX_Size() int {
	return xxx_messageInfo_QueryUsersByRoomReq.Size(m)
}
func (m *QueryUsersByRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUsersByRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUsersByRoomReq proto.InternalMessageInfo

func (m *QueryUsersByRoomReq) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

// 根据roomid查询  响应
type QueryUserByRoomRsp struct {
	UserInfos            []*QueryUserByRoomRsp_UserInfo `protobuf:"bytes,1,rep,name=userInfos,proto3" json:"userInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *QueryUserByRoomRsp) Reset()         { *m = QueryUserByRoomRsp{} }
func (m *QueryUserByRoomRsp) String() string { return proto.CompactTextString(m) }
func (*QueryUserByRoomRsp) ProtoMessage()    {}
func (*QueryUserByRoomRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_comapi_service_3a978815dda79e72, []int{3}
}
func (m *QueryUserByRoomRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserByRoomRsp.Unmarshal(m, b)
}
func (m *QueryUserByRoomRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserByRoomRsp.Marshal(b, m, deterministic)
}
func (dst *QueryUserByRoomRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserByRoomRsp.Merge(dst, src)
}
func (m *QueryUserByRoomRsp) XXX_Size() int {
	return xxx_messageInfo_QueryUserByRoomRsp.Size(m)
}
func (m *QueryUserByRoomRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserByRoomRsp.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserByRoomRsp proto.InternalMessageInfo

func (m *QueryUserByRoomRsp) GetUserInfos() []*QueryUserByRoomRsp_UserInfo {
	if m != nil {
		return m.UserInfos
	}
	return nil
}

type QueryUserByRoomRsp_UserInfo struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryUserByRoomRsp_UserInfo) Reset()         { *m = QueryUserByRoomRsp_UserInfo{} }
func (m *QueryUserByRoomRsp_UserInfo) String() string { return proto.CompactTextString(m) }
func (*QueryUserByRoomRsp_UserInfo) ProtoMessage()    {}
func (*QueryUserByRoomRsp_UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_comapi_service_3a978815dda79e72, []int{3, 0}
}
func (m *QueryUserByRoomRsp_UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryUserByRoomRsp_UserInfo.Unmarshal(m, b)
}
func (m *QueryUserByRoomRsp_UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryUserByRoomRsp_UserInfo.Marshal(b, m, deterministic)
}
func (dst *QueryUserByRoomRsp_UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserByRoomRsp_UserInfo.Merge(dst, src)
}
func (m *QueryUserByRoomRsp_UserInfo) XXX_Size() int {
	return xxx_messageInfo_QueryUserByRoomRsp_UserInfo.Size(m)
}
func (m *QueryUserByRoomRsp_UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserByRoomRsp_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserByRoomRsp_UserInfo proto.InternalMessageInfo

func (m *QueryUserByRoomRsp_UserInfo) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseRequest)(nil), "BaseRequest")
	proto.RegisterType((*BaseResponse)(nil), "BaseResponse")
	proto.RegisterType((*QueryUsersByRoomReq)(nil), "QueryUsersByRoomReq")
	proto.RegisterType((*QueryUserByRoomRsp)(nil), "QueryUserByRoomRsp")
	proto.RegisterType((*QueryUserByRoomRsp_UserInfo)(nil), "QueryUserByRoomRsp.UserInfo")
	proto.RegisterEnum("CMD", CMD_name, CMD_value)
	proto.RegisterEnum("ReturnCode", ReturnCode_name, ReturnCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ComApiServerClient is the client API for ComApiServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComApiServerClient interface {
	BaseInterface(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type comApiServerClient struct {
	cc *grpc.ClientConn
}

func NewComApiServerClient(cc *grpc.ClientConn) ComApiServerClient {
	return &comApiServerClient{cc}
}

func (c *comApiServerClient) BaseInterface(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/ComApiServer/BaseInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComApiServerServer is the server API for ComApiServer service.
type ComApiServerServer interface {
	BaseInterface(context.Context, *BaseRequest) (*BaseResponse, error)
}

func RegisterComApiServerServer(s *grpc.Server, srv ComApiServerServer) {
	s.RegisterService(&_ComApiServer_serviceDesc, srv)
}

func _ComApiServer_BaseInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComApiServerServer).BaseInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ComApiServer/BaseInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComApiServerServer).BaseInterface(ctx, req.(*BaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ComApiServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ComApiServer",
	HandlerType: (*ComApiServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BaseInterface",
			Handler:    _ComApiServer_BaseInterface_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comapi-service.proto",
}

func init() {
	proto.RegisterFile("comapi-service.proto", fileDescriptor_comapi_service_3a978815dda79e72)
}

var fileDescriptor_comapi_service_3a978815dda79e72 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x6b, 0xdb, 0x30,
	0x14, 0xc5, 0xab, 0xc6, 0xed, 0x96, 0x9b, 0xb4, 0x78, 0x4a, 0xb7, 0xa5, 0xa1, 0x6c, 0x21, 0x4f,
	0xa6, 0x50, 0x15, 0xb2, 0xb7, 0x0d, 0x06, 0xf1, 0x9f, 0x07, 0xb3, 0x24, 0xa6, 0xf2, 0xcc, 0x48,
	0x5f, 0x4c, 0x62, 0x5f, 0x97, 0xc0, 0x62, 0xb9, 0x92, 0x5d, 0xf0, 0xe3, 0x3e, 0xc9, 0xbe, 0xea,
	0x88, 0xed, 0x24, 0x85, 0x15, 0xfa, 0xa6, 0x7b, 0xce, 0x41, 0xbf, 0x83, 0x74, 0xe1, 0x22, 0x12,
	0x9b, 0x65, 0xb6, 0xbe, 0x51, 0x28, 0x9f, 0xd6, 0x11, 0xb2, 0x4c, 0x8a, 0x5c, 0x0c, 0x2e, 0x1f,
	0x84, 0x78, 0xf8, 0x8d, 0xb7, 0xd5, 0xb4, 0x2a, 0x92, 0xdb, 0x65, 0x5a, 0xd6, 0xd6, 0x68, 0x0d,
	0x1d, 0x73, 0xa9, 0x90, 0xe3, 0x63, 0x81, 0x2a, 0xa7, 0x57, 0xd0, 0x96, 0xf5, 0xd1, 0x8d, 0xfb,
	0x64, 0x48, 0x8c, 0x36, 0x3f, 0x08, 0x94, 0x02, 0x89, 0xfa, 0xc7, 0x43, 0x62, 0x9c, 0x8f, 0x35,
	0x66, 0xcd, 0x6c, 0x4e, 0x22, 0x6a, 0x80, 0xb6, 0x12, 0x71, 0xd9, 0x6f, 0x0d, 0x89, 0xd1, 0x19,
	0x5f, 0xb0, 0x1a, 0xc5, 0x76, 0x28, 0x36, 0x49, 0x4b, 0x5e, 0x25, 0x46, 0x7f, 0x09, 0x74, 0x6b,
	0x96, 0xca, 0x44, 0xaa, 0xf0, 0x15, 0xd8, 0x67, 0xd0, 0x22, 0x11, 0x63, 0xc3, 0xeb, 0x30, 0x8e,
	0x79, 0x21, 0x53, 0x4b, 0xc4, 0xc8, 0x2b, 0x83, 0x52, 0xd0, 0x62, 0x54, 0x51, 0x45, 0x6e, 0xf3,
	0xea, 0x5c, 0x37, 0xd4, 0x5e, 0x6e, 0x78, 0xf2, 0x6a, 0xc3, 0x1b, 0xe8, 0xdd, 0x15, 0x28, 0xcb,
	0x40, 0xa1, 0x54, 0x66, 0xc9, 0x85, 0xd8, 0x70, 0x7c, 0xa4, 0x1f, 0xe0, 0x54, 0x0a, 0xb1, 0xd9,
	0x97, 0x6c, 0xa6, 0x51, 0x0e, 0x74, 0x1f, 0x6f, 0xd2, 0x2a, 0xa3, 0x5f, 0xa1, 0x5d, 0x28, 0x94,
	0x6e, 0x9a, 0x08, 0xd5, 0x27, 0xc3, 0x96, 0xd1, 0x19, 0x5f, 0xb1, 0xff, 0x73, 0x2c, 0x68, 0x42,
	0xfc, 0x10, 0x1f, 0x8c, 0xe0, 0xed, 0x4e, 0xde, 0x52, 0x2b, 0x63, 0x4f, 0xad, 0xa7, 0xeb, 0x4f,
	0xd0, 0xb2, 0x66, 0x36, 0xfd, 0x08, 0xbd, 0xbb, 0xc0, 0xe1, 0x8b, 0x30, 0xf0, 0x1d, 0xee, 0x87,
	0xe6, 0x22, 0xe4, 0x9e, 0x37, 0xd3, 0x8f, 0xae, 0xff, 0x10, 0x80, 0xc3, 0x5b, 0xd1, 0x0e, 0xbc,
	0xf1, 0x03, 0xcb, 0x72, 0x7c, 0x5f, 0x3f, 0xa2, 0xe7, 0x00, 0xc1, 0xfc, 0x87, 0xf7, 0x6b, 0x1e,
	0x5a, 0x33, 0x5b, 0x27, 0x54, 0x87, 0xae, 0xe9, 0xd9, 0x8b, 0xd0, 0xf5, 0xc3, 0x79, 0x30, 0x9d,
	0xea, 0xc7, 0xf4, 0x12, 0xde, 0xdb, 0x8e, 0xef, 0x70, 0x77, 0x32, 0x75, 0xef, 0x27, 0x3f, 0x5d,
	0x6f, 0x1e, 0x3a, 0x9c, 0x7b, 0x5c, 0x6f, 0x6d, 0x89, 0x2f, 0x19, 0xda, 0xf6, 0x96, 0xe6, 0xd6,
	0x5a, 0x39, 0x19, 0x7f, 0x87, 0xae, 0x25, 0x36, 0x93, 0x6c, 0xed, 0xa3, 0x7c, 0x42, 0x49, 0x19,
	0x9c, 0x6d, 0x7f, 0xde, 0x4d, 0x73, 0x94, 0xc9, 0x32, 0x42, 0xda, 0x65, 0xcf, 0xb6, 0x6e, 0x70,
	0xc6, 0x9e, 0xef, 0xc5, 0xe8, 0xc8, 0xec, 0xdd, 0xbf, 0x63, 0xdf, 0x76, 0x1f, 0x94, 0xa0, 0x0c,
	0x63, 0x4c, 0x56, 0xa7, 0x95, 0xf0, 0xe5, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xb5, 0x53,
	0x0f, 0xe4, 0x02, 0x00, 0x00,
}
