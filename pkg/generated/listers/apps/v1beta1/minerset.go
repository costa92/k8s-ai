

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/costa92/k8s-ai/pkg/apis/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MinerSetLister helps list MinerSets.
// All objects returned here must be treated as read-only.
type MinerSetLister interface {
	// List lists all MinerSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.MinerSet, err error)
	// MinerSets returns an object that can list and get MinerSets.
	MinerSets(namespace string) MinerSetNamespaceLister
	MinerSetListerExpansion
}

// minerSetLister implements the MinerSetLister interface.
type minerSetLister struct {
	indexer cache.Indexer
}

// NewMinerSetLister returns a new MinerSetLister.
func NewMinerSetLister(indexer cache.Indexer) MinerSetLister {
	return &minerSetLister{indexer: indexer}
}

// List lists all MinerSets in the indexer.
func (s *minerSetLister) List(selector labels.Selector) (ret []*v1beta1.MinerSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.MinerSet))
	})
	return ret, err
}

// MinerSets returns an object that can list and get MinerSets.
func (s *minerSetLister) MinerSets(namespace string) MinerSetNamespaceLister {
	return minerSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MinerSetNamespaceLister helps list and get MinerSets.
// All objects returned here must be treated as read-only.
type MinerSetNamespaceLister interface {
	// List lists all MinerSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.MinerSet, err error)
	// Get retrieves the MinerSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.MinerSet, error)
	MinerSetNamespaceListerExpansion
}

// minerSetNamespaceLister implements the MinerSetNamespaceLister
// interface.
type minerSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MinerSets in the indexer for a given namespace.
func (s minerSetNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.MinerSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.MinerSet))
	})
	return ret, err
}

// Get retrieves the MinerSet from the indexer for a given namespace and name.
func (s minerSetNamespaceLister) Get(name string) (*v1beta1.MinerSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("minerset"), name)
	}
	return obj.(*v1beta1.MinerSet), nil
}
