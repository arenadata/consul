// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: private/pbservice/node.proto

package pbservice

import (
	pbcommon "github.com/arenadata/consul/proto/private/pbcommon"
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

// IndexedCheckServiceNodes is used to return multiple instances for a given service.
type IndexedCheckServiceNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index uint64              `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Nodes []*CheckServiceNode `protobuf:"bytes,2,rep,name=Nodes,proto3" json:"Nodes,omitempty"`
}

func (x *IndexedCheckServiceNodes) Reset() {
	*x = IndexedCheckServiceNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbservice_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexedCheckServiceNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexedCheckServiceNodes) ProtoMessage() {}

func (x *IndexedCheckServiceNodes) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbservice_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexedCheckServiceNodes.ProtoReflect.Descriptor instead.
func (*IndexedCheckServiceNodes) Descriptor() ([]byte, []int) {
	return file_private_pbservice_node_proto_rawDescGZIP(), []int{0}
}

func (x *IndexedCheckServiceNodes) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *IndexedCheckServiceNodes) GetNodes() []*CheckServiceNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// CheckServiceNode is used to provide the node, its service
// definition, as well as a HealthCheck that is associated.
type CheckServiceNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node    *Node          `protobuf:"bytes,1,opt,name=Node,proto3" json:"Node,omitempty"`
	Service *NodeService   `protobuf:"bytes,2,opt,name=Service,proto3" json:"Service,omitempty"`
	Checks  []*HealthCheck `protobuf:"bytes,3,rep,name=Checks,proto3" json:"Checks,omitempty"`
}

func (x *CheckServiceNode) Reset() {
	*x = CheckServiceNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbservice_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckServiceNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckServiceNode) ProtoMessage() {}

func (x *CheckServiceNode) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbservice_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckServiceNode.ProtoReflect.Descriptor instead.
func (*CheckServiceNode) Descriptor() ([]byte, []int) {
	return file_private_pbservice_node_proto_rawDescGZIP(), []int{1}
}

func (x *CheckServiceNode) GetNode() *Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *CheckServiceNode) GetService() *NodeService {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *CheckServiceNode) GetChecks() []*HealthCheck {
	if x != nil {
		return x.Checks
	}
	return nil
}

// Node contains information about a node.
//
// mog annotation:
//
// target=github.com/arenadata/consul/agent/structs.Node
// output=node.gen.go
// name=Structs
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// mog: func-to=NodeIDType func-from=string
	ID              string            `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Node            string            `protobuf:"bytes,2,opt,name=Node,proto3" json:"Node,omitempty"`
	Partition       string            `protobuf:"bytes,8,opt,name=Partition,proto3" json:"Partition,omitempty"`
	PeerName        string            `protobuf:"bytes,9,opt,name=PeerName,proto3" json:"PeerName,omitempty"`
	Address         string            `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	Datacenter      string            `protobuf:"bytes,4,opt,name=Datacenter,proto3" json:"Datacenter,omitempty"`
	TaggedAddresses map[string]string `protobuf:"bytes,5,rep,name=TaggedAddresses,proto3" json:"TaggedAddresses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Meta            map[string]string `protobuf:"bytes,6,rep,name=Meta,proto3" json:"Meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// mog: func-to=RaftIndexToStructs func-from=NewRaftIndexFromStructs
	RaftIndex *pbcommon.RaftIndex `protobuf:"bytes,7,opt,name=RaftIndex,proto3" json:"RaftIndex,omitempty"`
	// Locality identifies where the node is running.
	// mog: func-to=LocalityToStructs func-from=LocalityFromStructs
	Locality *pbcommon.Locality `protobuf:"bytes,10,opt,name=Locality,proto3" json:"Locality,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbservice_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbservice_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_private_pbservice_node_proto_rawDescGZIP(), []int{2}
}

func (x *Node) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Node) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *Node) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *Node) GetPeerName() string {
	if x != nil {
		return x.PeerName
	}
	return ""
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Node) GetDatacenter() string {
	if x != nil {
		return x.Datacenter
	}
	return ""
}

func (x *Node) GetTaggedAddresses() map[string]string {
	if x != nil {
		return x.TaggedAddresses
	}
	return nil
}

func (x *Node) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Node) GetRaftIndex() *pbcommon.RaftIndex {
	if x != nil {
		return x.RaftIndex
	}
	return nil
}

func (x *Node) GetLocality() *pbcommon.Locality {
	if x != nil {
		return x.Locality
	}
	return nil
}

// NodeService is a service provided by a node
//
// mog annotation:
//
// target=github.com/arenadata/consul/agent/structs.NodeService
// output=node.gen.go
// name=Structs
type NodeService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kind is the kind of service this is. Different kinds of services may
	// have differing validation, DNS behavior, etc. An empty kind will default
	// to the Default kind. See ServiceKind for the full list of kinds.
	// mog: func-to=structs.ServiceKind func-from=string
	Kind    string   `protobuf:"bytes,1,opt,name=Kind,proto3" json:"Kind,omitempty"`
	ID      string   `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Service string   `protobuf:"bytes,3,opt,name=Service,proto3" json:"Service,omitempty"`
	Tags    []string `protobuf:"bytes,4,rep,name=Tags,proto3" json:"Tags,omitempty"`
	Address string   `protobuf:"bytes,5,opt,name=Address,proto3" json:"Address,omitempty"`
	// mog: func-to=MapStringServiceAddressToStructs func-from=NewMapStringServiceAddressFromStructs
	TaggedAddresses map[string]*ServiceAddress `protobuf:"bytes,15,rep,name=TaggedAddresses,proto3" json:"TaggedAddresses,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Meta            map[string]string          `protobuf:"bytes,6,rep,name=Meta,proto3" json:"Meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// mog: func-to=int func-from=int32
	Port       int32  `protobuf:"varint,7,opt,name=Port,proto3" json:"Port,omitempty"`
	SocketPath string `protobuf:"bytes,17,opt,name=SocketPath,proto3" json:"SocketPath,omitempty"`
	// mog: func-to=WeightsPtrToStructs func-from=NewWeightsPtrFromStructs
	Weights           *Weights `protobuf:"bytes,8,opt,name=Weights,proto3" json:"Weights,omitempty"`
	EnableTagOverride bool     `protobuf:"varint,9,opt,name=EnableTagOverride,proto3" json:"EnableTagOverride,omitempty"`
	// Proxy is the configuration set for Kind = connect-proxy. It is mandatory in
	// that case and an error to be set for any other kind. This config is part of
	// a proxy service definition and is distinct from but shares some fields with
	// the Connect.Proxy which configures a managed proxy as part of the actual
	// service's definition. This duplication is ugly but seemed better than the
	// alternative which was to re-use the same struct fields for both cases even
	// though the semantics are different and the non-shred fields make no sense
	// in the other case. ProxyConfig may be a more natural name here, but it's
	// confusing for the UX because one of the fields in ConnectProxyConfig is
	// also called just "Config"
	Proxy *ConnectProxyConfig `protobuf:"bytes,11,opt,name=Proxy,proto3" json:"Proxy,omitempty"`
	// Connect are the Connect settings for a service. This is purposely NOT
	// a pointer so that we never have to nil-check this.
	Connect *ServiceConnect `protobuf:"bytes,12,opt,name=Connect,proto3" json:"Connect,omitempty"`
	// LocallyRegisteredAsSidecar is private as it is only used by a local agent
	// state to track if the service was registered from a nested sidecar_service
	// block. We need to track that so we can know whether we need to deregister
	// it automatically too if it's removed from the service definition or if the
	// parent service is deregistered. Relying only on ID would cause us to
	// deregister regular services if they happen to be registered using the same
	// ID scheme as our sidecars do by default. We could use meta but that gets
	// unpleasant because we can't use the consul- prefix from an agent (reserved
	// for use internally but in practice that means within the state store or in
	// responses only), and it leaks the detail publicly which people might rely
	// on which is a bit unpleasant for something that is meant to be config-file
	// syntax sugar. Note this is not translated to ServiceNode and friends and
	// may not be set on a NodeService that isn't the one the agent registered and
	// keeps in it's local state. We never want this rendered in JSON as it's
	// internal only. Right now our agent endpoints return api structs which don't
	// include it but this is a safety net incase we change that or there is
	// somewhere this is used in API output.
	LocallyRegisteredAsSidecar bool `protobuf:"varint,13,opt,name=LocallyRegisteredAsSidecar,proto3" json:"LocallyRegisteredAsSidecar,omitempty"`
	// mog: func-to=EnterpriseMetaToStructs func-from=NewEnterpriseMetaFromStructs
	EnterpriseMeta *pbcommon.EnterpriseMeta `protobuf:"bytes,16,opt,name=EnterpriseMeta,proto3" json:"EnterpriseMeta,omitempty"`
	PeerName       string                   `protobuf:"bytes,18,opt,name=PeerName,proto3" json:"PeerName,omitempty"`
	// mog: func-to=RaftIndexToStructs func-from=NewRaftIndexFromStructs
	RaftIndex *pbcommon.RaftIndex `protobuf:"bytes,14,opt,name=RaftIndex,proto3" json:"RaftIndex,omitempty"`
	// Locality identifies where the service is running.
	// mog: func-to=LocalityToStructs func-from=LocalityFromStructs
	Locality *pbcommon.Locality `protobuf:"bytes,19,opt,name=Locality,proto3" json:"Locality,omitempty"`
}

func (x *NodeService) Reset() {
	*x = NodeService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbservice_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeService) ProtoMessage() {}

func (x *NodeService) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbservice_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeService.ProtoReflect.Descriptor instead.
func (*NodeService) Descriptor() ([]byte, []int) {
	return file_private_pbservice_node_proto_rawDescGZIP(), []int{3}
}

func (x *NodeService) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *NodeService) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *NodeService) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *NodeService) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *NodeService) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *NodeService) GetTaggedAddresses() map[string]*ServiceAddress {
	if x != nil {
		return x.TaggedAddresses
	}
	return nil
}

func (x *NodeService) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *NodeService) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *NodeService) GetSocketPath() string {
	if x != nil {
		return x.SocketPath
	}
	return ""
}

func (x *NodeService) GetWeights() *Weights {
	if x != nil {
		return x.Weights
	}
	return nil
}

func (x *NodeService) GetEnableTagOverride() bool {
	if x != nil {
		return x.EnableTagOverride
	}
	return false
}

func (x *NodeService) GetProxy() *ConnectProxyConfig {
	if x != nil {
		return x.Proxy
	}
	return nil
}

func (x *NodeService) GetConnect() *ServiceConnect {
	if x != nil {
		return x.Connect
	}
	return nil
}

func (x *NodeService) GetLocallyRegisteredAsSidecar() bool {
	if x != nil {
		return x.LocallyRegisteredAsSidecar
	}
	return false
}

func (x *NodeService) GetEnterpriseMeta() *pbcommon.EnterpriseMeta {
	if x != nil {
		return x.EnterpriseMeta
	}
	return nil
}

func (x *NodeService) GetPeerName() string {
	if x != nil {
		return x.PeerName
	}
	return ""
}

func (x *NodeService) GetRaftIndex() *pbcommon.RaftIndex {
	if x != nil {
		return x.RaftIndex
	}
	return nil
}

func (x *NodeService) GetLocality() *pbcommon.Locality {
	if x != nil {
		return x.Locality
	}
	return nil
}

var File_private_pbservice_node_proto protoreflect.FileDescriptor

var file_private_pbservice_node_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x1d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x18, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x49, 0x0a, 0x05, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x22, 0xe1, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x46, 0x0a, 0x06, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x06, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x22, 0xdd, 0x04, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x65, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x65, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x61,
	0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x0f, 0x54, 0x61, 0x67, 0x67,
	0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3c, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x67, 0x65,
	0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0f, 0x54, 0x61, 0x67, 0x67, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x12, 0x45, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x49, 0x0a, 0x09, 0x52, 0x61, 0x66, 0x74, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x61,
	0x66, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x09, 0x52, 0x61, 0x66, 0x74, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x46, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x1a, 0x42, 0x0a, 0x14, 0x54, 0x61,
	0x67, 0x67, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37,
	0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf1, 0x08, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x6d, 0x0a, 0x0f, 0x54, 0x61, 0x67, 0x67, 0x65, 0x64, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x67,
	0x67, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0f, 0x54, 0x61, 0x67, 0x67, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x4c, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x44, 0x0a, 0x07, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x52, 0x07, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x67, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x67,
	0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x4b, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x4b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x3e, 0x0a, 0x1a, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x73, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x41, 0x73, 0x53, 0x69, 0x64, 0x65, 0x63,
	0x61, 0x72, 0x12, 0x58, 0x0a, 0x0e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0e, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x65, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x65, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x52, 0x61, 0x66, 0x74,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x61, 0x66, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x09, 0x52, 0x61, 0x66, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x46, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x1a, 0x75, 0x0a, 0x14, 0x54,
	0x61, 0x67, 0x67, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x47, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x8f, 0x02, 0x0a, 0x25,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xa2, 0x02, 0x04, 0x48, 0x43, 0x49, 0x53, 0xaa, 0x02,
	0x21, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0xca, 0x02, 0x21, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x2d, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x24, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_private_pbservice_node_proto_rawDescOnce sync.Once
	file_private_pbservice_node_proto_rawDescData = file_private_pbservice_node_proto_rawDesc
)

func file_private_pbservice_node_proto_rawDescGZIP() []byte {
	file_private_pbservice_node_proto_rawDescOnce.Do(func() {
		file_private_pbservice_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_private_pbservice_node_proto_rawDescData)
	})
	return file_private_pbservice_node_proto_rawDescData
}

var file_private_pbservice_node_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_private_pbservice_node_proto_goTypes = []interface{}{
	(*IndexedCheckServiceNodes)(nil), // 0: hashicorp.consul.internal.service.IndexedCheckServiceNodes
	(*CheckServiceNode)(nil),         // 1: hashicorp.consul.internal.service.CheckServiceNode
	(*Node)(nil),                     // 2: hashicorp.consul.internal.service.Node
	(*NodeService)(nil),              // 3: hashicorp.consul.internal.service.NodeService
	nil,                              // 4: hashicorp.consul.internal.service.Node.TaggedAddressesEntry
	nil,                              // 5: hashicorp.consul.internal.service.Node.MetaEntry
	nil,                              // 6: hashicorp.consul.internal.service.NodeService.TaggedAddressesEntry
	nil,                              // 7: hashicorp.consul.internal.service.NodeService.MetaEntry
	(*HealthCheck)(nil),              // 8: hashicorp.consul.internal.service.HealthCheck
	(*pbcommon.RaftIndex)(nil),       // 9: hashicorp.consul.internal.common.RaftIndex
	(*pbcommon.Locality)(nil),        // 10: hashicorp.consul.internal.common.Locality
	(*Weights)(nil),                  // 11: hashicorp.consul.internal.service.Weights
	(*ConnectProxyConfig)(nil),       // 12: hashicorp.consul.internal.service.ConnectProxyConfig
	(*ServiceConnect)(nil),           // 13: hashicorp.consul.internal.service.ServiceConnect
	(*pbcommon.EnterpriseMeta)(nil),  // 14: hashicorp.consul.internal.common.EnterpriseMeta
	(*ServiceAddress)(nil),           // 15: hashicorp.consul.internal.service.ServiceAddress
}
var file_private_pbservice_node_proto_depIdxs = []int32{
	1,  // 0: hashicorp.consul.internal.service.IndexedCheckServiceNodes.Nodes:type_name -> hashicorp.consul.internal.service.CheckServiceNode
	2,  // 1: hashicorp.consul.internal.service.CheckServiceNode.Node:type_name -> hashicorp.consul.internal.service.Node
	3,  // 2: hashicorp.consul.internal.service.CheckServiceNode.Service:type_name -> hashicorp.consul.internal.service.NodeService
	8,  // 3: hashicorp.consul.internal.service.CheckServiceNode.Checks:type_name -> hashicorp.consul.internal.service.HealthCheck
	4,  // 4: hashicorp.consul.internal.service.Node.TaggedAddresses:type_name -> hashicorp.consul.internal.service.Node.TaggedAddressesEntry
	5,  // 5: hashicorp.consul.internal.service.Node.Meta:type_name -> hashicorp.consul.internal.service.Node.MetaEntry
	9,  // 6: hashicorp.consul.internal.service.Node.RaftIndex:type_name -> hashicorp.consul.internal.common.RaftIndex
	10, // 7: hashicorp.consul.internal.service.Node.Locality:type_name -> hashicorp.consul.internal.common.Locality
	6,  // 8: hashicorp.consul.internal.service.NodeService.TaggedAddresses:type_name -> hashicorp.consul.internal.service.NodeService.TaggedAddressesEntry
	7,  // 9: hashicorp.consul.internal.service.NodeService.Meta:type_name -> hashicorp.consul.internal.service.NodeService.MetaEntry
	11, // 10: hashicorp.consul.internal.service.NodeService.Weights:type_name -> hashicorp.consul.internal.service.Weights
	12, // 11: hashicorp.consul.internal.service.NodeService.Proxy:type_name -> hashicorp.consul.internal.service.ConnectProxyConfig
	13, // 12: hashicorp.consul.internal.service.NodeService.Connect:type_name -> hashicorp.consul.internal.service.ServiceConnect
	14, // 13: hashicorp.consul.internal.service.NodeService.EnterpriseMeta:type_name -> hashicorp.consul.internal.common.EnterpriseMeta
	9,  // 14: hashicorp.consul.internal.service.NodeService.RaftIndex:type_name -> hashicorp.consul.internal.common.RaftIndex
	10, // 15: hashicorp.consul.internal.service.NodeService.Locality:type_name -> hashicorp.consul.internal.common.Locality
	15, // 16: hashicorp.consul.internal.service.NodeService.TaggedAddressesEntry.value:type_name -> hashicorp.consul.internal.service.ServiceAddress
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_private_pbservice_node_proto_init() }
func file_private_pbservice_node_proto_init() {
	if File_private_pbservice_node_proto != nil {
		return
	}
	file_private_pbservice_healthcheck_proto_init()
	file_private_pbservice_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_private_pbservice_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexedCheckServiceNodes); i {
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
		file_private_pbservice_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckServiceNode); i {
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
		file_private_pbservice_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_private_pbservice_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeService); i {
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
			RawDescriptor: file_private_pbservice_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_private_pbservice_node_proto_goTypes,
		DependencyIndexes: file_private_pbservice_node_proto_depIdxs,
		MessageInfos:      file_private_pbservice_node_proto_msgTypes,
	}.Build()
	File_private_pbservice_node_proto = out.File
	file_private_pbservice_node_proto_rawDesc = nil
	file_private_pbservice_node_proto_goTypes = nil
	file_private_pbservice_node_proto_depIdxs = nil
}
