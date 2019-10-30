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

// EgressOnlyInternetGatewayLister helps list EgressOnlyInternetGateways.
type EgressOnlyInternetGatewayLister interface {
	// List lists all EgressOnlyInternetGateways in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EgressOnlyInternetGateway, err error)
	// EgressOnlyInternetGateways returns an object that can list and get EgressOnlyInternetGateways.
	EgressOnlyInternetGateways(namespace string) EgressOnlyInternetGatewayNamespaceLister
	EgressOnlyInternetGatewayListerExpansion
}

// egressOnlyInternetGatewayLister implements the EgressOnlyInternetGatewayLister interface.
type egressOnlyInternetGatewayLister struct {
	indexer cache.Indexer
}

// NewEgressOnlyInternetGatewayLister returns a new EgressOnlyInternetGatewayLister.
func NewEgressOnlyInternetGatewayLister(indexer cache.Indexer) EgressOnlyInternetGatewayLister {
	return &egressOnlyInternetGatewayLister{indexer: indexer}
}

// List lists all EgressOnlyInternetGateways in the indexer.
func (s *egressOnlyInternetGatewayLister) List(selector labels.Selector) (ret []*v1alpha1.EgressOnlyInternetGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EgressOnlyInternetGateway))
	})
	return ret, err
}

// EgressOnlyInternetGateways returns an object that can list and get EgressOnlyInternetGateways.
func (s *egressOnlyInternetGatewayLister) EgressOnlyInternetGateways(namespace string) EgressOnlyInternetGatewayNamespaceLister {
	return egressOnlyInternetGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EgressOnlyInternetGatewayNamespaceLister helps list and get EgressOnlyInternetGateways.
type EgressOnlyInternetGatewayNamespaceLister interface {
	// List lists all EgressOnlyInternetGateways in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EgressOnlyInternetGateway, err error)
	// Get retrieves the EgressOnlyInternetGateway from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EgressOnlyInternetGateway, error)
	EgressOnlyInternetGatewayNamespaceListerExpansion
}

// egressOnlyInternetGatewayNamespaceLister implements the EgressOnlyInternetGatewayNamespaceLister
// interface.
type egressOnlyInternetGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EgressOnlyInternetGateways in the indexer for a given namespace.
func (s egressOnlyInternetGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EgressOnlyInternetGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EgressOnlyInternetGateway))
	})
	return ret, err
}

// Get retrieves the EgressOnlyInternetGateway from the indexer for a given namespace and name.
func (s egressOnlyInternetGatewayNamespaceLister) Get(name string) (*v1alpha1.EgressOnlyInternetGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("egressonlyinternetgateway"), name)
	}
	return obj.(*v1alpha1.EgressOnlyInternetGateway), nil
}
