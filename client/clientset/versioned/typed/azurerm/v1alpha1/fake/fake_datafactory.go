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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataFactories implements DataFactoryInterface
type FakeDataFactories struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var datafactoriesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "datafactories"}

var datafactoriesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "DataFactory"}

// Get takes name of the dataFactory, and returns the corresponding dataFactory object, and an error if there is any.
func (c *FakeDataFactories) Get(name string, options v1.GetOptions) (result *v1alpha1.DataFactory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datafactoriesResource, c.ns, name), &v1alpha1.DataFactory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactory), err
}

// List takes label and field selectors, and returns the list of DataFactories that match those selectors.
func (c *FakeDataFactories) List(opts v1.ListOptions) (result *v1alpha1.DataFactoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datafactoriesResource, datafactoriesKind, c.ns, opts), &v1alpha1.DataFactoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataFactoryList{ListMeta: obj.(*v1alpha1.DataFactoryList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataFactoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataFactories.
func (c *FakeDataFactories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datafactoriesResource, c.ns, opts))

}

// Create takes the representation of a dataFactory and creates it.  Returns the server's representation of the dataFactory, and an error, if there is any.
func (c *FakeDataFactories) Create(dataFactory *v1alpha1.DataFactory) (result *v1alpha1.DataFactory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datafactoriesResource, c.ns, dataFactory), &v1alpha1.DataFactory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactory), err
}

// Update takes the representation of a dataFactory and updates it. Returns the server's representation of the dataFactory, and an error, if there is any.
func (c *FakeDataFactories) Update(dataFactory *v1alpha1.DataFactory) (result *v1alpha1.DataFactory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datafactoriesResource, c.ns, dataFactory), &v1alpha1.DataFactory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactory), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataFactories) UpdateStatus(dataFactory *v1alpha1.DataFactory) (*v1alpha1.DataFactory, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datafactoriesResource, "status", c.ns, dataFactory), &v1alpha1.DataFactory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactory), err
}

// Delete takes name of the dataFactory and deletes it. Returns an error if one occurs.
func (c *FakeDataFactories) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datafactoriesResource, c.ns, name), &v1alpha1.DataFactory{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataFactories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datafactoriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataFactoryList{})
	return err
}

// Patch applies the patch and returns the patched dataFactory.
func (c *FakeDataFactories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataFactory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datafactoriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DataFactory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DataFactory), err
}
