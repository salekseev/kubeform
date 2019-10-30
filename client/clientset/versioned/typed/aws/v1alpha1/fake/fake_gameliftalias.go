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

// FakeGameliftAliases implements GameliftAliasInterface
type FakeGameliftAliases struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var gameliftaliasesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "gameliftaliases"}

var gameliftaliasesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "GameliftAlias"}

// Get takes name of the gameliftAlias, and returns the corresponding gameliftAlias object, and an error if there is any.
func (c *FakeGameliftAliases) Get(name string, options v1.GetOptions) (result *v1alpha1.GameliftAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gameliftaliasesResource, c.ns, name), &v1alpha1.GameliftAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameliftAlias), err
}

// List takes label and field selectors, and returns the list of GameliftAliases that match those selectors.
func (c *FakeGameliftAliases) List(opts v1.ListOptions) (result *v1alpha1.GameliftAliasList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gameliftaliasesResource, gameliftaliasesKind, c.ns, opts), &v1alpha1.GameliftAliasList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GameliftAliasList{ListMeta: obj.(*v1alpha1.GameliftAliasList).ListMeta}
	for _, item := range obj.(*v1alpha1.GameliftAliasList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gameliftAliases.
func (c *FakeGameliftAliases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gameliftaliasesResource, c.ns, opts))

}

// Create takes the representation of a gameliftAlias and creates it.  Returns the server's representation of the gameliftAlias, and an error, if there is any.
func (c *FakeGameliftAliases) Create(gameliftAlias *v1alpha1.GameliftAlias) (result *v1alpha1.GameliftAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gameliftaliasesResource, c.ns, gameliftAlias), &v1alpha1.GameliftAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameliftAlias), err
}

// Update takes the representation of a gameliftAlias and updates it. Returns the server's representation of the gameliftAlias, and an error, if there is any.
func (c *FakeGameliftAliases) Update(gameliftAlias *v1alpha1.GameliftAlias) (result *v1alpha1.GameliftAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gameliftaliasesResource, c.ns, gameliftAlias), &v1alpha1.GameliftAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameliftAlias), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGameliftAliases) UpdateStatus(gameliftAlias *v1alpha1.GameliftAlias) (*v1alpha1.GameliftAlias, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gameliftaliasesResource, "status", c.ns, gameliftAlias), &v1alpha1.GameliftAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameliftAlias), err
}

// Delete takes name of the gameliftAlias and deletes it. Returns an error if one occurs.
func (c *FakeGameliftAliases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gameliftaliasesResource, c.ns, name), &v1alpha1.GameliftAlias{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGameliftAliases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gameliftaliasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GameliftAliasList{})
	return err
}

// Patch applies the patch and returns the patched gameliftAlias.
func (c *FakeGameliftAliases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GameliftAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gameliftaliasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GameliftAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameliftAlias), err
}
