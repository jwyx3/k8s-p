// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	insectv1beta1 "k8s-practices/pkg/apis/insect/v1beta1"
	clientset "k8s-practices/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s-practices/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1beta1 "k8s-practices/pkg/client/listers_generated/insect/v1beta1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BeeInformer provides access to a shared informer and lister for
// Bees.
type BeeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.BeeLister
}

type beeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBeeInformer constructs a new informer for Bee type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBeeInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBeeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBeeInformer constructs a new informer for Bee type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBeeInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InsectV1beta1().Bees(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InsectV1beta1().Bees(namespace).Watch(options)
			},
		},
		&insectv1beta1.Bee{},
		resyncPeriod,
		indexers,
	)
}

func (f *beeInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBeeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *beeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&insectv1beta1.Bee{}, f.defaultInformer)
}

func (f *beeInformer) Lister() v1beta1.BeeLister {
	return v1beta1.NewBeeLister(f.Informer().GetIndexer())
}