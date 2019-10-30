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

// WafregionalWebACLLister helps list WafregionalWebACLs.
type WafregionalWebACLLister interface {
	// List lists all WafregionalWebACLs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalWebACL, err error)
	// WafregionalWebACLs returns an object that can list and get WafregionalWebACLs.
	WafregionalWebACLs(namespace string) WafregionalWebACLNamespaceLister
	WafregionalWebACLListerExpansion
}

// wafregionalWebACLLister implements the WafregionalWebACLLister interface.
type wafregionalWebACLLister struct {
	indexer cache.Indexer
}

// NewWafregionalWebACLLister returns a new WafregionalWebACLLister.
func NewWafregionalWebACLLister(indexer cache.Indexer) WafregionalWebACLLister {
	return &wafregionalWebACLLister{indexer: indexer}
}

// List lists all WafregionalWebACLs in the indexer.
func (s *wafregionalWebACLLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalWebACL, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalWebACL))
	})
	return ret, err
}

// WafregionalWebACLs returns an object that can list and get WafregionalWebACLs.
func (s *wafregionalWebACLLister) WafregionalWebACLs(namespace string) WafregionalWebACLNamespaceLister {
	return wafregionalWebACLNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WafregionalWebACLNamespaceLister helps list and get WafregionalWebACLs.
type WafregionalWebACLNamespaceLister interface {
	// List lists all WafregionalWebACLs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalWebACL, err error)
	// Get retrieves the WafregionalWebACL from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WafregionalWebACL, error)
	WafregionalWebACLNamespaceListerExpansion
}

// wafregionalWebACLNamespaceLister implements the WafregionalWebACLNamespaceLister
// interface.
type wafregionalWebACLNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WafregionalWebACLs in the indexer for a given namespace.
func (s wafregionalWebACLNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalWebACL, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalWebACL))
	})
	return ret, err
}

// Get retrieves the WafregionalWebACL from the indexer for a given namespace and name.
func (s wafregionalWebACLNamespaceLister) Get(name string) (*v1alpha1.WafregionalWebACL, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafregionalwebacl"), name)
	}
	return obj.(*v1alpha1.WafregionalWebACL), nil
}
