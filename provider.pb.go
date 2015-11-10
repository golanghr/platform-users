// Code generated by protoc-gen-go.
// source: provider.proto
// DO NOT EDIT!

package users

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Provider_Providers int32

const (
	Provider_STANDARD Provider_Providers = 0
	Provider_FACEBOOK Provider_Providers = 1
	Provider_TWITTER  Provider_Providers = 2
	Provider_GOOGLE   Provider_Providers = 3
	Provider_LINKEDIN Provider_Providers = 4
)

var Provider_Providers_name = map[int32]string{
	0: "STANDARD",
	1: "FACEBOOK",
	2: "TWITTER",
	3: "GOOGLE",
	4: "LINKEDIN",
}
var Provider_Providers_value = map[string]int32{
	"STANDARD": 0,
	"FACEBOOK": 1,
	"TWITTER":  2,
	"GOOGLE":   3,
	"LINKEDIN": 4,
}

func (x Provider_Providers) String() string {
	return proto.EnumName(Provider_Providers_name, int32(x))
}

type Provider struct {
	Id         int64              `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	User       *User              `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Provider   Provider_Providers `protobuf:"varint,3,opt,name=provider,enum=users.Provider_Providers" json:"provider,omitempty"`
	Username   string             `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
	Password   string             `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	ProviderId string             `protobuf:"bytes,6,opt,name=provider_id" json:"provider_id,omitempty"`
	Deleted    bool               `protobuf:"varint,7,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *Provider) Reset()         { *m = Provider{} }
func (m *Provider) String() string { return proto.CompactTextString(m) }
func (*Provider) ProtoMessage()    {}

func (m *Provider) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*Provider)(nil), "users.Provider")
	proto.RegisterEnum("users.Provider_Providers", Provider_Providers_name, Provider_Providers_value)
}
