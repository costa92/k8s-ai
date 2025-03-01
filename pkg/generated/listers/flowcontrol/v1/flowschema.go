

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/flowcontrol/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FlowSchemaLister helps list FlowSchemas.
// All objects returned here must be treated as read-only.
type FlowSchemaLister interface {
	// List lists all FlowSchemas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.FlowSchema, err error)
	// Get retrieves the FlowSchema from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.FlowSchema, error)
	FlowSchemaListerExpansion
}

// flowSchemaLister implements the FlowSchemaLister interface.
type flowSchemaLister struct {
	indexer cache.Indexer
}

// NewFlowSchemaLister returns a new FlowSchemaLister.
func NewFlowSchemaLister(indexer cache.Indexer) FlowSchemaLister {
	return &flowSchemaLister{indexer: indexer}
}

// List lists all FlowSchemas in the indexer.
func (s *flowSchemaLister) List(selector labels.Selector) (ret []*v1.FlowSchema, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.FlowSchema))
	})
	return ret, err
}

// Get retrieves the FlowSchema from the index for a given name.
func (s *flowSchemaLister) Get(name string) (*v1.FlowSchema, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("flowschema"), name)
	}
	return obj.(*v1.FlowSchema), nil
}
