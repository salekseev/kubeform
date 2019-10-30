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

// ComputeTargetTcpProxiesGetter has a method to return a ComputeTargetTcpProxyInterface.
// A group's client should implement this interface.
type ComputeTargetTcpProxiesGetter interface {
	ComputeTargetTcpProxies(namespace string) ComputeTargetTcpProxyInterface
}

// ComputeTargetTcpProxyInterface has methods to work with ComputeTargetTcpProxy resources.
type ComputeTargetTcpProxyInterface interface {
	Create(*v1alpha1.ComputeTargetTcpProxy) (*v1alpha1.ComputeTargetTcpProxy, error)
	Update(*v1alpha1.ComputeTargetTcpProxy) (*v1alpha1.ComputeTargetTcpProxy, error)
	UpdateStatus(*v1alpha1.ComputeTargetTcpProxy) (*v1alpha1.ComputeTargetTcpProxy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ComputeTargetTcpProxy, error)
	List(opts v1.ListOptions) (*v1alpha1.ComputeTargetTcpProxyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeTargetTcpProxy, err error)
	ComputeTargetTcpProxyExpansion
}

// computeTargetTcpProxies implements ComputeTargetTcpProxyInterface
type computeTargetTcpProxies struct {
	client rest.Interface
	ns     string
}

// newComputeTargetTcpProxies returns a ComputeTargetTcpProxies
func newComputeTargetTcpProxies(c *GoogleV1alpha1Client, namespace string) *computeTargetTcpProxies {
	return &computeTargetTcpProxies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeTargetTcpProxy, and returns the corresponding computeTargetTcpProxy object, and an error if there is any.
func (c *computeTargetTcpProxies) Get(name string, options v1.GetOptions) (result *v1alpha1.ComputeTargetTcpProxy, err error) {
	result = &v1alpha1.ComputeTargetTcpProxy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeTargetTcpProxies that match those selectors.
func (c *computeTargetTcpProxies) List(opts v1.ListOptions) (result *v1alpha1.ComputeTargetTcpProxyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComputeTargetTcpProxyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeTargetTcpProxies.
func (c *computeTargetTcpProxies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a computeTargetTcpProxy and creates it.  Returns the server's representation of the computeTargetTcpProxy, and an error, if there is any.
func (c *computeTargetTcpProxies) Create(computeTargetTcpProxy *v1alpha1.ComputeTargetTcpProxy) (result *v1alpha1.ComputeTargetTcpProxy, err error) {
	result = &v1alpha1.ComputeTargetTcpProxy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		Body(computeTargetTcpProxy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a computeTargetTcpProxy and updates it. Returns the server's representation of the computeTargetTcpProxy, and an error, if there is any.
func (c *computeTargetTcpProxies) Update(computeTargetTcpProxy *v1alpha1.ComputeTargetTcpProxy) (result *v1alpha1.ComputeTargetTcpProxy, err error) {
	result = &v1alpha1.ComputeTargetTcpProxy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		Name(computeTargetTcpProxy.Name).
		Body(computeTargetTcpProxy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *computeTargetTcpProxies) UpdateStatus(computeTargetTcpProxy *v1alpha1.ComputeTargetTcpProxy) (result *v1alpha1.ComputeTargetTcpProxy, err error) {
	result = &v1alpha1.ComputeTargetTcpProxy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		Name(computeTargetTcpProxy.Name).
		SubResource("status").
		Body(computeTargetTcpProxy).
		Do().
		Into(result)
	return
}

// Delete takes name of the computeTargetTcpProxy and deletes it. Returns an error if one occurs.
func (c *computeTargetTcpProxies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeTargetTcpProxies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched computeTargetTcpProxy.
func (c *computeTargetTcpProxies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ComputeTargetTcpProxy, err error) {
	result = &v1alpha1.ComputeTargetTcpProxy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
