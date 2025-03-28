// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package pbcommon

import "github.com/arenadata/consul/acl"

var DefaultEnterpriseMeta = &EnterpriseMeta{}

func NewEnterpriseMetaFromStructs(_ acl.EnterpriseMeta) *EnterpriseMeta {
	return &EnterpriseMeta{}
}
func EnterpriseMetaToStructs(s *EnterpriseMeta, t *acl.EnterpriseMeta) {
	if s == nil {
		return
	}
}
func EnterpriseMetaFromStructs(t *acl.EnterpriseMeta, s *EnterpriseMeta) {
	if s == nil {
		return
	}
}
