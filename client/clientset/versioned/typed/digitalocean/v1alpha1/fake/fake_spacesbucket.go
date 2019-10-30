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
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSpacesBuckets implements SpacesBucketInterface
type FakeSpacesBuckets struct {
	Fake *FakeDigitaloceanV1alpha1
	ns   string
}

var spacesbucketsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "spacesbuckets"}

var spacesbucketsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "SpacesBucket"}

// Get takes name of the spacesBucket, and returns the corresponding spacesBucket object, and an error if there is any.
func (c *FakeSpacesBuckets) Get(name string, options v1.GetOptions) (result *v1alpha1.SpacesBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(spacesbucketsResource, c.ns, name), &v1alpha1.SpacesBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpacesBucket), err
}

// List takes label and field selectors, and returns the list of SpacesBuckets that match those selectors.
func (c *FakeSpacesBuckets) List(opts v1.ListOptions) (result *v1alpha1.SpacesBucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(spacesbucketsResource, spacesbucketsKind, c.ns, opts), &v1alpha1.SpacesBucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SpacesBucketList{ListMeta: obj.(*v1alpha1.SpacesBucketList).ListMeta}
	for _, item := range obj.(*v1alpha1.SpacesBucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spacesBuckets.
func (c *FakeSpacesBuckets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(spacesbucketsResource, c.ns, opts))

}

// Create takes the representation of a spacesBucket and creates it.  Returns the server's representation of the spacesBucket, and an error, if there is any.
func (c *FakeSpacesBuckets) Create(spacesBucket *v1alpha1.SpacesBucket) (result *v1alpha1.SpacesBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(spacesbucketsResource, c.ns, spacesBucket), &v1alpha1.SpacesBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpacesBucket), err
}

// Update takes the representation of a spacesBucket and updates it. Returns the server's representation of the spacesBucket, and an error, if there is any.
func (c *FakeSpacesBuckets) Update(spacesBucket *v1alpha1.SpacesBucket) (result *v1alpha1.SpacesBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(spacesbucketsResource, c.ns, spacesBucket), &v1alpha1.SpacesBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpacesBucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpacesBuckets) UpdateStatus(spacesBucket *v1alpha1.SpacesBucket) (*v1alpha1.SpacesBucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(spacesbucketsResource, "status", c.ns, spacesBucket), &v1alpha1.SpacesBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpacesBucket), err
}

// Delete takes name of the spacesBucket and deletes it. Returns an error if one occurs.
func (c *FakeSpacesBuckets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(spacesbucketsResource, c.ns, name), &v1alpha1.SpacesBucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpacesBuckets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(spacesbucketsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SpacesBucketList{})
	return err
}

// Patch applies the patch and returns the patched spacesBucket.
func (c *FakeSpacesBuckets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpacesBucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(spacesbucketsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SpacesBucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpacesBucket), err
}
