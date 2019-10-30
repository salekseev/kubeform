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

// StreamAnalyticsStreamInputBlobLister helps list StreamAnalyticsStreamInputBlobs.
type StreamAnalyticsStreamInputBlobLister interface {
	// List lists all StreamAnalyticsStreamInputBlobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsStreamInputBlob, err error)
	// StreamAnalyticsStreamInputBlobs returns an object that can list and get StreamAnalyticsStreamInputBlobs.
	StreamAnalyticsStreamInputBlobs(namespace string) StreamAnalyticsStreamInputBlobNamespaceLister
	StreamAnalyticsStreamInputBlobListerExpansion
}

// streamAnalyticsStreamInputBlobLister implements the StreamAnalyticsStreamInputBlobLister interface.
type streamAnalyticsStreamInputBlobLister struct {
	indexer cache.Indexer
}

// NewStreamAnalyticsStreamInputBlobLister returns a new StreamAnalyticsStreamInputBlobLister.
func NewStreamAnalyticsStreamInputBlobLister(indexer cache.Indexer) StreamAnalyticsStreamInputBlobLister {
	return &streamAnalyticsStreamInputBlobLister{indexer: indexer}
}

// List lists all StreamAnalyticsStreamInputBlobs in the indexer.
func (s *streamAnalyticsStreamInputBlobLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsStreamInputBlob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsStreamInputBlob))
	})
	return ret, err
}

// StreamAnalyticsStreamInputBlobs returns an object that can list and get StreamAnalyticsStreamInputBlobs.
func (s *streamAnalyticsStreamInputBlobLister) StreamAnalyticsStreamInputBlobs(namespace string) StreamAnalyticsStreamInputBlobNamespaceLister {
	return streamAnalyticsStreamInputBlobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamAnalyticsStreamInputBlobNamespaceLister helps list and get StreamAnalyticsStreamInputBlobs.
type StreamAnalyticsStreamInputBlobNamespaceLister interface {
	// List lists all StreamAnalyticsStreamInputBlobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsStreamInputBlob, err error)
	// Get retrieves the StreamAnalyticsStreamInputBlob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StreamAnalyticsStreamInputBlob, error)
	StreamAnalyticsStreamInputBlobNamespaceListerExpansion
}

// streamAnalyticsStreamInputBlobNamespaceLister implements the StreamAnalyticsStreamInputBlobNamespaceLister
// interface.
type streamAnalyticsStreamInputBlobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StreamAnalyticsStreamInputBlobs in the indexer for a given namespace.
func (s streamAnalyticsStreamInputBlobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StreamAnalyticsStreamInputBlob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamAnalyticsStreamInputBlob))
	})
	return ret, err
}

// Get retrieves the StreamAnalyticsStreamInputBlob from the indexer for a given namespace and name.
func (s streamAnalyticsStreamInputBlobNamespaceLister) Get(name string) (*v1alpha1.StreamAnalyticsStreamInputBlob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("streamanalyticsstreaminputblob"), name)
	}
	return obj.(*v1alpha1.StreamAnalyticsStreamInputBlob), nil
}
