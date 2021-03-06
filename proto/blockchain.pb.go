// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/blockchain.proto

package in_memory_blockchain

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

type Balance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance float64 `protobuf:"fixed64,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Balance) Reset() {
	*x = Balance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Balance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Balance) ProtoMessage() {}

func (x *Balance) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Balance.ProtoReflect.Descriptor instead.
func (*Balance) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *Balance) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Balance) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type QueryParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlkHash string `protobuf:"bytes,1,opt,name=blk_hash,json=blkHash,proto3" json:"blk_hash,omitempty"`
	TxHash  string `protobuf:"bytes,2,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Range   int32  `protobuf:"varint,4,opt,name=range,proto3" json:"range,omitempty"`
}

func (x *QueryParams) Reset() {
	*x = QueryParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParams) ProtoMessage() {}

func (x *QueryParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParams.ProtoReflect.Descriptor instead.
func (*QueryParams) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{1}
}

func (x *QueryParams) GetBlkHash() string {
	if x != nil {
		return x.BlkHash
	}
	return ""
}

func (x *QueryParams) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *QueryParams) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *QueryParams) GetRange() int32 {
	if x != nil {
		return x.Range
	}
	return 0
}

type TxPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Value float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TxPayload) Reset() {
	*x = TxPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxPayload) ProtoMessage() {}

func (x *TxPayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxPayload.ProtoReflect.Descriptor instead.
func (*TxPayload) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{2}
}

func (x *TxPayload) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TxPayload) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *TxPayload) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash  string  `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	From  string  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To    string  `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value float64 `protobuf:"fixed64,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Tx) Reset() {
	*x = Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tx) ProtoMessage() {}

func (x *Tx) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tx.ProtoReflect.Descriptor instead.
func (*Tx) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{3}
}

func (x *Tx) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Tx) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Tx) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Tx) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TxList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Tx `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *TxList) Reset() {
	*x = TxList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxList) ProtoMessage() {}

func (x *TxList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxList.ProtoReflect.Descriptor instead.
func (*TxList) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{4}
}

func (x *TxList) GetTransactions() []*Tx {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash       string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash string   `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Height     int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	TxHashes   []string `protobuf:"bytes,4,rep,name=tx_hashes,json=txHashes,proto3" json:"tx_hashes,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{5}
}

func (x *Header) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Header) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *Header) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Header) GetTxHashes() []string {
	if x != nil {
		return x.TxHashes
	}
	return nil
}

type Blk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	TxList *TxList `protobuf:"bytes,2,opt,name=tx_list,json=txList,proto3" json:"tx_list,omitempty"`
}

func (x *Blk) Reset() {
	*x = Blk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blk) ProtoMessage() {}

func (x *Blk) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blk.ProtoReflect.Descriptor instead.
func (*Blk) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{6}
}

func (x *Blk) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Blk) GetTxList() *TxList {
	if x != nil {
		return x.TxList
	}
	return nil
}

type BlkList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Blk `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *BlkList) Reset() {
	*x = BlkList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_blockchain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlkList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlkList) ProtoMessage() {}

func (x *BlkList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_blockchain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlkList.ProtoReflect.Descriptor instead.
func (*BlkList) Descriptor() ([]byte, []int) {
	return file_proto_blockchain_proto_rawDescGZIP(), []int{7}
}

func (x *BlkList) GetBlocks() []*Blk {
	if x != nil {
		return x.Blocks
	}
	return nil
}

var File_proto_blockchain_proto protoreflect.FileDescriptor

var file_proto_blockchain_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x3d, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x71,
	0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x6c, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x6c, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67,
	0x65, 0x22, 0x45, 0x0a, 0x09, 0x54, 0x78, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x52, 0x0a, 0x02, 0x54, 0x78, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x37, 0x0a, 0x06,
	0x54, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x78, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x72, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x22, 0x54, 0x0a, 0x03, 0x42, 0x6c, 0x6b,
	0x12, 0x25, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x74, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0x2d, 0x0a, 0x07, 0x42, 0x6c, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x32, 0xde,
	0x01, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x32, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x78, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x78, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x42, 0x6c, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x42,
	0x17, 0x5a, 0x15, 0x2f, 0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_blockchain_proto_rawDescOnce sync.Once
	file_proto_blockchain_proto_rawDescData = file_proto_blockchain_proto_rawDesc
)

func file_proto_blockchain_proto_rawDescGZIP() []byte {
	file_proto_blockchain_proto_rawDescOnce.Do(func() {
		file_proto_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_blockchain_proto_rawDescData)
	})
	return file_proto_blockchain_proto_rawDescData
}

var file_proto_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_blockchain_proto_goTypes = []interface{}{
	(*Balance)(nil),     // 0: proto.Balance
	(*QueryParams)(nil), // 1: proto.QueryParams
	(*TxPayload)(nil),   // 2: proto.TxPayload
	(*Tx)(nil),          // 3: proto.Tx
	(*TxList)(nil),      // 4: proto.TxList
	(*Header)(nil),      // 5: proto.Header
	(*Blk)(nil),         // 6: proto.Blk
	(*BlkList)(nil),     // 7: proto.BlkList
}
var file_proto_blockchain_proto_depIdxs = []int32{
	3, // 0: proto.TxList.transactions:type_name -> proto.Tx
	5, // 1: proto.Blk.header:type_name -> proto.Header
	4, // 2: proto.Blk.tx_list:type_name -> proto.TxList
	6, // 3: proto.BlkList.blocks:type_name -> proto.Blk
	2, // 4: proto.Blockchain.CreateTransaction:input_type -> proto.TxPayload
	1, // 5: proto.Blockchain.GetTransactions:input_type -> proto.QueryParams
	1, // 6: proto.Blockchain.GetBlock:input_type -> proto.QueryParams
	1, // 7: proto.Blockchain.GetBalance:input_type -> proto.QueryParams
	3, // 8: proto.Blockchain.CreateTransaction:output_type -> proto.Tx
	4, // 9: proto.Blockchain.GetTransactions:output_type -> proto.TxList
	7, // 10: proto.Blockchain.GetBlock:output_type -> proto.BlkList
	0, // 11: proto.Blockchain.GetBalance:output_type -> proto.Balance
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_blockchain_proto_init() }
func file_proto_blockchain_proto_init() {
	if File_proto_blockchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_blockchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Balance); i {
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
		file_proto_blockchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParams); i {
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
		file_proto_blockchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxPayload); i {
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
		file_proto_blockchain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tx); i {
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
		file_proto_blockchain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxList); i {
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
		file_proto_blockchain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_proto_blockchain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blk); i {
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
		file_proto_blockchain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlkList); i {
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
			RawDescriptor: file_proto_blockchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_blockchain_proto_goTypes,
		DependencyIndexes: file_proto_blockchain_proto_depIdxs,
		MessageInfos:      file_proto_blockchain_proto_msgTypes,
	}.Build()
	File_proto_blockchain_proto = out.File
	file_proto_blockchain_proto_rawDesc = nil
	file_proto_blockchain_proto_goTypes = nil
	file_proto_blockchain_proto_depIdxs = nil
}
