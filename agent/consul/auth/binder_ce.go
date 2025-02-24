// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package auth

import (
	"github.com/shulutkov/yellow-pages/acl"
	"github.com/shulutkov/yellow-pages/agent/consul/authmethod"
	"github.com/shulutkov/yellow-pages/agent/structs"
)

func bindEnterpriseMeta(authMethod *structs.ACLAuthMethod, verifiedIdentity *authmethod.Identity) (acl.EnterpriseMeta, error) {
	return acl.EnterpriseMeta{}, nil
}
