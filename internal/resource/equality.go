// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import "github.com/arenadata/consul/proto-public/pbresource"

// EqualType compares two resource types for equality without reflection.
func EqualType(a, b *pbresource.Type) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Group == b.Group &&
		a.GroupVersion == b.GroupVersion &&
		a.Kind == b.Kind
}

// EqualType compares two resource tenancies for equality without reflection.
func EqualTenancy(a, b *pbresource.Tenancy) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Partition == b.Partition &&
		a.PeerName == b.PeerName &&
		a.Namespace == b.Namespace
}

// EqualType compares two resource IDs for equality without reflection.
func EqualID(a, b *pbresource.ID) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return EqualType(a.Type, b.Type) &&
		EqualTenancy(a.Tenancy, b.Tenancy) &&
		a.Name == b.Name &&
		a.Uid == b.Uid
}

// EqualStatus compares two statuses for equality without reflection.
//
// Pass true for compareUpdatedAt to compare the UpdatedAt timestamps, which you
// generally *don't* want when dirty checking the status in a controller.
func EqualStatus(a, b *pbresource.Status, compareUpdatedAt bool) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.ObservedGeneration != b.ObservedGeneration {
		return false
	}

	if compareUpdatedAt && !a.UpdatedAt.AsTime().Equal(b.UpdatedAt.AsTime()) {
		return false
	}

	if len(a.Conditions) != len(b.Conditions) {
		return false
	}

	for i, ac := range a.Conditions {
		bc := b.Conditions[i]

		if !EqualCondition(ac, bc) {
			return false
		}
	}

	return true
}

// EqualCondition compares two conditions for equality without reflection.
func EqualCondition(a, b *pbresource.Condition) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Type == b.Type &&
		a.State == b.State &&
		a.Reason == b.Reason &&
		a.Message == b.Message &&
		EqualReference(a.Resource, b.Resource)
}

// EqualReference compares two references for equality without reflection.
func EqualReference(a, b *pbresource.Reference) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return EqualType(a.Type, b.Type) &&
		EqualTenancy(a.Tenancy, b.Tenancy) &&
		a.Name == b.Name &&
		a.Section == b.Section
}

// EqualStatusMap compares two status maps for equality without reflection.
func EqualStatusMap(a, b map[string]*pbresource.Status) bool {
	if len(a) != len(b) {
		return false
	}

	compared := make(map[string]struct{})
	for k, av := range a {
		bv, ok := b[k]
		if !ok {
			return false
		}
		if !EqualStatus(av, bv, true) {
			return false
		}
		compared[k] = struct{}{}
	}

	for k, bv := range b {
		if _, skip := compared[k]; skip {
			continue
		}

		av, ok := a[k]
		if !ok {
			return false
		}

		if !EqualStatus(av, bv, true) {
			return false
		}
	}

	return true
}
