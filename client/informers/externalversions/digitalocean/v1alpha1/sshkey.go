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

	digitaloceanv1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
	versioned "kubeform.dev/kubeform/client/clientset/versioned"
	internalinterfaces "kubeform.dev/kubeform/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/kubeform/client/listers/digitalocean/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SshKeyInformer provides access to a shared informer and lister for
// SshKeys.
type SshKeyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SshKeyLister
}

type sshKeyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSshKeyInformer constructs a new informer for SshKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSshKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSshKeyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSshKeyInformer constructs a new informer for SshKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSshKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DigitaloceanV1alpha1().SshKeys(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DigitaloceanV1alpha1().SshKeys(namespace).Watch(options)
			},
		},
		&digitaloceanv1alpha1.SshKey{},
		resyncPeriod,
		indexers,
	)
}

func (f *sshKeyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSshKeyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sshKeyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&digitaloceanv1alpha1.SshKey{}, f.defaultInformer)
}

func (f *sshKeyInformer) Lister() v1alpha1.SshKeyLister {
	return v1alpha1.NewSshKeyLister(f.Informer().GetIndexer())
}
