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

// SpotInstanceRequestLister helps list SpotInstanceRequests.
type SpotInstanceRequestLister interface {
	// List lists all SpotInstanceRequests in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SpotInstanceRequest, err error)
	// SpotInstanceRequests returns an object that can list and get SpotInstanceRequests.
	SpotInstanceRequests(namespace string) SpotInstanceRequestNamespaceLister
	SpotInstanceRequestListerExpansion
}

// spotInstanceRequestLister implements the SpotInstanceRequestLister interface.
type spotInstanceRequestLister struct {
	indexer cache.Indexer
}

// NewSpotInstanceRequestLister returns a new SpotInstanceRequestLister.
func NewSpotInstanceRequestLister(indexer cache.Indexer) SpotInstanceRequestLister {
	return &spotInstanceRequestLister{indexer: indexer}
}

// List lists all SpotInstanceRequests in the indexer.
func (s *spotInstanceRequestLister) List(selector labels.Selector) (ret []*v1alpha1.SpotInstanceRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SpotInstanceRequest))
	})
	return ret, err
}

// SpotInstanceRequests returns an object that can list and get SpotInstanceRequests.
func (s *spotInstanceRequestLister) SpotInstanceRequests(namespace string) SpotInstanceRequestNamespaceLister {
	return spotInstanceRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SpotInstanceRequestNamespaceLister helps list and get SpotInstanceRequests.
type SpotInstanceRequestNamespaceLister interface {
	// List lists all SpotInstanceRequests in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SpotInstanceRequest, err error)
	// Get retrieves the SpotInstanceRequest from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SpotInstanceRequest, error)
	SpotInstanceRequestNamespaceListerExpansion
}

// spotInstanceRequestNamespaceLister implements the SpotInstanceRequestNamespaceLister
// interface.
type spotInstanceRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SpotInstanceRequests in the indexer for a given namespace.
func (s spotInstanceRequestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SpotInstanceRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SpotInstanceRequest))
	})
	return ret, err
}

// Get retrieves the SpotInstanceRequest from the indexer for a given namespace and name.
func (s spotInstanceRequestNamespaceLister) Get(name string) (*v1alpha1.SpotInstanceRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("spotinstancerequest"), name)
	}
	return obj.(*v1alpha1.SpotInstanceRequest), nil
}
