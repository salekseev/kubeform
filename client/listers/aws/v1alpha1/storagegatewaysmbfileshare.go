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

// StoragegatewaySmbFileShareLister helps list StoragegatewaySmbFileShares.
type StoragegatewaySmbFileShareLister interface {
	// List lists all StoragegatewaySmbFileShares in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StoragegatewaySmbFileShare, err error)
	// StoragegatewaySmbFileShares returns an object that can list and get StoragegatewaySmbFileShares.
	StoragegatewaySmbFileShares(namespace string) StoragegatewaySmbFileShareNamespaceLister
	StoragegatewaySmbFileShareListerExpansion
}

// storagegatewaySmbFileShareLister implements the StoragegatewaySmbFileShareLister interface.
type storagegatewaySmbFileShareLister struct {
	indexer cache.Indexer
}

// NewStoragegatewaySmbFileShareLister returns a new StoragegatewaySmbFileShareLister.
func NewStoragegatewaySmbFileShareLister(indexer cache.Indexer) StoragegatewaySmbFileShareLister {
	return &storagegatewaySmbFileShareLister{indexer: indexer}
}

// List lists all StoragegatewaySmbFileShares in the indexer.
func (s *storagegatewaySmbFileShareLister) List(selector labels.Selector) (ret []*v1alpha1.StoragegatewaySmbFileShare, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoragegatewaySmbFileShare))
	})
	return ret, err
}

// StoragegatewaySmbFileShares returns an object that can list and get StoragegatewaySmbFileShares.
func (s *storagegatewaySmbFileShareLister) StoragegatewaySmbFileShares(namespace string) StoragegatewaySmbFileShareNamespaceLister {
	return storagegatewaySmbFileShareNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StoragegatewaySmbFileShareNamespaceLister helps list and get StoragegatewaySmbFileShares.
type StoragegatewaySmbFileShareNamespaceLister interface {
	// List lists all StoragegatewaySmbFileShares in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StoragegatewaySmbFileShare, err error)
	// Get retrieves the StoragegatewaySmbFileShare from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StoragegatewaySmbFileShare, error)
	StoragegatewaySmbFileShareNamespaceListerExpansion
}

// storagegatewaySmbFileShareNamespaceLister implements the StoragegatewaySmbFileShareNamespaceLister
// interface.
type storagegatewaySmbFileShareNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StoragegatewaySmbFileShares in the indexer for a given namespace.
func (s storagegatewaySmbFileShareNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StoragegatewaySmbFileShare, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoragegatewaySmbFileShare))
	})
	return ret, err
}

// Get retrieves the StoragegatewaySmbFileShare from the indexer for a given namespace and name.
func (s storagegatewaySmbFileShareNamespaceLister) Get(name string) (*v1alpha1.StoragegatewaySmbFileShare, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagegatewaysmbfileshare"), name)
	}
	return obj.(*v1alpha1.StoragegatewaySmbFileShare), nil
}
