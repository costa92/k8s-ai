

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

type ConfigMapExpansion any

type EventExpansion any

type SecretExpansion any

type ComponentStatusExpansion any

type EndpointsExpansion any

type LimitRangeExpansion any

type NodeExpansion any

type PersistentVolumeExpansion any

type PersistentVolumeClaimExpansion any

type PodExpansion any

type PodTemplateExpansion any

type ReplicationControllerExpansion any

type ResourceQuotaExpansion any

type ServiceExpansion any

type ServiceAccountExpansion any

// The NamespaceExpansion interface allows manually adding extra methods to the NamespaceInterface.
type NamespaceExpansion interface {
	Finalize(ctx context.Context, item *v1.Namespace, opts metav1.UpdateOptions) (*v1.Namespace, error)
}

// Finalize takes the representation of a namespace to update.  Returns the server's representation of the namespace, and an error, if it occurs.
func (c *namespaces) Finalize(ctx context.Context, namespace *v1.Namespace, opts metav1.UpdateOptions) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
    err = c.GetClient().Put().Resource("namespaces").Name(namespace.Name).VersionedParams(&opts, scheme.ParameterCodec).SubResource("finalize").Body(namespace).Do(ctx).Into(result)
    return
}
