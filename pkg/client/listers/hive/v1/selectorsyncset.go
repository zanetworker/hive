// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/hive/pkg/apis/hive/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SelectorSyncSetLister helps list SelectorSyncSets.
// All objects returned here must be treated as read-only.
type SelectorSyncSetLister interface {
	// List lists all SelectorSyncSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.SelectorSyncSet, err error)
	// Get retrieves the SelectorSyncSet from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.SelectorSyncSet, error)
	SelectorSyncSetListerExpansion
}

// selectorSyncSetLister implements the SelectorSyncSetLister interface.
type selectorSyncSetLister struct {
	indexer cache.Indexer
}

// NewSelectorSyncSetLister returns a new SelectorSyncSetLister.
func NewSelectorSyncSetLister(indexer cache.Indexer) SelectorSyncSetLister {
	return &selectorSyncSetLister{indexer: indexer}
}

// List lists all SelectorSyncSets in the indexer.
func (s *selectorSyncSetLister) List(selector labels.Selector) (ret []*v1.SelectorSyncSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SelectorSyncSet))
	})
	return ret, err
}

// Get retrieves the SelectorSyncSet from the index for a given name.
func (s *selectorSyncSetLister) Get(name string) (*v1.SelectorSyncSet, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("selectorsyncset"), name)
	}
	return obj.(*v1.SelectorSyncSet), nil
}
