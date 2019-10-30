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

// RedshiftEventSubscriptionLister helps list RedshiftEventSubscriptions.
type RedshiftEventSubscriptionLister interface {
	// List lists all RedshiftEventSubscriptions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RedshiftEventSubscription, err error)
	// RedshiftEventSubscriptions returns an object that can list and get RedshiftEventSubscriptions.
	RedshiftEventSubscriptions(namespace string) RedshiftEventSubscriptionNamespaceLister
	RedshiftEventSubscriptionListerExpansion
}

// redshiftEventSubscriptionLister implements the RedshiftEventSubscriptionLister interface.
type redshiftEventSubscriptionLister struct {
	indexer cache.Indexer
}

// NewRedshiftEventSubscriptionLister returns a new RedshiftEventSubscriptionLister.
func NewRedshiftEventSubscriptionLister(indexer cache.Indexer) RedshiftEventSubscriptionLister {
	return &redshiftEventSubscriptionLister{indexer: indexer}
}

// List lists all RedshiftEventSubscriptions in the indexer.
func (s *redshiftEventSubscriptionLister) List(selector labels.Selector) (ret []*v1alpha1.RedshiftEventSubscription, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedshiftEventSubscription))
	})
	return ret, err
}

// RedshiftEventSubscriptions returns an object that can list and get RedshiftEventSubscriptions.
func (s *redshiftEventSubscriptionLister) RedshiftEventSubscriptions(namespace string) RedshiftEventSubscriptionNamespaceLister {
	return redshiftEventSubscriptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RedshiftEventSubscriptionNamespaceLister helps list and get RedshiftEventSubscriptions.
type RedshiftEventSubscriptionNamespaceLister interface {
	// List lists all RedshiftEventSubscriptions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RedshiftEventSubscription, err error)
	// Get retrieves the RedshiftEventSubscription from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RedshiftEventSubscription, error)
	RedshiftEventSubscriptionNamespaceListerExpansion
}

// redshiftEventSubscriptionNamespaceLister implements the RedshiftEventSubscriptionNamespaceLister
// interface.
type redshiftEventSubscriptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RedshiftEventSubscriptions in the indexer for a given namespace.
func (s redshiftEventSubscriptionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RedshiftEventSubscription, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedshiftEventSubscription))
	})
	return ret, err
}

// Get retrieves the RedshiftEventSubscription from the indexer for a given namespace and name.
func (s redshiftEventSubscriptionNamespaceLister) Get(name string) (*v1alpha1.RedshiftEventSubscription, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("redshifteventsubscription"), name)
	}
	return obj.(*v1alpha1.RedshiftEventSubscription), nil
}
