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

// FakeApiGatewayIntegrations implements ApiGatewayIntegrationInterface
type FakeApiGatewayIntegrations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var apigatewayintegrationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "apigatewayintegrations"}

var apigatewayintegrationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ApiGatewayIntegration"}

// Get takes name of the apiGatewayIntegration, and returns the corresponding apiGatewayIntegration object, and an error if there is any.
func (c *FakeApiGatewayIntegrations) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayIntegration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apigatewayintegrationsResource, c.ns, name), &v1alpha1.ApiGatewayIntegration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayIntegration), err
}

// List takes label and field selectors, and returns the list of ApiGatewayIntegrations that match those selectors.
func (c *FakeApiGatewayIntegrations) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayIntegrationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apigatewayintegrationsResource, apigatewayintegrationsKind, c.ns, opts), &v1alpha1.ApiGatewayIntegrationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiGatewayIntegrationList{ListMeta: obj.(*v1alpha1.ApiGatewayIntegrationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiGatewayIntegrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiGatewayIntegrations.
func (c *FakeApiGatewayIntegrations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apigatewayintegrationsResource, c.ns, opts))

}

// Create takes the representation of a apiGatewayIntegration and creates it.  Returns the server's representation of the apiGatewayIntegration, and an error, if there is any.
func (c *FakeApiGatewayIntegrations) Create(apiGatewayIntegration *v1alpha1.ApiGatewayIntegration) (result *v1alpha1.ApiGatewayIntegration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apigatewayintegrationsResource, c.ns, apiGatewayIntegration), &v1alpha1.ApiGatewayIntegration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayIntegration), err
}

// Update takes the representation of a apiGatewayIntegration and updates it. Returns the server's representation of the apiGatewayIntegration, and an error, if there is any.
func (c *FakeApiGatewayIntegrations) Update(apiGatewayIntegration *v1alpha1.ApiGatewayIntegration) (result *v1alpha1.ApiGatewayIntegration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apigatewayintegrationsResource, c.ns, apiGatewayIntegration), &v1alpha1.ApiGatewayIntegration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayIntegration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiGatewayIntegrations) UpdateStatus(apiGatewayIntegration *v1alpha1.ApiGatewayIntegration) (*v1alpha1.ApiGatewayIntegration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apigatewayintegrationsResource, "status", c.ns, apiGatewayIntegration), &v1alpha1.ApiGatewayIntegration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayIntegration), err
}

// Delete takes name of the apiGatewayIntegration and deletes it. Returns an error if one occurs.
func (c *FakeApiGatewayIntegrations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apigatewayintegrationsResource, c.ns, name), &v1alpha1.ApiGatewayIntegration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiGatewayIntegrations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apigatewayintegrationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiGatewayIntegrationList{})
	return err
}

// Patch applies the patch and returns the patched apiGatewayIntegration.
func (c *FakeApiGatewayIntegrations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayIntegration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apigatewayintegrationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiGatewayIntegration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiGatewayIntegration), err
}
