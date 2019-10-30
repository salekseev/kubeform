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
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDomainRecords implements DomainRecordInterface
type FakeDomainRecords struct {
	Fake *FakeLinodeV1alpha1
	ns   string
}

var domainrecordsResource = schema.GroupVersionResource{Group: "linode.kubeform.com", Version: "v1alpha1", Resource: "domainrecords"}

var domainrecordsKind = schema.GroupVersionKind{Group: "linode.kubeform.com", Version: "v1alpha1", Kind: "DomainRecord"}

// Get takes name of the domainRecord, and returns the corresponding domainRecord object, and an error if there is any.
func (c *FakeDomainRecords) Get(name string, options v1.GetOptions) (result *v1alpha1.DomainRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(domainrecordsResource, c.ns, name), &v1alpha1.DomainRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainRecord), err
}

// List takes label and field selectors, and returns the list of DomainRecords that match those selectors.
func (c *FakeDomainRecords) List(opts v1.ListOptions) (result *v1alpha1.DomainRecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(domainrecordsResource, domainrecordsKind, c.ns, opts), &v1alpha1.DomainRecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DomainRecordList{ListMeta: obj.(*v1alpha1.DomainRecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.DomainRecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested domainRecords.
func (c *FakeDomainRecords) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(domainrecordsResource, c.ns, opts))

}

// Create takes the representation of a domainRecord and creates it.  Returns the server's representation of the domainRecord, and an error, if there is any.
func (c *FakeDomainRecords) Create(domainRecord *v1alpha1.DomainRecord) (result *v1alpha1.DomainRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(domainrecordsResource, c.ns, domainRecord), &v1alpha1.DomainRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainRecord), err
}

// Update takes the representation of a domainRecord and updates it. Returns the server's representation of the domainRecord, and an error, if there is any.
func (c *FakeDomainRecords) Update(domainRecord *v1alpha1.DomainRecord) (result *v1alpha1.DomainRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(domainrecordsResource, c.ns, domainRecord), &v1alpha1.DomainRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainRecord), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDomainRecords) UpdateStatus(domainRecord *v1alpha1.DomainRecord) (*v1alpha1.DomainRecord, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(domainrecordsResource, "status", c.ns, domainRecord), &v1alpha1.DomainRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainRecord), err
}

// Delete takes name of the domainRecord and deletes it. Returns an error if one occurs.
func (c *FakeDomainRecords) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(domainrecordsResource, c.ns, name), &v1alpha1.DomainRecord{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDomainRecords) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(domainrecordsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DomainRecordList{})
	return err
}

// Patch applies the patch and returns the patched domainRecord.
func (c *FakeDomainRecords) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DomainRecord, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(domainrecordsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DomainRecord{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainRecord), err
}
