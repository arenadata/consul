// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package ssoauth

import (
	"fmt"

	"github.com/arenadata/consul/acl"
	"github.com/arenadata/consul/internal/go-sso/oidcauth"
)

func validateType(typ string) error {
	if typ != "jwt" {
		return fmt.Errorf("type should be %q", "jwt")
	}
	return nil
}

func (v *Validator) ssoEntMetaFromClaims(_ *oidcauth.Claims) *acl.EnterpriseMeta {
	return nil
}

type enterpriseConfig struct{}

func (c *Config) enterpriseConvertForLibrary(_ *oidcauth.Config) {}
