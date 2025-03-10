

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/coordination/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LeaseLister helps list Leases.
// All objects returned here must be treated as read-only.
type LeaseLister interface {
	// List lists all Leases in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Lease, err error)
	// Leases returns an object that can list and get Leases.
	Leases(namespace string) LeaseNamespaceLister
	LeaseListerExpansion
}

// leaseLister implements the LeaseLister interface.
type leaseLister struct {
	indexer cache.Indexer
}

// NewLeaseLister returns a new LeaseLister.
func NewLeaseLister(indexer cache.Indexer) LeaseLister {
	return &leaseLister{indexer: indexer}
}

// List lists all Leases in the indexer.
func (s *leaseLister) List(selector labels.Selector) (ret []*v1.Lease, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Lease))
	})
	return ret, err
}

// Leases returns an object that can list and get Leases.
func (s *leaseLister) Leases(namespace string) LeaseNamespaceLister {
	return leaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LeaseNamespaceLister helps list and get Leases.
// All objects returned here must be treated as read-only.
type LeaseNamespaceLister interface {
	// List lists all Leases in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Lease, err error)
	// Get retrieves the Lease from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Lease, error)
	LeaseNamespaceListerExpansion
}

// leaseNamespaceLister implements the LeaseNamespaceLister
// interface.
type leaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Leases in the indexer for a given namespace.
func (s leaseNamespaceLister) List(selector labels.Selector) (ret []*v1.Lease, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Lease))
	})
	return ret, err
}

// Get retrieves the Lease from the indexer for a given namespace and name.
func (s leaseNamespaceLister) Get(name string) (*v1.Lease, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("lease"), name)
	}
	return obj.(*v1.Lease), nil
}
