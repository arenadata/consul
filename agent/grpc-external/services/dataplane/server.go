// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplane

import (
	"google.golang.org/grpc"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-memdb"

	"github.com/arenadata/consul/acl"
	"github.com/arenadata/consul/acl/resolver"
	"github.com/arenadata/consul/agent/configentry"
	"github.com/arenadata/consul/agent/structs"
	"github.com/arenadata/consul/proto-public/pbdataplane"
)

type Server struct {
	Config
}

type Config struct {
	GetStore    func() StateStore
	Logger      hclog.Logger
	ACLResolver ACLResolver
	// Datacenter of the Consul server this gRPC server is hosted on
	Datacenter string
}

type StateStore interface {
	ServiceNode(string, string, string, *acl.EnterpriseMeta, string) (uint64, *structs.ServiceNode, error)
	ReadResolvedServiceConfigEntries(memdb.WatchSet, string, *acl.EnterpriseMeta, []structs.ServiceID, structs.ProxyMode) (uint64, *configentry.ResolvedServiceConfigSet, error)
}

//go:generate mockery --name ACLResolver --inpackage
type ACLResolver interface {
	ResolveTokenAndDefaultMeta(string, *acl.EnterpriseMeta, *acl.AuthorizerContext) (resolver.Result, error)
}

func NewServer(cfg Config) *Server {
	return &Server{cfg}
}

var _ pbdataplane.DataplaneServiceServer = (*Server)(nil)

func (s *Server) Register(grpcServer *grpc.Server) {
	pbdataplane.RegisterDataplaneServiceServer(grpcServer, s)
}
