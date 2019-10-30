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

// FakeOpsworksHaproxyLayers implements OpsworksHaproxyLayerInterface
type FakeOpsworksHaproxyLayers struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var opsworkshaproxylayersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "opsworkshaproxylayers"}

var opsworkshaproxylayersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OpsworksHaproxyLayer"}

// Get takes name of the opsworksHaproxyLayer, and returns the corresponding opsworksHaproxyLayer object, and an error if there is any.
func (c *FakeOpsworksHaproxyLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksHaproxyLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(opsworkshaproxylayersResource, c.ns, name), &v1alpha1.OpsworksHaproxyLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksHaproxyLayer), err
}

// List takes label and field selectors, and returns the list of OpsworksHaproxyLayers that match those selectors.
func (c *FakeOpsworksHaproxyLayers) List(opts v1.ListOptions) (result *v1alpha1.OpsworksHaproxyLayerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(opsworkshaproxylayersResource, opsworkshaproxylayersKind, c.ns, opts), &v1alpha1.OpsworksHaproxyLayerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpsworksHaproxyLayerList{ListMeta: obj.(*v1alpha1.OpsworksHaproxyLayerList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpsworksHaproxyLayerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested opsworksHaproxyLayers.
func (c *FakeOpsworksHaproxyLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(opsworkshaproxylayersResource, c.ns, opts))

}

// Create takes the representation of a opsworksHaproxyLayer and creates it.  Returns the server's representation of the opsworksHaproxyLayer, and an error, if there is any.
func (c *FakeOpsworksHaproxyLayers) Create(opsworksHaproxyLayer *v1alpha1.OpsworksHaproxyLayer) (result *v1alpha1.OpsworksHaproxyLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(opsworkshaproxylayersResource, c.ns, opsworksHaproxyLayer), &v1alpha1.OpsworksHaproxyLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksHaproxyLayer), err
}

// Update takes the representation of a opsworksHaproxyLayer and updates it. Returns the server's representation of the opsworksHaproxyLayer, and an error, if there is any.
func (c *FakeOpsworksHaproxyLayers) Update(opsworksHaproxyLayer *v1alpha1.OpsworksHaproxyLayer) (result *v1alpha1.OpsworksHaproxyLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(opsworkshaproxylayersResource, c.ns, opsworksHaproxyLayer), &v1alpha1.OpsworksHaproxyLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksHaproxyLayer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpsworksHaproxyLayers) UpdateStatus(opsworksHaproxyLayer *v1alpha1.OpsworksHaproxyLayer) (*v1alpha1.OpsworksHaproxyLayer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(opsworkshaproxylayersResource, "status", c.ns, opsworksHaproxyLayer), &v1alpha1.OpsworksHaproxyLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksHaproxyLayer), err
}

// Delete takes name of the opsworksHaproxyLayer and deletes it. Returns an error if one occurs.
func (c *FakeOpsworksHaproxyLayers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(opsworkshaproxylayersResource, c.ns, name), &v1alpha1.OpsworksHaproxyLayer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpsworksHaproxyLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(opsworkshaproxylayersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpsworksHaproxyLayerList{})
	return err
}

// Patch applies the patch and returns the patched opsworksHaproxyLayer.
func (c *FakeOpsworksHaproxyLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksHaproxyLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(opsworkshaproxylayersResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpsworksHaproxyLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksHaproxyLayer), err
}
