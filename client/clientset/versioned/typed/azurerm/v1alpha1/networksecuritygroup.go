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

// NetworkSecurityGroupsGetter has a method to return a NetworkSecurityGroupInterface.
// A group's client should implement this interface.
type NetworkSecurityGroupsGetter interface {
	NetworkSecurityGroups(namespace string) NetworkSecurityGroupInterface
}

// NetworkSecurityGroupInterface has methods to work with NetworkSecurityGroup resources.
type NetworkSecurityGroupInterface interface {
	Create(*v1alpha1.NetworkSecurityGroup) (*v1alpha1.NetworkSecurityGroup, error)
	Update(*v1alpha1.NetworkSecurityGroup) (*v1alpha1.NetworkSecurityGroup, error)
	UpdateStatus(*v1alpha1.NetworkSecurityGroup) (*v1alpha1.NetworkSecurityGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkSecurityGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkSecurityGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkSecurityGroup, err error)
	NetworkSecurityGroupExpansion
}

// networkSecurityGroups implements NetworkSecurityGroupInterface
type networkSecurityGroups struct {
	client rest.Interface
	ns     string
}

// newNetworkSecurityGroups returns a NetworkSecurityGroups
func newNetworkSecurityGroups(c *AzurermV1alpha1Client, namespace string) *networkSecurityGroups {
	return &networkSecurityGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkSecurityGroup, and returns the corresponding networkSecurityGroup object, and an error if there is any.
func (c *networkSecurityGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkSecurityGroup, err error) {
	result = &v1alpha1.NetworkSecurityGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networksecuritygroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkSecurityGroups that match those selectors.
func (c *networkSecurityGroups) List(opts v1.ListOptions) (result *v1alpha1.NetworkSecurityGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NetworkSecurityGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networksecuritygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkSecurityGroups.
func (c *networkSecurityGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networksecuritygroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a networkSecurityGroup and creates it.  Returns the server's representation of the networkSecurityGroup, and an error, if there is any.
func (c *networkSecurityGroups) Create(networkSecurityGroup *v1alpha1.NetworkSecurityGroup) (result *v1alpha1.NetworkSecurityGroup, err error) {
	result = &v1alpha1.NetworkSecurityGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networksecuritygroups").
		Body(networkSecurityGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkSecurityGroup and updates it. Returns the server's representation of the networkSecurityGroup, and an error, if there is any.
func (c *networkSecurityGroups) Update(networkSecurityGroup *v1alpha1.NetworkSecurityGroup) (result *v1alpha1.NetworkSecurityGroup, err error) {
	result = &v1alpha1.NetworkSecurityGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networksecuritygroups").
		Name(networkSecurityGroup.Name).
		Body(networkSecurityGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkSecurityGroups) UpdateStatus(networkSecurityGroup *v1alpha1.NetworkSecurityGroup) (result *v1alpha1.NetworkSecurityGroup, err error) {
	result = &v1alpha1.NetworkSecurityGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networksecuritygroups").
		Name(networkSecurityGroup.Name).
		SubResource("status").
		Body(networkSecurityGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkSecurityGroup and deletes it. Returns an error if one occurs.
func (c *networkSecurityGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networksecuritygroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkSecurityGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networksecuritygroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkSecurityGroup.
func (c *networkSecurityGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkSecurityGroup, err error) {
	result = &v1alpha1.NetworkSecurityGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networksecuritygroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
