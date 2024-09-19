// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: proto/income/income.proto

package income

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

type IncomeEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IncomeEmpty) Reset() {
	*x = IncomeEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_income_income_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomeEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomeEmpty) ProtoMessage() {}

func (x *IncomeEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_income_income_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomeEmpty.ProtoReflect.Descriptor instead.
func (*IncomeEmpty) Descriptor() ([]byte, []int) {
	return file_proto_income_income_proto_rawDescGZIP(), []int{0}
}

type ReportResponseIncome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Income     float64 `protobuf:"fixed64,1,opt,name=income,proto3" json:"income,omitempty"`
	Expenses   float64 `protobuf:"fixed64,2,opt,name=expenses,proto3" json:"expenses,omitempty"`
	NetSavings float64 `protobuf:"fixed64,3,opt,name=netSavings,proto3" json:"netSavings,omitempty"`
}

func (x *ReportResponseIncome) Reset() {
	*x = ReportResponseIncome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_income_income_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResponseIncome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponseIncome) ProtoMessage() {}

func (x *ReportResponseIncome) ProtoReflect() protoreflect.Message {
	mi := &file_proto_income_income_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponseIncome.ProtoReflect.Descriptor instead.
func (*ReportResponseIncome) Descriptor() ([]byte, []int) {
	return file_proto_income_income_proto_rawDescGZIP(), []int{1}
}

func (x *ReportResponseIncome) GetIncome() float64 {
	if x != nil {
		return x.Income
	}
	return 0
}

func (x *ReportResponseIncome) GetExpenses() float64 {
	if x != nil {
		return x.Expenses
	}
	return 0
}

func (x *ReportResponseIncome) GetNetSavings() float64 {
	if x != nil {
		return x.NetSavings
	}
	return 0
}

type TransactionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string  `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Category string  `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Currency string  `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount   float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransactionInfo) Reset() {
	*x = TransactionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_income_income_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionInfo) ProtoMessage() {}

func (x *TransactionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_income_income_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionInfo.ProtoReflect.Descriptor instead.
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return file_proto_income_income_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TransactionInfo) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *TransactionInfo) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransactionInfo) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransactionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TransactionID) Reset() {
	*x = TransactionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_income_income_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionID) ProtoMessage() {}

func (x *TransactionID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_income_income_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionID.ProtoReflect.Descriptor instead.
func (*TransactionID) Descriptor() ([]byte, []int) {
	return file_proto_income_income_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TransactionWithID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type     string  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Category string  `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Currency string  `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount   float64 `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Date     string  `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *TransactionWithID) Reset() {
	*x = TransactionWithID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_income_income_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionWithID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionWithID) ProtoMessage() {}

func (x *TransactionWithID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_income_income_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionWithID.ProtoReflect.Descriptor instead.
func (*TransactionWithID) Descriptor() ([]byte, []int) {
	return file_proto_income_income_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionWithID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TransactionWithID) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TransactionWithID) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *TransactionWithID) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransactionWithID) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransactionWithID) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type TransactionCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *TransactionCategory) Reset() {
	*x = TransactionCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_income_income_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionCategory) ProtoMessage() {}

func (x *TransactionCategory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_income_income_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionCategory.ProtoReflect.Descriptor instead.
func (*TransactionCategory) Descriptor() ([]byte, []int) {
	return file_proto_income_income_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionCategory) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

type ListTransactions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListTransactions []*TransactionWithID `protobuf:"bytes,1,rep,name=listTransactions,proto3" json:"listTransactions,omitempty"`
}

func (x *ListTransactions) Reset() {
	*x = ListTransactions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_income_income_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTransactions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTransactions) ProtoMessage() {}

func (x *ListTransactions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_income_income_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTransactions.ProtoReflect.Descriptor instead.
func (*ListTransactions) Descriptor() ([]byte, []int) {
	return file_proto_income_income_proto_rawDescGZIP(), []int{6}
}

func (x *ListTransactions) GetListTransactions() []*TransactionWithID {
	if x != nil {
		return x.ListTransactions
	}
	return nil
}

type TransactionDate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Start string `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End   string `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *TransactionDate) Reset() {
	*x = TransactionDate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_income_income_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionDate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionDate) ProtoMessage() {}

func (x *TransactionDate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_income_income_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionDate.ProtoReflect.Descriptor instead.
func (*TransactionDate) Descriptor() ([]byte, []int) {
	return file_proto_income_income_proto_rawDescGZIP(), []int{7}
}

func (x *TransactionDate) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TransactionDate) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *TransactionDate) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_income_income_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_income_income_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_proto_income_income_proto_rawDescGZIP(), []int{8}
}

func (x *TransactionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_income_income_proto protoreflect.FileDescriptor

var file_proto_income_income_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x2f, 0x69,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6a, 0x0a, 0x14, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x65, 0x78,
	0x70, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x53, 0x61, 0x76,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x53,
	0x61, 0x76, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x75, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1f, 0x0a,
	0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9b,
	0x01, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69,
	0x74, 0x68, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x31, 0x0a, 0x13,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22,
	0x52, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3e, 0x0a, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x49,
	0x44, 0x52, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x4d, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x6e, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xbc, 0x03, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x0e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x41, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x12, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x1a, 0x14,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x14, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x12, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x12, 0x44, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x1a, 0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x11, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x30, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0c, 0x2e,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_income_income_proto_rawDescOnce sync.Once
	file_proto_income_income_proto_rawDescData = file_proto_income_income_proto_rawDesc
)

func file_proto_income_income_proto_rawDescGZIP() []byte {
	file_proto_income_income_proto_rawDescOnce.Do(func() {
		file_proto_income_income_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_income_income_proto_rawDescData)
	})
	return file_proto_income_income_proto_rawDescData
}

var file_proto_income_income_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_income_income_proto_goTypes = []any{
	(*IncomeEmpty)(nil),          // 0: IncomeEmpty
	(*ReportResponseIncome)(nil), // 1: ReportResponseIncome
	(*TransactionInfo)(nil),      // 2: TransactionInfo
	(*TransactionID)(nil),        // 3: TransactionID
	(*TransactionWithID)(nil),    // 4: TransactionWithID
	(*TransactionCategory)(nil),  // 5: TransactionCategory
	(*ListTransactions)(nil),     // 6: ListTransactions
	(*TransactionDate)(nil),      // 7: TransactionDate
	(*TransactionResponse)(nil),  // 8: TransactionResponse
}
var file_proto_income_income_proto_depIdxs = []int32{
	4, // 0: ListTransactions.listTransactions:type_name -> TransactionWithID
	2, // 1: TransactionService.CreateTransaction:input_type -> TransactionInfo
	4, // 2: TransactionService.UpdateTransactionByID:input_type -> TransactionWithID
	3, // 3: TransactionService.DeleteTransactionByID:input_type -> TransactionID
	3, // 4: TransactionService.GetTransactionByID:input_type -> TransactionID
	5, // 5: TransactionService.GetTransactionsByCategory:input_type -> TransactionCategory
	7, // 6: TransactionService.GetTransactionByDate:input_type -> TransactionDate
	0, // 7: TransactionService.GetReport:input_type -> IncomeEmpty
	3, // 8: TransactionService.CreateTransaction:output_type -> TransactionID
	8, // 9: TransactionService.UpdateTransactionByID:output_type -> TransactionResponse
	8, // 10: TransactionService.DeleteTransactionByID:output_type -> TransactionResponse
	4, // 11: TransactionService.GetTransactionByID:output_type -> TransactionWithID
	6, // 12: TransactionService.GetTransactionsByCategory:output_type -> ListTransactions
	6, // 13: TransactionService.GetTransactionByDate:output_type -> ListTransactions
	1, // 14: TransactionService.GetReport:output_type -> ReportResponseIncome
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_income_income_proto_init() }
func file_proto_income_income_proto_init() {
	if File_proto_income_income_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_income_income_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IncomeEmpty); i {
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
		file_proto_income_income_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReportResponseIncome); i {
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
		file_proto_income_income_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionInfo); i {
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
		file_proto_income_income_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionID); i {
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
		file_proto_income_income_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionWithID); i {
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
		file_proto_income_income_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionCategory); i {
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
		file_proto_income_income_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListTransactions); i {
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
		file_proto_income_income_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionDate); i {
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
		file_proto_income_income_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*TransactionResponse); i {
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
			RawDescriptor: file_proto_income_income_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_income_income_proto_goTypes,
		DependencyIndexes: file_proto_income_income_proto_depIdxs,
		MessageInfos:      file_proto_income_income_proto_msgTypes,
	}.Build()
	File_proto_income_income_proto = out.File
	file_proto_income_income_proto_rawDesc = nil
	file_proto_income_income_proto_goTypes = nil
	file_proto_income_income_proto_depIdxs = nil
}
