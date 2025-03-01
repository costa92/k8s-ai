

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PodAntiAffinityApplyConfiguration represents a declarative configuration of the PodAntiAffinity type for use
// with apply.
type PodAntiAffinityApplyConfiguration struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []PodAffinityTermApplyConfiguration         `json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
	PreferredDuringSchedulingIgnoredDuringExecution []WeightedPodAffinityTermApplyConfiguration `json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
}

// PodAntiAffinityApplyConfiguration constructs a declarative configuration of the PodAntiAffinity type for use with
// apply.
func PodAntiAffinity() *PodAntiAffinityApplyConfiguration {
	return &PodAntiAffinityApplyConfiguration{}
}

// WithRequiredDuringSchedulingIgnoredDuringExecution adds the given value to the RequiredDuringSchedulingIgnoredDuringExecution field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RequiredDuringSchedulingIgnoredDuringExecution field.
func (b *PodAntiAffinityApplyConfiguration) WithRequiredDuringSchedulingIgnoredDuringExecution(values ...*PodAffinityTermApplyConfiguration) *PodAntiAffinityApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRequiredDuringSchedulingIgnoredDuringExecution")
		}
		b.RequiredDuringSchedulingIgnoredDuringExecution = append(b.RequiredDuringSchedulingIgnoredDuringExecution, *values[i])
	}
	return b
}

// WithPreferredDuringSchedulingIgnoredDuringExecution adds the given value to the PreferredDuringSchedulingIgnoredDuringExecution field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PreferredDuringSchedulingIgnoredDuringExecution field.
func (b *PodAntiAffinityApplyConfiguration) WithPreferredDuringSchedulingIgnoredDuringExecution(values ...*WeightedPodAffinityTermApplyConfiguration) *PodAntiAffinityApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPreferredDuringSchedulingIgnoredDuringExecution")
		}
		b.PreferredDuringSchedulingIgnoredDuringExecution = append(b.PreferredDuringSchedulingIgnoredDuringExecution, *values[i])
	}
	return b
}
