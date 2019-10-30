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

// GlueSecurityConfigurationLister helps list GlueSecurityConfigurations.
type GlueSecurityConfigurationLister interface {
	// List lists all GlueSecurityConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GlueSecurityConfiguration, err error)
	// GlueSecurityConfigurations returns an object that can list and get GlueSecurityConfigurations.
	GlueSecurityConfigurations(namespace string) GlueSecurityConfigurationNamespaceLister
	GlueSecurityConfigurationListerExpansion
}

// glueSecurityConfigurationLister implements the GlueSecurityConfigurationLister interface.
type glueSecurityConfigurationLister struct {
	indexer cache.Indexer
}

// NewGlueSecurityConfigurationLister returns a new GlueSecurityConfigurationLister.
func NewGlueSecurityConfigurationLister(indexer cache.Indexer) GlueSecurityConfigurationLister {
	return &glueSecurityConfigurationLister{indexer: indexer}
}

// List lists all GlueSecurityConfigurations in the indexer.
func (s *glueSecurityConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.GlueSecurityConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlueSecurityConfiguration))
	})
	return ret, err
}

// GlueSecurityConfigurations returns an object that can list and get GlueSecurityConfigurations.
func (s *glueSecurityConfigurationLister) GlueSecurityConfigurations(namespace string) GlueSecurityConfigurationNamespaceLister {
	return glueSecurityConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlueSecurityConfigurationNamespaceLister helps list and get GlueSecurityConfigurations.
type GlueSecurityConfigurationNamespaceLister interface {
	// List lists all GlueSecurityConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GlueSecurityConfiguration, err error)
	// Get retrieves the GlueSecurityConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GlueSecurityConfiguration, error)
	GlueSecurityConfigurationNamespaceListerExpansion
}

// glueSecurityConfigurationNamespaceLister implements the GlueSecurityConfigurationNamespaceLister
// interface.
type glueSecurityConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlueSecurityConfigurations in the indexer for a given namespace.
func (s glueSecurityConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GlueSecurityConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlueSecurityConfiguration))
	})
	return ret, err
}

// Get retrieves the GlueSecurityConfiguration from the indexer for a given namespace and name.
func (s glueSecurityConfigurationNamespaceLister) Get(name string) (*v1alpha1.GlueSecurityConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("gluesecurityconfiguration"), name)
	}
	return obj.(*v1alpha1.GlueSecurityConfiguration), nil
}
