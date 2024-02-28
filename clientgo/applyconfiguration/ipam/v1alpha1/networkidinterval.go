// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ironcore-dev/ipam/api/ipam/v1alpha1"
)

// NetworkIDIntervalApplyConfiguration represents an declarative configuration of the NetworkIDInterval type for use
// with apply.
type NetworkIDIntervalApplyConfiguration struct {
	Begin *v1alpha1.NetworkID `json:"begin,omitempty"`
	Exact *v1alpha1.NetworkID `json:"exact,omitempty"`
	End   *v1alpha1.NetworkID `json:"end,omitempty"`
}

// NetworkIDIntervalApplyConfiguration constructs an declarative configuration of the NetworkIDInterval type for use with
// apply.
func NetworkIDInterval() *NetworkIDIntervalApplyConfiguration {
	return &NetworkIDIntervalApplyConfiguration{}
}

// WithBegin sets the Begin field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Begin field is set to the value of the last call.
func (b *NetworkIDIntervalApplyConfiguration) WithBegin(value v1alpha1.NetworkID) *NetworkIDIntervalApplyConfiguration {
	b.Begin = &value
	return b
}

// WithExact sets the Exact field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Exact field is set to the value of the last call.
func (b *NetworkIDIntervalApplyConfiguration) WithExact(value v1alpha1.NetworkID) *NetworkIDIntervalApplyConfiguration {
	b.Exact = &value
	return b
}

// WithEnd sets the End field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the End field is set to the value of the last call.
func (b *NetworkIDIntervalApplyConfiguration) WithEnd(value v1alpha1.NetworkID) *NetworkIDIntervalApplyConfiguration {
	b.End = &value
	return b
}