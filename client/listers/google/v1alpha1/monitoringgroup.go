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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MonitoringGroupLister helps list MonitoringGroups.
type MonitoringGroupLister interface {
	// List lists all MonitoringGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MonitoringGroup, err error)
	// MonitoringGroups returns an object that can list and get MonitoringGroups.
	MonitoringGroups(namespace string) MonitoringGroupNamespaceLister
	MonitoringGroupListerExpansion
}

// monitoringGroupLister implements the MonitoringGroupLister interface.
type monitoringGroupLister struct {
	indexer cache.Indexer
}

// NewMonitoringGroupLister returns a new MonitoringGroupLister.
func NewMonitoringGroupLister(indexer cache.Indexer) MonitoringGroupLister {
	return &monitoringGroupLister{indexer: indexer}
}

// List lists all MonitoringGroups in the indexer.
func (s *monitoringGroupLister) List(selector labels.Selector) (ret []*v1alpha1.MonitoringGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitoringGroup))
	})
	return ret, err
}

// MonitoringGroups returns an object that can list and get MonitoringGroups.
func (s *monitoringGroupLister) MonitoringGroups(namespace string) MonitoringGroupNamespaceLister {
	return monitoringGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MonitoringGroupNamespaceLister helps list and get MonitoringGroups.
type MonitoringGroupNamespaceLister interface {
	// List lists all MonitoringGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MonitoringGroup, err error)
	// Get retrieves the MonitoringGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MonitoringGroup, error)
	MonitoringGroupNamespaceListerExpansion
}

// monitoringGroupNamespaceLister implements the MonitoringGroupNamespaceLister
// interface.
type monitoringGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MonitoringGroups in the indexer for a given namespace.
func (s monitoringGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MonitoringGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitoringGroup))
	})
	return ret, err
}

// Get retrieves the MonitoringGroup from the indexer for a given namespace and name.
func (s monitoringGroupNamespaceLister) Get(name string) (*v1alpha1.MonitoringGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("monitoringgroup"), name)
	}
	return obj.(*v1alpha1.MonitoringGroup), nil
}
