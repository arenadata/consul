// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serverdiscovery

import (
	"google.golang.org/grpc"

	"github.com/hashicorp/go-hclog"

	"github.com/arenadata/consul/acl"
	"github.com/arenadata/consul/acl/resolver"
	"github.com/arenadata/consul/agent/consul/stream"
	"github.com/arenadata/consul/proto-public/pbserverdiscovery"
)

type Server struct {
	Config
}

type Config struct {
	Publisher   EventPublisher
	Logger      hclog.Logger
	ACLResolver ACLResolver
}

type EventPublisher interface {
	Subscribe(*stream.SubscribeRequest) (*stream.Subscription, error)
}

//go:generate mockery --name ACLResolver --inpackage
type ACLResolver interface {
	ResolveTokenAndDefaultMeta(string, *acl.EnterpriseMeta, *acl.AuthorizerContext) (resolver.Result, error)
}

func NewServer(cfg Config) *Server {
	return &Server{cfg}
}

func (s *Server) Register(grpcServer *grpc.Server) {
	pbserverdiscovery.RegisterServerDiscoveryServiceServer(grpcServer, s)
}
