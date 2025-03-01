

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	types "k8s.io/apimachinery/pkg/types"
)

// ConfigMapNodeConfigSourceApplyConfiguration represents a declarative configuration of the ConfigMapNodeConfigSource type for use
// with apply.
type ConfigMapNodeConfigSourceApplyConfiguration struct {
	Namespace        *string    `json:"namespace,omitempty"`
	Name             *string    `json:"name,omitempty"`
	UID              *types.UID `json:"uid,omitempty"`
	ResourceVersion  *string    `json:"resourceVersion,omitempty"`
	KubeletConfigKey *string    `json:"kubeletConfigKey,omitempty"`
}

// ConfigMapNodeConfigSourceApplyConfiguration constructs a declarative configuration of the ConfigMapNodeConfigSource type for use with
// apply.
func ConfigMapNodeConfigSource() *ConfigMapNodeConfigSourceApplyConfiguration {
	return &ConfigMapNodeConfigSourceApplyConfiguration{}
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *ConfigMapNodeConfigSourceApplyConfiguration) WithNamespace(value string) *ConfigMapNodeConfigSourceApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ConfigMapNodeConfigSourceApplyConfiguration) WithName(value string) *ConfigMapNodeConfigSourceApplyConfiguration {
	b.Name = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *ConfigMapNodeConfigSourceApplyConfiguration) WithUID(value types.UID) *ConfigMapNodeConfigSourceApplyConfiguration {
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *ConfigMapNodeConfigSourceApplyConfiguration) WithResourceVersion(value string) *ConfigMapNodeConfigSourceApplyConfiguration {
	b.ResourceVersion = &value
	return b
}

// WithKubeletConfigKey sets the KubeletConfigKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KubeletConfigKey field is set to the value of the last call.
func (b *ConfigMapNodeConfigSourceApplyConfiguration) WithKubeletConfigKey(value string) *ConfigMapNodeConfigSourceApplyConfiguration {
	b.KubeletConfigKey = &value
	return b
}
