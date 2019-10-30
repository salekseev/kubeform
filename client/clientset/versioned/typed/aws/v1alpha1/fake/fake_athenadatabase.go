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

// FakeAthenaDatabases implements AthenaDatabaseInterface
type FakeAthenaDatabases struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var athenadatabasesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "athenadatabases"}

var athenadatabasesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AthenaDatabase"}

// Get takes name of the athenaDatabase, and returns the corresponding athenaDatabase object, and an error if there is any.
func (c *FakeAthenaDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.AthenaDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(athenadatabasesResource, c.ns, name), &v1alpha1.AthenaDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaDatabase), err
}

// List takes label and field selectors, and returns the list of AthenaDatabases that match those selectors.
func (c *FakeAthenaDatabases) List(opts v1.ListOptions) (result *v1alpha1.AthenaDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(athenadatabasesResource, athenadatabasesKind, c.ns, opts), &v1alpha1.AthenaDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AthenaDatabaseList{ListMeta: obj.(*v1alpha1.AthenaDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha1.AthenaDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested athenaDatabases.
func (c *FakeAthenaDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(athenadatabasesResource, c.ns, opts))

}

// Create takes the representation of a athenaDatabase and creates it.  Returns the server's representation of the athenaDatabase, and an error, if there is any.
func (c *FakeAthenaDatabases) Create(athenaDatabase *v1alpha1.AthenaDatabase) (result *v1alpha1.AthenaDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(athenadatabasesResource, c.ns, athenaDatabase), &v1alpha1.AthenaDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaDatabase), err
}

// Update takes the representation of a athenaDatabase and updates it. Returns the server's representation of the athenaDatabase, and an error, if there is any.
func (c *FakeAthenaDatabases) Update(athenaDatabase *v1alpha1.AthenaDatabase) (result *v1alpha1.AthenaDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(athenadatabasesResource, c.ns, athenaDatabase), &v1alpha1.AthenaDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAthenaDatabases) UpdateStatus(athenaDatabase *v1alpha1.AthenaDatabase) (*v1alpha1.AthenaDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(athenadatabasesResource, "status", c.ns, athenaDatabase), &v1alpha1.AthenaDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaDatabase), err
}

// Delete takes name of the athenaDatabase and deletes it. Returns an error if one occurs.
func (c *FakeAthenaDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(athenadatabasesResource, c.ns, name), &v1alpha1.AthenaDatabase{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAthenaDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(athenadatabasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AthenaDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched athenaDatabase.
func (c *FakeAthenaDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AthenaDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(athenadatabasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AthenaDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaDatabase), err
}
