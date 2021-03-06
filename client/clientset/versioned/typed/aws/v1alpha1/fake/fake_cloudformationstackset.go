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

// FakeCloudformationStackSets implements CloudformationStackSetInterface
type FakeCloudformationStackSets struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var cloudformationstacksetsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "cloudformationstacksets"}

var cloudformationstacksetsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "CloudformationStackSet"}

// Get takes name of the cloudformationStackSet, and returns the corresponding cloudformationStackSet object, and an error if there is any.
func (c *FakeCloudformationStackSets) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudformationStackSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudformationstacksetsResource, c.ns, name), &v1alpha1.CloudformationStackSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudformationStackSet), err
}

// List takes label and field selectors, and returns the list of CloudformationStackSets that match those selectors.
func (c *FakeCloudformationStackSets) List(opts v1.ListOptions) (result *v1alpha1.CloudformationStackSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudformationstacksetsResource, cloudformationstacksetsKind, c.ns, opts), &v1alpha1.CloudformationStackSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudformationStackSetList{ListMeta: obj.(*v1alpha1.CloudformationStackSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudformationStackSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudformationStackSets.
func (c *FakeCloudformationStackSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudformationstacksetsResource, c.ns, opts))

}

// Create takes the representation of a cloudformationStackSet and creates it.  Returns the server's representation of the cloudformationStackSet, and an error, if there is any.
func (c *FakeCloudformationStackSets) Create(cloudformationStackSet *v1alpha1.CloudformationStackSet) (result *v1alpha1.CloudformationStackSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudformationstacksetsResource, c.ns, cloudformationStackSet), &v1alpha1.CloudformationStackSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudformationStackSet), err
}

// Update takes the representation of a cloudformationStackSet and updates it. Returns the server's representation of the cloudformationStackSet, and an error, if there is any.
func (c *FakeCloudformationStackSets) Update(cloudformationStackSet *v1alpha1.CloudformationStackSet) (result *v1alpha1.CloudformationStackSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudformationstacksetsResource, c.ns, cloudformationStackSet), &v1alpha1.CloudformationStackSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudformationStackSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudformationStackSets) UpdateStatus(cloudformationStackSet *v1alpha1.CloudformationStackSet) (*v1alpha1.CloudformationStackSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudformationstacksetsResource, "status", c.ns, cloudformationStackSet), &v1alpha1.CloudformationStackSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudformationStackSet), err
}

// Delete takes name of the cloudformationStackSet and deletes it. Returns an error if one occurs.
func (c *FakeCloudformationStackSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudformationstacksetsResource, c.ns, name), &v1alpha1.CloudformationStackSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudformationStackSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudformationstacksetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudformationStackSetList{})
	return err
}

// Patch applies the patch and returns the patched cloudformationStackSet.
func (c *FakeCloudformationStackSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudformationStackSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudformationstacksetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudformationStackSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudformationStackSet), err
}
