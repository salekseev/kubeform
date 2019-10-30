/*
Copyright The Kubeform Authors.

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
	time "time"

	googlev1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SpannerDatabaseIamPolicyInformer provides access to a shared informer and lister for
// SpannerDatabaseIamPolicies.
type SpannerDatabaseIamPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SpannerDatabaseIamPolicyLister
}

type spannerDatabaseIamPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSpannerDatabaseIamPolicyInformer constructs a new informer for SpannerDatabaseIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSpannerDatabaseIamPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSpannerDatabaseIamPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSpannerDatabaseIamPolicyInformer constructs a new informer for SpannerDatabaseIamPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSpannerDatabaseIamPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GoogleV1alpha1().SpannerDatabaseIamPolicies(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GoogleV1alpha1().SpannerDatabaseIamPolicies(namespace).Watch(options)
			},
		},
		&googlev1alpha1.SpannerDatabaseIamPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *spannerDatabaseIamPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSpannerDatabaseIamPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *spannerDatabaseIamPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&googlev1alpha1.SpannerDatabaseIamPolicy{}, f.defaultInformer)
}

func (f *spannerDatabaseIamPolicyInformer) Lister() v1alpha1.SpannerDatabaseIamPolicyLister {
	return v1alpha1.NewSpannerDatabaseIamPolicyLister(f.Informer().GetIndexer())
}
