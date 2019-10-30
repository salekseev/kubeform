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

// StreamAnalyticsOutputEventhubLister helps list StreamAnalyticsOutputEventhubs.
type StreamAnalyticsOutputEventhubLister interface {
	// List lists all StreamAnalyticsOutputEventhubs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputEventhub, err error)
	// StreamAnalyticsOutputEventhubs returns an object that can list and get StreamAnalyticsOutputEventhubs.
	StreamAnalyticsOutputEventhubs(namespace string) StreamAnalyticsOutputEventhubNamespaceLister
	StreamAnalyticsOutputEventhubListerExpansion
}

// streamAnalyticsOutputEventhubLister implements the StreamAnalyticsOutputEventhubLister interface.
type streamAnalyticsOutputEventhubLister struct {
	indexer cache.Indexer
}

// NewStreamAnalyticsOutputEventhubLister returns a new StreamAnalyticsOutputEventhubLister.
func NewStreamAnalyticsOutputEventhubLister(indexer cache.Indexer) StreamAnalyticsOutputEventhubLister {
	return &streamAnalyticsOutputEventhubLister{indexer: indexer}
}

// List lists all StreamAnalyticsOutputEventhubs in the indexer.
func (s *streamAnalyticsOutputEventhubLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputEventhub, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsOutputEventhub))
	})
	return ret, err
}

// StreamAnalyticsOutputEventhubs returns an object that can list and get StreamAnalyticsOutputEventhubs.
func (s *streamAnalyticsOutputEventhubLister) StreamAnalyticsOutputEventhubs(namespace string) StreamAnalyticsOutputEventhubNamespaceLister {
	return streamAnalyticsOutputEventhubNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamAnalyticsOutputEventhubNamespaceLister helps list and get StreamAnalyticsOutputEventhubs.
type StreamAnalyticsOutputEventhubNamespaceLister interface {
	// List lists all StreamAnalyticsOutputEventhubs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputEventhub, err error)
	// Get retrieves the StreamAnalyticsOutputEventhub from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StreamAnalyticsOutputEventhub, error)
	StreamAnalyticsOutputEventhubNamespaceListerExpansion
}

// streamAnalyticsOutputEventhubNamespaceLister implements the StreamAnalyticsOutputEventhubNamespaceLister
// interface.
type streamAnalyticsOutputEventhubNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StreamAnalyticsOutputEventhubs in the indexer for a given namespace.
func (s streamAnalyticsOutputEventhubNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsOutputEventhub, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsOutputEventhub))
	})
	return ret, err
}

// Get retrieves the StreamAnalyticsOutputEventhub from the indexer for a given namespace and name.
func (s streamAnalyticsOutputEventhubNamespaceLister) Get(name string) (*v1alpha1.StreamAnalyticsOutputEventhub, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("streamanalyticsoutputeventhub"), name)
	}
	return obj.(*v1alpha1.StreamAnalyticsOutputEventhub), nil
}
