// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_redis_streams.proto

package args

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

type OffsetStart int32

const (
	OffsetStart_LATEST OffsetStart = 0
	OffsetStart_OLDEST OffsetStart = 1
)

var OffsetStart_name = map[int32]string{
	0: "LATEST",
	1: "OLDEST",
}

var OffsetStart_value = map[string]int32{
	"LATEST": 0,
	"OLDEST": 1,
}

func (x OffsetStart) String() string {
	return proto.EnumName(OffsetStart_name, int32(x))
}

func (OffsetStart) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ec80c3317f49688, []int{0}
}

type RedisStreamsConn struct {
	// @gotags: kong:"help='Address of redis server',default=localhost:6379,required,env='PLUMBER_RELAY_REDIS_STREAMS_ADDRESS'"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" kong:"help='Address of redis server',default=localhost:6379,required,env='PLUMBER_RELAY_REDIS_STREAMS_ADDRESS'"`
	// @gotags: kong:"help='Username (redis >= v6.0.0)',env='PLUMBER_RELAY_REDIS_STREAMS_USERNAME'"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" kong:"help='Username (redis >= v6.0.0)',env='PLUMBER_RELAY_REDIS_STREAMS_USERNAME'"`
	// @gotags: kong:"help='Password (redis >= v6.0.0)',env='PLUMBER_RELAY_REDIS_STREAMS_PASSWORD'"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" kong:"help='Password (redis >= v6.0.0)',env='PLUMBER_RELAY_REDIS_STREAMS_PASSWORD'"`
	// @gotags: kong:"help='Database (0-16)',env='PLUMBER_RELAY_REDIS_PUBSUB_DATABASE'"
	Database             uint32   `protobuf:"varint,4,opt,name=database,proto3" json:"database,omitempty" kong:"help='Database (0-16)',env='PLUMBER_RELAY_REDIS_PUBSUB_DATABASE'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedisStreamsConn) Reset()         { *m = RedisStreamsConn{} }
func (m *RedisStreamsConn) String() string { return proto.CompactTextString(m) }
func (*RedisStreamsConn) ProtoMessage()    {}
func (*RedisStreamsConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec80c3317f49688, []int{0}
}

func (m *RedisStreamsConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisStreamsConn.Unmarshal(m, b)
}
func (m *RedisStreamsConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisStreamsConn.Marshal(b, m, deterministic)
}
func (m *RedisStreamsConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisStreamsConn.Merge(m, src)
}
func (m *RedisStreamsConn) XXX_Size() int {
	return xxx_messageInfo_RedisStreamsConn.Size(m)
}
func (m *RedisStreamsConn) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisStreamsConn.DiscardUnknown(m)
}

var xxx_messageInfo_RedisStreamsConn proto.InternalMessageInfo

func (m *RedisStreamsConn) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RedisStreamsConn) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RedisStreamsConn) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RedisStreamsConn) GetDatabase() uint32 {
	if m != nil {
		return m.Database
	}
	return 0
}

type CreateConsumerConfig struct {
	// @gotags: kong:"help='Create the streams if creating a new consumer group',env='PLUMBER_RELAY_REDIS_STREAMS_CREATE_STREAMS'"
	CreateStreams bool `protobuf:"varint,1,opt,name=create_streams,json=createStreams,proto3" json:"create_streams,omitempty" kong:"help='Create the streams if creating a new consumer group',env='PLUMBER_RELAY_REDIS_STREAMS_CREATE_STREAMS'"`
	// @gotags: kong:"help='Recreate this consumer group if it does not exist',env='PLUMBER_RELAY_REDIS_STREAMS_RECREATE_CONSUMER_GROUP'"
	RecreateConsumerGroup bool `protobuf:"varint,2,opt,name=recreate_consumer_group,json=recreateConsumerGroup,proto3" json:"recreate_consumer_group,omitempty" kong:"help='Recreate this consumer group if it does not exist',env='PLUMBER_RELAY_REDIS_STREAMS_RECREATE_CONSUMER_GROUP'"`
	// @gotags: kong:"help='What offset to start reading at (options: latest oldest)',default=latest,required,env='PLUMBER_RELAY_REDIS_STREAMS_START_ID',type=pbenum,pbenum_lowercase"
	OffsetStart          OffsetStart `protobuf:"varint,3,opt,name=offset_start,json=offsetStart,proto3,enum=protos.args.OffsetStart" json:"offset_start,omitempty" kong:"help='What offset to start reading at (options: latest oldest)',default=latest,required,env='PLUMBER_RELAY_REDIS_STREAMS_START_ID',type=pbenum,pbenum_lowercase"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateConsumerConfig) Reset()         { *m = CreateConsumerConfig{} }
func (m *CreateConsumerConfig) String() string { return proto.CompactTextString(m) }
func (*CreateConsumerConfig) ProtoMessage()    {}
func (*CreateConsumerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec80c3317f49688, []int{1}
}

func (m *CreateConsumerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConsumerConfig.Unmarshal(m, b)
}
func (m *CreateConsumerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConsumerConfig.Marshal(b, m, deterministic)
}
func (m *CreateConsumerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConsumerConfig.Merge(m, src)
}
func (m *CreateConsumerConfig) XXX_Size() int {
	return xxx_messageInfo_CreateConsumerConfig.Size(m)
}
func (m *CreateConsumerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConsumerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConsumerConfig proto.InternalMessageInfo

func (m *CreateConsumerConfig) GetCreateStreams() bool {
	if m != nil {
		return m.CreateStreams
	}
	return false
}

func (m *CreateConsumerConfig) GetRecreateConsumerGroup() bool {
	if m != nil {
		return m.RecreateConsumerGroup
	}
	return false
}

func (m *CreateConsumerConfig) GetOffsetStart() OffsetStart {
	if m != nil {
		return m.OffsetStart
	}
	return OffsetStart_LATEST
}

type RedisStreamsReadArgs struct {
	// @gotags: kong:"help='Streams to read from',required,env='PLUMBER_RELAY_REDIS_STREAMS_STREAMS'"
	Streams []string `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams,omitempty" kong:"help='Streams to read from',required,env='PLUMBER_RELAY_REDIS_STREAMS_STREAMS'"`
	// @gotags: kong:"help='Consumer group name',env='PLUMBER_RELAY_REDIS_STREAMS_CONSUMER_GROUP',default=plumber"
	ConsumerGroup string `protobuf:"bytes,2,opt,name=consumer_group,json=consumerGroup,proto3" json:"consumer_group,omitempty" kong:"help='Consumer group name',env='PLUMBER_RELAY_REDIS_STREAMS_CONSUMER_GROUP',default=plumber"`
	// @gotags: kong:"help='Consumer name',env='PLUMBER_RELAY_REDIS_STREAMS_CONSUMER_NAME',default=plumber-consumer-1"
	ConsumerName string `protobuf:"bytes,3,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty" kong:"help='Consumer name',env='PLUMBER_RELAY_REDIS_STREAMS_CONSUMER_NAME',default=plumber-consumer-1"`
	// @gotags: kong:"help='Number of records to read from stream(s) per read',env='PLUMBER_RELAY_REDIS_STREAMS_COUNT',default=10"
	Count uint32 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty" kong:"help='Number of records to read from stream(s) per read',env='PLUMBER_RELAY_REDIS_STREAMS_COUNT',default=10"`
	// @gotags: kong:"embed"
	CreateConsumerConfig *CreateConsumerConfig `protobuf:"bytes,5,opt,name=create_consumer_config,json=createConsumerConfig,proto3" json:"create_consumer_config,omitempty" kong:"embed"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RedisStreamsReadArgs) Reset()         { *m = RedisStreamsReadArgs{} }
func (m *RedisStreamsReadArgs) String() string { return proto.CompactTextString(m) }
func (*RedisStreamsReadArgs) ProtoMessage()    {}
func (*RedisStreamsReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec80c3317f49688, []int{2}
}

func (m *RedisStreamsReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisStreamsReadArgs.Unmarshal(m, b)
}
func (m *RedisStreamsReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisStreamsReadArgs.Marshal(b, m, deterministic)
}
func (m *RedisStreamsReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisStreamsReadArgs.Merge(m, src)
}
func (m *RedisStreamsReadArgs) XXX_Size() int {
	return xxx_messageInfo_RedisStreamsReadArgs.Size(m)
}
func (m *RedisStreamsReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisStreamsReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RedisStreamsReadArgs proto.InternalMessageInfo

func (m *RedisStreamsReadArgs) GetStreams() []string {
	if m != nil {
		return m.Streams
	}
	return nil
}

func (m *RedisStreamsReadArgs) GetConsumerGroup() string {
	if m != nil {
		return m.ConsumerGroup
	}
	return ""
}

func (m *RedisStreamsReadArgs) GetConsumerName() string {
	if m != nil {
		return m.ConsumerName
	}
	return ""
}

func (m *RedisStreamsReadArgs) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *RedisStreamsReadArgs) GetCreateConsumerConfig() *CreateConsumerConfig {
	if m != nil {
		return m.CreateConsumerConfig
	}
	return nil
}

type RedisStreamsWriteArgs struct {
	// @gotags: kong:"help='What redis ID to use for input data (* = auto-generate)',default='*'"
	WriteId string `protobuf:"bytes,1,opt,name=write_id,json=writeId,proto3" json:"write_id,omitempty" kong:"help='What redis ID to use for input data (* = auto-generate)',default='*'"`
	// @gotags: kong:"help='Streams to write to'"
	Streams []string `protobuf:"bytes,2,rep,name=streams,proto3" json:"streams,omitempty" kong:"help='Streams to write to'"`
	// @gotags: kong:"help='Key name to write input data to'"
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty" kong:"help='Key name to write input data to'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedisStreamsWriteArgs) Reset()         { *m = RedisStreamsWriteArgs{} }
func (m *RedisStreamsWriteArgs) String() string { return proto.CompactTextString(m) }
func (*RedisStreamsWriteArgs) ProtoMessage()    {}
func (*RedisStreamsWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ec80c3317f49688, []int{3}
}

func (m *RedisStreamsWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisStreamsWriteArgs.Unmarshal(m, b)
}
func (m *RedisStreamsWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisStreamsWriteArgs.Marshal(b, m, deterministic)
}
func (m *RedisStreamsWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisStreamsWriteArgs.Merge(m, src)
}
func (m *RedisStreamsWriteArgs) XXX_Size() int {
	return xxx_messageInfo_RedisStreamsWriteArgs.Size(m)
}
func (m *RedisStreamsWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisStreamsWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RedisStreamsWriteArgs proto.InternalMessageInfo

func (m *RedisStreamsWriteArgs) GetWriteId() string {
	if m != nil {
		return m.WriteId
	}
	return ""
}

func (m *RedisStreamsWriteArgs) GetStreams() []string {
	if m != nil {
		return m.Streams
	}
	return nil
}

func (m *RedisStreamsWriteArgs) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterEnum("protos.args.OffsetStart", OffsetStart_name, OffsetStart_value)
	proto.RegisterType((*RedisStreamsConn)(nil), "protos.args.RedisStreamsConn")
	proto.RegisterType((*CreateConsumerConfig)(nil), "protos.args.CreateConsumerConfig")
	proto.RegisterType((*RedisStreamsReadArgs)(nil), "protos.args.RedisStreamsReadArgs")
	proto.RegisterType((*RedisStreamsWriteArgs)(nil), "protos.args.RedisStreamsWriteArgs")
}

func init() { proto.RegisterFile("ps_args_redis_streams.proto", fileDescriptor_7ec80c3317f49688) }

var fileDescriptor_7ec80c3317f49688 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x6b, 0xd4, 0x40,
	0x14, 0x85, 0x4d, 0xd7, 0xd6, 0xed, 0xdd, 0x6e, 0x59, 0xc2, 0x56, 0xa3, 0xbe, 0xac, 0x2b, 0x85,
	0x45, 0x30, 0x81, 0x0a, 0x82, 0xf4, 0xa9, 0x46, 0x11, 0xa1, 0x58, 0x98, 0x2d, 0x14, 0x7c, 0x30,
	0x4c, 0x66, 0x66, 0xb3, 0xc1, 0x26, 0x13, 0xee, 0x9d, 0x50, 0x7c, 0xf3, 0x37, 0xf9, 0xb7, 0xfc,
	0x13, 0x32, 0x93, 0x64, 0x9b, 0xb5, 0xfb, 0x94, 0x7b, 0xee, 0xb9, 0x0c, 0xe7, 0x7c, 0x04, 0x5e,
	0x56, 0x94, 0x70, 0xcc, 0x28, 0x41, 0x25, 0x73, 0x4a, 0xc8, 0xa0, 0xe2, 0x05, 0x85, 0x15, 0x6a,
	0xa3, 0xfd, 0x91, 0xfb, 0x50, 0x68, 0x0f, 0xe6, 0xbf, 0x3d, 0x98, 0x30, 0x7b, 0xb4, 0x6c, 0x6e,
	0x62, 0x5d, 0x96, 0x7e, 0x00, 0x4f, 0xb8, 0x94, 0xa8, 0x88, 0x02, 0x6f, 0xe6, 0x2d, 0x0e, 0x59,
	0x27, 0xfd, 0x17, 0x30, 0xac, 0x49, 0x61, 0xc9, 0x0b, 0x15, 0xec, 0x39, 0x6b, 0xa3, 0xad, 0x57,
	0x71, 0xa2, 0x3b, 0x8d, 0x32, 0x18, 0x34, 0x5e, 0xa7, 0xad, 0x27, 0xb9, 0xe1, 0x29, 0x27, 0x15,
	0x3c, 0x9e, 0x79, 0x8b, 0x31, 0xdb, 0xe8, 0xf9, 0x1f, 0x0f, 0xa6, 0x31, 0x2a, 0x6e, 0x54, 0xac,
	0x4b, 0xaa, 0x0b, 0x85, 0xb1, 0x2e, 0x57, 0x79, 0xe6, 0x9f, 0xc2, 0xb1, 0x70, 0xfb, 0xae, 0x80,
	0x4b, 0x33, 0x64, 0xe3, 0x66, 0xdb, 0x26, 0xf6, 0xdf, 0xc3, 0x33, 0x54, 0xed, 0xa1, 0x68, 0x5f,
	0x48, 0x32, 0xd4, 0x75, 0xe5, 0x22, 0x0e, 0xd9, 0x49, 0x67, 0x77, 0xef, 0x7f, 0xb1, 0xa6, 0x7f,
	0x0e, 0x47, 0x7a, 0xb5, 0x22, 0x65, 0x12, 0x32, 0x1c, 0x8d, 0xcb, 0x7c, 0x7c, 0x16, 0x84, 0x3d,
	0x3c, 0xe1, 0x95, 0x3b, 0x58, 0x5a, 0x9f, 0x8d, 0xf4, 0xbd, 0x98, 0xff, 0xf5, 0x60, 0xda, 0xe7,
	0xc6, 0x14, 0x97, 0x17, 0x98, 0x91, 0x65, 0x77, 0x9f, 0x76, 0x60, 0xd9, 0xb5, 0xd2, 0xd5, 0x79,
	0x18, 0xef, 0x90, 0x8d, 0xc5, 0x56, 0xac, 0xd7, 0xb0, 0x59, 0x24, 0x8e, 0x73, 0xc3, 0xf2, 0xa8,
	0x5b, 0x7e, 0xb3, 0xac, 0xa7, 0xb0, 0x2f, 0x74, 0x5d, 0x9a, 0x16, 0x66, 0x23, 0xfc, 0x1b, 0x78,
	0xfa, 0x3f, 0x07, 0xe1, 0x50, 0x06, 0xfb, 0x33, 0x6f, 0x31, 0x3a, 0x7b, 0xb5, 0xd5, 0x6d, 0x17,
	0x73, 0x36, 0x15, 0x3b, 0xb6, 0xf3, 0x1f, 0x70, 0xd2, 0x2f, 0x7b, 0x83, 0xb9, 0x51, 0xae, 0xed,
	0x73, 0x18, 0xde, 0x59, 0x91, 0xe4, 0xb2, 0xfb, 0x55, 0x9c, 0xfe, 0x2a, 0xfb, 0x20, 0xf6, 0xb6,
	0x41, 0x4c, 0x60, 0xf0, 0x53, 0xfd, 0x6a, 0x7b, 0xd9, 0xf1, 0xcd, 0x29, 0x8c, 0x7a, 0xa4, 0x7d,
	0x80, 0x83, 0xcb, 0x8b, 0xeb, 0xcf, 0xcb, 0xeb, 0xc9, 0x23, 0x3b, 0x5f, 0x5d, 0x7e, 0xb2, 0xb3,
	0xf7, 0xf1, 0xfc, 0xfb, 0x87, 0x2c, 0x37, 0xeb, 0x3a, 0x0d, 0x85, 0x2e, 0xa2, 0x94, 0x1b, 0xb1,
	0x16, 0x1a, 0xab, 0xa8, 0xba, 0xad, 0x8b, 0x54, 0xe1, 0x5b, 0x12, 0x6b, 0x55, 0x70, 0x8a, 0xd2,
	0x3a, 0xbf, 0x95, 0x51, 0xa6, 0xa3, 0xa6, 0x6e, 0x64, 0xeb, 0xa6, 0x07, 0x4e, 0xbc, 0xfb, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x8e, 0xba, 0x17, 0xde, 0x1c, 0x03, 0x00, 0x00,
}