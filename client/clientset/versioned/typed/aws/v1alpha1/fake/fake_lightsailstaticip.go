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

// FakeLightsailStaticIPs implements LightsailStaticIPInterface
type FakeLightsailStaticIPs struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var lightsailstaticipsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "lightsailstaticips"}

var lightsailstaticipsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LightsailStaticIP"}

// Get takes name of the lightsailStaticIP, and returns the corresponding lightsailStaticIP object, and an error if there is any.
func (c *FakeLightsailStaticIPs) Get(name string, options v1.GetOptions) (result *v1alpha1.LightsailStaticIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lightsailstaticipsResource, c.ns, name), &v1alpha1.LightsailStaticIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIP), err
}

// List takes label and field selectors, and returns the list of LightsailStaticIPs that match those selectors.
func (c *FakeLightsailStaticIPs) List(opts v1.ListOptions) (result *v1alpha1.LightsailStaticIPList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lightsailstaticipsResource, lightsailstaticipsKind, c.ns, opts), &v1alpha1.LightsailStaticIPList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LightsailStaticIPList{ListMeta: obj.(*v1alpha1.LightsailStaticIPList).ListMeta}
	for _, item := range obj.(*v1alpha1.LightsailStaticIPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lightsailStaticIPs.
func (c *FakeLightsailStaticIPs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lightsailstaticipsResource, c.ns, opts))

}

// Create takes the representation of a lightsailStaticIP and creates it.  Returns the server's representation of the lightsailStaticIP, and an error, if there is any.
func (c *FakeLightsailStaticIPs) Create(lightsailStaticIP *v1alpha1.LightsailStaticIP) (result *v1alpha1.LightsailStaticIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lightsailstaticipsResource, c.ns, lightsailStaticIP), &v1alpha1.LightsailStaticIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIP), err
}

// Update takes the representation of a lightsailStaticIP and updates it. Returns the server's representation of the lightsailStaticIP, and an error, if there is any.
func (c *FakeLightsailStaticIPs) Update(lightsailStaticIP *v1alpha1.LightsailStaticIP) (result *v1alpha1.LightsailStaticIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lightsailstaticipsResource, c.ns, lightsailStaticIP), &v1alpha1.LightsailStaticIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLightsailStaticIPs) UpdateStatus(lightsailStaticIP *v1alpha1.LightsailStaticIP) (*v1alpha1.LightsailStaticIP, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lightsailstaticipsResource, "status", c.ns, lightsailStaticIP), &v1alpha1.LightsailStaticIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIP), err
}

// Delete takes name of the lightsailStaticIP and deletes it. Returns an error if one occurs.
func (c *FakeLightsailStaticIPs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lightsailstaticipsResource, c.ns, name), &v1alpha1.LightsailStaticIP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLightsailStaticIPs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lightsailstaticipsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LightsailStaticIPList{})
	return err
}

// Patch applies the patch and returns the patched lightsailStaticIP.
func (c *FakeLightsailStaticIPs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LightsailStaticIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lightsailstaticipsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LightsailStaticIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LightsailStaticIP), err
}
