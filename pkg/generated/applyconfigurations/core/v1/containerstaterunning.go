

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContainerStateRunningApplyConfiguration represents a declarative configuration of the ContainerStateRunning type for use
// with apply.
type ContainerStateRunningApplyConfiguration struct {
	StartedAt *v1.Time `json:"startedAt,omitempty"`
}

// ContainerStateRunningApplyConfiguration constructs a declarative configuration of the ContainerStateRunning type for use with
// apply.
func ContainerStateRunning() *ContainerStateRunningApplyConfiguration {
	return &ContainerStateRunningApplyConfiguration{}
}

// WithStartedAt sets the StartedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartedAt field is set to the value of the last call.
func (b *ContainerStateRunningApplyConfiguration) WithStartedAt(value v1.Time) *ContainerStateRunningApplyConfiguration {
	b.StartedAt = &value
	return b
}
