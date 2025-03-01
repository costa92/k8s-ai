

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	versioned "github.com/costa92/k8s-ai/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/costa92/k8s-ai/pkg/generated/informers/internalinterfaces"
	corev1 "github.com/costa92/k8s-ai/pkg/generated/listers/core/v1"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceAccountInformer provides access to a shared informer and lister for
// ServiceAccounts.
type ServiceAccountInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() corev1.ServiceAccountLister
}

type serviceAccountInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceAccountInformer constructs a new informer for ServiceAccount type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceAccountInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceAccountInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceAccountInformer constructs a new informer for ServiceAccount type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceAccountInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().ServiceAccounts(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1().ServiceAccounts(namespace).Watch(context.TODO(), options)
			},
		},
		&apicorev1.ServiceAccount{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceAccountInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceAccountInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceAccountInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apicorev1.ServiceAccount{}, f.defaultInformer)
}

func (f *serviceAccountInformer) Lister() corev1.ServiceAccountLister {
	return corev1.NewServiceAccountLister(f.Informer().GetIndexer())
}
