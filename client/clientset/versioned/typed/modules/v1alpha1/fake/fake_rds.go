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
	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRDSs implements RDSInterface
type FakeRDSs struct {
	Fake *FakeModulesV1alpha1
	ns   string
}

var rdssResource = schema.GroupVersionResource{Group: "modules.kubeform.com", Version: "v1alpha1", Resource: "rdss"}

var rdssKind = schema.GroupVersionKind{Group: "modules.kubeform.com", Version: "v1alpha1", Kind: "RDS"}

// Get takes name of the rDS, and returns the corresponding rDS object, and an error if there is any.
func (c *FakeRDSs) Get(name string, options v1.GetOptions) (result *v1alpha1.RDS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rdssResource, c.ns, name), &v1alpha1.RDS{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RDS), err
}

// List takes label and field selectors, and returns the list of RDSs that match those selectors.
func (c *FakeRDSs) List(opts v1.ListOptions) (result *v1alpha1.RDSList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rdssResource, rdssKind, c.ns, opts), &v1alpha1.RDSList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RDSList{ListMeta: obj.(*v1alpha1.RDSList).ListMeta}
	for _, item := range obj.(*v1alpha1.RDSList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rDSs.
func (c *FakeRDSs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rdssResource, c.ns, opts))

}

// Create takes the representation of a rDS and creates it.  Returns the server's representation of the rDS, and an error, if there is any.
func (c *FakeRDSs) Create(rDS *v1alpha1.RDS) (result *v1alpha1.RDS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rdssResource, c.ns, rDS), &v1alpha1.RDS{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RDS), err
}

// Update takes the representation of a rDS and updates it. Returns the server's representation of the rDS, and an error, if there is any.
func (c *FakeRDSs) Update(rDS *v1alpha1.RDS) (result *v1alpha1.RDS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rdssResource, c.ns, rDS), &v1alpha1.RDS{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RDS), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRDSs) UpdateStatus(rDS *v1alpha1.RDS) (*v1alpha1.RDS, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rdssResource, "status", c.ns, rDS), &v1alpha1.RDS{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RDS), err
}

// Delete takes name of the rDS and deletes it. Returns an error if one occurs.
func (c *FakeRDSs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rdssResource, c.ns, name), &v1alpha1.RDS{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRDSs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rdssResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RDSList{})
	return err
}

// Patch applies the patch and returns the patched rDS.
func (c *FakeRDSs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RDS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rdssResource, c.ns, name, pt, data, subresources...), &v1alpha1.RDS{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RDS), err
}
