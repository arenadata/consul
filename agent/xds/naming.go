// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package xds

import (
	"fmt"

	"github.com/arenadata/consul/agent/structs"
)

func CustomizeClusterName(clusterName string, chain *structs.CompiledDiscoveryChain) string {
	if chain == nil || chain.CustomizationHash == "" {
		return clusterName
	}
	// Use a colon to delimit this prefix instead of a dot to avoid a
	// theoretical collision problem with subsets.
	return fmt.Sprintf("%s~%s", chain.CustomizationHash, clusterName)
}
