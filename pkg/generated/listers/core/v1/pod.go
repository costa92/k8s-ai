

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodLister helps list Pods.
// All objects returned here must be treated as read-only.
type PodLister interface {
	// List lists all Pods in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Pod, err error)
	// Pods returns an object that can list and get Pods.
	Pods(namespace string) PodNamespaceLister
	PodListerExpansion
}

// podLister implements the PodLister interface.
type podLister struct {
	indexer cache.Indexer
}

// NewPodLister returns a new PodLister.
func NewPodLister(indexer cache.Indexer) PodLister {
	return &podLister{indexer: indexer}
}

// List lists all Pods in the indexer.
func (s *podLister) List(selector labels.Selector) (ret []*v1.Pod, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Pod))
	})
	return ret, err
}

// Pods returns an object that can list and get Pods.
func (s *podLister) Pods(namespace string) PodNamespaceLister {
	return podNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodNamespaceLister helps list and get Pods.
// All objects returned here must be treated as read-only.
type PodNamespaceLister interface {
	// List lists all Pods in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Pod, err error)
	// Get retrieves the Pod from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Pod, error)
	PodNamespaceListerExpansion
}

// podNamespaceLister implements the PodNamespaceLister
// interface.
type podNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Pods in the indexer for a given namespace.
func (s podNamespaceLister) List(selector labels.Selector) (ret []*v1.Pod, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Pod))
	})
	return ret, err
}

// Get retrieves the Pod from the indexer for a given namespace and name.
func (s podNamespaceLister) Get(name string) (*v1.Pod, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("pod"), name)
	}
	return obj.(*v1.Pod), nil
}
