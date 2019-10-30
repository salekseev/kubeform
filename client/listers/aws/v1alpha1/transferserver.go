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

// TransferServerLister helps list TransferServers.
type TransferServerLister interface {
	// List lists all TransferServers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TransferServer, err error)
	// TransferServers returns an object that can list and get TransferServers.
	TransferServers(namespace string) TransferServerNamespaceLister
	TransferServerListerExpansion
}

// transferServerLister implements the TransferServerLister interface.
type transferServerLister struct {
	indexer cache.Indexer
}

// NewTransferServerLister returns a new TransferServerLister.
func NewTransferServerLister(indexer cache.Indexer) TransferServerLister {
	return &transferServerLister{indexer: indexer}
}

// List lists all TransferServers in the indexer.
func (s *transferServerLister) List(selector labels.Selector) (ret []*v1alpha1.TransferServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TransferServer))
	})
	return ret, err
}

// TransferServers returns an object that can list and get TransferServers.
func (s *transferServerLister) TransferServers(namespace string) TransferServerNamespaceLister {
	return transferServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TransferServerNamespaceLister helps list and get TransferServers.
type TransferServerNamespaceLister interface {
	// List lists all TransferServers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TransferServer, err error)
	// Get retrieves the TransferServer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TransferServer, error)
	TransferServerNamespaceListerExpansion
}

// transferServerNamespaceLister implements the TransferServerNamespaceLister
// interface.
type transferServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TransferServers in the indexer for a given namespace.
func (s transferServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TransferServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TransferServer))
	})
	return ret, err
}

// Get retrieves the TransferServer from the indexer for a given namespace and name.
func (s transferServerNamespaceLister) Get(name string) (*v1alpha1.TransferServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("transferserver"), name)
	}
	return obj.(*v1alpha1.TransferServer), nil
}
