// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rate.proto

package ratepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Rate struct {
	UUID                 string               `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Rating               int32                `protobuf:"varint,2,opt,name=Rating,proto3" json:"Rating,omitempty"`
	UserUUID             string               `protobuf:"bytes,3,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
	CommentUUID          string               `protobuf:"bytes,4,opt,name=CommentUUID,proto3" json:"CommentUUID,omitempty"`
	ProductUUID          string               `protobuf:"bytes,5,opt,name=ProductUUID,proto3" json:"ProductUUID,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Rate) Reset()         { *m = Rate{} }
func (m *Rate) String() string { return proto.CompactTextString(m) }
func (*Rate) ProtoMessage()    {}
func (*Rate) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{0}
}

func (m *Rate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rate.Unmarshal(m, b)
}
func (m *Rate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rate.Marshal(b, m, deterministic)
}
func (m *Rate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rate.Merge(m, src)
}
func (m *Rate) XXX_Size() int {
	return xxx_messageInfo_Rate.Size(m)
}
func (m *Rate) XXX_DiscardUnknown() {
	xxx_messageInfo_Rate.DiscardUnknown(m)
}

var xxx_messageInfo_Rate proto.InternalMessageInfo

func (m *Rate) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Rate) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Rate) GetUserUUID() string {
	if m != nil {
		return m.UserUUID
	}
	return ""
}

func (m *Rate) GetCommentUUID() string {
	if m != nil {
		return m.CommentUUID
	}
	return ""
}

func (m *Rate) GetProductUUID() string {
	if m != nil {
		return m.ProductUUID
	}
	return ""
}

func (m *Rate) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Rate) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Rate) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type GetRequest struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{1}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type GetResponse struct {
	Rate                 *Rate    `protobuf:"bytes,1,opt,name=rate,proto3" json:"rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{2}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetRate() *Rate {
	if m != nil {
		return m.Rate
	}
	return nil
}

type SetRequest struct {
	Rate                 *Rate    `protobuf:"bytes,1,opt,name=rate,proto3" json:"rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{3}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetRate() *Rate {
	if m != nil {
		return m.Rate
	}
	return nil
}

type SetResponse struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{4}
}

func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (m *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(m, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type UpdateRequest struct {
	Rate                 *Rate    `protobuf:"bytes,1,opt,name=rate,proto3" json:"rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetRate() *Rate {
	if m != nil {
		return m.Rate
	}
	return nil
}

type DeleteRequest struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_426335fde4cae2d1, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func init() {
	proto.RegisterType((*Rate)(nil), "ratepb.Rate")
	proto.RegisterType((*GetRequest)(nil), "ratepb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "ratepb.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "ratepb.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "ratepb.SetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "ratepb.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "ratepb.DeleteRequest")
}

func init() { proto.RegisterFile("rate.proto", fileDescriptor_426335fde4cae2d1) }

var fileDescriptor_426335fde4cae2d1 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x52, 0xf2, 0x30,
	0x10, 0xc7, 0xbf, 0x42, 0x29, 0xb0, 0xfd, 0xb8, 0xc4, 0x91, 0x61, 0xea, 0xc1, 0x5a, 0x2f, 0x9c,
	0x82, 0xe2, 0x45, 0x8f, 0x8c, 0x38, 0x0c, 0x37, 0x26, 0xb5, 0x0f, 0x50, 0xe8, 0xca, 0x30, 0x43,
	0x49, 0x6d, 0xd3, 0x83, 0x6f, 0xe7, 0xbb, 0xf8, 0x22, 0x4e, 0x12, 0x4a, 0x5b, 0x45, 0xf1, 0x96,
	0xec, 0xfe, 0x7e, 0x69, 0xfe, 0xd9, 0x29, 0x40, 0x1a, 0x0a, 0xa4, 0x49, 0xca, 0x05, 0x27, 0x96,
	0x5c, 0x27, 0x4b, 0xe7, 0x62, 0xcd, 0xf9, 0x7a, 0x8b, 0x23, 0x55, 0x5d, 0xe6, 0x2f, 0x23, 0x8c,
	0x13, 0xf1, 0xa6, 0x21, 0xe7, 0xf2, 0x6b, 0x53, 0x6c, 0x62, 0xcc, 0x44, 0x18, 0x27, 0x1a, 0xf0,
	0xde, 0x1b, 0x60, 0xb2, 0x50, 0x20, 0x21, 0x60, 0x06, 0xc1, 0x7c, 0x3a, 0x30, 0x5c, 0x63, 0xd8,
	0x65, 0x6a, 0x4d, 0xfa, 0x60, 0xb1, 0x50, 0x6c, 0x76, 0xeb, 0x41, 0xc3, 0x35, 0x86, 0x2d, 0xb6,
	0xdf, 0x11, 0x07, 0x3a, 0x41, 0x86, 0xa9, 0xe2, 0x9b, 0x8a, 0x3f, 0xec, 0x89, 0x0b, 0xf6, 0x23,
	0x8f, 0x63, 0xdc, 0x09, 0xd5, 0x36, 0x55, 0xbb, 0x5a, 0x92, 0xc4, 0x22, 0xe5, 0x51, 0xbe, 0xd2,
	0x44, 0x4b, 0x13, 0x95, 0x12, 0xb9, 0x87, 0xee, 0x2a, 0xc5, 0x50, 0x60, 0x34, 0x11, 0x03, 0xcb,
	0x35, 0x86, 0xf6, 0xd8, 0xa1, 0x3a, 0x09, 0x2d, 0x92, 0xd0, 0xe7, 0x22, 0x09, 0x2b, 0x61, 0x69,
	0xe6, 0x49, 0xb4, 0x37, 0xdb, 0xa7, 0xcd, 0x03, 0x2c, 0xcd, 0x08, 0xb7, 0xa8, 0xcd, 0xce, 0x69,
	0xf3, 0x00, 0x7b, 0x2e, 0xc0, 0x0c, 0x05, 0xc3, 0xd7, 0x1c, 0x33, 0x71, 0xec, 0x1d, 0xbd, 0x11,
	0xd8, 0x8a, 0xc8, 0x12, 0xbe, 0xcb, 0x90, 0xb8, 0x60, 0xca, 0xd9, 0x29, 0xc4, 0x1e, 0xff, 0xa7,
	0x7a, 0x90, 0x54, 0x8e, 0x81, 0xa9, 0x8e, 0x47, 0x01, 0xfc, 0xf2, 0xc8, 0xd3, 0xfc, 0x15, 0xd8,
	0x7e, 0xe5, 0x03, 0xc7, 0xee, 0x70, 0x0b, 0xbd, 0x40, 0x85, 0xfd, 0xfb, 0xa9, 0xd7, 0xd0, 0x9b,
	0xaa, 0x94, 0xbf, 0x64, 0x1b, 0x7f, 0x18, 0xd0, 0x96, 0xce, 0x64, 0x31, 0x27, 0x37, 0xd0, 0x9c,
	0xa1, 0x20, 0xa4, 0x38, 0xab, 0x7c, 0x16, 0xe7, 0xac, 0x56, 0xd3, 0xf7, 0xf4, 0xfe, 0x49, 0xc3,
	0xaf, 0x1a, 0xfe, 0x11, 0xc3, 0xaf, 0x19, 0x0f, 0x60, 0xe9, 0x1c, 0xe4, 0xbc, 0x00, 0x6a, 0xb9,
	0x9c, 0xfe, 0xb7, 0xa9, 0x3d, 0xc9, 0x1f, 0x42, 0xab, 0x3a, 0x4f, 0xa9, 0xd6, 0xf2, 0xfd, 0xac,
	0x2e, 0x2d, 0x55, 0xb9, 0xfb, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xaf, 0x8d, 0xf1, 0x81, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateAPIClient is the client API for RateAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateAPIClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type rateAPIClient struct {
	cc *grpc.ClientConn
}

func NewRateAPIClient(cc *grpc.ClientConn) RateAPIClient {
	return &rateAPIClient{cc}
}

func (c *rateAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/ratepb.RateAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateAPIClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/ratepb.RateAPI/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ratepb.RateAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rateAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ratepb.RateAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateAPIServer is the server API for RateAPI service.
type RateAPIServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Update(context.Context, *UpdateRequest) (*empty.Empty, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
}

// UnimplementedRateAPIServer can be embedded to have forward compatible implementations.
type UnimplementedRateAPIServer struct {
}

func (*UnimplementedRateAPIServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRateAPIServer) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedRateAPIServer) Update(ctx context.Context, req *UpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedRateAPIServer) Delete(ctx context.Context, req *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterRateAPIServer(s *grpc.Server, srv RateAPIServer) {
	s.RegisterService(&_RateAPI_serviceDesc, srv)
}

func _RateAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratepb.RateAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateAPI_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateAPIServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratepb.RateAPI/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateAPIServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratepb.RateAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RateAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ratepb.RateAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ratepb.RateAPI",
	HandlerType: (*RateAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RateAPI_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _RateAPI_Set_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RateAPI_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RateAPI_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rate.proto",
}
