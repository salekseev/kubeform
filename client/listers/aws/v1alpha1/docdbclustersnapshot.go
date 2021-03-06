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

// DocdbClusterSnapshotLister helps list DocdbClusterSnapshots.
type DocdbClusterSnapshotLister interface {
	// List lists all DocdbClusterSnapshots in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DocdbClusterSnapshot, err error)
	// DocdbClusterSnapshots returns an object that can list and get DocdbClusterSnapshots.
	DocdbClusterSnapshots(namespace string) DocdbClusterSnapshotNamespaceLister
	DocdbClusterSnapshotListerExpansion
}

// docdbClusterSnapshotLister implements the DocdbClusterSnapshotLister interface.
type docdbClusterSnapshotLister struct {
	indexer cache.Indexer
}

// NewDocdbClusterSnapshotLister returns a new DocdbClusterSnapshotLister.
func NewDocdbClusterSnapshotLister(indexer cache.Indexer) DocdbClusterSnapshotLister {
	return &docdbClusterSnapshotLister{indexer: indexer}
}

// List lists all DocdbClusterSnapshots in the indexer.
func (s *docdbClusterSnapshotLister) List(selector labels.Selector) (ret []*v1alpha1.DocdbClusterSnapshot, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DocdbClusterSnapshot))
	})
	return ret, err
}

// DocdbClusterSnapshots returns an object that can list and get DocdbClusterSnapshots.
func (s *docdbClusterSnapshotLister) DocdbClusterSnapshots(namespace string) DocdbClusterSnapshotNamespaceLister {
	return docdbClusterSnapshotNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DocdbClusterSnapshotNamespaceLister helps list and get DocdbClusterSnapshots.
type DocdbClusterSnapshotNamespaceLister interface {
	// List lists all DocdbClusterSnapshots in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DocdbClusterSnapshot, err error)
	// Get retrieves the DocdbClusterSnapshot from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DocdbClusterSnapshot, error)
	DocdbClusterSnapshotNamespaceListerExpansion
}

// docdbClusterSnapshotNamespaceLister implements the DocdbClusterSnapshotNamespaceLister
// interface.
type docdbClusterSnapshotNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DocdbClusterSnapshots in the indexer for a given namespace.
func (s docdbClusterSnapshotNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DocdbClusterSnapshot, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DocdbClusterSnapshot))
	})
	return ret, err
}

// Get retrieves the DocdbClusterSnapshot from the indexer for a given namespace and name.
func (s docdbClusterSnapshotNamespaceLister) Get(name string) (*v1alpha1.DocdbClusterSnapshot, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("docdbclustersnapshot"), name)
	}
	return obj.(*v1alpha1.DocdbClusterSnapshot), nil
}
