// Copyright 2020-2021 Clastix Labs
// SPDX-License-Identifier: Apache-2.0

package v1beta1

const (
	DenyWildcard = "capsule.clastix.io/deny-wildcard"
)

func (t *Tenant) IsWildcardDenied() bool {
	if v, ok := t.Annotations[DenyWildcard]; ok && v == "true" {
		return true
	}

	return false
}
