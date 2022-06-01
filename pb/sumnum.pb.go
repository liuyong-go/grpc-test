// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.0
// source: proto/sumnum.proto

//protoc -I. --go_out=. --go-grpc_out=.  ./proto/sumnum.proto

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

type SumTwoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int64 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int64 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *SumTwoRequest) Reset() {
	*x = SumTwoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sumnum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumTwoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumTwoRequest) ProtoMessage() {}

func (x *SumTwoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sumnum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumTwoRequest.ProtoReflect.Descriptor instead.
func (*SumTwoRequest) Descriptor() ([]byte, []int) {
	return file_proto_sumnum_proto_rawDescGZIP(), []int{0}
}

func (x *SumTwoRequest) GetNum1() int64 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *SumTwoRequest) GetNum2() int64 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type SumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sumvalue int64 `protobuf:"varint,1,opt,name=sumvalue,proto3" json:"sumvalue,omitempty"`
}

func (x *SumReply) Reset() {
	*x = SumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sumnum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumReply) ProtoMessage() {}

func (x *SumReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sumnum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumReply.ProtoReflect.Descriptor instead.
func (*SumReply) Descriptor() ([]byte, []int) {
	return file_proto_sumnum_proto_rawDescGZIP(), []int{1}
}

func (x *SumReply) GetSumvalue() int64 {
	if x != nil {
		return x.Sumvalue
	}
	return 0
}

type NumsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int64 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *NumsRequest) Reset() {
	*x = NumsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sumnum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumsRequest) ProtoMessage() {}

func (x *NumsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sumnum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumsRequest.ProtoReflect.Descriptor instead.
func (*NumsRequest) Descriptor() ([]byte, []int) {
	return file_proto_sumnum_proto_rawDescGZIP(), []int{2}
}

func (x *NumsRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_proto_sumnum_proto protoreflect.FileDescriptor

var file_proto_sumnum_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x6d, 0x6e, 0x75, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x75, 0x6d, 0x6e, 0x75, 0x6d, 0x22, 0x37, 0x0a, 0x0d,
	0x53, 0x75, 0x6d, 0x54, 0x77, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6e, 0x75, 0x6d,
	0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x6e, 0x75, 0x6d, 0x32, 0x22, 0x26, 0x0a, 0x08, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x6d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x75, 0x6d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1f, 0x0a,
	0x0b, 0x4e, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x32, 0x77,
	0x0a, 0x06, 0x53, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x12, 0x37, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x54,
	0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x15, 0x2e, 0x73, 0x75, 0x6d, 0x6e, 0x75, 0x6d, 0x2e,
	0x53, 0x75, 0x6d, 0x54, 0x77, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x73, 0x75, 0x6d, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x13, 0x2e, 0x73,
	0x75, 0x6d, 0x6e, 0x75, 0x6d, 0x2e, 0x4e, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x73, 0x75, 0x6d, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sumnum_proto_rawDescOnce sync.Once
	file_proto_sumnum_proto_rawDescData = file_proto_sumnum_proto_rawDesc
)

func file_proto_sumnum_proto_rawDescGZIP() []byte {
	file_proto_sumnum_proto_rawDescOnce.Do(func() {
		file_proto_sumnum_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sumnum_proto_rawDescData)
	})
	return file_proto_sumnum_proto_rawDescData
}

var file_proto_sumnum_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_sumnum_proto_goTypes = []interface{}{
	(*SumTwoRequest)(nil), // 0: sumnum.SumTwoRequest
	(*SumReply)(nil),      // 1: sumnum.SumReply
	(*NumsRequest)(nil),   // 2: sumnum.NumsRequest
}
var file_proto_sumnum_proto_depIdxs = []int32{
	0, // 0: sumnum.SumNum.SumTwoNums:input_type -> sumnum.SumTwoRequest
	2, // 1: sumnum.SumNum.SumNums:input_type -> sumnum.NumsRequest
	1, // 2: sumnum.SumNum.SumTwoNums:output_type -> sumnum.SumReply
	1, // 3: sumnum.SumNum.SumNums:output_type -> sumnum.SumReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_sumnum_proto_init() }
func file_proto_sumnum_proto_init() {
	if File_proto_sumnum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sumnum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumTwoRequest); i {
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
		file_proto_sumnum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumReply); i {
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
		file_proto_sumnum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumsRequest); i {
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
			RawDescriptor: file_proto_sumnum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sumnum_proto_goTypes,
		DependencyIndexes: file_proto_sumnum_proto_depIdxs,
		MessageInfos:      file_proto_sumnum_proto_msgTypes,
	}.Build()
	File_proto_sumnum_proto = out.File
	file_proto_sumnum_proto_rawDesc = nil
	file_proto_sumnum_proto_goTypes = nil
	file_proto_sumnum_proto_depIdxs = nil
}
