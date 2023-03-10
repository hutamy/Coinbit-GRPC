// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: flag.proto

package pb

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

type FlagEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlagRemoved bool  `protobuf:"varint,1,opt,name=flag_removed,json=flagRemoved,proto3" json:"flag_removed,omitempty"`
	StartPeriod int64 `protobuf:"varint,2,opt,name=start_period,json=startPeriod,proto3" json:"start_period,omitempty"`
}

func (x *FlagEvent) Reset() {
	*x = FlagEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagEvent) ProtoMessage() {}

func (x *FlagEvent) ProtoReflect() protoreflect.Message {
	mi := &file_flag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagEvent.ProtoReflect.Descriptor instead.
func (*FlagEvent) Descriptor() ([]byte, []int) {
	return file_flag_proto_rawDescGZIP(), []int{0}
}

func (x *FlagEvent) GetFlagRemoved() bool {
	if x != nil {
		return x.FlagRemoved
	}
	return false
}

func (x *FlagEvent) GetStartPeriod() int64 {
	if x != nil {
		return x.StartPeriod
	}
	return 0
}

type FlagValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flagged     bool  `protobuf:"varint,1,opt,name=flagged,proto3" json:"flagged,omitempty"`
	StartPeriod int64 `protobuf:"varint,2,opt,name=start_period,json=startPeriod,proto3" json:"start_period,omitempty"`
}

func (x *FlagValue) Reset() {
	*x = FlagValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlagValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlagValue) ProtoMessage() {}

func (x *FlagValue) ProtoReflect() protoreflect.Message {
	mi := &file_flag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlagValue.ProtoReflect.Descriptor instead.
func (*FlagValue) Descriptor() ([]byte, []int) {
	return file_flag_proto_rawDescGZIP(), []int{1}
}

func (x *FlagValue) GetFlagged() bool {
	if x != nil {
		return x.Flagged
	}
	return false
}

func (x *FlagValue) GetStartPeriod() int64 {
	if x != nil {
		return x.StartPeriod
	}
	return 0
}

var File_flag_proto protoreflect.FileDescriptor

var file_flag_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6c, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x09,
	0x46, 0x6c, 0x61, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6c, 0x61,
	0x67, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x66, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22,
	0x48, 0x0a, 0x09, 0x46, 0x6c, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x6c, 0x61, 0x67, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x66,
	0x6c, 0x61, 0x67, 0x67, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flag_proto_rawDescOnce sync.Once
	file_flag_proto_rawDescData = file_flag_proto_rawDesc
)

func file_flag_proto_rawDescGZIP() []byte {
	file_flag_proto_rawDescOnce.Do(func() {
		file_flag_proto_rawDescData = protoimpl.X.CompressGZIP(file_flag_proto_rawDescData)
	})
	return file_flag_proto_rawDescData
}

var file_flag_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flag_proto_goTypes = []interface{}{
	(*FlagEvent)(nil), // 0: FlagEvent
	(*FlagValue)(nil), // 1: FlagValue
}
var file_flag_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flag_proto_init() }
func file_flag_proto_init() {
	if File_flag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagEvent); i {
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
		file_flag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlagValue); i {
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
			RawDescriptor: file_flag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flag_proto_goTypes,
		DependencyIndexes: file_flag_proto_depIdxs,
		MessageInfos:      file_flag_proto_msgTypes,
	}.Build()
	File_flag_proto = out.File
	file_flag_proto_rawDesc = nil
	file_flag_proto_goTypes = nil
	file_flag_proto_depIdxs = nil
}
