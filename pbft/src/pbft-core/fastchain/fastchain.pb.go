// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fastchain.proto

package fastchain

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Inner                *Request_Inner `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
	Dig                  []byte         `protobuf:"bytes,2,opt,name=dig,proto3" json:"dig,omitempty"`
	Outer                []byte         `protobuf:"bytes,3,opt,name=outer,proto3" json:"outer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_05184b12c21fa7e0, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetInner() *Request_Inner {
	if m != nil {
		return m.Inner
	}
	return nil
}

func (m *Request) GetDig() []byte {
	if m != nil {
		return m.Dig
	}
	return nil
}

func (m *Request) GetOuter() []byte {
	if m != nil {
		return m.Outer
	}
	return nil
}

type Request_Inner struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Seq                  int32      `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	View                 int32      `protobuf:"varint,3,opt,name=view,proto3" json:"view,omitempty"`
	Type                 int32      `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Block                *PbftBlock `protobuf:"bytes,5,opt,name=block,proto3" json:"block,omitempty"`
	Timestamp            int64      `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Request_Inner) Reset()         { *m = Request_Inner{} }
func (m *Request_Inner) String() string { return proto.CompactTextString(m) }
func (*Request_Inner) ProtoMessage()    {}
func (*Request_Inner) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_05184b12c21fa7e0, []int{0, 0}
}
func (m *Request_Inner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_Inner.Unmarshal(m, b)
}
func (m *Request_Inner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_Inner.Marshal(b, m, deterministic)
}
func (dst *Request_Inner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_Inner.Merge(dst, src)
}
func (m *Request_Inner) XXX_Size() int {
	return xxx_messageInfo_Request_Inner.Size(m)
}
func (m *Request_Inner) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_Inner.DiscardUnknown(m)
}

var xxx_messageInfo_Request_Inner proto.InternalMessageInfo

func (m *Request_Inner) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Request_Inner) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Request_Inner) GetView() int32 {
	if m != nil {
		return m.View
	}
	return 0
}

func (m *Request_Inner) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Request_Inner) GetBlock() *PbftBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *Request_Inner) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PbftNode struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Pubkey               string   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Privkey              string   `protobuf:"bytes,3,opt,name=privkey,proto3" json:"privkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbftNode) Reset()         { *m = PbftNode{} }
func (m *PbftNode) String() string { return proto.CompactTextString(m) }
func (*PbftNode) ProtoMessage()    {}
func (*PbftNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_05184b12c21fa7e0, []int{1}
}
func (m *PbftNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftNode.Unmarshal(m, b)
}
func (m *PbftNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftNode.Marshal(b, m, deterministic)
}
func (dst *PbftNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftNode.Merge(dst, src)
}
func (m *PbftNode) XXX_Size() int {
	return xxx_messageInfo_PbftNode.Size(m)
}
func (m *PbftNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftNode.DiscardUnknown(m)
}

var xxx_messageInfo_PbftNode proto.InternalMessageInfo

func (m *PbftNode) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *PbftNode) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *PbftNode) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

type Nodes struct {
	Nodes                []*PbftNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Nodes) Reset()         { *m = Nodes{} }
func (m *Nodes) String() string { return proto.CompactTextString(m) }
func (*Nodes) ProtoMessage()    {}
func (*Nodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_05184b12c21fa7e0, []int{2}
}
func (m *Nodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nodes.Unmarshal(m, b)
}
func (m *Nodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nodes.Marshal(b, m, deterministic)
}
func (dst *Nodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nodes.Merge(dst, src)
}
func (m *Nodes) XXX_Size() int {
	return xxx_messageInfo_Nodes.Size(m)
}
func (m *Nodes) XXX_DiscardUnknown() {
	xxx_messageInfo_Nodes.DiscardUnknown(m)
}

var xxx_messageInfo_Nodes proto.InternalMessageInfo

func (m *Nodes) GetNodes() []*PbftNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type TxnData struct {
	AccountNonce         uint64   `protobuf:"varint,1,opt,name=AccountNonce,json=accountNonce,proto3" json:"AccountNonce,omitempty"`
	Price                int64    `protobuf:"varint,2,opt,name=Price,json=price,proto3" json:"Price,omitempty"`
	GasLimit             int64    `protobuf:"varint,3,opt,name=GasLimit,json=gasLimit,proto3" json:"GasLimit,omitempty"`
	Recipient            []byte   `protobuf:"bytes,4,opt,name=Recipient,json=recipient,proto3" json:"Recipient,omitempty"`
	Amount               int64    `protobuf:"varint,5,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
	Payload              []byte   `protobuf:"bytes,6,opt,name=Payload,json=payload,proto3" json:"Payload,omitempty"`
	Signature            []byte   `protobuf:"bytes,7,opt,name=Signature,json=signature,proto3" json:"Signature,omitempty"`
	Hash                 []byte   `protobuf:"bytes,8,opt,name=Hash,json=hash,proto3" json:"Hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxnData) Reset()         { *m = TxnData{} }
func (m *TxnData) String() string { return proto.CompactTextString(m) }
func (*TxnData) ProtoMessage()    {}
func (*TxnData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_05184b12c21fa7e0, []int{3}
}
func (m *TxnData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxnData.Unmarshal(m, b)
}
func (m *TxnData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxnData.Marshal(b, m, deterministic)
}
func (dst *TxnData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxnData.Merge(dst, src)
}
func (m *TxnData) XXX_Size() int {
	return xxx_messageInfo_TxnData.Size(m)
}
func (m *TxnData) XXX_DiscardUnknown() {
	xxx_messageInfo_TxnData.DiscardUnknown(m)
}

var xxx_messageInfo_TxnData proto.InternalMessageInfo

func (m *TxnData) GetAccountNonce() uint64 {
	if m != nil {
		return m.AccountNonce
	}
	return 0
}

func (m *TxnData) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TxnData) GetGasLimit() int64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *TxnData) GetRecipient() []byte {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *TxnData) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TxnData) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TxnData) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *TxnData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Transaction struct {
	Data                 *TxnData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_05184b12c21fa7e0, []int{4}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetData() *TxnData {
	if m != nil {
		return m.Data
	}
	return nil
}

type PbftBlockHeader struct {
	Number               int64    `protobuf:"varint,1,opt,name=Number,json=number,proto3" json:"Number,omitempty"`
	GasLimit             int64    `protobuf:"varint,2,opt,name=GasLimit,json=gasLimit,proto3" json:"GasLimit,omitempty"`
	GasUsed              int64    `protobuf:"varint,3,opt,name=GasUsed,json=gasUsed,proto3" json:"GasUsed,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=Timestamp,json=timestamp,proto3" json:"Timestamp,omitempty"`
	ParentHash           []byte   `protobuf:"bytes,5,opt,name=ParentHash,json=parentHash,proto3" json:"ParentHash,omitempty"`
	TxnsHash             []byte   `protobuf:"bytes,6,opt,name=TxnsHash,json=txnsHash,proto3" json:"TxnsHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbftBlockHeader) Reset()         { *m = PbftBlockHeader{} }
func (m *PbftBlockHeader) String() string { return proto.CompactTextString(m) }
func (*PbftBlockHeader) ProtoMessage()    {}
func (*PbftBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_05184b12c21fa7e0, []int{5}
}
func (m *PbftBlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftBlockHeader.Unmarshal(m, b)
}
func (m *PbftBlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftBlockHeader.Marshal(b, m, deterministic)
}
func (dst *PbftBlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftBlockHeader.Merge(dst, src)
}
func (m *PbftBlockHeader) XXX_Size() int {
	return xxx_messageInfo_PbftBlockHeader.Size(m)
}
func (m *PbftBlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftBlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_PbftBlockHeader proto.InternalMessageInfo

func (m *PbftBlockHeader) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *PbftBlockHeader) GetGasLimit() int64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *PbftBlockHeader) GetGasUsed() int64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *PbftBlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PbftBlockHeader) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *PbftBlockHeader) GetTxnsHash() []byte {
	if m != nil {
		return m.TxnsHash
	}
	return nil
}

type PbftBlock struct {
	Header               *PbftBlockHeader `protobuf:"bytes,1,opt,name=Header,json=header,proto3" json:"Header,omitempty"`
	Txns                 []*Transaction   `protobuf:"bytes,2,rep,name=Txns,json=txns,proto3" json:"Txns,omitempty"`
	Signs                []string         `protobuf:"bytes,3,rep,name=Signs,json=signs,proto3" json:"Signs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PbftBlock) Reset()         { *m = PbftBlock{} }
func (m *PbftBlock) String() string { return proto.CompactTextString(m) }
func (*PbftBlock) ProtoMessage()    {}
func (*PbftBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_05184b12c21fa7e0, []int{6}
}
func (m *PbftBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbftBlock.Unmarshal(m, b)
}
func (m *PbftBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbftBlock.Marshal(b, m, deterministic)
}
func (dst *PbftBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbftBlock.Merge(dst, src)
}
func (m *PbftBlock) XXX_Size() int {
	return xxx_messageInfo_PbftBlock.Size(m)
}
func (m *PbftBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_PbftBlock.DiscardUnknown(m)
}

var xxx_messageInfo_PbftBlock proto.InternalMessageInfo

func (m *PbftBlock) GetHeader() *PbftBlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PbftBlock) GetTxns() []*Transaction {
	if m != nil {
		return m.Txns
	}
	return nil
}

func (m *PbftBlock) GetSigns() []string {
	if m != nil {
		return m.Signs
	}
	return nil
}

type TrueChain struct {
	Blocks               []*PbftBlock     `protobuf:"bytes,1,rep,name=Blocks,json=blocks,proto3" json:"Blocks,omitempty"`
	LastBlockHeader      *PbftBlockHeader `protobuf:"bytes,2,opt,name=LastBlockHeader,json=lastBlockHeader,proto3" json:"LastBlockHeader,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TrueChain) Reset()         { *m = TrueChain{} }
func (m *TrueChain) String() string { return proto.CompactTextString(m) }
func (*TrueChain) ProtoMessage()    {}
func (*TrueChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_05184b12c21fa7e0, []int{7}
}
func (m *TrueChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrueChain.Unmarshal(m, b)
}
func (m *TrueChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrueChain.Marshal(b, m, deterministic)
}
func (dst *TrueChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrueChain.Merge(dst, src)
}
func (m *TrueChain) XXX_Size() int {
	return xxx_messageInfo_TrueChain.Size(m)
}
func (m *TrueChain) XXX_DiscardUnknown() {
	xxx_messageInfo_TrueChain.DiscardUnknown(m)
}

var xxx_messageInfo_TrueChain proto.InternalMessageInfo

func (m *TrueChain) GetBlocks() []*PbftBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *TrueChain) GetLastBlockHeader() *PbftBlockHeader {
	if m != nil {
		return m.LastBlockHeader
	}
	return nil
}

type GenericResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=Msg,json=msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResp) Reset()         { *m = GenericResp{} }
func (m *GenericResp) String() string { return proto.CompactTextString(m) }
func (*GenericResp) ProtoMessage()    {}
func (*GenericResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_05184b12c21fa7e0, []int{8}
}
func (m *GenericResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResp.Unmarshal(m, b)
}
func (m *GenericResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResp.Marshal(b, m, deterministic)
}
func (dst *GenericResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResp.Merge(dst, src)
}
func (m *GenericResp) XXX_Size() int {
	return xxx_messageInfo_GenericResp.Size(m)
}
func (m *GenericResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResp.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResp proto.InternalMessageInfo

func (m *GenericResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "fastchain.Request")
	proto.RegisterType((*Request_Inner)(nil), "fastchain.Request.Inner")
	proto.RegisterType((*PbftNode)(nil), "fastchain.PbftNode")
	proto.RegisterType((*Nodes)(nil), "fastchain.Nodes")
	proto.RegisterType((*TxnData)(nil), "fastchain.TxnData")
	proto.RegisterType((*Transaction)(nil), "fastchain.Transaction")
	proto.RegisterType((*PbftBlockHeader)(nil), "fastchain.PbftBlockHeader")
	proto.RegisterType((*PbftBlock)(nil), "fastchain.PbftBlock")
	proto.RegisterType((*TrueChain)(nil), "fastchain.TrueChain")
	proto.RegisterType((*GenericResp)(nil), "fastchain.GenericResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FastChainClient is the client API for FastChain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FastChainClient interface {
	// Send new transaction to presumed leader node
	NewTxnRequest(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*GenericResp, error)
}

type fastChainClient struct {
	cc *grpc.ClientConn
}

func NewFastChainClient(cc *grpc.ClientConn) FastChainClient {
	return &fastChainClient{cc}
}

func (c *fastChainClient) NewTxnRequest(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*GenericResp, error) {
	out := new(GenericResp)
	err := c.cc.Invoke(ctx, "/fastchain.FastChain/NewTxnRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FastChainServer is the server API for FastChain service.
type FastChainServer interface {
	// Send new transaction to presumed leader node
	NewTxnRequest(context.Context, *Transaction) (*GenericResp, error)
}

func RegisterFastChainServer(s *grpc.Server, srv FastChainServer) {
	s.RegisterService(&_FastChain_serviceDesc, srv)
}

func _FastChain_NewTxnRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FastChainServer).NewTxnRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fastchain.FastChain/NewTxnRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FastChainServer).NewTxnRequest(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _FastChain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fastchain.FastChain",
	HandlerType: (*FastChainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewTxnRequest",
			Handler:    _FastChain_NewTxnRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fastchain.proto",
}

func init() { proto.RegisterFile("fastchain.proto", fileDescriptor_fastchain_05184b12c21fa7e0) }

var fileDescriptor_fastchain_05184b12c21fa7e0 = []byte{
	// 664 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x5b, 0x6e, 0x13, 0x3d,
	0x14, 0xfe, 0x27, 0x73, 0x49, 0xe6, 0x24, 0x3f, 0x41, 0xa6, 0xaa, 0x46, 0x11, 0x82, 0x68, 0x1e,
	0x50, 0xa8, 0x50, 0x1e, 0x82, 0x58, 0x40, 0xa1, 0xa2, 0x45, 0x2a, 0x51, 0x64, 0xc2, 0x02, 0x9c,
	0x19, 0x37, 0xb1, 0xda, 0x78, 0xa6, 0x63, 0xa7, 0x97, 0x27, 0xc4, 0x3a, 0x58, 0x09, 0xcb, 0x61,
	0x25, 0xa0, 0x73, 0x3c, 0x49, 0xa7, 0x55, 0xe1, 0xed, 0x5c, 0x3f, 0x7f, 0xe7, 0xf3, 0xb1, 0xa1,
	0x7f, 0x26, 0x8c, 0xcd, 0x56, 0x42, 0xe9, 0x71, 0x59, 0x15, 0xb6, 0x60, 0xf1, 0x2e, 0x90, 0xfe,
	0xf6, 0xa0, 0xcd, 0xe5, 0xe5, 0x46, 0x1a, 0xcb, 0xc6, 0x10, 0x2a, 0xad, 0x65, 0x95, 0x78, 0x43,
	0x6f, 0xd4, 0x9d, 0x24, 0xe3, 0xbb, 0xbe, 0xba, 0x64, 0xfc, 0x09, 0xf3, 0xdc, 0x95, 0xb1, 0xa7,
	0xe0, 0xe7, 0x6a, 0x99, 0xb4, 0x86, 0xde, 0xa8, 0xc7, 0xd1, 0x64, 0x7b, 0x10, 0x16, 0x1b, 0x2b,
	0xab, 0xc4, 0xa7, 0x98, 0x73, 0x06, 0x3f, 0x3c, 0x08, 0xa9, 0x91, 0x3d, 0x81, 0x96, 0xca, 0x09,
	0x3e, 0xe4, 0x2d, 0x95, 0x23, 0x82, 0x91, 0x97, 0x84, 0x10, 0x72, 0x34, 0x19, 0x83, 0xe0, 0x4a,
	0xc9, 0x6b, 0x02, 0x08, 0x39, 0xd9, 0x18, 0xb3, 0xb7, 0xa5, 0x4c, 0x02, 0x17, 0x43, 0x9b, 0x1d,
	0x40, 0xb8, 0xb8, 0x28, 0xb2, 0xf3, 0x24, 0x24, 0xae, 0x7b, 0x0d, 0xae, 0xb3, 0xc5, 0x99, 0x7d,
	0x8f, 0x39, 0xee, 0x4a, 0xd8, 0x73, 0x88, 0xad, 0x5a, 0x4b, 0x63, 0xc5, 0xba, 0x4c, 0xa2, 0xa1,
	0x37, 0xf2, 0xf9, 0x5d, 0x20, 0x9d, 0x41, 0x07, 0x3b, 0xa6, 0x45, 0x2e, 0xf1, 0x24, 0x91, 0xe7,
	0x4e, 0x80, 0x98, 0x93, 0xcd, 0xf6, 0x21, 0x2a, 0x37, 0x8b, 0x73, 0x79, 0x4b, 0x34, 0x63, 0x5e,
	0x7b, 0x2c, 0x81, 0x76, 0x59, 0xa9, 0x2b, 0x4c, 0xf8, 0x94, 0xd8, 0xba, 0xe9, 0x04, 0x42, 0x44,
	0x33, 0xec, 0x35, 0x84, 0x1a, 0x8d, 0xc4, 0x1b, 0xfa, 0xa3, 0xee, 0xe4, 0xd9, 0x03, 0x92, 0x58,
	0xc4, 0x5d, 0x45, 0xfa, 0xcb, 0x83, 0xf6, 0xfc, 0x46, 0x1f, 0x09, 0x2b, 0x58, 0x0a, 0xbd, 0xc3,
	0x2c, 0x2b, 0x36, 0xda, 0x4e, 0x0b, 0x9d, 0x49, 0x62, 0x13, 0xf0, 0x9e, 0x68, 0xc4, 0x50, 0xe9,
	0x59, 0xa5, 0x32, 0x49, 0xa4, 0x7c, 0x1e, 0x96, 0xe8, 0xb0, 0x01, 0x74, 0x8e, 0x85, 0x39, 0x55,
	0x6b, 0x65, 0x89, 0x94, 0xcf, 0x3b, 0xcb, 0xda, 0x47, 0x15, 0xb8, 0xcc, 0x54, 0xa9, 0xa4, 0xb6,
	0x24, 0x65, 0x8f, 0xc7, 0xd5, 0x36, 0x80, 0x53, 0x1e, 0xae, 0x11, 0x9e, 0x04, 0xf5, 0x79, 0x24,
	0xc8, 0xc3, 0x29, 0x67, 0xe2, 0xf6, 0xa2, 0x10, 0x39, 0x29, 0xd7, 0xe3, 0xed, 0xd2, 0xb9, 0x88,
	0xf7, 0x45, 0x2d, 0xb5, 0xb0, 0x9b, 0x4a, 0x26, 0x6d, 0x87, 0x67, 0xb6, 0x01, 0x54, 0xf2, 0x44,
	0x98, 0x55, 0xd2, 0xa1, 0x44, 0xb0, 0x12, 0x66, 0x95, 0xbe, 0x83, 0xee, 0xbc, 0x12, 0xda, 0x88,
	0xcc, 0xaa, 0x42, 0xb3, 0x57, 0x10, 0xe4, 0xc2, 0x8a, 0x7a, 0xdb, 0x58, 0x43, 0x9c, 0x5a, 0x08,
	0x4e, 0xf9, 0xf4, 0xa7, 0x07, 0xfd, 0xdd, 0x9d, 0x9e, 0x48, 0x91, 0x4b, 0xba, 0x94, 0xe9, 0x66,
	0xbd, 0xa8, 0x77, 0xd5, 0xe7, 0x91, 0x26, 0xef, 0x9e, 0x00, 0xad, 0x07, 0x02, 0x24, 0xd0, 0x3e,
	0x16, 0xe6, 0xab, 0x91, 0x79, 0xad, 0x4d, 0x7b, 0xe9, 0x5c, 0x1c, 0x65, 0xbe, 0x5b, 0x90, 0xe0,
	0xc1, 0x82, 0xb0, 0x17, 0x00, 0x33, 0x51, 0x49, 0x6d, 0x69, 0xa0, 0x90, 0x06, 0x82, 0x72, 0x17,
	0xc1, 0x33, 0xe7, 0x37, 0xda, 0x50, 0xd6, 0x69, 0xd4, 0xb1, 0xb5, 0x9f, 0x7e, 0xf7, 0x20, 0xde,
	0x71, 0x67, 0x13, 0x88, 0x1c, 0xff, 0x7a, 0xe6, 0xc1, 0x63, 0x5b, 0xeb, 0x2a, 0x78, 0xb4, 0x72,
	0x93, 0x1e, 0x40, 0x80, 0xe8, 0x49, 0x8b, 0x56, 0x68, 0xbf, 0xa9, 0xd2, 0x9d, 0x96, 0x3c, 0xc0,
	0x13, 0x71, 0x29, 0xf0, 0x4a, 0x4c, 0xe2, 0x0f, 0xfd, 0x51, 0xcc, 0x43, 0xbc, 0x0e, 0x93, 0x7e,
	0x83, 0x78, 0x5e, 0x6d, 0xe4, 0x07, 0x6c, 0x62, 0x6f, 0x20, 0xa2, 0x53, 0xb6, 0x3b, 0xf9, 0xf8,
	0xc3, 0x89, 0xe8, 0xe1, 0x18, 0x76, 0x04, 0xfd, 0x53, 0x61, 0x9a, 0xbc, 0x48, 0xd5, 0x7f, 0x33,
	0xef, 0x5f, 0xdc, 0x6f, 0x49, 0x5f, 0x42, 0xf7, 0x58, 0x6a, 0x59, 0xa9, 0x8c, 0x4b, 0x53, 0xe2,
	0xa3, 0xff, 0x6c, 0x96, 0xf5, 0x1b, 0xf3, 0xd7, 0x66, 0x39, 0x99, 0x42, 0xfc, 0x51, 0x18, 0xeb,
	0x18, 0x1e, 0xc2, 0xff, 0x53, 0x79, 0x3d, 0xbf, 0xd1, 0xdb, 0x6f, 0xe9, 0x2f, 0x33, 0x0f, 0x9a,
	0xf1, 0x06, 0x7e, 0xfa, 0xdf, 0x22, 0xa2, 0x6f, 0xee, 0xed, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd9, 0x31, 0x7e, 0x3e, 0xf9, 0x04, 0x00, 0x00,
}