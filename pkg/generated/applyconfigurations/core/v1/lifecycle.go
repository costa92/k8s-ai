

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// LifecycleApplyConfiguration represents a declarative configuration of the Lifecycle type for use
// with apply.
type LifecycleApplyConfiguration struct {
	PostStart *LifecycleHandlerApplyConfiguration `json:"postStart,omitempty"`
	PreStop   *LifecycleHandlerApplyConfiguration `json:"preStop,omitempty"`
}

// LifecycleApplyConfiguration constructs a declarative configuration of the Lifecycle type for use with
// apply.
func Lifecycle() *LifecycleApplyConfiguration {
	return &LifecycleApplyConfiguration{}
}

// WithPostStart sets the PostStart field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PostStart field is set to the value of the last call.
func (b *LifecycleApplyConfiguration) WithPostStart(value *LifecycleHandlerApplyConfiguration) *LifecycleApplyConfiguration {
	b.PostStart = value
	return b
}

// WithPreStop sets the PreStop field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PreStop field is set to the value of the last call.
func (b *LifecycleApplyConfiguration) WithPreStop(value *LifecycleHandlerApplyConfiguration) *LifecycleApplyConfiguration {
	b.PreStop = value
	return b
}
