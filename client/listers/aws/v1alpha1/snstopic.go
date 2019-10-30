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

// SnsTopicLister helps list SnsTopics.
type SnsTopicLister interface {
	// List lists all SnsTopics in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SnsTopic, err error)
	// SnsTopics returns an object that can list and get SnsTopics.
	SnsTopics(namespace string) SnsTopicNamespaceLister
	SnsTopicListerExpansion
}

// snsTopicLister implements the SnsTopicLister interface.
type snsTopicLister struct {
	indexer cache.Indexer
}

// NewSnsTopicLister returns a new SnsTopicLister.
func NewSnsTopicLister(indexer cache.Indexer) SnsTopicLister {
	return &snsTopicLister{indexer: indexer}
}

// List lists all SnsTopics in the indexer.
func (s *snsTopicLister) List(selector labels.Selector) (ret []*v1alpha1.SnsTopic, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SnsTopic))
	})
	return ret, err
}

// SnsTopics returns an object that can list and get SnsTopics.
func (s *snsTopicLister) SnsTopics(namespace string) SnsTopicNamespaceLister {
	return snsTopicNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SnsTopicNamespaceLister helps list and get SnsTopics.
type SnsTopicNamespaceLister interface {
	// List lists all SnsTopics in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SnsTopic, err error)
	// Get retrieves the SnsTopic from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SnsTopic, error)
	SnsTopicNamespaceListerExpansion
}

// snsTopicNamespaceLister implements the SnsTopicNamespaceLister
// interface.
type snsTopicNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SnsTopics in the indexer for a given namespace.
func (s snsTopicNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SnsTopic, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SnsTopic))
	})
	return ret, err
}

// Get retrieves the SnsTopic from the indexer for a given namespace and name.
func (s snsTopicNamespaceLister) Get(name string) (*v1alpha1.SnsTopic, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("snstopic"), name)
	}
	return obj.(*v1alpha1.SnsTopic), nil
}
