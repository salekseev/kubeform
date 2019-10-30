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

// PubsubSubscriptionIamMembersGetter has a method to return a PubsubSubscriptionIamMemberInterface.
// A group's client should implement this interface.
type PubsubSubscriptionIamMembersGetter interface {
	PubsubSubscriptionIamMembers(namespace string) PubsubSubscriptionIamMemberInterface
}

// PubsubSubscriptionIamMemberInterface has methods to work with PubsubSubscriptionIamMember resources.
type PubsubSubscriptionIamMemberInterface interface {
	Create(*v1alpha1.PubsubSubscriptionIamMember) (*v1alpha1.PubsubSubscriptionIamMember, error)
	Update(*v1alpha1.PubsubSubscriptionIamMember) (*v1alpha1.PubsubSubscriptionIamMember, error)
	UpdateStatus(*v1alpha1.PubsubSubscriptionIamMember) (*v1alpha1.PubsubSubscriptionIamMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PubsubSubscriptionIamMember, error)
	List(opts v1.ListOptions) (*v1alpha1.PubsubSubscriptionIamMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PubsubSubscriptionIamMember, err error)
	PubsubSubscriptionIamMemberExpansion
}

// pubsubSubscriptionIamMembers implements PubsubSubscriptionIamMemberInterface
type pubsubSubscriptionIamMembers struct {
	client rest.Interface
	ns     string
}

// newPubsubSubscriptionIamMembers returns a PubsubSubscriptionIamMembers
func newPubsubSubscriptionIamMembers(c *GoogleV1alpha1Client, namespace string) *pubsubSubscriptionIamMembers {
	return &pubsubSubscriptionIamMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pubsubSubscriptionIamMember, and returns the corresponding pubsubSubscriptionIamMember object, and an error if there is any.
func (c *pubsubSubscriptionIamMembers) Get(name string, options v1.GetOptions) (result *v1alpha1.PubsubSubscriptionIamMember, err error) {
	result = &v1alpha1.PubsubSubscriptionIamMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pubsubsubscriptioniammembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PubsubSubscriptionIamMembers that match those selectors.
func (c *pubsubSubscriptionIamMembers) List(opts v1.ListOptions) (result *v1alpha1.PubsubSubscriptionIamMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PubsubSubscriptionIamMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pubsubsubscriptioniammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pubsubSubscriptionIamMembers.
func (c *pubsubSubscriptionIamMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pubsubsubscriptioniammembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a pubsubSubscriptionIamMember and creates it.  Returns the server's representation of the pubsubSubscriptionIamMember, and an error, if there is any.
func (c *pubsubSubscriptionIamMembers) Create(pubsubSubscriptionIamMember *v1alpha1.PubsubSubscriptionIamMember) (result *v1alpha1.PubsubSubscriptionIamMember, err error) {
	result = &v1alpha1.PubsubSubscriptionIamMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pubsubsubscriptioniammembers").
		Body(pubsubSubscriptionIamMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pubsubSubscriptionIamMember and updates it. Returns the server's representation of the pubsubSubscriptionIamMember, and an error, if there is any.
func (c *pubsubSubscriptionIamMembers) Update(pubsubSubscriptionIamMember *v1alpha1.PubsubSubscriptionIamMember) (result *v1alpha1.PubsubSubscriptionIamMember, err error) {
	result = &v1alpha1.PubsubSubscriptionIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pubsubsubscriptioniammembers").
		Name(pubsubSubscriptionIamMember.Name).
		Body(pubsubSubscriptionIamMember).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pubsubSubscriptionIamMembers) UpdateStatus(pubsubSubscriptionIamMember *v1alpha1.PubsubSubscriptionIamMember) (result *v1alpha1.PubsubSubscriptionIamMember, err error) {
	result = &v1alpha1.PubsubSubscriptionIamMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pubsubsubscriptioniammembers").
		Name(pubsubSubscriptionIamMember.Name).
		SubResource("status").
		Body(pubsubSubscriptionIamMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the pubsubSubscriptionIamMember and deletes it. Returns an error if one occurs.
func (c *pubsubSubscriptionIamMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pubsubsubscriptioniammembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pubsubSubscriptionIamMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pubsubsubscriptioniammembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pubsubSubscriptionIamMember.
func (c *pubsubSubscriptionIamMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PubsubSubscriptionIamMember, err error) {
	result = &v1alpha1.PubsubSubscriptionIamMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pubsubsubscriptioniammembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
