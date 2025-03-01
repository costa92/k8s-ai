

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NamespaceLister helps list Namespaces.
// All objects returned here must be treated as read-only.
type NamespaceLister interface {
	// List lists all Namespaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Namespace, err error)
	// Get retrieves the Namespace from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Namespace, error)
	NamespaceListerExpansion
}

// namespaceLister implements the NamespaceLister interface.
type namespaceLister struct {
	indexer cache.Indexer
}

// NewNamespaceLister returns a new NamespaceLister.
func NewNamespaceLister(indexer cache.Indexer) NamespaceLister {
	return &namespaceLister{indexer: indexer}
}

// List lists all Namespaces in the indexer.
func (s *namespaceLister) List(selector labels.Selector) (ret []*v1.Namespace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Namespace))
	})
	return ret, err
}

// Get retrieves the Namespace from the index for a given name.
func (s *namespaceLister) Get(name string) (*v1.Namespace, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("namespace"), name)
	}
	return obj.(*v1.Namespace), nil
}
