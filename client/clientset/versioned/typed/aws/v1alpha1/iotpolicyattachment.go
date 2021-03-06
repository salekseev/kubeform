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

// IotPolicyAttachmentsGetter has a method to return a IotPolicyAttachmentInterface.
// A group's client should implement this interface.
type IotPolicyAttachmentsGetter interface {
	IotPolicyAttachments(namespace string) IotPolicyAttachmentInterface
}

// IotPolicyAttachmentInterface has methods to work with IotPolicyAttachment resources.
type IotPolicyAttachmentInterface interface {
	Create(*v1alpha1.IotPolicyAttachment) (*v1alpha1.IotPolicyAttachment, error)
	Update(*v1alpha1.IotPolicyAttachment) (*v1alpha1.IotPolicyAttachment, error)
	UpdateStatus(*v1alpha1.IotPolicyAttachment) (*v1alpha1.IotPolicyAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IotPolicyAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.IotPolicyAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotPolicyAttachment, err error)
	IotPolicyAttachmentExpansion
}

// iotPolicyAttachments implements IotPolicyAttachmentInterface
type iotPolicyAttachments struct {
	client rest.Interface
	ns     string
}

// newIotPolicyAttachments returns a IotPolicyAttachments
func newIotPolicyAttachments(c *AwsV1alpha1Client, namespace string) *iotPolicyAttachments {
	return &iotPolicyAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iotPolicyAttachment, and returns the corresponding iotPolicyAttachment object, and an error if there is any.
func (c *iotPolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.IotPolicyAttachment, err error) {
	result = &v1alpha1.IotPolicyAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iotpolicyattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IotPolicyAttachments that match those selectors.
func (c *iotPolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.IotPolicyAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IotPolicyAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iotpolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iotPolicyAttachments.
func (c *iotPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iotpolicyattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iotPolicyAttachment and creates it.  Returns the server's representation of the iotPolicyAttachment, and an error, if there is any.
func (c *iotPolicyAttachments) Create(iotPolicyAttachment *v1alpha1.IotPolicyAttachment) (result *v1alpha1.IotPolicyAttachment, err error) {
	result = &v1alpha1.IotPolicyAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iotpolicyattachments").
		Body(iotPolicyAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iotPolicyAttachment and updates it. Returns the server's representation of the iotPolicyAttachment, and an error, if there is any.
func (c *iotPolicyAttachments) Update(iotPolicyAttachment *v1alpha1.IotPolicyAttachment) (result *v1alpha1.IotPolicyAttachment, err error) {
	result = &v1alpha1.IotPolicyAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iotpolicyattachments").
		Name(iotPolicyAttachment.Name).
		Body(iotPolicyAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iotPolicyAttachments) UpdateStatus(iotPolicyAttachment *v1alpha1.IotPolicyAttachment) (result *v1alpha1.IotPolicyAttachment, err error) {
	result = &v1alpha1.IotPolicyAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iotpolicyattachments").
		Name(iotPolicyAttachment.Name).
		SubResource("status").
		Body(iotPolicyAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the iotPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *iotPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iotpolicyattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iotPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iotpolicyattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iotPolicyAttachment.
func (c *iotPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IotPolicyAttachment, err error) {
	result = &v1alpha1.IotPolicyAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iotpolicyattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
