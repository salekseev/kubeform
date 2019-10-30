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

// FakeIamUserSSHKeys implements IamUserSSHKeyInterface
type FakeIamUserSSHKeys struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iamusersshkeysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iamusersshkeys"}

var iamusersshkeysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamUserSSHKey"}

// Get takes name of the iamUserSSHKey, and returns the corresponding iamUserSSHKey object, and an error if there is any.
func (c *FakeIamUserSSHKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.IamUserSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamusersshkeysResource, c.ns, name), &v1alpha1.IamUserSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserSSHKey), err
}

// List takes label and field selectors, and returns the list of IamUserSSHKeys that match those selectors.
func (c *FakeIamUserSSHKeys) List(opts v1.ListOptions) (result *v1alpha1.IamUserSSHKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamusersshkeysResource, iamusersshkeysKind, c.ns, opts), &v1alpha1.IamUserSSHKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamUserSSHKeyList{ListMeta: obj.(*v1alpha1.IamUserSSHKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamUserSSHKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamUserSSHKeys.
func (c *FakeIamUserSSHKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamusersshkeysResource, c.ns, opts))

}

// Create takes the representation of a iamUserSSHKey and creates it.  Returns the server's representation of the iamUserSSHKey, and an error, if there is any.
func (c *FakeIamUserSSHKeys) Create(iamUserSSHKey *v1alpha1.IamUserSSHKey) (result *v1alpha1.IamUserSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamusersshkeysResource, c.ns, iamUserSSHKey), &v1alpha1.IamUserSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserSSHKey), err
}

// Update takes the representation of a iamUserSSHKey and updates it. Returns the server's representation of the iamUserSSHKey, and an error, if there is any.
func (c *FakeIamUserSSHKeys) Update(iamUserSSHKey *v1alpha1.IamUserSSHKey) (result *v1alpha1.IamUserSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamusersshkeysResource, c.ns, iamUserSSHKey), &v1alpha1.IamUserSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserSSHKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamUserSSHKeys) UpdateStatus(iamUserSSHKey *v1alpha1.IamUserSSHKey) (*v1alpha1.IamUserSSHKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamusersshkeysResource, "status", c.ns, iamUserSSHKey), &v1alpha1.IamUserSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserSSHKey), err
}

// Delete takes name of the iamUserSSHKey and deletes it. Returns an error if one occurs.
func (c *FakeIamUserSSHKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iamusersshkeysResource, c.ns, name), &v1alpha1.IamUserSSHKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamUserSSHKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamusersshkeysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamUserSSHKeyList{})
	return err
}

// Patch applies the patch and returns the patched iamUserSSHKey.
func (c *FakeIamUserSSHKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamUserSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamusersshkeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamUserSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamUserSSHKey), err
}
