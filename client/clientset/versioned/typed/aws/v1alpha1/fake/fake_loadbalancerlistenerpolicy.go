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

// FakeLoadBalancerListenerPolicies implements LoadBalancerListenerPolicyInterface
type FakeLoadBalancerListenerPolicies struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var loadbalancerlistenerpoliciesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "loadbalancerlistenerpolicies"}

var loadbalancerlistenerpoliciesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "LoadBalancerListenerPolicy"}

// Get takes name of the loadBalancerListenerPolicy, and returns the corresponding loadBalancerListenerPolicy object, and an error if there is any.
func (c *FakeLoadBalancerListenerPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.LoadBalancerListenerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(loadbalancerlistenerpoliciesResource, c.ns, name), &v1alpha1.LoadBalancerListenerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancerListenerPolicy), err
}

// List takes label and field selectors, and returns the list of LoadBalancerListenerPolicies that match those selectors.
func (c *FakeLoadBalancerListenerPolicies) List(opts v1.ListOptions) (result *v1alpha1.LoadBalancerListenerPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(loadbalancerlistenerpoliciesResource, loadbalancerlistenerpoliciesKind, c.ns, opts), &v1alpha1.LoadBalancerListenerPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LoadBalancerListenerPolicyList{ListMeta: obj.(*v1alpha1.LoadBalancerListenerPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.LoadBalancerListenerPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loadBalancerListenerPolicies.
func (c *FakeLoadBalancerListenerPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(loadbalancerlistenerpoliciesResource, c.ns, opts))

}

// Create takes the representation of a loadBalancerListenerPolicy and creates it.  Returns the server's representation of the loadBalancerListenerPolicy, and an error, if there is any.
func (c *FakeLoadBalancerListenerPolicies) Create(loadBalancerListenerPolicy *v1alpha1.LoadBalancerListenerPolicy) (result *v1alpha1.LoadBalancerListenerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(loadbalancerlistenerpoliciesResource, c.ns, loadBalancerListenerPolicy), &v1alpha1.LoadBalancerListenerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancerListenerPolicy), err
}

// Update takes the representation of a loadBalancerListenerPolicy and updates it. Returns the server's representation of the loadBalancerListenerPolicy, and an error, if there is any.
func (c *FakeLoadBalancerListenerPolicies) Update(loadBalancerListenerPolicy *v1alpha1.LoadBalancerListenerPolicy) (result *v1alpha1.LoadBalancerListenerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(loadbalancerlistenerpoliciesResource, c.ns, loadBalancerListenerPolicy), &v1alpha1.LoadBalancerListenerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancerListenerPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoadBalancerListenerPolicies) UpdateStatus(loadBalancerListenerPolicy *v1alpha1.LoadBalancerListenerPolicy) (*v1alpha1.LoadBalancerListenerPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(loadbalancerlistenerpoliciesResource, "status", c.ns, loadBalancerListenerPolicy), &v1alpha1.LoadBalancerListenerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancerListenerPolicy), err
}

// Delete takes name of the loadBalancerListenerPolicy and deletes it. Returns an error if one occurs.
func (c *FakeLoadBalancerListenerPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(loadbalancerlistenerpoliciesResource, c.ns, name), &v1alpha1.LoadBalancerListenerPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoadBalancerListenerPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(loadbalancerlistenerpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LoadBalancerListenerPolicyList{})
	return err
}

// Patch applies the patch and returns the patched loadBalancerListenerPolicy.
func (c *FakeLoadBalancerListenerPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoadBalancerListenerPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(loadbalancerlistenerpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LoadBalancerListenerPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LoadBalancerListenerPolicy), err
}
