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

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudhsmV2HsmsGetter has a method to return a CloudhsmV2HsmInterface.
// A group's client should implement this interface.
type CloudhsmV2HsmsGetter interface {
	CloudhsmV2Hsms(namespace string) CloudhsmV2HsmInterface
}

// CloudhsmV2HsmInterface has methods to work with CloudhsmV2Hsm resources.
type CloudhsmV2HsmInterface interface {
	Create(*v1alpha1.CloudhsmV2Hsm) (*v1alpha1.CloudhsmV2Hsm, error)
	Update(*v1alpha1.CloudhsmV2Hsm) (*v1alpha1.CloudhsmV2Hsm, error)
	UpdateStatus(*v1alpha1.CloudhsmV2Hsm) (*v1alpha1.CloudhsmV2Hsm, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudhsmV2Hsm, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudhsmV2HsmList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudhsmV2Hsm, err error)
	CloudhsmV2HsmExpansion
}

// cloudhsmV2Hsms implements CloudhsmV2HsmInterface
type cloudhsmV2Hsms struct {
	client rest.Interface
	ns     string
}

// newCloudhsmV2Hsms returns a CloudhsmV2Hsms
func newCloudhsmV2Hsms(c *AwsV1alpha1Client, namespace string) *cloudhsmV2Hsms {
	return &cloudhsmV2Hsms{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudhsmV2Hsm, and returns the corresponding cloudhsmV2Hsm object, and an error if there is any.
func (c *cloudhsmV2Hsms) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudhsmV2Hsm, err error) {
	result = &v1alpha1.CloudhsmV2Hsm{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudhsmv2hsms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudhsmV2Hsms that match those selectors.
func (c *cloudhsmV2Hsms) List(opts v1.ListOptions) (result *v1alpha1.CloudhsmV2HsmList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudhsmV2HsmList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudhsmv2hsms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudhsmV2Hsms.
func (c *cloudhsmV2Hsms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudhsmv2hsms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudhsmV2Hsm and creates it.  Returns the server's representation of the cloudhsmV2Hsm, and an error, if there is any.
func (c *cloudhsmV2Hsms) Create(cloudhsmV2Hsm *v1alpha1.CloudhsmV2Hsm) (result *v1alpha1.CloudhsmV2Hsm, err error) {
	result = &v1alpha1.CloudhsmV2Hsm{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudhsmv2hsms").
		Body(cloudhsmV2Hsm).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudhsmV2Hsm and updates it. Returns the server's representation of the cloudhsmV2Hsm, and an error, if there is any.
func (c *cloudhsmV2Hsms) Update(cloudhsmV2Hsm *v1alpha1.CloudhsmV2Hsm) (result *v1alpha1.CloudhsmV2Hsm, err error) {
	result = &v1alpha1.CloudhsmV2Hsm{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudhsmv2hsms").
		Name(cloudhsmV2Hsm.Name).
		Body(cloudhsmV2Hsm).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudhsmV2Hsms) UpdateStatus(cloudhsmV2Hsm *v1alpha1.CloudhsmV2Hsm) (result *v1alpha1.CloudhsmV2Hsm, err error) {
	result = &v1alpha1.CloudhsmV2Hsm{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudhsmv2hsms").
		Name(cloudhsmV2Hsm.Name).
		SubResource("status").
		Body(cloudhsmV2Hsm).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudhsmV2Hsm and deletes it. Returns an error if one occurs.
func (c *cloudhsmV2Hsms) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudhsmv2hsms").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudhsmV2Hsms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudhsmv2hsms").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudhsmV2Hsm.
func (c *cloudhsmV2Hsms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudhsmV2Hsm, err error) {
	result = &v1alpha1.CloudhsmV2Hsm{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudhsmv2hsms").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
