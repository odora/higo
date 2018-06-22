// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unique_id.proto

package go_micro_srv_unique_id

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

type UniqueIdRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UniqueIdRequest) Reset()         { *m = UniqueIdRequest{} }
func (m *UniqueIdRequest) String() string { return proto.CompactTextString(m) }
func (*UniqueIdRequest) ProtoMessage()    {}
func (*UniqueIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_unique_id_89c71fdfa1d85a6e, []int{0}
}
func (m *UniqueIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UniqueIdRequest.Unmarshal(m, b)
}
func (m *UniqueIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UniqueIdRequest.Marshal(b, m, deterministic)
}
func (dst *UniqueIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniqueIdRequest.Merge(dst, src)
}
func (m *UniqueIdRequest) XXX_Size() int {
	return xxx_messageInfo_UniqueIdRequest.Size(m)
}
func (m *UniqueIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UniqueIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UniqueIdRequest proto.InternalMessageInfo

func (m *UniqueIdRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type UniqueIdResponse struct {
	Value                int64    `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UniqueIdResponse) Reset()         { *m = UniqueIdResponse{} }
func (m *UniqueIdResponse) String() string { return proto.CompactTextString(m) }
func (*UniqueIdResponse) ProtoMessage()    {}
func (*UniqueIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_unique_id_89c71fdfa1d85a6e, []int{1}
}
func (m *UniqueIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UniqueIdResponse.Unmarshal(m, b)
}
func (m *UniqueIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UniqueIdResponse.Marshal(b, m, deterministic)
}
func (dst *UniqueIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniqueIdResponse.Merge(dst, src)
}
func (m *UniqueIdResponse) XXX_Size() int {
	return xxx_messageInfo_UniqueIdResponse.Size(m)
}
func (m *UniqueIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UniqueIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UniqueIdResponse proto.InternalMessageInfo

func (m *UniqueIdResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*UniqueIdRequest)(nil), "go.micro.srv.unique_id.UniqueIdRequest")
	proto.RegisterType((*UniqueIdResponse)(nil), "go.micro.srv.unique_id.UniqueIdResponse")
}

func init() { proto.RegisterFile("unique_id.proto", fileDescriptor_unique_id_89c71fdfa1d85a6e) }

var fileDescriptor_unique_id_89c71fdfa1d85a6e = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xcd, 0xcb, 0x2c,
	0x2c, 0x4d, 0x8d, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4b, 0xcf, 0xd7,
	0xcb, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x83, 0xcb, 0x2a, 0x29, 0x73, 0xf1,
	0x87, 0x82, 0x39, 0x9e, 0x29, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x02, 0x5c, 0xcc,
	0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x92, 0x06, 0x97, 0x00,
	0x42, 0x51, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69,
	0x2a, 0x58, 0x1d, 0x73, 0x10, 0x84, 0x63, 0x54, 0xcd, 0x25, 0xe0, 0x9e, 0x9a, 0x97, 0x5a, 0x94,
	0x58, 0x92, 0x0a, 0xd3, 0x21, 0x94, 0x8e, 0x45, 0x4c, 0x5d, 0x0f, 0xbb, 0x7b, 0xf4, 0xd0, 0x1c,
	0x23, 0xa5, 0x41, 0x58, 0x21, 0xc4, 0x41, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xaf, 0x1a, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa5, 0xae, 0x19, 0x11, 0xfd, 0x00, 0x00, 0x00,
}
