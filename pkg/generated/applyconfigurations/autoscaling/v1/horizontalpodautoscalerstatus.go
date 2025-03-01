

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HorizontalPodAutoscalerStatusApplyConfiguration represents a declarative configuration of the HorizontalPodAutoscalerStatus type for use
// with apply.
type HorizontalPodAutoscalerStatusApplyConfiguration struct {
	ObservedGeneration              *int64   `json:"observedGeneration,omitempty"`
	LastScaleTime                   *v1.Time `json:"lastScaleTime,omitempty"`
	CurrentReplicas                 *int32   `json:"currentReplicas,omitempty"`
	DesiredReplicas                 *int32   `json:"desiredReplicas,omitempty"`
	CurrentCPUUtilizationPercentage *int32   `json:"currentCPUUtilizationPercentage,omitempty"`
}

// HorizontalPodAutoscalerStatusApplyConfiguration constructs a declarative configuration of the HorizontalPodAutoscalerStatus type for use with
// apply.
func HorizontalPodAutoscalerStatus() *HorizontalPodAutoscalerStatusApplyConfiguration {
	return &HorizontalPodAutoscalerStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithObservedGeneration(value int64) *HorizontalPodAutoscalerStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithLastScaleTime sets the LastScaleTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastScaleTime field is set to the value of the last call.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithLastScaleTime(value v1.Time) *HorizontalPodAutoscalerStatusApplyConfiguration {
	b.LastScaleTime = &value
	return b
}

// WithCurrentReplicas sets the CurrentReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentReplicas field is set to the value of the last call.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithCurrentReplicas(value int32) *HorizontalPodAutoscalerStatusApplyConfiguration {
	b.CurrentReplicas = &value
	return b
}

// WithDesiredReplicas sets the DesiredReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredReplicas field is set to the value of the last call.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithDesiredReplicas(value int32) *HorizontalPodAutoscalerStatusApplyConfiguration {
	b.DesiredReplicas = &value
	return b
}

// WithCurrentCPUUtilizationPercentage sets the CurrentCPUUtilizationPercentage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentCPUUtilizationPercentage field is set to the value of the last call.
func (b *HorizontalPodAutoscalerStatusApplyConfiguration) WithCurrentCPUUtilizationPercentage(value int32) *HorizontalPodAutoscalerStatusApplyConfiguration {
	b.CurrentCPUUtilizationPercentage = &value
	return b
}
