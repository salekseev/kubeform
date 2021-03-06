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

// WafXssMatchSetLister helps list WafXssMatchSets.
type WafXssMatchSetLister interface {
	// List lists all WafXssMatchSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WafXssMatchSet, err error)
	// WafXssMatchSets returns an object that can list and get WafXssMatchSets.
	WafXssMatchSets(namespace string) WafXssMatchSetNamespaceLister
	WafXssMatchSetListerExpansion
}

// wafXssMatchSetLister implements the WafXssMatchSetLister interface.
type wafXssMatchSetLister struct {
	indexer cache.Indexer
}

// NewWafXssMatchSetLister returns a new WafXssMatchSetLister.
func NewWafXssMatchSetLister(indexer cache.Indexer) WafXssMatchSetLister {
	return &wafXssMatchSetLister{indexer: indexer}
}

// List lists all WafXssMatchSets in the indexer.
func (s *wafXssMatchSetLister) List(selector labels.Selector) (ret []*v1alpha1.WafXssMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafXssMatchSet))
	})
	return ret, err
}

// WafXssMatchSets returns an object that can list and get WafXssMatchSets.
func (s *wafXssMatchSetLister) WafXssMatchSets(namespace string) WafXssMatchSetNamespaceLister {
	return wafXssMatchSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WafXssMatchSetNamespaceLister helps list and get WafXssMatchSets.
type WafXssMatchSetNamespaceLister interface {
	// List lists all WafXssMatchSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WafXssMatchSet, err error)
	// Get retrieves the WafXssMatchSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WafXssMatchSet, error)
	WafXssMatchSetNamespaceListerExpansion
}

// wafXssMatchSetNamespaceLister implements the WafXssMatchSetNamespaceLister
// interface.
type wafXssMatchSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WafXssMatchSets in the indexer for a given namespace.
func (s wafXssMatchSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WafXssMatchSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafXssMatchSet))
	})
	return ret, err
}

// Get retrieves the WafXssMatchSet from the indexer for a given namespace and name.
func (s wafXssMatchSetNamespaceLister) Get(name string) (*v1alpha1.WafXssMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafxssmatchset"), name)
	}
	return obj.(*v1alpha1.WafXssMatchSet), nil
}
