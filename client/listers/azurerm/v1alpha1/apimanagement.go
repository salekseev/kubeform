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

// ApiManagementLister helps list ApiManagements.
type ApiManagementLister interface {
	// List lists all ApiManagements in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagement, err error)
	// ApiManagements returns an object that can list and get ApiManagements.
	ApiManagements(namespace string) ApiManagementNamespaceLister
	ApiManagementListerExpansion
}

// apiManagementLister implements the ApiManagementLister interface.
type apiManagementLister struct {
	indexer cache.Indexer
}

// NewApiManagementLister returns a new ApiManagementLister.
func NewApiManagementLister(indexer cache.Indexer) ApiManagementLister {
	return &apiManagementLister{indexer: indexer}
}

// List lists all ApiManagements in the indexer.
func (s *apiManagementLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagement, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagement))
	})
	return ret, err
}

// ApiManagements returns an object that can list and get ApiManagements.
func (s *apiManagementLister) ApiManagements(namespace string) ApiManagementNamespaceLister {
	return apiManagementNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiManagementNamespaceLister helps list and get ApiManagements.
type ApiManagementNamespaceLister interface {
	// List lists all ApiManagements in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiManagement, err error)
	// Get retrieves the ApiManagement from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiManagement, error)
	ApiManagementNamespaceListerExpansion
}

// apiManagementNamespaceLister implements the ApiManagementNamespaceLister
// interface.
type apiManagementNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiManagements in the indexer for a given namespace.
func (s apiManagementNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiManagement, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiManagement))
	})
	return ret, err
}

// Get retrieves the ApiManagement from the indexer for a given namespace and name.
func (s apiManagementNamespaceLister) Get(name string) (*v1alpha1.ApiManagement, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimanagement"), name)
	}
	return obj.(*v1alpha1.ApiManagement), nil
}
