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

// MonitorActivityLogAlertLister helps list MonitorActivityLogAlerts.
type MonitorActivityLogAlertLister interface {
	// List lists all MonitorActivityLogAlerts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MonitorActivityLogAlert, err error)
	// MonitorActivityLogAlerts returns an object that can list and get MonitorActivityLogAlerts.
	MonitorActivityLogAlerts(namespace string) MonitorActivityLogAlertNamespaceLister
	MonitorActivityLogAlertListerExpansion
}

// monitorActivityLogAlertLister implements the MonitorActivityLogAlertLister interface.
type monitorActivityLogAlertLister struct {
	indexer cache.Indexer
}

// NewMonitorActivityLogAlertLister returns a new MonitorActivityLogAlertLister.
func NewMonitorActivityLogAlertLister(indexer cache.Indexer) MonitorActivityLogAlertLister {
	return &monitorActivityLogAlertLister{indexer: indexer}
}

// List lists all MonitorActivityLogAlerts in the indexer.
func (s *monitorActivityLogAlertLister) List(selector labels.Selector) (ret []*v1alpha1.MonitorActivityLogAlert, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitorActivityLogAlert))
	})
	return ret, err
}

// MonitorActivityLogAlerts returns an object that can list and get MonitorActivityLogAlerts.
func (s *monitorActivityLogAlertLister) MonitorActivityLogAlerts(namespace string) MonitorActivityLogAlertNamespaceLister {
	return monitorActivityLogAlertNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MonitorActivityLogAlertNamespaceLister helps list and get MonitorActivityLogAlerts.
type MonitorActivityLogAlertNamespaceLister interface {
	// List lists all MonitorActivityLogAlerts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MonitorActivityLogAlert, err error)
	// Get retrieves the MonitorActivityLogAlert from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MonitorActivityLogAlert, error)
	MonitorActivityLogAlertNamespaceListerExpansion
}

// monitorActivityLogAlertNamespaceLister implements the MonitorActivityLogAlertNamespaceLister
// interface.
type monitorActivityLogAlertNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MonitorActivityLogAlerts in the indexer for a given namespace.
func (s monitorActivityLogAlertNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MonitorActivityLogAlert, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MonitorActivityLogAlert))
	})
	return ret, err
}

// Get retrieves the MonitorActivityLogAlert from the indexer for a given namespace and name.
func (s monitorActivityLogAlertNamespaceLister) Get(name string) (*v1alpha1.MonitorActivityLogAlert, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("monitoractivitylogalert"), name)
	}
	return obj.(*v1alpha1.MonitorActivityLogAlert), nil
}
