

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ConfigMapEnvSourceApplyConfiguration represents a declarative configuration of the ConfigMapEnvSource type for use
// with apply.
type ConfigMapEnvSourceApplyConfiguration struct {
	LocalObjectReferenceApplyConfiguration `json:",inline"`
	Optional                               *bool `json:"optional,omitempty"`
}

// ConfigMapEnvSourceApplyConfiguration constructs a declarative configuration of the ConfigMapEnvSource type for use with
// apply.
func ConfigMapEnvSource() *ConfigMapEnvSourceApplyConfiguration {
	return &ConfigMapEnvSourceApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ConfigMapEnvSourceApplyConfiguration) WithName(value string) *ConfigMapEnvSourceApplyConfiguration {
	b.Name = &value
	return b
}

// WithOptional sets the Optional field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Optional field is set to the value of the last call.
func (b *ConfigMapEnvSourceApplyConfiguration) WithOptional(value bool) *ConfigMapEnvSourceApplyConfiguration {
	b.Optional = &value
	return b
}
