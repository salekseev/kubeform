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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SwfDomainLister helps list SwfDomains.
type SwfDomainLister interface {
	// List lists all SwfDomains in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SwfDomain, err error)
	// SwfDomains returns an object that can list and get SwfDomains.
	SwfDomains(namespace string) SwfDomainNamespaceLister
	SwfDomainListerExpansion
}

// swfDomainLister implements the SwfDomainLister interface.
type swfDomainLister struct {
	indexer cache.Indexer
}

// NewSwfDomainLister returns a new SwfDomainLister.
func NewSwfDomainLister(indexer cache.Indexer) SwfDomainLister {
	return &swfDomainLister{indexer: indexer}
}

// List lists all SwfDomains in the indexer.
func (s *swfDomainLister) List(selector labels.Selector) (ret []*v1alpha1.SwfDomain, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SwfDomain))
	})
	return ret, err
}

// SwfDomains returns an object that can list and get SwfDomains.
func (s *swfDomainLister) SwfDomains(namespace string) SwfDomainNamespaceLister {
	return swfDomainNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SwfDomainNamespaceLister helps list and get SwfDomains.
type SwfDomainNamespaceLister interface {
	// List lists all SwfDomains in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SwfDomain, err error)
	// Get retrieves the SwfDomain from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SwfDomain, error)
	SwfDomainNamespaceListerExpansion
}

// swfDomainNamespaceLister implements the SwfDomainNamespaceLister
// interface.
type swfDomainNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SwfDomains in the indexer for a given namespace.
func (s swfDomainNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SwfDomain, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SwfDomain))
	})
	return ret, err
}

// Get retrieves the SwfDomain from the indexer for a given namespace and name.
func (s swfDomainNamespaceLister) Get(name string) (*v1alpha1.SwfDomain, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("swfdomain"), name)
	}
	return obj.(*v1alpha1.SwfDomain), nil
}
