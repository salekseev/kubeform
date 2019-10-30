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

// CodedeployAppLister helps list CodedeployApps.
type CodedeployAppLister interface {
	// List lists all CodedeployApps in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CodedeployApp, err error)
	// CodedeployApps returns an object that can list and get CodedeployApps.
	CodedeployApps(namespace string) CodedeployAppNamespaceLister
	CodedeployAppListerExpansion
}

// codedeployAppLister implements the CodedeployAppLister interface.
type codedeployAppLister struct {
	indexer cache.Indexer
}

// NewCodedeployAppLister returns a new CodedeployAppLister.
func NewCodedeployAppLister(indexer cache.Indexer) CodedeployAppLister {
	return &codedeployAppLister{indexer: indexer}
}

// List lists all CodedeployApps in the indexer.
func (s *codedeployAppLister) List(selector labels.Selector) (ret []*v1alpha1.CodedeployApp, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodedeployApp))
	})
	return ret, err
}

// CodedeployApps returns an object that can list and get CodedeployApps.
func (s *codedeployAppLister) CodedeployApps(namespace string) CodedeployAppNamespaceLister {
	return codedeployAppNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CodedeployAppNamespaceLister helps list and get CodedeployApps.
type CodedeployAppNamespaceLister interface {
	// List lists all CodedeployApps in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CodedeployApp, err error)
	// Get retrieves the CodedeployApp from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CodedeployApp, error)
	CodedeployAppNamespaceListerExpansion
}

// codedeployAppNamespaceLister implements the CodedeployAppNamespaceLister
// interface.
type codedeployAppNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CodedeployApps in the indexer for a given namespace.
func (s codedeployAppNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CodedeployApp, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodedeployApp))
	})
	return ret, err
}

// Get retrieves the CodedeployApp from the indexer for a given namespace and name.
func (s codedeployAppNamespaceLister) Get(name string) (*v1alpha1.CodedeployApp, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("codedeployapp"), name)
	}
	return obj.(*v1alpha1.CodedeployApp), nil
}
