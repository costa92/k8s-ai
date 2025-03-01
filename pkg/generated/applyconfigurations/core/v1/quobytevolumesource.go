

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// QuobyteVolumeSourceApplyConfiguration represents a declarative configuration of the QuobyteVolumeSource type for use
// with apply.
type QuobyteVolumeSourceApplyConfiguration struct {
	Registry *string `json:"registry,omitempty"`
	Volume   *string `json:"volume,omitempty"`
	ReadOnly *bool   `json:"readOnly,omitempty"`
	User     *string `json:"user,omitempty"`
	Group    *string `json:"group,omitempty"`
	Tenant   *string `json:"tenant,omitempty"`
}

// QuobyteVolumeSourceApplyConfiguration constructs a declarative configuration of the QuobyteVolumeSource type for use with
// apply.
func QuobyteVolumeSource() *QuobyteVolumeSourceApplyConfiguration {
	return &QuobyteVolumeSourceApplyConfiguration{}
}

// WithRegistry sets the Registry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Registry field is set to the value of the last call.
func (b *QuobyteVolumeSourceApplyConfiguration) WithRegistry(value string) *QuobyteVolumeSourceApplyConfiguration {
	b.Registry = &value
	return b
}

// WithVolume sets the Volume field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Volume field is set to the value of the last call.
func (b *QuobyteVolumeSourceApplyConfiguration) WithVolume(value string) *QuobyteVolumeSourceApplyConfiguration {
	b.Volume = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *QuobyteVolumeSourceApplyConfiguration) WithReadOnly(value bool) *QuobyteVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithUser sets the User field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the User field is set to the value of the last call.
func (b *QuobyteVolumeSourceApplyConfiguration) WithUser(value string) *QuobyteVolumeSourceApplyConfiguration {
	b.User = &value
	return b
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *QuobyteVolumeSourceApplyConfiguration) WithGroup(value string) *QuobyteVolumeSourceApplyConfiguration {
	b.Group = &value
	return b
}

// WithTenant sets the Tenant field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Tenant field is set to the value of the last call.
func (b *QuobyteVolumeSourceApplyConfiguration) WithTenant(value string) *QuobyteVolumeSourceApplyConfiguration {
	b.Tenant = &value
	return b
}
