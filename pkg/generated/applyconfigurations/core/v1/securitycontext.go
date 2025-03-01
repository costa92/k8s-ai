

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// SecurityContextApplyConfiguration represents a declarative configuration of the SecurityContext type for use
// with apply.
type SecurityContextApplyConfiguration struct {
	Capabilities             *CapabilitiesApplyConfiguration                  `json:"capabilities,omitempty"`
	Privileged               *bool                                            `json:"privileged,omitempty"`
	SELinuxOptions           *SELinuxOptionsApplyConfiguration                `json:"seLinuxOptions,omitempty"`
	WindowsOptions           *WindowsSecurityContextOptionsApplyConfiguration `json:"windowsOptions,omitempty"`
	RunAsUser                *int64                                           `json:"runAsUser,omitempty"`
	RunAsGroup               *int64                                           `json:"runAsGroup,omitempty"`
	RunAsNonRoot             *bool                                            `json:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem   *bool                                            `json:"readOnlyRootFilesystem,omitempty"`
	AllowPrivilegeEscalation *bool                                            `json:"allowPrivilegeEscalation,omitempty"`
	ProcMount                *corev1.ProcMountType                            `json:"procMount,omitempty"`
	SeccompProfile           *SeccompProfileApplyConfiguration                `json:"seccompProfile,omitempty"`
	AppArmorProfile          *AppArmorProfileApplyConfiguration               `json:"appArmorProfile,omitempty"`
}

// SecurityContextApplyConfiguration constructs a declarative configuration of the SecurityContext type for use with
// apply.
func SecurityContext() *SecurityContextApplyConfiguration {
	return &SecurityContextApplyConfiguration{}
}

// WithCapabilities sets the Capabilities field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Capabilities field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithCapabilities(value *CapabilitiesApplyConfiguration) *SecurityContextApplyConfiguration {
	b.Capabilities = value
	return b
}

// WithPrivileged sets the Privileged field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Privileged field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithPrivileged(value bool) *SecurityContextApplyConfiguration {
	b.Privileged = &value
	return b
}

// WithSELinuxOptions sets the SELinuxOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SELinuxOptions field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithSELinuxOptions(value *SELinuxOptionsApplyConfiguration) *SecurityContextApplyConfiguration {
	b.SELinuxOptions = value
	return b
}

// WithWindowsOptions sets the WindowsOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WindowsOptions field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithWindowsOptions(value *WindowsSecurityContextOptionsApplyConfiguration) *SecurityContextApplyConfiguration {
	b.WindowsOptions = value
	return b
}

// WithRunAsUser sets the RunAsUser field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunAsUser field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithRunAsUser(value int64) *SecurityContextApplyConfiguration {
	b.RunAsUser = &value
	return b
}

// WithRunAsGroup sets the RunAsGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunAsGroup field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithRunAsGroup(value int64) *SecurityContextApplyConfiguration {
	b.RunAsGroup = &value
	return b
}

// WithRunAsNonRoot sets the RunAsNonRoot field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunAsNonRoot field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithRunAsNonRoot(value bool) *SecurityContextApplyConfiguration {
	b.RunAsNonRoot = &value
	return b
}

// WithReadOnlyRootFilesystem sets the ReadOnlyRootFilesystem field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnlyRootFilesystem field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithReadOnlyRootFilesystem(value bool) *SecurityContextApplyConfiguration {
	b.ReadOnlyRootFilesystem = &value
	return b
}

// WithAllowPrivilegeEscalation sets the AllowPrivilegeEscalation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AllowPrivilegeEscalation field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithAllowPrivilegeEscalation(value bool) *SecurityContextApplyConfiguration {
	b.AllowPrivilegeEscalation = &value
	return b
}

// WithProcMount sets the ProcMount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProcMount field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithProcMount(value corev1.ProcMountType) *SecurityContextApplyConfiguration {
	b.ProcMount = &value
	return b
}

// WithSeccompProfile sets the SeccompProfile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SeccompProfile field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithSeccompProfile(value *SeccompProfileApplyConfiguration) *SecurityContextApplyConfiguration {
	b.SeccompProfile = value
	return b
}

// WithAppArmorProfile sets the AppArmorProfile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AppArmorProfile field is set to the value of the last call.
func (b *SecurityContextApplyConfiguration) WithAppArmorProfile(value *AppArmorProfileApplyConfiguration) *SecurityContextApplyConfiguration {
	b.AppArmorProfile = value
	return b
}
