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

// AlbLister helps list Albs.
type AlbLister interface {
	// List lists all Albs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Alb, err error)
	// Albs returns an object that can list and get Albs.
	Albs(namespace string) AlbNamespaceLister
	AlbListerExpansion
}

// albLister implements the AlbLister interface.
type albLister struct {
	indexer cache.Indexer
}

// NewAlbLister returns a new AlbLister.
func NewAlbLister(indexer cache.Indexer) AlbLister {
	return &albLister{indexer: indexer}
}

// List lists all Albs in the indexer.
func (s *albLister) List(selector labels.Selector) (ret []*v1alpha1.Alb, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Alb))
	})
	return ret, err
}

// Albs returns an object that can list and get Albs.
func (s *albLister) Albs(namespace string) AlbNamespaceLister {
	return albNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlbNamespaceLister helps list and get Albs.
type AlbNamespaceLister interface {
	// List lists all Albs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Alb, err error)
	// Get retrieves the Alb from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Alb, error)
	AlbNamespaceListerExpansion
}

// albNamespaceLister implements the AlbNamespaceLister
// interface.
type albNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Albs in the indexer for a given namespace.
func (s albNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Alb, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Alb))
	})
	return ret, err
}

// Get retrieves the Alb from the indexer for a given namespace and name.
func (s albNamespaceLister) Get(name string) (*v1alpha1.Alb, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alb"), name)
	}
	return obj.(*v1alpha1.Alb), nil
}
