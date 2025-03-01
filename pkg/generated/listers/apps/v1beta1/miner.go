

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/costa92/k8s-ai/pkg/apis/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MinerLister helps list Miners.
// All objects returned here must be treated as read-only.
type MinerLister interface {
	// List lists all Miners in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Miner, err error)
	// Miners returns an object that can list and get Miners.
	Miners(namespace string) MinerNamespaceLister
	MinerListerExpansion
}

// minerLister implements the MinerLister interface.
type minerLister struct {
	indexer cache.Indexer
}

// NewMinerLister returns a new MinerLister.
func NewMinerLister(indexer cache.Indexer) MinerLister {
	return &minerLister{indexer: indexer}
}

// List lists all Miners in the indexer.
func (s *minerLister) List(selector labels.Selector) (ret []*v1beta1.Miner, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Miner))
	})
	return ret, err
}

// Miners returns an object that can list and get Miners.
func (s *minerLister) Miners(namespace string) MinerNamespaceLister {
	return minerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MinerNamespaceLister helps list and get Miners.
// All objects returned here must be treated as read-only.
type MinerNamespaceLister interface {
	// List lists all Miners in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Miner, err error)
	// Get retrieves the Miner from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Miner, error)
	MinerNamespaceListerExpansion
}

// minerNamespaceLister implements the MinerNamespaceLister
// interface.
type minerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Miners in the indexer for a given namespace.
func (s minerNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Miner, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Miner))
	})
	return ret, err
}

// Get retrieves the Miner from the indexer for a given namespace and name.
func (s minerNamespaceLister) Get(name string) (*v1beta1.Miner, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("miner"), name)
	}
	return obj.(*v1beta1.Miner), nil
}
