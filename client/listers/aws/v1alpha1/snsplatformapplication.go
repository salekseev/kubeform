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

// SnsPlatformApplicationLister helps list SnsPlatformApplications.
type SnsPlatformApplicationLister interface {
	// List lists all SnsPlatformApplications in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SnsPlatformApplication, err error)
	// SnsPlatformApplications returns an object that can list and get SnsPlatformApplications.
	SnsPlatformApplications(namespace string) SnsPlatformApplicationNamespaceLister
	SnsPlatformApplicationListerExpansion
}

// snsPlatformApplicationLister implements the SnsPlatformApplicationLister interface.
type snsPlatformApplicationLister struct {
	indexer cache.Indexer
}

// NewSnsPlatformApplicationLister returns a new SnsPlatformApplicationLister.
func NewSnsPlatformApplicationLister(indexer cache.Indexer) SnsPlatformApplicationLister {
	return &snsPlatformApplicationLister{indexer: indexer}
}

// List lists all SnsPlatformApplications in the indexer.
func (s *snsPlatformApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.SnsPlatformApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SnsPlatformApplication))
	})
	return ret, err
}

// SnsPlatformApplications returns an object that can list and get SnsPlatformApplications.
func (s *snsPlatformApplicationLister) SnsPlatformApplications(namespace string) SnsPlatformApplicationNamespaceLister {
	return snsPlatformApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SnsPlatformApplicationNamespaceLister helps list and get SnsPlatformApplications.
type SnsPlatformApplicationNamespaceLister interface {
	// List lists all SnsPlatformApplications in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SnsPlatformApplication, err error)
	// Get retrieves the SnsPlatformApplication from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SnsPlatformApplication, error)
	SnsPlatformApplicationNamespaceListerExpansion
}

// snsPlatformApplicationNamespaceLister implements the SnsPlatformApplicationNamespaceLister
// interface.
type snsPlatformApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SnsPlatformApplications in the indexer for a given namespace.
func (s snsPlatformApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SnsPlatformApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SnsPlatformApplication))
	})
	return ret, err
}

// Get retrieves the SnsPlatformApplication from the indexer for a given namespace and name.
func (s snsPlatformApplicationNamespaceLister) Get(name string) (*v1alpha1.SnsPlatformApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("snsplatformapplication"), name)
	}
	return obj.(*v1alpha1.SnsPlatformApplication), nil
}
