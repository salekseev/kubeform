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

// NotificationHubNamespace_sGetter has a method to return a NotificationHubNamespace_Interface.
// A group's client should implement this interface.
type NotificationHubNamespace_sGetter interface {
	NotificationHubNamespace_s(namespace string) NotificationHubNamespace_Interface
}

// NotificationHubNamespace_Interface has methods to work with NotificationHubNamespace_ resources.
type NotificationHubNamespace_Interface interface {
	Create(*v1alpha1.NotificationHubNamespace_) (*v1alpha1.NotificationHubNamespace_, error)
	Update(*v1alpha1.NotificationHubNamespace_) (*v1alpha1.NotificationHubNamespace_, error)
	UpdateStatus(*v1alpha1.NotificationHubNamespace_) (*v1alpha1.NotificationHubNamespace_, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NotificationHubNamespace_, error)
	List(opts v1.ListOptions) (*v1alpha1.NotificationHubNamespace_List, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NotificationHubNamespace_, err error)
	NotificationHubNamespace_Expansion
}

// notificationHubNamespace_s implements NotificationHubNamespace_Interface
type notificationHubNamespace_s struct {
	client rest.Interface
	ns     string
}

// newNotificationHubNamespace_s returns a NotificationHubNamespace_s
func newNotificationHubNamespace_s(c *AzurermV1alpha1Client, namespace string) *notificationHubNamespace_s {
	return &notificationHubNamespace_s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the notificationHubNamespace_, and returns the corresponding notificationHubNamespace_ object, and an error if there is any.
func (c *notificationHubNamespace_s) Get(name string, options v1.GetOptions) (result *v1alpha1.NotificationHubNamespace_, err error) {
	result = &v1alpha1.NotificationHubNamespace_{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubnamespace_s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NotificationHubNamespace_s that match those selectors.
func (c *notificationHubNamespace_s) List(opts v1.ListOptions) (result *v1alpha1.NotificationHubNamespace_List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NotificationHubNamespace_List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubnamespace_s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested notificationHubNamespace_s.
func (c *notificationHubNamespace_s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubnamespace_s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a notificationHubNamespace_ and creates it.  Returns the server's representation of the notificationHubNamespace_, and an error, if there is any.
func (c *notificationHubNamespace_s) Create(notificationHubNamespace_ *v1alpha1.NotificationHubNamespace_) (result *v1alpha1.NotificationHubNamespace_, err error) {
	result = &v1alpha1.NotificationHubNamespace_{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("notificationhubnamespace_s").
		Body(notificationHubNamespace_).
		Do().
		Into(result)
	return
}

// Update takes the representation of a notificationHubNamespace_ and updates it. Returns the server's representation of the notificationHubNamespace_, and an error, if there is any.
func (c *notificationHubNamespace_s) Update(notificationHubNamespace_ *v1alpha1.NotificationHubNamespace_) (result *v1alpha1.NotificationHubNamespace_, err error) {
	result = &v1alpha1.NotificationHubNamespace_{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notificationhubnamespace_s").
		Name(notificationHubNamespace_.Name).
		Body(notificationHubNamespace_).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *notificationHubNamespace_s) UpdateStatus(notificationHubNamespace_ *v1alpha1.NotificationHubNamespace_) (result *v1alpha1.NotificationHubNamespace_, err error) {
	result = &v1alpha1.NotificationHubNamespace_{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notificationhubnamespace_s").
		Name(notificationHubNamespace_.Name).
		SubResource("status").
		Body(notificationHubNamespace_).
		Do().
		Into(result)
	return
}

// Delete takes name of the notificationHubNamespace_ and deletes it. Returns an error if one occurs.
func (c *notificationHubNamespace_s) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notificationhubnamespace_s").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *notificationHubNamespace_s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notificationhubnamespace_s").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched notificationHubNamespace_.
func (c *notificationHubNamespace_s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NotificationHubNamespace_, err error) {
	result = &v1alpha1.NotificationHubNamespace_{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("notificationhubnamespace_s").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
