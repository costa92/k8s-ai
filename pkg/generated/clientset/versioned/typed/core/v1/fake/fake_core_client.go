

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/costa92/k8s-ai/pkg/generated/clientset/versioned/typed/core/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCoreV1 struct {
	*testing.Fake
}

func (c *FakeCoreV1) ComponentStatuses() v1.ComponentStatusInterface {
	return newFakeComponentStatuses(c)
}

func (c *FakeCoreV1) ConfigMaps(namespace string) v1.ConfigMapInterface {
	return newFakeConfigMaps(c, namespace)
}

func (c *FakeCoreV1) Endpoints(namespace string) v1.EndpointsInterface {
	return newFakeEndpoints(c, namespace)
}

func (c *FakeCoreV1) Events(namespace string) v1.EventInterface {
	return newFakeEvents(c, namespace)
}

func (c *FakeCoreV1) LimitRanges(namespace string) v1.LimitRangeInterface {
	return newFakeLimitRanges(c, namespace)
}

func (c *FakeCoreV1) Namespaces() v1.NamespaceInterface {
	return newFakeNamespaces(c)
}

func (c *FakeCoreV1) Nodes() v1.NodeInterface {
	return newFakeNodes(c)
}

func (c *FakeCoreV1) PersistentVolumes() v1.PersistentVolumeInterface {
	return newFakePersistentVolumes(c)
}

func (c *FakeCoreV1) PersistentVolumeClaims(namespace string) v1.PersistentVolumeClaimInterface {
	return newFakePersistentVolumeClaims(c, namespace)
}

func (c *FakeCoreV1) Pods(namespace string) v1.PodInterface {
	return newFakePods(c, namespace)
}

func (c *FakeCoreV1) PodTemplates(namespace string) v1.PodTemplateInterface {
	return newFakePodTemplates(c, namespace)
}

func (c *FakeCoreV1) ReplicationControllers(namespace string) v1.ReplicationControllerInterface {
	return newFakeReplicationControllers(c, namespace)
}

func (c *FakeCoreV1) ResourceQuotas(namespace string) v1.ResourceQuotaInterface {
	return newFakeResourceQuotas(c, namespace)
}

func (c *FakeCoreV1) Secrets(namespace string) v1.SecretInterface {
	return newFakeSecrets(c, namespace)
}

func (c *FakeCoreV1) Services(namespace string) v1.ServiceInterface {
	return newFakeServices(c, namespace)
}

func (c *FakeCoreV1) ServiceAccounts(namespace string) v1.ServiceAccountInterface {
	return newFakeServiceAccounts(c, namespace)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCoreV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
