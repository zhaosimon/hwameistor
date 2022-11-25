// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LocalDiskNodeLister helps list LocalDiskNodes.
// All objects returned here must be treated as read-only.
type LocalDiskNodeLister interface {
	// List lists all LocalDiskNodes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LocalDiskNode, err error)
	// Get retrieves the LocalDiskNode from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LocalDiskNode, error)
	LocalDiskNodeListerExpansion
}

// localDiskNodeLister implements the LocalDiskNodeLister interface.
type localDiskNodeLister struct {
	indexer cache.Indexer
}

// NewLocalDiskNodeLister returns a new LocalDiskNodeLister.
func NewLocalDiskNodeLister(indexer cache.Indexer) LocalDiskNodeLister {
	return &localDiskNodeLister{indexer: indexer}
}

// List lists all LocalDiskNodes in the indexer.
func (s *localDiskNodeLister) List(selector labels.Selector) (ret []*v1alpha1.LocalDiskNode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LocalDiskNode))
	})
	return ret, err
}

// Get retrieves the LocalDiskNode from the index for a given name.
func (s *localDiskNodeLister) Get(name string) (*v1alpha1.LocalDiskNode, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("localdisknode"), name)
	}
	return obj.(*v1alpha1.LocalDiskNode), nil
}
