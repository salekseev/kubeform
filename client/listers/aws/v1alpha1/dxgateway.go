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

// DxGatewayLister helps list DxGateways.
type DxGatewayLister interface {
	// List lists all DxGateways in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DxGateway, err error)
	// DxGateways returns an object that can list and get DxGateways.
	DxGateways(namespace string) DxGatewayNamespaceLister
	DxGatewayListerExpansion
}

// dxGatewayLister implements the DxGatewayLister interface.
type dxGatewayLister struct {
	indexer cache.Indexer
}

// NewDxGatewayLister returns a new DxGatewayLister.
func NewDxGatewayLister(indexer cache.Indexer) DxGatewayLister {
	return &dxGatewayLister{indexer: indexer}
}

// List lists all DxGateways in the indexer.
func (s *dxGatewayLister) List(selector labels.Selector) (ret []*v1alpha1.DxGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxGateway))
	})
	return ret, err
}

// DxGateways returns an object that can list and get DxGateways.
func (s *dxGatewayLister) DxGateways(namespace string) DxGatewayNamespaceLister {
	return dxGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DxGatewayNamespaceLister helps list and get DxGateways.
type DxGatewayNamespaceLister interface {
	// List lists all DxGateways in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DxGateway, err error)
	// Get retrieves the DxGateway from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DxGateway, error)
	DxGatewayNamespaceListerExpansion
}

// dxGatewayNamespaceLister implements the DxGatewayNamespaceLister
// interface.
type dxGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DxGateways in the indexer for a given namespace.
func (s dxGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DxGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DxGateway))
	})
	return ret, err
}

// Get retrieves the DxGateway from the indexer for a given namespace and name.
func (s dxGatewayNamespaceLister) Get(name string) (*v1alpha1.DxGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dxgateway"), name)
	}
	return obj.(*v1alpha1.DxGateway), nil
}
