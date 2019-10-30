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

// IamUserLoginProfileLister helps list IamUserLoginProfiles.
type IamUserLoginProfileLister interface {
	// List lists all IamUserLoginProfiles in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamUserLoginProfile, err error)
	// IamUserLoginProfiles returns an object that can list and get IamUserLoginProfiles.
	IamUserLoginProfiles(namespace string) IamUserLoginProfileNamespaceLister
	IamUserLoginProfileListerExpansion
}

// iamUserLoginProfileLister implements the IamUserLoginProfileLister interface.
type iamUserLoginProfileLister struct {
	indexer cache.Indexer
}

// NewIamUserLoginProfileLister returns a new IamUserLoginProfileLister.
func NewIamUserLoginProfileLister(indexer cache.Indexer) IamUserLoginProfileLister {
	return &iamUserLoginProfileLister{indexer: indexer}
}

// List lists all IamUserLoginProfiles in the indexer.
func (s *iamUserLoginProfileLister) List(selector labels.Selector) (ret []*v1alpha1.IamUserLoginProfile, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUserLoginProfile))
	})
	return ret, err
}

// IamUserLoginProfiles returns an object that can list and get IamUserLoginProfiles.
func (s *iamUserLoginProfileLister) IamUserLoginProfiles(namespace string) IamUserLoginProfileNamespaceLister {
	return iamUserLoginProfileNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamUserLoginProfileNamespaceLister helps list and get IamUserLoginProfiles.
type IamUserLoginProfileNamespaceLister interface {
	// List lists all IamUserLoginProfiles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamUserLoginProfile, err error)
	// Get retrieves the IamUserLoginProfile from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamUserLoginProfile, error)
	IamUserLoginProfileNamespaceListerExpansion
}

// iamUserLoginProfileNamespaceLister implements the IamUserLoginProfileNamespaceLister
// interface.
type iamUserLoginProfileNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamUserLoginProfiles in the indexer for a given namespace.
func (s iamUserLoginProfileNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamUserLoginProfile, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUserLoginProfile))
	})
	return ret, err
}

// Get retrieves the IamUserLoginProfile from the indexer for a given namespace and name.
func (s iamUserLoginProfileNamespaceLister) Get(name string) (*v1alpha1.IamUserLoginProfile, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamuserloginprofile"), name)
	}
	return obj.(*v1alpha1.IamUserLoginProfile), nil
}
