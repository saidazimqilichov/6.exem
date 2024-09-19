// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: proto/notification/notification.proto

package notification

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

type NotifyEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyEmpty) Reset() {
	*x = NotifyEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notification_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyEmpty) ProtoMessage() {}

func (x *NotifyEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyEmpty.ProtoReflect.Descriptor instead.
func (*NotifyEmpty) Descriptor() ([]byte, []int) {
	return file_proto_notification_notification_proto_rawDescGZIP(), []int{0}
}

type Notify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message string  `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Report  *Report `protobuf:"bytes,3,opt,name=report,proto3" json:"report,omitempty"`
	Date    string  `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Notify) Reset() {
	*x = Notify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notification_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notify) ProtoMessage() {}

func (x *Notify) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notify.ProtoReflect.Descriptor instead.
func (*Notify) Descriptor() ([]byte, []int) {
	return file_proto_notification_notification_proto_rawDescGZIP(), []int{1}
}

func (x *Notify) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Notify) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Notify) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

func (x *Notify) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Income     float64 `protobuf:"fixed64,1,opt,name=income,proto3" json:"income,omitempty"`
	Expenses   float64 `protobuf:"fixed64,2,opt,name=expenses,proto3" json:"expenses,omitempty"`
	NetSavings float64 `protobuf:"fixed64,3,opt,name=netSavings,proto3" json:"netSavings,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notification_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_proto_notification_notification_proto_rawDescGZIP(), []int{2}
}

func (x *Report) GetIncome() float64 {
	if x != nil {
		return x.Income
	}
	return 0
}

func (x *Report) GetExpenses() float64 {
	if x != nil {
		return x.Expenses
	}
	return 0
}

func (x *Report) GetNetSavings() float64 {
	if x != nil {
		return x.NetSavings
	}
	return 0
}

type NotifyList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotifyList []*Notify `protobuf:"bytes,1,rep,name=notifyList,proto3" json:"notifyList,omitempty"`
}

func (x *NotifyList) Reset() {
	*x = NotifyList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notification_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyList) ProtoMessage() {}

func (x *NotifyList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notification_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyList.ProtoReflect.Descriptor instead.
func (*NotifyList) Descriptor() ([]byte, []int) {
	return file_proto_notification_notification_proto_rawDescGZIP(), []int{3}
}

func (x *NotifyList) GetNotifyList() []*Notify {
	if x != nil {
		return x.NotifyList
	}
	return nil
}

var File_proto_notification_notification_proto protoreflect.FileDescriptor

var file_proto_notification_notification_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x70, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x5c, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x65, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x53, 0x61, 0x76,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x53,
	0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x35, 0x0a, 0x0a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x76, 0x0a,
	0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x4e, 0x6f,
	0x74, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0c, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_notification_notification_proto_rawDescOnce sync.Once
	file_proto_notification_notification_proto_rawDescData = file_proto_notification_notification_proto_rawDesc
)

func file_proto_notification_notification_proto_rawDescGZIP() []byte {
	file_proto_notification_notification_proto_rawDescOnce.Do(func() {
		file_proto_notification_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_notification_notification_proto_rawDescData)
	})
	return file_proto_notification_notification_proto_rawDescData
}

var file_proto_notification_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_notification_notification_proto_goTypes = []any{
	(*NotifyEmpty)(nil), // 0: NotifyEmpty
	(*Notify)(nil),      // 1: Notify
	(*Report)(nil),      // 2: Report
	(*NotifyList)(nil),  // 3: NotifyList
}
var file_proto_notification_notification_proto_depIdxs = []int32{
	2, // 0: Notify.report:type_name -> Report
	1, // 1: NotifyList.notifyList:type_name -> Notify
	0, // 2: NotificationService.GetNotfication:input_type -> NotifyEmpty
	0, // 3: NotificationService.GetUnreadNotfications:input_type -> NotifyEmpty
	3, // 4: NotificationService.GetNotfication:output_type -> NotifyList
	3, // 5: NotificationService.GetUnreadNotfications:output_type -> NotifyList
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_notification_notification_proto_init() }
func file_proto_notification_notification_proto_init() {
	if File_proto_notification_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_notification_notification_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyEmpty); i {
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
		file_proto_notification_notification_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Notify); i {
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
		file_proto_notification_notification_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Report); i {
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
		file_proto_notification_notification_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*NotifyList); i {
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
			RawDescriptor: file_proto_notification_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_notification_notification_proto_goTypes,
		DependencyIndexes: file_proto_notification_notification_proto_depIdxs,
		MessageInfos:      file_proto_notification_notification_proto_msgTypes,
	}.Build()
	File_proto_notification_notification_proto = out.File
	file_proto_notification_notification_proto_rawDesc = nil
	file_proto_notification_notification_proto_goTypes = nil
	file_proto_notification_notification_proto_depIdxs = nil
}