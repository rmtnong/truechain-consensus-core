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
	return fileDescriptor_fastchain_a2f3fdeaa9c87286, []int{0}
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
	return fileDescriptor_fastchain_a2f3fdeaa9c87286, []int{0, 0}
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
	return fileDescriptor_fastchain_a2f3fdeaa9c87286, []int{1}
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
	return fileDescriptor_fastchain_a2f3fdeaa9c87286, []int{2}
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
	return fileDescriptor_fastchain_a2f3fdeaa9c87286, []int{3}
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
	return fileDescriptor_fastchain_a2f3fdeaa9c87286, []int{4}
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
	Time                 int64    `protobuf:"varint,4,opt,name=Time,json=time,proto3" json:"Time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbftBlockHeader) Reset()         { *m = PbftBlockHeader{} }
func (m *PbftBlockHeader) String() string { return proto.CompactTextString(m) }
func (*PbftBlockHeader) ProtoMessage()    {}
func (*PbftBlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_a2f3fdeaa9c87286, []int{5}
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

func (m *PbftBlockHeader) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type PbftBlock struct {
	Header               *PbftBlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Txns                 []*Transaction   `protobuf:"bytes,2,rep,name=txns,proto3" json:"txns,omitempty"`
	Signs                []string         `protobuf:"bytes,3,rep,name=signs,proto3" json:"signs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PbftBlock) Reset()         { *m = PbftBlock{} }
func (m *PbftBlock) String() string { return proto.CompactTextString(m) }
func (*PbftBlock) ProtoMessage()    {}
func (*PbftBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_a2f3fdeaa9c87286, []int{6}
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
	Blocks               []*PbftBlock     `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Lastheader           *PbftBlockHeader `protobuf:"bytes,2,opt,name=lastheader,proto3" json:"lastheader,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TrueChain) Reset()         { *m = TrueChain{} }
func (m *TrueChain) String() string { return proto.CompactTextString(m) }
func (*TrueChain) ProtoMessage()    {}
func (*TrueChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_a2f3fdeaa9c87286, []int{7}
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

func (m *TrueChain) GetLastheader() *PbftBlockHeader {
	if m != nil {
		return m.Lastheader
	}
	return nil
}

type GenericResp struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResp) Reset()         { *m = GenericResp{} }
func (m *GenericResp) String() string { return proto.CompactTextString(m) }
func (*GenericResp) ProtoMessage()    {}
func (*GenericResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_fastchain_a2f3fdeaa9c87286, []int{8}
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

func init() { proto.RegisterFile("fastchain.proto", fileDescriptor_fastchain_a2f3fdeaa9c87286) }

var fileDescriptor_fastchain_a2f3fdeaa9c87286 = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0x5e, 0x59, 0x7f, 0xd6, 0xd8, 0xbb, 0x59, 0x70, 0x83, 0x80, 0x30, 0x16, 0xa8, 0xa1, 0x43,
	0xe1, 0x06, 0x85, 0x0f, 0x2e, 0x7a, 0xe9, 0x2d, 0x6d, 0xd1, 0xa4, 0x40, 0x61, 0x18, 0xac, 0xfb,
	0x00, 0xb4, 0xc4, 0xd8, 0x44, 0x22, 0x4a, 0x11, 0xa9, 0xfc, 0x1c, 0xfb, 0x1c, 0x7d, 0xba, 0x3e,
	0x49, 0x8b, 0x19, 0xc9, 0x89, 0x1a, 0xa4, 0xe8, 0x6d, 0xbe, 0x99, 0xe1, 0xc7, 0x6f, 0x3e, 0x8e,
	0x04, 0x07, 0xe7, 0xd2, 0xba, 0x6c, 0x27, 0xb5, 0x99, 0x57, 0x75, 0xe9, 0x4a, 0x96, 0xdc, 0x27,
	0xd2, 0x1f, 0x1e, 0xc4, 0x42, 0x5d, 0x35, 0xca, 0x3a, 0x36, 0x87, 0x50, 0x1b, 0xa3, 0x6a, 0xee,
	0x4d, 0xbd, 0xd9, 0x68, 0xc1, 0xe7, 0x0f, 0xe7, 0xba, 0x96, 0xf9, 0x47, 0xac, 0x8b, 0xb6, 0x8d,
	0xfd, 0x0b, 0x7e, 0xae, 0xb7, 0x7c, 0x30, 0xf5, 0x66, 0x63, 0x81, 0x21, 0x3b, 0x84, 0xb0, 0x6c,
	0x9c, 0xaa, 0xb9, 0x4f, 0xb9, 0x16, 0x4c, 0xbe, 0x79, 0x10, 0xd2, 0x41, 0xf6, 0x0f, 0x0c, 0x74,
	0x4e, 0xf4, 0xa1, 0x18, 0xe8, 0x1c, 0x19, 0xac, 0xba, 0x22, 0x86, 0x50, 0x60, 0xc8, 0x18, 0x04,
	0xd7, 0x5a, 0xdd, 0x10, 0x41, 0x28, 0x28, 0xc6, 0x9c, 0xbb, 0xab, 0x14, 0x0f, 0xda, 0x1c, 0xc6,
	0xec, 0x18, 0xc2, 0xcd, 0x65, 0x99, 0x5d, 0xf0, 0x90, 0xb4, 0x1e, 0xf6, 0xb4, 0xae, 0x36, 0xe7,
	0xee, 0x2d, 0xd6, 0x44, 0xdb, 0xc2, 0xfe, 0x87, 0xc4, 0xe9, 0x42, 0x59, 0x27, 0x8b, 0x8a, 0x47,
	0x53, 0x6f, 0xe6, 0x8b, 0x87, 0x44, 0xba, 0x82, 0x21, 0x9e, 0x58, 0x96, 0xb9, 0xc2, 0x9b, 0x64,
	0x9e, 0xb7, 0x06, 0x24, 0x82, 0x62, 0x76, 0x04, 0x51, 0xd5, 0x6c, 0x2e, 0xd4, 0x1d, 0xc9, 0x4c,
	0x44, 0x87, 0x18, 0x87, 0xb8, 0xaa, 0xf5, 0x35, 0x16, 0x7c, 0x2a, 0xec, 0x61, 0xba, 0x80, 0x10,
	0xd9, 0x2c, 0x7b, 0x01, 0xa1, 0xc1, 0x80, 0x7b, 0x53, 0x7f, 0x36, 0x5a, 0xfc, 0xf7, 0x48, 0x24,
	0x36, 0x89, 0xb6, 0x23, 0xfd, 0xee, 0x41, 0xbc, 0xbe, 0x35, 0xef, 0xa5, 0x93, 0x2c, 0x85, 0xf1,
	0x49, 0x96, 0x95, 0x8d, 0x71, 0xcb, 0xd2, 0x64, 0x8a, 0xd4, 0x04, 0x62, 0x2c, 0x7b, 0x39, 0x74,
	0x7a, 0x55, 0xeb, 0x4c, 0x91, 0x28, 0x5f, 0x84, 0x15, 0x02, 0x36, 0x81, 0xe1, 0xa9, 0xb4, 0x9f,
	0x74, 0xa1, 0x1d, 0x89, 0xf2, 0xc5, 0x70, 0xdb, 0x61, 0x74, 0x41, 0xa8, 0x4c, 0x57, 0x5a, 0x19,
	0x47, 0x56, 0x8e, 0x45, 0x52, 0xef, 0x13, 0x38, 0xe5, 0x49, 0x81, 0xf4, 0x64, 0xa8, 0x2f, 0x22,
	0x49, 0x08, 0xa7, 0x5c, 0xc9, 0xbb, 0xcb, 0x52, 0xe6, 0xe4, 0xdc, 0x58, 0xc4, 0x55, 0x0b, 0x91,
	0xef, 0xb3, 0xde, 0x1a, 0xe9, 0x9a, 0x5a, 0xf1, 0xb8, 0xe5, 0xb3, 0xfb, 0x04, 0x3a, 0x79, 0x26,
	0xed, 0x8e, 0x0f, 0xa9, 0x10, 0xec, 0xa4, 0xdd, 0xa5, 0xaf, 0x61, 0xb4, 0xae, 0xa5, 0xb1, 0x32,
	0x73, 0xba, 0x34, 0xec, 0x39, 0x04, 0xb9, 0x74, 0xb2, 0xdb, 0x36, 0xd6, 0x33, 0xa7, 0x33, 0x42,
	0x50, 0x3d, 0xb5, 0x70, 0x70, 0xff, 0xa4, 0x67, 0x4a, 0xe6, 0x8a, 0xde, 0x64, 0xd9, 0x14, 0x9b,
	0x6e, 0x55, 0x7d, 0x11, 0x19, 0x42, 0xbf, 0xcc, 0x3f, 0x78, 0x34, 0x3f, 0x87, 0xf8, 0x54, 0xda,
	0x2f, 0x56, 0xe5, 0x9d, 0x35, 0xf1, 0xb6, 0x85, 0xa8, 0x75, 0xad, 0x8b, 0x76, 0xbf, 0x7c, 0x11,
	0xe0, 0x6a, 0xa4, 0x5f, 0x3d, 0x48, 0xee, 0x6f, 0x65, 0x0b, 0x88, 0x76, 0x74, 0x73, 0x27, 0x76,
	0xf2, 0xd4, 0xba, 0xb5, 0xda, 0x44, 0xd7, 0xc9, 0x8e, 0x21, 0x70, 0xb7, 0xc6, 0xf2, 0x01, 0xbd,
	0xfd, 0x51, 0x7f, 0xbc, 0x07, 0x13, 0x04, 0xf5, 0xe0, 0x6b, 0xa2, 0x75, 0x96, 0xfb, 0x53, 0x7f,
	0x96, 0x88, 0x16, 0xa4, 0x0d, 0x24, 0xeb, 0xba, 0x51, 0xef, 0xf0, 0x10, 0x7b, 0x09, 0x11, 0x6d,
	0xf3, 0x7e, 0x99, 0x9e, 0xde, 0xf8, 0xae, 0x87, 0xbd, 0x01, 0xb8, 0x94, 0xd6, 0x75, 0xa2, 0x07,
	0x7f, 0x14, 0xdd, 0xeb, 0x4e, 0x9f, 0xc1, 0xe8, 0x54, 0x19, 0x55, 0xeb, 0x4c, 0x28, 0x5b, 0xe1,
	0x37, 0x5a, 0xd8, 0x6d, 0xf7, 0x49, 0x60, 0xb8, 0x58, 0x42, 0xf2, 0x41, 0x5a, 0xd7, 0xea, 0x3a,
	0x81, 0xbf, 0x97, 0xea, 0x66, 0x7d, 0x6b, 0xf6, 0x7f, 0x91, 0xdf, 0x4c, 0x3a, 0xe9, 0xe7, 0x7b,
	0xfc, 0xe9, 0x5f, 0x9b, 0x88, 0xfe, 0x4a, 0xaf, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x4d,
	0xef, 0xb2, 0xa8, 0x04, 0x00, 0x00,
}
