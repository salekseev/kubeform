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

// LogicAppActionHTTPLister helps list LogicAppActionHTTPs.
type LogicAppActionHTTPLister interface {
	// List lists all LogicAppActionHTTPs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LogicAppActionHTTP, err error)
	// LogicAppActionHTTPs returns an object that can list and get LogicAppActionHTTPs.
	LogicAppActionHTTPs(namespace string) LogicAppActionHTTPNamespaceLister
	LogicAppActionHTTPListerExpansion
}

// logicAppActionHTTPLister implements the LogicAppActionHTTPLister interface.
type logicAppActionHTTPLister struct {
	indexer cache.Indexer
}

// NewLogicAppActionHTTPLister returns a new LogicAppActionHTTPLister.
func NewLogicAppActionHTTPLister(indexer cache.Indexer) LogicAppActionHTTPLister {
	return &logicAppActionHTTPLister{indexer: indexer}
}

// List lists all LogicAppActionHTTPs in the indexer.
func (s *logicAppActionHTTPLister) List(selector labels.Selector) (ret []*v1alpha1.LogicAppActionHTTP, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogicAppActionHTTP))
	})
	return ret, err
}

// LogicAppActionHTTPs returns an object that can list and get LogicAppActionHTTPs.
func (s *logicAppActionHTTPLister) LogicAppActionHTTPs(namespace string) LogicAppActionHTTPNamespaceLister {
	return logicAppActionHTTPNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LogicAppActionHTTPNamespaceLister helps list and get LogicAppActionHTTPs.
type LogicAppActionHTTPNamespaceLister interface {
	// List lists all LogicAppActionHTTPs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LogicAppActionHTTP, err error)
	// Get retrieves the LogicAppActionHTTP from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LogicAppActionHTTP, error)
	LogicAppActionHTTPNamespaceListerExpansion
}

// logicAppActionHTTPNamespaceLister implements the LogicAppActionHTTPNamespaceLister
// interface.
type logicAppActionHTTPNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LogicAppActionHTTPs in the indexer for a given namespace.
func (s logicAppActionHTTPNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LogicAppActionHTTP, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LogicAppActionHTTP))
	})
	return ret, err
}

// Get retrieves the LogicAppActionHTTP from the indexer for a given namespace and name.
func (s logicAppActionHTTPNamespaceLister) Get(name string) (*v1alpha1.LogicAppActionHTTP, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("logicappactionhttp"), name)
	}
	return obj.(*v1alpha1.LogicAppActionHTTP), nil
}
