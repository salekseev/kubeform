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

// FakePubsubSubscriptionIamMembers implements PubsubSubscriptionIamMemberInterface
type FakePubsubSubscriptionIamMembers struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var pubsubsubscriptioniammembersResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "pubsubsubscriptioniammembers"}

var pubsubsubscriptioniammembersKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "PubsubSubscriptionIamMember"}

// Get takes name of the pubsubSubscriptionIamMember, and returns the corresponding pubsubSubscriptionIamMember object, and an error if there is any.
func (c *FakePubsubSubscriptionIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.PubsubSubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pubsubsubscriptioniammembersResource, c.ns, name), &v1alpha1.PubsubSubscriptionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamMember), err
}

// List takes label and field selectors, and returns the list of PubsubSubscriptionIamMembers that match those selectors.
func (c *FakePubsubSubscriptionIamMembers) List(opts v1.ListOptions) (result *v1alpha1.PubsubSubscriptionIamMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pubsubsubscriptioniammembersResource, pubsubsubscriptioniammembersKind, c.ns, opts), &v1alpha1.PubsubSubscriptionIamMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PubsubSubscriptionIamMemberList{ListMeta: obj.(*v1alpha1.PubsubSubscriptionIamMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.PubsubSubscriptionIamMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pubsubSubscriptionIamMembers.
func (c *FakePubsubSubscriptionIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pubsubsubscriptioniammembersResource, c.ns, opts))

}

// Create takes the representation of a pubsubSubscriptionIamMember and creates it.  Returns the server's representation of the pubsubSubscriptionIamMember, and an error, if there is any.
func (c *FakePubsubSubscriptionIamMembers) Create(pubsubSubscriptionIamMember *v1alpha1.PubsubSubscriptionIamMember) (result *v1alpha1.PubsubSubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pubsubsubscriptioniammembersResource, c.ns, pubsubSubscriptionIamMember), &v1alpha1.PubsubSubscriptionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamMember), err
}

// Update takes the representation of a pubsubSubscriptionIamMember and updates it. Returns the server's representation of the pubsubSubscriptionIamMember, and an error, if there is any.
func (c *FakePubsubSubscriptionIamMembers) Update(pubsubSubscriptionIamMember *v1alpha1.PubsubSubscriptionIamMember) (result *v1alpha1.PubsubSubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pubsubsubscriptioniammembersResource, c.ns, pubsubSubscriptionIamMember), &v1alpha1.PubsubSubscriptionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePubsubSubscriptionIamMembers) UpdateStatus(pubsubSubscriptionIamMember *v1alpha1.PubsubSubscriptionIamMember) (*v1alpha1.PubsubSubscriptionIamMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pubsubsubscriptioniammembersResource, "status", c.ns, pubsubSubscriptionIamMember), &v1alpha1.PubsubSubscriptionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamMember), err
}

// Delete takes name of the pubsubSubscriptionIamMember and deletes it. Returns an error if one occurs.
func (c *FakePubsubSubscriptionIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pubsubsubscriptioniammembersResource, c.ns, name), &v1alpha1.PubsubSubscriptionIamMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePubsubSubscriptionIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pubsubsubscriptioniammembersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PubsubSubscriptionIamMemberList{})
	return err
}

// Patch applies the patch and returns the patched pubsubSubscriptionIamMember.
func (c *FakePubsubSubscriptionIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PubsubSubscriptionIamMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pubsubsubscriptioniammembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.PubsubSubscriptionIamMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamMember), err
}
