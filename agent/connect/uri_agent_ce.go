// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package connect

import (
	"fmt"

	"github.com/arenadata/consul/acl"
)

// GetEnterpriseMeta will synthesize an EnterpriseMeta struct from the SpiffeIDAgent.
// in CE this just returns an empty (but never nil) struct pointer
func (id SpiffeIDAgent) GetEnterpriseMeta() *acl.EnterpriseMeta {
	return &acl.EnterpriseMeta{}
}

func (id SpiffeIDAgent) uriPath() string {
	return fmt.Sprintf("/agent/client/dc/%s/id/%s", id.Datacenter, id.Agent)
}
