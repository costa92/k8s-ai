

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// ResourceQuotaStatusApplyConfiguration represents a declarative configuration of the ResourceQuotaStatus type for use
// with apply.
type ResourceQuotaStatusApplyConfiguration struct {
	Hard *v1.ResourceList `json:"hard,omitempty"`
	Used *v1.ResourceList `json:"used,omitempty"`
}

// ResourceQuotaStatusApplyConfiguration constructs a declarative configuration of the ResourceQuotaStatus type for use with
// apply.
func ResourceQuotaStatus() *ResourceQuotaStatusApplyConfiguration {
	return &ResourceQuotaStatusApplyConfiguration{}
}

// WithHard sets the Hard field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hard field is set to the value of the last call.
func (b *ResourceQuotaStatusApplyConfiguration) WithHard(value v1.ResourceList) *ResourceQuotaStatusApplyConfiguration {
	b.Hard = &value
	return b
}

// WithUsed sets the Used field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Used field is set to the value of the last call.
func (b *ResourceQuotaStatusApplyConfiguration) WithUsed(value v1.ResourceList) *ResourceQuotaStatusApplyConfiguration {
	b.Used = &value
	return b
}
