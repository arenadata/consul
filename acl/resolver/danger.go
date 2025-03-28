// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resolver

import "github.com/arenadata/consul/acl"

// DANGER_NO_AUTH implements an ACL resolver short-circuit authorization in
// cases where it is handled somewhere else or expressly not required.
type DANGER_NO_AUTH struct{}

// ResolveTokenAndDefaultMeta returns an authorizer with unfettered permissions.
func (DANGER_NO_AUTH) ResolveTokenAndDefaultMeta(string, *acl.EnterpriseMeta, *acl.AuthorizerContext) (Result, error) {
	return Result{Authorizer: acl.ManageAll()}, nil
}
