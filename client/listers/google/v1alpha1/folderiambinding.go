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

// FolderIamBindingLister helps list FolderIamBindings.
type FolderIamBindingLister interface {
	// List lists all FolderIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FolderIamBinding, err error)
	// FolderIamBindings returns an object that can list and get FolderIamBindings.
	FolderIamBindings(namespace string) FolderIamBindingNamespaceLister
	FolderIamBindingListerExpansion
}

// folderIamBindingLister implements the FolderIamBindingLister interface.
type folderIamBindingLister struct {
	indexer cache.Indexer
}

// NewFolderIamBindingLister returns a new FolderIamBindingLister.
func NewFolderIamBindingLister(indexer cache.Indexer) FolderIamBindingLister {
	return &folderIamBindingLister{indexer: indexer}
}

// List lists all FolderIamBindings in the indexer.
func (s *folderIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.FolderIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FolderIamBinding))
	})
	return ret, err
}

// FolderIamBindings returns an object that can list and get FolderIamBindings.
func (s *folderIamBindingLister) FolderIamBindings(namespace string) FolderIamBindingNamespaceLister {
	return folderIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FolderIamBindingNamespaceLister helps list and get FolderIamBindings.
type FolderIamBindingNamespaceLister interface {
	// List lists all FolderIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FolderIamBinding, err error)
	// Get retrieves the FolderIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FolderIamBinding, error)
	FolderIamBindingNamespaceListerExpansion
}

// folderIamBindingNamespaceLister implements the FolderIamBindingNamespaceLister
// interface.
type folderIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FolderIamBindings in the indexer for a given namespace.
func (s folderIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FolderIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FolderIamBinding))
	})
	return ret, err
}

// Get retrieves the FolderIamBinding from the indexer for a given namespace and name.
func (s folderIamBindingNamespaceLister) Get(name string) (*v1alpha1.FolderIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("folderiambinding"), name)
	}
	return obj.(*v1alpha1.FolderIamBinding), nil
}
