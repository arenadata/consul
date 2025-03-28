// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package agent

import (
	"net/http"

	"github.com/arenadata/consul/acl"
)

func (s *HTTPHandlers) validateRequestPartition(_ http.ResponseWriter, _ *acl.EnterpriseMeta) bool {
	return true
}

func (s *HTTPHandlers) defaultMetaPartitionToAgent(entMeta *acl.EnterpriseMeta) {
}
