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

// EbsSnapshotCopiesGetter has a method to return a EbsSnapshotCopyInterface.
// A group's client should implement this interface.
type EbsSnapshotCopiesGetter interface {
	EbsSnapshotCopies(namespace string) EbsSnapshotCopyInterface
}

// EbsSnapshotCopyInterface has methods to work with EbsSnapshotCopy resources.
type EbsSnapshotCopyInterface interface {
	Create(*v1alpha1.EbsSnapshotCopy) (*v1alpha1.EbsSnapshotCopy, error)
	Update(*v1alpha1.EbsSnapshotCopy) (*v1alpha1.EbsSnapshotCopy, error)
	UpdateStatus(*v1alpha1.EbsSnapshotCopy) (*v1alpha1.EbsSnapshotCopy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EbsSnapshotCopy, error)
	List(opts v1.ListOptions) (*v1alpha1.EbsSnapshotCopyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EbsSnapshotCopy, err error)
	EbsSnapshotCopyExpansion
}

// ebsSnapshotCopies implements EbsSnapshotCopyInterface
type ebsSnapshotCopies struct {
	client rest.Interface
	ns     string
}

// newEbsSnapshotCopies returns a EbsSnapshotCopies
func newEbsSnapshotCopies(c *AwsV1alpha1Client, namespace string) *ebsSnapshotCopies {
	return &ebsSnapshotCopies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ebsSnapshotCopy, and returns the corresponding ebsSnapshotCopy object, and an error if there is any.
func (c *ebsSnapshotCopies) Get(name string, options v1.GetOptions) (result *v1alpha1.EbsSnapshotCopy, err error) {
	result = &v1alpha1.EbsSnapshotCopy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ebssnapshotcopies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EbsSnapshotCopies that match those selectors.
func (c *ebsSnapshotCopies) List(opts v1.ListOptions) (result *v1alpha1.EbsSnapshotCopyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EbsSnapshotCopyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ebssnapshotcopies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ebsSnapshotCopies.
func (c *ebsSnapshotCopies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ebssnapshotcopies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ebsSnapshotCopy and creates it.  Returns the server's representation of the ebsSnapshotCopy, and an error, if there is any.
func (c *ebsSnapshotCopies) Create(ebsSnapshotCopy *v1alpha1.EbsSnapshotCopy) (result *v1alpha1.EbsSnapshotCopy, err error) {
	result = &v1alpha1.EbsSnapshotCopy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ebssnapshotcopies").
		Body(ebsSnapshotCopy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ebsSnapshotCopy and updates it. Returns the server's representation of the ebsSnapshotCopy, and an error, if there is any.
func (c *ebsSnapshotCopies) Update(ebsSnapshotCopy *v1alpha1.EbsSnapshotCopy) (result *v1alpha1.EbsSnapshotCopy, err error) {
	result = &v1alpha1.EbsSnapshotCopy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ebssnapshotcopies").
		Name(ebsSnapshotCopy.Name).
		Body(ebsSnapshotCopy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ebsSnapshotCopies) UpdateStatus(ebsSnapshotCopy *v1alpha1.EbsSnapshotCopy) (result *v1alpha1.EbsSnapshotCopy, err error) {
	result = &v1alpha1.EbsSnapshotCopy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ebssnapshotcopies").
		Name(ebsSnapshotCopy.Name).
		SubResource("status").
		Body(ebsSnapshotCopy).
		Do().
		Into(result)
	return
}

// Delete takes name of the ebsSnapshotCopy and deletes it. Returns an error if one occurs.
func (c *ebsSnapshotCopies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ebssnapshotcopies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ebsSnapshotCopies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ebssnapshotcopies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ebsSnapshotCopy.
func (c *ebsSnapshotCopies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EbsSnapshotCopy, err error) {
	result = &v1alpha1.EbsSnapshotCopy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ebssnapshotcopies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
