// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/vpp/nat/nat.proto

package vpp_nat

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DNat44_Protocol int32

const (
	DNat44_TCP  DNat44_Protocol = 0
	DNat44_UDP  DNat44_Protocol = 1
	DNat44_ICMP DNat44_Protocol = 2
)

var DNat44_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
	2: "ICMP",
}

var DNat44_Protocol_value = map[string]int32{
	"TCP":  0,
	"UDP":  1,
	"ICMP": 2,
}

func (x DNat44_Protocol) String() string {
	return proto.EnumName(DNat44_Protocol_name, int32(x))
}

func (DNat44_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_74abe0ababd815c7, []int{1, 0}
}

type DNat44_StaticMapping_TwiceNatMode int32

const (
	DNat44_StaticMapping_DISABLED DNat44_StaticMapping_TwiceNatMode = 0
	DNat44_StaticMapping_ENABLED  DNat44_StaticMapping_TwiceNatMode = 1
	DNat44_StaticMapping_SELF     DNat44_StaticMapping_TwiceNatMode = 2
)

var DNat44_StaticMapping_TwiceNatMode_name = map[int32]string{
	0: "DISABLED",
	1: "ENABLED",
	2: "SELF",
}

var DNat44_StaticMapping_TwiceNatMode_value = map[string]int32{
	"DISABLED": 0,
	"ENABLED":  1,
	"SELF":     2,
}

func (x DNat44_StaticMapping_TwiceNatMode) String() string {
	return proto.EnumName(DNat44_StaticMapping_TwiceNatMode_name, int32(x))
}

func (DNat44_StaticMapping_TwiceNatMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_74abe0ababd815c7, []int{1, 0, 0}
}

type Nat44Global struct {
	Forwarding           bool                     `protobuf:"varint,1,opt,name=forwarding,proto3" json:"forwarding,omitempty"`
	NatInterfaces        []*Nat44Global_Interface `protobuf:"bytes,2,rep,name=nat_interfaces,json=natInterfaces,proto3" json:"nat_interfaces,omitempty"`
	AddressPool          []*Nat44Global_Address   `protobuf:"bytes,3,rep,name=address_pool,json=addressPool,proto3" json:"address_pool,omitempty"`
	VirtualReassembly    *VirtualReassembly       `protobuf:"bytes,4,opt,name=virtual_reassembly,json=virtualReassembly,proto3" json:"virtual_reassembly,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Nat44Global) Reset()         { *m = Nat44Global{} }
func (m *Nat44Global) String() string { return proto.CompactTextString(m) }
func (*Nat44Global) ProtoMessage()    {}
func (*Nat44Global) Descriptor() ([]byte, []int) {
	return fileDescriptor_74abe0ababd815c7, []int{0}
}

func (m *Nat44Global) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nat44Global.Unmarshal(m, b)
}
func (m *Nat44Global) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nat44Global.Marshal(b, m, deterministic)
}
func (m *Nat44Global) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nat44Global.Merge(m, src)
}
func (m *Nat44Global) XXX_Size() int {
	return xxx_messageInfo_Nat44Global.Size(m)
}
func (m *Nat44Global) XXX_DiscardUnknown() {
	xxx_messageInfo_Nat44Global.DiscardUnknown(m)
}

var xxx_messageInfo_Nat44Global proto.InternalMessageInfo

func (m *Nat44Global) GetForwarding() bool {
	if m != nil {
		return m.Forwarding
	}
	return false
}

func (m *Nat44Global) GetNatInterfaces() []*Nat44Global_Interface {
	if m != nil {
		return m.NatInterfaces
	}
	return nil
}

func (m *Nat44Global) GetAddressPool() []*Nat44Global_Address {
	if m != nil {
		return m.AddressPool
	}
	return nil
}

func (m *Nat44Global) GetVirtualReassembly() *VirtualReassembly {
	if m != nil {
		return m.VirtualReassembly
	}
	return nil
}

type Nat44Global_Interface struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsInside             bool     `protobuf:"varint,2,opt,name=is_inside,json=isInside,proto3" json:"is_inside,omitempty"`
	OutputFeature        bool     `protobuf:"varint,3,opt,name=output_feature,json=outputFeature,proto3" json:"output_feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nat44Global_Interface) Reset()         { *m = Nat44Global_Interface{} }
func (m *Nat44Global_Interface) String() string { return proto.CompactTextString(m) }
func (*Nat44Global_Interface) ProtoMessage()    {}
func (*Nat44Global_Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_74abe0ababd815c7, []int{0, 0}
}

func (m *Nat44Global_Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nat44Global_Interface.Unmarshal(m, b)
}
func (m *Nat44Global_Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nat44Global_Interface.Marshal(b, m, deterministic)
}
func (m *Nat44Global_Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nat44Global_Interface.Merge(m, src)
}
func (m *Nat44Global_Interface) XXX_Size() int {
	return xxx_messageInfo_Nat44Global_Interface.Size(m)
}
func (m *Nat44Global_Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_Nat44Global_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_Nat44Global_Interface proto.InternalMessageInfo

func (m *Nat44Global_Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Nat44Global_Interface) GetIsInside() bool {
	if m != nil {
		return m.IsInside
	}
	return false
}

func (m *Nat44Global_Interface) GetOutputFeature() bool {
	if m != nil {
		return m.OutputFeature
	}
	return false
}

type Nat44Global_Address struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	VrfId                uint32   `protobuf:"varint,2,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	TwiceNat             bool     `protobuf:"varint,3,opt,name=twice_nat,json=twiceNat,proto3" json:"twice_nat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nat44Global_Address) Reset()         { *m = Nat44Global_Address{} }
func (m *Nat44Global_Address) String() string { return proto.CompactTextString(m) }
func (*Nat44Global_Address) ProtoMessage()    {}
func (*Nat44Global_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_74abe0ababd815c7, []int{0, 1}
}

func (m *Nat44Global_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nat44Global_Address.Unmarshal(m, b)
}
func (m *Nat44Global_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nat44Global_Address.Marshal(b, m, deterministic)
}
func (m *Nat44Global_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nat44Global_Address.Merge(m, src)
}
func (m *Nat44Global_Address) XXX_Size() int {
	return xxx_messageInfo_Nat44Global_Address.Size(m)
}
func (m *Nat44Global_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Nat44Global_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Nat44Global_Address proto.InternalMessageInfo

func (m *Nat44Global_Address) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Nat44Global_Address) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Nat44Global_Address) GetTwiceNat() bool {
	if m != nil {
		return m.TwiceNat
	}
	return false
}

type DNat44 struct {
	Label                string                    `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	StMappings           []*DNat44_StaticMapping   `protobuf:"bytes,2,rep,name=st_mappings,json=stMappings,proto3" json:"st_mappings,omitempty"`
	IdMappings           []*DNat44_IdentityMapping `protobuf:"bytes,3,rep,name=id_mappings,json=idMappings,proto3" json:"id_mappings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DNat44) Reset()         { *m = DNat44{} }
func (m *DNat44) String() string { return proto.CompactTextString(m) }
func (*DNat44) ProtoMessage()    {}
func (*DNat44) Descriptor() ([]byte, []int) {
	return fileDescriptor_74abe0ababd815c7, []int{1}
}

func (m *DNat44) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNat44.Unmarshal(m, b)
}
func (m *DNat44) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNat44.Marshal(b, m, deterministic)
}
func (m *DNat44) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNat44.Merge(m, src)
}
func (m *DNat44) XXX_Size() int {
	return xxx_messageInfo_DNat44.Size(m)
}
func (m *DNat44) XXX_DiscardUnknown() {
	xxx_messageInfo_DNat44.DiscardUnknown(m)
}

var xxx_messageInfo_DNat44 proto.InternalMessageInfo

func (m *DNat44) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *DNat44) GetStMappings() []*DNat44_StaticMapping {
	if m != nil {
		return m.StMappings
	}
	return nil
}

func (m *DNat44) GetIdMappings() []*DNat44_IdentityMapping {
	if m != nil {
		return m.IdMappings
	}
	return nil
}

type DNat44_StaticMapping struct {
	ExternalInterface    string                            `protobuf:"bytes,1,opt,name=external_interface,json=externalInterface,proto3" json:"external_interface,omitempty"`
	ExternalIp           string                            `protobuf:"bytes,2,opt,name=external_ip,json=externalIp,proto3" json:"external_ip,omitempty"`
	ExternalPort         uint32                            `protobuf:"varint,3,opt,name=external_port,json=externalPort,proto3" json:"external_port,omitempty"`
	LocalIps             []*DNat44_StaticMapping_LocalIP   `protobuf:"bytes,4,rep,name=local_ips,json=localIps,proto3" json:"local_ips,omitempty"`
	Protocol             DNat44_Protocol                   `protobuf:"varint,5,opt,name=protocol,proto3,enum=vpp.nat.DNat44_Protocol" json:"protocol,omitempty"`
	TwiceNat             DNat44_StaticMapping_TwiceNatMode `protobuf:"varint,6,opt,name=twice_nat,json=twiceNat,proto3,enum=vpp.nat.DNat44_StaticMapping_TwiceNatMode" json:"twice_nat,omitempty"`
	SessionAffinity      uint32                            `protobuf:"varint,7,opt,name=session_affinity,json=sessionAffinity,proto3" json:"session_affinity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *DNat44_StaticMapping) Reset()         { *m = DNat44_StaticMapping{} }
func (m *DNat44_StaticMapping) String() string { return proto.CompactTextString(m) }
func (*DNat44_StaticMapping) ProtoMessage()    {}
func (*DNat44_StaticMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_74abe0ababd815c7, []int{1, 0}
}

func (m *DNat44_StaticMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNat44_StaticMapping.Unmarshal(m, b)
}
func (m *DNat44_StaticMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNat44_StaticMapping.Marshal(b, m, deterministic)
}
func (m *DNat44_StaticMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNat44_StaticMapping.Merge(m, src)
}
func (m *DNat44_StaticMapping) XXX_Size() int {
	return xxx_messageInfo_DNat44_StaticMapping.Size(m)
}
func (m *DNat44_StaticMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_DNat44_StaticMapping.DiscardUnknown(m)
}

var xxx_messageInfo_DNat44_StaticMapping proto.InternalMessageInfo

func (m *DNat44_StaticMapping) GetExternalInterface() string {
	if m != nil {
		return m.ExternalInterface
	}
	return ""
}

func (m *DNat44_StaticMapping) GetExternalIp() string {
	if m != nil {
		return m.ExternalIp
	}
	return ""
}

func (m *DNat44_StaticMapping) GetExternalPort() uint32 {
	if m != nil {
		return m.ExternalPort
	}
	return 0
}

func (m *DNat44_StaticMapping) GetLocalIps() []*DNat44_StaticMapping_LocalIP {
	if m != nil {
		return m.LocalIps
	}
	return nil
}

func (m *DNat44_StaticMapping) GetProtocol() DNat44_Protocol {
	if m != nil {
		return m.Protocol
	}
	return DNat44_TCP
}

func (m *DNat44_StaticMapping) GetTwiceNat() DNat44_StaticMapping_TwiceNatMode {
	if m != nil {
		return m.TwiceNat
	}
	return DNat44_StaticMapping_DISABLED
}

func (m *DNat44_StaticMapping) GetSessionAffinity() uint32 {
	if m != nil {
		return m.SessionAffinity
	}
	return 0
}

type DNat44_StaticMapping_LocalIP struct {
	VrfId                uint32   `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	LocalIp              string   `protobuf:"bytes,2,opt,name=local_ip,json=localIp,proto3" json:"local_ip,omitempty"`
	LocalPort            uint32   `protobuf:"varint,3,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	Probability          uint32   `protobuf:"varint,4,opt,name=probability,proto3" json:"probability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNat44_StaticMapping_LocalIP) Reset()         { *m = DNat44_StaticMapping_LocalIP{} }
func (m *DNat44_StaticMapping_LocalIP) String() string { return proto.CompactTextString(m) }
func (*DNat44_StaticMapping_LocalIP) ProtoMessage()    {}
func (*DNat44_StaticMapping_LocalIP) Descriptor() ([]byte, []int) {
	return fileDescriptor_74abe0ababd815c7, []int{1, 0, 0}
}

func (m *DNat44_StaticMapping_LocalIP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNat44_StaticMapping_LocalIP.Unmarshal(m, b)
}
func (m *DNat44_StaticMapping_LocalIP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNat44_StaticMapping_LocalIP.Marshal(b, m, deterministic)
}
func (m *DNat44_StaticMapping_LocalIP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNat44_StaticMapping_LocalIP.Merge(m, src)
}
func (m *DNat44_StaticMapping_LocalIP) XXX_Size() int {
	return xxx_messageInfo_DNat44_StaticMapping_LocalIP.Size(m)
}
func (m *DNat44_StaticMapping_LocalIP) XXX_DiscardUnknown() {
	xxx_messageInfo_DNat44_StaticMapping_LocalIP.DiscardUnknown(m)
}

var xxx_messageInfo_DNat44_StaticMapping_LocalIP proto.InternalMessageInfo

func (m *DNat44_StaticMapping_LocalIP) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *DNat44_StaticMapping_LocalIP) GetLocalIp() string {
	if m != nil {
		return m.LocalIp
	}
	return ""
}

func (m *DNat44_StaticMapping_LocalIP) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *DNat44_StaticMapping_LocalIP) GetProbability() uint32 {
	if m != nil {
		return m.Probability
	}
	return 0
}

type DNat44_IdentityMapping struct {
	VrfId                uint32          `protobuf:"varint,1,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	Interface            string          `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	IpAddress            string          `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Port                 uint32          `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Protocol             DNat44_Protocol `protobuf:"varint,5,opt,name=protocol,proto3,enum=vpp.nat.DNat44_Protocol" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DNat44_IdentityMapping) Reset()         { *m = DNat44_IdentityMapping{} }
func (m *DNat44_IdentityMapping) String() string { return proto.CompactTextString(m) }
func (*DNat44_IdentityMapping) ProtoMessage()    {}
func (*DNat44_IdentityMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_74abe0ababd815c7, []int{1, 1}
}

func (m *DNat44_IdentityMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNat44_IdentityMapping.Unmarshal(m, b)
}
func (m *DNat44_IdentityMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNat44_IdentityMapping.Marshal(b, m, deterministic)
}
func (m *DNat44_IdentityMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNat44_IdentityMapping.Merge(m, src)
}
func (m *DNat44_IdentityMapping) XXX_Size() int {
	return xxx_messageInfo_DNat44_IdentityMapping.Size(m)
}
func (m *DNat44_IdentityMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_DNat44_IdentityMapping.DiscardUnknown(m)
}

var xxx_messageInfo_DNat44_IdentityMapping proto.InternalMessageInfo

func (m *DNat44_IdentityMapping) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *DNat44_IdentityMapping) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *DNat44_IdentityMapping) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *DNat44_IdentityMapping) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *DNat44_IdentityMapping) GetProtocol() DNat44_Protocol {
	if m != nil {
		return m.Protocol
	}
	return DNat44_TCP
}

type VirtualReassembly struct {
	Timeout              uint32   `protobuf:"varint,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	MaxReassemblies      uint32   `protobuf:"varint,2,opt,name=max_reassemblies,json=maxReassemblies,proto3" json:"max_reassemblies,omitempty"`
	MaxFragments         uint32   `protobuf:"varint,3,opt,name=max_fragments,json=maxFragments,proto3" json:"max_fragments,omitempty"`
	DropFragments        bool     `protobuf:"varint,4,opt,name=drop_fragments,json=dropFragments,proto3" json:"drop_fragments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualReassembly) Reset()         { *m = VirtualReassembly{} }
func (m *VirtualReassembly) String() string { return proto.CompactTextString(m) }
func (*VirtualReassembly) ProtoMessage()    {}
func (*VirtualReassembly) Descriptor() ([]byte, []int) {
	return fileDescriptor_74abe0ababd815c7, []int{2}
}

func (m *VirtualReassembly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualReassembly.Unmarshal(m, b)
}
func (m *VirtualReassembly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualReassembly.Marshal(b, m, deterministic)
}
func (m *VirtualReassembly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualReassembly.Merge(m, src)
}
func (m *VirtualReassembly) XXX_Size() int {
	return xxx_messageInfo_VirtualReassembly.Size(m)
}
func (m *VirtualReassembly) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualReassembly.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualReassembly proto.InternalMessageInfo

func (m *VirtualReassembly) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *VirtualReassembly) GetMaxReassemblies() uint32 {
	if m != nil {
		return m.MaxReassemblies
	}
	return 0
}

func (m *VirtualReassembly) GetMaxFragments() uint32 {
	if m != nil {
		return m.MaxFragments
	}
	return 0
}

func (m *VirtualReassembly) GetDropFragments() bool {
	if m != nil {
		return m.DropFragments
	}
	return false
}

func init() {
	proto.RegisterEnum("vpp.nat.DNat44_Protocol", DNat44_Protocol_name, DNat44_Protocol_value)
	proto.RegisterEnum("vpp.nat.DNat44_StaticMapping_TwiceNatMode", DNat44_StaticMapping_TwiceNatMode_name, DNat44_StaticMapping_TwiceNatMode_value)
	proto.RegisterType((*Nat44Global)(nil), "vpp.nat.Nat44Global")
	proto.RegisterType((*Nat44Global_Interface)(nil), "vpp.nat.Nat44Global.Interface")
	proto.RegisterType((*Nat44Global_Address)(nil), "vpp.nat.Nat44Global.Address")
	proto.RegisterType((*DNat44)(nil), "vpp.nat.DNat44")
	proto.RegisterType((*DNat44_StaticMapping)(nil), "vpp.nat.DNat44.StaticMapping")
	proto.RegisterType((*DNat44_StaticMapping_LocalIP)(nil), "vpp.nat.DNat44.StaticMapping.LocalIP")
	proto.RegisterType((*DNat44_IdentityMapping)(nil), "vpp.nat.DNat44.IdentityMapping")
	proto.RegisterType((*VirtualReassembly)(nil), "vpp.nat.VirtualReassembly")
}

func init() { proto.RegisterFile("models/vpp/nat/nat.proto", fileDescriptor_74abe0ababd815c7) }

var fileDescriptor_74abe0ababd815c7 = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x6b, 0x23, 0x37,
	0x14, 0x5d, 0xc7, 0x4e, 0x66, 0x7c, 0x27, 0xce, 0x3a, 0xa2, 0x85, 0xa9, 0xbb, 0x1f, 0xc6, 0x65,
	0x4b, 0x5a, 0xd8, 0x31, 0x64, 0x43, 0x29, 0x14, 0xda, 0x26, 0x9b, 0x64, 0x19, 0x48, 0x82, 0x99,
	0x6c, 0x5b, 0xe8, 0xcb, 0x20, 0x7b, 0x34, 0xae, 0x60, 0x46, 0x12, 0x92, 0xec, 0x4d, 0xa0, 0xbf,
	0xa6, 0xb4, 0xcf, 0xfd, 0x07, 0xfd, 0x6d, 0x65, 0x24, 0xcd, 0xc7, 0xa6, 0xed, 0x3e, 0xf4, 0xc1,
	0x20, 0x1d, 0x1d, 0x1d, 0xdd, 0x7b, 0xee, 0x9d, 0x6b, 0x08, 0x4b, 0x9e, 0x91, 0x42, 0xcd, 0xb7,
	0x42, 0xcc, 0x19, 0xd6, 0xd5, 0x2f, 0x12, 0x92, 0x6b, 0x8e, 0xbc, 0xad, 0x10, 0x11, 0xc3, 0x7a,
	0xf6, 0x57, 0x1f, 0x82, 0x1b, 0xac, 0x4f, 0x4e, 0xde, 0x14, 0x7c, 0x89, 0x0b, 0xf4, 0x0c, 0x20,
	0xe7, 0xf2, 0x1d, 0x96, 0x19, 0x65, 0xeb, 0xb0, 0x37, 0xed, 0x1d, 0xf9, 0x49, 0x07, 0x41, 0x17,
	0x70, 0xc0, 0xb0, 0x4e, 0x29, 0xd3, 0x44, 0xe6, 0x78, 0x45, 0x54, 0xb8, 0x33, 0xed, 0x1f, 0x05,
	0xc7, 0xcf, 0x22, 0xa7, 0x18, 0x75, 0xd4, 0xa2, 0xb8, 0xa6, 0x25, 0x23, 0x86, 0x75, 0xb3, 0x53,
	0xe8, 0x3b, 0xd8, 0xc7, 0x59, 0x26, 0x89, 0x52, 0xa9, 0xe0, 0xbc, 0x08, 0xfb, 0x46, 0xe4, 0xc9,
	0xbf, 0x8a, 0x9c, 0x5a, 0x62, 0x12, 0xb8, 0x1b, 0x0b, 0xce, 0x0b, 0x14, 0x03, 0xda, 0x52, 0xa9,
	0x37, 0xb8, 0x48, 0x25, 0xc1, 0x4a, 0x91, 0x72, 0x59, 0xdc, 0x87, 0x83, 0x69, 0xef, 0x28, 0x38,
	0x9e, 0x34, 0x32, 0x3f, 0x5a, 0x4a, 0xd2, 0x30, 0x92, 0xc3, 0xed, 0x43, 0x68, 0xb2, 0x82, 0x61,
	0x13, 0x19, 0x42, 0x30, 0x60, 0xb8, 0x24, 0x26, 0xf3, 0x61, 0x62, 0xd6, 0xe8, 0x53, 0x18, 0x52,
	0x95, 0x52, 0xa6, 0x68, 0x46, 0xc2, 0x1d, 0x63, 0x89, 0x4f, 0x55, 0x6c, 0xf6, 0xe8, 0x05, 0x1c,
	0xf0, 0x8d, 0x16, 0x1b, 0x9d, 0xe6, 0x04, 0xeb, 0x8d, 0x24, 0x61, 0xdf, 0x30, 0x46, 0x16, 0xbd,
	0xb4, 0xe0, 0xe4, 0x27, 0xf0, 0x5c, 0x1e, 0x28, 0x04, 0xcf, 0x65, 0xe2, 0x5e, 0xa9, 0xb7, 0xe8,
	0x63, 0xd8, 0xdb, 0xca, 0x3c, 0xa5, 0x99, 0x79, 0x65, 0x94, 0xec, 0x6e, 0x65, 0x1e, 0x67, 0xd5,
	0xfb, 0xfa, 0x1d, 0x5d, 0x91, 0x94, 0x61, 0xed, 0xd4, 0x7d, 0x03, 0xdc, 0x60, 0x3d, 0xfb, 0xcd,
	0x83, 0xbd, 0x73, 0x63, 0x17, 0xfa, 0x08, 0x76, 0x0b, 0xbc, 0x24, 0x85, 0x93, 0xb5, 0x1b, 0xf4,
	0x2d, 0x04, 0x4a, 0xa7, 0x25, 0x16, 0x82, 0xb2, 0x75, 0x5d, 0xae, 0xa7, 0x8d, 0x45, 0xf6, 0x6e,
	0x74, 0xab, 0xb1, 0xa6, 0xab, 0x6b, 0xcb, 0x4a, 0x40, 0x69, 0xb7, 0x54, 0xe8, 0x7b, 0x08, 0x68,
	0xd6, 0xde, 0xb7, 0x95, 0x7a, 0xfe, 0xf0, 0x7e, 0x9c, 0x11, 0xa6, 0xa9, 0xbe, 0x6f, 0x14, 0x68,
	0x56, 0x2b, 0x4c, 0xfe, 0x18, 0xc0, 0xe8, 0x3d, 0x7d, 0xf4, 0x12, 0x10, 0xb9, 0xd3, 0x44, 0x32,
	0x5c, 0xb4, 0xad, 0xe4, 0xc2, 0x3e, 0xac, 0x4f, 0xda, 0xa2, 0x3c, 0x87, 0xa0, 0xa5, 0x0b, 0x63,
	0xce, 0x30, 0x81, 0x86, 0x27, 0xd0, 0x67, 0x30, 0x6a, 0x08, 0x82, 0x4b, 0xeb, 0xd2, 0x28, 0xd9,
	0xaf, 0xc1, 0x05, 0x97, 0x1a, 0x9d, 0xc1, 0xb0, 0xe0, 0x2b, 0x23, 0xa1, 0xc2, 0x81, 0x49, 0xe3,
	0xc5, 0x07, 0x6d, 0x88, 0xae, 0x2a, 0x7a, 0xbc, 0x48, 0x7c, 0x73, 0x2f, 0x16, 0x0a, 0x9d, 0x80,
	0x6f, 0x3e, 0xa0, 0x15, 0x2f, 0xc2, 0xdd, 0x69, 0xef, 0xe8, 0xe0, 0x38, 0x7c, 0x28, 0xb1, 0x70,
	0xe7, 0x49, 0xc3, 0x44, 0x6f, 0xba, 0x05, 0xdc, 0x33, 0xd7, 0xbe, 0xfc, 0xf0, 0xcb, 0x6f, 0x5d,
	0x79, 0xaf, 0x79, 0x46, 0xda, 0x62, 0xa3, 0x2f, 0x60, 0xac, 0x88, 0x52, 0x94, 0xb3, 0x14, 0xe7,
	0x39, 0x65, 0x54, 0xdf, 0x87, 0x9e, 0x49, 0xf5, 0xb1, 0xc3, 0x4f, 0x1d, 0x3c, 0xf9, 0x15, 0x3c,
	0x17, 0x7e, 0xa7, 0xad, 0x7a, 0xdd, 0xb6, 0xfa, 0x04, 0xfc, 0xda, 0x0f, 0x67, 0xa9, 0xe7, 0xf2,
	0x44, 0x4f, 0x01, 0xec, 0x51, 0xc7, 0x4c, 0x6b, 0x9e, 0x71, 0x72, 0x0a, 0x81, 0x90, 0x7c, 0x89,
	0x97, 0xb4, 0xa8, 0x22, 0x18, 0x98, 0xf3, 0x2e, 0x34, 0x7b, 0x05, 0xfb, 0xdd, 0x14, 0xd0, 0x3e,
	0xf8, 0xe7, 0xf1, 0xed, 0xe9, 0xd9, 0xd5, 0xc5, 0xf9, 0xf8, 0x11, 0x0a, 0xc0, 0xbb, 0xb8, 0xb1,
	0x9b, 0x1e, 0xf2, 0x61, 0x70, 0x7b, 0x71, 0x75, 0x39, 0xde, 0x99, 0xfc, 0xd9, 0x83, 0xc7, 0x0f,
	0xfa, 0xe8, 0xbf, 0x62, 0x7f, 0x02, 0xc3, 0xb6, 0x6f, 0x6c, 0xf0, 0x2d, 0x50, 0x85, 0x4f, 0x45,
	0x5a, 0x7f, 0x64, 0x7d, 0x77, 0x2c, 0xea, 0x0f, 0x10, 0xc1, 0xc0, 0xe4, 0x65, 0xe3, 0x36, 0xeb,
	0xff, 0x57, 0xd8, 0xd9, 0xe7, 0xe0, 0xd7, 0x28, 0xf2, 0xa0, 0xff, 0xf6, 0xf5, 0x62, 0xfc, 0xa8,
	0x5a, 0xfc, 0x70, 0xbe, 0xb0, 0x99, 0xc5, 0xaf, 0xaf, 0x17, 0xe3, 0x9d, 0xd9, 0xef, 0x3d, 0x38,
	0xfc, 0xc7, 0x2c, 0xaa, 0x06, 0x81, 0xa6, 0x25, 0xe1, 0x1b, 0xed, 0x92, 0xab, 0xb7, 0x55, 0x9d,
	0x4b, 0x7c, 0xd7, 0x4e, 0x36, 0x6a, 0xe6, 0xac, 0xa9, 0x73, 0x89, 0xef, 0x92, 0x0e, 0x5c, 0xb5,
	0x7e, 0x45, 0xcd, 0x25, 0x5e, 0x97, 0x84, 0x69, 0x55, 0xb7, 0x7e, 0x89, 0xef, 0x2e, 0x6b, 0xac,
	0x1a, 0x52, 0x99, 0xe4, 0xa2, 0xc3, 0x1a, 0xd8, 0x21, 0x55, 0xa1, 0x0d, 0xed, 0xec, 0xeb, 0x9f,
	0xbf, 0x5a, 0x53, 0xfd, 0xcb, 0x66, 0x19, 0xad, 0x78, 0x39, 0x2f, 0xe8, 0x1a, 0x6b, 0x5e, 0xfd,
	0x79, 0xbc, 0xc4, 0x6b, 0xc2, 0xf4, 0x1c, 0x0b, 0x3a, 0x7f, 0xff, 0x1f, 0xe5, 0x9b, 0xad, 0x10,
	0x55, 0x53, 0x2f, 0xf7, 0x8c, 0x25, 0xaf, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x73, 0x32, 0x8d,
	0xd2, 0x72, 0x06, 0x00, 0x00,
}
