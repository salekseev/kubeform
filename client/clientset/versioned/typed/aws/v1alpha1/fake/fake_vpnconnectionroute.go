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

// FakeVpnConnectionRoutes implements VpnConnectionRouteInterface
type FakeVpnConnectionRoutes struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var vpnconnectionroutesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "vpnconnectionroutes"}

var vpnconnectionroutesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "VpnConnectionRoute"}

// Get takes name of the vpnConnectionRoute, and returns the corresponding vpnConnectionRoute object, and an error if there is any.
func (c *FakeVpnConnectionRoutes) Get(name string, options v1.GetOptions) (result *v1alpha1.VpnConnectionRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpnconnectionroutesResource, c.ns, name), &v1alpha1.VpnConnectionRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnConnectionRoute), err
}

// List takes label and field selectors, and returns the list of VpnConnectionRoutes that match those selectors.
func (c *FakeVpnConnectionRoutes) List(opts v1.ListOptions) (result *v1alpha1.VpnConnectionRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpnconnectionroutesResource, vpnconnectionroutesKind, c.ns, opts), &v1alpha1.VpnConnectionRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpnConnectionRouteList{ListMeta: obj.(*v1alpha1.VpnConnectionRouteList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpnConnectionRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpnConnectionRoutes.
func (c *FakeVpnConnectionRoutes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpnconnectionroutesResource, c.ns, opts))

}

// Create takes the representation of a vpnConnectionRoute and creates it.  Returns the server's representation of the vpnConnectionRoute, and an error, if there is any.
func (c *FakeVpnConnectionRoutes) Create(vpnConnectionRoute *v1alpha1.VpnConnectionRoute) (result *v1alpha1.VpnConnectionRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpnconnectionroutesResource, c.ns, vpnConnectionRoute), &v1alpha1.VpnConnectionRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnConnectionRoute), err
}

// Update takes the representation of a vpnConnectionRoute and updates it. Returns the server's representation of the vpnConnectionRoute, and an error, if there is any.
func (c *FakeVpnConnectionRoutes) Update(vpnConnectionRoute *v1alpha1.VpnConnectionRoute) (result *v1alpha1.VpnConnectionRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpnconnectionroutesResource, c.ns, vpnConnectionRoute), &v1alpha1.VpnConnectionRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnConnectionRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpnConnectionRoutes) UpdateStatus(vpnConnectionRoute *v1alpha1.VpnConnectionRoute) (*v1alpha1.VpnConnectionRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpnconnectionroutesResource, "status", c.ns, vpnConnectionRoute), &v1alpha1.VpnConnectionRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnConnectionRoute), err
}

// Delete takes name of the vpnConnectionRoute and deletes it. Returns an error if one occurs.
func (c *FakeVpnConnectionRoutes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpnconnectionroutesResource, c.ns, name), &v1alpha1.VpnConnectionRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpnConnectionRoutes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpnconnectionroutesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpnConnectionRouteList{})
	return err
}

// Patch applies the patch and returns the patched vpnConnectionRoute.
func (c *FakeVpnConnectionRoutes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpnConnectionRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpnconnectionroutesResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpnConnectionRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpnConnectionRoute), err
}
