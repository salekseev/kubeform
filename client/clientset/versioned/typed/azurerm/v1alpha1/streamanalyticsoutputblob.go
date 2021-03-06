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

// StreamAnalyticsOutputBlobsGetter has a method to return a StreamAnalyticsOutputBlobInterface.
// A group's client should implement this interface.
type StreamAnalyticsOutputBlobsGetter interface {
	StreamAnalyticsOutputBlobs(namespace string) StreamAnalyticsOutputBlobInterface
}

// StreamAnalyticsOutputBlobInterface has methods to work with StreamAnalyticsOutputBlob resources.
type StreamAnalyticsOutputBlobInterface interface {
	Create(*v1alpha1.StreamAnalyticsOutputBlob) (*v1alpha1.StreamAnalyticsOutputBlob, error)
	Update(*v1alpha1.StreamAnalyticsOutputBlob) (*v1alpha1.StreamAnalyticsOutputBlob, error)
	UpdateStatus(*v1alpha1.StreamAnalyticsOutputBlob) (*v1alpha1.StreamAnalyticsOutputBlob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.StreamAnalyticsOutputBlob, error)
	List(opts v1.ListOptions) (*v1alpha1.StreamAnalyticsOutputBlobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StreamAnalyticsOutputBlob, err error)
	StreamAnalyticsOutputBlobExpansion
}

// streamAnalyticsOutputBlobs implements StreamAnalyticsOutputBlobInterface
type streamAnalyticsOutputBlobs struct {
	client rest.Interface
	ns     string
}

// newStreamAnalyticsOutputBlobs returns a StreamAnalyticsOutputBlobs
func newStreamAnalyticsOutputBlobs(c *AzurermV1alpha1Client, namespace string) *streamAnalyticsOutputBlobs {
	return &streamAnalyticsOutputBlobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the streamAnalyticsOutputBlob, and returns the corresponding streamAnalyticsOutputBlob object, and an error if there is any.
func (c *streamAnalyticsOutputBlobs) Get(name string, options v1.GetOptions) (result *v1alpha1.StreamAnalyticsOutputBlob, err error) {
	result = &v1alpha1.StreamAnalyticsOutputBlob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("streamanalyticsoutputblobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StreamAnalyticsOutputBlobs that match those selectors.
func (c *streamAnalyticsOutputBlobs) List(opts v1.ListOptions) (result *v1alpha1.StreamAnalyticsOutputBlobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.StreamAnalyticsOutputBlobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("streamanalyticsoutputblobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested streamAnalyticsOutputBlobs.
func (c *streamAnalyticsOutputBlobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("streamanalyticsoutputblobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a streamAnalyticsOutputBlob and creates it.  Returns the server's representation of the streamAnalyticsOutputBlob, and an error, if there is any.
func (c *streamAnalyticsOutputBlobs) Create(streamAnalyticsOutputBlob *v1alpha1.StreamAnalyticsOutputBlob) (result *v1alpha1.StreamAnalyticsOutputBlob, err error) {
	result = &v1alpha1.StreamAnalyticsOutputBlob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("streamanalyticsoutputblobs").
		Body(streamAnalyticsOutputBlob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a streamAnalyticsOutputBlob and updates it. Returns the server's representation of the streamAnalyticsOutputBlob, and an error, if there is any.
func (c *streamAnalyticsOutputBlobs) Update(streamAnalyticsOutputBlob *v1alpha1.StreamAnalyticsOutputBlob) (result *v1alpha1.StreamAnalyticsOutputBlob, err error) {
	result = &v1alpha1.StreamAnalyticsOutputBlob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("streamanalyticsoutputblobs").
		Name(streamAnalyticsOutputBlob.Name).
		Body(streamAnalyticsOutputBlob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *streamAnalyticsOutputBlobs) UpdateStatus(streamAnalyticsOutputBlob *v1alpha1.StreamAnalyticsOutputBlob) (result *v1alpha1.StreamAnalyticsOutputBlob, err error) {
	result = &v1alpha1.StreamAnalyticsOutputBlob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("streamanalyticsoutputblobs").
		Name(streamAnalyticsOutputBlob.Name).
		SubResource("status").
		Body(streamAnalyticsOutputBlob).
		Do().
		Into(result)
	return
}

// Delete takes name of the streamAnalyticsOutputBlob and deletes it. Returns an error if one occurs.
func (c *streamAnalyticsOutputBlobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("streamanalyticsoutputblobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *streamAnalyticsOutputBlobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("streamanalyticsoutputblobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched streamAnalyticsOutputBlob.
func (c *streamAnalyticsOutputBlobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StreamAnalyticsOutputBlob, err error) {
	result = &v1alpha1.StreamAnalyticsOutputBlob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("streamanalyticsoutputblobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
