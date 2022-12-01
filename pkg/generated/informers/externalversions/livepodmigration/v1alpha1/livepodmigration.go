/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	livepodmigrationv1alpha1 "github.com/ubombar/live-pod-migration/pkg/apis/livepodmigration/v1alpha1"
	versioned "github.com/ubombar/live-pod-migration/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/ubombar/live-pod-migration/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/ubombar/live-pod-migration/pkg/generated/listers/livepodmigration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LivePodMigrationInformer provides access to a shared informer and lister for
// LivePodMigrations.
type LivePodMigrationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LivePodMigrationLister
}

type livePodMigrationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLivePodMigrationInformer constructs a new informer for LivePodMigration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLivePodMigrationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLivePodMigrationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLivePodMigrationInformer constructs a new informer for LivePodMigration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLivePodMigrationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LivepodmigrationV1alpha1().LivePodMigrations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LivepodmigrationV1alpha1().LivePodMigrations(namespace).Watch(context.TODO(), options)
			},
		},
		&livepodmigrationv1alpha1.LivePodMigration{},
		resyncPeriod,
		indexers,
	)
}

func (f *livePodMigrationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLivePodMigrationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *livePodMigrationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&livepodmigrationv1alpha1.LivePodMigration{}, f.defaultInformer)
}

func (f *livePodMigrationInformer) Lister() v1alpha1.LivePodMigrationLister {
	return v1alpha1.NewLivePodMigrationLister(f.Informer().GetIndexer())
}
