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

// IamUserGroupMembershipLister helps list IamUserGroupMemberships.
type IamUserGroupMembershipLister interface {
	// List lists all IamUserGroupMemberships in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamUserGroupMembership, err error)
	// IamUserGroupMemberships returns an object that can list and get IamUserGroupMemberships.
	IamUserGroupMemberships(namespace string) IamUserGroupMembershipNamespaceLister
	IamUserGroupMembershipListerExpansion
}

// iamUserGroupMembershipLister implements the IamUserGroupMembershipLister interface.
type iamUserGroupMembershipLister struct {
	indexer cache.Indexer
}

// NewIamUserGroupMembershipLister returns a new IamUserGroupMembershipLister.
func NewIamUserGroupMembershipLister(indexer cache.Indexer) IamUserGroupMembershipLister {
	return &iamUserGroupMembershipLister{indexer: indexer}
}

// List lists all IamUserGroupMemberships in the indexer.
func (s *iamUserGroupMembershipLister) List(selector labels.Selector) (ret []*v1alpha1.IamUserGroupMembership, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUserGroupMembership))
	})
	return ret, err
}

// IamUserGroupMemberships returns an object that can list and get IamUserGroupMemberships.
func (s *iamUserGroupMembershipLister) IamUserGroupMemberships(namespace string) IamUserGroupMembershipNamespaceLister {
	return iamUserGroupMembershipNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamUserGroupMembershipNamespaceLister helps list and get IamUserGroupMemberships.
type IamUserGroupMembershipNamespaceLister interface {
	// List lists all IamUserGroupMemberships in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamUserGroupMembership, err error)
	// Get retrieves the IamUserGroupMembership from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamUserGroupMembership, error)
	IamUserGroupMembershipNamespaceListerExpansion
}

// iamUserGroupMembershipNamespaceLister implements the IamUserGroupMembershipNamespaceLister
// interface.
type iamUserGroupMembershipNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamUserGroupMemberships in the indexer for a given namespace.
func (s iamUserGroupMembershipNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamUserGroupMembership, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUserGroupMembership))
	})
	return ret, err
}

// Get retrieves the IamUserGroupMembership from the indexer for a given namespace and name.
func (s iamUserGroupMembershipNamespaceLister) Get(name string) (*v1alpha1.IamUserGroupMembership, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamusergroupmembership"), name)
	}
	return obj.(*v1alpha1.IamUserGroupMembership), nil
}
