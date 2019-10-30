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

// ElasticacheSubnetGroupLister helps list ElasticacheSubnetGroups.
type ElasticacheSubnetGroupLister interface {
	// List lists all ElasticacheSubnetGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticacheSubnetGroup, err error)
	// ElasticacheSubnetGroups returns an object that can list and get ElasticacheSubnetGroups.
	ElasticacheSubnetGroups(namespace string) ElasticacheSubnetGroupNamespaceLister
	ElasticacheSubnetGroupListerExpansion
}

// elasticacheSubnetGroupLister implements the ElasticacheSubnetGroupLister interface.
type elasticacheSubnetGroupLister struct {
	indexer cache.Indexer
}

// NewElasticacheSubnetGroupLister returns a new ElasticacheSubnetGroupLister.
func NewElasticacheSubnetGroupLister(indexer cache.Indexer) ElasticacheSubnetGroupLister {
	return &elasticacheSubnetGroupLister{indexer: indexer}
}

// List lists all ElasticacheSubnetGroups in the indexer.
func (s *elasticacheSubnetGroupLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticacheSubnetGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticacheSubnetGroup))
	})
	return ret, err
}

// ElasticacheSubnetGroups returns an object that can list and get ElasticacheSubnetGroups.
func (s *elasticacheSubnetGroupLister) ElasticacheSubnetGroups(namespace string) ElasticacheSubnetGroupNamespaceLister {
	return elasticacheSubnetGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ElasticacheSubnetGroupNamespaceLister helps list and get ElasticacheSubnetGroups.
type ElasticacheSubnetGroupNamespaceLister interface {
	// List lists all ElasticacheSubnetGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticacheSubnetGroup, err error)
	// Get retrieves the ElasticacheSubnetGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ElasticacheSubnetGroup, error)
	ElasticacheSubnetGroupNamespaceListerExpansion
}

// elasticacheSubnetGroupNamespaceLister implements the ElasticacheSubnetGroupNamespaceLister
// interface.
type elasticacheSubnetGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ElasticacheSubnetGroups in the indexer for a given namespace.
func (s elasticacheSubnetGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticacheSubnetGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticacheSubnetGroup))
	})
	return ret, err
}

// Get retrieves the ElasticacheSubnetGroup from the indexer for a given namespace and name.
func (s elasticacheSubnetGroupNamespaceLister) Get(name string) (*v1alpha1.ElasticacheSubnetGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("elasticachesubnetgroup"), name)
	}
	return obj.(*v1alpha1.ElasticacheSubnetGroup), nil
}
