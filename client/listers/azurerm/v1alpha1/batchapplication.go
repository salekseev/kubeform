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

// BatchApplicationLister helps list BatchApplications.
type BatchApplicationLister interface {
	// List lists all BatchApplications in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BatchApplication, err error)
	// BatchApplications returns an object that can list and get BatchApplications.
	BatchApplications(namespace string) BatchApplicationNamespaceLister
	BatchApplicationListerExpansion
}

// batchApplicationLister implements the BatchApplicationLister interface.
type batchApplicationLister struct {
	indexer cache.Indexer
}

// NewBatchApplicationLister returns a new BatchApplicationLister.
func NewBatchApplicationLister(indexer cache.Indexer) BatchApplicationLister {
	return &batchApplicationLister{indexer: indexer}
}

// List lists all BatchApplications in the indexer.
func (s *batchApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.BatchApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BatchApplication))
	})
	return ret, err
}

// BatchApplications returns an object that can list and get BatchApplications.
func (s *batchApplicationLister) BatchApplications(namespace string) BatchApplicationNamespaceLister {
	return batchApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BatchApplicationNamespaceLister helps list and get BatchApplications.
type BatchApplicationNamespaceLister interface {
	// List lists all BatchApplications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.BatchApplication, err error)
	// Get retrieves the BatchApplication from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.BatchApplication, error)
	BatchApplicationNamespaceListerExpansion
}

// batchApplicationNamespaceLister implements the BatchApplicationNamespaceLister
// interface.
type batchApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BatchApplications in the indexer for a given namespace.
func (s batchApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BatchApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BatchApplication))
	})
	return ret, err
}

// Get retrieves the BatchApplication from the indexer for a given namespace and name.
func (s batchApplicationNamespaceLister) Get(name string) (*v1alpha1.BatchApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("batchapplication"), name)
	}
	return obj.(*v1alpha1.BatchApplication), nil
}
