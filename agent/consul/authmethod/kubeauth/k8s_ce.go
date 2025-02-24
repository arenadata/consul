// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package kubeauth

import (
	"github.com/shulutkov/yellow-pages/acl"
	"github.com/shulutkov/yellow-pages/agent/structs"
)

type enterpriseConfig struct{}

func enterpriseValidation(method *structs.ACLAuthMethod, config *Config) error {
	return nil
}

func (v *Validator) k8sEntMetaFromFields(fields map[string]string) *acl.EnterpriseMeta {
	return nil
}
