// Code generated by protoc-gen-go.
// source: raft_serverpb.proto
// DO NOT EDIT!

/*
Package raft_serverpb is a generated protocol buffer package.

It is generated from these files:
	raft_serverpb.proto

It has these top-level messages:
	RaftMessage
	RaftTruncatedState
	KeyValue
	RaftSnapshotData
	StoreIdent
	RaftLocalState
	RaftApplyState
	RegionLocalState
*/
package raft_serverpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import raftpb "github.com/pingcap/kvproto/pkg/raftpb"
import metapb "github.com/pingcap/kvproto/pkg/metapb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type PeerState int32

const (
	PeerState_Normal    PeerState = 0
	PeerState_Applying  PeerState = 1
	PeerState_Tombstone PeerState = 2
)

var PeerState_name = map[int32]string{
	0: "Normal",
	1: "Applying",
	2: "Tombstone",
}
var PeerState_value = map[string]int32{
	"Normal":    0,
	"Applying":  1,
	"Tombstone": 2,
}

func (x PeerState) Enum() *PeerState {
	p := new(PeerState)
	*p = x
	return p
}
func (x PeerState) String() string {
	return proto.EnumName(PeerState_name, int32(x))
}
func (x *PeerState) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PeerState_value, data, "PeerState")
	if err != nil {
		return err
	}
	*x = PeerState(value)
	return nil
}
func (PeerState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RaftMessage struct {
	RegionId    *uint64             `protobuf:"varint,1,opt,name=region_id" json:"region_id,omitempty"`
	FromPeer    *metapb.Peer        `protobuf:"bytes,2,opt,name=from_peer" json:"from_peer,omitempty"`
	ToPeer      *metapb.Peer        `protobuf:"bytes,3,opt,name=to_peer" json:"to_peer,omitempty"`
	Message     *raftpb.Message     `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	RegionEpoch *metapb.RegionEpoch `protobuf:"bytes,5,opt,name=region_epoch" json:"region_epoch,omitempty"`
	// true means to_peer is a tombstone peer and it should remove itself.
	IsTombstone      *bool  `protobuf:"varint,6,opt,name=is_tombstone" json:"is_tombstone,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RaftMessage) Reset()                    { *m = RaftMessage{} }
func (m *RaftMessage) String() string            { return proto.CompactTextString(m) }
func (*RaftMessage) ProtoMessage()               {}
func (*RaftMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RaftMessage) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

func (m *RaftMessage) GetFromPeer() *metapb.Peer {
	if m != nil {
		return m.FromPeer
	}
	return nil
}

func (m *RaftMessage) GetToPeer() *metapb.Peer {
	if m != nil {
		return m.ToPeer
	}
	return nil
}

func (m *RaftMessage) GetMessage() *raftpb.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *RaftMessage) GetRegionEpoch() *metapb.RegionEpoch {
	if m != nil {
		return m.RegionEpoch
	}
	return nil
}

func (m *RaftMessage) GetIsTombstone() bool {
	if m != nil && m.IsTombstone != nil {
		return *m.IsTombstone
	}
	return false
}

type RaftTruncatedState struct {
	Index            *uint64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Term             *uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RaftTruncatedState) Reset()                    { *m = RaftTruncatedState{} }
func (m *RaftTruncatedState) String() string            { return proto.CompactTextString(m) }
func (*RaftTruncatedState) ProtoMessage()               {}
func (*RaftTruncatedState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RaftTruncatedState) GetIndex() uint64 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *RaftTruncatedState) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

type KeyValue struct {
	Key              []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RaftSnapshotData struct {
	Region           *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	FileSize         *uint64        `protobuf:"varint,2,opt,name=file_size" json:"file_size,omitempty"`
	Data             []*KeyValue    `protobuf:"bytes,3,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *RaftSnapshotData) Reset()                    { *m = RaftSnapshotData{} }
func (m *RaftSnapshotData) String() string            { return proto.CompactTextString(m) }
func (*RaftSnapshotData) ProtoMessage()               {}
func (*RaftSnapshotData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RaftSnapshotData) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *RaftSnapshotData) GetFileSize() uint64 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *RaftSnapshotData) GetData() []*KeyValue {
	if m != nil {
		return m.Data
	}
	return nil
}

type StoreIdent struct {
	ClusterId        *uint64 `protobuf:"varint,1,opt,name=cluster_id" json:"cluster_id,omitempty"`
	StoreId          *uint64 `protobuf:"varint,2,opt,name=store_id" json:"store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StoreIdent) Reset()                    { *m = StoreIdent{} }
func (m *StoreIdent) String() string            { return proto.CompactTextString(m) }
func (*StoreIdent) ProtoMessage()               {}
func (*StoreIdent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StoreIdent) GetClusterId() uint64 {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return 0
}

func (m *StoreIdent) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

type RaftLocalState struct {
	HardState        *raftpb.HardState `protobuf:"bytes,1,opt,name=hard_state" json:"hard_state,omitempty"`
	LastIndex        *uint64           `protobuf:"varint,2,opt,name=last_index" json:"last_index,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *RaftLocalState) Reset()                    { *m = RaftLocalState{} }
func (m *RaftLocalState) String() string            { return proto.CompactTextString(m) }
func (*RaftLocalState) ProtoMessage()               {}
func (*RaftLocalState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RaftLocalState) GetHardState() *raftpb.HardState {
	if m != nil {
		return m.HardState
	}
	return nil
}

func (m *RaftLocalState) GetLastIndex() uint64 {
	if m != nil && m.LastIndex != nil {
		return *m.LastIndex
	}
	return 0
}

type RaftApplyState struct {
	AppliedIndex     *uint64             `protobuf:"varint,1,opt,name=applied_index" json:"applied_index,omitempty"`
	TruncatedState   *RaftTruncatedState `protobuf:"bytes,2,opt,name=truncated_state" json:"truncated_state,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *RaftApplyState) Reset()                    { *m = RaftApplyState{} }
func (m *RaftApplyState) String() string            { return proto.CompactTextString(m) }
func (*RaftApplyState) ProtoMessage()               {}
func (*RaftApplyState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RaftApplyState) GetAppliedIndex() uint64 {
	if m != nil && m.AppliedIndex != nil {
		return *m.AppliedIndex
	}
	return 0
}

func (m *RaftApplyState) GetTruncatedState() *RaftTruncatedState {
	if m != nil {
		return m.TruncatedState
	}
	return nil
}

type RegionLocalState struct {
	State            *PeerState     `protobuf:"varint,1,opt,name=state,enum=raft_serverpb.PeerState" json:"state,omitempty"`
	Region           *metapb.Region `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *RegionLocalState) Reset()                    { *m = RegionLocalState{} }
func (m *RegionLocalState) String() string            { return proto.CompactTextString(m) }
func (*RegionLocalState) ProtoMessage()               {}
func (*RegionLocalState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RegionLocalState) GetState() PeerState {
	if m != nil && m.State != nil {
		return *m.State
	}
	return PeerState_Normal
}

func (m *RegionLocalState) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func init() {
	proto.RegisterType((*RaftMessage)(nil), "raft_serverpb.RaftMessage")
	proto.RegisterType((*RaftTruncatedState)(nil), "raft_serverpb.RaftTruncatedState")
	proto.RegisterType((*KeyValue)(nil), "raft_serverpb.KeyValue")
	proto.RegisterType((*RaftSnapshotData)(nil), "raft_serverpb.RaftSnapshotData")
	proto.RegisterType((*StoreIdent)(nil), "raft_serverpb.StoreIdent")
	proto.RegisterType((*RaftLocalState)(nil), "raft_serverpb.RaftLocalState")
	proto.RegisterType((*RaftApplyState)(nil), "raft_serverpb.RaftApplyState")
	proto.RegisterType((*RegionLocalState)(nil), "raft_serverpb.RegionLocalState")
	proto.RegisterEnum("raft_serverpb.PeerState", PeerState_name, PeerState_value)
}

var fileDescriptor0 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x25, 0x89, 0x93, 0x26, 0x13, 0x27, 0x75, 0xb7, 0x20, 0xac, 0x4a, 0x40, 0xb1, 0x54, 0xbe,
	0x0e, 0x91, 0x88, 0x38, 0x71, 0x43, 0x02, 0x09, 0x54, 0x40, 0xa8, 0xa9, 0xb8, 0x70, 0xb0, 0xb6,
	0xf6, 0x24, 0x59, 0xb1, 0xf6, 0x5a, 0xbb, 0x9b, 0x8a, 0xf0, 0xf7, 0xf8, 0x63, 0xec, 0x87, 0x57,
	0x69, 0xda, 0x5e, 0x2c, 0xcd, 0xbc, 0x37, 0xcf, 0x6f, 0xde, 0x2c, 0x1c, 0x4b, 0xba, 0xd4, 0xb9,
	0x42, 0x79, 0x8d, 0xb2, 0xb9, 0x9a, 0x35, 0x52, 0x68, 0x41, 0x26, 0x7b, 0xcd, 0x93, 0xd8, 0x96,
	0x01, 0x3c, 0x89, 0x2b, 0xd4, 0x34, 0x54, 0xd9, 0xbf, 0x0e, 0x8c, 0x2f, 0x0c, 0xfc, 0x0d, 0x95,
	0xa2, 0x2b, 0x24, 0x47, 0x30, 0x92, 0xb8, 0x62, 0xa2, 0xce, 0x59, 0x99, 0x76, 0x4e, 0x3b, 0xaf,
	0x22, 0xf2, 0x0c, 0x46, 0x4b, 0x29, 0xaa, 0xbc, 0x41, 0x94, 0x69, 0xd7, 0xb4, 0xc6, 0xf3, 0x78,
	0xd6, 0x8a, 0xfc, 0x30, 0x3d, 0xf2, 0x04, 0x0e, 0xb4, 0xf0, 0x70, 0xef, 0x1e, 0xf8, 0x14, 0x0e,
	0x2a, 0xaf, 0x9e, 0x46, 0x0e, 0x3e, 0x9c, 0xb5, 0x86, 0xc2, 0x4f, 0x5f, 0x43, 0xdc, 0xfe, 0x14,
	0x1b, 0x51, 0xac, 0xd3, 0xbe, 0xa3, 0x1d, 0x07, 0x95, 0x0b, 0x87, 0x7d, 0xb2, 0x10, 0x79, 0x08,
	0x31, 0x53, 0xb9, 0x16, 0xd5, 0x95, 0xd2, 0xa2, 0xc6, 0x74, 0x60, 0xa8, 0xc3, 0xec, 0x2d, 0x10,
	0xbb, 0xc4, 0xa5, 0xdc, 0xd4, 0x05, 0xd5, 0x58, 0x2e, 0xb4, 0xf9, 0x92, 0x09, 0xf4, 0x59, 0x5d,
	0xe2, 0x9f, 0x76, 0x8f, 0x18, 0x22, 0x8d, 0xb2, 0x72, 0x2b, 0x44, 0xd9, 0x0b, 0x18, 0x9e, 0xe3,
	0xf6, 0x27, 0xe5, 0x1b, 0x24, 0x63, 0xe8, 0xfd, 0xc6, 0xad, 0xa3, 0xc5, 0x76, 0xea, 0xda, 0x76,
	0x1d, 0x2f, 0xce, 0x38, 0x24, 0x56, 0x7a, 0x51, 0xd3, 0x46, 0xad, 0x85, 0xfe, 0x48, 0x35, 0x25,
	0x4f, 0x61, 0xe0, 0xfd, 0xba, 0x91, 0xf1, 0x7c, 0xba, 0xef, 0xd4, 0x86, 0xb8, 0x64, 0x1c, 0x73,
	0xc5, 0xfe, 0x7a, 0x99, 0x88, 0x9c, 0x41, 0x54, 0x9a, 0x51, 0x13, 0x50, 0xcf, 0x0c, 0x3c, 0x9e,
	0xed, 0x9f, 0x2d, 0x38, 0xc9, 0xe6, 0x00, 0x0b, 0x2d, 0x24, 0x7e, 0x29, 0xb1, 0xd6, 0x84, 0x00,
	0x14, 0x7c, 0xa3, 0x8c, 0xeb, 0xdd, 0x35, 0x12, 0x18, 0x2a, 0xcb, 0xb0, 0x1d, 0xbf, 0xc9, 0x39,
	0x4c, 0xad, 0xc3, 0xaf, 0xa2, 0xa0, 0xdc, 0x2f, 0x7e, 0x06, 0xb0, 0xa6, 0xb2, 0xcc, 0x95, 0xad,
	0x5a, 0x8f, 0x47, 0x21, 0xf4, 0xcf, 0x06, 0xf1, 0x34, 0x23, 0xcf, 0xa9, 0xd2, 0xb9, 0x0f, 0xc9,
	0x8b, 0x15, 0x5e, 0xec, 0x43, 0xd3, 0xf0, 0xad, 0x67, 0x3d, 0x82, 0x09, 0x35, 0x15, 0xc3, 0x32,
	0xbf, 0x99, 0xe6, 0x7b, 0x38, 0xd4, 0x21, 0xee, 0xf6, 0x47, 0xfe, 0x6d, 0x3c, 0xbf, 0xb5, 0xdb,
	0xdd, 0xc3, 0x64, 0xbf, 0x4c, 0xa6, 0x2e, 0xa9, 0x1b, 0x9e, 0x5f, 0x42, 0x7f, 0x67, 0x77, 0x3a,
	0x4f, 0x6f, 0xa9, 0xd8, 0x97, 0xe4, 0x89, 0xbb, 0xf0, 0xbb, 0xf7, 0x85, 0xff, 0xe6, 0x1d, 0x8c,
	0x76, 0x64, 0x80, 0xc1, 0x77, 0x21, 0x2b, 0xca, 0x93, 0x07, 0xe6, 0xfe, 0x43, 0xb7, 0x16, 0xab,
	0x57, 0x49, 0xc7, 0x9c, 0x79, 0x74, 0x19, 0x5e, 0x51, 0xd2, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff,
	0x2d, 0x8d, 0x51, 0xf5, 0x48, 0x03, 0x00, 0x00,
}