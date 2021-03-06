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

	v1alpha1 "kubeform.dev/kubeform/apis/modules/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RDSsGetter has a method to return a RDSInterface.
// A group's client should implement this interface.
type RDSsGetter interface {
	RDSs(namespace string) RDSInterface
}

// RDSInterface has methods to work with RDS resources.
type RDSInterface interface {
	Create(*v1alpha1.RDS) (*v1alpha1.RDS, error)
	Update(*v1alpha1.RDS) (*v1alpha1.RDS, error)
	UpdateStatus(*v1alpha1.RDS) (*v1alpha1.RDS, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RDS, error)
	List(opts v1.ListOptions) (*v1alpha1.RDSList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RDS, err error)
	RDSExpansion
}

// rDSs implements RDSInterface
type rDSs struct {
	client rest.Interface
	ns     string
}

// newRDSs returns a RDSs
func newRDSs(c *ModulesV1alpha1Client, namespace string) *rDSs {
	return &rDSs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rDS, and returns the corresponding rDS object, and an error if there is any.
func (c *rDSs) Get(name string, options v1.GetOptions) (result *v1alpha1.RDS, err error) {
	result = &v1alpha1.RDS{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rdss").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RDSs that match those selectors.
func (c *rDSs) List(opts v1.ListOptions) (result *v1alpha1.RDSList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RDSList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rdss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rDSs.
func (c *rDSs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rdss").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a rDS and creates it.  Returns the server's representation of the rDS, and an error, if there is any.
func (c *rDSs) Create(rDS *v1alpha1.RDS) (result *v1alpha1.RDS, err error) {
	result = &v1alpha1.RDS{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rdss").
		Body(rDS).
		Do().
		Into(result)
	return
}

// Update takes the representation of a rDS and updates it. Returns the server's representation of the rDS, and an error, if there is any.
func (c *rDSs) Update(rDS *v1alpha1.RDS) (result *v1alpha1.RDS, err error) {
	result = &v1alpha1.RDS{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rdss").
		Name(rDS.Name).
		Body(rDS).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *rDSs) UpdateStatus(rDS *v1alpha1.RDS) (result *v1alpha1.RDS, err error) {
	result = &v1alpha1.RDS{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rdss").
		Name(rDS.Name).
		SubResource("status").
		Body(rDS).
		Do().
		Into(result)
	return
}

// Delete takes name of the rDS and deletes it. Returns an error if one occurs.
func (c *rDSs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rdss").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rDSs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rdss").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched rDS.
func (c *rDSs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RDS, err error) {
	result = &v1alpha1.RDS{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rdss").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
