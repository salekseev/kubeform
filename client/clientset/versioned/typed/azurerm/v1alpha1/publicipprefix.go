/*
Copyright 2019 The Kubeform Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// PublicIpPrefixesGetter has a method to return a PublicIpPrefixInterface.
// A group's client should implement this interface.
type PublicIpPrefixesGetter interface {
	PublicIpPrefixes() PublicIpPrefixInterface
}

// PublicIpPrefixInterface has methods to work with PublicIpPrefix resources.
type PublicIpPrefixInterface interface {
	Create(*v1alpha1.PublicIpPrefix) (*v1alpha1.PublicIpPrefix, error)
	Update(*v1alpha1.PublicIpPrefix) (*v1alpha1.PublicIpPrefix, error)
	UpdateStatus(*v1alpha1.PublicIpPrefix) (*v1alpha1.PublicIpPrefix, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PublicIpPrefix, error)
	List(opts v1.ListOptions) (*v1alpha1.PublicIpPrefixList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PublicIpPrefix, err error)
	PublicIpPrefixExpansion
}

// publicIpPrefixes implements PublicIpPrefixInterface
type publicIpPrefixes struct {
	client rest.Interface
}

// newPublicIpPrefixes returns a PublicIpPrefixes
func newPublicIpPrefixes(c *AzurermV1alpha1Client) *publicIpPrefixes {
	return &publicIpPrefixes{
		client: c.RESTClient(),
	}
}

// Get takes name of the publicIpPrefix, and returns the corresponding publicIpPrefix object, and an error if there is any.
func (c *publicIpPrefixes) Get(name string, options v1.GetOptions) (result *v1alpha1.PublicIpPrefix, err error) {
	result = &v1alpha1.PublicIpPrefix{}
	err = c.client.Get().
		Resource("publicipprefixes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PublicIpPrefixes that match those selectors.
func (c *publicIpPrefixes) List(opts v1.ListOptions) (result *v1alpha1.PublicIpPrefixList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PublicIpPrefixList{}
	err = c.client.Get().
		Resource("publicipprefixes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested publicIpPrefixes.
func (c *publicIpPrefixes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("publicipprefixes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a publicIpPrefix and creates it.  Returns the server's representation of the publicIpPrefix, and an error, if there is any.
func (c *publicIpPrefixes) Create(publicIpPrefix *v1alpha1.PublicIpPrefix) (result *v1alpha1.PublicIpPrefix, err error) {
	result = &v1alpha1.PublicIpPrefix{}
	err = c.client.Post().
		Resource("publicipprefixes").
		Body(publicIpPrefix).
		Do().
		Into(result)
	return
}

// Update takes the representation of a publicIpPrefix and updates it. Returns the server's representation of the publicIpPrefix, and an error, if there is any.
func (c *publicIpPrefixes) Update(publicIpPrefix *v1alpha1.PublicIpPrefix) (result *v1alpha1.PublicIpPrefix, err error) {
	result = &v1alpha1.PublicIpPrefix{}
	err = c.client.Put().
		Resource("publicipprefixes").
		Name(publicIpPrefix.Name).
		Body(publicIpPrefix).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *publicIpPrefixes) UpdateStatus(publicIpPrefix *v1alpha1.PublicIpPrefix) (result *v1alpha1.PublicIpPrefix, err error) {
	result = &v1alpha1.PublicIpPrefix{}
	err = c.client.Put().
		Resource("publicipprefixes").
		Name(publicIpPrefix.Name).
		SubResource("status").
		Body(publicIpPrefix).
		Do().
		Into(result)
	return
}

// Delete takes name of the publicIpPrefix and deletes it. Returns an error if one occurs.
func (c *publicIpPrefixes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("publicipprefixes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *publicIpPrefixes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("publicipprefixes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched publicIpPrefix.
func (c *publicIpPrefixes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PublicIpPrefix, err error) {
	result = &v1alpha1.PublicIpPrefix{}
	err = c.client.Patch(pt).
		Resource("publicipprefixes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}