// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/message.proto

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

type MessageTypes int32

const (
	MessageTypes_DEFAULT    MessageTypes = 0
	MessageTypes_ENCRYPTED  MessageTypes = 1
	MessageTypes_PLAIN_TEXT MessageTypes = 2
	MessageTypes_GOSSIP     MessageTypes = 3
	MessageTypes_HANDSHAKE  MessageTypes = 4
	MessageTypes_SIGNATURE  MessageTypes = 5
	MessageTypes_REQUEST    MessageTypes = 6
	MessageTypes_AUTH       MessageTypes = 7
	MessageTypes_ERROR      MessageTypes = 8
	MessageTypes_TIME       MessageTypes = 9
	MessageTypes_TX         MessageTypes = 10
	MessageTypes_BATCH_SIG  MessageTypes = 11
	MessageTypes_BATCH_STD  MessageTypes = 12
)

// Enum value maps for MessageTypes.
var (
	MessageTypes_name = map[int32]string{
		0:  "DEFAULT",
		1:  "ENCRYPTED",
		2:  "PLAIN_TEXT",
		3:  "GOSSIP",
		4:  "HANDSHAKE",
		5:  "SIGNATURE",
		6:  "REQUEST",
		7:  "AUTH",
		8:  "ERROR",
		9:  "TIME",
		10: "TX",
		11: "BATCH_SIG",
		12: "BATCH_STD",
	}
	MessageTypes_value = map[string]int32{
		"DEFAULT":    0,
		"ENCRYPTED":  1,
		"PLAIN_TEXT": 2,
		"GOSSIP":     3,
		"HANDSHAKE":  4,
		"SIGNATURE":  5,
		"REQUEST":    6,
		"AUTH":       7,
		"ERROR":      8,
		"TIME":       9,
		"TX":         10,
		"BATCH_SIG":  11,
		"BATCH_STD":  12,
	}
)

func (x MessageTypes) Enum() *MessageTypes {
	p := new(MessageTypes)
	*p = x
	return p
}

func (x MessageTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_message_proto_enumTypes[0].Descriptor()
}

func (MessageTypes) Type() protoreflect.EnumType {
	return &file_proto_message_proto_enumTypes[0]
}

func (x MessageTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageTypes.Descriptor instead.
func (MessageTypes) EnumDescriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{0}
}

type SIGNATURE_ALGS int32

const (
	SIGNATURE_ALGS_SIG_DEFAULT     SIGNATURE_ALGS = 0
	SIGNATURE_ALGS_SIG_EDWARD25519 SIGNATURE_ALGS = 1
	SIGNATURE_ALGS_SIG_FALCON      SIGNATURE_ALGS = 2
	SIGNATURE_ALGS_SIG_HPKE        SIGNATURE_ALGS = 3
)

// Enum value maps for SIGNATURE_ALGS.
var (
	SIGNATURE_ALGS_name = map[int32]string{
		0: "SIG_DEFAULT",
		1: "SIG_EDWARD25519",
		2: "SIG_FALCON",
		3: "SIG_HPKE",
	}
	SIGNATURE_ALGS_value = map[string]int32{
		"SIG_DEFAULT":     0,
		"SIG_EDWARD25519": 1,
		"SIG_FALCON":      2,
		"SIG_HPKE":        3,
	}
)

func (x SIGNATURE_ALGS) Enum() *SIGNATURE_ALGS {
	p := new(SIGNATURE_ALGS)
	*p = x
	return p
}

func (x SIGNATURE_ALGS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SIGNATURE_ALGS) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_message_proto_enumTypes[1].Descriptor()
}

func (SIGNATURE_ALGS) Type() protoreflect.EnumType {
	return &file_proto_message_proto_enumTypes[1]
}

func (x SIGNATURE_ALGS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SIGNATURE_ALGS.Descriptor instead.
func (SIGNATURE_ALGS) EnumDescriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{1}
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Signer    string         `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
	PubKey    string         `protobuf:"bytes,3,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Signature string         `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Alg       SIGNATURE_ALGS `protobuf:"varint,5,opt,name=alg,proto3,enum=SIGNATURE_ALGS" json:"alg,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{0}
}

func (x *Signature) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Signature) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

func (x *Signature) GetPubKey() string {
	if x != nil {
		return x.PubKey
	}
	return ""
}

func (x *Signature) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Signature) GetAlg() SIGNATURE_ALGS {
	if x != nil {
		return x.Alg
	}
	return SIGNATURE_ALGS_SIG_DEFAULT
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp int64        `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Type      MessageTypes `protobuf:"varint,3,opt,name=type,proto3,enum=MessageTypes" json:"type,omitempty"`
	Signature *Signature   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetType() MessageTypes {
	if x != nil {
		return x.Type
	}
	return MessageTypes_DEFAULT
}

func (x *Message) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

type EncryptedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgData *Message `protobuf:"bytes,1,opt,name=msgData,proto3" json:"msgData,omitempty"`
	PeerID  string   `protobuf:"bytes,2,opt,name=peerID,proto3" json:"peerID,omitempty"`
	Aad     string   `protobuf:"bytes,3,opt,name=aad,proto3" json:"aad,omitempty"`
	Data    string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EncryptedMessage) Reset() {
	*x = EncryptedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedMessage) ProtoMessage() {}

func (x *EncryptedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedMessage.ProtoReflect.Descriptor instead.
func (*EncryptedMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{2}
}

func (x *EncryptedMessage) GetMsgData() *Message {
	if x != nil {
		return x.MsgData
	}
	return nil
}

func (x *EncryptedMessage) GetPeerID() string {
	if x != nil {
		return x.PeerID
	}
	return ""
}

func (x *EncryptedMessage) GetAad() string {
	if x != nil {
		return x.Aad
	}
	return ""
}

func (x *EncryptedMessage) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgData  *Message `protobuf:"bytes,1,opt,name=msgData,proto3" json:"msgData,omitempty"`
	Request  []byte   `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	FromPeer string   `protobuf:"bytes,3,opt,name=fromPeer,proto3" json:"fromPeer,omitempty"`
	Referrer string   `protobuf:"bytes,4,opt,name=referrer,proto3" json:"referrer,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{3}
}

func (x *Request) GetMsgData() *Message {
	if x != nil {
		return x.MsgData
	}
	return nil
}

func (x *Request) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Request) GetFromPeer() string {
	if x != nil {
		return x.FromPeer
	}
	return ""
}

func (x *Request) GetReferrer() string {
	if x != nil {
		return x.Referrer
	}
	return ""
}

type Errors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    int32  `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Timestamp    int64  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Errors) Reset() {
	*x = Errors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Errors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Errors) ProtoMessage() {}

func (x *Errors) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Errors.ProtoReflect.Descriptor instead.
func (*Errors) Descriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{4}
}

func (x *Errors) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *Errors) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *Errors) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgData  *Message `protobuf:"bytes,1,opt,name=msgData,proto3" json:"msgData,omitempty"`
	Response []byte   `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	Errors   *Errors  `protobuf:"bytes,3,opt,name=errors,proto3" json:"errors,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_message_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetMsgData() *Message {
	if x != nil {
		return x.MsgData
	}
	return nil
}

func (x *Response) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *Response) GetErrors() *Errors {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_proto_message_proto protoreflect.FileDescriptor

var file_proto_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x21, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x41, 0x4c, 0x47, 0x53, 0x52,
	0x03, 0x61, 0x6c, 0x67, 0x22, 0x84, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x21,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x74, 0x0a, 0x10, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x22, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x61, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x7f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07,
	0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72,
	0x6f, 0x6d, 0x50, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72,
	0x6f, 0x6d, 0x50, 0x65, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x22, 0x68, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x6b, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2a, 0xb6, 0x01, 0x0a, 0x0c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x4e, 0x43, 0x52, 0x59,
	0x50, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x5f,
	0x54, 0x45, 0x58, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4f, 0x53, 0x53, 0x49, 0x50,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x41, 0x4e, 0x44, 0x53, 0x48, 0x41, 0x4b, 0x45, 0x10,
	0x04, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x05,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x06, 0x12, 0x08, 0x0a,
	0x04, 0x41, 0x55, 0x54, 0x48, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x09, 0x12, 0x06, 0x0a, 0x02,
	0x54, 0x58, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x49,
	0x47, 0x10, 0x0b, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x53, 0x54, 0x44,
	0x10, 0x0c, 0x2a, 0x54, 0x0a, 0x0e, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f,
	0x41, 0x4c, 0x47, 0x53, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x49, 0x47, 0x5f, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x49, 0x47, 0x5f, 0x45, 0x44, 0x57,
	0x41, 0x52, 0x44, 0x32, 0x35, 0x35, 0x31, 0x39, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x49,
	0x47, 0x5f, 0x46, 0x41, 0x4c, 0x43, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x49,
	0x47, 0x5f, 0x48, 0x50, 0x4b, 0x45, 0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_proto_rawDescOnce sync.Once
	file_proto_message_proto_rawDescData = file_proto_message_proto_rawDesc
)

func file_proto_message_proto_rawDescGZIP() []byte {
	file_proto_message_proto_rawDescOnce.Do(func() {
		file_proto_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_proto_rawDescData)
	})
	return file_proto_message_proto_rawDescData
}

var file_proto_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_message_proto_goTypes = []interface{}{
	(MessageTypes)(0),        // 0: MessageTypes
	(SIGNATURE_ALGS)(0),      // 1: SIGNATURE_ALGS
	(*Signature)(nil),        // 2: Signature
	(*Message)(nil),          // 3: Message
	(*EncryptedMessage)(nil), // 4: EncryptedMessage
	(*Request)(nil),          // 5: Request
	(*Errors)(nil),           // 6: Errors
	(*Response)(nil),         // 7: Response
}
var file_proto_message_proto_depIdxs = []int32{
	1, // 0: Signature.alg:type_name -> SIGNATURE_ALGS
	0, // 1: Message.type:type_name -> MessageTypes
	2, // 2: Message.signature:type_name -> Signature
	3, // 3: EncryptedMessage.msgData:type_name -> Message
	3, // 4: Request.msgData:type_name -> Message
	3, // 5: Response.msgData:type_name -> Message
	6, // 6: Response.errors:type_name -> Errors
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_proto_message_proto_init() }
func file_proto_message_proto_init() {
	if File_proto_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_proto_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedMessage); i {
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
		file_proto_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Errors); i {
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
		file_proto_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_message_proto_goTypes,
		DependencyIndexes: file_proto_message_proto_depIdxs,
		EnumInfos:         file_proto_message_proto_enumTypes,
		MessageInfos:      file_proto_message_proto_msgTypes,
	}.Build()
	File_proto_message_proto = out.File
	file_proto_message_proto_rawDesc = nil
	file_proto_message_proto_goTypes = nil
	file_proto_message_proto_depIdxs = nil
}
