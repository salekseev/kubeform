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

// ProjectServiceLister helps list ProjectServices.
type ProjectServiceLister interface {
	// List lists all ProjectServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ProjectService, err error)
	// ProjectServices returns an object that can list and get ProjectServices.
	ProjectServices(namespace string) ProjectServiceNamespaceLister
	ProjectServiceListerExpansion
}

// projectServiceLister implements the ProjectServiceLister interface.
type projectServiceLister struct {
	indexer cache.Indexer
}

// NewProjectServiceLister returns a new ProjectServiceLister.
func NewProjectServiceLister(indexer cache.Indexer) ProjectServiceLister {
	return &projectServiceLister{indexer: indexer}
}

// List lists all ProjectServices in the indexer.
func (s *projectServiceLister) List(selector labels.Selector) (ret []*v1alpha1.ProjectService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProjectService))
	})
	return ret, err
}

// ProjectServices returns an object that can list and get ProjectServices.
func (s *projectServiceLister) ProjectServices(namespace string) ProjectServiceNamespaceLister {
	return projectServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProjectServiceNamespaceLister helps list and get ProjectServices.
type ProjectServiceNamespaceLister interface {
	// List lists all ProjectServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ProjectService, err error)
	// Get retrieves the ProjectService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ProjectService, error)
	ProjectServiceNamespaceListerExpansion
}

// projectServiceNamespaceLister implements the ProjectServiceNamespaceLister
// interface.
type projectServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProjectServices in the indexer for a given namespace.
func (s projectServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProjectService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProjectService))
	})
	return ret, err
}

// Get retrieves the ProjectService from the indexer for a given namespace and name.
func (s projectServiceNamespaceLister) Get(name string) (*v1alpha1.ProjectService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("projectservice"), name)
	}
	return obj.(*v1alpha1.ProjectService), nil
}
