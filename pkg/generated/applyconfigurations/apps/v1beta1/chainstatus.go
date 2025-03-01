

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	appsv1beta1 "github.com/costa92/k8s-ai/pkg/apis/apps/v1beta1"
)

// ChainStatusApplyConfiguration represents a declarative configuration of the ChainStatus type for use
// with apply.
type ChainStatusApplyConfiguration struct {
	ConfigMapRef       *LocalObjectReferenceApplyConfiguration `json:"configMapRef,omitempty"`
	MinerRef           *LocalObjectReferenceApplyConfiguration `json:"minerRef,omitempty"`
	ObservedGeneration *int64                                  `json:"observedGeneration,omitempty"`
	Conditions         *appsv1beta1.Conditions                 `json:"conditions,omitempty"`
}

// ChainStatusApplyConfiguration constructs a declarative configuration of the ChainStatus type for use with
// apply.
func ChainStatus() *ChainStatusApplyConfiguration {
	return &ChainStatusApplyConfiguration{}
}

// WithConfigMapRef sets the ConfigMapRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ConfigMapRef field is set to the value of the last call.
func (b *ChainStatusApplyConfiguration) WithConfigMapRef(value *LocalObjectReferenceApplyConfiguration) *ChainStatusApplyConfiguration {
	b.ConfigMapRef = value
	return b
}

// WithMinerRef sets the MinerRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinerRef field is set to the value of the last call.
func (b *ChainStatusApplyConfiguration) WithMinerRef(value *LocalObjectReferenceApplyConfiguration) *ChainStatusApplyConfiguration {
	b.MinerRef = value
	return b
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *ChainStatusApplyConfiguration) WithObservedGeneration(value int64) *ChainStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithConditions sets the Conditions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Conditions field is set to the value of the last call.
func (b *ChainStatusApplyConfiguration) WithConditions(value appsv1beta1.Conditions) *ChainStatusApplyConfiguration {
	b.Conditions = &value
	return b
}
