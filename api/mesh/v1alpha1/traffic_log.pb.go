// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/traffic_log.proto

package v1alpha1

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

// TrafficLog defines traffic log for selected dataplanes.
type TrafficLog struct {
	// List of selectors to match dataplanes.
	Selectors []*Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// Configuration of the logging.
	Conf                 *TrafficLog_Conf `protobuf:"bytes,3,opt,name=conf,proto3" json:"conf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TrafficLog) Reset()         { *m = TrafficLog{} }
func (m *TrafficLog) String() string { return proto.CompactTextString(m) }
func (*TrafficLog) ProtoMessage()    {}
func (*TrafficLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c4f4c9c894eeed, []int{0}
}

func (m *TrafficLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficLog.Unmarshal(m, b)
}
func (m *TrafficLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficLog.Marshal(b, m, deterministic)
}
func (m *TrafficLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficLog.Merge(m, src)
}
func (m *TrafficLog) XXX_Size() int {
	return xxx_messageInfo_TrafficLog.Size(m)
}
func (m *TrafficLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficLog.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficLog proto.InternalMessageInfo

func (m *TrafficLog) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *TrafficLog) GetConf() *TrafficLog_Conf {
	if m != nil {
		return m.Conf
	}
	return nil
}

// Configuration defines settings of the logging.
type TrafficLog_Conf struct {
	// Backend defined in the Mesh entity.
	Backend              string   `protobuf:"bytes,1,opt,name=backend,proto3" json:"backend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrafficLog_Conf) Reset()         { *m = TrafficLog_Conf{} }
func (m *TrafficLog_Conf) String() string { return proto.CompactTextString(m) }
func (*TrafficLog_Conf) ProtoMessage()    {}
func (*TrafficLog_Conf) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c4f4c9c894eeed, []int{0, 0}
}

func (m *TrafficLog_Conf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficLog_Conf.Unmarshal(m, b)
}
func (m *TrafficLog_Conf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficLog_Conf.Marshal(b, m, deterministic)
}
func (m *TrafficLog_Conf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficLog_Conf.Merge(m, src)
}
func (m *TrafficLog_Conf) XXX_Size() int {
	return xxx_messageInfo_TrafficLog_Conf.Size(m)
}
func (m *TrafficLog_Conf) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficLog_Conf.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficLog_Conf proto.InternalMessageInfo

func (m *TrafficLog_Conf) GetBackend() string {
	if m != nil {
		return m.Backend
	}
	return ""
}

func init() {
	proto.RegisterType((*TrafficLog)(nil), "kuma.mesh.v1alpha1.TrafficLog")
	proto.RegisterType((*TrafficLog_Conf)(nil), "kuma.mesh.v1alpha1.TrafficLog.Conf")
}

func init() { proto.RegisterFile("mesh/v1alpha1/traffic_log.proto", fileDescriptor_47c4f4c9c894eeed) }

var fileDescriptor_47c4f4c9c894eeed = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0x4d, 0x2d, 0xce,
	0xd0, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x2f, 0x29, 0x4a, 0x4c, 0x4b, 0xcb,
	0x4c, 0x8e, 0xcf, 0xc9, 0x4f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xca, 0x2e, 0xcd,
	0x4d, 0xd4, 0x03, 0xa9, 0xd2, 0x83, 0xa9, 0x92, 0x92, 0x41, 0xd5, 0x54, 0x9c, 0x9a, 0x93, 0x9a,
	0x5c, 0x92, 0x5f, 0x04, 0xd1, 0xa1, 0xb4, 0x98, 0x91, 0x8b, 0x2b, 0x04, 0x62, 0x8e, 0x4f, 0x7e,
	0xba, 0x90, 0x15, 0x17, 0x27, 0x4c, 0x41, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x8c,
	0x1e, 0xa6, 0xa1, 0x7a, 0xc1, 0x50, 0x45, 0x41, 0x08, 0xe5, 0x42, 0xe6, 0x5c, 0x2c, 0xc9, 0xf9,
	0x79, 0x69, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xca, 0xd8, 0xb4, 0x21, 0x6c, 0xd2, 0x73,
	0xce, 0xcf, 0x4b, 0x0b, 0x02, 0x6b, 0x90, 0x52, 0xe0, 0x62, 0x01, 0xf1, 0x84, 0x24, 0xb8, 0xd8,
	0x93, 0x12, 0x93, 0xb3, 0x53, 0xf3, 0x52, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c,
	0x27, 0xae, 0x28, 0x0e, 0x98, 0x19, 0x49, 0x6c, 0x60, 0x87, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xf4, 0xbf, 0x14, 0xe3, 0x0d, 0x01, 0x00, 0x00,
}
