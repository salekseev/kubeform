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

// FakeVpnGatewayRoutePropagations implements VpnGatewayRoutePropagationInterface
type FakeVpnGatewayRoutePropagations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var vpngatewayroutepropagationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "vpngatewayroutepropagations"}

var vpngatewayroutepropagationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "VpnGatewayRoutePropagation"}

// Get takes name of the vpnGatewayRoutePropagation, and returns the corresponding vpnGatewayRoutePropagation object, and an error if there is any.
func (c *FakeVpnGatewayRoutePropagations) Get(name string, options v1.GetOptions) (result *v1alpha1.VpnGatewayRoutePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpngatewayroutepropagationsResource, c.ns, name), &v1alpha1.VpnGatewayRoutePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnGatewayRoutePropagation), err
}

// List takes label and field selectors, and returns the list of VpnGatewayRoutePropagations that match those selectors.
func (c *FakeVpnGatewayRoutePropagations) List(opts v1.ListOptions) (result *v1alpha1.VpnGatewayRoutePropagationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpngatewayroutepropagationsResource, vpngatewayroutepropagationsKind, c.ns, opts), &v1alpha1.VpnGatewayRoutePropagationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpnGatewayRoutePropagationList{ListMeta: obj.(*v1alpha1.VpnGatewayRoutePropagationList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpnGatewayRoutePropagationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpnGatewayRoutePropagations.
func (c *FakeVpnGatewayRoutePropagations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpngatewayroutepropagationsResource, c.ns, opts))

}

// Create takes the representation of a vpnGatewayRoutePropagation and creates it.  Returns the server's representation of the vpnGatewayRoutePropagation, and an error, if there is any.
func (c *FakeVpnGatewayRoutePropagations) Create(vpnGatewayRoutePropagation *v1alpha1.VpnGatewayRoutePropagation) (result *v1alpha1.VpnGatewayRoutePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpngatewayroutepropagationsResource, c.ns, vpnGatewayRoutePropagation), &v1alpha1.VpnGatewayRoutePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnGatewayRoutePropagation), err
}

// Update takes the representation of a vpnGatewayRoutePropagation and updates it. Returns the server's representation of the vpnGatewayRoutePropagation, and an error, if there is any.
func (c *FakeVpnGatewayRoutePropagations) Update(vpnGatewayRoutePropagation *v1alpha1.VpnGatewayRoutePropagation) (result *v1alpha1.VpnGatewayRoutePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpngatewayroutepropagationsResource, c.ns, vpnGatewayRoutePropagation), &v1alpha1.VpnGatewayRoutePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnGatewayRoutePropagation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpnGatewayRoutePropagations) UpdateStatus(vpnGatewayRoutePropagation *v1alpha1.VpnGatewayRoutePropagation) (*v1alpha1.VpnGatewayRoutePropagation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpngatewayroutepropagationsResource, "status", c.ns, vpnGatewayRoutePropagation), &v1alpha1.VpnGatewayRoutePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnGatewayRoutePropagation), err
}

// Delete takes name of the vpnGatewayRoutePropagation and deletes it. Returns an error if one occurs.
func (c *FakeVpnGatewayRoutePropagations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpngatewayroutepropagationsResource, c.ns, name), &v1alpha1.VpnGatewayRoutePropagation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpnGatewayRoutePropagations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpngatewayroutepropagationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpnGatewayRoutePropagationList{})
	return err
}

// Patch applies the patch and returns the patched vpnGatewayRoutePropagation.
func (c *FakeVpnGatewayRoutePropagations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpnGatewayRoutePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpngatewayroutepropagationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpnGatewayRoutePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnGatewayRoutePropagation), err
}
