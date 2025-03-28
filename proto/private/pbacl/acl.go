// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pbacl

import (
	"github.com/arenadata/consul/api"
)

func (a *ACLLink) ToAPI() api.ACLLink {
	return api.ACLLink{
		ID:   a.ID,
		Name: a.Name,
	}
}

func ACLLinkFromAPI(a api.ACLLink) *ACLLink {
	return &ACLLink{
		ID:   a.ID,
		Name: a.Name,
	}
}
