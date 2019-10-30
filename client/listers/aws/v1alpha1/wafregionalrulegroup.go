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

// WafregionalRuleGroupLister helps list WafregionalRuleGroups.
type WafregionalRuleGroupLister interface {
	// List lists all WafregionalRuleGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalRuleGroup, err error)
	// WafregionalRuleGroups returns an object that can list and get WafregionalRuleGroups.
	WafregionalRuleGroups(namespace string) WafregionalRuleGroupNamespaceLister
	WafregionalRuleGroupListerExpansion
}

// wafregionalRuleGroupLister implements the WafregionalRuleGroupLister interface.
type wafregionalRuleGroupLister struct {
	indexer cache.Indexer
}

// NewWafregionalRuleGroupLister returns a new WafregionalRuleGroupLister.
func NewWafregionalRuleGroupLister(indexer cache.Indexer) WafregionalRuleGroupLister {
	return &wafregionalRuleGroupLister{indexer: indexer}
}

// List lists all WafregionalRuleGroups in the indexer.
func (s *wafregionalRuleGroupLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalRuleGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalRuleGroup))
	})
	return ret, err
}

// WafregionalRuleGroups returns an object that can list and get WafregionalRuleGroups.
func (s *wafregionalRuleGroupLister) WafregionalRuleGroups(namespace string) WafregionalRuleGroupNamespaceLister {
	return wafregionalRuleGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WafregionalRuleGroupNamespaceLister helps list and get WafregionalRuleGroups.
type WafregionalRuleGroupNamespaceLister interface {
	// List lists all WafregionalRuleGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WafregionalRuleGroup, err error)
	// Get retrieves the WafregionalRuleGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WafregionalRuleGroup, error)
	WafregionalRuleGroupNamespaceListerExpansion
}

// wafregionalRuleGroupNamespaceLister implements the WafregionalRuleGroupNamespaceLister
// interface.
type wafregionalRuleGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WafregionalRuleGroups in the indexer for a given namespace.
func (s wafregionalRuleGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WafregionalRuleGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WafregionalRuleGroup))
	})
	return ret, err
}

// Get retrieves the WafregionalRuleGroup from the indexer for a given namespace and name.
func (s wafregionalRuleGroupNamespaceLister) Get(name string) (*v1alpha1.WafregionalRuleGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("wafregionalrulegroup"), name)
	}
	return obj.(*v1alpha1.WafregionalRuleGroup), nil
}
