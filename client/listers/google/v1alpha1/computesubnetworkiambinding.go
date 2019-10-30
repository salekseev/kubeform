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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ComputeSubnetworkIamBindingLister helps list ComputeSubnetworkIamBindings.
type ComputeSubnetworkIamBindingLister interface {
	// List lists all ComputeSubnetworkIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeSubnetworkIamBinding, err error)
	// ComputeSubnetworkIamBindings returns an object that can list and get ComputeSubnetworkIamBindings.
	ComputeSubnetworkIamBindings(namespace string) ComputeSubnetworkIamBindingNamespaceLister
	ComputeSubnetworkIamBindingListerExpansion
}

// computeSubnetworkIamBindingLister implements the ComputeSubnetworkIamBindingLister interface.
type computeSubnetworkIamBindingLister struct {
	indexer cache.Indexer
}

// NewComputeSubnetworkIamBindingLister returns a new ComputeSubnetworkIamBindingLister.
func NewComputeSubnetworkIamBindingLister(indexer cache.Indexer) ComputeSubnetworkIamBindingLister {
	return &computeSubnetworkIamBindingLister{indexer: indexer}
}

// List lists all ComputeSubnetworkIamBindings in the indexer.
func (s *computeSubnetworkIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeSubnetworkIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeSubnetworkIamBinding))
	})
	return ret, err
}

// ComputeSubnetworkIamBindings returns an object that can list and get ComputeSubnetworkIamBindings.
func (s *computeSubnetworkIamBindingLister) ComputeSubnetworkIamBindings(namespace string) ComputeSubnetworkIamBindingNamespaceLister {
	return computeSubnetworkIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeSubnetworkIamBindingNamespaceLister helps list and get ComputeSubnetworkIamBindings.
type ComputeSubnetworkIamBindingNamespaceLister interface {
	// List lists all ComputeSubnetworkIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeSubnetworkIamBinding, err error)
	// Get retrieves the ComputeSubnetworkIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeSubnetworkIamBinding, error)
	ComputeSubnetworkIamBindingNamespaceListerExpansion
}

// computeSubnetworkIamBindingNamespaceLister implements the ComputeSubnetworkIamBindingNamespaceLister
// interface.
type computeSubnetworkIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeSubnetworkIamBindings in the indexer for a given namespace.
func (s computeSubnetworkIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeSubnetworkIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeSubnetworkIamBinding))
	})
	return ret, err
}

// Get retrieves the ComputeSubnetworkIamBinding from the indexer for a given namespace and name.
func (s computeSubnetworkIamBindingNamespaceLister) Get(name string) (*v1alpha1.ComputeSubnetworkIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computesubnetworkiambinding"), name)
	}
	return obj.(*v1alpha1.ComputeSubnetworkIamBinding), nil
}
