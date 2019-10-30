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

package v1alpha1

import (
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SubnetRouteTableAssociationsGetter has a method to return a SubnetRouteTableAssociationInterface.
// A group's client should implement this interface.
type SubnetRouteTableAssociationsGetter interface {
	SubnetRouteTableAssociations(namespace string) SubnetRouteTableAssociationInterface
}

// SubnetRouteTableAssociationInterface has methods to work with SubnetRouteTableAssociation resources.
type SubnetRouteTableAssociationInterface interface {
	Create(*v1alpha1.SubnetRouteTableAssociation) (*v1alpha1.SubnetRouteTableAssociation, error)
	Update(*v1alpha1.SubnetRouteTableAssociation) (*v1alpha1.SubnetRouteTableAssociation, error)
	UpdateStatus(*v1alpha1.SubnetRouteTableAssociation) (*v1alpha1.SubnetRouteTableAssociation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SubnetRouteTableAssociation, error)
	List(opts v1.ListOptions) (*v1alpha1.SubnetRouteTableAssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SubnetRouteTableAssociation, err error)
	SubnetRouteTableAssociationExpansion
}

// subnetRouteTableAssociations implements SubnetRouteTableAssociationInterface
type subnetRouteTableAssociations struct {
	client rest.Interface
	ns     string
}

// newSubnetRouteTableAssociations returns a SubnetRouteTableAssociations
func newSubnetRouteTableAssociations(c *AzurermV1alpha1Client, namespace string) *subnetRouteTableAssociations {
	return &subnetRouteTableAssociations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the subnetRouteTableAssociation, and returns the corresponding subnetRouteTableAssociation object, and an error if there is any.
func (c *subnetRouteTableAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.SubnetRouteTableAssociation, err error) {
	result = &v1alpha1.SubnetRouteTableAssociation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subnetroutetableassociations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SubnetRouteTableAssociations that match those selectors.
func (c *subnetRouteTableAssociations) List(opts v1.ListOptions) (result *v1alpha1.SubnetRouteTableAssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SubnetRouteTableAssociationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subnetroutetableassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested subnetRouteTableAssociations.
func (c *subnetRouteTableAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("subnetroutetableassociations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a subnetRouteTableAssociation and creates it.  Returns the server's representation of the subnetRouteTableAssociation, and an error, if there is any.
func (c *subnetRouteTableAssociations) Create(subnetRouteTableAssociation *v1alpha1.SubnetRouteTableAssociation) (result *v1alpha1.SubnetRouteTableAssociation, err error) {
	result = &v1alpha1.SubnetRouteTableAssociation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("subnetroutetableassociations").
		Body(subnetRouteTableAssociation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a subnetRouteTableAssociation and updates it. Returns the server's representation of the subnetRouteTableAssociation, and an error, if there is any.
func (c *subnetRouteTableAssociations) Update(subnetRouteTableAssociation *v1alpha1.SubnetRouteTableAssociation) (result *v1alpha1.SubnetRouteTableAssociation, err error) {
	result = &v1alpha1.SubnetRouteTableAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subnetroutetableassociations").
		Name(subnetRouteTableAssociation.Name).
		Body(subnetRouteTableAssociation).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *subnetRouteTableAssociations) UpdateStatus(subnetRouteTableAssociation *v1alpha1.SubnetRouteTableAssociation) (result *v1alpha1.SubnetRouteTableAssociation, err error) {
	result = &v1alpha1.SubnetRouteTableAssociation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subnetroutetableassociations").
		Name(subnetRouteTableAssociation.Name).
		SubResource("status").
		Body(subnetRouteTableAssociation).
		Do().
		Into(result)
	return
}

// Delete takes name of the subnetRouteTableAssociation and deletes it. Returns an error if one occurs.
func (c *subnetRouteTableAssociations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subnetroutetableassociations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *subnetRouteTableAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subnetroutetableassociations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched subnetRouteTableAssociation.
func (c *subnetRouteTableAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SubnetRouteTableAssociation, err error) {
	result = &v1alpha1.SubnetRouteTableAssociation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("subnetroutetableassociations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
