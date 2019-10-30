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

// FakeAmiCopies implements AmiCopyInterface
type FakeAmiCopies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var amicopiesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "amicopies"}

var amicopiesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AmiCopy"}

// Get takes name of the amiCopy, and returns the corresponding amiCopy object, and an error if there is any.
func (c *FakeAmiCopies) Get(name string, options v1.GetOptions) (result *v1alpha1.AmiCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(amicopiesResource, c.ns, name), &v1alpha1.AmiCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AmiCopy), err
}

// List takes label and field selectors, and returns the list of AmiCopies that match those selectors.
func (c *FakeAmiCopies) List(opts v1.ListOptions) (result *v1alpha1.AmiCopyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(amicopiesResource, amicopiesKind, c.ns, opts), &v1alpha1.AmiCopyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AmiCopyList{ListMeta: obj.(*v1alpha1.AmiCopyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AmiCopyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested amiCopies.
func (c *FakeAmiCopies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(amicopiesResource, c.ns, opts))

}

// Create takes the representation of a amiCopy and creates it.  Returns the server's representation of the amiCopy, and an error, if there is any.
func (c *FakeAmiCopies) Create(amiCopy *v1alpha1.AmiCopy) (result *v1alpha1.AmiCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(amicopiesResource, c.ns, amiCopy), &v1alpha1.AmiCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AmiCopy), err
}

// Update takes the representation of a amiCopy and updates it. Returns the server's representation of the amiCopy, and an error, if there is any.
func (c *FakeAmiCopies) Update(amiCopy *v1alpha1.AmiCopy) (result *v1alpha1.AmiCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(amicopiesResource, c.ns, amiCopy), &v1alpha1.AmiCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AmiCopy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAmiCopies) UpdateStatus(amiCopy *v1alpha1.AmiCopy) (*v1alpha1.AmiCopy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(amicopiesResource, "status", c.ns, amiCopy), &v1alpha1.AmiCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AmiCopy), err
}

// Delete takes name of the amiCopy and deletes it. Returns an error if one occurs.
func (c *FakeAmiCopies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(amicopiesResource, c.ns, name), &v1alpha1.AmiCopy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAmiCopies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(amicopiesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AmiCopyList{})
	return err
}

// Patch applies the patch and returns the patched amiCopy.
func (c *FakeAmiCopies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AmiCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(amicopiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AmiCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AmiCopy), err
}
