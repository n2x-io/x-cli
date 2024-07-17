// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0--rc1
// source: n2x/protobuf/resources/v1/services/thirdParty/integrations.proto

package thirdParty

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Integrations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Github    *GitHub    `protobuf:"bytes,1,opt,name=github,proto3" json:"github,omitempty"`
	Pagerduty *PagerDuty `protobuf:"bytes,11,opt,name=pagerduty,proto3" json:"pagerduty,omitempty"`
	Slack     *Slack     `protobuf:"bytes,21,opt,name=slack,proto3" json:"slack,omitempty"`
	Unmasked  bool       `protobuf:"varint,1001,opt,name=unmasked,proto3" json:"unmasked,omitempty"`
}

func (x *Integrations) Reset() {
	*x = Integrations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integrations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integrations) ProtoMessage() {}

func (x *Integrations) ProtoReflect() protoreflect.Message {
	mi := &file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integrations.ProtoReflect.Descriptor instead.
func (*Integrations) Descriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDescGZIP(), []int{0}
}

func (x *Integrations) GetGithub() *GitHub {
	if x != nil {
		return x.Github
	}
	return nil
}

func (x *Integrations) GetPagerduty() *PagerDuty {
	if x != nil {
		return x.Pagerduty
	}
	return nil
}

func (x *Integrations) GetSlack() *Slack {
	if x != nil {
		return x.Slack
	}
	return nil
}

func (x *Integrations) GetUnmasked() bool {
	if x != nil {
		return x.Unmasked
	}
	return false
}

var File_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto protoreflect.FileDescriptor

var file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDesc = []byte{
	0x0a, 0x40, 0x6e, 0x32, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x1a, 0x38,
	0x6e, 0x32, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x73, 0x61,
	0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x0c, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x52, 0x06, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x12, 0x33, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x72, 0x64, 0x75,
	0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44, 0x75, 0x74, 0x79, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x72, 0x64, 0x75, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x6c,
	0x61, 0x63, 0x6b, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x52, 0x05, 0x73, 0x6c,
	0x61, 0x63, 0x6b, 0x12, 0x1b, 0x0a, 0x08, 0x75, 0x6e, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64, 0x18,
	0xe9, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x6e, 0x6d, 0x61, 0x73, 0x6b, 0x65, 0x64,
	0x42, 0x35, 0x5a, 0x33, 0x6e, 0x32, 0x78, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x78, 0x2d, 0x61, 0x70,
	0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDescOnce sync.Once
	file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDescData = file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDesc
)

func file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDescGZIP() []byte {
	file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDescOnce.Do(func() {
		file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDescData = protoimpl.X.CompressGZIP(file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDescData)
	})
	return file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDescData
}

var file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_goTypes = []any{
	(*Integrations)(nil), // 0: thirdParty.Integrations
	(*GitHub)(nil),       // 1: thirdParty.GitHub
	(*PagerDuty)(nil),    // 2: thirdParty.PagerDuty
	(*Slack)(nil),        // 3: thirdParty.Slack
}
var file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_depIdxs = []int32{
	1, // 0: thirdParty.Integrations.github:type_name -> thirdParty.GitHub
	2, // 1: thirdParty.Integrations.pagerduty:type_name -> thirdParty.PagerDuty
	3, // 2: thirdParty.Integrations.slack:type_name -> thirdParty.Slack
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_init() }
func file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_init() {
	if File_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto != nil {
		return
	}
	file_n2x_protobuf_resources_v1_services_thirdParty_saas_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Integrations); i {
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
			RawDescriptor: file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_goTypes,
		DependencyIndexes: file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_depIdxs,
		MessageInfos:      file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_msgTypes,
	}.Build()
	File_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto = out.File
	file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_rawDesc = nil
	file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_goTypes = nil
	file_n2x_protobuf_resources_v1_services_thirdParty_integrations_proto_depIdxs = nil
}
