// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package agent

import (
	"github.com/arenadata/consul/acl"
	"github.com/arenadata/consul/agent/structs"
	"github.com/arenadata/consul/types"
)

// persistedCheck is used to serialize a check and write it to disk
// so that it may be restored later on.
type persistedCheck struct {
	Check   *structs.HealthCheck
	ChkType *structs.CheckType
	Token   string
	Source  string
}

// persistedCheckState is used to persist the current state of a given
// check. This is different from the check definition, and includes an
// expiration timestamp which is used to determine staleness on later
// agent restarts.
type persistedCheckState struct {
	CheckID types.CheckID
	Output  string
	Status  string
	Expires int64
	acl.EnterpriseMeta
}
