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

// ServiceDiscoveryPrivateDNSNamespacesGetter has a method to return a ServiceDiscoveryPrivateDNSNamespaceInterface.
// A group's client should implement this interface.
type ServiceDiscoveryPrivateDNSNamespacesGetter interface {
	ServiceDiscoveryPrivateDNSNamespaces(namespace string) ServiceDiscoveryPrivateDNSNamespaceInterface
}

// ServiceDiscoveryPrivateDNSNamespaceInterface has methods to work with ServiceDiscoveryPrivateDNSNamespace resources.
type ServiceDiscoveryPrivateDNSNamespaceInterface interface {
	Create(*v1alpha1.ServiceDiscoveryPrivateDNSNamespace) (*v1alpha1.ServiceDiscoveryPrivateDNSNamespace, error)
	Update(*v1alpha1.ServiceDiscoveryPrivateDNSNamespace) (*v1alpha1.ServiceDiscoveryPrivateDNSNamespace, error)
	UpdateStatus(*v1alpha1.ServiceDiscoveryPrivateDNSNamespace) (*v1alpha1.ServiceDiscoveryPrivateDNSNamespace, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServiceDiscoveryPrivateDNSNamespace, error)
	List(opts v1.ListOptions) (*v1alpha1.ServiceDiscoveryPrivateDNSNamespaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespace, err error)
	ServiceDiscoveryPrivateDNSNamespaceExpansion
}

// serviceDiscoveryPrivateDNSNamespaces implements ServiceDiscoveryPrivateDNSNamespaceInterface
type serviceDiscoveryPrivateDNSNamespaces struct {
	client rest.Interface
	ns     string
}

// newServiceDiscoveryPrivateDNSNamespaces returns a ServiceDiscoveryPrivateDNSNamespaces
func newServiceDiscoveryPrivateDNSNamespaces(c *AwsV1alpha1Client, namespace string) *serviceDiscoveryPrivateDNSNamespaces {
	return &serviceDiscoveryPrivateDNSNamespaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceDiscoveryPrivateDNSNamespace, and returns the corresponding serviceDiscoveryPrivateDNSNamespace object, and an error if there is any.
func (c *serviceDiscoveryPrivateDNSNamespaces) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespace, err error) {
	result = &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicediscoveryprivatednsnamespaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceDiscoveryPrivateDNSNamespaces that match those selectors.
func (c *serviceDiscoveryPrivateDNSNamespaces) List(opts v1.ListOptions) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceDiscoveryPrivateDNSNamespaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicediscoveryprivatednsnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceDiscoveryPrivateDNSNamespaces.
func (c *serviceDiscoveryPrivateDNSNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicediscoveryprivatednsnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a serviceDiscoveryPrivateDNSNamespace and creates it.  Returns the server's representation of the serviceDiscoveryPrivateDNSNamespace, and an error, if there is any.
func (c *serviceDiscoveryPrivateDNSNamespaces) Create(serviceDiscoveryPrivateDNSNamespace *v1alpha1.ServiceDiscoveryPrivateDNSNamespace) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespace, err error) {
	result = &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicediscoveryprivatednsnamespaces").
		Body(serviceDiscoveryPrivateDNSNamespace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceDiscoveryPrivateDNSNamespace and updates it. Returns the server's representation of the serviceDiscoveryPrivateDNSNamespace, and an error, if there is any.
func (c *serviceDiscoveryPrivateDNSNamespaces) Update(serviceDiscoveryPrivateDNSNamespace *v1alpha1.ServiceDiscoveryPrivateDNSNamespace) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespace, err error) {
	result = &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicediscoveryprivatednsnamespaces").
		Name(serviceDiscoveryPrivateDNSNamespace.Name).
		Body(serviceDiscoveryPrivateDNSNamespace).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *serviceDiscoveryPrivateDNSNamespaces) UpdateStatus(serviceDiscoveryPrivateDNSNamespace *v1alpha1.ServiceDiscoveryPrivateDNSNamespace) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespace, err error) {
	result = &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicediscoveryprivatednsnamespaces").
		Name(serviceDiscoveryPrivateDNSNamespace.Name).
		SubResource("status").
		Body(serviceDiscoveryPrivateDNSNamespace).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceDiscoveryPrivateDNSNamespace and deletes it. Returns an error if one occurs.
func (c *serviceDiscoveryPrivateDNSNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicediscoveryprivatednsnamespaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceDiscoveryPrivateDNSNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicediscoveryprivatednsnamespaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceDiscoveryPrivateDNSNamespace.
func (c *serviceDiscoveryPrivateDNSNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespace, err error) {
	result = &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicediscoveryprivatednsnamespaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
