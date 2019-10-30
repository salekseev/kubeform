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

// Cloud9EnvironmentEc2Lister helps list Cloud9EnvironmentEc2s.
type Cloud9EnvironmentEc2Lister interface {
	// List lists all Cloud9EnvironmentEc2s in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Cloud9EnvironmentEc2, err error)
	// Cloud9EnvironmentEc2s returns an object that can list and get Cloud9EnvironmentEc2s.
	Cloud9EnvironmentEc2s(namespace string) Cloud9EnvironmentEc2NamespaceLister
	Cloud9EnvironmentEc2ListerExpansion
}

// cloud9EnvironmentEc2Lister implements the Cloud9EnvironmentEc2Lister interface.
type cloud9EnvironmentEc2Lister struct {
	indexer cache.Indexer
}

// NewCloud9EnvironmentEc2Lister returns a new Cloud9EnvironmentEc2Lister.
func NewCloud9EnvironmentEc2Lister(indexer cache.Indexer) Cloud9EnvironmentEc2Lister {
	return &cloud9EnvironmentEc2Lister{indexer: indexer}
}

// List lists all Cloud9EnvironmentEc2s in the indexer.
func (s *cloud9EnvironmentEc2Lister) List(selector labels.Selector) (ret []*v1alpha1.Cloud9EnvironmentEc2, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Cloud9EnvironmentEc2))
	})
	return ret, err
}

// Cloud9EnvironmentEc2s returns an object that can list and get Cloud9EnvironmentEc2s.
func (s *cloud9EnvironmentEc2Lister) Cloud9EnvironmentEc2s(namespace string) Cloud9EnvironmentEc2NamespaceLister {
	return cloud9EnvironmentEc2NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// Cloud9EnvironmentEc2NamespaceLister helps list and get Cloud9EnvironmentEc2s.
type Cloud9EnvironmentEc2NamespaceLister interface {
	// List lists all Cloud9EnvironmentEc2s in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Cloud9EnvironmentEc2, err error)
	// Get retrieves the Cloud9EnvironmentEc2 from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Cloud9EnvironmentEc2, error)
	Cloud9EnvironmentEc2NamespaceListerExpansion
}

// cloud9EnvironmentEc2NamespaceLister implements the Cloud9EnvironmentEc2NamespaceLister
// interface.
type cloud9EnvironmentEc2NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Cloud9EnvironmentEc2s in the indexer for a given namespace.
func (s cloud9EnvironmentEc2NamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Cloud9EnvironmentEc2, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Cloud9EnvironmentEc2))
	})
	return ret, err
}

// Get retrieves the Cloud9EnvironmentEc2 from the indexer for a given namespace and name.
func (s cloud9EnvironmentEc2NamespaceLister) Get(name string) (*v1alpha1.Cloud9EnvironmentEc2, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloud9environmentec2"), name)
	}
	return obj.(*v1alpha1.Cloud9EnvironmentEc2), nil
}
