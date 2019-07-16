/*
Copyright 2019 The Kubeform Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

// ElastictranscoderPipelinesGetter has a method to return a ElastictranscoderPipelineInterface.
// A group's client should implement this interface.
type ElastictranscoderPipelinesGetter interface {
	ElastictranscoderPipelines() ElastictranscoderPipelineInterface
}

// ElastictranscoderPipelineInterface has methods to work with ElastictranscoderPipeline resources.
type ElastictranscoderPipelineInterface interface {
	Create(*v1alpha1.ElastictranscoderPipeline) (*v1alpha1.ElastictranscoderPipeline, error)
	Update(*v1alpha1.ElastictranscoderPipeline) (*v1alpha1.ElastictranscoderPipeline, error)
	UpdateStatus(*v1alpha1.ElastictranscoderPipeline) (*v1alpha1.ElastictranscoderPipeline, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ElastictranscoderPipeline, error)
	List(opts v1.ListOptions) (*v1alpha1.ElastictranscoderPipelineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElastictranscoderPipeline, err error)
	ElastictranscoderPipelineExpansion
}

// elastictranscoderPipelines implements ElastictranscoderPipelineInterface
type elastictranscoderPipelines struct {
	client rest.Interface
}

// newElastictranscoderPipelines returns a ElastictranscoderPipelines
func newElastictranscoderPipelines(c *AwsV1alpha1Client) *elastictranscoderPipelines {
	return &elastictranscoderPipelines{
		client: c.RESTClient(),
	}
}

// Get takes name of the elastictranscoderPipeline, and returns the corresponding elastictranscoderPipeline object, and an error if there is any.
func (c *elastictranscoderPipelines) Get(name string, options v1.GetOptions) (result *v1alpha1.ElastictranscoderPipeline, err error) {
	result = &v1alpha1.ElastictranscoderPipeline{}
	err = c.client.Get().
		Resource("elastictranscoderpipelines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ElastictranscoderPipelines that match those selectors.
func (c *elastictranscoderPipelines) List(opts v1.ListOptions) (result *v1alpha1.ElastictranscoderPipelineList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ElastictranscoderPipelineList{}
	err = c.client.Get().
		Resource("elastictranscoderpipelines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested elastictranscoderPipelines.
func (c *elastictranscoderPipelines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("elastictranscoderpipelines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a elastictranscoderPipeline and creates it.  Returns the server's representation of the elastictranscoderPipeline, and an error, if there is any.
func (c *elastictranscoderPipelines) Create(elastictranscoderPipeline *v1alpha1.ElastictranscoderPipeline) (result *v1alpha1.ElastictranscoderPipeline, err error) {
	result = &v1alpha1.ElastictranscoderPipeline{}
	err = c.client.Post().
		Resource("elastictranscoderpipelines").
		Body(elastictranscoderPipeline).
		Do().
		Into(result)
	return
}

// Update takes the representation of a elastictranscoderPipeline and updates it. Returns the server's representation of the elastictranscoderPipeline, and an error, if there is any.
func (c *elastictranscoderPipelines) Update(elastictranscoderPipeline *v1alpha1.ElastictranscoderPipeline) (result *v1alpha1.ElastictranscoderPipeline, err error) {
	result = &v1alpha1.ElastictranscoderPipeline{}
	err = c.client.Put().
		Resource("elastictranscoderpipelines").
		Name(elastictranscoderPipeline.Name).
		Body(elastictranscoderPipeline).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *elastictranscoderPipelines) UpdateStatus(elastictranscoderPipeline *v1alpha1.ElastictranscoderPipeline) (result *v1alpha1.ElastictranscoderPipeline, err error) {
	result = &v1alpha1.ElastictranscoderPipeline{}
	err = c.client.Put().
		Resource("elastictranscoderpipelines").
		Name(elastictranscoderPipeline.Name).
		SubResource("status").
		Body(elastictranscoderPipeline).
		Do().
		Into(result)
	return
}

// Delete takes name of the elastictranscoderPipeline and deletes it. Returns an error if one occurs.
func (c *elastictranscoderPipelines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("elastictranscoderpipelines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *elastictranscoderPipelines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("elastictranscoderpipelines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched elastictranscoderPipeline.
func (c *elastictranscoderPipelines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElastictranscoderPipeline, err error) {
	result = &v1alpha1.ElastictranscoderPipeline{}
	err = c.client.Patch(pt).
		Resource("elastictranscoderpipelines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}