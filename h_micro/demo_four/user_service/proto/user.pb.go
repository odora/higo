// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package go_micro_srv_user

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

type RegisterRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone" json:"phone,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_16d80bb2443c8295, []int{0}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_16d80bb2443c8295, []int{1}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(dst, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "go.micro.srv.user.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "go.micro.srv.user.RegisterResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_user_16d80bb2443c8295) }

var fileDescriptor_user_16d80bb2443c8295 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x4f, 0x8f, 0x82, 0x40,
	0x0c, 0xc5, 0x17, 0x36, 0xbb, 0x61, 0xbb, 0x89, 0x7f, 0x26, 0x1e, 0x08, 0x27, 0x33, 0x5e, 0x3c,
	0xcd, 0x41, 0xbf, 0x09, 0x86, 0x83, 0x47, 0x64, 0x1a, 0x9c, 0x44, 0xe8, 0xd8, 0x02, 0x7e, 0x7d,
	0x03, 0xf8, 0x27, 0xd1, 0xc4, 0xe3, 0xef, 0xbd, 0xf6, 0xf5, 0x15, 0xa0, 0x15, 0x64, 0xe3, 0x99,
	0x1a, 0x52, 0xf3, 0x92, 0x4c, 0xe5, 0x0a, 0x26, 0x23, 0xdc, 0x99, 0xde, 0xd0, 0x7b, 0x98, 0xa6,
	0x58, 0x3a, 0x69, 0x90, 0x53, 0x3c, 0xb7, 0x28, 0x8d, 0x5a, 0xc0, 0x0f, 0x56, 0xb9, 0x3b, 0xc5,
	0xc1, 0x32, 0x58, 0xff, 0xa5, 0x23, 0xf4, 0xaa, 0x3f, 0x52, 0x8d, 0x71, 0x38, 0xaa, 0x03, 0xa8,
	0x04, 0x22, 0x9f, 0x8b, 0x5c, 0x88, 0x6d, 0xfc, 0x3d, 0x18, 0x0f, 0xd6, 0x1a, 0x66, 0xcf, 0x68,
	0xf1, 0x54, 0x0b, 0xaa, 0x09, 0x84, 0xce, 0xde, 0x82, 0x43, 0x67, 0x37, 0x16, 0xfe, 0x33, 0x41,
	0xde, 0x21, 0x77, 0xae, 0x40, 0x95, 0x41, 0x74, 0x5f, 0x51, 0xda, 0xbc, 0xb5, 0x35, 0x2f, 0x55,
	0x93, 0xd5, 0xc7, 0x99, 0xf1, 0xa6, 0xfe, 0x3a, 0xfc, 0x0e, 0xef, 0x6f, 0xaf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xd2, 0x86, 0xa1, 0x43, 0x0c, 0x01, 0x00, 0x00,
}