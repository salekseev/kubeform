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

// EcsClusterLister helps list EcsClusters.
type EcsClusterLister interface {
	// List lists all EcsClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EcsCluster, err error)
	// EcsClusters returns an object that can list and get EcsClusters.
	EcsClusters(namespace string) EcsClusterNamespaceLister
	EcsClusterListerExpansion
}

// ecsClusterLister implements the EcsClusterLister interface.
type ecsClusterLister struct {
	indexer cache.Indexer
}

// NewEcsClusterLister returns a new EcsClusterLister.
func NewEcsClusterLister(indexer cache.Indexer) EcsClusterLister {
	return &ecsClusterLister{indexer: indexer}
}

// List lists all EcsClusters in the indexer.
func (s *ecsClusterLister) List(selector labels.Selector) (ret []*v1alpha1.EcsCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EcsCluster))
	})
	return ret, err
}

// EcsClusters returns an object that can list and get EcsClusters.
func (s *ecsClusterLister) EcsClusters(namespace string) EcsClusterNamespaceLister {
	return ecsClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EcsClusterNamespaceLister helps list and get EcsClusters.
type EcsClusterNamespaceLister interface {
	// List lists all EcsClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EcsCluster, err error)
	// Get retrieves the EcsCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EcsCluster, error)
	EcsClusterNamespaceListerExpansion
}

// ecsClusterNamespaceLister implements the EcsClusterNamespaceLister
// interface.
type ecsClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EcsClusters in the indexer for a given namespace.
func (s ecsClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EcsCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EcsCluster))
	})
	return ret, err
}

// Get retrieves the EcsCluster from the indexer for a given namespace and name.
func (s ecsClusterNamespaceLister) Get(name string) (*v1alpha1.EcsCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ecscluster"), name)
	}
	return obj.(*v1alpha1.EcsCluster), nil
}
