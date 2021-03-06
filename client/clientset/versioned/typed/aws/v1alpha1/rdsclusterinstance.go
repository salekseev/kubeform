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

// RdsClusterInstancesGetter has a method to return a RdsClusterInstanceInterface.
// A group's client should implement this interface.
type RdsClusterInstancesGetter interface {
	RdsClusterInstances(namespace string) RdsClusterInstanceInterface
}

// RdsClusterInstanceInterface has methods to work with RdsClusterInstance resources.
type RdsClusterInstanceInterface interface {
	Create(*v1alpha1.RdsClusterInstance) (*v1alpha1.RdsClusterInstance, error)
	Update(*v1alpha1.RdsClusterInstance) (*v1alpha1.RdsClusterInstance, error)
	UpdateStatus(*v1alpha1.RdsClusterInstance) (*v1alpha1.RdsClusterInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RdsClusterInstance, error)
	List(opts v1.ListOptions) (*v1alpha1.RdsClusterInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RdsClusterInstance, err error)
	RdsClusterInstanceExpansion
}

// rdsClusterInstances implements RdsClusterInstanceInterface
type rdsClusterInstances struct {
	client rest.Interface
	ns     string
}

// newRdsClusterInstances returns a RdsClusterInstances
func newRdsClusterInstances(c *AwsV1alpha1Client, namespace string) *rdsClusterInstances {
	return &rdsClusterInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rdsClusterInstance, and returns the corresponding rdsClusterInstance object, and an error if there is any.
func (c *rdsClusterInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.RdsClusterInstance, err error) {
	result = &v1alpha1.RdsClusterInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rdsclusterinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RdsClusterInstances that match those selectors.
func (c *rdsClusterInstances) List(opts v1.ListOptions) (result *v1alpha1.RdsClusterInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RdsClusterInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rdsclusterinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rdsClusterInstances.
func (c *rdsClusterInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rdsclusterinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a rdsClusterInstance and creates it.  Returns the server's representation of the rdsClusterInstance, and an error, if there is any.
func (c *rdsClusterInstances) Create(rdsClusterInstance *v1alpha1.RdsClusterInstance) (result *v1alpha1.RdsClusterInstance, err error) {
	result = &v1alpha1.RdsClusterInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rdsclusterinstances").
		Body(rdsClusterInstance).
		Do().
		Into(result)
	return
}

// Update takes the representation of a rdsClusterInstance and updates it. Returns the server's representation of the rdsClusterInstance, and an error, if there is any.
func (c *rdsClusterInstances) Update(rdsClusterInstance *v1alpha1.RdsClusterInstance) (result *v1alpha1.RdsClusterInstance, err error) {
	result = &v1alpha1.RdsClusterInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rdsclusterinstances").
		Name(rdsClusterInstance.Name).
		Body(rdsClusterInstance).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *rdsClusterInstances) UpdateStatus(rdsClusterInstance *v1alpha1.RdsClusterInstance) (result *v1alpha1.RdsClusterInstance, err error) {
	result = &v1alpha1.RdsClusterInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rdsclusterinstances").
		Name(rdsClusterInstance.Name).
		SubResource("status").
		Body(rdsClusterInstance).
		Do().
		Into(result)
	return
}

// Delete takes name of the rdsClusterInstance and deletes it. Returns an error if one occurs.
func (c *rdsClusterInstances) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rdsclusterinstances").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rdsClusterInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rdsclusterinstances").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched rdsClusterInstance.
func (c *rdsClusterInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RdsClusterInstance, err error) {
	result = &v1alpha1.RdsClusterInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rdsclusterinstances").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
