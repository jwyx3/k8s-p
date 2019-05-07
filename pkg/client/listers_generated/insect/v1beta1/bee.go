// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "k8s-practices/pkg/apis/insect/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BeeLister helps list Bees.
type BeeLister interface {
	// List lists all Bees in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.Bee, err error)
	// Bees returns an object that can list and get Bees.
	Bees(namespace string) BeeNamespaceLister
	BeeListerExpansion
}

// beeLister implements the BeeLister interface.
type beeLister struct {
	indexer cache.Indexer
}

// NewBeeLister returns a new BeeLister.
func NewBeeLister(indexer cache.Indexer) BeeLister {
	return &beeLister{indexer: indexer}
}

// List lists all Bees in the indexer.
func (s *beeLister) List(selector labels.Selector) (ret []*v1beta1.Bee, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Bee))
	})
	return ret, err
}

// Bees returns an object that can list and get Bees.
func (s *beeLister) Bees(namespace string) BeeNamespaceLister {
	return beeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BeeNamespaceLister helps list and get Bees.
type BeeNamespaceLister interface {
	// List lists all Bees in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.Bee, err error)
	// Get retrieves the Bee from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.Bee, error)
	BeeNamespaceListerExpansion
}

// beeNamespaceLister implements the BeeNamespaceLister
// interface.
type beeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Bees in the indexer for a given namespace.
func (s beeNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Bee, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Bee))
	})
	return ret, err
}

// Get retrieves the Bee from the indexer for a given namespace and name.
func (s beeNamespaceLister) Get(name string) (*v1beta1.Bee, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("bee"), name)
	}
	return obj.(*v1beta1.Bee), nil
}