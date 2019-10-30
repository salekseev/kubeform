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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SqlFirewallRuleLister helps list SqlFirewallRules.
type SqlFirewallRuleLister interface {
	// List lists all SqlFirewallRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SqlFirewallRule, err error)
	// SqlFirewallRules returns an object that can list and get SqlFirewallRules.
	SqlFirewallRules(namespace string) SqlFirewallRuleNamespaceLister
	SqlFirewallRuleListerExpansion
}

// sqlFirewallRuleLister implements the SqlFirewallRuleLister interface.
type sqlFirewallRuleLister struct {
	indexer cache.Indexer
}

// NewSqlFirewallRuleLister returns a new SqlFirewallRuleLister.
func NewSqlFirewallRuleLister(indexer cache.Indexer) SqlFirewallRuleLister {
	return &sqlFirewallRuleLister{indexer: indexer}
}

// List lists all SqlFirewallRules in the indexer.
func (s *sqlFirewallRuleLister) List(selector labels.Selector) (ret []*v1alpha1.SqlFirewallRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SqlFirewallRule))
	})
	return ret, err
}

// SqlFirewallRules returns an object that can list and get SqlFirewallRules.
func (s *sqlFirewallRuleLister) SqlFirewallRules(namespace string) SqlFirewallRuleNamespaceLister {
	return sqlFirewallRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SqlFirewallRuleNamespaceLister helps list and get SqlFirewallRules.
type SqlFirewallRuleNamespaceLister interface {
	// List lists all SqlFirewallRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SqlFirewallRule, err error)
	// Get retrieves the SqlFirewallRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SqlFirewallRule, error)
	SqlFirewallRuleNamespaceListerExpansion
}

// sqlFirewallRuleNamespaceLister implements the SqlFirewallRuleNamespaceLister
// interface.
type sqlFirewallRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SqlFirewallRules in the indexer for a given namespace.
func (s sqlFirewallRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SqlFirewallRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SqlFirewallRule))
	})
	return ret, err
}

// Get retrieves the SqlFirewallRule from the indexer for a given namespace and name.
func (s sqlFirewallRuleNamespaceLister) Get(name string) (*v1alpha1.SqlFirewallRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sqlfirewallrule"), name)
	}
	return obj.(*v1alpha1.SqlFirewallRule), nil
}
