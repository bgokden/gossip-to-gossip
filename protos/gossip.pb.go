// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gossip.proto

package protos

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

type RegistrationRequest struct {
	Token                string          `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Mode                 uint32          `protobuf:"varint,2,opt,name=mode" json:"mode,omitempty"`
	ConnectionInfo       *ConnectionInfo `protobuf:"bytes,3,opt,name=connectionInfo" json:"connectionInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegistrationRequest) Reset()         { *m = RegistrationRequest{} }
func (m *RegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*RegistrationRequest) ProtoMessage()    {}
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{0}
}
func (m *RegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationRequest.Unmarshal(m, b)
}
func (m *RegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationRequest.Marshal(b, m, deterministic)
}
func (dst *RegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationRequest.Merge(dst, src)
}
func (m *RegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_RegistrationRequest.Size(m)
}
func (m *RegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationRequest proto.InternalMessageInfo

func (m *RegistrationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RegistrationRequest) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *RegistrationRequest) GetConnectionInfo() *ConnectionInfo {
	if m != nil {
		return m.ConnectionInfo
	}
	return nil
}

type ConnectionInfo struct {
	Addresses            []string     `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	Certs                []*BytesPair `protobuf:"bytes,2,rep,name=certs" json:"certs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConnectionInfo) Reset()         { *m = ConnectionInfo{} }
func (m *ConnectionInfo) String() string { return proto.CompactTextString(m) }
func (*ConnectionInfo) ProtoMessage()    {}
func (*ConnectionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{1}
}
func (m *ConnectionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionInfo.Unmarshal(m, b)
}
func (m *ConnectionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionInfo.Marshal(b, m, deterministic)
}
func (dst *ConnectionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionInfo.Merge(dst, src)
}
func (m *ConnectionInfo) XXX_Size() int {
	return xxx_messageInfo_ConnectionInfo.Size(m)
}
func (m *ConnectionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionInfo proto.InternalMessageInfo

func (m *ConnectionInfo) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ConnectionInfo) GetCerts() []*BytesPair {
	if m != nil {
		return m.Certs
	}
	return nil
}

type RegistrationResponse struct {
	Code                 uint32          `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Error                string          `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	ConnectionInfo       *ConnectionInfo `protobuf:"bytes,3,opt,name=connectionInfo" json:"connectionInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegistrationResponse) Reset()         { *m = RegistrationResponse{} }
func (m *RegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*RegistrationResponse) ProtoMessage()    {}
func (*RegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{2}
}
func (m *RegistrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrationResponse.Unmarshal(m, b)
}
func (m *RegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrationResponse.Marshal(b, m, deterministic)
}
func (dst *RegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationResponse.Merge(dst, src)
}
func (m *RegistrationResponse) XXX_Size() int {
	return xxx_messageInfo_RegistrationResponse.Size(m)
}
func (m *RegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationResponse proto.InternalMessageInfo

func (m *RegistrationResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegistrationResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RegistrationResponse) GetConnectionInfo() *ConnectionInfo {
	if m != nil {
		return m.ConnectionInfo
	}
	return nil
}

type ServiceMessage struct {
	Services             []string `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceMessage) Reset()         { *m = ServiceMessage{} }
func (m *ServiceMessage) String() string { return proto.CompactTextString(m) }
func (*ServiceMessage) ProtoMessage()    {}
func (*ServiceMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{3}
}
func (m *ServiceMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceMessage.Unmarshal(m, b)
}
func (m *ServiceMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceMessage.Marshal(b, m, deterministic)
}
func (dst *ServiceMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceMessage.Merge(dst, src)
}
func (m *ServiceMessage) XXX_Size() int {
	return xxx_messageInfo_ServiceMessage.Size(m)
}
func (m *ServiceMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceMessage proto.InternalMessageInfo

func (m *ServiceMessage) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

type Peer struct {
	ConnectionInfo       *ConnectionInfo `protobuf:"bytes,1,opt,name=connectionInfo" json:"connectionInfo,omitempty"`
	Mode                 uint32          `protobuf:"varint,2,opt,name=mode" json:"mode,omitempty"`
	Timestamp            uint64          `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Version              string          `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	Signature            []float64       `protobuf:"fixed64,5,rep,packed,name=signature" json:"signature,omitempty"`
	DoubleInfo           []*DoublePair   `protobuf:"bytes,6,rep,name=doubleInfo" json:"doubleInfo,omitempty"`
	StringInfo           []*StringPair   `protobuf:"bytes,7,rep,name=stringInfo" json:"stringInfo,omitempty"`
	BytesInfo            []*BytesPair    `protobuf:"bytes,8,rep,name=bytesInfo" json:"bytesInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{4}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (dst *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(dst, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetConnectionInfo() *ConnectionInfo {
	if m != nil {
		return m.ConnectionInfo
	}
	return nil
}

func (m *Peer) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *Peer) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Peer) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Peer) GetSignature() []float64 {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Peer) GetDoubleInfo() []*DoublePair {
	if m != nil {
		return m.DoubleInfo
	}
	return nil
}

func (m *Peer) GetStringInfo() []*StringPair {
	if m != nil {
		return m.StringInfo
	}
	return nil
}

func (m *Peer) GetBytesInfo() []*BytesPair {
	if m != nil {
		return m.BytesInfo
	}
	return nil
}

type DoublePair struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoublePair) Reset()         { *m = DoublePair{} }
func (m *DoublePair) String() string { return proto.CompactTextString(m) }
func (*DoublePair) ProtoMessage()    {}
func (*DoublePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{5}
}
func (m *DoublePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoublePair.Unmarshal(m, b)
}
func (m *DoublePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoublePair.Marshal(b, m, deterministic)
}
func (dst *DoublePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoublePair.Merge(dst, src)
}
func (m *DoublePair) XXX_Size() int {
	return xxx_messageInfo_DoublePair.Size(m)
}
func (m *DoublePair) XXX_DiscardUnknown() {
	xxx_messageInfo_DoublePair.DiscardUnknown(m)
}

var xxx_messageInfo_DoublePair proto.InternalMessageInfo

func (m *DoublePair) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DoublePair) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StringPair struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringPair) Reset()         { *m = StringPair{} }
func (m *StringPair) String() string { return proto.CompactTextString(m) }
func (*StringPair) ProtoMessage()    {}
func (*StringPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{6}
}
func (m *StringPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringPair.Unmarshal(m, b)
}
func (m *StringPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringPair.Marshal(b, m, deterministic)
}
func (dst *StringPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringPair.Merge(dst, src)
}
func (m *StringPair) XXX_Size() int {
	return xxx_messageInfo_StringPair.Size(m)
}
func (m *StringPair) XXX_DiscardUnknown() {
	xxx_messageInfo_StringPair.DiscardUnknown(m)
}

var xxx_messageInfo_StringPair proto.InternalMessageInfo

func (m *StringPair) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StringPair) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type BytesPair struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesPair) Reset()         { *m = BytesPair{} }
func (m *BytesPair) String() string { return proto.CompactTextString(m) }
func (*BytesPair) ProtoMessage()    {}
func (*BytesPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{7}
}
func (m *BytesPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesPair.Unmarshal(m, b)
}
func (m *BytesPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesPair.Marshal(b, m, deterministic)
}
func (dst *BytesPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesPair.Merge(dst, src)
}
func (m *BytesPair) XXX_Size() int {
	return xxx_messageInfo_BytesPair.Size(m)
}
func (m *BytesPair) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesPair.DiscardUnknown(m)
}

var xxx_messageInfo_BytesPair proto.InternalMessageInfo

func (m *BytesPair) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BytesPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type BytesMessage struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesMessage) Reset()         { *m = BytesMessage{} }
func (m *BytesMessage) String() string { return proto.CompactTextString(m) }
func (*BytesMessage) ProtoMessage()    {}
func (*BytesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{8}
}
func (m *BytesMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesMessage.Unmarshal(m, b)
}
func (m *BytesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesMessage.Marshal(b, m, deterministic)
}
func (dst *BytesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesMessage.Merge(dst, src)
}
func (m *BytesMessage) XXX_Size() int {
	return xxx_messageInfo_BytesMessage.Size(m)
}
func (m *BytesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BytesMessage proto.InternalMessageInfo

func (m *BytesMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PeerMessage struct {
	Peers                []*Peer  `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerMessage) Reset()         { *m = PeerMessage{} }
func (m *PeerMessage) String() string { return proto.CompactTextString(m) }
func (*PeerMessage) ProtoMessage()    {}
func (*PeerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{9}
}
func (m *PeerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerMessage.Unmarshal(m, b)
}
func (m *PeerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerMessage.Marshal(b, m, deterministic)
}
func (dst *PeerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerMessage.Merge(dst, src)
}
func (m *PeerMessage) XXX_Size() int {
	return xxx_messageInfo_PeerMessage.Size(m)
}
func (m *PeerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PeerMessage proto.InternalMessageInfo

func (m *PeerMessage) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type JoinRequest struct {
	Peer                 *Peer    `protobuf:"bytes,1,opt,name=peer" json:"peer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{10}
}
func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (dst *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(dst, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetPeer() *Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type JoinResponse struct {
	Address              string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinResponse) Reset()         { *m = JoinResponse{} }
func (m *JoinResponse) String() string { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()    {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gossip_df2c0fff6f1e2b34, []int{11}
}
func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResponse.Unmarshal(m, b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
}
func (dst *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(dst, src)
}
func (m *JoinResponse) XXX_Size() int {
	return xxx_messageInfo_JoinResponse.Size(m)
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func (m *JoinResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*RegistrationRequest)(nil), "protos.RegistrationRequest")
	proto.RegisterType((*ConnectionInfo)(nil), "protos.ConnectionInfo")
	proto.RegisterType((*RegistrationResponse)(nil), "protos.RegistrationResponse")
	proto.RegisterType((*ServiceMessage)(nil), "protos.ServiceMessage")
	proto.RegisterType((*Peer)(nil), "protos.Peer")
	proto.RegisterType((*DoublePair)(nil), "protos.DoublePair")
	proto.RegisterType((*StringPair)(nil), "protos.StringPair")
	proto.RegisterType((*BytesPair)(nil), "protos.BytesPair")
	proto.RegisterType((*BytesMessage)(nil), "protos.BytesMessage")
	proto.RegisterType((*PeerMessage)(nil), "protos.PeerMessage")
	proto.RegisterType((*JoinRequest)(nil), "protos.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "protos.JoinResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GossipClient is the client API for Gossip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GossipClient interface {
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	ExchangeServices(ctx context.Context, in *ServiceMessage, opts ...grpc.CallOption) (*ServiceMessage, error)
	ExchangePeers(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*PeerMessage, error)
}

type gossipClient struct {
	cc *grpc.ClientConn
}

func NewGossipClient(cc *grpc.ClientConn) GossipClient {
	return &gossipClient{cc}
}

func (c *gossipClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/protos.Gossip/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gossipClient) ExchangeServices(ctx context.Context, in *ServiceMessage, opts ...grpc.CallOption) (*ServiceMessage, error) {
	out := new(ServiceMessage)
	err := c.cc.Invoke(ctx, "/protos.Gossip/ExchangeServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gossipClient) ExchangePeers(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*PeerMessage, error) {
	out := new(PeerMessage)
	err := c.cc.Invoke(ctx, "/protos.Gossip/ExchangePeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GossipServer is the server API for Gossip service.
type GossipServer interface {
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	ExchangeServices(context.Context, *ServiceMessage) (*ServiceMessage, error)
	ExchangePeers(context.Context, *PeerMessage) (*PeerMessage, error)
}

func RegisterGossipServer(s *grpc.Server, srv GossipServer) {
	s.RegisterService(&_Gossip_serviceDesc, srv)
}

func _Gossip_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Gossip/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gossip_ExchangeServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).ExchangeServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Gossip/ExchangeServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).ExchangeServices(ctx, req.(*ServiceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gossip_ExchangePeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).ExchangePeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Gossip/ExchangePeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).ExchangePeers(ctx, req.(*PeerMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gossip_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Gossip",
	HandlerType: (*GossipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Gossip_Join_Handler,
		},
		{
			MethodName: "ExchangeServices",
			Handler:    _Gossip_ExchangeServices_Handler,
		},
		{
			MethodName: "ExchangePeers",
			Handler:    _Gossip_ExchangePeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gossip.proto",
}

// RegistrationClient is the client API for Registration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistrationClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (Registration_RegisterClient, error)
}

type registrationClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationClient(cc *grpc.ClientConn) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) Register(ctx context.Context, opts ...grpc.CallOption) (Registration_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Registration_serviceDesc.Streams[0], "/protos.Registration/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &registrationRegisterClient{stream}
	return x, nil
}

type Registration_RegisterClient interface {
	Send(*BytesMessage) error
	Recv() (*BytesMessage, error)
	grpc.ClientStream
}

type registrationRegisterClient struct {
	grpc.ClientStream
}

func (x *registrationRegisterClient) Send(m *BytesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *registrationRegisterClient) Recv() (*BytesMessage, error) {
	m := new(BytesMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegistrationServer is the server API for Registration service.
type RegistrationServer interface {
	Register(Registration_RegisterServer) error
}

func RegisterRegistrationServer(s *grpc.Server, srv RegistrationServer) {
	s.RegisterService(&_Registration_serviceDesc, srv)
}

func _Registration_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RegistrationServer).Register(&registrationRegisterServer{stream})
}

type Registration_RegisterServer interface {
	Send(*BytesMessage) error
	Recv() (*BytesMessage, error)
	grpc.ServerStream
}

type registrationRegisterServer struct {
	grpc.ServerStream
}

func (x *registrationRegisterServer) Send(m *BytesMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *registrationRegisterServer) Recv() (*BytesMessage, error) {
	m := new(BytesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Registration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _Registration_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gossip.proto",
}

func init() { proto.RegisterFile("gossip.proto", fileDescriptor_gossip_df2c0fff6f1e2b34) }

var fileDescriptor_gossip_df2c0fff6f1e2b34 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xae, 0xd7, 0x74, 0x6d, 0x5e, 0xb3, 0x0a, 0xbc, 0x0a, 0x45, 0xd5, 0x0e, 0x91, 0x85, 0x44,
	0x0e, 0x68, 0x83, 0x4c, 0x70, 0x01, 0x71, 0x80, 0x21, 0x04, 0x02, 0x69, 0xf2, 0x0e, 0x9c, 0xd3,
	0xf4, 0x11, 0x22, 0xd6, 0xb8, 0xd8, 0x6e, 0x05, 0x27, 0xf8, 0xd3, 0xb8, 0xf2, 0x5f, 0x21, 0xdb,
	0x71, 0xd3, 0xa0, 0x4e, 0x68, 0xe2, 0x94, 0xbc, 0xf7, 0xbe, 0xcf, 0xcf, 0xdf, 0xfb, 0x61, 0x88,
	0x4a, 0xa1, 0x54, 0xb5, 0x3a, 0x5d, 0x49, 0xa1, 0x05, 0x3d, 0xb4, 0x1f, 0xc5, 0x7e, 0xc0, 0x31,
	0xc7, 0xb2, 0x52, 0x5a, 0xe6, 0xba, 0x12, 0x35, 0xc7, 0xaf, 0x6b, 0x54, 0x9a, 0x4e, 0x61, 0xa0,
	0xc5, 0x17, 0xac, 0x63, 0x92, 0x90, 0x34, 0xe4, 0xce, 0xa0, 0x14, 0x82, 0xa5, 0x58, 0x60, 0x7c,
	0x90, 0x90, 0xf4, 0x88, 0xdb, 0x7f, 0xfa, 0x02, 0x26, 0x85, 0xa8, 0x6b, 0x2c, 0x0c, 0xfd, 0x6d,
	0xfd, 0x49, 0xc4, 0xfd, 0x84, 0xa4, 0xe3, 0xec, 0x9e, 0x4b, 0xa4, 0x4e, 0x5f, 0x75, 0xa2, 0xfc,
	0x2f, 0x34, 0xfb, 0x08, 0x93, 0x2e, 0x82, 0x9e, 0x40, 0x98, 0x2f, 0x16, 0x12, 0x95, 0x42, 0x15,
	0x93, 0xa4, 0x9f, 0x86, 0xbc, 0x75, 0xd0, 0x07, 0x30, 0x28, 0x50, 0x6a, 0x15, 0x1f, 0x24, 0xfd,
	0x74, 0x9c, 0xdd, 0xf5, 0x69, 0x5e, 0x7e, 0xd7, 0xa8, 0x2e, 0xf3, 0x4a, 0x72, 0x17, 0x67, 0x3f,
	0x09, 0x4c, 0xbb, 0xd2, 0xd4, 0x4a, 0xd4, 0x0a, 0x8d, 0x8a, 0xc2, 0xa8, 0x20, 0x4e, 0x85, 0xf9,
	0x37, 0x7a, 0x51, 0x4a, 0x21, 0xad, 0xb4, 0x90, 0x3b, 0xe3, 0xbf, 0xb5, 0x3d, 0x84, 0xc9, 0x15,
	0xca, 0x4d, 0x55, 0xe0, 0x07, 0x54, 0x2a, 0x2f, 0x91, 0xce, 0x60, 0xa4, 0x9c, 0xc7, 0x4b, 0xdb,
	0xda, 0xec, 0xf7, 0x01, 0x04, 0x97, 0x88, 0xfb, 0xd2, 0x92, 0xdb, 0xa4, 0xdd, 0xdb, 0xa6, 0x13,
	0x08, 0x75, 0xb5, 0x44, 0xa5, 0xf3, 0xe5, 0xca, 0xaa, 0x08, 0x78, 0xeb, 0xa0, 0x31, 0x0c, 0x37,
	0x28, 0x55, 0x25, 0xea, 0x38, 0xb0, 0x05, 0xf0, 0xa6, 0xe1, 0xa9, 0xaa, 0xac, 0x73, 0xbd, 0x96,
	0x18, 0x0f, 0x92, 0x7e, 0x4a, 0x78, 0xeb, 0xa0, 0x19, 0xc0, 0x42, 0xac, 0xe7, 0xd7, 0x68, 0x6f,
	0x79, 0x68, 0x3b, 0x42, 0xfd, 0x2d, 0x2f, 0x6c, 0xc4, 0xb6, 0x64, 0x07, 0x65, 0x38, 0x4a, 0xcb,
	0xaa, 0x2e, 0x2d, 0x67, 0xd8, 0xe5, 0x5c, 0xd9, 0x88, 0xe3, 0xb4, 0x28, 0x7a, 0x06, 0xe1, 0xdc,
	0xf4, 0xd7, 0x52, 0x46, 0x37, 0x35, 0xbe, 0xc5, 0xb0, 0xa7, 0x00, 0x6d, 0x7a, 0x53, 0x90, 0x3a,
	0x5f, 0x62, 0x33, 0xcc, 0xf6, 0xdf, 0x74, 0x7c, 0x93, 0x5f, 0xaf, 0x5d, 0x95, 0x08, 0x77, 0x86,
	0xe1, 0xb5, 0x57, 0xb8, 0x05, 0xef, 0x09, 0x84, 0xdb, 0x7b, 0xfc, 0x9b, 0x16, 0x79, 0xda, 0x7d,
	0x88, 0x2c, 0xcd, 0x8f, 0xc7, 0x16, 0x45, 0x76, 0x51, 0x8f, 0x61, 0x6c, 0xe6, 0xc2, 0x83, 0x18,
	0x0c, 0x56, 0x88, 0xd2, 0x0d, 0xd0, 0x38, 0x8b, 0x7c, 0x21, 0x0c, 0x86, 0xbb, 0x10, 0x3b, 0x83,
	0xf1, 0x3b, 0x51, 0x6d, 0xd7, 0x39, 0x81, 0xc0, 0xf8, 0x9b, 0x39, 0xea, 0x32, 0x6c, 0x84, 0xa5,
	0x10, 0x39, 0x42, 0xb3, 0x24, 0x31, 0x0c, 0x9b, 0x9d, 0x6b, 0x64, 0x78, 0x33, 0xfb, 0x45, 0xe0,
	0xf0, 0x8d, 0x7d, 0x4a, 0xe8, 0x39, 0x04, 0x86, 0x44, 0x8f, 0xfd, 0x81, 0x3b, 0x39, 0x67, 0xd3,
	0xae, 0xd3, 0x9d, 0xcb, 0x7a, 0xf4, 0x02, 0xee, 0xbc, 0xfe, 0x56, 0x7c, 0xce, 0xeb, 0x12, 0x9b,
	0xe5, 0x50, 0x74, 0x3b, 0xd9, 0xdd, 0x75, 0x99, 0xdd, 0xe0, 0x67, 0x3d, 0xfa, 0x0c, 0x8e, 0xfc,
	0x29, 0x46, 0x85, 0x6a, 0xef, 0xb0, 0x53, 0xaa, 0xd9, 0x3e, 0x27, 0xeb, 0x65, 0xef, 0x21, 0xda,
	0x7d, 0x19, 0xe8, 0x73, 0x18, 0x39, 0x1b, 0x25, 0x9d, 0x76, 0xe6, 0xca, 0x1f, 0xb4, 0xd7, 0xcb,
	0x7a, 0x29, 0x79, 0x44, 0xe6, 0xee, 0x29, 0x3d, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x93,
	0xf0, 0x3f, 0x61, 0x05, 0x00, 0x00,
}
