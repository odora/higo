// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

/*
Package go_micro_srv_user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	UserRequest
	UserResponse
*/
package go_micro_srv_user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CreateUser service

type CreateUserService interface {
	CreateUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
}

type createUserService struct {
	c    client.Client
	name string
}

func NewCreateUserService(name string, c client.Client) CreateUserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.user"
	}
	return &createUserService{
		c:    c,
		name: name,
	}
}

func (c *createUserService) CreateUser(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "CreateUser.CreateUser", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CreateUser service

type CreateUserHandler interface {
	CreateUser(context.Context, *UserRequest, *UserResponse) error
}

func RegisterCreateUserHandler(s server.Server, hdlr CreateUserHandler, opts ...server.HandlerOption) {
	type createUser interface {
		CreateUser(ctx context.Context, in *UserRequest, out *UserResponse) error
	}
	type CreateUser struct {
		createUser
	}
	h := &createUserHandler{hdlr}
	s.Handle(s.NewHandler(&CreateUser{h}, opts...))
}

type createUserHandler struct {
	CreateUserHandler
}

func (h *createUserHandler) CreateUser(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.CreateUserHandler.CreateUser(ctx, in, out)
}
