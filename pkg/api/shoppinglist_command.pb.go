// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shoppinglist_command/v1/shoppinglist_command.proto

package api

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type NewShoppingListRequest struct {
	NewShoppingList      *ShoppingList `protobuf:"bytes,1,opt,name=new_shopping_list,json=newShoppingList,proto3" json:"new_shopping_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NewShoppingListRequest) Reset()         { *m = NewShoppingListRequest{} }
func (m *NewShoppingListRequest) String() string { return proto.CompactTextString(m) }
func (*NewShoppingListRequest) ProtoMessage()    {}
func (*NewShoppingListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d35231c372be15, []int{0}
}

func (m *NewShoppingListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewShoppingListRequest.Unmarshal(m, b)
}
func (m *NewShoppingListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewShoppingListRequest.Marshal(b, m, deterministic)
}
func (m *NewShoppingListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewShoppingListRequest.Merge(m, src)
}
func (m *NewShoppingListRequest) XXX_Size() int {
	return xxx_messageInfo_NewShoppingListRequest.Size(m)
}
func (m *NewShoppingListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewShoppingListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewShoppingListRequest proto.InternalMessageInfo

func (m *NewShoppingListRequest) GetNewShoppingList() *ShoppingList {
	if m != nil {
		return m.NewShoppingList
	}
	return nil
}

type SetActiveShoppingListRequest struct {
	ActiveShoppingListName string   `protobuf:"bytes,1,opt,name=active_shoppingList_name,json=activeShoppingListName,proto3" json:"active_shoppingList_name,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *SetActiveShoppingListRequest) Reset()         { *m = SetActiveShoppingListRequest{} }
func (m *SetActiveShoppingListRequest) String() string { return proto.CompactTextString(m) }
func (*SetActiveShoppingListRequest) ProtoMessage()    {}
func (*SetActiveShoppingListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d35231c372be15, []int{1}
}

func (m *SetActiveShoppingListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetActiveShoppingListRequest.Unmarshal(m, b)
}
func (m *SetActiveShoppingListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetActiveShoppingListRequest.Marshal(b, m, deterministic)
}
func (m *SetActiveShoppingListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetActiveShoppingListRequest.Merge(m, src)
}
func (m *SetActiveShoppingListRequest) XXX_Size() int {
	return xxx_messageInfo_SetActiveShoppingListRequest.Size(m)
}
func (m *SetActiveShoppingListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetActiveShoppingListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetActiveShoppingListRequest proto.InternalMessageInfo

func (m *SetActiveShoppingListRequest) GetActiveShoppingListName() string {
	if m != nil {
		return m.ActiveShoppingListName
	}
	return ""
}

type FinishShoppingRequest struct {
	UpdatedShoppingList  *ShoppingList `protobuf:"bytes,1,opt,name=updated_shopping_list,json=updatedShoppingList,proto3" json:"updated_shopping_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FinishShoppingRequest) Reset()         { *m = FinishShoppingRequest{} }
func (m *FinishShoppingRequest) String() string { return proto.CompactTextString(m) }
func (*FinishShoppingRequest) ProtoMessage()    {}
func (*FinishShoppingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d35231c372be15, []int{2}
}

func (m *FinishShoppingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinishShoppingRequest.Unmarshal(m, b)
}
func (m *FinishShoppingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinishShoppingRequest.Marshal(b, m, deterministic)
}
func (m *FinishShoppingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinishShoppingRequest.Merge(m, src)
}
func (m *FinishShoppingRequest) XXX_Size() int {
	return xxx_messageInfo_FinishShoppingRequest.Size(m)
}
func (m *FinishShoppingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FinishShoppingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FinishShoppingRequest proto.InternalMessageInfo

func (m *FinishShoppingRequest) GetUpdatedShoppingList() *ShoppingList {
	if m != nil {
		return m.UpdatedShoppingList
	}
	return nil
}

type FinishShoppingResponse struct {
	ShoppingList         *ShoppingList `protobuf:"bytes,1,opt,name=shopping_list,json=shoppingList,proto3" json:"shopping_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FinishShoppingResponse) Reset()         { *m = FinishShoppingResponse{} }
func (m *FinishShoppingResponse) String() string { return proto.CompactTextString(m) }
func (*FinishShoppingResponse) ProtoMessage()    {}
func (*FinishShoppingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d35231c372be15, []int{3}
}

func (m *FinishShoppingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinishShoppingResponse.Unmarshal(m, b)
}
func (m *FinishShoppingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinishShoppingResponse.Marshal(b, m, deterministic)
}
func (m *FinishShoppingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinishShoppingResponse.Merge(m, src)
}
func (m *FinishShoppingResponse) XXX_Size() int {
	return xxx_messageInfo_FinishShoppingResponse.Size(m)
}
func (m *FinishShoppingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FinishShoppingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FinishShoppingResponse proto.InternalMessageInfo

func (m *FinishShoppingResponse) GetShoppingList() *ShoppingList {
	if m != nil {
		return m.ShoppingList
	}
	return nil
}

type AddNewItemRequest struct {
	ItemToAdd            *Item    `protobuf:"bytes,1,opt,name=item_to_add,json=itemToAdd,proto3" json:"item_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddNewItemRequest) Reset()         { *m = AddNewItemRequest{} }
func (m *AddNewItemRequest) String() string { return proto.CompactTextString(m) }
func (*AddNewItemRequest) ProtoMessage()    {}
func (*AddNewItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d35231c372be15, []int{4}
}

func (m *AddNewItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddNewItemRequest.Unmarshal(m, b)
}
func (m *AddNewItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddNewItemRequest.Marshal(b, m, deterministic)
}
func (m *AddNewItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNewItemRequest.Merge(m, src)
}
func (m *AddNewItemRequest) XXX_Size() int {
	return xxx_messageInfo_AddNewItemRequest.Size(m)
}
func (m *AddNewItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNewItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddNewItemRequest proto.InternalMessageInfo

func (m *AddNewItemRequest) GetItemToAdd() *Item {
	if m != nil {
		return m.ItemToAdd
	}
	return nil
}

type AddItemToShoppingListRequest struct {
	ItemToAdd            *Item         `protobuf:"bytes,1,opt,name=item_to_add,json=itemToAdd,proto3" json:"item_to_add,omitempty"`
	ShoppingList         *ShoppingList `protobuf:"bytes,2,opt,name=shopping_list,json=shoppingList,proto3" json:"shopping_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AddItemToShoppingListRequest) Reset()         { *m = AddItemToShoppingListRequest{} }
func (m *AddItemToShoppingListRequest) String() string { return proto.CompactTextString(m) }
func (*AddItemToShoppingListRequest) ProtoMessage()    {}
func (*AddItemToShoppingListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27d35231c372be15, []int{5}
}

func (m *AddItemToShoppingListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddItemToShoppingListRequest.Unmarshal(m, b)
}
func (m *AddItemToShoppingListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddItemToShoppingListRequest.Marshal(b, m, deterministic)
}
func (m *AddItemToShoppingListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddItemToShoppingListRequest.Merge(m, src)
}
func (m *AddItemToShoppingListRequest) XXX_Size() int {
	return xxx_messageInfo_AddItemToShoppingListRequest.Size(m)
}
func (m *AddItemToShoppingListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddItemToShoppingListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddItemToShoppingListRequest proto.InternalMessageInfo

func (m *AddItemToShoppingListRequest) GetItemToAdd() *Item {
	if m != nil {
		return m.ItemToAdd
	}
	return nil
}

func (m *AddItemToShoppingListRequest) GetShoppingList() *ShoppingList {
	if m != nil {
		return m.ShoppingList
	}
	return nil
}

func init() {
	proto.RegisterType((*NewShoppingListRequest)(nil), "buyarella.shoppinglist.command.NewShoppingListRequest")
	proto.RegisterType((*SetActiveShoppingListRequest)(nil), "buyarella.shoppinglist.command.SetActiveShoppingListRequest")
	proto.RegisterType((*FinishShoppingRequest)(nil), "buyarella.shoppinglist.command.FinishShoppingRequest")
	proto.RegisterType((*FinishShoppingResponse)(nil), "buyarella.shoppinglist.command.FinishShoppingResponse")
	proto.RegisterType((*AddNewItemRequest)(nil), "buyarella.shoppinglist.command.AddNewItemRequest")
	proto.RegisterType((*AddItemToShoppingListRequest)(nil), "buyarella.shoppinglist.command.AddItemToShoppingListRequest")
}

func init() {
	proto.RegisterFile("shoppinglist_command/v1/shoppinglist_command.proto", fileDescriptor_27d35231c372be15)
}

var fileDescriptor_27d35231c372be15 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x8a, 0xda, 0x40,
	0x14, 0xc6, 0x4d, 0x29, 0x2d, 0x1e, 0xdb, 0x8a, 0xd3, 0x9a, 0x4a, 0x90, 0x52, 0x72, 0xd5, 0xab,
	0x04, 0x2d, 0x95, 0x5e, 0x78, 0x93, 0x5a, 0x0a, 0xd2, 0xe2, 0x85, 0x96, 0x65, 0x59, 0x58, 0xc2,
	0xe8, 0x0c, 0x3a, 0x6c, 0x92, 0x99, 0x75, 0x46, 0x65, 0x2f, 0xf7, 0x01, 0xf6, 0x7d, 0xf6, 0xf1,
	0x96, 0xfc, 0x5b, 0x13, 0xcd, 0x26, 0xe4, 0x2e, 0xcc, 0x39, 0xdf, 0xf7, 0x9d, 0x73, 0xf8, 0x11,
	0x18, 0xca, 0x0d, 0x17, 0x82, 0x05, 0x6b, 0x8f, 0x49, 0xe5, 0xae, 0xb8, 0xef, 0xe3, 0x80, 0xd8,
	0xfb, 0x81, 0x5d, 0xf4, 0x6e, 0x89, 0x2d, 0x57, 0x1c, 0x7d, 0x59, 0xee, 0xee, 0xf0, 0x96, 0x7a,
	0x1e, 0xb6, 0xb2, 0x5d, 0x56, 0xd2, 0x65, 0xe8, 0x3e, 0x27, 0xd4, 0x93, 0xa1, 0x4b, 0xfc, 0x15,
	0xeb, 0xcc, 0x6b, 0xd0, 0x67, 0xf4, 0xb0, 0x48, 0x24, 0xff, 0x98, 0x54, 0x73, 0x7a, 0xbb, 0xa3,
	0x52, 0xa1, 0x09, 0x74, 0x02, 0x7a, 0x70, 0x53, 0x37, 0x37, 0xb4, 0xeb, 0x69, 0x5f, 0xb5, 0x6f,
	0xad, 0xe1, 0x67, 0xeb, 0x98, 0x96, 0x93, 0xb6, 0x83, 0xbc, 0x97, 0x79, 0x09, 0xfd, 0x05, 0x55,
	0xce, 0x4a, 0xb1, 0x3d, 0x2d, 0x0a, 0xf9, 0x09, 0x3d, 0x1c, 0x15, 0x9f, 0x73, 0xc2, 0xaa, 0x1b,
	0x60, 0x9f, 0x46, 0x59, 0xcd, 0xb9, 0x8e, 0xcf, 0xc4, 0x33, 0xec, 0x53, 0x93, 0x40, 0xf7, 0x0f,
	0x0b, 0x98, 0xdc, 0xa4, 0x95, 0xd4, 0xf2, 0x2f, 0x74, 0x77, 0x82, 0x60, 0x45, 0x49, 0xbd, 0xd9,
	0x3f, 0x26, 0xaa, 0xdc, 0xfc, 0x17, 0xa0, 0x9f, 0xa6, 0x48, 0xc1, 0x03, 0x49, 0xd1, 0x18, 0xde,
	0xd7, 0xb2, 0x7f, 0x97, 0x5d, 0xd0, 0xfc, 0x0d, 0x1d, 0x87, 0x90, 0x19, 0x3d, 0x4c, 0x15, 0xf5,
	0xd3, 0xc9, 0x6d, 0x68, 0x31, 0x45, 0x7d, 0x57, 0x71, 0x17, 0x13, 0x92, 0x18, 0xb6, 0x33, 0x86,
	0x51, 0x73, 0x33, 0xec, 0xf9, 0xcf, 0x1d, 0x42, 0xcc, 0x07, 0x0d, 0xfa, 0x0e, 0x21, 0xd3, 0xe8,
	0xa1, 0xe8, 0xbc, 0x75, 0x1d, 0xcf, 0xb7, 0x7a, 0x55, 0x63, 0xab, 0xe1, 0xe3, 0x6b, 0xf8, 0x94,
	0x2d, 0x4f, 0x62, 0xf8, 0x24, 0x22, 0xd0, 0x3e, 0xa1, 0x0c, 0x8d, 0xac, 0x72, 0x62, 0xad, 0x62,
	0x2c, 0x8d, 0x97, 0x46, 0x31, 0x1b, 0x48, 0x40, 0xb7, 0x10, 0x36, 0x34, 0xae, 0xca, 0x2a, 0x63,
	0xb4, 0x2c, 0xf1, 0x5e, 0x83, 0x0f, 0x79, 0x3e, 0xd0, 0x8f, 0xaa, 0xac, 0x42, 0x6a, 0x8d, 0x51,
	0x5d, 0x59, 0x8c, 0xa1, 0xd9, 0x40, 0x73, 0x80, 0x23, 0x4a, 0x68, 0x50, 0xe5, 0x73, 0x86, 0x9d,
	0x71, 0xca, 0x43, 0x7c, 0xc9, 0x42, 0xae, 0xaa, 0x2f, 0x59, 0x86, 0x63, 0xc9, 0x25, 0x7f, 0x35,
	0xaf, 0xde, 0x8a, 0x9b, 0xb5, 0x8d, 0x05, 0x5b, 0xbe, 0x89, 0xfe, 0x4c, 0xdf, 0x9f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xa7, 0x30, 0xa2, 0x07, 0x07, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShoppingListCommandsClient is the client API for ShoppingListCommands service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShoppingListCommandsClient interface {
	NewShoppingList(ctx context.Context, in *NewShoppingListRequest, opts ...grpc.CallOption) (*ShoppingList, error)
	SetActiveShoppingList(ctx context.Context, in *SetActiveShoppingListRequest, opts ...grpc.CallOption) (*ShoppingList, error)
	FinishShopping(ctx context.Context, in *FinishShoppingRequest, opts ...grpc.CallOption) (*FinishShoppingResponse, error)
	AddNewItem(ctx context.Context, in *AddNewItemRequest, opts ...grpc.CallOption) (*Item, error)
	AddItemToShoppingList(ctx context.Context, in *AddItemToShoppingListRequest, opts ...grpc.CallOption) (*ShoppingList, error)
}

type shoppingListCommandsClient struct {
	cc *grpc.ClientConn
}

func NewShoppingListCommandsClient(cc *grpc.ClientConn) ShoppingListCommandsClient {
	return &shoppingListCommandsClient{cc}
}

func (c *shoppingListCommandsClient) NewShoppingList(ctx context.Context, in *NewShoppingListRequest, opts ...grpc.CallOption) (*ShoppingList, error) {
	out := new(ShoppingList)
	err := c.cc.Invoke(ctx, "/buyarella.shoppinglist.command.ShoppingListCommands/NewShoppingList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingListCommandsClient) SetActiveShoppingList(ctx context.Context, in *SetActiveShoppingListRequest, opts ...grpc.CallOption) (*ShoppingList, error) {
	out := new(ShoppingList)
	err := c.cc.Invoke(ctx, "/buyarella.shoppinglist.command.ShoppingListCommands/SetActiveShoppingList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingListCommandsClient) FinishShopping(ctx context.Context, in *FinishShoppingRequest, opts ...grpc.CallOption) (*FinishShoppingResponse, error) {
	out := new(FinishShoppingResponse)
	err := c.cc.Invoke(ctx, "/buyarella.shoppinglist.command.ShoppingListCommands/FinishShopping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingListCommandsClient) AddNewItem(ctx context.Context, in *AddNewItemRequest, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/buyarella.shoppinglist.command.ShoppingListCommands/AddNewItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shoppingListCommandsClient) AddItemToShoppingList(ctx context.Context, in *AddItemToShoppingListRequest, opts ...grpc.CallOption) (*ShoppingList, error) {
	out := new(ShoppingList)
	err := c.cc.Invoke(ctx, "/buyarella.shoppinglist.command.ShoppingListCommands/AddItemToShoppingList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingListCommandsServer is the server API for ShoppingListCommands service.
type ShoppingListCommandsServer interface {
	NewShoppingList(context.Context, *NewShoppingListRequest) (*ShoppingList, error)
	SetActiveShoppingList(context.Context, *SetActiveShoppingListRequest) (*ShoppingList, error)
	FinishShopping(context.Context, *FinishShoppingRequest) (*FinishShoppingResponse, error)
	AddNewItem(context.Context, *AddNewItemRequest) (*Item, error)
	AddItemToShoppingList(context.Context, *AddItemToShoppingListRequest) (*ShoppingList, error)
}

// UnimplementedShoppingListCommandsServer can be embedded to have forward compatible implementations.
type UnimplementedShoppingListCommandsServer struct {
}

func (*UnimplementedShoppingListCommandsServer) NewShoppingList(ctx context.Context, req *NewShoppingListRequest) (*ShoppingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewShoppingList not implemented")
}
func (*UnimplementedShoppingListCommandsServer) SetActiveShoppingList(ctx context.Context, req *SetActiveShoppingListRequest) (*ShoppingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetActiveShoppingList not implemented")
}
func (*UnimplementedShoppingListCommandsServer) FinishShopping(ctx context.Context, req *FinishShoppingRequest) (*FinishShoppingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishShopping not implemented")
}
func (*UnimplementedShoppingListCommandsServer) AddNewItem(ctx context.Context, req *AddNewItemRequest) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewItem not implemented")
}
func (*UnimplementedShoppingListCommandsServer) AddItemToShoppingList(ctx context.Context, req *AddItemToShoppingListRequest) (*ShoppingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItemToShoppingList not implemented")
}

func RegisterShoppingListCommandsServer(s *grpc.Server, srv ShoppingListCommandsServer) {
	s.RegisterService(&_ShoppingListCommands_serviceDesc, srv)
}

func _ShoppingListCommands_NewShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListCommandsServer).NewShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buyarella.shoppinglist.command.ShoppingListCommands/NewShoppingList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListCommandsServer).NewShoppingList(ctx, req.(*NewShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingListCommands_SetActiveShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetActiveShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListCommandsServer).SetActiveShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buyarella.shoppinglist.command.ShoppingListCommands/SetActiveShoppingList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListCommandsServer).SetActiveShoppingList(ctx, req.(*SetActiveShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingListCommands_FinishShopping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishShoppingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListCommandsServer).FinishShopping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buyarella.shoppinglist.command.ShoppingListCommands/FinishShopping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListCommandsServer).FinishShopping(ctx, req.(*FinishShoppingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingListCommands_AddNewItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListCommandsServer).AddNewItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buyarella.shoppinglist.command.ShoppingListCommands/AddNewItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListCommandsServer).AddNewItem(ctx, req.(*AddNewItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShoppingListCommands_AddItemToShoppingList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemToShoppingListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingListCommandsServer).AddItemToShoppingList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buyarella.shoppinglist.command.ShoppingListCommands/AddItemToShoppingList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingListCommandsServer).AddItemToShoppingList(ctx, req.(*AddItemToShoppingListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShoppingListCommands_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buyarella.shoppinglist.command.ShoppingListCommands",
	HandlerType: (*ShoppingListCommandsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewShoppingList",
			Handler:    _ShoppingListCommands_NewShoppingList_Handler,
		},
		{
			MethodName: "SetActiveShoppingList",
			Handler:    _ShoppingListCommands_SetActiveShoppingList_Handler,
		},
		{
			MethodName: "FinishShopping",
			Handler:    _ShoppingListCommands_FinishShopping_Handler,
		},
		{
			MethodName: "AddNewItem",
			Handler:    _ShoppingListCommands_AddNewItem_Handler,
		},
		{
			MethodName: "AddItemToShoppingList",
			Handler:    _ShoppingListCommands_AddItemToShoppingList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shoppinglist_command/v1/shoppinglist_command.proto",
}