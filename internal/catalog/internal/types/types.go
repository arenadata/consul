// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"github.com/arenadata/consul/internal/resource"
)

const (
	GroupName       = "catalog"
	VersionV1Alpha1 = "v1alpha1"
	CurrentVersion  = VersionV1Alpha1
)

func Register(r resource.Registry) {
	RegisterWorkload(r)
	RegisterService(r)
	RegisterServiceEndpoints(r)
	RegisterNode(r)
	RegisterHealthStatus(r)
	RegisterHealthChecks(r)
	RegisterDNSPolicy(r)
	RegisterVirtualIPs(r)
}
