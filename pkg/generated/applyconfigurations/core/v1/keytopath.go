

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// KeyToPathApplyConfiguration represents a declarative configuration of the KeyToPath type for use
// with apply.
type KeyToPathApplyConfiguration struct {
	Key  *string `json:"key,omitempty"`
	Path *string `json:"path,omitempty"`
	Mode *int32  `json:"mode,omitempty"`
}

// KeyToPathApplyConfiguration constructs a declarative configuration of the KeyToPath type for use with
// apply.
func KeyToPath() *KeyToPathApplyConfiguration {
	return &KeyToPathApplyConfiguration{}
}

// WithKey sets the Key field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Key field is set to the value of the last call.
func (b *KeyToPathApplyConfiguration) WithKey(value string) *KeyToPathApplyConfiguration {
	b.Key = &value
	return b
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *KeyToPathApplyConfiguration) WithPath(value string) *KeyToPathApplyConfiguration {
	b.Path = &value
	return b
}

// WithMode sets the Mode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Mode field is set to the value of the last call.
func (b *KeyToPathApplyConfiguration) WithMode(value int32) *KeyToPathApplyConfiguration {
	b.Mode = &value
	return b
}
