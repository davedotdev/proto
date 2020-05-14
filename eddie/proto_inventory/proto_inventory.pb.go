// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto_inventory/proto_inventory.proto

package proto_inventory

import (
	fmt "fmt"
	proto_common "github.com/davedotdev/proto/eddie/proto_common"
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

type DEVICE struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Vendor               string                 `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Platform             string                 `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	Routeengine          []*RouteEngine         `protobuf:"bytes,4,rep,name=routeengine,proto3" json:"routeengine,omitempty"`
	Linecard             []*Linecard            `protobuf:"bytes,5,rep,name=linecard,proto3" json:"linecard,omitempty"`
	Operation            proto_common.Operation `protobuf:"varint,6,opt,name=operation,proto3,enum=Operation" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DEVICE) Reset()         { *m = DEVICE{} }
func (m *DEVICE) String() string { return proto.CompactTextString(m) }
func (*DEVICE) ProtoMessage()    {}
func (*DEVICE) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd0852c13d032805, []int{0}
}

func (m *DEVICE) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DEVICE.Unmarshal(m, b)
}
func (m *DEVICE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DEVICE.Marshal(b, m, deterministic)
}
func (m *DEVICE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DEVICE.Merge(m, src)
}
func (m *DEVICE) XXX_Size() int {
	return xxx_messageInfo_DEVICE.Size(m)
}
func (m *DEVICE) XXX_DiscardUnknown() {
	xxx_messageInfo_DEVICE.DiscardUnknown(m)
}

var xxx_messageInfo_DEVICE proto.InternalMessageInfo

func (m *DEVICE) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DEVICE) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *DEVICE) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *DEVICE) GetRouteengine() []*RouteEngine {
	if m != nil {
		return m.Routeengine
	}
	return nil
}

func (m *DEVICE) GetLinecard() []*Linecard {
	if m != nil {
		return m.Linecard
	}
	return nil
}

func (m *DEVICE) GetOperation() proto_common.Operation {
	if m != nil {
		return m.Operation
	}
	return proto_common.Operation_CREATE
}

type Linecard struct {
	Slot                 uint32                 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Model                string                 `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	PortType             string                 `protobuf:"bytes,3,opt,name=portType,proto3" json:"portType,omitempty"`
	Ports                uint32                 `protobuf:"varint,4,opt,name=ports,proto3" json:"ports,omitempty"`
	Explicit             bool                   `protobuf:"varint,5,opt,name=explicit,proto3" json:"explicit,omitempty"`
	Operation            proto_common.Operation `protobuf:"varint,6,opt,name=operation,proto3,enum=Operation" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Linecard) Reset()         { *m = Linecard{} }
func (m *Linecard) String() string { return proto.CompactTextString(m) }
func (*Linecard) ProtoMessage()    {}
func (*Linecard) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd0852c13d032805, []int{1}
}

func (m *Linecard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Linecard.Unmarshal(m, b)
}
func (m *Linecard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Linecard.Marshal(b, m, deterministic)
}
func (m *Linecard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Linecard.Merge(m, src)
}
func (m *Linecard) XXX_Size() int {
	return xxx_messageInfo_Linecard.Size(m)
}
func (m *Linecard) XXX_DiscardUnknown() {
	xxx_messageInfo_Linecard.DiscardUnknown(m)
}

var xxx_messageInfo_Linecard proto.InternalMessageInfo

func (m *Linecard) GetSlot() uint32 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *Linecard) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Linecard) GetPortType() string {
	if m != nil {
		return m.PortType
	}
	return ""
}

func (m *Linecard) GetPorts() uint32 {
	if m != nil {
		return m.Ports
	}
	return 0
}

func (m *Linecard) GetExplicit() bool {
	if m != nil {
		return m.Explicit
	}
	return false
}

func (m *Linecard) GetOperation() proto_common.Operation {
	if m != nil {
		return m.Operation
	}
	return proto_common.Operation_CREATE
}

type RouteEngine struct {
	Slot                 uint32                 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Model                string                 `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	MgmtIf               string                 `protobuf:"bytes,3,opt,name=mgmt_if,json=mgmtIf,proto3" json:"mgmt_if,omitempty"`
	Operation            proto_common.Operation `protobuf:"varint,4,opt,name=operation,proto3,enum=Operation" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteEngine) Reset()         { *m = RouteEngine{} }
func (m *RouteEngine) String() string { return proto.CompactTextString(m) }
func (*RouteEngine) ProtoMessage()    {}
func (*RouteEngine) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd0852c13d032805, []int{2}
}

func (m *RouteEngine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteEngine.Unmarshal(m, b)
}
func (m *RouteEngine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteEngine.Marshal(b, m, deterministic)
}
func (m *RouteEngine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteEngine.Merge(m, src)
}
func (m *RouteEngine) XXX_Size() int {
	return xxx_messageInfo_RouteEngine.Size(m)
}
func (m *RouteEngine) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteEngine.DiscardUnknown(m)
}

var xxx_messageInfo_RouteEngine proto.InternalMessageInfo

func (m *RouteEngine) GetSlot() uint32 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *RouteEngine) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *RouteEngine) GetMgmtIf() string {
	if m != nil {
		return m.MgmtIf
	}
	return ""
}

func (m *RouteEngine) GetOperation() proto_common.Operation {
	if m != nil {
		return m.Operation
	}
	return proto_common.Operation_CREATE
}

func init() {
	proto.RegisterType((*DEVICE)(nil), "DEVICE")
	proto.RegisterType((*Linecard)(nil), "Linecard")
	proto.RegisterType((*RouteEngine)(nil), "RouteEngine")
}

func init() {
	proto.RegisterFile("proto_inventory/proto_inventory.proto", fileDescriptor_fd0852c13d032805)
}

var fileDescriptor_fd0852c13d032805 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0x86, 0xf1, 0xc6, 0xf1, 0x3a, 0xca, 0xee, 0x1e, 0xc4, 0xd2, 0x8a, 0x5c, 0x6a, 0x02, 0x01,
	0x9f, 0x6c, 0x9a, 0xbc, 0x41, 0xdb, 0x1c, 0x02, 0x85, 0x82, 0x28, 0x3d, 0xf4, 0x12, 0x1c, 0x6b,
	0x92, 0x0a, 0x2c, 0x8d, 0x51, 0x14, 0xd3, 0xf4, 0x99, 0xfa, 0x32, 0x7d, 0xa3, 0x62, 0xd9, 0x71,
	0xd3, 0x1c, 0x0a, 0xb9, 0xcd, 0xf7, 0xcf, 0x68, 0x66, 0x7e, 0x31, 0x64, 0x52, 0x1a, 0xb4, 0xb8,
	0x94, 0xba, 0x02, 0x6d, 0xd1, 0xec, 0xd3, 0x13, 0x4e, 0x1c, 0x8f, 0xae, 0x1a, 0x39, 0x47, 0xa5,
	0x50, 0xa7, 0xc7, 0xd0, 0x14, 0x8c, 0x3f, 0x3c, 0x12, 0xdc, 0xcd, 0x9f, 0x16, 0xb7, 0x73, 0x4a,
	0x89, 0xaf, 0x33, 0x05, 0xcc, 0x8b, 0xbc, 0x78, 0xc0, 0x5d, 0x4c, 0x2f, 0x48, 0x50, 0x81, 0x16,
	0x68, 0xd8, 0x2f, 0xa7, 0xb6, 0x44, 0x47, 0x24, 0x2c, 0x8b, 0xcc, 0xae, 0xd1, 0x28, 0xd6, 0x73,
	0x99, 0x8e, 0x69, 0x42, 0x86, 0x06, 0x77, 0x16, 0x40, 0x6f, 0xa4, 0x06, 0xe6, 0x47, 0xbd, 0x78,
	0x38, 0xfd, 0x93, 0xf0, 0x5a, 0x9b, 0x3b, 0x8d, 0x1f, 0x17, 0xd0, 0x09, 0x09, 0x0b, 0xa9, 0x21,
	0xcf, 0x8c, 0x60, 0x7d, 0x57, 0x3c, 0x48, 0xee, 0x5b, 0x81, 0x77, 0x29, 0x1a, 0x93, 0x01, 0x96,
	0x60, 0x32, 0x2b, 0x51, 0xb3, 0x20, 0xf2, 0xe2, 0x7f, 0x53, 0x92, 0x3c, 0x1c, 0x14, 0xfe, 0x95,
	0x1c, 0xbf, 0x7b, 0x24, 0x3c, 0x34, 0xa8, 0x5d, 0x6d, 0x0b, 0xb4, 0xce, 0xd5, 0x5f, 0xee, 0x62,
	0xfa, 0x9f, 0xf4, 0x15, 0x0a, 0x28, 0x5a, 0x53, 0x0d, 0x38, 0x4f, 0x68, 0xec, 0xe3, 0xbe, 0x84,
	0xce, 0x53, 0xcb, 0xf5, 0x8b, 0x3a, 0xde, 0x32, 0xdf, 0xb5, 0x69, 0xa0, 0x7e, 0x01, 0xaf, 0x65,
	0x21, 0x73, 0x69, 0x59, 0x3f, 0xf2, 0xe2, 0x90, 0x77, 0x7c, 0xc6, 0xba, 0x6f, 0x64, 0x78, 0xf4,
	0x37, 0x67, 0x2c, 0x7c, 0x49, 0x7e, 0xab, 0x8d, 0xb2, 0x4b, 0xb9, 0x6e, 0xf7, 0x0d, 0x6a, 0x5c,
	0xac, 0xbf, 0xcf, 0xf6, 0x7f, 0x98, 0x7d, 0x33, 0x7b, 0xbe, 0xde, 0x48, 0xfb, 0xb2, 0x5b, 0x25,
	0x39, 0xaa, 0x54, 0x64, 0x15, 0x08, 0xb4, 0x02, 0xaa, 0xe6, 0x54, 0x52, 0x10, 0x42, 0xc2, 0xe9,
	0x69, 0xad, 0x02, 0x27, 0xcc, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x08, 0x5f, 0xa9, 0x38, 0x84,
	0x02, 0x00, 0x00,
}
