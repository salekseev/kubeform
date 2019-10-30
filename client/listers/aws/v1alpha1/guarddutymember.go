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

// GuarddutyMemberLister helps list GuarddutyMembers.
type GuarddutyMemberLister interface {
	// List lists all GuarddutyMembers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GuarddutyMember, err error)
	// GuarddutyMembers returns an object that can list and get GuarddutyMembers.
	GuarddutyMembers(namespace string) GuarddutyMemberNamespaceLister
	GuarddutyMemberListerExpansion
}

// guarddutyMemberLister implements the GuarddutyMemberLister interface.
type guarddutyMemberLister struct {
	indexer cache.Indexer
}

// NewGuarddutyMemberLister returns a new GuarddutyMemberLister.
func NewGuarddutyMemberLister(indexer cache.Indexer) GuarddutyMemberLister {
	return &guarddutyMemberLister{indexer: indexer}
}

// List lists all GuarddutyMembers in the indexer.
func (s *guarddutyMemberLister) List(selector labels.Selector) (ret []*v1alpha1.GuarddutyMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GuarddutyMember))
	})
	return ret, err
}

// GuarddutyMembers returns an object that can list and get GuarddutyMembers.
func (s *guarddutyMemberLister) GuarddutyMembers(namespace string) GuarddutyMemberNamespaceLister {
	return guarddutyMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GuarddutyMemberNamespaceLister helps list and get GuarddutyMembers.
type GuarddutyMemberNamespaceLister interface {
	// List lists all GuarddutyMembers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GuarddutyMember, err error)
	// Get retrieves the GuarddutyMember from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GuarddutyMember, error)
	GuarddutyMemberNamespaceListerExpansion
}

// guarddutyMemberNamespaceLister implements the GuarddutyMemberNamespaceLister
// interface.
type guarddutyMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GuarddutyMembers in the indexer for a given namespace.
func (s guarddutyMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GuarddutyMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GuarddutyMember))
	})
	return ret, err
}

// Get retrieves the GuarddutyMember from the indexer for a given namespace and name.
func (s guarddutyMemberNamespaceLister) Get(name string) (*v1alpha1.GuarddutyMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("guarddutymember"), name)
	}
	return obj.(*v1alpha1.GuarddutyMember), nil
}
