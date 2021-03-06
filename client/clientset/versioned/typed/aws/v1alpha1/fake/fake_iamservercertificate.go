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

package fake

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIamServerCertificates implements IamServerCertificateInterface
type FakeIamServerCertificates struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iamservercertificatesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iamservercertificates"}

var iamservercertificatesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamServerCertificate"}

// Get takes name of the iamServerCertificate, and returns the corresponding iamServerCertificate object, and an error if there is any.
func (c *FakeIamServerCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.IamServerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamservercertificatesResource, c.ns, name), &v1alpha1.IamServerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamServerCertificate), err
}

// List takes label and field selectors, and returns the list of IamServerCertificates that match those selectors.
func (c *FakeIamServerCertificates) List(opts v1.ListOptions) (result *v1alpha1.IamServerCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamservercertificatesResource, iamservercertificatesKind, c.ns, opts), &v1alpha1.IamServerCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamServerCertificateList{ListMeta: obj.(*v1alpha1.IamServerCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamServerCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamServerCertificates.
func (c *FakeIamServerCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamservercertificatesResource, c.ns, opts))

}

// Create takes the representation of a iamServerCertificate and creates it.  Returns the server's representation of the iamServerCertificate, and an error, if there is any.
func (c *FakeIamServerCertificates) Create(iamServerCertificate *v1alpha1.IamServerCertificate) (result *v1alpha1.IamServerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamservercertificatesResource, c.ns, iamServerCertificate), &v1alpha1.IamServerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamServerCertificate), err
}

// Update takes the representation of a iamServerCertificate and updates it. Returns the server's representation of the iamServerCertificate, and an error, if there is any.
func (c *FakeIamServerCertificates) Update(iamServerCertificate *v1alpha1.IamServerCertificate) (result *v1alpha1.IamServerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamservercertificatesResource, c.ns, iamServerCertificate), &v1alpha1.IamServerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamServerCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamServerCertificates) UpdateStatus(iamServerCertificate *v1alpha1.IamServerCertificate) (*v1alpha1.IamServerCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamservercertificatesResource, "status", c.ns, iamServerCertificate), &v1alpha1.IamServerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamServerCertificate), err
}

// Delete takes name of the iamServerCertificate and deletes it. Returns an error if one occurs.
func (c *FakeIamServerCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iamservercertificatesResource, c.ns, name), &v1alpha1.IamServerCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamServerCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamservercertificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamServerCertificateList{})
	return err
}

// Patch applies the patch and returns the patched iamServerCertificate.
func (c *FakeIamServerCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamServerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamservercertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamServerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamServerCertificate), err
}
