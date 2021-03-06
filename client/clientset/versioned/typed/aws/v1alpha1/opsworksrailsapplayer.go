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

// OpsworksRailsAppLayersGetter has a method to return a OpsworksRailsAppLayerInterface.
// A group's client should implement this interface.
type OpsworksRailsAppLayersGetter interface {
	OpsworksRailsAppLayers(namespace string) OpsworksRailsAppLayerInterface
}

// OpsworksRailsAppLayerInterface has methods to work with OpsworksRailsAppLayer resources.
type OpsworksRailsAppLayerInterface interface {
	Create(*v1alpha1.OpsworksRailsAppLayer) (*v1alpha1.OpsworksRailsAppLayer, error)
	Update(*v1alpha1.OpsworksRailsAppLayer) (*v1alpha1.OpsworksRailsAppLayer, error)
	UpdateStatus(*v1alpha1.OpsworksRailsAppLayer) (*v1alpha1.OpsworksRailsAppLayer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.OpsworksRailsAppLayer, error)
	List(opts v1.ListOptions) (*v1alpha1.OpsworksRailsAppLayerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksRailsAppLayer, err error)
	OpsworksRailsAppLayerExpansion
}

// opsworksRailsAppLayers implements OpsworksRailsAppLayerInterface
type opsworksRailsAppLayers struct {
	client rest.Interface
	ns     string
}

// newOpsworksRailsAppLayers returns a OpsworksRailsAppLayers
func newOpsworksRailsAppLayers(c *AwsV1alpha1Client, namespace string) *opsworksRailsAppLayers {
	return &opsworksRailsAppLayers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the opsworksRailsAppLayer, and returns the corresponding opsworksRailsAppLayer object, and an error if there is any.
func (c *opsworksRailsAppLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksRailsAppLayer, err error) {
	result = &v1alpha1.OpsworksRailsAppLayer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("opsworksrailsapplayers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpsworksRailsAppLayers that match those selectors.
func (c *opsworksRailsAppLayers) List(opts v1.ListOptions) (result *v1alpha1.OpsworksRailsAppLayerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OpsworksRailsAppLayerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("opsworksrailsapplayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested opsworksRailsAppLayers.
func (c *opsworksRailsAppLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("opsworksrailsapplayers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a opsworksRailsAppLayer and creates it.  Returns the server's representation of the opsworksRailsAppLayer, and an error, if there is any.
func (c *opsworksRailsAppLayers) Create(opsworksRailsAppLayer *v1alpha1.OpsworksRailsAppLayer) (result *v1alpha1.OpsworksRailsAppLayer, err error) {
	result = &v1alpha1.OpsworksRailsAppLayer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("opsworksrailsapplayers").
		Body(opsworksRailsAppLayer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a opsworksRailsAppLayer and updates it. Returns the server's representation of the opsworksRailsAppLayer, and an error, if there is any.
func (c *opsworksRailsAppLayers) Update(opsworksRailsAppLayer *v1alpha1.OpsworksRailsAppLayer) (result *v1alpha1.OpsworksRailsAppLayer, err error) {
	result = &v1alpha1.OpsworksRailsAppLayer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("opsworksrailsapplayers").
		Name(opsworksRailsAppLayer.Name).
		Body(opsworksRailsAppLayer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *opsworksRailsAppLayers) UpdateStatus(opsworksRailsAppLayer *v1alpha1.OpsworksRailsAppLayer) (result *v1alpha1.OpsworksRailsAppLayer, err error) {
	result = &v1alpha1.OpsworksRailsAppLayer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("opsworksrailsapplayers").
		Name(opsworksRailsAppLayer.Name).
		SubResource("status").
		Body(opsworksRailsAppLayer).
		Do().
		Into(result)
	return
}

// Delete takes name of the opsworksRailsAppLayer and deletes it. Returns an error if one occurs.
func (c *opsworksRailsAppLayers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("opsworksrailsapplayers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *opsworksRailsAppLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("opsworksrailsapplayers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched opsworksRailsAppLayer.
func (c *opsworksRailsAppLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksRailsAppLayer, err error) {
	result = &v1alpha1.OpsworksRailsAppLayer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("opsworksrailsapplayers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
