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

// SqlDatabasesGetter has a method to return a SqlDatabaseInterface.
// A group's client should implement this interface.
type SqlDatabasesGetter interface {
	SqlDatabases(namespace string) SqlDatabaseInterface
}

// SqlDatabaseInterface has methods to work with SqlDatabase resources.
type SqlDatabaseInterface interface {
	Create(*v1alpha1.SqlDatabase) (*v1alpha1.SqlDatabase, error)
	Update(*v1alpha1.SqlDatabase) (*v1alpha1.SqlDatabase, error)
	UpdateStatus(*v1alpha1.SqlDatabase) (*v1alpha1.SqlDatabase, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SqlDatabase, error)
	List(opts v1.ListOptions) (*v1alpha1.SqlDatabaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqlDatabase, err error)
	SqlDatabaseExpansion
}

// sqlDatabases implements SqlDatabaseInterface
type sqlDatabases struct {
	client rest.Interface
	ns     string
}

// newSqlDatabases returns a SqlDatabases
func newSqlDatabases(c *GoogleV1alpha1Client, namespace string) *sqlDatabases {
	return &sqlDatabases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sqlDatabase, and returns the corresponding sqlDatabase object, and an error if there is any.
func (c *sqlDatabases) Get(name string, options v1.GetOptions) (result *v1alpha1.SqlDatabase, err error) {
	result = &v1alpha1.SqlDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqldatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SqlDatabases that match those selectors.
func (c *sqlDatabases) List(opts v1.ListOptions) (result *v1alpha1.SqlDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SqlDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqldatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sqlDatabases.
func (c *sqlDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sqldatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sqlDatabase and creates it.  Returns the server's representation of the sqlDatabase, and an error, if there is any.
func (c *sqlDatabases) Create(sqlDatabase *v1alpha1.SqlDatabase) (result *v1alpha1.SqlDatabase, err error) {
	result = &v1alpha1.SqlDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sqldatabases").
		Body(sqlDatabase).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sqlDatabase and updates it. Returns the server's representation of the sqlDatabase, and an error, if there is any.
func (c *sqlDatabases) Update(sqlDatabase *v1alpha1.SqlDatabase) (result *v1alpha1.SqlDatabase, err error) {
	result = &v1alpha1.SqlDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sqldatabases").
		Name(sqlDatabase.Name).
		Body(sqlDatabase).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *sqlDatabases) UpdateStatus(sqlDatabase *v1alpha1.SqlDatabase) (result *v1alpha1.SqlDatabase, err error) {
	result = &v1alpha1.SqlDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sqldatabases").
		Name(sqlDatabase.Name).
		SubResource("status").
		Body(sqlDatabase).
		Do().
		Into(result)
	return
}

// Delete takes name of the sqlDatabase and deletes it. Returns an error if one occurs.
func (c *sqlDatabases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqldatabases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sqlDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqldatabases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sqlDatabase.
func (c *sqlDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqlDatabase, err error) {
	result = &v1alpha1.SqlDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sqldatabases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
