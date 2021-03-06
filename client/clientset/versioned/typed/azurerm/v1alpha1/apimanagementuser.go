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

// ApiManagementUsersGetter has a method to return a ApiManagementUserInterface.
// A group's client should implement this interface.
type ApiManagementUsersGetter interface {
	ApiManagementUsers(namespace string) ApiManagementUserInterface
}

// ApiManagementUserInterface has methods to work with ApiManagementUser resources.
type ApiManagementUserInterface interface {
	Create(*v1alpha1.ApiManagementUser) (*v1alpha1.ApiManagementUser, error)
	Update(*v1alpha1.ApiManagementUser) (*v1alpha1.ApiManagementUser, error)
	UpdateStatus(*v1alpha1.ApiManagementUser) (*v1alpha1.ApiManagementUser, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiManagementUser, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiManagementUserList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementUser, err error)
	ApiManagementUserExpansion
}

// apiManagementUsers implements ApiManagementUserInterface
type apiManagementUsers struct {
	client rest.Interface
	ns     string
}

// newApiManagementUsers returns a ApiManagementUsers
func newApiManagementUsers(c *AzurermV1alpha1Client, namespace string) *apiManagementUsers {
	return &apiManagementUsers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiManagementUser, and returns the corresponding apiManagementUser object, and an error if there is any.
func (c *apiManagementUsers) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiManagementUser, err error) {
	result = &v1alpha1.ApiManagementUser{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiManagementUsers that match those selectors.
func (c *apiManagementUsers) List(opts v1.ListOptions) (result *v1alpha1.ApiManagementUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiManagementUserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiManagementUsers.
func (c *apiManagementUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apimanagementusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiManagementUser and creates it.  Returns the server's representation of the apiManagementUser, and an error, if there is any.
func (c *apiManagementUsers) Create(apiManagementUser *v1alpha1.ApiManagementUser) (result *v1alpha1.ApiManagementUser, err error) {
	result = &v1alpha1.ApiManagementUser{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apimanagementusers").
		Body(apiManagementUser).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiManagementUser and updates it. Returns the server's representation of the apiManagementUser, and an error, if there is any.
func (c *apiManagementUsers) Update(apiManagementUser *v1alpha1.ApiManagementUser) (result *v1alpha1.ApiManagementUser, err error) {
	result = &v1alpha1.ApiManagementUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementusers").
		Name(apiManagementUser.Name).
		Body(apiManagementUser).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiManagementUsers) UpdateStatus(apiManagementUser *v1alpha1.ApiManagementUser) (result *v1alpha1.ApiManagementUser, err error) {
	result = &v1alpha1.ApiManagementUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apimanagementusers").
		Name(apiManagementUser.Name).
		SubResource("status").
		Body(apiManagementUser).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiManagementUser and deletes it. Returns an error if one occurs.
func (c *apiManagementUsers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementusers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiManagementUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apimanagementusers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiManagementUser.
func (c *apiManagementUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiManagementUser, err error) {
	result = &v1alpha1.ApiManagementUser{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apimanagementusers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
