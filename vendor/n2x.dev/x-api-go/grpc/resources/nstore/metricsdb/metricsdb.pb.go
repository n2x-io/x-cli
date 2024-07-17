// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0--rc1
// source: n2x/protobuf/resources/v1/nstore/metricsdb/metricsdb.proto

package metricsdb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	nstore "n2x.dev/x-api-go/grpc/resources/nstore"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HostMetricType int32

const (
	HostMetricType_UNDEFINED       HostMetricType = 0
	HostMetricType_NET_RX_BYTES    HostMetricType = 11
	HostMetricType_NET_TX_BYTES    HostMetricType = 12
	HostMetricType_HOST_LOAD_AVG   HostMetricType = 111
	HostMetricType_HOST_CPU_USAGE  HostMetricType = 121
	HostMetricType_HOST_MEM_USAGE  HostMetricType = 131
	HostMetricType_HOST_DISK_USAGE HostMetricType = 141
)

// Enum value maps for HostMetricType.
var (
	HostMetricType_name = map[int32]string{
		0:   "UNDEFINED",
		11:  "NET_RX_BYTES",
		12:  "NET_TX_BYTES",
		111: "HOST_LOAD_AVG",
		121: "HOST_CPU_USAGE",
		131: "HOST_MEM_USAGE",
		141: "HOST_DISK_USAGE",
	}
	HostMetricType_value = map[string]int32{
		"UNDEFINED":       0,
		"NET_RX_BYTES":    11,
		"NET_TX_BYTES":    12,
		"HOST_LOAD_AVG":   111,
		"HOST_CPU_USAGE":  121,
		"HOST_MEM_USAGE":  131,
		"HOST_DISK_USAGE": 141,
	}
)

func (x HostMetricType) Enum() *HostMetricType {
	p := new(HostMetricType)
	*p = x
	return p
}

func (x HostMetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HostMetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_enumTypes[0].Descriptor()
}

func (HostMetricType) Type() protoreflect.EnumType {
	return &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_enumTypes[0]
}

func (x HostMetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HostMetricType.Descriptor instead.
func (HostMetricType) EnumDescriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescGZIP(), []int{0}
}

type HostMetricsQueryType int32

const (
	HostMetricsQueryType_UNKNOWN_QUERY    HostMetricsQueryType = 0
	HostMetricsQueryType_QUERY_NET_USAGE  HostMetricsQueryType = 11
	HostMetricsQueryType_QUERY_LOAD_AVG   HostMetricsQueryType = 21
	HostMetricsQueryType_QUERY_CPU_USAGE  HostMetricsQueryType = 31
	HostMetricsQueryType_QUERY_MEM_USAGE  HostMetricsQueryType = 41
	HostMetricsQueryType_QUERY_DISK_USAGE HostMetricsQueryType = 51
)

// Enum value maps for HostMetricsQueryType.
var (
	HostMetricsQueryType_name = map[int32]string{
		0:  "UNKNOWN_QUERY",
		11: "QUERY_NET_USAGE",
		21: "QUERY_LOAD_AVG",
		31: "QUERY_CPU_USAGE",
		41: "QUERY_MEM_USAGE",
		51: "QUERY_DISK_USAGE",
	}
	HostMetricsQueryType_value = map[string]int32{
		"UNKNOWN_QUERY":    0,
		"QUERY_NET_USAGE":  11,
		"QUERY_LOAD_AVG":   21,
		"QUERY_CPU_USAGE":  31,
		"QUERY_MEM_USAGE":  41,
		"QUERY_DISK_USAGE": 51,
	}
)

func (x HostMetricsQueryType) Enum() *HostMetricsQueryType {
	p := new(HostMetricsQueryType)
	*p = x
	return p
}

func (x HostMetricsQueryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HostMetricsQueryType) Descriptor() protoreflect.EnumDescriptor {
	return file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_enumTypes[1].Descriptor()
}

func (HostMetricsQueryType) Type() protoreflect.EnumType {
	return &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_enumTypes[1]
}

func (x HostMetricsQueryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HostMetricsQueryType.Descriptor instead.
func (HostMetricsQueryType) EnumDescriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescGZIP(), []int{1}
}

type HostMetricDataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64            `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TimeRange nstore.TimeRange `protobuf:"varint,11,opt,name=timeRange,proto3,enum=nstore.TimeRange" json:"timeRange,omitempty"`   // namespace
	Metric    HostMetricType   `protobuf:"varint,21,opt,name=metric,proto3,enum=metricsdb.HostMetricType" json:"metric,omitempty"` // key
	Value     float64          `protobuf:"fixed64,101,opt,name=value,proto3" json:"value,omitempty"`                               // value
}

func (x *HostMetricDataPoint) Reset() {
	*x = HostMetricDataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostMetricDataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostMetricDataPoint) ProtoMessage() {}

func (x *HostMetricDataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostMetricDataPoint.ProtoReflect.Descriptor instead.
func (*HostMetricDataPoint) Descriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescGZIP(), []int{0}
}

func (x *HostMetricDataPoint) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *HostMetricDataPoint) GetTimeRange() nstore.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nstore.TimeRange(0)
}

func (x *HostMetricDataPoint) GetMetric() HostMetricType {
	if x != nil {
		return x.Metric
	}
	return HostMetricType_UNDEFINED
}

func (x *HostMetricDataPoint) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type HostMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request   *nstore.DataRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Type      HostMetricsQueryType `protobuf:"varint,11,opt,name=type,proto3,enum=metricsdb.HostMetricsQueryType" json:"type,omitempty"`
	TimeRange nstore.TimeRange     `protobuf:"varint,21,opt,name=timeRange,proto3,enum=nstore.TimeRange" json:"timeRange,omitempty"` // namespace
}

func (x *HostMetricsRequest) Reset() {
	*x = HostMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostMetricsRequest) ProtoMessage() {}

func (x *HostMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostMetricsRequest.ProtoReflect.Descriptor instead.
func (*HostMetricsRequest) Descriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescGZIP(), []int{1}
}

func (x *HostMetricsRequest) GetRequest() *nstore.DataRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *HostMetricsRequest) GetType() HostMetricsQueryType {
	if x != nil {
		return x.Type
	}
	return HostMetricsQueryType_UNKNOWN_QUERY
}

func (x *HostMetricsRequest) GetTimeRange() nstore.TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nstore.TimeRange(0)
}

type HostMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string         `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string         `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NodeID    string         `protobuf:"bytes,5,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	QueryID   string         `protobuf:"bytes,11,opt,name=queryID,proto3" json:"queryID,omitempty"`
	Metrics   []*HostMetrics `protobuf:"bytes,101,rep,name=metrics,proto3" json:"metrics,omitempty"`
	Timestamp int64          `protobuf:"varint,1001,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *HostMetricsResponse) Reset() {
	*x = HostMetricsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostMetricsResponse) ProtoMessage() {}

func (x *HostMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostMetricsResponse.ProtoReflect.Descriptor instead.
func (*HostMetricsResponse) Descriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescGZIP(), []int{2}
}

func (x *HostMetricsResponse) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *HostMetricsResponse) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *HostMetricsResponse) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *HostMetricsResponse) GetQueryID() string {
	if x != nil {
		return x.QueryID
	}
	return ""
}

func (x *HostMetricsResponse) GetMetrics() []*HostMetrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *HostMetricsResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type HostMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64                   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      map[string]*MetricValue `protobuf:"bytes,11,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // map[HostMetricType]*MetricValue
}

func (x *HostMetrics) Reset() {
	*x = HostMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostMetrics) ProtoMessage() {}

func (x *HostMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostMetrics.ProtoReflect.Descriptor instead.
func (*HostMetrics) Descriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescGZIP(), []int{3}
}

func (x *HostMetrics) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *HostMetrics) GetData() map[string]*MetricValue {
	if x != nil {
		return x.Data
	}
	return nil
}

type MetricValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,11,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MetricValue) Reset() {
	*x = MetricValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricValue) ProtoMessage() {}

func (x *MetricValue) ProtoReflect() protoreflect.Message {
	mi := &file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricValue.ProtoReflect.Descriptor instead.
func (*MetricValue) Descriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescGZIP(), []int{4}
}

func (x *MetricValue) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto protoreflect.FileDescriptor

var file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x6e, 0x32, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x62, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x62, 0x1a, 0x2d, 0x6e, 0x32, 0x78, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x13, 0x48, 0x6f, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2f, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x31, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x12, 0x48, 0x6f, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x64, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2f, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x13, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x64, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xb2, 0x01, 0x0a, 0x0b, 0x48, 0x6f, 0x73, 0x74,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x64, 0x62, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x4f, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x64, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x23, 0x0a, 0x0b,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x2a, 0x95, 0x01, 0x0a, 0x0e, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x45, 0x54, 0x5f, 0x52, 0x58, 0x5f, 0x42, 0x59,
	0x54, 0x45, 0x53, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x45, 0x54, 0x5f, 0x54, 0x58, 0x5f,
	0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x48, 0x4f, 0x53, 0x54, 0x5f,
	0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x41, 0x56, 0x47, 0x10, 0x6f, 0x12, 0x12, 0x0a, 0x0e, 0x48, 0x4f,
	0x53, 0x54, 0x5f, 0x43, 0x50, 0x55, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x79, 0x12, 0x13,
	0x0a, 0x0e, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x4d, 0x45, 0x4d, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45,
	0x10, 0x83, 0x01, 0x12, 0x14, 0x0a, 0x0f, 0x48, 0x4f, 0x53, 0x54, 0x5f, 0x44, 0x49, 0x53, 0x4b,
	0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x8d, 0x01, 0x2a, 0x92, 0x01, 0x0a, 0x14, 0x48, 0x6f,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x51, 0x55,
	0x45, 0x52, 0x59, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x4e,
	0x45, 0x54, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x12, 0x0a, 0x0e, 0x51, 0x55,
	0x45, 0x52, 0x59, 0x5f, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x41, 0x56, 0x47, 0x10, 0x15, 0x12, 0x13,
	0x0a, 0x0f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x43, 0x50, 0x55, 0x5f, 0x55, 0x53, 0x41, 0x47,
	0x45, 0x10, 0x1f, 0x12, 0x13, 0x0a, 0x0f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x4d, 0x45, 0x4d,
	0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x29, 0x12, 0x14, 0x0a, 0x10, 0x51, 0x55, 0x45, 0x52,
	0x59, 0x5f, 0x44, 0x49, 0x53, 0x4b, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0x33, 0x42, 0x32,
	0x5a, 0x30, 0x6e, 0x32, 0x78, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x78, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescOnce sync.Once
	file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescData = file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDesc
)

func file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescGZIP() []byte {
	file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescOnce.Do(func() {
		file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescData)
	})
	return file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDescData
}

var file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_goTypes = []any{
	(HostMetricType)(0),         // 0: metricsdb.HostMetricType
	(HostMetricsQueryType)(0),   // 1: metricsdb.HostMetricsQueryType
	(*HostMetricDataPoint)(nil), // 2: metricsdb.HostMetricDataPoint
	(*HostMetricsRequest)(nil),  // 3: metricsdb.HostMetricsRequest
	(*HostMetricsResponse)(nil), // 4: metricsdb.HostMetricsResponse
	(*HostMetrics)(nil),         // 5: metricsdb.HostMetrics
	(*MetricValue)(nil),         // 6: metricsdb.MetricValue
	nil,                         // 7: metricsdb.HostMetrics.DataEntry
	(nstore.TimeRange)(0),       // 8: nstore.TimeRange
	(*nstore.DataRequest)(nil),  // 9: nstore.DataRequest
}
var file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_depIdxs = []int32{
	8, // 0: metricsdb.HostMetricDataPoint.timeRange:type_name -> nstore.TimeRange
	0, // 1: metricsdb.HostMetricDataPoint.metric:type_name -> metricsdb.HostMetricType
	9, // 2: metricsdb.HostMetricsRequest.request:type_name -> nstore.DataRequest
	1, // 3: metricsdb.HostMetricsRequest.type:type_name -> metricsdb.HostMetricsQueryType
	8, // 4: metricsdb.HostMetricsRequest.timeRange:type_name -> nstore.TimeRange
	5, // 5: metricsdb.HostMetricsResponse.metrics:type_name -> metricsdb.HostMetrics
	7, // 6: metricsdb.HostMetrics.data:type_name -> metricsdb.HostMetrics.DataEntry
	6, // 7: metricsdb.HostMetrics.DataEntry.value:type_name -> metricsdb.MetricValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_init() }
func file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_init() {
	if File_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HostMetricDataPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HostMetricsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HostMetricsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*HostMetrics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MetricValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_goTypes,
		DependencyIndexes: file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_depIdxs,
		EnumInfos:         file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_enumTypes,
		MessageInfos:      file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_msgTypes,
	}.Build()
	File_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto = out.File
	file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_rawDesc = nil
	file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_goTypes = nil
	file_n2x_protobuf_resources_v1_nstore_metricsdb_metricsdb_proto_depIdxs = nil
}
