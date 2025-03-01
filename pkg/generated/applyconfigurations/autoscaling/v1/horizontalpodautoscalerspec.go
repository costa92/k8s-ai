

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// HorizontalPodAutoscalerSpecApplyConfiguration represents a declarative configuration of the HorizontalPodAutoscalerSpec type for use
// with apply.
type HorizontalPodAutoscalerSpecApplyConfiguration struct {
	ScaleTargetRef                 *CrossVersionObjectReferenceApplyConfiguration `json:"scaleTargetRef,omitempty"`
	MinReplicas                    *int32                                         `json:"minReplicas,omitempty"`
	MaxReplicas                    *int32                                         `json:"maxReplicas,omitempty"`
	TargetCPUUtilizationPercentage *int32                                         `json:"targetCPUUtilizationPercentage,omitempty"`
}

// HorizontalPodAutoscalerSpecApplyConfiguration constructs a declarative configuration of the HorizontalPodAutoscalerSpec type for use with
// apply.
func HorizontalPodAutoscalerSpec() *HorizontalPodAutoscalerSpecApplyConfiguration {
	return &HorizontalPodAutoscalerSpecApplyConfiguration{}
}

// WithScaleTargetRef sets the ScaleTargetRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScaleTargetRef field is set to the value of the last call.
func (b *HorizontalPodAutoscalerSpecApplyConfiguration) WithScaleTargetRef(value *CrossVersionObjectReferenceApplyConfiguration) *HorizontalPodAutoscalerSpecApplyConfiguration {
	b.ScaleTargetRef = value
	return b
}

// WithMinReplicas sets the MinReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinReplicas field is set to the value of the last call.
func (b *HorizontalPodAutoscalerSpecApplyConfiguration) WithMinReplicas(value int32) *HorizontalPodAutoscalerSpecApplyConfiguration {
	b.MinReplicas = &value
	return b
}

// WithMaxReplicas sets the MaxReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxReplicas field is set to the value of the last call.
func (b *HorizontalPodAutoscalerSpecApplyConfiguration) WithMaxReplicas(value int32) *HorizontalPodAutoscalerSpecApplyConfiguration {
	b.MaxReplicas = &value
	return b
}

// WithTargetCPUUtilizationPercentage sets the TargetCPUUtilizationPercentage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetCPUUtilizationPercentage field is set to the value of the last call.
func (b *HorizontalPodAutoscalerSpecApplyConfiguration) WithTargetCPUUtilizationPercentage(value int32) *HorizontalPodAutoscalerSpecApplyConfiguration {
	b.TargetCPUUtilizationPercentage = &value
	return b
}
