// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: deposit.proto

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

type Deposit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletId string  `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	Amount   float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Deposit) Reset() {
	*x = Deposit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deposit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deposit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deposit) ProtoMessage() {}

func (x *Deposit) ProtoReflect() protoreflect.Message {
	mi := &file_deposit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deposit.ProtoReflect.Descriptor instead.
func (*Deposit) Descriptor() ([]byte, []int) {
	return file_deposit_proto_rawDescGZIP(), []int{0}
}

func (x *Deposit) GetWalletId() string {
	if x != nil {
		return x.WalletId
	}
	return ""
}

func (x *Deposit) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type DepositHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletId string     `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	Deposits []*Deposit `protobuf:"bytes,2,rep,name=deposits,proto3" json:"deposits,omitempty"`
}

func (x *DepositHistory) Reset() {
	*x = DepositHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deposit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepositHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositHistory) ProtoMessage() {}

func (x *DepositHistory) ProtoReflect() protoreflect.Message {
	mi := &file_deposit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositHistory.ProtoReflect.Descriptor instead.
func (*DepositHistory) Descriptor() ([]byte, []int) {
	return file_deposit_proto_rawDescGZIP(), []int{1}
}

func (x *DepositHistory) GetWalletId() string {
	if x != nil {
		return x.WalletId
	}
	return ""
}

func (x *DepositHistory) GetDeposits() []*Deposit {
	if x != nil {
		return x.Deposits
	}
	return nil
}

var File_deposit_proto protoreflect.FileDescriptor

var file_deposit_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3e, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x53, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x08, 0x64, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_deposit_proto_rawDescOnce sync.Once
	file_deposit_proto_rawDescData = file_deposit_proto_rawDesc
)

func file_deposit_proto_rawDescGZIP() []byte {
	file_deposit_proto_rawDescOnce.Do(func() {
		file_deposit_proto_rawDescData = protoimpl.X.CompressGZIP(file_deposit_proto_rawDescData)
	})
	return file_deposit_proto_rawDescData
}

var file_deposit_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_deposit_proto_goTypes = []interface{}{
	(*Deposit)(nil),        // 0: Deposit
	(*DepositHistory)(nil), // 1: DepositHistory
}
var file_deposit_proto_depIdxs = []int32{
	0, // 0: DepositHistory.deposits:type_name -> Deposit
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_deposit_proto_init() }
func file_deposit_proto_init() {
	if File_deposit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deposit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deposit); i {
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
		file_deposit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepositHistory); i {
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
			RawDescriptor: file_deposit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deposit_proto_goTypes,
		DependencyIndexes: file_deposit_proto_depIdxs,
		MessageInfos:      file_deposit_proto_msgTypes,
	}.Build()
	File_deposit_proto = out.File
	file_deposit_proto_rawDesc = nil
	file_deposit_proto_goTypes = nil
	file_deposit_proto_depIdxs = nil
}
