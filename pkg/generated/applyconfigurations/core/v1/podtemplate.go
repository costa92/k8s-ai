

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	internal "github.com/costa92/k8s-ai/pkg/generated/applyconfigurations/internal"
	v1 "github.com/costa92/k8s-ai/pkg/generated/applyconfigurations/meta/v1"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
)

// PodTemplateApplyConfiguration represents a declarative configuration of the PodTemplate type for use
// with apply.
type PodTemplateApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Template                         *PodTemplateSpecApplyConfiguration `json:"template,omitempty"`
}

// PodTemplate constructs a declarative configuration of the PodTemplate type for use with
// apply.
func PodTemplate(name, namespace string) *PodTemplateApplyConfiguration {
	b := &PodTemplateApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("PodTemplate")
	b.WithAPIVersion("v1")
	return b
}

// ExtractPodTemplate extracts the applied configuration owned by fieldManager from
// podTemplate. If no managedFields are found in podTemplate for fieldManager, a
// PodTemplateApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// podTemplate must be a unmodified PodTemplate API object that was retrieved from the Kubernetes API.
// ExtractPodTemplate provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractPodTemplate(podTemplate *apicorev1.PodTemplate, fieldManager string) (*PodTemplateApplyConfiguration, error) {
	return extractPodTemplate(podTemplate, fieldManager, "")
}

// ExtractPodTemplateStatus is the same as ExtractPodTemplate except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractPodTemplateStatus(podTemplate *apicorev1.PodTemplate, fieldManager string) (*PodTemplateApplyConfiguration, error) {
	return extractPodTemplate(podTemplate, fieldManager, "status")
}

func extractPodTemplate(podTemplate *apicorev1.PodTemplate, fieldManager string, subresource string) (*PodTemplateApplyConfiguration, error) {
	b := &PodTemplateApplyConfiguration{}
	err := managedfields.ExtractInto(podTemplate, internal.Parser().Type("io.k8s.api.core.v1.PodTemplate"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(podTemplate.Name)
	b.WithNamespace(podTemplate.Namespace)

	b.WithKind("PodTemplate")
	b.WithAPIVersion("v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithKind(value string) *PodTemplateApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithAPIVersion(value string) *PodTemplateApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithName(value string) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithGenerateName(value string) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithNamespace(value string) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithUID(value types.UID) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithResourceVersion(value string) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithGeneration(value int64) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithCreationTimestamp(value metav1.Time) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *PodTemplateApplyConfiguration) WithLabels(entries map[string]string) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *PodTemplateApplyConfiguration) WithAnnotations(entries map[string]string) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *PodTemplateApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *PodTemplateApplyConfiguration) WithFinalizers(values ...string) *PodTemplateApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *PodTemplateApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *PodTemplateApplyConfiguration) WithTemplate(value *PodTemplateSpecApplyConfiguration) *PodTemplateApplyConfiguration {
	b.Template = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *PodTemplateApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.Name
}
