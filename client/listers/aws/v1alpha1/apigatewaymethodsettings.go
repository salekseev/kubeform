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

// ApiGatewayMethodSettingsLister helps list ApiGatewayMethodSettingses.
type ApiGatewayMethodSettingsLister interface {
	// List lists all ApiGatewayMethodSettingses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayMethodSettings, err error)
	// ApiGatewayMethodSettingses returns an object that can list and get ApiGatewayMethodSettingses.
	ApiGatewayMethodSettingses(namespace string) ApiGatewayMethodSettingsNamespaceLister
	ApiGatewayMethodSettingsListerExpansion
}

// apiGatewayMethodSettingsLister implements the ApiGatewayMethodSettingsLister interface.
type apiGatewayMethodSettingsLister struct {
	indexer cache.Indexer
}

// NewApiGatewayMethodSettingsLister returns a new ApiGatewayMethodSettingsLister.
func NewApiGatewayMethodSettingsLister(indexer cache.Indexer) ApiGatewayMethodSettingsLister {
	return &apiGatewayMethodSettingsLister{indexer: indexer}
}

// List lists all ApiGatewayMethodSettingses in the indexer.
func (s *apiGatewayMethodSettingsLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayMethodSettings, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayMethodSettings))
	})
	return ret, err
}

// ApiGatewayMethodSettingses returns an object that can list and get ApiGatewayMethodSettingses.
func (s *apiGatewayMethodSettingsLister) ApiGatewayMethodSettingses(namespace string) ApiGatewayMethodSettingsNamespaceLister {
	return apiGatewayMethodSettingsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiGatewayMethodSettingsNamespaceLister helps list and get ApiGatewayMethodSettingses.
type ApiGatewayMethodSettingsNamespaceLister interface {
	// List lists all ApiGatewayMethodSettingses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayMethodSettings, err error)
	// Get retrieves the ApiGatewayMethodSettings from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiGatewayMethodSettings, error)
	ApiGatewayMethodSettingsNamespaceListerExpansion
}

// apiGatewayMethodSettingsNamespaceLister implements the ApiGatewayMethodSettingsNamespaceLister
// interface.
type apiGatewayMethodSettingsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiGatewayMethodSettingses in the indexer for a given namespace.
func (s apiGatewayMethodSettingsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayMethodSettings, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayMethodSettings))
	})
	return ret, err
}

// Get retrieves the ApiGatewayMethodSettings from the indexer for a given namespace and name.
func (s apiGatewayMethodSettingsNamespaceLister) Get(name string) (*v1alpha1.ApiGatewayMethodSettings, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apigatewaymethodsettings"), name)
	}
	return obj.(*v1alpha1.ApiGatewayMethodSettings), nil
}
