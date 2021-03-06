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

// AutoscalingGroupLister helps list AutoscalingGroups.
type AutoscalingGroupLister interface {
	// List lists all AutoscalingGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AutoscalingGroup, err error)
	// AutoscalingGroups returns an object that can list and get AutoscalingGroups.
	AutoscalingGroups(namespace string) AutoscalingGroupNamespaceLister
	AutoscalingGroupListerExpansion
}

// autoscalingGroupLister implements the AutoscalingGroupLister interface.
type autoscalingGroupLister struct {
	indexer cache.Indexer
}

// NewAutoscalingGroupLister returns a new AutoscalingGroupLister.
func NewAutoscalingGroupLister(indexer cache.Indexer) AutoscalingGroupLister {
	return &autoscalingGroupLister{indexer: indexer}
}

// List lists all AutoscalingGroups in the indexer.
func (s *autoscalingGroupLister) List(selector labels.Selector) (ret []*v1alpha1.AutoscalingGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutoscalingGroup))
	})
	return ret, err
}

// AutoscalingGroups returns an object that can list and get AutoscalingGroups.
func (s *autoscalingGroupLister) AutoscalingGroups(namespace string) AutoscalingGroupNamespaceLister {
	return autoscalingGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AutoscalingGroupNamespaceLister helps list and get AutoscalingGroups.
type AutoscalingGroupNamespaceLister interface {
	// List lists all AutoscalingGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AutoscalingGroup, err error)
	// Get retrieves the AutoscalingGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AutoscalingGroup, error)
	AutoscalingGroupNamespaceListerExpansion
}

// autoscalingGroupNamespaceLister implements the AutoscalingGroupNamespaceLister
// interface.
type autoscalingGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AutoscalingGroups in the indexer for a given namespace.
func (s autoscalingGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AutoscalingGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutoscalingGroup))
	})
	return ret, err
}

// Get retrieves the AutoscalingGroup from the indexer for a given namespace and name.
func (s autoscalingGroupNamespaceLister) Get(name string) (*v1alpha1.AutoscalingGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("autoscalinggroup"), name)
	}
	return obj.(*v1alpha1.AutoscalingGroup), nil
}
