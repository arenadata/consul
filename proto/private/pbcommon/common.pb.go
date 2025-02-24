// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: private/pbcommon/common.proto

package pbcommon

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RaftIndex is used to track the index used while creating
// or modifying a given struct type.
//
// mog annotation:
//
// target=github.com/shulutkov/yellow-pages/agent/structs.RaftIndex
// output=common.gen.go
// name=Structs
// ignore-fields=state,sizeCache,unknownFields
type RaftIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: bexpr:"-"
	CreateIndex uint64 `protobuf:"varint,1,opt,name=CreateIndex,proto3" json:"CreateIndex,omitempty" bexpr:"-"`
	// @gotags: bexpr:"-"
	ModifyIndex uint64 `protobuf:"varint,2,opt,name=ModifyIndex,proto3" json:"ModifyIndex,omitempty" bexpr:"-"`
}

func (x *RaftIndex) Reset() {
	*x = RaftIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbcommon_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaftIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaftIndex) ProtoMessage() {}

func (x *RaftIndex) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbcommon_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaftIndex.ProtoReflect.Descriptor instead.
func (*RaftIndex) Descriptor() ([]byte, []int) {
	return file_private_pbcommon_common_proto_rawDescGZIP(), []int{0}
}

func (x *RaftIndex) GetCreateIndex() uint64 {
	if x != nil {
		return x.CreateIndex
	}
	return 0
}

func (x *RaftIndex) GetModifyIndex() uint64 {
	if x != nil {
		return x.ModifyIndex
	}
	return 0
}

// TargetDatacenter is intended to be used within other messages used for RPC routing
// amongst the various Consul datacenters
type TargetDatacenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datacenter string `protobuf:"bytes,1,opt,name=Datacenter,proto3" json:"Datacenter,omitempty"`
}

func (x *TargetDatacenter) Reset() {
	*x = TargetDatacenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbcommon_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetDatacenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetDatacenter) ProtoMessage() {}

func (x *TargetDatacenter) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbcommon_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetDatacenter.ProtoReflect.Descriptor instead.
func (*TargetDatacenter) Descriptor() ([]byte, []int) {
	return file_private_pbcommon_common_proto_rawDescGZIP(), []int{1}
}

func (x *TargetDatacenter) GetDatacenter() string {
	if x != nil {
		return x.Datacenter
	}
	return ""
}

// mog annotation:
//
// target=github.com/shulutkov/yellow-pages/agent/structs.WriteRequest
// output=common.gen.go
// name=Structs
// ignore-fields=state,sizeCache,unknownFields
type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token is the ACL token ID. If not provided, the 'anonymous'
	// token is assumed for backwards compatibility.
	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbcommon_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbcommon_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_private_pbcommon_common_proto_rawDescGZIP(), []int{2}
}

func (x *WriteRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// ReadRequest is a type that may be embedded into any requests for read
// operations.
// It is a replacement for QueryOptions now that we no longer need any of those
// fields because we are moving away from using blocking queries.
// It is also similar to WriteRequest. It is a separate type so that in the
// future we can introduce fields that may only be relevant for reads.
type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token is the ACL token ID. If not provided, the 'anonymous'
	// token is assumed for backwards compatibility.
	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	// RequireConsistent indicates that the request must be sent to the leader.
	RequireConsistent bool `protobuf:"varint,2,opt,name=RequireConsistent,proto3" json:"RequireConsistent,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbcommon_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbcommon_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_private_pbcommon_common_proto_rawDescGZIP(), []int{3}
}

func (x *ReadRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ReadRequest) GetRequireConsistent() bool {
	if x != nil {
		return x.RequireConsistent
	}
	return false
}

// QueryOptions is used to specify various flags for read queries
//
// mog annotation:
//
// target=github.com/shulutkov/yellow-pages/agent/structs.QueryOptions
// output=common.gen.go
// name=Structs
// ignore-fields=StaleIfError,AllowNotModifiedResponse,state,sizeCache,unknownFields
type QueryOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token is the ACL token ID. If not provided, the 'anonymous'
	// token is assumed for backwards compatibility.
	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	// If set, wait until query exceeds given index. Must be provided
	// with MaxQueryTime.
	MinQueryIndex uint64 `protobuf:"varint,2,opt,name=MinQueryIndex,proto3" json:"MinQueryIndex,omitempty"`
	// Provided with MinQueryIndex to wait for change.
	// mog: func-to=structs.DurationFromProto func-from=structs.DurationToProto
	MaxQueryTime *durationpb.Duration `protobuf:"bytes,3,opt,name=MaxQueryTime,proto3" json:"MaxQueryTime,omitempty"`
	// If set, any follower can service the request. Results
	// may be arbitrarily stale.
	AllowStale bool `protobuf:"varint,4,opt,name=AllowStale,proto3" json:"AllowStale,omitempty"`
	// If set, the leader must verify leadership prior to
	// servicing the request. Prevents a stale read.
	RequireConsistent bool `protobuf:"varint,5,opt,name=RequireConsistent,proto3" json:"RequireConsistent,omitempty"`
	// If set, the local agent may respond with an arbitrarily stale locally
	// cached response. The semantics differ from AllowStale since the agent may
	// be entirely partitioned from the servers and still considered "healthy" by
	// operators. Stale responses from Servers are also arbitrarily stale, but can
	// provide additional bounds on the last contact time from the leader. It's
	// expected that servers that are partitioned are noticed and replaced in a
	// timely way by operators while the same may not be true for client agents.
	UseCache bool `protobuf:"varint,6,opt,name=UseCache,proto3" json:"UseCache,omitempty"`
	// If set and AllowStale is true, will try first a stale
	// read, and then will perform a consistent read if stale
	// read is older than value.
	// mog: func-to=structs.DurationFromProto func-from=structs.DurationToProto
	MaxStaleDuration *durationpb.Duration `protobuf:"bytes,7,opt,name=MaxStaleDuration,proto3" json:"MaxStaleDuration,omitempty"`
	// MaxAge limits how old a cached value will be returned if UseCache is true.
	// If there is a cached response that is older than the MaxAge, it is treated
	// as a cache miss and a new fetch invoked. If the fetch fails, the error is
	// returned. Clients that wish to allow for stale results on error can set
	// StaleIfError to a longer duration to change this behavior. It is ignored
	// if the endpoint supports background refresh caching. See
	// https://www.consul.io/api/index.html#agent-caching for more details.
	// mog: func-to=structs.DurationFromProto func-from=structs.DurationToProto
	MaxAge *durationpb.Duration `protobuf:"bytes,8,opt,name=MaxAge,proto3" json:"MaxAge,omitempty"`
	// MustRevalidate forces the agent to fetch a fresh version of a cached
	// resource or at least validate that the cached version is still fresh. It is
	// implied by either max-age=0 or must-revalidate Cache-Control headers. It
	// only makes sense when UseCache is true. We store it since MaxAge = 0 is the
	// default unset value.
	MustRevalidate bool `protobuf:"varint,9,opt,name=MustRevalidate,proto3" json:"MustRevalidate,omitempty"`
	// StaleIfError specifies how stale the client will accept a cached response
	// if the servers are unavailable to fetch a fresh one. Only makes sense when
	// UseCache is true and MaxAge is set to a lower, non-zero value. It is
	// ignored if the endpoint supports background refresh caching. See
	// https://www.consul.io/api/index.html#agent-caching for more details.
	StaleIfError *durationpb.Duration `protobuf:"bytes,10,opt,name=StaleIfError,proto3" json:"StaleIfError,omitempty"`
	// Filter specifies the go-bexpr filter expression to be used for
	// filtering the data prior to returning a response
	Filter string `protobuf:"bytes,11,opt,name=Filter,proto3" json:"Filter,omitempty"`
}

func (x *QueryOptions) Reset() {
	*x = QueryOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbcommon_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOptions) ProtoMessage() {}

func (x *QueryOptions) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbcommon_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOptions.ProtoReflect.Descriptor instead.
func (*QueryOptions) Descriptor() ([]byte, []int) {
	return file_private_pbcommon_common_proto_rawDescGZIP(), []int{4}
}

func (x *QueryOptions) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *QueryOptions) GetMinQueryIndex() uint64 {
	if x != nil {
		return x.MinQueryIndex
	}
	return 0
}

func (x *QueryOptions) GetMaxQueryTime() *durationpb.Duration {
	if x != nil {
		return x.MaxQueryTime
	}
	return nil
}

func (x *QueryOptions) GetAllowStale() bool {
	if x != nil {
		return x.AllowStale
	}
	return false
}

func (x *QueryOptions) GetRequireConsistent() bool {
	if x != nil {
		return x.RequireConsistent
	}
	return false
}

func (x *QueryOptions) GetUseCache() bool {
	if x != nil {
		return x.UseCache
	}
	return false
}

func (x *QueryOptions) GetMaxStaleDuration() *durationpb.Duration {
	if x != nil {
		return x.MaxStaleDuration
	}
	return nil
}

func (x *QueryOptions) GetMaxAge() *durationpb.Duration {
	if x != nil {
		return x.MaxAge
	}
	return nil
}

func (x *QueryOptions) GetMustRevalidate() bool {
	if x != nil {
		return x.MustRevalidate
	}
	return false
}

func (x *QueryOptions) GetStaleIfError() *durationpb.Duration {
	if x != nil {
		return x.StaleIfError
	}
	return nil
}

func (x *QueryOptions) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// QueryMeta allows a query response to include potentially
// useful metadata about a query
//
// mog annotation:
//
// target=github.com/shulutkov/yellow-pages/agent/structs.QueryMeta
// output=common.gen.go
// name=Structs
// ignore-fields=NotModified,Backend,state,sizeCache,unknownFields
type QueryMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the index associated with the read
	Index uint64 `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	// If AllowStale is used, this is time elapsed since
	// last contact between the follower and leader. This
	// can be used to gauge staleness.
	// mog: func-to=structs.DurationFromProto func-from=structs.DurationToProto
	LastContact *durationpb.Duration `protobuf:"bytes,2,opt,name=LastContact,proto3" json:"LastContact,omitempty"`
	// Used to indicate if there is a known leader node
	KnownLeader bool `protobuf:"varint,3,opt,name=KnownLeader,proto3" json:"KnownLeader,omitempty"`
	// Consistencylevel returns the consistency used to serve the query
	// Having `discovery_max_stale` on the agent can affect whether
	// the request was served by a leader.
	ConsistencyLevel string `protobuf:"bytes,4,opt,name=ConsistencyLevel,proto3" json:"ConsistencyLevel,omitempty"`
	// ResultsFilteredByACLs is true when some of the query's results were
	// filtered out by enforcing ACLs. It may be false because nothing was
	// removed, or because the endpoint does not yet support this flag.
	ResultsFilteredByACLs bool `protobuf:"varint,7,opt,name=ResultsFilteredByACLs,proto3" json:"ResultsFilteredByACLs,omitempty"`
}

func (x *QueryMeta) Reset() {
	*x = QueryMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbcommon_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMeta) ProtoMessage() {}

func (x *QueryMeta) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbcommon_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMeta.ProtoReflect.Descriptor instead.
func (*QueryMeta) Descriptor() ([]byte, []int) {
	return file_private_pbcommon_common_proto_rawDescGZIP(), []int{5}
}

func (x *QueryMeta) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *QueryMeta) GetLastContact() *durationpb.Duration {
	if x != nil {
		return x.LastContact
	}
	return nil
}

func (x *QueryMeta) GetKnownLeader() bool {
	if x != nil {
		return x.KnownLeader
	}
	return false
}

func (x *QueryMeta) GetConsistencyLevel() string {
	if x != nil {
		return x.ConsistencyLevel
	}
	return ""
}

func (x *QueryMeta) GetResultsFilteredByACLs() bool {
	if x != nil {
		return x.ResultsFilteredByACLs
	}
	return false
}

// EnterpriseMeta contains metadata that is only used by the Enterprise version
// of Consul.
type EnterpriseMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Namespace in which the entity exists.
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	// Partition in which the entity exists.
	Partition string `protobuf:"bytes,2,opt,name=Partition,proto3" json:"Partition,omitempty"`
}

func (x *EnterpriseMeta) Reset() {
	*x = EnterpriseMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbcommon_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterpriseMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterpriseMeta) ProtoMessage() {}

func (x *EnterpriseMeta) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbcommon_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterpriseMeta.ProtoReflect.Descriptor instead.
func (*EnterpriseMeta) Descriptor() ([]byte, []int) {
	return file_private_pbcommon_common_proto_rawDescGZIP(), []int{6}
}

func (x *EnterpriseMeta) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *EnterpriseMeta) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

// mog annotation:
//
// target=github.com/shulutkov/yellow-pages/agent/structs.EnvoyExtension
// output=common.gen.go
// name=Structs
type EnvoyExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Required bool   `protobuf:"varint,2,opt,name=Required,proto3" json:"Required,omitempty"`
	// mog: func-to=ProtobufTypesStructToMapStringInterface func-from=MapStringInterfaceToProtobufTypesStruct
	Arguments     *structpb.Struct `protobuf:"bytes,3,opt,name=Arguments,proto3" json:"Arguments,omitempty"`
	ConsulVersion string           `protobuf:"bytes,4,opt,name=ConsulVersion,proto3" json:"ConsulVersion,omitempty"`
	EnvoyVersion  string           `protobuf:"bytes,5,opt,name=EnvoyVersion,proto3" json:"EnvoyVersion,omitempty"`
}

func (x *EnvoyExtension) Reset() {
	*x = EnvoyExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbcommon_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvoyExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvoyExtension) ProtoMessage() {}

func (x *EnvoyExtension) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbcommon_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvoyExtension.ProtoReflect.Descriptor instead.
func (*EnvoyExtension) Descriptor() ([]byte, []int) {
	return file_private_pbcommon_common_proto_rawDescGZIP(), []int{7}
}

func (x *EnvoyExtension) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnvoyExtension) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *EnvoyExtension) GetArguments() *structpb.Struct {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *EnvoyExtension) GetConsulVersion() string {
	if x != nil {
		return x.ConsulVersion
	}
	return ""
}

func (x *EnvoyExtension) GetEnvoyVersion() string {
	if x != nil {
		return x.EnvoyVersion
	}
	return ""
}

// mog annotation:
//
// target=github.com/shulutkov/yellow-pages/agent/structs.Locality
// output=common.gen.go
// name=Structs
type Locality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Region is region the zone belongs to.
	Region string `protobuf:"bytes,1,opt,name=Region,proto3" json:"Region,omitempty"`
	// Zone is the zone the entity is running in.
	Zone string `protobuf:"bytes,2,opt,name=Zone,proto3" json:"Zone,omitempty"`
}

func (x *Locality) Reset() {
	*x = Locality{}
	if protoimpl.UnsafeEnabled {
		mi := &file_private_pbcommon_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Locality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Locality) ProtoMessage() {}

func (x *Locality) ProtoReflect() protoreflect.Message {
	mi := &file_private_pbcommon_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Locality.ProtoReflect.Descriptor instead.
func (*Locality) Descriptor() ([]byte, []int) {
	return file_private_pbcommon_common_proto_rawDescGZIP(), []int{8}
}

func (x *Locality) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Locality) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

var File_private_pbcommon_common_proto protoreflect.FileDescriptor

var file_private_pbcommon_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4f, 0x0a, 0x09, 0x52, 0x61, 0x66, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20,
	0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x32, 0x0a, 0x10, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x51, 0x0a, 0x0b, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x2c, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xec, 0x03,
	0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x4d, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x4d, 0x69, 0x6e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x3d, 0x0a, 0x0c, 0x4d, 0x61,
	0x78, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x4d, 0x61, 0x78,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x6c, 0x6c,
	0x6f, 0x77, 0x53, 0x74, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x55, 0x73, 0x65, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x4d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x6c, 0x65, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x4d, 0x61, 0x78, 0x53, 0x74, 0x61,
	0x6c, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x4d, 0x61,
	0x78, 0x41, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x4d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x4d, 0x75, 0x73, 0x74, 0x52, 0x65, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x4d, 0x75, 0x73, 0x74, 0x52, 0x65, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x6c, 0x65, 0x49, 0x66,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x53, 0x74, 0x61, 0x6c, 0x65, 0x49, 0x66, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xee, 0x01, 0x0a,
	0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x3b, 0x0a, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x43, 0x6f, 0x6e, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x34, 0x0a, 0x15, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x42, 0x79,
	0x41, 0x43, 0x4c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x65, 0x64, 0x42, 0x79, 0x41, 0x43, 0x4c,
	0x73, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x22, 0x4c, 0x0a,
	0x0e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x72, 0x69, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc1, 0x01, 0x0a, 0x0e,
	0x45, 0x6e, 0x76, 0x6f, 0x79, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x35,
	0x0a, 0x09, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x09, 0x41, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x45,
	0x6e, 0x76, 0x6f, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x45, 0x6e, 0x76, 0x6f, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x36, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x42, 0x8b, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x04, 0x48, 0x43, 0x49, 0x43, 0xaa, 0x02, 0x20, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x20,
	0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0xe2, 0x02, 0x2c, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x23, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x3a, 0x3a, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_private_pbcommon_common_proto_rawDescOnce sync.Once
	file_private_pbcommon_common_proto_rawDescData = file_private_pbcommon_common_proto_rawDesc
)

func file_private_pbcommon_common_proto_rawDescGZIP() []byte {
	file_private_pbcommon_common_proto_rawDescOnce.Do(func() {
		file_private_pbcommon_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_private_pbcommon_common_proto_rawDescData)
	})
	return file_private_pbcommon_common_proto_rawDescData
}

var file_private_pbcommon_common_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_private_pbcommon_common_proto_goTypes = []interface{}{
	(*RaftIndex)(nil),           // 0: hashicorp.consul.internal.common.RaftIndex
	(*TargetDatacenter)(nil),    // 1: hashicorp.consul.internal.common.TargetDatacenter
	(*WriteRequest)(nil),        // 2: hashicorp.consul.internal.common.WriteRequest
	(*ReadRequest)(nil),         // 3: hashicorp.consul.internal.common.ReadRequest
	(*QueryOptions)(nil),        // 4: hashicorp.consul.internal.common.QueryOptions
	(*QueryMeta)(nil),           // 5: hashicorp.consul.internal.common.QueryMeta
	(*EnterpriseMeta)(nil),      // 6: hashicorp.consul.internal.common.EnterpriseMeta
	(*EnvoyExtension)(nil),      // 7: hashicorp.consul.internal.common.EnvoyExtension
	(*Locality)(nil),            // 8: hashicorp.consul.internal.common.Locality
	(*durationpb.Duration)(nil), // 9: google.protobuf.Duration
	(*structpb.Struct)(nil),     // 10: google.protobuf.Struct
}
var file_private_pbcommon_common_proto_depIdxs = []int32{
	9,  // 0: hashicorp.consul.internal.common.QueryOptions.MaxQueryTime:type_name -> google.protobuf.Duration
	9,  // 1: hashicorp.consul.internal.common.QueryOptions.MaxStaleDuration:type_name -> google.protobuf.Duration
	9,  // 2: hashicorp.consul.internal.common.QueryOptions.MaxAge:type_name -> google.protobuf.Duration
	9,  // 3: hashicorp.consul.internal.common.QueryOptions.StaleIfError:type_name -> google.protobuf.Duration
	9,  // 4: hashicorp.consul.internal.common.QueryMeta.LastContact:type_name -> google.protobuf.Duration
	10, // 5: hashicorp.consul.internal.common.EnvoyExtension.Arguments:type_name -> google.protobuf.Struct
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_private_pbcommon_common_proto_init() }
func file_private_pbcommon_common_proto_init() {
	if File_private_pbcommon_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_private_pbcommon_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RaftIndex); i {
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
		file_private_pbcommon_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetDatacenter); i {
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
		file_private_pbcommon_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_private_pbcommon_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_private_pbcommon_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOptions); i {
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
		file_private_pbcommon_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMeta); i {
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
		file_private_pbcommon_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterpriseMeta); i {
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
		file_private_pbcommon_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvoyExtension); i {
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
		file_private_pbcommon_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Locality); i {
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
			RawDescriptor: file_private_pbcommon_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_private_pbcommon_common_proto_goTypes,
		DependencyIndexes: file_private_pbcommon_common_proto_depIdxs,
		MessageInfos:      file_private_pbcommon_common_proto_msgTypes,
	}.Build()
	File_private_pbcommon_common_proto = out.File
	file_private_pbcommon_common_proto_rawDesc = nil
	file_private_pbcommon_common_proto_goTypes = nil
	file_private_pbcommon_common_proto_depIdxs = nil
}
