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

// FakeIamPolicies implements IamPolicyInterface
type FakeIamPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iampoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iampolicies"}

var iampoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamPolicy"}

// Get takes name of the iamPolicy, and returns the corresponding iamPolicy object, and an error if there is any.
func (c *FakeIamPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.IamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iampoliciesResource, c.ns, name), &v1alpha1.IamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamPolicy), err
}

// List takes label and field selectors, and returns the list of IamPolicies that match those selectors.
func (c *FakeIamPolicies) List(opts v1.ListOptions) (result *v1alpha1.IamPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iampoliciesResource, iampoliciesKind, c.ns, opts), &v1alpha1.IamPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamPolicyList{ListMeta: obj.(*v1alpha1.IamPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamPolicies.
func (c *FakeIamPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iampoliciesResource, c.ns, opts))

}

// Create takes the representation of a iamPolicy and creates it.  Returns the server's representation of the iamPolicy, and an error, if there is any.
func (c *FakeIamPolicies) Create(iamPolicy *v1alpha1.IamPolicy) (result *v1alpha1.IamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iampoliciesResource, c.ns, iamPolicy), &v1alpha1.IamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamPolicy), err
}

// Update takes the representation of a iamPolicy and updates it. Returns the server's representation of the iamPolicy, and an error, if there is any.
func (c *FakeIamPolicies) Update(iamPolicy *v1alpha1.IamPolicy) (result *v1alpha1.IamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iampoliciesResource, c.ns, iamPolicy), &v1alpha1.IamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamPolicies) UpdateStatus(iamPolicy *v1alpha1.IamPolicy) (*v1alpha1.IamPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iampoliciesResource, "status", c.ns, iamPolicy), &v1alpha1.IamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamPolicy), err
}

// Delete takes name of the iamPolicy and deletes it. Returns an error if one occurs.
func (c *FakeIamPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iampoliciesResource, c.ns, name), &v1alpha1.IamPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iampoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamPolicyList{})
	return err
}

// Patch applies the patch and returns the patched iamPolicy.
func (c *FakeIamPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iampoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamPolicy), err
}
