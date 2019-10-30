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

	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudfunctionsFunctionsGetter has a method to return a CloudfunctionsFunctionInterface.
// A group's client should implement this interface.
type CloudfunctionsFunctionsGetter interface {
	CloudfunctionsFunctions(namespace string) CloudfunctionsFunctionInterface
}

// CloudfunctionsFunctionInterface has methods to work with CloudfunctionsFunction resources.
type CloudfunctionsFunctionInterface interface {
	Create(*v1alpha1.CloudfunctionsFunction) (*v1alpha1.CloudfunctionsFunction, error)
	Update(*v1alpha1.CloudfunctionsFunction) (*v1alpha1.CloudfunctionsFunction, error)
	UpdateStatus(*v1alpha1.CloudfunctionsFunction) (*v1alpha1.CloudfunctionsFunction, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudfunctionsFunction, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudfunctionsFunctionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudfunctionsFunction, err error)
	CloudfunctionsFunctionExpansion
}

// cloudfunctionsFunctions implements CloudfunctionsFunctionInterface
type cloudfunctionsFunctions struct {
	client rest.Interface
	ns     string
}

// newCloudfunctionsFunctions returns a CloudfunctionsFunctions
func newCloudfunctionsFunctions(c *GoogleV1alpha1Client, namespace string) *cloudfunctionsFunctions {
	return &cloudfunctionsFunctions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudfunctionsFunction, and returns the corresponding cloudfunctionsFunction object, and an error if there is any.
func (c *cloudfunctionsFunctions) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudfunctionsFunction, err error) {
	result = &v1alpha1.CloudfunctionsFunction{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudfunctionsFunctions that match those selectors.
func (c *cloudfunctionsFunctions) List(opts v1.ListOptions) (result *v1alpha1.CloudfunctionsFunctionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudfunctionsFunctionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudfunctionsFunctions.
func (c *cloudfunctionsFunctions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudfunctionsFunction and creates it.  Returns the server's representation of the cloudfunctionsFunction, and an error, if there is any.
func (c *cloudfunctionsFunctions) Create(cloudfunctionsFunction *v1alpha1.CloudfunctionsFunction) (result *v1alpha1.CloudfunctionsFunction, err error) {
	result = &v1alpha1.CloudfunctionsFunction{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctions").
		Body(cloudfunctionsFunction).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudfunctionsFunction and updates it. Returns the server's representation of the cloudfunctionsFunction, and an error, if there is any.
func (c *cloudfunctionsFunctions) Update(cloudfunctionsFunction *v1alpha1.CloudfunctionsFunction) (result *v1alpha1.CloudfunctionsFunction, err error) {
	result = &v1alpha1.CloudfunctionsFunction{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctions").
		Name(cloudfunctionsFunction.Name).
		Body(cloudfunctionsFunction).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudfunctionsFunctions) UpdateStatus(cloudfunctionsFunction *v1alpha1.CloudfunctionsFunction) (result *v1alpha1.CloudfunctionsFunction, err error) {
	result = &v1alpha1.CloudfunctionsFunction{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctions").
		Name(cloudfunctionsFunction.Name).
		SubResource("status").
		Body(cloudfunctionsFunction).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudfunctionsFunction and deletes it. Returns an error if one occurs.
func (c *cloudfunctionsFunctions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudfunctionsFunctions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudfunctionsfunctions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudfunctionsFunction.
func (c *cloudfunctionsFunctions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudfunctionsFunction, err error) {
	result = &v1alpha1.CloudfunctionsFunction{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudfunctionsfunctions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
