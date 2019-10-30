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

// FakeAthenaNamedQueries implements AthenaNamedQueryInterface
type FakeAthenaNamedQueries struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var athenanamedqueriesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "athenanamedqueries"}

var athenanamedqueriesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AthenaNamedQuery"}

// Get takes name of the athenaNamedQuery, and returns the corresponding athenaNamedQuery object, and an error if there is any.
func (c *FakeAthenaNamedQueries) Get(name string, options v1.GetOptions) (result *v1alpha1.AthenaNamedQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(athenanamedqueriesResource, c.ns, name), &v1alpha1.AthenaNamedQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaNamedQuery), err
}

// List takes label and field selectors, and returns the list of AthenaNamedQueries that match those selectors.
func (c *FakeAthenaNamedQueries) List(opts v1.ListOptions) (result *v1alpha1.AthenaNamedQueryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(athenanamedqueriesResource, athenanamedqueriesKind, c.ns, opts), &v1alpha1.AthenaNamedQueryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AthenaNamedQueryList{ListMeta: obj.(*v1alpha1.AthenaNamedQueryList).ListMeta}
	for _, item := range obj.(*v1alpha1.AthenaNamedQueryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested athenaNamedQueries.
func (c *FakeAthenaNamedQueries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(athenanamedqueriesResource, c.ns, opts))

}

// Create takes the representation of a athenaNamedQuery and creates it.  Returns the server's representation of the athenaNamedQuery, and an error, if there is any.
func (c *FakeAthenaNamedQueries) Create(athenaNamedQuery *v1alpha1.AthenaNamedQuery) (result *v1alpha1.AthenaNamedQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(athenanamedqueriesResource, c.ns, athenaNamedQuery), &v1alpha1.AthenaNamedQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaNamedQuery), err
}

// Update takes the representation of a athenaNamedQuery and updates it. Returns the server's representation of the athenaNamedQuery, and an error, if there is any.
func (c *FakeAthenaNamedQueries) Update(athenaNamedQuery *v1alpha1.AthenaNamedQuery) (result *v1alpha1.AthenaNamedQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(athenanamedqueriesResource, c.ns, athenaNamedQuery), &v1alpha1.AthenaNamedQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaNamedQuery), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAthenaNamedQueries) UpdateStatus(athenaNamedQuery *v1alpha1.AthenaNamedQuery) (*v1alpha1.AthenaNamedQuery, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(athenanamedqueriesResource, "status", c.ns, athenaNamedQuery), &v1alpha1.AthenaNamedQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaNamedQuery), err
}

// Delete takes name of the athenaNamedQuery and deletes it. Returns an error if one occurs.
func (c *FakeAthenaNamedQueries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(athenanamedqueriesResource, c.ns, name), &v1alpha1.AthenaNamedQuery{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAthenaNamedQueries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(athenanamedqueriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AthenaNamedQueryList{})
	return err
}

// Patch applies the patch and returns the patched athenaNamedQuery.
func (c *FakeAthenaNamedQueries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AthenaNamedQuery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(athenanamedqueriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AthenaNamedQuery{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AthenaNamedQuery), err
}
