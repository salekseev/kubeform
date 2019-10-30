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

// FakeSsmMaintenanceWindows implements SsmMaintenanceWindowInterface
type FakeSsmMaintenanceWindows struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var ssmmaintenancewindowsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "ssmmaintenancewindows"}

var ssmmaintenancewindowsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SsmMaintenanceWindow"}

// Get takes name of the ssmMaintenanceWindow, and returns the corresponding ssmMaintenanceWindow object, and an error if there is any.
func (c *FakeSsmMaintenanceWindows) Get(name string, options v1.GetOptions) (result *v1alpha1.SsmMaintenanceWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ssmmaintenancewindowsResource, c.ns, name), &v1alpha1.SsmMaintenanceWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindow), err
}

// List takes label and field selectors, and returns the list of SsmMaintenanceWindows that match those selectors.
func (c *FakeSsmMaintenanceWindows) List(opts v1.ListOptions) (result *v1alpha1.SsmMaintenanceWindowList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ssmmaintenancewindowsResource, ssmmaintenancewindowsKind, c.ns, opts), &v1alpha1.SsmMaintenanceWindowList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SsmMaintenanceWindowList{ListMeta: obj.(*v1alpha1.SsmMaintenanceWindowList).ListMeta}
	for _, item := range obj.(*v1alpha1.SsmMaintenanceWindowList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ssmMaintenanceWindows.
func (c *FakeSsmMaintenanceWindows) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ssmmaintenancewindowsResource, c.ns, opts))

}

// Create takes the representation of a ssmMaintenanceWindow and creates it.  Returns the server's representation of the ssmMaintenanceWindow, and an error, if there is any.
func (c *FakeSsmMaintenanceWindows) Create(ssmMaintenanceWindow *v1alpha1.SsmMaintenanceWindow) (result *v1alpha1.SsmMaintenanceWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ssmmaintenancewindowsResource, c.ns, ssmMaintenanceWindow), &v1alpha1.SsmMaintenanceWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindow), err
}

// Update takes the representation of a ssmMaintenanceWindow and updates it. Returns the server's representation of the ssmMaintenanceWindow, and an error, if there is any.
func (c *FakeSsmMaintenanceWindows) Update(ssmMaintenanceWindow *v1alpha1.SsmMaintenanceWindow) (result *v1alpha1.SsmMaintenanceWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ssmmaintenancewindowsResource, c.ns, ssmMaintenanceWindow), &v1alpha1.SsmMaintenanceWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindow), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSsmMaintenanceWindows) UpdateStatus(ssmMaintenanceWindow *v1alpha1.SsmMaintenanceWindow) (*v1alpha1.SsmMaintenanceWindow, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ssmmaintenancewindowsResource, "status", c.ns, ssmMaintenanceWindow), &v1alpha1.SsmMaintenanceWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindow), err
}

// Delete takes name of the ssmMaintenanceWindow and deletes it. Returns an error if one occurs.
func (c *FakeSsmMaintenanceWindows) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ssmmaintenancewindowsResource, c.ns, name), &v1alpha1.SsmMaintenanceWindow{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSsmMaintenanceWindows) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ssmmaintenancewindowsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SsmMaintenanceWindowList{})
	return err
}

// Patch applies the patch and returns the patched ssmMaintenanceWindow.
func (c *FakeSsmMaintenanceWindows) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SsmMaintenanceWindow, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ssmmaintenancewindowsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SsmMaintenanceWindow{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SsmMaintenanceWindow), err
}
