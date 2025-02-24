// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package config

import (
	"github.com/shulutkov/yellow-pages/acl"
	"github.com/shulutkov/yellow-pages/agent/structs"
)

// EnterpriseMeta stub
type EnterpriseMeta struct{}

func (_ *EnterpriseMeta) ToStructs() acl.EnterpriseMeta {
	return *structs.DefaultEnterpriseMetaInDefaultPartition()
}
