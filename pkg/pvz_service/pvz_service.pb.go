// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pvz_service/pvz_service.proto

package pvz_service

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Reception_ReceptionStatus int32

const (
	Reception_RECEPTION_STATUS_IN_PROGRESS Reception_ReceptionStatus = 0
	Reception_RECEPTION_STATUS_CLOSED      Reception_ReceptionStatus = 1
)

// Enum value maps for Reception_ReceptionStatus.
var (
	Reception_ReceptionStatus_name = map[int32]string{
		0: "RECEPTION_STATUS_IN_PROGRESS",
		1: "RECEPTION_STATUS_CLOSED",
	}
	Reception_ReceptionStatus_value = map[string]int32{
		"RECEPTION_STATUS_IN_PROGRESS": 0,
		"RECEPTION_STATUS_CLOSED":      1,
	}
)

func (x Reception_ReceptionStatus) Enum() *Reception_ReceptionStatus {
	p := new(Reception_ReceptionStatus)
	*p = x
	return p
}

func (x Reception_ReceptionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reception_ReceptionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pvz_service_pvz_service_proto_enumTypes[0].Descriptor()
}

func (Reception_ReceptionStatus) Type() protoreflect.EnumType {
	return &file_pvz_service_pvz_service_proto_enumTypes[0]
}

func (x Reception_ReceptionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reception_ReceptionStatus.Descriptor instead.
func (Reception_ReceptionStatus) EnumDescriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{1, 0}
}

type PVZ struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	IdPVZ            string                 `protobuf:"bytes,1,opt,name=idPVZ,proto3" json:"idPVZ,omitempty"`
	RegistrationDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=registration_date,json=registrationDate,proto3" json:"registration_date,omitempty"`
	City             string                 `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *PVZ) Reset() {
	*x = PVZ{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PVZ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PVZ) ProtoMessage() {}

func (x *PVZ) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PVZ.ProtoReflect.Descriptor instead.
func (*PVZ) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{0}
}

func (x *PVZ) GetIdPVZ() string {
	if x != nil {
		return x.IdPVZ
	}
	return ""
}

func (x *PVZ) GetRegistrationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RegistrationDate
	}
	return nil
}

func (x *PVZ) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type Reception struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	IdReception   string                    `protobuf:"bytes,1,opt,name=idReception,proto3" json:"idReception,omitempty"`
	OpeningTime   *timestamppb.Timestamp    `protobuf:"bytes,2,opt,name=openingTime,proto3" json:"openingTime,omitempty"`
	ClosingTime   *timestamppb.Timestamp    `protobuf:"bytes,3,opt,name=closingTime,proto3" json:"closingTime,omitempty"`
	IdPVZ         string                    `protobuf:"bytes,4,opt,name=idPVZ,proto3" json:"idPVZ,omitempty"`
	Items         []*Item                   `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	Status        Reception_ReceptionStatus `protobuf:"varint,6,opt,name=status,proto3,enum=pvz_service.Reception_ReceptionStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Reception) Reset() {
	*x = Reception{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Reception) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reception) ProtoMessage() {}

func (x *Reception) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reception.ProtoReflect.Descriptor instead.
func (*Reception) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{1}
}

func (x *Reception) GetIdReception() string {
	if x != nil {
		return x.IdReception
	}
	return ""
}

func (x *Reception) GetOpeningTime() *timestamppb.Timestamp {
	if x != nil {
		return x.OpeningTime
	}
	return nil
}

func (x *Reception) GetClosingTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ClosingTime
	}
	return nil
}

func (x *Reception) GetIdPVZ() string {
	if x != nil {
		return x.IdPVZ
	}
	return ""
}

func (x *Reception) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Reception) GetStatus() Reception_ReceptionStatus {
	if x != nil {
		return x.Status
	}
	return Reception_RECEPTION_STATUS_IN_PROGRESS
}

type Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IdItem        string                 `protobuf:"bytes,1,opt,name=idItem,proto3" json:"idItem,omitempty"`
	ReceptionTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=receptionTime,proto3" json:"receptionTime,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetIdItem() string {
	if x != nil {
		return x.IdItem
	}
	return ""
}

func (x *Item) GetReceptionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceptionTime
	}
	return nil
}

func (x *Item) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role          string                 `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CreatePVZRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	City          string                 `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePVZRequest) Reset() {
	*x = CreatePVZRequest{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePVZRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePVZRequest) ProtoMessage() {}

func (x *CreatePVZRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePVZRequest.ProtoReflect.Descriptor instead.
func (*CreatePVZRequest) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePVZRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreatePVZRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type CreatePVZResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IdPVZ         string                 `protobuf:"bytes,1,opt,name=idPVZ,proto3" json:"idPVZ,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePVZResponse) Reset() {
	*x = CreatePVZResponse{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePVZResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePVZResponse) ProtoMessage() {}

func (x *CreatePVZResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePVZResponse.ProtoReflect.Descriptor instead.
func (*CreatePVZResponse) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreatePVZResponse) GetIdPVZ() string {
	if x != nil {
		return x.IdPVZ
	}
	return ""
}

type OpenReceptionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IdPVZ         string                 `protobuf:"bytes,2,opt,name=idPVZ,proto3" json:"idPVZ,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OpenReceptionRequest) Reset() {
	*x = OpenReceptionRequest{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OpenReceptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenReceptionRequest) ProtoMessage() {}

func (x *OpenReceptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenReceptionRequest.ProtoReflect.Descriptor instead.
func (*OpenReceptionRequest) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{8}
}

func (x *OpenReceptionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *OpenReceptionRequest) GetIdPVZ() string {
	if x != nil {
		return x.IdPVZ
	}
	return ""
}

type AddItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IdPVZ         string                 `protobuf:"bytes,2,opt,name=idPVZ,proto3" json:"idPVZ,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddItemRequest) Reset() {
	*x = AddItemRequest{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemRequest) ProtoMessage() {}

func (x *AddItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemRequest.ProtoReflect.Descriptor instead.
func (*AddItemRequest) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{9}
}

func (x *AddItemRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AddItemRequest) GetIdPVZ() string {
	if x != nil {
		return x.IdPVZ
	}
	return ""
}

func (x *AddItemRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AddItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IdItem        string                 `protobuf:"bytes,1,opt,name=idItem,proto3" json:"idItem,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddItemResponse) Reset() {
	*x = AddItemResponse{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemResponse) ProtoMessage() {}

func (x *AddItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemResponse.ProtoReflect.Descriptor instead.
func (*AddItemResponse) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{10}
}

func (x *AddItemResponse) GetIdItem() string {
	if x != nil {
		return x.IdItem
	}
	return ""
}

type RemoveItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IdItem        string                 `protobuf:"bytes,2,opt,name=idItem,proto3" json:"idItem,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveItemRequest) Reset() {
	*x = RemoveItemRequest{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveItemRequest) ProtoMessage() {}

func (x *RemoveItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveItemRequest.ProtoReflect.Descriptor instead.
func (*RemoveItemRequest) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{11}
}

func (x *RemoveItemRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RemoveItemRequest) GetIdItem() string {
	if x != nil {
		return x.IdItem
	}
	return ""
}

type CloseReceptionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IdPVZ         string                 `protobuf:"bytes,2,opt,name=idPVZ,proto3" json:"idPVZ,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CloseReceptionRequest) Reset() {
	*x = CloseReceptionRequest{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CloseReceptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseReceptionRequest) ProtoMessage() {}

func (x *CloseReceptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseReceptionRequest.ProtoReflect.Descriptor instead.
func (*CloseReceptionRequest) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{12}
}

func (x *CloseReceptionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CloseReceptionRequest) GetIdPVZ() string {
	if x != nil {
		return x.IdPVZ
	}
	return ""
}

type GetPVZDataRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Start         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	Finish        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=finish,proto3" json:"finish,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPVZDataRequest) Reset() {
	*x = GetPVZDataRequest{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPVZDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPVZDataRequest) ProtoMessage() {}

func (x *GetPVZDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPVZDataRequest.ProtoReflect.Descriptor instead.
func (*GetPVZDataRequest) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{13}
}

func (x *GetPVZDataRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetPVZDataRequest) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *GetPVZDataRequest) GetFinish() *timestamppb.Timestamp {
	if x != nil {
		return x.Finish
	}
	return nil
}

type GetPVZDataResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pvzs          []*PVZ                 `protobuf:"bytes,1,rep,name=pvzs,proto3" json:"pvzs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPVZDataResponse) Reset() {
	*x = GetPVZDataResponse{}
	mi := &file_pvz_service_pvz_service_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPVZDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPVZDataResponse) ProtoMessage() {}

func (x *GetPVZDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pvz_service_pvz_service_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPVZDataResponse.ProtoReflect.Descriptor instead.
func (*GetPVZDataResponse) Descriptor() ([]byte, []int) {
	return file_pvz_service_pvz_service_proto_rawDescGZIP(), []int{14}
}

func (x *GetPVZDataResponse) GetPvzs() []*PVZ {
	if x != nil {
		return x.Pvzs
	}
	return nil
}

var File_pvz_service_pvz_service_proto protoreflect.FileDescriptor

const file_pvz_service_pvz_service_proto_rawDesc = "" +
	"\n" +
	"\x1dpvz_service/pvz_service.proto\x12\vpvz_service\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xba\x01\n" +
	"\x03PVZ\x12\x14\n" +
	"\x05idPVZ\x18\x01 \x01(\tR\x05idPVZ\x12G\n" +
	"\x11registration_date\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x10registrationDate\x12T\n" +
	"\x04city\x18\x03 \x01(\tB@\xfaB=r;R\fМоскваR\x1dСанкт-ПетербургR\fКазаньR\x04city\"\xfa\x02\n" +
	"\tReception\x12 \n" +
	"\vidReception\x18\x01 \x01(\tR\vidReception\x12<\n" +
	"\vopeningTime\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\vopeningTime\x12<\n" +
	"\vclosingTime\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\vclosingTime\x12\x14\n" +
	"\x05idPVZ\x18\x04 \x01(\tR\x05idPVZ\x12'\n" +
	"\x05items\x18\x05 \x03(\v2\x11.pvz_service.ItemR\x05items\x12>\n" +
	"\x06status\x18\x06 \x01(\x0e2&.pvz_service.Reception.ReceptionStatusR\x06status\"P\n" +
	"\x0fReceptionStatus\x12 \n" +
	"\x1cRECEPTION_STATUS_IN_PROGRESS\x10\x00\x12\x1b\n" +
	"\x17RECEPTION_STATUS_CLOSED\x10\x01\"\xad\x01\n" +
	"\x04Item\x12\x16\n" +
	"\x06idItem\x18\x01 \x01(\tR\x06idItem\x12@\n" +
	"\rreceptionTime\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\rreceptionTime\x12K\n" +
	"\x04type\x18\x03 \x01(\tB7\xfaB4r2R\x16электроникаR\fодеждаR\n" +
	"обувьR\x04type\"\x98\x01\n" +
	"\x0fRegisterRequest\x12\x14\n" +
	"\x05login\x18\x01 \x01(\tR\x05login\x12%\n" +
	"\bpassword\x18\x02 \x01(\tB\t\xfaB\x06r\x04\x10\b\x18 R\bpassword\x12H\n" +
	"\x04role\x18\x03 \x01(\tB4\xfaB1r/R\x12модераторR\x19сотрудник ПВЗR\x04role\"K\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05login\x18\x01 \x01(\tR\x05login\x12%\n" +
	"\bpassword\x18\x02 \x01(\tB\t\xfaB\x06r\x04\x10\b\x18 R\bpassword\"%\n" +
	"\rLoginResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"~\n" +
	"\x10CreatePVZRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12T\n" +
	"\x04city\x18\x02 \x01(\tB@\xfaB=r;R\fМоскваR\x1dСанкт-ПетербургR\fКазаньR\x04city\")\n" +
	"\x11CreatePVZResponse\x12\x14\n" +
	"\x05idPVZ\x18\x01 \x01(\tR\x05idPVZ\"B\n" +
	"\x14OpenReceptionRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x14\n" +
	"\x05idPVZ\x18\x02 \x01(\tR\x05idPVZ\"\x89\x01\n" +
	"\x0eAddItemRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x14\n" +
	"\x05idPVZ\x18\x02 \x01(\tR\x05idPVZ\x12K\n" +
	"\x04type\x18\x03 \x01(\tB7\xfaB4r2R\x16электроникаR\fодеждаR\n" +
	"обувьR\x04type\")\n" +
	"\x0fAddItemResponse\x12\x16\n" +
	"\x06idItem\x18\x01 \x01(\tR\x06idItem\"A\n" +
	"\x11RemoveItemRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x16\n" +
	"\x06idItem\x18\x02 \x01(\tR\x06idItem\"C\n" +
	"\x15CloseReceptionRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x14\n" +
	"\x05idPVZ\x18\x02 \x01(\tR\x05idPVZ\"\x8f\x01\n" +
	"\x11GetPVZDataRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x120\n" +
	"\x05start\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x05start\x122\n" +
	"\x06finish\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\x06finish\":\n" +
	"\x12GetPVZDataResponse\x12$\n" +
	"\x04pvzs\x18\x01 \x03(\v2\x10.pvz_service.PVZR\x04pvzs2\xc9\x04\n" +
	"\n" +
	"PVZService\x12@\n" +
	"\bRegister\x12\x1c.pvz_service.RegisterRequest\x1a\x16.google.protobuf.Empty\x12>\n" +
	"\x05Login\x12\x19.pvz_service.LoginRequest\x1a\x1a.pvz_service.LoginResponse\x12J\n" +
	"\tCreatePVZ\x12\x1d.pvz_service.CreatePVZRequest\x1a\x1e.pvz_service.CreatePVZResponse\x12J\n" +
	"\rOpenReception\x12!.pvz_service.OpenReceptionRequest\x1a\x16.google.protobuf.Empty\x12>\n" +
	"\aAddItem\x12\x1b.pvz_service.AddItemRequest\x1a\x16.google.protobuf.Empty\x12D\n" +
	"\n" +
	"RemoveItem\x12\x1e.pvz_service.RemoveItemRequest\x1a\x16.google.protobuf.Empty\x12L\n" +
	"\x0eCloseReception\x12\".pvz_service.CloseReceptionRequest\x1a\x16.google.protobuf.Empty\x12M\n" +
	"\n" +
	"GetPVZData\x12\x1e.pvz_service.GetPVZDataRequest\x1a\x1f.pvz_service.GetPVZDataResponseB>Z<github.com/Turalchik/pvz-service/pkg/pvz_service;pvz_serviceb\x06proto3"

var (
	file_pvz_service_pvz_service_proto_rawDescOnce sync.Once
	file_pvz_service_pvz_service_proto_rawDescData []byte
)

func file_pvz_service_pvz_service_proto_rawDescGZIP() []byte {
	file_pvz_service_pvz_service_proto_rawDescOnce.Do(func() {
		file_pvz_service_pvz_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pvz_service_pvz_service_proto_rawDesc), len(file_pvz_service_pvz_service_proto_rawDesc)))
	})
	return file_pvz_service_pvz_service_proto_rawDescData
}

var file_pvz_service_pvz_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pvz_service_pvz_service_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_pvz_service_pvz_service_proto_goTypes = []any{
	(Reception_ReceptionStatus)(0), // 0: pvz_service.Reception.ReceptionStatus
	(*PVZ)(nil),                    // 1: pvz_service.PVZ
	(*Reception)(nil),              // 2: pvz_service.Reception
	(*Item)(nil),                   // 3: pvz_service.Item
	(*RegisterRequest)(nil),        // 4: pvz_service.RegisterRequest
	(*LoginRequest)(nil),           // 5: pvz_service.LoginRequest
	(*LoginResponse)(nil),          // 6: pvz_service.LoginResponse
	(*CreatePVZRequest)(nil),       // 7: pvz_service.CreatePVZRequest
	(*CreatePVZResponse)(nil),      // 8: pvz_service.CreatePVZResponse
	(*OpenReceptionRequest)(nil),   // 9: pvz_service.OpenReceptionRequest
	(*AddItemRequest)(nil),         // 10: pvz_service.AddItemRequest
	(*AddItemResponse)(nil),        // 11: pvz_service.AddItemResponse
	(*RemoveItemRequest)(nil),      // 12: pvz_service.RemoveItemRequest
	(*CloseReceptionRequest)(nil),  // 13: pvz_service.CloseReceptionRequest
	(*GetPVZDataRequest)(nil),      // 14: pvz_service.GetPVZDataRequest
	(*GetPVZDataResponse)(nil),     // 15: pvz_service.GetPVZDataResponse
	(*timestamppb.Timestamp)(nil),  // 16: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 17: google.protobuf.Empty
}
var file_pvz_service_pvz_service_proto_depIdxs = []int32{
	16, // 0: pvz_service.PVZ.registration_date:type_name -> google.protobuf.Timestamp
	16, // 1: pvz_service.Reception.openingTime:type_name -> google.protobuf.Timestamp
	16, // 2: pvz_service.Reception.closingTime:type_name -> google.protobuf.Timestamp
	3,  // 3: pvz_service.Reception.items:type_name -> pvz_service.Item
	0,  // 4: pvz_service.Reception.status:type_name -> pvz_service.Reception.ReceptionStatus
	16, // 5: pvz_service.Item.receptionTime:type_name -> google.protobuf.Timestamp
	16, // 6: pvz_service.GetPVZDataRequest.start:type_name -> google.protobuf.Timestamp
	16, // 7: pvz_service.GetPVZDataRequest.finish:type_name -> google.protobuf.Timestamp
	1,  // 8: pvz_service.GetPVZDataResponse.pvzs:type_name -> pvz_service.PVZ
	4,  // 9: pvz_service.PVZService.Register:input_type -> pvz_service.RegisterRequest
	5,  // 10: pvz_service.PVZService.Login:input_type -> pvz_service.LoginRequest
	7,  // 11: pvz_service.PVZService.CreatePVZ:input_type -> pvz_service.CreatePVZRequest
	9,  // 12: pvz_service.PVZService.OpenReception:input_type -> pvz_service.OpenReceptionRequest
	10, // 13: pvz_service.PVZService.AddItem:input_type -> pvz_service.AddItemRequest
	12, // 14: pvz_service.PVZService.RemoveItem:input_type -> pvz_service.RemoveItemRequest
	13, // 15: pvz_service.PVZService.CloseReception:input_type -> pvz_service.CloseReceptionRequest
	14, // 16: pvz_service.PVZService.GetPVZData:input_type -> pvz_service.GetPVZDataRequest
	17, // 17: pvz_service.PVZService.Register:output_type -> google.protobuf.Empty
	6,  // 18: pvz_service.PVZService.Login:output_type -> pvz_service.LoginResponse
	8,  // 19: pvz_service.PVZService.CreatePVZ:output_type -> pvz_service.CreatePVZResponse
	17, // 20: pvz_service.PVZService.OpenReception:output_type -> google.protobuf.Empty
	17, // 21: pvz_service.PVZService.AddItem:output_type -> google.protobuf.Empty
	17, // 22: pvz_service.PVZService.RemoveItem:output_type -> google.protobuf.Empty
	17, // 23: pvz_service.PVZService.CloseReception:output_type -> google.protobuf.Empty
	15, // 24: pvz_service.PVZService.GetPVZData:output_type -> pvz_service.GetPVZDataResponse
	17, // [17:25] is the sub-list for method output_type
	9,  // [9:17] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pvz_service_pvz_service_proto_init() }
func file_pvz_service_pvz_service_proto_init() {
	if File_pvz_service_pvz_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pvz_service_pvz_service_proto_rawDesc), len(file_pvz_service_pvz_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pvz_service_pvz_service_proto_goTypes,
		DependencyIndexes: file_pvz_service_pvz_service_proto_depIdxs,
		EnumInfos:         file_pvz_service_pvz_service_proto_enumTypes,
		MessageInfos:      file_pvz_service_pvz_service_proto_msgTypes,
	}.Build()
	File_pvz_service_pvz_service_proto = out.File
	file_pvz_service_pvz_service_proto_goTypes = nil
	file_pvz_service_pvz_service_proto_depIdxs = nil
}
