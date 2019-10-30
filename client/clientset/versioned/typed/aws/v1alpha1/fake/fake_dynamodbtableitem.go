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

// FakeDynamodbTableItems implements DynamodbTableItemInterface
type FakeDynamodbTableItems struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var dynamodbtableitemsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "dynamodbtableitems"}

var dynamodbtableitemsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DynamodbTableItem"}

// Get takes name of the dynamodbTableItem, and returns the corresponding dynamodbTableItem object, and an error if there is any.
func (c *FakeDynamodbTableItems) Get(name string, options v1.GetOptions) (result *v1alpha1.DynamodbTableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dynamodbtableitemsResource, c.ns, name), &v1alpha1.DynamodbTableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DynamodbTableItem), err
}

// List takes label and field selectors, and returns the list of DynamodbTableItems that match those selectors.
func (c *FakeDynamodbTableItems) List(opts v1.ListOptions) (result *v1alpha1.DynamodbTableItemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dynamodbtableitemsResource, dynamodbtableitemsKind, c.ns, opts), &v1alpha1.DynamodbTableItemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DynamodbTableItemList{ListMeta: obj.(*v1alpha1.DynamodbTableItemList).ListMeta}
	for _, item := range obj.(*v1alpha1.DynamodbTableItemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dynamodbTableItems.
func (c *FakeDynamodbTableItems) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dynamodbtableitemsResource, c.ns, opts))

}

// Create takes the representation of a dynamodbTableItem and creates it.  Returns the server's representation of the dynamodbTableItem, and an error, if there is any.
func (c *FakeDynamodbTableItems) Create(dynamodbTableItem *v1alpha1.DynamodbTableItem) (result *v1alpha1.DynamodbTableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dynamodbtableitemsResource, c.ns, dynamodbTableItem), &v1alpha1.DynamodbTableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DynamodbTableItem), err
}

// Update takes the representation of a dynamodbTableItem and updates it. Returns the server's representation of the dynamodbTableItem, and an error, if there is any.
func (c *FakeDynamodbTableItems) Update(dynamodbTableItem *v1alpha1.DynamodbTableItem) (result *v1alpha1.DynamodbTableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dynamodbtableitemsResource, c.ns, dynamodbTableItem), &v1alpha1.DynamodbTableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DynamodbTableItem), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDynamodbTableItems) UpdateStatus(dynamodbTableItem *v1alpha1.DynamodbTableItem) (*v1alpha1.DynamodbTableItem, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dynamodbtableitemsResource, "status", c.ns, dynamodbTableItem), &v1alpha1.DynamodbTableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DynamodbTableItem), err
}

// Delete takes name of the dynamodbTableItem and deletes it. Returns an error if one occurs.
func (c *FakeDynamodbTableItems) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dynamodbtableitemsResource, c.ns, name), &v1alpha1.DynamodbTableItem{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDynamodbTableItems) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dynamodbtableitemsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DynamodbTableItemList{})
	return err
}

// Patch applies the patch and returns the patched dynamodbTableItem.
func (c *FakeDynamodbTableItems) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DynamodbTableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dynamodbtableitemsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DynamodbTableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DynamodbTableItem), err
}
