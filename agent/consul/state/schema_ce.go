// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package state

import (
	"fmt"

	"github.com/arenadata/consul/acl"
	"github.com/arenadata/consul/agent/structs"
)

func partitionedIndexEntryName(entry string, _ string) string {
	return entry
}

func partitionedAndNamespacedIndexEntryName(entry string, _ *acl.EnterpriseMeta) string {
	return entry
}

// peeredIndexEntryName returns the peered index key for an importable entity (e.g. checks, services, or nodes).
func peeredIndexEntryName(entry, peerName string) string {
	if peerName == "" {
		peerName = structs.LocalPeerKeyword
	}
	return fmt.Sprintf("peer.%s:%s", peerName, entry)
}
