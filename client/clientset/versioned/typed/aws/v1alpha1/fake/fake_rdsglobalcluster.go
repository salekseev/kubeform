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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRdsGlobalClusters implements RdsGlobalClusterInterface
type FakeRdsGlobalClusters struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var rdsglobalclustersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "rdsglobalclusters"}

var rdsglobalclustersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "RdsGlobalCluster"}

// Get takes name of the rdsGlobalCluster, and returns the corresponding rdsGlobalCluster object, and an error if there is any.
func (c *FakeRdsGlobalClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.RdsGlobalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rdsglobalclustersResource, c.ns, name), &v1alpha1.RdsGlobalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RdsGlobalCluster), err
}

// List takes label and field selectors, and returns the list of RdsGlobalClusters that match those selectors.
func (c *FakeRdsGlobalClusters) List(opts v1.ListOptions) (result *v1alpha1.RdsGlobalClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rdsglobalclustersResource, rdsglobalclustersKind, c.ns, opts), &v1alpha1.RdsGlobalClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RdsGlobalClusterList{ListMeta: obj.(*v1alpha1.RdsGlobalClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.RdsGlobalClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rdsGlobalClusters.
func (c *FakeRdsGlobalClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rdsglobalclustersResource, c.ns, opts))

}

// Create takes the representation of a rdsGlobalCluster and creates it.  Returns the server's representation of the rdsGlobalCluster, and an error, if there is any.
func (c *FakeRdsGlobalClusters) Create(rdsGlobalCluster *v1alpha1.RdsGlobalCluster) (result *v1alpha1.RdsGlobalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rdsglobalclustersResource, c.ns, rdsGlobalCluster), &v1alpha1.RdsGlobalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RdsGlobalCluster), err
}

// Update takes the representation of a rdsGlobalCluster and updates it. Returns the server's representation of the rdsGlobalCluster, and an error, if there is any.
func (c *FakeRdsGlobalClusters) Update(rdsGlobalCluster *v1alpha1.RdsGlobalCluster) (result *v1alpha1.RdsGlobalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rdsglobalclustersResource, c.ns, rdsGlobalCluster), &v1alpha1.RdsGlobalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RdsGlobalCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRdsGlobalClusters) UpdateStatus(rdsGlobalCluster *v1alpha1.RdsGlobalCluster) (*v1alpha1.RdsGlobalCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rdsglobalclustersResource, "status", c.ns, rdsGlobalCluster), &v1alpha1.RdsGlobalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RdsGlobalCluster), err
}

// Delete takes name of the rdsGlobalCluster and deletes it. Returns an error if one occurs.
func (c *FakeRdsGlobalClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rdsglobalclustersResource, c.ns, name), &v1alpha1.RdsGlobalCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRdsGlobalClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rdsglobalclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RdsGlobalClusterList{})
	return err
}

// Patch applies the patch and returns the patched rdsGlobalCluster.
func (c *FakeRdsGlobalClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RdsGlobalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rdsglobalclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.RdsGlobalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RdsGlobalCluster), err
}
