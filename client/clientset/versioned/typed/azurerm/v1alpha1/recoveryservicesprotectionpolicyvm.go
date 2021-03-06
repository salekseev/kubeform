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

// RecoveryServicesProtectionPolicyVmsGetter has a method to return a RecoveryServicesProtectionPolicyVmInterface.
// A group's client should implement this interface.
type RecoveryServicesProtectionPolicyVmsGetter interface {
	RecoveryServicesProtectionPolicyVms(namespace string) RecoveryServicesProtectionPolicyVmInterface
}

// RecoveryServicesProtectionPolicyVmInterface has methods to work with RecoveryServicesProtectionPolicyVm resources.
type RecoveryServicesProtectionPolicyVmInterface interface {
	Create(*v1alpha1.RecoveryServicesProtectionPolicyVm) (*v1alpha1.RecoveryServicesProtectionPolicyVm, error)
	Update(*v1alpha1.RecoveryServicesProtectionPolicyVm) (*v1alpha1.RecoveryServicesProtectionPolicyVm, error)
	UpdateStatus(*v1alpha1.RecoveryServicesProtectionPolicyVm) (*v1alpha1.RecoveryServicesProtectionPolicyVm, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RecoveryServicesProtectionPolicyVm, error)
	List(opts v1.ListOptions) (*v1alpha1.RecoveryServicesProtectionPolicyVmList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RecoveryServicesProtectionPolicyVm, err error)
	RecoveryServicesProtectionPolicyVmExpansion
}

// recoveryServicesProtectionPolicyVms implements RecoveryServicesProtectionPolicyVmInterface
type recoveryServicesProtectionPolicyVms struct {
	client rest.Interface
	ns     string
}

// newRecoveryServicesProtectionPolicyVms returns a RecoveryServicesProtectionPolicyVms
func newRecoveryServicesProtectionPolicyVms(c *AzurermV1alpha1Client, namespace string) *recoveryServicesProtectionPolicyVms {
	return &recoveryServicesProtectionPolicyVms{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the recoveryServicesProtectionPolicyVm, and returns the corresponding recoveryServicesProtectionPolicyVm object, and an error if there is any.
func (c *recoveryServicesProtectionPolicyVms) Get(name string, options v1.GetOptions) (result *v1alpha1.RecoveryServicesProtectionPolicyVm, err error) {
	result = &v1alpha1.RecoveryServicesProtectionPolicyVm{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("recoveryservicesprotectionpolicyvms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RecoveryServicesProtectionPolicyVms that match those selectors.
func (c *recoveryServicesProtectionPolicyVms) List(opts v1.ListOptions) (result *v1alpha1.RecoveryServicesProtectionPolicyVmList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RecoveryServicesProtectionPolicyVmList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("recoveryservicesprotectionpolicyvms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested recoveryServicesProtectionPolicyVms.
func (c *recoveryServicesProtectionPolicyVms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("recoveryservicesprotectionpolicyvms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a recoveryServicesProtectionPolicyVm and creates it.  Returns the server's representation of the recoveryServicesProtectionPolicyVm, and an error, if there is any.
func (c *recoveryServicesProtectionPolicyVms) Create(recoveryServicesProtectionPolicyVm *v1alpha1.RecoveryServicesProtectionPolicyVm) (result *v1alpha1.RecoveryServicesProtectionPolicyVm, err error) {
	result = &v1alpha1.RecoveryServicesProtectionPolicyVm{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("recoveryservicesprotectionpolicyvms").
		Body(recoveryServicesProtectionPolicyVm).
		Do().
		Into(result)
	return
}

// Update takes the representation of a recoveryServicesProtectionPolicyVm and updates it. Returns the server's representation of the recoveryServicesProtectionPolicyVm, and an error, if there is any.
func (c *recoveryServicesProtectionPolicyVms) Update(recoveryServicesProtectionPolicyVm *v1alpha1.RecoveryServicesProtectionPolicyVm) (result *v1alpha1.RecoveryServicesProtectionPolicyVm, err error) {
	result = &v1alpha1.RecoveryServicesProtectionPolicyVm{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("recoveryservicesprotectionpolicyvms").
		Name(recoveryServicesProtectionPolicyVm.Name).
		Body(recoveryServicesProtectionPolicyVm).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *recoveryServicesProtectionPolicyVms) UpdateStatus(recoveryServicesProtectionPolicyVm *v1alpha1.RecoveryServicesProtectionPolicyVm) (result *v1alpha1.RecoveryServicesProtectionPolicyVm, err error) {
	result = &v1alpha1.RecoveryServicesProtectionPolicyVm{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("recoveryservicesprotectionpolicyvms").
		Name(recoveryServicesProtectionPolicyVm.Name).
		SubResource("status").
		Body(recoveryServicesProtectionPolicyVm).
		Do().
		Into(result)
	return
}

// Delete takes name of the recoveryServicesProtectionPolicyVm and deletes it. Returns an error if one occurs.
func (c *recoveryServicesProtectionPolicyVms) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("recoveryservicesprotectionpolicyvms").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *recoveryServicesProtectionPolicyVms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("recoveryservicesprotectionpolicyvms").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched recoveryServicesProtectionPolicyVm.
func (c *recoveryServicesProtectionPolicyVms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RecoveryServicesProtectionPolicyVm, err error) {
	result = &v1alpha1.RecoveryServicesProtectionPolicyVm{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("recoveryservicesprotectionpolicyvms").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
