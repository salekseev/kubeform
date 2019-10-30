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

// SpotDatafeedSubscriptionLister helps list SpotDatafeedSubscriptions.
type SpotDatafeedSubscriptionLister interface {
	// List lists all SpotDatafeedSubscriptions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SpotDatafeedSubscription, err error)
	// SpotDatafeedSubscriptions returns an object that can list and get SpotDatafeedSubscriptions.
	SpotDatafeedSubscriptions(namespace string) SpotDatafeedSubscriptionNamespaceLister
	SpotDatafeedSubscriptionListerExpansion
}

// spotDatafeedSubscriptionLister implements the SpotDatafeedSubscriptionLister interface.
type spotDatafeedSubscriptionLister struct {
	indexer cache.Indexer
}

// NewSpotDatafeedSubscriptionLister returns a new SpotDatafeedSubscriptionLister.
func NewSpotDatafeedSubscriptionLister(indexer cache.Indexer) SpotDatafeedSubscriptionLister {
	return &spotDatafeedSubscriptionLister{indexer: indexer}
}

// List lists all SpotDatafeedSubscriptions in the indexer.
func (s *spotDatafeedSubscriptionLister) List(selector labels.Selector) (ret []*v1alpha1.SpotDatafeedSubscription, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SpotDatafeedSubscription))
	})
	return ret, err
}

// SpotDatafeedSubscriptions returns an object that can list and get SpotDatafeedSubscriptions.
func (s *spotDatafeedSubscriptionLister) SpotDatafeedSubscriptions(namespace string) SpotDatafeedSubscriptionNamespaceLister {
	return spotDatafeedSubscriptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SpotDatafeedSubscriptionNamespaceLister helps list and get SpotDatafeedSubscriptions.
type SpotDatafeedSubscriptionNamespaceLister interface {
	// List lists all SpotDatafeedSubscriptions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SpotDatafeedSubscription, err error)
	// Get retrieves the SpotDatafeedSubscription from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SpotDatafeedSubscription, error)
	SpotDatafeedSubscriptionNamespaceListerExpansion
}

// spotDatafeedSubscriptionNamespaceLister implements the SpotDatafeedSubscriptionNamespaceLister
// interface.
type spotDatafeedSubscriptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SpotDatafeedSubscriptions in the indexer for a given namespace.
func (s spotDatafeedSubscriptionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SpotDatafeedSubscription, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SpotDatafeedSubscription))
	})
	return ret, err
}

// Get retrieves the SpotDatafeedSubscription from the indexer for a given namespace and name.
func (s spotDatafeedSubscriptionNamespaceLister) Get(name string) (*v1alpha1.SpotDatafeedSubscription, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("spotdatafeedsubscription"), name)
	}
	return obj.(*v1alpha1.SpotDatafeedSubscription), nil
}
