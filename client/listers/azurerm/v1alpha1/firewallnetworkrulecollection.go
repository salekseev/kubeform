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

// FirewallNetworkRuleCollectionLister helps list FirewallNetworkRuleCollections.
type FirewallNetworkRuleCollectionLister interface {
	// List lists all FirewallNetworkRuleCollections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FirewallNetworkRuleCollection, err error)
	// FirewallNetworkRuleCollections returns an object that can list and get FirewallNetworkRuleCollections.
	FirewallNetworkRuleCollections(namespace string) FirewallNetworkRuleCollectionNamespaceLister
	FirewallNetworkRuleCollectionListerExpansion
}

// firewallNetworkRuleCollectionLister implements the FirewallNetworkRuleCollectionLister interface.
type firewallNetworkRuleCollectionLister struct {
	indexer cache.Indexer
}

// NewFirewallNetworkRuleCollectionLister returns a new FirewallNetworkRuleCollectionLister.
func NewFirewallNetworkRuleCollectionLister(indexer cache.Indexer) FirewallNetworkRuleCollectionLister {
	return &firewallNetworkRuleCollectionLister{indexer: indexer}
}

// List lists all FirewallNetworkRuleCollections in the indexer.
func (s *firewallNetworkRuleCollectionLister) List(selector labels.Selector) (ret []*v1alpha1.FirewallNetworkRuleCollection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FirewallNetworkRuleCollection))
	})
	return ret, err
}

// FirewallNetworkRuleCollections returns an object that can list and get FirewallNetworkRuleCollections.
func (s *firewallNetworkRuleCollectionLister) FirewallNetworkRuleCollections(namespace string) FirewallNetworkRuleCollectionNamespaceLister {
	return firewallNetworkRuleCollectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FirewallNetworkRuleCollectionNamespaceLister helps list and get FirewallNetworkRuleCollections.
type FirewallNetworkRuleCollectionNamespaceLister interface {
	// List lists all FirewallNetworkRuleCollections in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FirewallNetworkRuleCollection, err error)
	// Get retrieves the FirewallNetworkRuleCollection from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FirewallNetworkRuleCollection, error)
	FirewallNetworkRuleCollectionNamespaceListerExpansion
}

// firewallNetworkRuleCollectionNamespaceLister implements the FirewallNetworkRuleCollectionNamespaceLister
// interface.
type firewallNetworkRuleCollectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FirewallNetworkRuleCollections in the indexer for a given namespace.
func (s firewallNetworkRuleCollectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FirewallNetworkRuleCollection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FirewallNetworkRuleCollection))
	})
	return ret, err
}

// Get retrieves the FirewallNetworkRuleCollection from the indexer for a given namespace and name.
func (s firewallNetworkRuleCollectionNamespaceLister) Get(name string) (*v1alpha1.FirewallNetworkRuleCollection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("firewallnetworkrulecollection"), name)
	}
	return obj.(*v1alpha1.FirewallNetworkRuleCollection), nil
}
