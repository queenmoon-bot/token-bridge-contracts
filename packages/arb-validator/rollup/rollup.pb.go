// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rollup.proto

package rollup

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protocol "github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	value "github.com/offchainlabs/arbitrum/packages/arb-util/value"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NodeGraphBuf struct {
	Nodes                []*NodeBuf       `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	OldestNodeHash       *value.HashBuf   `protobuf:"bytes,2,opt,name=oldestNodeHash,proto3" json:"oldestNodeHash,omitempty"`
	LatestConfirmedHash  *value.HashBuf   `protobuf:"bytes,3,opt,name=latestConfirmedHash,proto3" json:"latestConfirmedHash,omitempty"`
	LeafHashes           []*value.HashBuf `protobuf:"bytes,4,rep,name=leafHashes,proto3" json:"leafHashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NodeGraphBuf) Reset()         { *m = NodeGraphBuf{} }
func (m *NodeGraphBuf) String() string { return proto.CompactTextString(m) }
func (*NodeGraphBuf) ProtoMessage()    {}
func (*NodeGraphBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{0}
}

func (m *NodeGraphBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeGraphBuf.Unmarshal(m, b)
}
func (m *NodeGraphBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeGraphBuf.Marshal(b, m, deterministic)
}
func (m *NodeGraphBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeGraphBuf.Merge(m, src)
}
func (m *NodeGraphBuf) XXX_Size() int {
	return xxx_messageInfo_NodeGraphBuf.Size(m)
}
func (m *NodeGraphBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeGraphBuf.DiscardUnknown(m)
}

var xxx_messageInfo_NodeGraphBuf proto.InternalMessageInfo

func (m *NodeGraphBuf) GetNodes() []*NodeBuf {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *NodeGraphBuf) GetOldestNodeHash() *value.HashBuf {
	if m != nil {
		return m.OldestNodeHash
	}
	return nil
}

func (m *NodeGraphBuf) GetLatestConfirmedHash() *value.HashBuf {
	if m != nil {
		return m.LatestConfirmedHash
	}
	return nil
}

func (m *NodeGraphBuf) GetLeafHashes() []*value.HashBuf {
	if m != nil {
		return m.LeafHashes
	}
	return nil
}

type StakedNodeGraphBuf struct {
	NodeGraph            *NodeGraphBuf   `protobuf:"bytes,1,opt,name=nodeGraph,proto3" json:"nodeGraph,omitempty"`
	Stakers              []*StakerBuf    `protobuf:"bytes,2,rep,name=stakers,proto3" json:"stakers,omitempty"`
	Challenges           []*ChallengeBuf `protobuf:"bytes,3,rep,name=challenges,proto3" json:"challenges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StakedNodeGraphBuf) Reset()         { *m = StakedNodeGraphBuf{} }
func (m *StakedNodeGraphBuf) String() string { return proto.CompactTextString(m) }
func (*StakedNodeGraphBuf) ProtoMessage()    {}
func (*StakedNodeGraphBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{1}
}

func (m *StakedNodeGraphBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakedNodeGraphBuf.Unmarshal(m, b)
}
func (m *StakedNodeGraphBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakedNodeGraphBuf.Marshal(b, m, deterministic)
}
func (m *StakedNodeGraphBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakedNodeGraphBuf.Merge(m, src)
}
func (m *StakedNodeGraphBuf) XXX_Size() int {
	return xxx_messageInfo_StakedNodeGraphBuf.Size(m)
}
func (m *StakedNodeGraphBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_StakedNodeGraphBuf.DiscardUnknown(m)
}

var xxx_messageInfo_StakedNodeGraphBuf proto.InternalMessageInfo

func (m *StakedNodeGraphBuf) GetNodeGraph() *NodeGraphBuf {
	if m != nil {
		return m.NodeGraph
	}
	return nil
}

func (m *StakedNodeGraphBuf) GetStakers() []*StakerBuf {
	if m != nil {
		return m.Stakers
	}
	return nil
}

func (m *StakedNodeGraphBuf) GetChallenges() []*ChallengeBuf {
	if m != nil {
		return m.Challenges
	}
	return nil
}

type ChainObserverBuf struct {
	StakedNodeGraph      *StakedNodeGraphBuf `protobuf:"bytes,1,opt,name=stakedNodeGraph,proto3" json:"stakedNodeGraph,omitempty"`
	ContractAddress      []byte              `protobuf:"bytes,2,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	VmParams             *ChainParamsBuf     `protobuf:"bytes,3,opt,name=vmParams,proto3" json:"vmParams,omitempty"`
	PendingInbox         *PendingInboxBuf    `protobuf:"bytes,4,opt,name=pendingInbox,proto3" json:"pendingInbox,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ChainObserverBuf) Reset()         { *m = ChainObserverBuf{} }
func (m *ChainObserverBuf) String() string { return proto.CompactTextString(m) }
func (*ChainObserverBuf) ProtoMessage()    {}
func (*ChainObserverBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{2}
}

func (m *ChainObserverBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainObserverBuf.Unmarshal(m, b)
}
func (m *ChainObserverBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainObserverBuf.Marshal(b, m, deterministic)
}
func (m *ChainObserverBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainObserverBuf.Merge(m, src)
}
func (m *ChainObserverBuf) XXX_Size() int {
	return xxx_messageInfo_ChainObserverBuf.Size(m)
}
func (m *ChainObserverBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainObserverBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChainObserverBuf proto.InternalMessageInfo

func (m *ChainObserverBuf) GetStakedNodeGraph() *StakedNodeGraphBuf {
	if m != nil {
		return m.StakedNodeGraph
	}
	return nil
}

func (m *ChainObserverBuf) GetContractAddress() []byte {
	if m != nil {
		return m.ContractAddress
	}
	return nil
}

func (m *ChainObserverBuf) GetVmParams() *ChainParamsBuf {
	if m != nil {
		return m.VmParams
	}
	return nil
}

func (m *ChainObserverBuf) GetPendingInbox() *PendingInboxBuf {
	if m != nil {
		return m.PendingInbox
	}
	return nil
}

type ChainParamsBuf struct {
	StakeRequirement     *value.BigIntegerBuf `protobuf:"bytes,1,opt,name=stakeRequirement,proto3" json:"stakeRequirement,omitempty"`
	GracePeriod          *RollupTimeBuf       `protobuf:"bytes,2,opt,name=gracePeriod,proto3" json:"gracePeriod,omitempty"`
	MaxExecutionSteps    uint32               `protobuf:"varint,3,opt,name=maxExecutionSteps,proto3" json:"maxExecutionSteps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ChainParamsBuf) Reset()         { *m = ChainParamsBuf{} }
func (m *ChainParamsBuf) String() string { return proto.CompactTextString(m) }
func (*ChainParamsBuf) ProtoMessage()    {}
func (*ChainParamsBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{3}
}

func (m *ChainParamsBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainParamsBuf.Unmarshal(m, b)
}
func (m *ChainParamsBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainParamsBuf.Marshal(b, m, deterministic)
}
func (m *ChainParamsBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainParamsBuf.Merge(m, src)
}
func (m *ChainParamsBuf) XXX_Size() int {
	return xxx_messageInfo_ChainParamsBuf.Size(m)
}
func (m *ChainParamsBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainParamsBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChainParamsBuf proto.InternalMessageInfo

func (m *ChainParamsBuf) GetStakeRequirement() *value.BigIntegerBuf {
	if m != nil {
		return m.StakeRequirement
	}
	return nil
}

func (m *ChainParamsBuf) GetGracePeriod() *RollupTimeBuf {
	if m != nil {
		return m.GracePeriod
	}
	return nil
}

func (m *ChainParamsBuf) GetMaxExecutionSteps() uint32 {
	if m != nil {
		return m.MaxExecutionSteps
	}
	return 0
}

type VMProtoDataBuf struct {
	MachineHash          *value.HashBuf       `protobuf:"bytes,1,opt,name=machineHash,proto3" json:"machineHash,omitempty"`
	InboxHash            *value.HashBuf       `protobuf:"bytes,2,opt,name=inboxHash,proto3" json:"inboxHash,omitempty"`
	PendingTop           *value.HashBuf       `protobuf:"bytes,3,opt,name=pendingTop,proto3" json:"pendingTop,omitempty"`
	PendingCount         *value.BigIntegerBuf `protobuf:"bytes,4,opt,name=pendingCount,proto3" json:"pendingCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VMProtoDataBuf) Reset()         { *m = VMProtoDataBuf{} }
func (m *VMProtoDataBuf) String() string { return proto.CompactTextString(m) }
func (*VMProtoDataBuf) ProtoMessage()    {}
func (*VMProtoDataBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{4}
}

func (m *VMProtoDataBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMProtoDataBuf.Unmarshal(m, b)
}
func (m *VMProtoDataBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMProtoDataBuf.Marshal(b, m, deterministic)
}
func (m *VMProtoDataBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMProtoDataBuf.Merge(m, src)
}
func (m *VMProtoDataBuf) XXX_Size() int {
	return xxx_messageInfo_VMProtoDataBuf.Size(m)
}
func (m *VMProtoDataBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_VMProtoDataBuf.DiscardUnknown(m)
}

var xxx_messageInfo_VMProtoDataBuf proto.InternalMessageInfo

func (m *VMProtoDataBuf) GetMachineHash() *value.HashBuf {
	if m != nil {
		return m.MachineHash
	}
	return nil
}

func (m *VMProtoDataBuf) GetInboxHash() *value.HashBuf {
	if m != nil {
		return m.InboxHash
	}
	return nil
}

func (m *VMProtoDataBuf) GetPendingTop() *value.HashBuf {
	if m != nil {
		return m.PendingTop
	}
	return nil
}

func (m *VMProtoDataBuf) GetPendingCount() *value.BigIntegerBuf {
	if m != nil {
		return m.PendingCount
	}
	return nil
}

type NodeBuf struct {
	Depth                uint64             `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	Hash                 *value.HashBuf     `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	DisputableNode       *DisputableNodeBuf `protobuf:"bytes,3,opt,name=disputableNode,proto3" json:"disputableNode,omitempty"`
	VmProtoData          *VMProtoDataBuf    `protobuf:"bytes,4,opt,name=vmProtoData,proto3" json:"vmProtoData,omitempty"`
	LinkType             uint32             `protobuf:"varint,5,opt,name=linkType,proto3" json:"linkType,omitempty"`
	PrevHash             *value.HashBuf     `protobuf:"bytes,6,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NodeBuf) Reset()         { *m = NodeBuf{} }
func (m *NodeBuf) String() string { return proto.CompactTextString(m) }
func (*NodeBuf) ProtoMessage()    {}
func (*NodeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{5}
}

func (m *NodeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeBuf.Unmarshal(m, b)
}
func (m *NodeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeBuf.Marshal(b, m, deterministic)
}
func (m *NodeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeBuf.Merge(m, src)
}
func (m *NodeBuf) XXX_Size() int {
	return xxx_messageInfo_NodeBuf.Size(m)
}
func (m *NodeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_NodeBuf proto.InternalMessageInfo

func (m *NodeBuf) GetDepth() uint64 {
	if m != nil {
		return m.Depth
	}
	return 0
}

func (m *NodeBuf) GetHash() *value.HashBuf {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *NodeBuf) GetDisputableNode() *DisputableNodeBuf {
	if m != nil {
		return m.DisputableNode
	}
	return nil
}

func (m *NodeBuf) GetVmProtoData() *VMProtoDataBuf {
	if m != nil {
		return m.VmProtoData
	}
	return nil
}

func (m *NodeBuf) GetLinkType() uint32 {
	if m != nil {
		return m.LinkType
	}
	return 0
}

func (m *NodeBuf) GetPrevHash() *value.HashBuf {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

type StakerBuf struct {
	Address              []byte         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Location             *value.HashBuf `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	CreationTime         *RollupTimeBuf `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	ChallengeAddr        []byte         `protobuf:"bytes,4,opt,name=challengeAddr,proto3" json:"challengeAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StakerBuf) Reset()         { *m = StakerBuf{} }
func (m *StakerBuf) String() string { return proto.CompactTextString(m) }
func (*StakerBuf) ProtoMessage()    {}
func (*StakerBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{6}
}

func (m *StakerBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakerBuf.Unmarshal(m, b)
}
func (m *StakerBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakerBuf.Marshal(b, m, deterministic)
}
func (m *StakerBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakerBuf.Merge(m, src)
}
func (m *StakerBuf) XXX_Size() int {
	return xxx_messageInfo_StakerBuf.Size(m)
}
func (m *StakerBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_StakerBuf.DiscardUnknown(m)
}

var xxx_messageInfo_StakerBuf proto.InternalMessageInfo

func (m *StakerBuf) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *StakerBuf) GetLocation() *value.HashBuf {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *StakerBuf) GetCreationTime() *RollupTimeBuf {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *StakerBuf) GetChallengeAddr() []byte {
	if m != nil {
		return m.ChallengeAddr
	}
	return nil
}

type ChallengeBuf struct {
	Contract             []byte   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Asserter             []byte   `protobuf:"bytes,2,opt,name=asserter,proto3" json:"asserter,omitempty"`
	Challenger           []byte   `protobuf:"bytes,3,opt,name=challenger,proto3" json:"challenger,omitempty"`
	Kind                 uint32   `protobuf:"varint,4,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChallengeBuf) Reset()         { *m = ChallengeBuf{} }
func (m *ChallengeBuf) String() string { return proto.CompactTextString(m) }
func (*ChallengeBuf) ProtoMessage()    {}
func (*ChallengeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{7}
}

func (m *ChallengeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChallengeBuf.Unmarshal(m, b)
}
func (m *ChallengeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChallengeBuf.Marshal(b, m, deterministic)
}
func (m *ChallengeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChallengeBuf.Merge(m, src)
}
func (m *ChallengeBuf) XXX_Size() int {
	return xxx_messageInfo_ChallengeBuf.Size(m)
}
func (m *ChallengeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_ChallengeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_ChallengeBuf proto.InternalMessageInfo

func (m *ChallengeBuf) GetContract() []byte {
	if m != nil {
		return m.Contract
	}
	return nil
}

func (m *ChallengeBuf) GetAsserter() []byte {
	if m != nil {
		return m.Asserter
	}
	return nil
}

func (m *ChallengeBuf) GetChallenger() []byte {
	if m != nil {
		return m.Challenger
	}
	return nil
}

func (m *ChallengeBuf) GetKind() uint32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

type DisputableNodeBuf struct {
	NumSteps              uint32                           `protobuf:"varint,1,opt,name=numSteps,proto3" json:"numSteps,omitempty"`
	TimeLowerBound        *RollupTimeBuf                   `protobuf:"bytes,2,opt,name=timeLowerBound,proto3" json:"timeLowerBound,omitempty"`
	TimeUpperBound        *RollupTimeBuf                   `protobuf:"bytes,3,opt,name=timeUpperBound,proto3" json:"timeUpperBound,omitempty"`
	ImportedMessageCount  *value.BigIntegerBuf             `protobuf:"bytes,4,opt,name=importedMessageCount,proto3" json:"importedMessageCount,omitempty"`
	AfterPendingTop       *value.HashBuf                   `protobuf:"bytes,5,opt,name=afterPendingTop,proto3" json:"afterPendingTop,omitempty"`
	ImportedMessagesSlice *value.HashBuf                   `protobuf:"bytes,6,opt,name=importedMessagesSlice,proto3" json:"importedMessagesSlice,omitempty"`
	AssertionStub         *protocol.ExecutionAssertionStub `protobuf:"bytes,7,opt,name=assertionStub,proto3" json:"assertionStub,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                         `json:"-"`
	XXX_unrecognized      []byte                           `json:"-"`
	XXX_sizecache         int32                            `json:"-"`
}

func (m *DisputableNodeBuf) Reset()         { *m = DisputableNodeBuf{} }
func (m *DisputableNodeBuf) String() string { return proto.CompactTextString(m) }
func (*DisputableNodeBuf) ProtoMessage()    {}
func (*DisputableNodeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{8}
}

func (m *DisputableNodeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisputableNodeBuf.Unmarshal(m, b)
}
func (m *DisputableNodeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisputableNodeBuf.Marshal(b, m, deterministic)
}
func (m *DisputableNodeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisputableNodeBuf.Merge(m, src)
}
func (m *DisputableNodeBuf) XXX_Size() int {
	return xxx_messageInfo_DisputableNodeBuf.Size(m)
}
func (m *DisputableNodeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_DisputableNodeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_DisputableNodeBuf proto.InternalMessageInfo

func (m *DisputableNodeBuf) GetNumSteps() uint32 {
	if m != nil {
		return m.NumSteps
	}
	return 0
}

func (m *DisputableNodeBuf) GetTimeLowerBound() *RollupTimeBuf {
	if m != nil {
		return m.TimeLowerBound
	}
	return nil
}

func (m *DisputableNodeBuf) GetTimeUpperBound() *RollupTimeBuf {
	if m != nil {
		return m.TimeUpperBound
	}
	return nil
}

func (m *DisputableNodeBuf) GetImportedMessageCount() *value.BigIntegerBuf {
	if m != nil {
		return m.ImportedMessageCount
	}
	return nil
}

func (m *DisputableNodeBuf) GetAfterPendingTop() *value.HashBuf {
	if m != nil {
		return m.AfterPendingTop
	}
	return nil
}

func (m *DisputableNodeBuf) GetImportedMessagesSlice() *value.HashBuf {
	if m != nil {
		return m.ImportedMessagesSlice
	}
	return nil
}

func (m *DisputableNodeBuf) GetAssertionStub() *protocol.ExecutionAssertionStub {
	if m != nil {
		return m.AssertionStub
	}
	return nil
}

type PendingInboxBuf struct {
	Items                [][]byte       `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	HashOfRest           *value.HashBuf `protobuf:"bytes,2,opt,name=hashOfRest,proto3" json:"hashOfRest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PendingInboxBuf) Reset()         { *m = PendingInboxBuf{} }
func (m *PendingInboxBuf) String() string { return proto.CompactTextString(m) }
func (*PendingInboxBuf) ProtoMessage()    {}
func (*PendingInboxBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{9}
}

func (m *PendingInboxBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PendingInboxBuf.Unmarshal(m, b)
}
func (m *PendingInboxBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PendingInboxBuf.Marshal(b, m, deterministic)
}
func (m *PendingInboxBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingInboxBuf.Merge(m, src)
}
func (m *PendingInboxBuf) XXX_Size() int {
	return xxx_messageInfo_PendingInboxBuf.Size(m)
}
func (m *PendingInboxBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingInboxBuf.DiscardUnknown(m)
}

var xxx_messageInfo_PendingInboxBuf proto.InternalMessageInfo

func (m *PendingInboxBuf) GetItems() [][]byte {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *PendingInboxBuf) GetHashOfRest() *value.HashBuf {
	if m != nil {
		return m.HashOfRest
	}
	return nil
}

type RollupTimeBuf struct {
	Val                  *value.BigIntegerBuf `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RollupTimeBuf) Reset()         { *m = RollupTimeBuf{} }
func (m *RollupTimeBuf) String() string { return proto.CompactTextString(m) }
func (*RollupTimeBuf) ProtoMessage()    {}
func (*RollupTimeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_037f188b50610c79, []int{10}
}

func (m *RollupTimeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RollupTimeBuf.Unmarshal(m, b)
}
func (m *RollupTimeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RollupTimeBuf.Marshal(b, m, deterministic)
}
func (m *RollupTimeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RollupTimeBuf.Merge(m, src)
}
func (m *RollupTimeBuf) XXX_Size() int {
	return xxx_messageInfo_RollupTimeBuf.Size(m)
}
func (m *RollupTimeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_RollupTimeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_RollupTimeBuf proto.InternalMessageInfo

func (m *RollupTimeBuf) GetVal() *value.BigIntegerBuf {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeGraphBuf)(nil), "rollup.NodeGraphBuf")
	proto.RegisterType((*StakedNodeGraphBuf)(nil), "rollup.StakedNodeGraphBuf")
	proto.RegisterType((*ChainObserverBuf)(nil), "rollup.ChainObserverBuf")
	proto.RegisterType((*ChainParamsBuf)(nil), "rollup.ChainParamsBuf")
	proto.RegisterType((*VMProtoDataBuf)(nil), "rollup.VMProtoDataBuf")
	proto.RegisterType((*NodeBuf)(nil), "rollup.NodeBuf")
	proto.RegisterType((*StakerBuf)(nil), "rollup.StakerBuf")
	proto.RegisterType((*ChallengeBuf)(nil), "rollup.ChallengeBuf")
	proto.RegisterType((*DisputableNodeBuf)(nil), "rollup.DisputableNodeBuf")
	proto.RegisterType((*PendingInboxBuf)(nil), "rollup.PendingInboxBuf")
	proto.RegisterType((*RollupTimeBuf)(nil), "rollup.RollupTimeBuf")
}

func init() { proto.RegisterFile("rollup.proto", fileDescriptor_037f188b50610c79) }

var fileDescriptor_037f188b50610c79 = []byte{
	// 944 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0x9b, 0xf4, 0xef, 0xc4, 0x49, 0xb6, 0x43, 0x97, 0x35, 0xbd, 0x40, 0x95, 0x05, 0xa8,
	0x82, 0x25, 0x41, 0x05, 0xb1, 0x45, 0x08, 0xd4, 0xa6, 0x05, 0x76, 0x25, 0x96, 0x8d, 0xdc, 0x02,
	0x12, 0x77, 0x13, 0xfb, 0x24, 0x19, 0xd5, 0x9e, 0x31, 0x33, 0xe3, 0x50, 0x78, 0x02, 0xde, 0x85,
	0x3b, 0xee, 0xb8, 0xe2, 0x3d, 0xb8, 0xe1, 0x21, 0x78, 0x01, 0x34, 0xe3, 0x9f, 0xd8, 0x69, 0xb2,
	0xbd, 0x69, 0xe6, 0xcc, 0xf9, 0xce, 0x9c, 0xdf, 0xef, 0xb8, 0xe0, 0x4a, 0x11, 0xc7, 0x59, 0x3a,
	0x48, 0xa5, 0xd0, 0x82, 0xec, 0xe4, 0xd2, 0xd1, 0xc1, 0x82, 0xc6, 0x19, 0x0e, 0xed, 0xdf, 0x5c,
	0x75, 0xf4, 0xc4, 0xfe, 0x84, 0x22, 0x1e, 0x96, 0x87, 0x5c, 0xe1, 0xff, 0xeb, 0x80, 0xfb, 0x9d,
	0x88, 0xf0, 0x1b, 0x49, 0xd3, 0xf9, 0x28, 0x9b, 0x92, 0x77, 0x61, 0x9b, 0x8b, 0x08, 0x95, 0xe7,
	0x1c, 0xb7, 0x4e, 0x3a, 0xa7, 0xfd, 0x41, 0xe1, 0xc2, 0x80, 0x46, 0xd9, 0x34, 0xc8, 0xb5, 0xe4,
	0x53, 0xe8, 0x89, 0x38, 0x42, 0xa5, 0xcd, 0xfd, 0x73, 0xaa, 0xe6, 0xde, 0xd6, 0xb1, 0x73, 0xd2,
	0x39, 0xed, 0x0d, 0x72, 0xb7, 0xe6, 0xca, 0xc0, 0x57, 0x50, 0xe4, 0x1c, 0xde, 0x88, 0xa9, 0x46,
	0xa5, 0x2f, 0x05, 0x9f, 0x32, 0x99, 0x60, 0x64, 0x8d, 0x5b, 0x6b, 0x8d, 0xd7, 0x41, 0xc9, 0x00,
	0x20, 0x46, 0x3a, 0x35, 0x67, 0x54, 0x5e, 0xdb, 0x46, 0xb9, 0x6a, 0x58, 0x43, 0xf8, 0x7f, 0x38,
	0x40, 0xae, 0x35, 0xbd, 0xc5, 0xa8, 0x91, 0xe7, 0x29, 0xec, 0xf3, 0x52, 0xf6, 0x1c, 0xeb, 0xfe,
	0xb0, 0x9e, 0x6b, 0x09, 0x0c, 0x96, 0x30, 0xf2, 0x01, 0xec, 0x2a, 0xf3, 0x92, 0x54, 0xde, 0x96,
	0xf5, 0x7b, 0x50, 0x5a, 0x58, 0x07, 0xd2, 0xc0, 0x4b, 0x04, 0xf9, 0x04, 0x20, 0x9c, 0xd3, 0x38,
	0x46, 0x3e, 0x43, 0xe5, 0xb5, 0x2c, 0xbe, 0xf2, 0x70, 0x59, 0x6a, 0x6c, 0xb4, 0x4b, 0x9c, 0xff,
	0x9f, 0x03, 0x8f, 0x2e, 0xe7, 0x94, 0xf1, 0x57, 0x13, 0x85, 0x72, 0x61, 0xdf, 0x24, 0x57, 0xd0,
	0x57, 0xcd, 0x0c, 0x8a, 0x88, 0x8f, 0x1a, 0xfe, 0x1b, 0x09, 0x06, 0xab, 0x26, 0xe4, 0x04, 0xfa,
	0xa1, 0xe0, 0x5a, 0xd2, 0x50, 0x5f, 0x44, 0x91, 0x44, 0xa5, 0x6c, 0xcf, 0xdc, 0x60, 0xf5, 0x9a,
	0x9c, 0xc2, 0xde, 0x22, 0x19, 0x53, 0x49, 0x13, 0x55, 0x74, 0xe6, 0xcd, 0x5a, 0xe0, 0x8c, 0xe7,
	0x2a, 0xe3, 0xa4, 0xc2, 0x91, 0xcf, 0xc1, 0x4d, 0x91, 0x47, 0x8c, 0xcf, 0x5e, 0xf0, 0x89, 0xb8,
	0xf3, 0xda, 0xd6, 0xee, 0x49, 0x69, 0x37, 0xae, 0xe9, 0x8c, 0x61, 0x03, 0xec, 0xff, 0xe5, 0x40,
	0xaf, 0xf9, 0x32, 0x39, 0x87, 0x47, 0x36, 0x81, 0x00, 0x7f, 0xce, 0x98, 0xc4, 0x04, 0xb9, 0xae,
	0xda, 0x94, 0x37, 0x7b, 0xc4, 0x66, 0x2f, 0xb8, 0xc6, 0x59, 0x5e, 0xf7, 0x7b, 0x68, 0xf2, 0x0c,
	0x3a, 0x33, 0x49, 0x43, 0x1c, 0xa3, 0x64, 0x22, 0x2a, 0xe6, 0xf3, 0x71, 0x19, 0x50, 0x60, 0x7f,
	0x6e, 0x58, 0x62, 0x5b, 0x50, 0x47, 0x92, 0xa7, 0x70, 0x90, 0xd0, 0xbb, 0xaf, 0xee, 0x30, 0xcc,
	0x34, 0x13, 0xfc, 0x5a, 0x63, 0x9a, 0xd7, 0xa1, 0x1b, 0xdc, 0x57, 0xf8, 0xff, 0x38, 0xd0, 0xfb,
	0xe1, 0xe5, 0xd8, 0xb0, 0xe9, 0x8a, 0x6a, 0x6a, 0x62, 0xff, 0x08, 0x3a, 0x09, 0x0d, 0xe7, 0x8c,
	0xe7, 0xcc, 0x70, 0xd6, 0x0e, 0x77, 0x1d, 0x42, 0x9e, 0xc2, 0x3e, 0x33, 0x95, 0x78, 0x0d, 0x93,
	0x96, 0x00, 0x43, 0x81, 0xa2, 0x7c, 0x37, 0x22, 0xdd, 0xc0, 0x9d, 0x1a, 0x82, 0x9c, 0x55, 0xbd,
	0xb9, 0x14, 0x19, 0xd7, 0x45, 0x6f, 0xd6, 0xd7, 0xb1, 0x81, 0xf4, 0x7f, 0xdf, 0x82, 0xdd, 0x82,
	0xf9, 0xe4, 0x10, 0xb6, 0x23, 0x4c, 0x75, 0x9e, 0x4f, 0x3b, 0xc8, 0x05, 0xe2, 0x43, 0x7b, 0xbe,
	0x39, 0x68, 0xab, 0x23, 0x17, 0xd0, 0x8b, 0x98, 0x4a, 0x33, 0x4d, 0x27, 0x31, 0x9a, 0xe7, 0x8a,
	0x98, 0xdf, 0x2a, 0x9b, 0x71, 0xd5, 0xd0, 0xda, 0xbd, 0xd1, 0x34, 0x20, 0x67, 0xd0, 0x59, 0x24,
	0x55, 0x91, 0x8b, 0x0c, 0xaa, 0xa9, 0x6c, 0xd6, 0x3f, 0xa8, 0x43, 0xc9, 0x11, 0xec, 0xc5, 0x8c,
	0xdf, 0xde, 0xfc, 0x9a, 0xa2, 0xb7, 0x6d, 0x9b, 0x58, 0xc9, 0xe4, 0x7d, 0xd8, 0x4b, 0x25, 0x2e,
	0x6c, 0xd5, 0x77, 0xd6, 0x26, 0x50, 0xe9, 0xfd, 0x3f, 0x1d, 0xd8, 0xaf, 0x68, 0x4e, 0x3c, 0xd8,
	0xa5, 0x05, 0x89, 0x1c, 0x4b, 0xa2, 0x52, 0x34, 0x6f, 0xc6, 0x22, 0xa4, 0x66, 0x40, 0x36, 0x14,
	0xa5, 0xd2, 0x93, 0xcf, 0xc0, 0x0d, 0x25, 0xda, 0xb3, 0x99, 0xc4, 0xa2, 0x2c, 0x1b, 0x66, 0xb4,
	0x01, 0x25, 0xef, 0x40, 0xb7, 0x5a, 0x1b, 0x86, 0xb7, 0xb6, 0x24, 0x6e, 0xd0, 0xbc, 0xf4, 0x7f,
	0x03, 0xb7, 0xbe, 0x6a, 0x4c, 0x31, 0x4a, 0xb2, 0x17, 0x71, 0x57, 0xb2, 0xd1, 0x51, 0xa5, 0x50,
	0x6a, 0x94, 0xc5, 0x62, 0xa8, 0x64, 0xf2, 0x76, 0x6d, 0x99, 0x49, 0x1b, 0xa6, 0x5b, 0x5b, 0x5b,
	0x92, 0x10, 0x68, 0xdf, 0x32, 0x1e, 0xd9, 0x20, 0xba, 0x81, 0x3d, 0xfb, 0x7f, 0xb7, 0xe0, 0xe0,
	0x5e, 0x63, 0x8d, 0x17, 0x9e, 0x25, 0x39, 0xa7, 0x9c, 0xbc, 0x1d, 0xa5, 0x4c, 0xbe, 0x80, 0x9e,
	0x66, 0x09, 0x7e, 0x2b, 0x7e, 0x41, 0x39, 0x12, 0x19, 0x7f, 0x80, 0xb4, 0x2b, 0xe0, 0xd2, 0xfc,
	0xfb, 0x34, 0x2d, 0xcd, 0x5b, 0x0f, 0x9a, 0x2f, 0xc1, 0xe4, 0x39, 0x1c, 0xb2, 0x24, 0x15, 0x52,
	0x63, 0xf4, 0x12, 0x95, 0xa2, 0x33, 0x7c, 0x98, 0x2d, 0x6b, 0x2d, 0xc8, 0x19, 0xf4, 0xe9, 0x54,
	0xa3, 0x1c, 0x2f, 0x49, 0xba, 0xbd, 0x76, 0x12, 0x56, 0x61, 0xe4, 0x0a, 0x1e, 0xaf, 0xbc, 0xa8,
	0xae, 0x63, 0x16, 0xe2, 0x86, 0xe9, 0x5c, 0x0f, 0x26, 0x5f, 0x43, 0x37, 0xef, 0x9c, 0x5d, 0x52,
	0xd9, 0xc4, 0xdb, 0xb5, 0xd6, 0xc7, 0x83, 0xea, 0xe3, 0x5f, 0xed, 0xb0, 0x8b, 0x3a, 0x2e, 0x68,
	0x9a, 0xf9, 0x3f, 0x42, 0x7f, 0x65, 0x6f, 0x9b, 0x25, 0xc0, 0x34, 0x26, 0xf9, 0xbf, 0x07, 0x6e,
	0x90, 0x0b, 0x66, 0x21, 0x19, 0xa2, 0xbf, 0x9a, 0x06, 0xa8, 0xf4, 0x86, 0xa9, 0xaf, 0x21, 0xfc,
	0x67, 0xd0, 0x6d, 0xf4, 0x82, 0xbc, 0x07, 0xad, 0x05, 0x8d, 0x5f, 0xbb, 0xe0, 0x0d, 0x60, 0x74,
	0xfe, 0xd3, 0x97, 0x33, 0xa6, 0xe7, 0xd9, 0x64, 0x10, 0x8a, 0x64, 0x28, 0xa6, 0xd3, 0xd0, 0x7c,
	0x35, 0x62, 0x3a, 0x51, 0x43, 0x2a, 0x27, 0x4c, 0xcb, 0x2c, 0x19, 0xa6, 0x34, 0xbc, 0x35, 0xb5,
	0x30, 0x37, 0x1f, 0x2e, 0x68, 0xcc, 0x22, 0xaa, 0x85, 0x1c, 0xe6, 0x33, 0x30, 0xd9, 0xb1, 0x35,
	0xf8, 0xf8, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0xab, 0x1a, 0x32, 0x3b, 0x09, 0x00, 0x00,
}
