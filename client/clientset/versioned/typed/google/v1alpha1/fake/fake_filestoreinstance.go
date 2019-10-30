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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFilestoreInstances implements FilestoreInstanceInterface
type FakeFilestoreInstances struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var filestoreinstancesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "filestoreinstances"}

var filestoreinstancesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "FilestoreInstance"}

// Get takes name of the filestoreInstance, and returns the corresponding filestoreInstance object, and an error if there is any.
func (c *FakeFilestoreInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.FilestoreInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(filestoreinstancesResource, c.ns, name), &v1alpha1.FilestoreInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FilestoreInstance), err
}

// List takes label and field selectors, and returns the list of FilestoreInstances that match those selectors.
func (c *FakeFilestoreInstances) List(opts v1.ListOptions) (result *v1alpha1.FilestoreInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(filestoreinstancesResource, filestoreinstancesKind, c.ns, opts), &v1alpha1.FilestoreInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FilestoreInstanceList{ListMeta: obj.(*v1alpha1.FilestoreInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.FilestoreInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested filestoreInstances.
func (c *FakeFilestoreInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(filestoreinstancesResource, c.ns, opts))

}

// Create takes the representation of a filestoreInstance and creates it.  Returns the server's representation of the filestoreInstance, and an error, if there is any.
func (c *FakeFilestoreInstances) Create(filestoreInstance *v1alpha1.FilestoreInstance) (result *v1alpha1.FilestoreInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(filestoreinstancesResource, c.ns, filestoreInstance), &v1alpha1.FilestoreInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FilestoreInstance), err
}

// Update takes the representation of a filestoreInstance and updates it. Returns the server's representation of the filestoreInstance, and an error, if there is any.
func (c *FakeFilestoreInstances) Update(filestoreInstance *v1alpha1.FilestoreInstance) (result *v1alpha1.FilestoreInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(filestoreinstancesResource, c.ns, filestoreInstance), &v1alpha1.FilestoreInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FilestoreInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFilestoreInstances) UpdateStatus(filestoreInstance *v1alpha1.FilestoreInstance) (*v1alpha1.FilestoreInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(filestoreinstancesResource, "status", c.ns, filestoreInstance), &v1alpha1.FilestoreInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FilestoreInstance), err
}

// Delete takes name of the filestoreInstance and deletes it. Returns an error if one occurs.
func (c *FakeFilestoreInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(filestoreinstancesResource, c.ns, name), &v1alpha1.FilestoreInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFilestoreInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(filestoreinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FilestoreInstanceList{})
	return err
}

// Patch applies the patch and returns the patched filestoreInstance.
func (c *FakeFilestoreInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FilestoreInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(filestoreinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FilestoreInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FilestoreInstance), err
}
