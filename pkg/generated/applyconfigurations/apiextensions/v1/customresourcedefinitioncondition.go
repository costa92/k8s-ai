

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CustomResourceDefinitionConditionApplyConfiguration represents a declarative configuration of the CustomResourceDefinitionCondition type for use
// with apply.
type CustomResourceDefinitionConditionApplyConfiguration struct {
	Type               *v1.CustomResourceDefinitionConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus                       `json:"status,omitempty"`
	LastTransitionTime *metav1.Time                              `json:"lastTransitionTime,omitempty"`
	Reason             *string                                   `json:"reason,omitempty"`
	Message            *string                                   `json:"message,omitempty"`
}

// CustomResourceDefinitionConditionApplyConfiguration constructs a declarative configuration of the CustomResourceDefinitionCondition type for use with
// apply.
func CustomResourceDefinitionCondition() *CustomResourceDefinitionConditionApplyConfiguration {
	return &CustomResourceDefinitionConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *CustomResourceDefinitionConditionApplyConfiguration) WithType(value v1.CustomResourceDefinitionConditionType) *CustomResourceDefinitionConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *CustomResourceDefinitionConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *CustomResourceDefinitionConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *CustomResourceDefinitionConditionApplyConfiguration) WithLastTransitionTime(value metav1.Time) *CustomResourceDefinitionConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *CustomResourceDefinitionConditionApplyConfiguration) WithReason(value string) *CustomResourceDefinitionConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *CustomResourceDefinitionConditionApplyConfiguration) WithMessage(value string) *CustomResourceDefinitionConditionApplyConfiguration {
	b.Message = &value
	return b
}
