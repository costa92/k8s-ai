

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
)

// ResourceQuotaSpecApplyConfiguration represents a declarative configuration of the ResourceQuotaSpec type for use
// with apply.
type ResourceQuotaSpecApplyConfiguration struct {
	Hard          *v1.ResourceList                 `json:"hard,omitempty"`
	Scopes        []v1.ResourceQuotaScope          `json:"scopes,omitempty"`
	ScopeSelector *ScopeSelectorApplyConfiguration `json:"scopeSelector,omitempty"`
}

// ResourceQuotaSpecApplyConfiguration constructs a declarative configuration of the ResourceQuotaSpec type for use with
// apply.
func ResourceQuotaSpec() *ResourceQuotaSpecApplyConfiguration {
	return &ResourceQuotaSpecApplyConfiguration{}
}

// WithHard sets the Hard field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Hard field is set to the value of the last call.
func (b *ResourceQuotaSpecApplyConfiguration) WithHard(value v1.ResourceList) *ResourceQuotaSpecApplyConfiguration {
	b.Hard = &value
	return b
}

// WithScopes adds the given value to the Scopes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Scopes field.
func (b *ResourceQuotaSpecApplyConfiguration) WithScopes(values ...v1.ResourceQuotaScope) *ResourceQuotaSpecApplyConfiguration {
	for i := range values {
		b.Scopes = append(b.Scopes, values[i])
	}
	return b
}

// WithScopeSelector sets the ScopeSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScopeSelector field is set to the value of the last call.
func (b *ResourceQuotaSpecApplyConfiguration) WithScopeSelector(value *ScopeSelectorApplyConfiguration) *ResourceQuotaSpecApplyConfiguration {
	b.ScopeSelector = value
	return b
}
