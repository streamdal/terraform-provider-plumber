// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/ps_records_nats_jetstream.proto

package records

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

type NatsJetstream struct {
	Stream               string   `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ConsumerName         string   `protobuf:"bytes,3,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty"`
	Sequence             int64    `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NatsJetstream) Reset()         { *m = NatsJetstream{} }
func (m *NatsJetstream) String() string { return proto.CompactTextString(m) }
func (*NatsJetstream) ProtoMessage()    {}
func (*NatsJetstream) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fb39dca56fdde41, []int{0}
}

func (m *NatsJetstream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NatsJetstream.Unmarshal(m, b)
}
func (m *NatsJetstream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NatsJetstream.Marshal(b, m, deterministic)
}
func (m *NatsJetstream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NatsJetstream.Merge(m, src)
}
func (m *NatsJetstream) XXX_Size() int {
	return xxx_messageInfo_NatsJetstream.Size(m)
}
func (m *NatsJetstream) XXX_DiscardUnknown() {
	xxx_messageInfo_NatsJetstream.DiscardUnknown(m)
}

var xxx_messageInfo_NatsJetstream proto.InternalMessageInfo

func (m *NatsJetstream) GetStream() string {
	if m != nil {
		return m.Stream
	}
	return ""
}

func (m *NatsJetstream) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *NatsJetstream) GetConsumerName() string {
	if m != nil {
		return m.ConsumerName
	}
	return ""
}

func (m *NatsJetstream) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*NatsJetstream)(nil), "protos.records.NatsJetstream")
}

func init() {
	proto.RegisterFile("records/ps_records_nats_jetstream.proto", fileDescriptor_2fb39dca56fdde41)
}

var fileDescriptor_2fb39dca56fdde41 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xbf, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x65, 0x0a, 0x15, 0x58, 0x2d, 0x83, 0x85, 0x90, 0xc5, 0x14, 0xc1, 0x40, 0x16, 0xe2,
	0x81, 0x15, 0x31, 0x30, 0x32, 0x74, 0xc8, 0xc8, 0x62, 0x9d, 0xdd, 0x53, 0x13, 0x14, 0xff, 0xc0,
	0x67, 0x33, 0xf2, 0xb7, 0x23, 0x25, 0x4e, 0xa7, 0x7b, 0xdf, 0xe9, 0x3d, 0xe9, 0xe3, 0xcf, 0x09,
	0x6d, 0x48, 0x47, 0x52, 0x91, 0x74, 0x8d, 0xda, 0x43, 0x26, 0xfd, 0x8d, 0x99, 0x72, 0x42, 0x70,
	0x5d, 0x4c, 0x21, 0x07, 0x71, 0x3b, 0x1f, 0xea, 0x6a, 0xe9, 0xf1, 0x8f, 0xef, 0x0f, 0x90, 0xe9,
	0x73, 0xad, 0x89, 0x7b, 0xbe, 0x5d, 0x92, 0x64, 0x0d, 0x6b, 0x6f, 0xfa, 0x4a, 0xe2, 0x8e, 0x5f,
	0xfd, 0xc2, 0x54, 0x50, 0x5e, 0x34, 0xac, 0xdd, 0xf5, 0x0b, 0x88, 0x27, 0xbe, 0xb7, 0xc1, 0x53,
	0x71, 0x98, 0xb4, 0x07, 0x87, 0x72, 0x33, 0x8f, 0x76, 0xeb, 0xf3, 0x00, 0x0e, 0xc5, 0x03, 0xbf,
	0x26, 0xfc, 0x29, 0xe8, 0x2d, 0xca, 0xcb, 0x86, 0xb5, 0x9b, 0xfe, 0xcc, 0x1f, 0xef, 0x5f, 0x6f,
	0xa7, 0x31, 0x0f, 0xc5, 0x74, 0x36, 0x38, 0x65, 0x20, 0xdb, 0xc1, 0x86, 0x14, 0x55, 0x9c, 0x8a,
	0x33, 0x98, 0x5e, 0xc8, 0x0e, 0xe8, 0x80, 0x94, 0x29, 0xe3, 0x74, 0x54, 0xa7, 0xa0, 0x16, 0x7f,
	0x55, 0xfd, 0xcd, 0x76, 0xe6, 0xd7, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xeb, 0x84, 0x09,
	0x01, 0x01, 0x00, 0x00,
}
