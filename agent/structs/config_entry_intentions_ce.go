// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !consulent
// +build !consulent

package structs

import (
	"fmt"

	"github.com/arenadata/consul/acl"
)

func validateSourceIntentionEnterpriseMeta(_, _ *acl.EnterpriseMeta) error {
	return nil
}

func (s *SourceIntention) validateSamenessGroup() error {
	if s.SamenessGroup != "" {
		return fmt.Errorf("Sameness groups are a Consul Enterprise feature.")
	}

	return nil
}
