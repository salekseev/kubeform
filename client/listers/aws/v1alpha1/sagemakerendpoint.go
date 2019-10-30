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

// SagemakerEndpointLister helps list SagemakerEndpoints.
type SagemakerEndpointLister interface {
	// List lists all SagemakerEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SagemakerEndpoint, err error)
	// SagemakerEndpoints returns an object that can list and get SagemakerEndpoints.
	SagemakerEndpoints(namespace string) SagemakerEndpointNamespaceLister
	SagemakerEndpointListerExpansion
}

// sagemakerEndpointLister implements the SagemakerEndpointLister interface.
type sagemakerEndpointLister struct {
	indexer cache.Indexer
}

// NewSagemakerEndpointLister returns a new SagemakerEndpointLister.
func NewSagemakerEndpointLister(indexer cache.Indexer) SagemakerEndpointLister {
	return &sagemakerEndpointLister{indexer: indexer}
}

// List lists all SagemakerEndpoints in the indexer.
func (s *sagemakerEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.SagemakerEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SagemakerEndpoint))
	})
	return ret, err
}

// SagemakerEndpoints returns an object that can list and get SagemakerEndpoints.
func (s *sagemakerEndpointLister) SagemakerEndpoints(namespace string) SagemakerEndpointNamespaceLister {
	return sagemakerEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SagemakerEndpointNamespaceLister helps list and get SagemakerEndpoints.
type SagemakerEndpointNamespaceLister interface {
	// List lists all SagemakerEndpoints in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SagemakerEndpoint, err error)
	// Get retrieves the SagemakerEndpoint from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SagemakerEndpoint, error)
	SagemakerEndpointNamespaceListerExpansion
}

// sagemakerEndpointNamespaceLister implements the SagemakerEndpointNamespaceLister
// interface.
type sagemakerEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SagemakerEndpoints in the indexer for a given namespace.
func (s sagemakerEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SagemakerEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SagemakerEndpoint))
	})
	return ret, err
}

// Get retrieves the SagemakerEndpoint from the indexer for a given namespace and name.
func (s sagemakerEndpointNamespaceLister) Get(name string) (*v1alpha1.SagemakerEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sagemakerendpoint"), name)
	}
	return obj.(*v1alpha1.SagemakerEndpoint), nil
}
