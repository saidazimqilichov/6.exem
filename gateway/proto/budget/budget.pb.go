// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: proto/budget/budget.proto

package budget

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

type SpentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SpentRequest) Reset() {
	*x = SpentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_budget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpentRequest) ProtoMessage() {}

func (x *SpentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_budget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpentRequest.ProtoReflect.Descriptor instead.
func (*SpentRequest) Descriptor() ([]byte, []int) {
	return file_proto_budget_budget_proto_rawDescGZIP(), []int{0}
}

func (x *SpentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SpentRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BudgetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string  `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Amount   float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Spent    float64 `protobuf:"fixed64,3,opt,name=spent,proto3" json:"spent,omitempty"`
	Currency string  `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *BudgetInfo) Reset() {
	*x = BudgetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_budget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetInfo) ProtoMessage() {}

func (x *BudgetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_budget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetInfo.ProtoReflect.Descriptor instead.
func (*BudgetInfo) Descriptor() ([]byte, []int) {
	return file_proto_budget_budget_proto_rawDescGZIP(), []int{1}
}

func (x *BudgetInfo) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *BudgetInfo) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BudgetInfo) GetSpent() float64 {
	if x != nil {
		return x.Spent
	}
	return 0
}

func (x *BudgetInfo) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type BudgetID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BudgetID) Reset() {
	*x = BudgetID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_budget_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetID) ProtoMessage() {}

func (x *BudgetID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_budget_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetID.ProtoReflect.Descriptor instead.
func (*BudgetID) Descriptor() ([]byte, []int) {
	return file_proto_budget_budget_proto_rawDescGZIP(), []int{2}
}

func (x *BudgetID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BudgetWithID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Category string  `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Amount   float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Spent    float64 `protobuf:"fixed64,4,opt,name=spent,proto3" json:"spent,omitempty"`
	Currency string  `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *BudgetWithID) Reset() {
	*x = BudgetWithID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_budget_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetWithID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetWithID) ProtoMessage() {}

func (x *BudgetWithID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_budget_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetWithID.ProtoReflect.Descriptor instead.
func (*BudgetWithID) Descriptor() ([]byte, []int) {
	return file_proto_budget_budget_proto_rawDescGZIP(), []int{3}
}

func (x *BudgetWithID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BudgetWithID) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *BudgetWithID) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BudgetWithID) GetSpent() float64 {
	if x != nil {
		return x.Spent
	}
	return 0
}

func (x *BudgetWithID) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type BudgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BudgetResponse) Reset() {
	*x = BudgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_budget_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetResponse) ProtoMessage() {}

func (x *BudgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_budget_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetResponse.ProtoReflect.Descriptor instead.
func (*BudgetResponse) Descriptor() ([]byte, []int) {
	return file_proto_budget_budget_proto_rawDescGZIP(), []int{4}
}

func (x *BudgetResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BudgetUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *BudgetUpdate) Reset() {
	*x = BudgetUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_budget_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetUpdate) ProtoMessage() {}

func (x *BudgetUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_budget_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetUpdate.ProtoReflect.Descriptor instead.
func (*BudgetUpdate) Descriptor() ([]byte, []int) {
	return file_proto_budget_budget_proto_rawDescGZIP(), []int{5}
}

func (x *BudgetUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BudgetUpdate) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BudgetReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalBudget     float64 `protobuf:"fixed64,1,opt,name=total_budget,json=totalBudget,proto3" json:"total_budget,omitempty"`
	TotalSpent      float64 `protobuf:"fixed64,2,opt,name=total_spent,json=totalSpent,proto3" json:"total_spent,omitempty"`
	RemainingBudget float64 `protobuf:"fixed64,3,opt,name=remaining_budget,json=remainingBudget,proto3" json:"remaining_budget,omitempty"`
}

func (x *BudgetReport) Reset() {
	*x = BudgetReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_budget_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BudgetReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BudgetReport) ProtoMessage() {}

func (x *BudgetReport) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_budget_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BudgetReport.ProtoReflect.Descriptor instead.
func (*BudgetReport) Descriptor() ([]byte, []int) {
	return file_proto_budget_budget_proto_rawDescGZIP(), []int{6}
}

func (x *BudgetReport) GetTotalBudget() float64 {
	if x != nil {
		return x.TotalBudget
	}
	return 0
}

func (x *BudgetReport) GetTotalSpent() float64 {
	if x != nil {
		return x.TotalSpent
	}
	return 0
}

func (x *BudgetReport) GetRemainingBudget() float64 {
	if x != nil {
		return x.RemainingBudget
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_budget_budget_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_budget_budget_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_budget_budget_proto_rawDescGZIP(), []int{7}
}

var File_proto_budget_budget_proto protoreflect.FileDescriptor

var file_proto_budget_budget_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2f, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0c, 0x53,
	0x70, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x72, 0x0a, 0x0a, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x1a, 0x0a, 0x08, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x70, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x2a, 0x0a, 0x0e, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x36, 0x0a, 0x0c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7d,
	0x0a, 0x0c, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x70, 0x65,
	0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x72, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xa1, 0x02, 0x0a, 0x0d, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49, 0x44,
	0x12, 0x34, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0f, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x49, 0x44, 0x30, 0x01, 0x12, 0x2e, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x09, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x49, 0x44, 0x1a, 0x0f, 0x2e, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x30, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x53,
	0x70, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x53, 0x70, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_budget_budget_proto_rawDescOnce sync.Once
	file_proto_budget_budget_proto_rawDescData = file_proto_budget_budget_proto_rawDesc
)

func file_proto_budget_budget_proto_rawDescGZIP() []byte {
	file_proto_budget_budget_proto_rawDescOnce.Do(func() {
		file_proto_budget_budget_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_budget_budget_proto_rawDescData)
	})
	return file_proto_budget_budget_proto_rawDescData
}

var file_proto_budget_budget_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_budget_budget_proto_goTypes = []any{
	(*SpentRequest)(nil),   // 0: SpentRequest
	(*BudgetInfo)(nil),     // 1: BudgetInfo
	(*BudgetID)(nil),       // 2: BudgetID
	(*BudgetWithID)(nil),   // 3: BudgetWithID
	(*BudgetResponse)(nil), // 4: BudgetResponse
	(*BudgetUpdate)(nil),   // 5: BudgetUpdate
	(*BudgetReport)(nil),   // 6: BudgetReport
	(*Empty)(nil),          // 7: Empty
}
var file_proto_budget_budget_proto_depIdxs = []int32{
	1, // 0: BudgetService.CreateBudget:input_type -> BudgetInfo
	5, // 1: BudgetService.UpdateBudgetAmount:input_type -> BudgetUpdate
	7, // 2: BudgetService.GetBudgets:input_type -> Empty
	2, // 3: BudgetService.DeleteBudgetByID:input_type -> BudgetID
	7, // 4: BudgetService.GetBudgetReports:input_type -> Empty
	0, // 5: BudgetService.AddSpentAmount:input_type -> SpentRequest
	2, // 6: BudgetService.CreateBudget:output_type -> BudgetID
	4, // 7: BudgetService.UpdateBudgetAmount:output_type -> BudgetResponse
	3, // 8: BudgetService.GetBudgets:output_type -> BudgetWithID
	4, // 9: BudgetService.DeleteBudgetByID:output_type -> BudgetResponse
	6, // 10: BudgetService.GetBudgetReports:output_type -> BudgetReport
	4, // 11: BudgetService.AddSpentAmount:output_type -> BudgetResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_budget_budget_proto_init() }
func file_proto_budget_budget_proto_init() {
	if File_proto_budget_budget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_budget_budget_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SpentRequest); i {
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
		file_proto_budget_budget_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BudgetInfo); i {
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
		file_proto_budget_budget_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BudgetID); i {
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
		file_proto_budget_budget_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BudgetWithID); i {
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
		file_proto_budget_budget_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BudgetResponse); i {
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
		file_proto_budget_budget_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*BudgetUpdate); i {
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
		file_proto_budget_budget_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*BudgetReport); i {
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
		file_proto_budget_budget_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_proto_budget_budget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_budget_budget_proto_goTypes,
		DependencyIndexes: file_proto_budget_budget_proto_depIdxs,
		MessageInfos:      file_proto_budget_budget_proto_msgTypes,
	}.Build()
	File_proto_budget_budget_proto = out.File
	file_proto_budget_budget_proto_rawDesc = nil
	file_proto_budget_budget_proto_goTypes = nil
	file_proto_budget_budget_proto_depIdxs = nil
}