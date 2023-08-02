// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/ps_records_base.proto

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

// Returned for read requests (server & cli)
type ReadRecord struct {
	// Unique id automatically created by plumber
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// Plumber counts the number of messages it reads; this number represents
	// the message number (useful for CLI).
	Num int64 `protobuf:"varint,3,opt,name=num,proto3" json:"num,omitempty"`
	// Metadata may contain properties that cannot be found in the Raw message.
	// For example: read lag in Kafka.
	//
	// Metadata may also contain data such as "count" which is an incremental
	// number that plumber assigns to each message it receives. This is used
	// with read via CLI functionality to allow the user to quickly discern
	// whether this is message #1 or #500, etc.
	Metadata map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// UTC unix timestamp of when plumber received the message; a backend record
	// entry might have its own timestamp as well. This should be seconds.
	ReceivedAtUnixTsUtc int64 `protobuf:"varint,6,opt,name=received_at_unix_ts_utc,json=receivedAtUnixTsUtc,proto3" json:"received_at_unix_ts_utc,omitempty"`
	// Set _outside_ the backend; will contain the final value, regardless of
	// whether decoding options were specified for a read.
	// _This_ is what both CLI and desktop should display for the payload.
	Payload []byte `protobuf:"bytes,99,opt,name=payload,proto3" json:"payload,omitempty"`
	// Types that are valid to be assigned to Record:
	//	*ReadRecord_Kafka
	//	*ReadRecord_Rabbit
	//	*ReadRecord_Activemq
	//	*ReadRecord_AwsSqs
	//	*ReadRecord_AzureEventHub
	//	*ReadRecord_AzureServiceBus
	//	*ReadRecord_GcpPubsub
	//	*ReadRecord_Kubemq
	//	*ReadRecord_Mongo
	//	*ReadRecord_Mqtt
	//	*ReadRecord_Nats
	//	*ReadRecord_NatsStreaming
	//	*ReadRecord_Nsq
	//	*ReadRecord_Postgres
	//	*ReadRecord_Pulsar
	//	*ReadRecord_RabbitStreams
	//	*ReadRecord_RedisPubsub
	//	*ReadRecord_RedisStreams
	//	*ReadRecord_NatsJetstream
	//	*ReadRecord_AwsKinesis
	//	*ReadRecord_Memphis
	Record isReadRecord_Record `protobuf_oneof:"Record"`
	// Original backend message (encoded with gob, ie. *skafka.Message, etc.).
	// In most cases, you should use the oneof record instead of the raw message.
	XRaw []byte `protobuf:"bytes,1000,opt,name=_raw,json=Raw,proto3" json:"_raw,omitempty"`
	// Identifies which plumber instance received the event (set outside the backend)
	XPlumberId           string   `protobuf:"bytes,1001,opt,name=_plumber_id,json=PlumberId,proto3" json:"_plumber_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRecord) Reset()         { *m = ReadRecord{} }
func (m *ReadRecord) String() string { return proto.CompactTextString(m) }
func (*ReadRecord) ProtoMessage()    {}
func (*ReadRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_79a95b759e1a93f6, []int{0}
}

func (m *ReadRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRecord.Unmarshal(m, b)
}
func (m *ReadRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRecord.Marshal(b, m, deterministic)
}
func (m *ReadRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRecord.Merge(m, src)
}
func (m *ReadRecord) XXX_Size() int {
	return xxx_messageInfo_ReadRecord.Size(m)
}
func (m *ReadRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRecord proto.InternalMessageInfo

func (m *ReadRecord) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *ReadRecord) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *ReadRecord) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ReadRecord) GetReceivedAtUnixTsUtc() int64 {
	if m != nil {
		return m.ReceivedAtUnixTsUtc
	}
	return 0
}

func (m *ReadRecord) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type isReadRecord_Record interface {
	isReadRecord_Record()
}

type ReadRecord_Kafka struct {
	Kafka *Kafka `protobuf:"bytes,100,opt,name=kafka,proto3,oneof"`
}

type ReadRecord_Rabbit struct {
	Rabbit *Rabbit `protobuf:"bytes,101,opt,name=rabbit,proto3,oneof"`
}

type ReadRecord_Activemq struct {
	Activemq *ActiveMQ `protobuf:"bytes,102,opt,name=activemq,proto3,oneof"`
}

type ReadRecord_AwsSqs struct {
	AwsSqs *AWSSQS `protobuf:"bytes,103,opt,name=aws_sqs,json=awsSqs,proto3,oneof"`
}

type ReadRecord_AzureEventHub struct {
	AzureEventHub *AzureEventHub `protobuf:"bytes,104,opt,name=azure_event_hub,json=azureEventHub,proto3,oneof"`
}

type ReadRecord_AzureServiceBus struct {
	AzureServiceBus *AzureServiceBus `protobuf:"bytes,105,opt,name=azure_service_bus,json=azureServiceBus,proto3,oneof"`
}

type ReadRecord_GcpPubsub struct {
	GcpPubsub *GCPPubSub `protobuf:"bytes,106,opt,name=gcp_pubsub,json=gcpPubsub,proto3,oneof"`
}

type ReadRecord_Kubemq struct {
	Kubemq *KubeMQ `protobuf:"bytes,107,opt,name=kubemq,proto3,oneof"`
}

type ReadRecord_Mongo struct {
	Mongo *Mongo `protobuf:"bytes,108,opt,name=mongo,proto3,oneof"`
}

type ReadRecord_Mqtt struct {
	Mqtt *MQTT `protobuf:"bytes,109,opt,name=mqtt,proto3,oneof"`
}

type ReadRecord_Nats struct {
	Nats *Nats `protobuf:"bytes,110,opt,name=nats,proto3,oneof"`
}

type ReadRecord_NatsStreaming struct {
	NatsStreaming *NatsStreaming `protobuf:"bytes,111,opt,name=nats_streaming,json=natsStreaming,proto3,oneof"`
}

type ReadRecord_Nsq struct {
	Nsq *NSQ `protobuf:"bytes,112,opt,name=nsq,proto3,oneof"`
}

type ReadRecord_Postgres struct {
	Postgres *Postgres `protobuf:"bytes,113,opt,name=postgres,proto3,oneof"`
}

type ReadRecord_Pulsar struct {
	Pulsar *Pulsar `protobuf:"bytes,114,opt,name=pulsar,proto3,oneof"`
}

type ReadRecord_RabbitStreams struct {
	RabbitStreams *RabbitStreams `protobuf:"bytes,115,opt,name=rabbit_streams,json=rabbitStreams,proto3,oneof"`
}

type ReadRecord_RedisPubsub struct {
	RedisPubsub *RedisPubsub `protobuf:"bytes,116,opt,name=redis_pubsub,json=redisPubsub,proto3,oneof"`
}

type ReadRecord_RedisStreams struct {
	RedisStreams *RedisStreams `protobuf:"bytes,117,opt,name=redis_streams,json=redisStreams,proto3,oneof"`
}

type ReadRecord_NatsJetstream struct {
	NatsJetstream *NatsJetstream `protobuf:"bytes,118,opt,name=nats_jetstream,json=natsJetstream,proto3,oneof"`
}

type ReadRecord_AwsKinesis struct {
	AwsKinesis *AWSKinesis `protobuf:"bytes,119,opt,name=aws_kinesis,json=awsKinesis,proto3,oneof"`
}

type ReadRecord_Memphis struct {
	Memphis *Memphis `protobuf:"bytes,120,opt,name=memphis,proto3,oneof"`
}

func (*ReadRecord_Kafka) isReadRecord_Record() {}

func (*ReadRecord_Rabbit) isReadRecord_Record() {}

func (*ReadRecord_Activemq) isReadRecord_Record() {}

func (*ReadRecord_AwsSqs) isReadRecord_Record() {}

func (*ReadRecord_AzureEventHub) isReadRecord_Record() {}

func (*ReadRecord_AzureServiceBus) isReadRecord_Record() {}

func (*ReadRecord_GcpPubsub) isReadRecord_Record() {}

func (*ReadRecord_Kubemq) isReadRecord_Record() {}

func (*ReadRecord_Mongo) isReadRecord_Record() {}

func (*ReadRecord_Mqtt) isReadRecord_Record() {}

func (*ReadRecord_Nats) isReadRecord_Record() {}

func (*ReadRecord_NatsStreaming) isReadRecord_Record() {}

func (*ReadRecord_Nsq) isReadRecord_Record() {}

func (*ReadRecord_Postgres) isReadRecord_Record() {}

func (*ReadRecord_Pulsar) isReadRecord_Record() {}

func (*ReadRecord_RabbitStreams) isReadRecord_Record() {}

func (*ReadRecord_RedisPubsub) isReadRecord_Record() {}

func (*ReadRecord_RedisStreams) isReadRecord_Record() {}

func (*ReadRecord_NatsJetstream) isReadRecord_Record() {}

func (*ReadRecord_AwsKinesis) isReadRecord_Record() {}

func (*ReadRecord_Memphis) isReadRecord_Record() {}

func (m *ReadRecord) GetRecord() isReadRecord_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *ReadRecord) GetKafka() *Kafka {
	if x, ok := m.GetRecord().(*ReadRecord_Kafka); ok {
		return x.Kafka
	}
	return nil
}

func (m *ReadRecord) GetRabbit() *Rabbit {
	if x, ok := m.GetRecord().(*ReadRecord_Rabbit); ok {
		return x.Rabbit
	}
	return nil
}

func (m *ReadRecord) GetActivemq() *ActiveMQ {
	if x, ok := m.GetRecord().(*ReadRecord_Activemq); ok {
		return x.Activemq
	}
	return nil
}

func (m *ReadRecord) GetAwsSqs() *AWSSQS {
	if x, ok := m.GetRecord().(*ReadRecord_AwsSqs); ok {
		return x.AwsSqs
	}
	return nil
}

func (m *ReadRecord) GetAzureEventHub() *AzureEventHub {
	if x, ok := m.GetRecord().(*ReadRecord_AzureEventHub); ok {
		return x.AzureEventHub
	}
	return nil
}

func (m *ReadRecord) GetAzureServiceBus() *AzureServiceBus {
	if x, ok := m.GetRecord().(*ReadRecord_AzureServiceBus); ok {
		return x.AzureServiceBus
	}
	return nil
}

func (m *ReadRecord) GetGcpPubsub() *GCPPubSub {
	if x, ok := m.GetRecord().(*ReadRecord_GcpPubsub); ok {
		return x.GcpPubsub
	}
	return nil
}

func (m *ReadRecord) GetKubemq() *KubeMQ {
	if x, ok := m.GetRecord().(*ReadRecord_Kubemq); ok {
		return x.Kubemq
	}
	return nil
}

func (m *ReadRecord) GetMongo() *Mongo {
	if x, ok := m.GetRecord().(*ReadRecord_Mongo); ok {
		return x.Mongo
	}
	return nil
}

func (m *ReadRecord) GetMqtt() *MQTT {
	if x, ok := m.GetRecord().(*ReadRecord_Mqtt); ok {
		return x.Mqtt
	}
	return nil
}

func (m *ReadRecord) GetNats() *Nats {
	if x, ok := m.GetRecord().(*ReadRecord_Nats); ok {
		return x.Nats
	}
	return nil
}

func (m *ReadRecord) GetNatsStreaming() *NatsStreaming {
	if x, ok := m.GetRecord().(*ReadRecord_NatsStreaming); ok {
		return x.NatsStreaming
	}
	return nil
}

func (m *ReadRecord) GetNsq() *NSQ {
	if x, ok := m.GetRecord().(*ReadRecord_Nsq); ok {
		return x.Nsq
	}
	return nil
}

func (m *ReadRecord) GetPostgres() *Postgres {
	if x, ok := m.GetRecord().(*ReadRecord_Postgres); ok {
		return x.Postgres
	}
	return nil
}

func (m *ReadRecord) GetPulsar() *Pulsar {
	if x, ok := m.GetRecord().(*ReadRecord_Pulsar); ok {
		return x.Pulsar
	}
	return nil
}

func (m *ReadRecord) GetRabbitStreams() *RabbitStreams {
	if x, ok := m.GetRecord().(*ReadRecord_RabbitStreams); ok {
		return x.RabbitStreams
	}
	return nil
}

func (m *ReadRecord) GetRedisPubsub() *RedisPubsub {
	if x, ok := m.GetRecord().(*ReadRecord_RedisPubsub); ok {
		return x.RedisPubsub
	}
	return nil
}

func (m *ReadRecord) GetRedisStreams() *RedisStreams {
	if x, ok := m.GetRecord().(*ReadRecord_RedisStreams); ok {
		return x.RedisStreams
	}
	return nil
}

func (m *ReadRecord) GetNatsJetstream() *NatsJetstream {
	if x, ok := m.GetRecord().(*ReadRecord_NatsJetstream); ok {
		return x.NatsJetstream
	}
	return nil
}

func (m *ReadRecord) GetAwsKinesis() *AWSKinesis {
	if x, ok := m.GetRecord().(*ReadRecord_AwsKinesis); ok {
		return x.AwsKinesis
	}
	return nil
}

func (m *ReadRecord) GetMemphis() *Memphis {
	if x, ok := m.GetRecord().(*ReadRecord_Memphis); ok {
		return x.Memphis
	}
	return nil
}

func (m *ReadRecord) GetXRaw() []byte {
	if m != nil {
		return m.XRaw
	}
	return nil
}

func (m *ReadRecord) GetXPlumberId() string {
	if m != nil {
		return m.XPlumberId
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReadRecord) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReadRecord_Kafka)(nil),
		(*ReadRecord_Rabbit)(nil),
		(*ReadRecord_Activemq)(nil),
		(*ReadRecord_AwsSqs)(nil),
		(*ReadRecord_AzureEventHub)(nil),
		(*ReadRecord_AzureServiceBus)(nil),
		(*ReadRecord_GcpPubsub)(nil),
		(*ReadRecord_Kubemq)(nil),
		(*ReadRecord_Mongo)(nil),
		(*ReadRecord_Mqtt)(nil),
		(*ReadRecord_Nats)(nil),
		(*ReadRecord_NatsStreaming)(nil),
		(*ReadRecord_Nsq)(nil),
		(*ReadRecord_Postgres)(nil),
		(*ReadRecord_Pulsar)(nil),
		(*ReadRecord_RabbitStreams)(nil),
		(*ReadRecord_RedisPubsub)(nil),
		(*ReadRecord_RedisStreams)(nil),
		(*ReadRecord_NatsJetstream)(nil),
		(*ReadRecord_AwsKinesis)(nil),
		(*ReadRecord_Memphis)(nil),
	}
}

// Used as an arg for write requests (server & cli)
type WriteRecord struct {
	// If encoding options are provided, this value will be updated by plumber
	// to contain the encoded payload _before_ passing it to the backend.
	// @gotags: kong:"help='Input string',name=input,xor=input,default"
	Input string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty" kong:"help='Input string',name=input,xor=input,default"`
	// @gotags: kong:"help='Optional metadata a backend may use (key1=value,key2=value,etc)',name=input-metadata"
	InputMetadata        map[string]string `protobuf:"bytes,2,rep,name=input_metadata,json=inputMetadata,proto3" json:"input_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" kong:"help='Optional metadata a backend may use (key1=value,key2=value,etc)',name=input-metadata"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *WriteRecord) Reset()         { *m = WriteRecord{} }
func (m *WriteRecord) String() string { return proto.CompactTextString(m) }
func (*WriteRecord) ProtoMessage()    {}
func (*WriteRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_79a95b759e1a93f6, []int{1}
}

func (m *WriteRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRecord.Unmarshal(m, b)
}
func (m *WriteRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRecord.Marshal(b, m, deterministic)
}
func (m *WriteRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRecord.Merge(m, src)
}
func (m *WriteRecord) XXX_Size() int {
	return xxx_messageInfo_WriteRecord.Size(m)
}
func (m *WriteRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRecord.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRecord proto.InternalMessageInfo

func (m *WriteRecord) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *WriteRecord) GetInputMetadata() map[string]string {
	if m != nil {
		return m.InputMetadata
	}
	return nil
}

// Used for communicating errors that occur during a read, write, relay, etc.
type ErrorRecord struct {
	OccurredAtUnixTsUtc  int64             `protobuf:"varint,1,opt,name=occurred_at_unix_ts_utc,json=occurredAtUnixTsUtc,proto3" json:"occurred_at_unix_ts_utc,omitempty"`
	Error                string            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Metadata             map[string][]byte `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ErrorRecord) Reset()         { *m = ErrorRecord{} }
func (m *ErrorRecord) String() string { return proto.CompactTextString(m) }
func (*ErrorRecord) ProtoMessage()    {}
func (*ErrorRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_79a95b759e1a93f6, []int{2}
}

func (m *ErrorRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorRecord.Unmarshal(m, b)
}
func (m *ErrorRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorRecord.Marshal(b, m, deterministic)
}
func (m *ErrorRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorRecord.Merge(m, src)
}
func (m *ErrorRecord) XXX_Size() int {
	return xxx_messageInfo_ErrorRecord.Size(m)
}
func (m *ErrorRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorRecord proto.InternalMessageInfo

func (m *ErrorRecord) GetOccurredAtUnixTsUtc() int64 {
	if m != nil {
		return m.OccurredAtUnixTsUtc
	}
	return 0
}

func (m *ErrorRecord) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ErrorRecord) GetMetadata() map[string][]byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadRecord)(nil), "protos.records.ReadRecord")
	proto.RegisterMapType((map[string]string)(nil), "protos.records.ReadRecord.MetadataEntry")
	proto.RegisterType((*WriteRecord)(nil), "protos.records.WriteRecord")
	proto.RegisterMapType((map[string]string)(nil), "protos.records.WriteRecord.InputMetadataEntry")
	proto.RegisterType((*ErrorRecord)(nil), "protos.records.ErrorRecord")
	proto.RegisterMapType((map[string][]byte)(nil), "protos.records.ErrorRecord.MetadataEntry")
}

func init() { proto.RegisterFile("records/ps_records_base.proto", fileDescriptor_79a95b759e1a93f6) }

var fileDescriptor_79a95b759e1a93f6 = []byte{
	// 1037 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0x6b, 0x6f, 0x1b, 0x45,
	0x17, 0xc7, 0xed, 0x38, 0xd7, 0x71, 0x9c, 0xe7, 0x61, 0x1a, 0xc8, 0x60, 0x1a, 0x62, 0xc2, 0xa5,
	0xa6, 0x52, 0x6d, 0x68, 0x2b, 0x84, 0xca, 0x45, 0x4d, 0x4b, 0xe8, 0x86, 0x28, 0xc8, 0x59, 0x27,
	0xaa, 0xc4, 0x9b, 0xd5, 0xec, 0xee, 0xd4, 0xde, 0xda, 0x7b, 0xf1, 0x5c, 0xec, 0x84, 0x0f, 0xc1,
	0x27, 0xe2, 0xab, 0xf0, 0x1e, 0xbe, 0x05, 0x9a, 0xcb, 0xda, 0x6b, 0x7b, 0x1c, 0xd4, 0x57, 0xd9,
	0x99, 0xf3, 0x3b, 0x67, 0xce, 0x99, 0x33, 0xfe, 0x9f, 0x80, 0x43, 0x4a, 0x82, 0x94, 0x86, 0xac,
	0x9d, 0x31, 0xcf, 0x7c, 0x7a, 0x3e, 0x66, 0xa4, 0x95, 0xd1, 0x94, 0xa7, 0x70, 0x4f, 0xfd, 0x61,
	0x2d, 0x63, 0xaa, 0x7f, 0x62, 0xc1, 0x71, 0xc0, 0xa3, 0x31, 0x89, 0x47, 0xda, 0xa5, 0xfe, 0x99,
	0x0d, 0x99, 0x30, 0x6f, 0x10, 0x25, 0x84, 0x45, 0xcc, 0x50, 0x8d, 0x15, 0x14, 0x1b, 0xe5, 0x44,
	0xd3, 0x46, 0xfc, 0x2e, 0x28, 0xf1, 0xc8, 0x98, 0x24, 0xdc, 0xeb, 0x0b, 0xdf, 0x90, 0x0f, 0x57,
	0x92, 0x8c, 0xd0, 0x71, 0x14, 0x10, 0xcf, 0x17, 0x79, 0xd4, 0x4f, 0x2d, 0x6c, 0x2f, 0xc8, 0xbc,
	0x4c, 0xf8, 0x6c, 0x1a, 0xf0, 0x63, 0x0b, 0x34, 0xc0, 0x6f, 0x06, 0xd8, 0xd8, 0x8f, 0x6c, 0x76,
	0xe1, 0xcf, 0xee, 0xc0, 0x56, 0x5d, 0x4c, 0xe2, 0xac, 0x3f, 0xad, 0xdf, 0x76, 0x44, 0x9c, 0x26,
	0xbd, 0xd4, 0xd8, 0x6d, 0x7d, 0x89, 0x47, 0x9c, 0xdf, 0x61, 0x4e, 0x30, 0xcf, 0xa3, 0x3f, 0x58,
	0x61, 0xf6, 0xde, 0x12, 0xce, 0x38, 0x25, 0x38, 0xfe, 0x2f, 0x50, 0x53, 0x51, 0xd2, 0x33, 0xe0,
	0x7d, 0x1b, 0xc8, 0xf2, 0x7a, 0x6d, 0xcf, 0x22, 0x4b, 0x19, 0xef, 0x51, 0xc2, 0xee, 0xb8, 0xb3,
	0x4c, 0x0c, 0x19, 0xa6, 0x77, 0x00, 0x14, 0xfb, 0x7e, 0xc4, 0xef, 0xc8, 0x55, 0x03, 0x26, 0xdb,
	0xfc, 0xa8, 0xcf, 0x6d, 0x20, 0x09, 0x23, 0x36, 0xdf, 0xe5, 0x2f, 0x56, 0x62, 0x73, 0xe1, 0x8e,
	0xff, 0xd8, 0x05, 0xc0, 0x25, 0x38, 0x74, 0x15, 0x03, 0x0f, 0x01, 0x88, 0x09, 0x63, 0xb8, 0x47,
	0xbc, 0x28, 0x44, 0xe5, 0x46, 0xb9, 0xb9, 0xe3, 0xee, 0x98, 0x9d, 0xb3, 0x10, 0xfe, 0x1f, 0x54,
	0x12, 0x11, 0xa3, 0x4a, 0xa3, 0xdc, 0xac, 0xb8, 0xf2, 0x13, 0xfe, 0x04, 0xb6, 0x63, 0xc2, 0x71,
	0x88, 0x39, 0x46, 0xeb, 0x8d, 0x4a, 0xb3, 0xfa, 0xb8, 0xd9, 0x9a, 0xff, 0x59, 0xb5, 0x66, 0xe1,
	0x5b, 0x17, 0x06, 0x3d, 0x4d, 0x38, 0xbd, 0x75, 0xa7, 0x9e, 0xf0, 0x29, 0x38, 0xa0, 0x24, 0x20,
	0xd1, 0x98, 0x84, 0x1e, 0xe6, 0x9e, 0x48, 0xa2, 0x1b, 0x8f, 0x33, 0x4f, 0xf0, 0x00, 0x6d, 0xaa,
	0xb3, 0xee, 0xe5, 0xe6, 0x13, 0x7e, 0x9d, 0x44, 0x37, 0x57, 0xec, 0x9a, 0x07, 0x10, 0x81, 0xad,
	0x0c, 0xdf, 0x0e, 0x53, 0x1c, 0xa2, 0xa0, 0x51, 0x6e, 0xee, 0xba, 0xf9, 0x12, 0x3e, 0x02, 0x1b,
	0xea, 0x49, 0xa3, 0xb0, 0x51, 0x6e, 0x56, 0x1f, 0xbf, 0xbf, 0x98, 0xd2, 0xb9, 0x34, 0x3a, 0x25,
	0x57, 0x53, 0xf0, 0x2b, 0xb0, 0xa9, 0xef, 0x1a, 0x11, 0xc5, 0x7f, 0xb0, 0x54, 0x82, 0xb2, 0x3a,
	0x25, 0xd7, 0x70, 0xf0, 0x1b, 0xb0, 0x9d, 0x2b, 0x03, 0x7a, 0xa3, 0x7c, 0xd0, 0xa2, 0xcf, 0x89,
	0xb2, 0x5f, 0x5c, 0x3a, 0x25, 0x77, 0xca, 0xc2, 0xaf, 0xc1, 0x96, 0x11, 0x02, 0xd4, 0xb3, 0x1f,
	0x75, 0xf2, 0xba, 0xdb, 0xbd, 0xec, 0xca, 0xa3, 0xf0, 0x84, 0x75, 0x47, 0x0c, 0xbe, 0x02, 0xff,
	0x5b, 0x50, 0x06, 0xd4, 0x57, 0xae, 0x87, 0x4b, 0xae, 0x12, 0x3b, 0x95, 0x94, 0x23, 0x7c, 0xa7,
	0xe4, 0xd6, 0x70, 0x71, 0x03, 0x5e, 0x80, 0xf7, 0x96, 0x84, 0x03, 0x45, 0x2a, 0xd4, 0x91, 0x35,
	0x54, 0x57, 0x73, 0x2f, 0x04, 0x73, 0x4a, 0xae, 0x4e, 0x62, 0xb6, 0x05, 0x9f, 0x01, 0x30, 0xd3,
	0x16, 0xf4, 0x56, 0xc5, 0xf9, 0x70, 0x31, 0xce, 0xab, 0x97, 0x9d, 0x8e, 0xf0, 0xbb, 0x2a, 0x9d,
	0x9d, 0x5e, 0x90, 0x75, 0x14, 0x2d, 0x2f, 0x5c, 0x4b, 0x0a, 0x1a, 0xd8, 0x6f, 0xe1, 0x5c, 0xf8,
	0xfa, 0xea, 0x0c, 0x27, 0x3b, 0xaa, 0x14, 0x04, 0x0d, 0xed, 0x1d, 0xbd, 0x90, 0x46, 0xd9, 0x51,
	0x45, 0xc1, 0x87, 0x60, 0x5d, 0x0a, 0x0a, 0x8a, 0x15, 0xbd, 0xbf, 0x44, 0x5f, 0x5e, 0x5d, 0x39,
	0x25, 0x57, 0x31, 0x92, 0x95, 0xaa, 0x80, 0x12, 0x3b, 0xfb, 0x2b, 0xe6, 0xb2, 0x7e, 0xc5, 0xc0,
	0x9f, 0xc1, 0xde, 0xbc, 0x82, 0xa0, 0xd4, 0xde, 0x0b, 0xe9, 0xd5, 0xcd, 0x21, 0xd9, 0x8b, 0xa4,
	0xb8, 0x01, 0x1f, 0x80, 0x4a, 0xc2, 0x46, 0x28, 0x53, 0xce, 0xf7, 0x96, 0x9c, 0xbb, 0xb2, 0x74,
	0x49, 0xc8, 0x87, 0x96, 0x6b, 0x0d, 0x1a, 0xd9, 0x1f, 0x5a, 0xc7, 0xd8, 0xe5, 0x43, 0xcb, 0x59,
	0x79, 0xc3, 0x5a, 0x80, 0x10, 0xb5, 0xdf, 0x70, 0x47, 0x59, 0xe5, 0x0d, 0x6b, 0x4e, 0x96, 0x36,
	0x2f, 0x38, 0x88, 0xd9, 0x4b, 0xd3, 0x3f, 0x06, 0x5d, 0x8b, 0x3c, 0xb4, 0x46, 0x8b, 0x1b, 0xf0,
	0x39, 0xd8, 0x2d, 0xea, 0x11, 0xe2, 0x2a, 0xca, 0x47, 0xcb, 0xaa, 0x10, 0x46, 0x4c, 0x3f, 0x07,
	0xa7, 0xe4, 0x56, 0xe9, 0x6c, 0x09, 0x5f, 0x82, 0xda, 0x9c, 0x54, 0x21, 0xa1, 0x42, 0xdc, 0xb7,
	0x86, 0x98, 0xe5, 0xa1, 0x8f, 0xcd, 0xd3, 0xc8, 0x3b, 0x35, 0x1d, 0x0a, 0x68, 0xbc, 0xba, 0x53,
	0xbf, 0xe4, 0x50, 0xde, 0xa9, 0xe9, 0x06, 0xfc, 0x01, 0x54, 0x0b, 0x03, 0x1e, 0x4d, 0x54, 0x90,
	0xba, 0xe5, 0x57, 0x7b, 0xae, 0x09, 0xa7, 0xe4, 0x02, 0x3c, 0x61, 0x66, 0x05, 0x9f, 0x80, 0x2d,
	0x33, 0x1b, 0xd1, 0x8d, 0x72, 0x3d, 0x58, 0x7a, 0x8b, 0xda, 0xec, 0x94, 0xdc, 0x9c, 0x84, 0x10,
	0xac, 0x7b, 0x14, 0x4f, 0xd0, 0xdf, 0x5b, 0x4a, 0xd6, 0x2a, 0x2e, 0x9e, 0xc0, 0x23, 0x50, 0xf5,
	0xb2, 0xa1, 0x88, 0x7d, 0x42, 0xa5, 0x34, 0xff, 0xb3, 0xa5, 0xb5, 0xb9, 0xa3, 0xb7, 0xce, 0xc2,
	0xfa, 0x77, 0xa0, 0x36, 0x27, 0xaf, 0x52, 0xac, 0x07, 0xe4, 0xd6, 0x88, 0xb8, 0xfc, 0x84, 0xfb,
	0x60, 0x63, 0x8c, 0x87, 0x82, 0xa0, 0x35, 0xb5, 0xa7, 0x17, 0xcf, 0xd6, 0xbe, 0x2d, 0xbf, 0xd8,
	0x06, 0x9b, 0x5a, 0xa2, 0x8f, 0xff, 0x2c, 0x83, 0xea, 0x6b, 0x1a, 0x71, 0x62, 0x26, 0xc2, 0x3e,
	0xd8, 0x88, 0x92, 0x4c, 0x70, 0x13, 0x47, 0x2f, 0xe0, 0x35, 0xd8, 0x53, 0x1f, 0xde, 0x54, 0xfc,
	0xd7, 0x94, 0xf8, 0xb7, 0x16, 0xab, 0x2b, 0x84, 0x6a, 0x9d, 0x49, 0x8f, 0xf9, 0x11, 0x50, 0x8b,
	0x8a, 0x7b, 0xf5, 0xe7, 0x00, 0x2e, 0x43, 0xef, 0x52, 0xc8, 0xf1, 0x5f, 0x65, 0x50, 0x3d, 0xa5,
	0x34, 0xa5, 0x26, 0xfd, 0xa7, 0xe0, 0x20, 0x0d, 0x02, 0x41, 0xe9, 0xf2, 0x64, 0x29, 0xeb, 0xc9,
	0x92, 0x9b, 0x8b, 0x93, 0x65, 0x1f, 0x6c, 0x10, 0x19, 0x24, 0x8f, 0xaf, 0x16, 0xf0, 0xb4, 0x30,
	0xeb, 0x2a, 0xaa, 0xdc, 0x2f, 0x17, 0xcb, 0x2d, 0x1c, 0xbd, 0x6a, 0xd8, 0xbd, 0x73, 0xa3, 0x76,
	0x8b, 0x8d, 0xfa, 0xf1, 0xb7, 0xef, 0x7b, 0x11, 0x97, 0xff, 0x1e, 0x06, 0x69, 0xdc, 0xf6, 0x31,
	0x0f, 0xfa, 0x41, 0x4a, 0xb3, 0xb6, 0x79, 0x1a, 0x8f, 0x58, 0xd0, 0x27, 0x31, 0x66, 0x6d, 0x5f,
	0x44, 0xc3, 0xb0, 0xdd, 0x4b, 0xdb, 0x3a, 0xc1, 0xb6, 0x49, 0xd0, 0xdf, 0x54, 0xeb, 0x27, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x9b, 0x97, 0x55, 0x1b, 0x0b, 0x00, 0x00,
}
