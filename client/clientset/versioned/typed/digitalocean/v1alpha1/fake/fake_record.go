/*
Copyright 2019 The Kubeform Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/digitalocean/v1alpha1"
)

// FakeRecords implements RecordInterface
type FakeRecords struct {
	Fake *FakeDigitaloceanV1alpha1
}

var recordsResource = schema.GroupVersionResource{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Resource: "records"}

var recordsKind = schema.GroupVersionKind{Group: "digitalocean.kubeform.com", Version: "v1alpha1", Kind: "Record"}

// Get takes name of the record, and returns the corresponding record object, and an error if there is any.
func (c *FakeRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.Record, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(recordsResource, name), &v1alpha1.Record{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Record), err
}

// List takes label and field selectors, and returns the list of Records that match those selectors.
func (c *FakeRecords) List(opts v1.ListOptions) (result *v1alpha1.RecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(recordsResource, recordsKind, opts), &v1alpha1.RecordList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RecordList{ListMeta: obj.(*v1alpha1.RecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.RecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested records.
func (c *FakeRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(recordsResource, opts))
}

// Create takes the representation of a record and creates it.  Returns the server's representation of the record, and an error, if there is any.
func (c *FakeRecords) Create(record *v1alpha1.Record) (result *v1alpha1.Record, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(recordsResource, record), &v1alpha1.Record{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Record), err
}

// Update takes the representation of a record and updates it. Returns the server's representation of the record, and an error, if there is any.
func (c *FakeRecords) Update(record *v1alpha1.Record) (result *v1alpha1.Record, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(recordsResource, record), &v1alpha1.Record{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Record), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRecords) UpdateStatus(record *v1alpha1.Record) (*v1alpha1.Record, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(recordsResource, "status", record), &v1alpha1.Record{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Record), err
}

// Delete takes name of the record and deletes it. Returns an error if one occurs.
func (c *FakeRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(recordsResource, name), &v1alpha1.Record{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(recordsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RecordList{})
	return err
}

// Patch applies the patch and returns the patched record.
func (c *FakeRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Record, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(recordsResource, name, pt, data, subresources...), &v1alpha1.Record{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Record), err
}