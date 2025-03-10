

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/costa92/k8s-ai/pkg/apis/apps/v1beta1"
)

// ChargeRequestStatusApplyConfiguration represents a declarative configuration of the ChargeRequestStatus type for use
// with apply.
type ChargeRequestStatusApplyConfiguration struct {
	Conditions *v1beta1.Conditions `json:"conditions,omitempty"`
}

// ChargeRequestStatusApplyConfiguration constructs a declarative configuration of the ChargeRequestStatus type for use with
// apply.
func ChargeRequestStatus() *ChargeRequestStatusApplyConfiguration {
	return &ChargeRequestStatusApplyConfiguration{}
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *ChargeRequestStatusApplyConfiguration) WithConditions(value v1beta1.Conditions) *ChargeRequestStatusApplyConfiguration {
	b.Conditions = &value
	return b
}
