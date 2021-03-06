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

// ElasticacheSecurityGroupLister helps list ElasticacheSecurityGroups.
type ElasticacheSecurityGroupLister interface {
	// List lists all ElasticacheSecurityGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticacheSecurityGroup, err error)
	// ElasticacheSecurityGroups returns an object that can list and get ElasticacheSecurityGroups.
	ElasticacheSecurityGroups(namespace string) ElasticacheSecurityGroupNamespaceLister
	ElasticacheSecurityGroupListerExpansion
}

// elasticacheSecurityGroupLister implements the ElasticacheSecurityGroupLister interface.
type elasticacheSecurityGroupLister struct {
	indexer cache.Indexer
}

// NewElasticacheSecurityGroupLister returns a new ElasticacheSecurityGroupLister.
func NewElasticacheSecurityGroupLister(indexer cache.Indexer) ElasticacheSecurityGroupLister {
	return &elasticacheSecurityGroupLister{indexer: indexer}
}

// List lists all ElasticacheSecurityGroups in the indexer.
func (s *elasticacheSecurityGroupLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticacheSecurityGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticacheSecurityGroup))
	})
	return ret, err
}

// ElasticacheSecurityGroups returns an object that can list and get ElasticacheSecurityGroups.
func (s *elasticacheSecurityGroupLister) ElasticacheSecurityGroups(namespace string) ElasticacheSecurityGroupNamespaceLister {
	return elasticacheSecurityGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ElasticacheSecurityGroupNamespaceLister helps list and get ElasticacheSecurityGroups.
type ElasticacheSecurityGroupNamespaceLister interface {
	// List lists all ElasticacheSecurityGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticacheSecurityGroup, err error)
	// Get retrieves the ElasticacheSecurityGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ElasticacheSecurityGroup, error)
	ElasticacheSecurityGroupNamespaceListerExpansion
}

// elasticacheSecurityGroupNamespaceLister implements the ElasticacheSecurityGroupNamespaceLister
// interface.
type elasticacheSecurityGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ElasticacheSecurityGroups in the indexer for a given namespace.
func (s elasticacheSecurityGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticacheSecurityGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticacheSecurityGroup))
	})
	return ret, err
}

// Get retrieves the ElasticacheSecurityGroup from the indexer for a given namespace and name.
func (s elasticacheSecurityGroupNamespaceLister) Get(name string) (*v1alpha1.ElasticacheSecurityGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("elasticachesecuritygroup"), name)
	}
	return obj.(*v1alpha1.ElasticacheSecurityGroup), nil
}
