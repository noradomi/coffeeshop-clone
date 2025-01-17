// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: common.proto

package gen

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type OrderSource int32

const (
	OrderSource_COUNTER OrderSource = 0
	OrderSource_WEB     OrderSource = 1
)

// Enum value maps for OrderSource.
var (
	OrderSource_name = map[int32]string{
		0: "COUNTER",
		1: "WEB",
	}
	OrderSource_value = map[string]int32{
		"COUNTER": 0,
		"WEB":     1,
	}
)

func (x OrderSource) Enum() *OrderSource {
	p := new(OrderSource)
	*p = x
	return p
}

func (x OrderSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderSource) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (OrderSource) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x OrderSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderSource.Descriptor instead.
func (OrderSource) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_PLACED      Status = 0
	Status_IN_PROGRESS Status = 1
	Status_FULFILLED   Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "PLACED",
		1: "IN_PROGRESS",
		2: "FULFILLED",
	}
	Status_value = map[string]int32{
		"PLACED":      0,
		"IN_PROGRESS": 1,
		"FULFILLED":   2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

type Location int32

const (
	Location_ATLANTA   Location = 0
	Location_CHARLOTTE Location = 1
	Location_RALEIGH   Location = 2
)

// Enum value maps for Location.
var (
	Location_name = map[int32]string{
		0: "ATLANTA",
		1: "CHARLOTTE",
		2: "RALEIGH",
	}
	Location_value = map[string]int32{
		"ATLANTA":   0,
		"CHARLOTTE": 1,
		"RALEIGH":   2,
	}
)

func (x Location) Enum() *Location {
	p := new(Location)
	*p = x
	return p
}

func (x Location) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Location) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (Location) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x Location) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Location.Descriptor instead.
func (Location) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

type ItemType int32

const (
	// Beverages
	ItemType_CAPPUCCINO       ItemType = 0
	ItemType_COFFEE_BLACK     ItemType = 1
	ItemType_COFFEE_WITH_ROOM ItemType = 2
	ItemType_ESPRESSO         ItemType = 3
	ItemType_ESPRESSO_DOUBLE  ItemType = 4
	ItemType_LATTE            ItemType = 5
	// Food
	ItemType_CAKEPOP             ItemType = 6
	ItemType_CROISSANT           ItemType = 7
	ItemType_MUFFIN              ItemType = 8
	ItemType_CROISSANT_CHOCOLATE ItemType = 9
)

// Enum value maps for ItemType.
var (
	ItemType_name = map[int32]string{
		0: "CAPPUCCINO",
		1: "COFFEE_BLACK",
		2: "COFFEE_WITH_ROOM",
		3: "ESPRESSO",
		4: "ESPRESSO_DOUBLE",
		5: "LATTE",
		6: "CAKEPOP",
		7: "CROISSANT",
		8: "MUFFIN",
		9: "CROISSANT_CHOCOLATE",
	}
	ItemType_value = map[string]int32{
		"CAPPUCCINO":          0,
		"COFFEE_BLACK":        1,
		"COFFEE_WITH_ROOM":    2,
		"ESPRESSO":            3,
		"ESPRESSO_DOUBLE":     4,
		"LATTE":               5,
		"CAKEPOP":             6,
		"CROISSANT":           7,
		"MUFFIN":              8,
		"CROISSANT_CHOCOLATE": 9,
	}
)

func (x ItemType) Enum() *ItemType {
	p := new(ItemType)
	*p = x
	return p
}

func (x ItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[3].Descriptor()
}

func (ItemType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[3]
}

func (x ItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemType.Descriptor instead.
func (ItemType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

type CommandType int32

const (
	CommandType_PLACE_ORDER CommandType = 0
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0: "PLACE_ORDER",
	}
	CommandType_value = map[string]int32{
		"PLACE_ORDER": 0,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[4].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[4]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x23, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x45, 0x52, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42, 0x10, 0x01, 0x2a, 0x34, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x4c, 0x41, 0x43, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x55, 0x4c, 0x46, 0x49, 0x4c, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x2a, 0x33, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0b, 0x0a, 0x07, 0x41, 0x54, 0x4c, 0x41, 0x4e, 0x54, 0x41, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x48, 0x41, 0x52, 0x4c, 0x4f, 0x54, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52,
	0x41, 0x4c, 0x45, 0x49, 0x47, 0x48, 0x10, 0x02, 0x2a, 0xb1, 0x01, 0x0a, 0x08, 0x49, 0x74, 0x65,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x50, 0x50, 0x55, 0x43, 0x43,
	0x49, 0x4e, 0x4f, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x46, 0x46, 0x45, 0x45, 0x5f,
	0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x46, 0x46, 0x45,
	0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x45, 0x53, 0x50, 0x52, 0x45, 0x53, 0x53, 0x4f, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x53, 0x50, 0x52, 0x45, 0x53, 0x53, 0x4f, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x04,
	0x12, 0x09, 0x0a, 0x05, 0x4c, 0x41, 0x54, 0x54, 0x45, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x41, 0x4b, 0x45, 0x50, 0x4f, 0x50, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x52, 0x4f, 0x49,
	0x53, 0x53, 0x41, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x55, 0x46, 0x46, 0x49,
	0x4e, 0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x52, 0x4f, 0x49, 0x53, 0x53, 0x41, 0x4e, 0x54,
	0x5f, 0x43, 0x48, 0x4f, 0x43, 0x4f, 0x4c, 0x41, 0x54, 0x45, 0x10, 0x09, 0x2a, 0x1e, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x50,
	0x4c, 0x41, 0x43, 0x45, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x00, 0x42, 0x87, 0x01, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x72, 0x61,
	0x64, 0x6f, 0x6d, 0x69, 0x2f, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2d,
	0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x92,
	0x41, 0x54, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x72, 0x48, 0x0a, 0x18,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x20, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x72,
	0x61, 0x64, 0x6f, 0x6d, 0x69, 0x2f, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70,
	0x2d, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_common_proto_goTypes = []interface{}{
	(OrderSource)(0), // 0: go.coffeeshop.proto.common.OrderSource
	(Status)(0),      // 1: go.coffeeshop.proto.common.Status
	(Location)(0),    // 2: go.coffeeshop.proto.common.Location
	(ItemType)(0),    // 3: go.coffeeshop.proto.common.ItemType
	(CommandType)(0), // 4: go.coffeeshop.proto.common.CommandType
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
