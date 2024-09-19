// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: proto/report/report.proto

package report

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

type ReportEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportEmpty) Reset() {
	*x = ReportEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_report_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEmpty) ProtoMessage() {}

func (x *ReportEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEmpty.ProtoReflect.Descriptor instead.
func (*ReportEmpty) Descriptor() ([]byte, []int) {
	return file_proto_report_report_proto_rawDescGZIP(), []int{0}
}

type ReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Income     float64 `protobuf:"fixed64,1,opt,name=income,proto3" json:"income,omitempty"`
	Expenses   float64 `protobuf:"fixed64,2,opt,name=expenses,proto3" json:"expenses,omitempty"`
	NetSavings float64 `protobuf:"fixed64,3,opt,name=netSavings,proto3" json:"netSavings,omitempty"`
	Budget     *Budget `protobuf:"bytes,4,opt,name=budget,proto3" json:"budget,omitempty"`
}

func (x *ReportResponse) Reset() {
	*x = ReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_report_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponse.ProtoReflect.Descriptor instead.
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return file_proto_report_report_proto_rawDescGZIP(), []int{1}
}

func (x *ReportResponse) GetIncome() float64 {
	if x != nil {
		return x.Income
	}
	return 0
}

func (x *ReportResponse) GetExpenses() float64 {
	if x != nil {
		return x.Expenses
	}
	return 0
}

func (x *ReportResponse) GetNetSavings() float64 {
	if x != nil {
		return x.NetSavings
	}
	return 0
}

func (x *ReportResponse) GetBudget() *Budget {
	if x != nil {
		return x.Budget
	}
	return nil
}

type Budget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalAmount     float64 `protobuf:"fixed64,1,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	TotalSpent      float64 `protobuf:"fixed64,2,opt,name=totalSpent,proto3" json:"totalSpent,omitempty"`
	RemainingBudget float64 `protobuf:"fixed64,3,opt,name=remainingBudget,proto3" json:"remainingBudget,omitempty"`
}

func (x *Budget) Reset() {
	*x = Budget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_report_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Budget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budget) ProtoMessage() {}

func (x *Budget) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Budget.ProtoReflect.Descriptor instead.
func (*Budget) Descriptor() ([]byte, []int) {
	return file_proto_report_report_proto_rawDescGZIP(), []int{2}
}

func (x *Budget) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Budget) GetTotalSpent() float64 {
	if x != nil {
		return x.TotalSpent
	}
	return 0
}

func (x *Budget) GetRemainingBudget() float64 {
	if x != nil {
		return x.RemainingBudget
	}
	return 0
}

type ReportID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReportID) Reset() {
	*x = ReportID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_report_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportID) ProtoMessage() {}

func (x *ReportID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportID.ProtoReflect.Descriptor instead.
func (*ReportID) Descriptor() ([]byte, []int) {
	return file_proto_report_report_proto_rawDescGZIP(), []int{3}
}

func (x *ReportID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_proto_report_report_proto protoreflect.FileDescriptor

var file_proto_report_report_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x85, 0x01, 0x0a, 0x0e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x69,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x53, 0x61, 0x76, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x1f, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x22, 0x74, 0x0a, 0x06, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x6e, 0x74, 0x12, 0x28,
	0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x32, 0x3b, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x0c, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0f, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_report_report_proto_rawDescOnce sync.Once
	file_proto_report_report_proto_rawDescData = file_proto_report_report_proto_rawDesc
)

func file_proto_report_report_proto_rawDescGZIP() []byte {
	file_proto_report_report_proto_rawDescOnce.Do(func() {
		file_proto_report_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_report_report_proto_rawDescData)
	})
	return file_proto_report_report_proto_rawDescData
}

var file_proto_report_report_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_report_report_proto_goTypes = []any{
	(*ReportEmpty)(nil),    // 0: ReportEmpty
	(*ReportResponse)(nil), // 1: ReportResponse
	(*Budget)(nil),         // 2: Budget
	(*ReportID)(nil),       // 3: ReportID
}
var file_proto_report_report_proto_depIdxs = []int32{
	2, // 0: ReportResponse.budget:type_name -> Budget
	0, // 1: ReportService.GetReport:input_type -> ReportEmpty
	1, // 2: ReportService.GetReport:output_type -> ReportResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_report_report_proto_init() }
func file_proto_report_report_proto_init() {
	if File_proto_report_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_report_report_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ReportEmpty); i {
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
		file_proto_report_report_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReportResponse); i {
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
		file_proto_report_report_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Budget); i {
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
		file_proto_report_report_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ReportID); i {
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
			RawDescriptor: file_proto_report_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_report_report_proto_goTypes,
		DependencyIndexes: file_proto_report_report_proto_depIdxs,
		MessageInfos:      file_proto_report_report_proto_msgTypes,
	}.Build()
	File_proto_report_report_proto = out.File
	file_proto_report_report_proto_rawDesc = nil
	file_proto_report_report_proto_goTypes = nil
	file_proto_report_report_proto_depIdxs = nil
}
