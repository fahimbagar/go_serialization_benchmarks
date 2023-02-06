// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: structdef.proto

package goserbench

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

type ProtoBufA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	BirthDay *int64   `protobuf:"varint,2,req,name=birthDay" json:"birthDay,omitempty"`
	Phone    *string  `protobuf:"bytes,3,req,name=phone" json:"phone,omitempty"`
	Siblings *int32   `protobuf:"varint,4,req,name=siblings" json:"siblings,omitempty"`
	Spouse   *bool    `protobuf:"varint,5,req,name=spouse" json:"spouse,omitempty"`
	Money    *float64 `protobuf:"fixed64,6,req,name=money" json:"money,omitempty"`
}

func (x *ProtoBufA) Reset() {
	*x = ProtoBufA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_structdef_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoBufA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoBufA) ProtoMessage() {}

func (x *ProtoBufA) ProtoReflect() protoreflect.Message {
	mi := &file_structdef_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoBufA.ProtoReflect.Descriptor instead.
func (*ProtoBufA) Descriptor() ([]byte, []int) {
	return file_structdef_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoBufA) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ProtoBufA) GetBirthDay() int64 {
	if x != nil && x.BirthDay != nil {
		return *x.BirthDay
	}
	return 0
}

func (x *ProtoBufA) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *ProtoBufA) GetSiblings() int32 {
	if x != nil && x.Siblings != nil {
		return *x.Siblings
	}
	return 0
}

func (x *ProtoBufA) GetSpouse() bool {
	if x != nil && x.Spouse != nil {
		return *x.Spouse
	}
	return false
}

func (x *ProtoBufA) GetMoney() float64 {
	if x != nil && x.Money != nil {
		return *x.Money
	}
	return 0
}

var File_structdef_proto protoreflect.FileDescriptor

var file_structdef_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x73, 0x65, 0x72, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x22, 0x9b, 0x01,
	0x0a, 0x09, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x41, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x03, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x08, 0x73, 0x69, 0x62, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x70, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x70, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06,
	0x20, 0x02, 0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x42, 0x3e, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x63, 0x74, 0x68,
	0x6f, 0x6d, 0x61, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x3b, 0x67, 0x6f, 0x73, 0x65, 0x72, 0x62, 0x65, 0x6e, 0x63, 0x68,
}

var (
	file_structdef_proto_rawDescOnce sync.Once
	file_structdef_proto_rawDescData = file_structdef_proto_rawDesc
)

func file_structdef_proto_rawDescGZIP() []byte {
	file_structdef_proto_rawDescOnce.Do(func() {
		file_structdef_proto_rawDescData = protoimpl.X.CompressGZIP(file_structdef_proto_rawDescData)
	})
	return file_structdef_proto_rawDescData
}

var file_structdef_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_structdef_proto_goTypes = []interface{}{
	(*ProtoBufA)(nil), // 0: goserbench.ProtoBufA
}
var file_structdef_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_structdef_proto_init() }
func file_structdef_proto_init() {
	if File_structdef_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_structdef_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoBufA); i {
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
			RawDescriptor: file_structdef_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_structdef_proto_goTypes,
		DependencyIndexes: file_structdef_proto_depIdxs,
		MessageInfos:      file_structdef_proto_msgTypes,
	}.Build()
	File_structdef_proto = out.File
	file_structdef_proto_rawDesc = nil
	file_structdef_proto_goTypes = nil
	file_structdef_proto_depIdxs = nil
}
