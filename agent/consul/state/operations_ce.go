// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package state

import (
	"github.com/hashicorp/go-memdb"

	"github.com/arenadata/consul/acl"
)

func getCompoundWithTxn(tx ReadTxn, table, index string,
	_ *acl.EnterpriseMeta, idxVals ...interface{}) (memdb.ResultIterator, error) {

	return tx.Get(table, index, idxVals...)
}
