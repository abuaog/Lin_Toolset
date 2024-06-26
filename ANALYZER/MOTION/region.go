// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: motion.proto

package motion

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows    int32  `protobuf:"varint,1,opt,name=rows,proto3" json:"rows,omitempty"`
	Columns int32  `protobuf:"varint,2,opt,name=columns,proto3" json:"columns,omitempty"`
	Cells   string `protobuf:"bytes,3,opt,name=cells,proto3" json:"cells,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_motion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_motion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_motion_proto_rawDescGZIP(), []int{0}
}

func (x *Region) GetRows() int32 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *Region) GetColumns() int32 {
	if x != nil {
		return x.Columns
	}
	return 0
}

func (x *Region) GetCells() string {
	if x != nil {
		return x.Cells
	}
	return ""
}

var File_motion_proto protoreflect.FileDescriptor

var file_motion_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63,
	0x65, 0x6c, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_motion_proto_rawDescOnce sync.Once
	file_motion_proto_rawDescData = file_motion_proto_rawDesc
)

func file_motion_proto_rawDescGZIP() []byte {
	file_motion_proto_rawDescOnce.Do(func() {
		file_motion_proto_rawDescData = protoimpl.X.CompressGZIP(file_motion_proto_rawDescData)
	})
	return file_motion_proto_rawDescData
}

var (
	file_motion_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
	file_motion_proto_goTypes  = []interface{}{
		(*Region)(nil), // 0: motion.Region
	}
)

var file_motion_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_motion_proto_init() }
func file_motion_proto_init() {
	if File_motion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_motion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
			RawDescriptor: file_motion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_motion_proto_goTypes,
		DependencyIndexes: file_motion_proto_depIdxs,
		MessageInfos:      file_motion_proto_msgTypes,
	}.Build()
	File_motion_proto = out.File
	file_motion_proto_rawDesc = nil
	file_motion_proto_goTypes = nil
	file_motion_proto_depIdxs = nil
}
