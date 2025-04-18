// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: proto/shortlink.proto

package shortlink

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortlink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortlink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_proto_shortlink_proto_rawDescGZIP(), []int{0}
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortlink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortlink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_proto_shortlink_proto_rawDescGZIP(), []int{1}
}

func (x *Link) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Link) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type Links struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link []*Link `protobuf:"bytes,1,rep,name=link,proto3" json:"link,omitempty"`
}

func (x *Links) Reset() {
	*x = Links{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortlink_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Links) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Links) ProtoMessage() {}

func (x *Links) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortlink_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Links.ProtoReflect.Descriptor instead.
func (*Links) Descriptor() ([]byte, []int) {
	return file_proto_shortlink_proto_rawDescGZIP(), []int{2}
}

func (x *Links) GetLink() []*Link {
	if x != nil {
		return x.Link
	}
	return nil
}

type LimitOffset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *LimitOffset) Reset() {
	*x = LimitOffset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shortlink_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitOffset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitOffset) ProtoMessage() {}

func (x *LimitOffset) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shortlink_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitOffset.ProtoReflect.Descriptor instead.
func (*LimitOffset) Descriptor() ([]byte, []int) {
	return file_proto_shortlink_proto_rawDescGZIP(), []int{3}
}

func (x *LimitOffset) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *LimitOffset) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_proto_shortlink_proto protoreflect.FileDescriptor

var file_proto_shortlink_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x3c, 0x0a,
	0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x2f, 0x0a, 0x05, 0x4c,
	0x69, 0x6e, 0x6b, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x3b, 0x0a, 0x0b,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32, 0xa6, 0x03, 0x0a, 0x09, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x12, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x3d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x12,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x12, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x12,
	0x38, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x12, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x3a, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x12, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x12, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_shortlink_proto_rawDescOnce sync.Once
	file_proto_shortlink_proto_rawDescData = file_proto_shortlink_proto_rawDesc
)

func file_proto_shortlink_proto_rawDescGZIP() []byte {
	file_proto_shortlink_proto_rawDescOnce.Do(func() {
		file_proto_shortlink_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shortlink_proto_rawDescData)
	})
	return file_proto_shortlink_proto_rawDescData
}

var file_proto_shortlink_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_shortlink_proto_goTypes = []any{
	(*Void)(nil),                   // 0: shortlink.v1.Void
	(*Link)(nil),                   // 1: shortlink.v1.Link
	(*Links)(nil),                  // 2: shortlink.v1.Links
	(*LimitOffset)(nil),            // 3: shortlink.v1.LimitOffset
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*wrapperspb.UInt64Value)(nil), // 5: google.protobuf.UInt64Value
}
var file_proto_shortlink_proto_depIdxs = []int32{
	1, // 0: shortlink.v1.Links.link:type_name -> shortlink.v1.Link
	4, // 1: shortlink.v1.ShortLink.Create:input_type -> google.protobuf.StringValue
	4, // 2: shortlink.v1.ShortLink.GetByHash:input_type -> google.protobuf.StringValue
	5, // 3: shortlink.v1.ShortLink.GetById:input_type -> google.protobuf.UInt64Value
	3, // 4: shortlink.v1.ShortLink.GetAll:input_type -> shortlink.v1.LimitOffset
	1, // 5: shortlink.v1.ShortLink.Update:input_type -> shortlink.v1.Link
	5, // 6: shortlink.v1.ShortLink.Delete:input_type -> google.protobuf.UInt64Value
	0, // 7: shortlink.v1.ShortLink.Count:input_type -> shortlink.v1.Void
	1, // 8: shortlink.v1.ShortLink.Create:output_type -> shortlink.v1.Link
	1, // 9: shortlink.v1.ShortLink.GetByHash:output_type -> shortlink.v1.Link
	1, // 10: shortlink.v1.ShortLink.GetById:output_type -> shortlink.v1.Link
	2, // 11: shortlink.v1.ShortLink.GetAll:output_type -> shortlink.v1.Links
	1, // 12: shortlink.v1.ShortLink.Update:output_type -> shortlink.v1.Link
	0, // 13: shortlink.v1.ShortLink.Delete:output_type -> shortlink.v1.Void
	5, // 14: shortlink.v1.ShortLink.Count:output_type -> google.protobuf.UInt64Value
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_shortlink_proto_init() }
func file_proto_shortlink_proto_init() {
	if File_proto_shortlink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shortlink_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Void); i {
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
		file_proto_shortlink_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Link); i {
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
		file_proto_shortlink_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Links); i {
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
		file_proto_shortlink_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LimitOffset); i {
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
			RawDescriptor: file_proto_shortlink_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shortlink_proto_goTypes,
		DependencyIndexes: file_proto_shortlink_proto_depIdxs,
		MessageInfos:      file_proto_shortlink_proto_msgTypes,
	}.Build()
	File_proto_shortlink_proto = out.File
	file_proto_shortlink_proto_rawDesc = nil
	file_proto_shortlink_proto_goTypes = nil
	file_proto_shortlink_proto_depIdxs = nil
}
