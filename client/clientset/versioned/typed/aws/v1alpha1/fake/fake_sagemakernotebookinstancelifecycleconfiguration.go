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

// FakeSagemakerNotebookInstanceLifecycleConfigurations implements SagemakerNotebookInstanceLifecycleConfigurationInterface
type FakeSagemakerNotebookInstanceLifecycleConfigurations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var sagemakernotebookinstancelifecycleconfigurationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "sagemakernotebookinstancelifecycleconfigurations"}

var sagemakernotebookinstancelifecycleconfigurationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SagemakerNotebookInstanceLifecycleConfiguration"}

// Get takes name of the sagemakerNotebookInstanceLifecycleConfiguration, and returns the corresponding sagemakerNotebookInstanceLifecycleConfiguration object, and an error if there is any.
func (c *FakeSagemakerNotebookInstanceLifecycleConfigurations) Get(name string, options v1.GetOptions) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sagemakernotebookinstancelifecycleconfigurationsResource, c.ns, name), &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration), err
}

// List takes label and field selectors, and returns the list of SagemakerNotebookInstanceLifecycleConfigurations that match those selectors.
func (c *FakeSagemakerNotebookInstanceLifecycleConfigurations) List(opts v1.ListOptions) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sagemakernotebookinstancelifecycleconfigurationsResource, sagemakernotebookinstancelifecycleconfigurationsKind, c.ns, opts), &v1alpha1.SagemakerNotebookInstanceLifecycleConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SagemakerNotebookInstanceLifecycleConfigurationList{ListMeta: obj.(*v1alpha1.SagemakerNotebookInstanceLifecycleConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.SagemakerNotebookInstanceLifecycleConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sagemakerNotebookInstanceLifecycleConfigurations.
func (c *FakeSagemakerNotebookInstanceLifecycleConfigurations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sagemakernotebookinstancelifecycleconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a sagemakerNotebookInstanceLifecycleConfiguration and creates it.  Returns the server's representation of the sagemakerNotebookInstanceLifecycleConfiguration, and an error, if there is any.
func (c *FakeSagemakerNotebookInstanceLifecycleConfigurations) Create(sagemakerNotebookInstanceLifecycleConfiguration *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sagemakernotebookinstancelifecycleconfigurationsResource, c.ns, sagemakerNotebookInstanceLifecycleConfiguration), &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration), err
}

// Update takes the representation of a sagemakerNotebookInstanceLifecycleConfiguration and updates it. Returns the server's representation of the sagemakerNotebookInstanceLifecycleConfiguration, and an error, if there is any.
func (c *FakeSagemakerNotebookInstanceLifecycleConfigurations) Update(sagemakerNotebookInstanceLifecycleConfiguration *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sagemakernotebookinstancelifecycleconfigurationsResource, c.ns, sagemakerNotebookInstanceLifecycleConfiguration), &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSagemakerNotebookInstanceLifecycleConfigurations) UpdateStatus(sagemakerNotebookInstanceLifecycleConfiguration *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration) (*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sagemakernotebookinstancelifecycleconfigurationsResource, "status", c.ns, sagemakerNotebookInstanceLifecycleConfiguration), &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration), err
}

// Delete takes name of the sagemakerNotebookInstanceLifecycleConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeSagemakerNotebookInstanceLifecycleConfigurations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sagemakernotebookinstancelifecycleconfigurationsResource, c.ns, name), &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSagemakerNotebookInstanceLifecycleConfigurations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sagemakernotebookinstancelifecycleconfigurationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SagemakerNotebookInstanceLifecycleConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched sagemakerNotebookInstanceLifecycleConfiguration.
func (c *FakeSagemakerNotebookInstanceLifecycleConfigurations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sagemakernotebookinstancelifecycleconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SagemakerNotebookInstanceLifecycleConfiguration), err
}
