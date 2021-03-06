// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

package users

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import platform "github.com/golanghr/platform/protos"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User_Gender int32

const (
	User_UNKNOWN User_Gender = 0
	User_FEMALE  User_Gender = 1
	User_MALE    User_Gender = 2
)

var User_Gender_name = map[int32]string{
	0: "UNKNOWN",
	1: "FEMALE",
	2: "MALE",
}
var User_Gender_value = map[string]int32{
	"UNKNOWN": 0,
	"FEMALE":  1,
	"MALE":    2,
}

func (x User_Gender) String() string {
	return proto.EnumName(User_Gender_name, int32(x))
}

type User_Relationship int32

const (
	User_IRRELEVANT      User_Relationship = 0
	User_SINGLE          User_Relationship = 1
	User_IN_RELATIONSHIP User_Relationship = 2
	User_DIVORCED        User_Relationship = 3
	User_MARRIED         User_Relationship = 4
	User_COMPLICATED     User_Relationship = 5
)

var User_Relationship_name = map[int32]string{
	0: "IRRELEVANT",
	1: "SINGLE",
	2: "IN_RELATIONSHIP",
	3: "DIVORCED",
	4: "MARRIED",
	5: "COMPLICATED",
}
var User_Relationship_value = map[string]int32{
	"IRRELEVANT":      0,
	"SINGLE":          1,
	"IN_RELATIONSHIP": 2,
	"DIVORCED":        3,
	"MARRIED":         4,
	"COMPLICATED":     5,
}

func (x User_Relationship) String() string {
	return proto.EnumName(User_Relationship_name, int32(x))
}

type User struct {
	Id           int64             `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Email        string            `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	FirstName    string            `protobuf:"bytes,3,opt,name=first_name" json:"first_name,omitempty"`
	LastName     string            `protobuf:"bytes,4,opt,name=last_name" json:"last_name,omitempty"`
	Dob          string            `protobuf:"bytes,5,opt,name=dob" json:"dob,omitempty"`
	Geo          *User_Geo         `protobuf:"bytes,6,opt,name=geo" json:"geo,omitempty"`
	Gender       User_Gender       `protobuf:"varint,7,opt,name=gender,enum=users.User_Gender" json:"gender,omitempty"`
	Relationship User_Relationship `protobuf:"varint,8,opt,name=relationship,enum=users.User_Relationship" json:"relationship,omitempty"`
	PhoneNumber  string            `protobuf:"bytes,9,opt,name=phone_number" json:"phone_number,omitempty"`
	Suspended    bool              `protobuf:"varint,10,opt,name=suspended" json:"suspended,omitempty"`
	Deleted      bool              `protobuf:"varint,11,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}

func (m *User) GetGeo() *User_Geo {
	if m != nil {
		return m.Geo
	}
	return nil
}

type User_Geo struct {
	Latitude  int64 `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude int64 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *User_Geo) Reset()         { *m = User_Geo{} }
func (m *User_Geo) String() string { return proto.CompactTextString(m) }
func (*User_Geo) ProtoMessage()    {}

func init() {
	proto.RegisterType((*User)(nil), "users.User")
	proto.RegisterType((*User_Geo)(nil), "users.User.Geo")
	proto.RegisterEnum("users.User_Gender", User_Gender_name, User_Gender_value)
	proto.RegisterEnum("users.User_Relationship", User_Relationship_name, User_Relationship_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Users service

type UsersClient interface {
	SignIn(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*platform.Response, error)
	SignOut(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*platform.Response, error)
	Register(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*platform.Response, error)
	Recover(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*platform.Response, error)
	Profile(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*platform.Response, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) SignIn(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*platform.Response, error) {
	out := new(platform.Response)
	err := grpc.Invoke(ctx, "/users.Users/SignIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) SignOut(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*platform.Response, error) {
	out := new(platform.Response)
	err := grpc.Invoke(ctx, "/users.Users/SignOut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Register(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*platform.Response, error) {
	out := new(platform.Response)
	err := grpc.Invoke(ctx, "/users.Users/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Recover(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*platform.Response, error) {
	out := new(platform.Response)
	err := grpc.Invoke(ctx, "/users.Users/Recover", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Profile(ctx context.Context, in *platform.Request, opts ...grpc.CallOption) (*platform.Response, error) {
	out := new(platform.Response)
	err := grpc.Invoke(ctx, "/users.Users/Profile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersServer interface {
	SignIn(context.Context, *platform.Request) (*platform.Response, error)
	SignOut(context.Context, *platform.Request) (*platform.Response, error)
	Register(context.Context, *platform.Request) (*platform.Response, error)
	Recover(context.Context, *platform.Request) (*platform.Response, error)
	Profile(context.Context, *platform.Request) (*platform.Response, error)
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(platform.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(UsersServer).SignIn(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Users_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(platform.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(UsersServer).SignOut(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Users_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(platform.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(UsersServer).Register(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Users_Recover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(platform.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(UsersServer).Recover(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Users_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(platform.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(UsersServer).Profile(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _Users_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _Users_SignOut_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Users_Register_Handler,
		},
		{
			MethodName: "Recover",
			Handler:    _Users_Recover_Handler,
		},
		{
			MethodName: "Profile",
			Handler:    _Users_Profile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
