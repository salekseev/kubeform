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

// GlueJobsGetter has a method to return a GlueJobInterface.
// A group's client should implement this interface.
type GlueJobsGetter interface {
	GlueJobs(namespace string) GlueJobInterface
}

// GlueJobInterface has methods to work with GlueJob resources.
type GlueJobInterface interface {
	Create(*v1alpha1.GlueJob) (*v1alpha1.GlueJob, error)
	Update(*v1alpha1.GlueJob) (*v1alpha1.GlueJob, error)
	UpdateStatus(*v1alpha1.GlueJob) (*v1alpha1.GlueJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GlueJob, error)
	List(opts v1.ListOptions) (*v1alpha1.GlueJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlueJob, err error)
	GlueJobExpansion
}

// glueJobs implements GlueJobInterface
type glueJobs struct {
	client rest.Interface
	ns     string
}

// newGlueJobs returns a GlueJobs
func newGlueJobs(c *AwsV1alpha1Client, namespace string) *glueJobs {
	return &glueJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the glueJob, and returns the corresponding glueJob object, and an error if there is any.
func (c *glueJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.GlueJob, err error) {
	result = &v1alpha1.GlueJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gluejobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlueJobs that match those selectors.
func (c *glueJobs) List(opts v1.ListOptions) (result *v1alpha1.GlueJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GlueJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gluejobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested glueJobs.
func (c *glueJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gluejobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a glueJob and creates it.  Returns the server's representation of the glueJob, and an error, if there is any.
func (c *glueJobs) Create(glueJob *v1alpha1.GlueJob) (result *v1alpha1.GlueJob, err error) {
	result = &v1alpha1.GlueJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gluejobs").
		Body(glueJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a glueJob and updates it. Returns the server's representation of the glueJob, and an error, if there is any.
func (c *glueJobs) Update(glueJob *v1alpha1.GlueJob) (result *v1alpha1.GlueJob, err error) {
	result = &v1alpha1.GlueJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gluejobs").
		Name(glueJob.Name).
		Body(glueJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *glueJobs) UpdateStatus(glueJob *v1alpha1.GlueJob) (result *v1alpha1.GlueJob, err error) {
	result = &v1alpha1.GlueJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gluejobs").
		Name(glueJob.Name).
		SubResource("status").
		Body(glueJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the glueJob and deletes it. Returns an error if one occurs.
func (c *glueJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gluejobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *glueJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gluejobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched glueJob.
func (c *glueJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlueJob, err error) {
	result = &v1alpha1.GlueJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gluejobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
