// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"github.com/arenadata/consul/internal/resource"
	pbmesh "github.com/arenadata/consul/proto-public/pbmesh/v1alpha1"
	"github.com/arenadata/consul/proto-public/pbresource"
)

const (
	ProxyConfigurationKind = "ProxyConfiguration"
)

var (
	ProxyConfigurationV1Alpha1Type = &pbresource.Type{
		Group:        GroupName,
		GroupVersion: CurrentVersion,
		Kind:         ProxyConfigurationKind,
	}

	ProxyConfigurationType = ProxyConfigurationV1Alpha1Type
)

func RegisterProxyConfiguration(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     ProxyConfigurationV1Alpha1Type,
		Proto:    &pbmesh.ProxyConfiguration{},
		Validate: nil,
	})
}
