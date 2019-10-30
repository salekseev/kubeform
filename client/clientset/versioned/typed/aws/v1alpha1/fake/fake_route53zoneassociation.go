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

// FakeRoute53ZoneAssociations implements Route53ZoneAssociationInterface
type FakeRoute53ZoneAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var route53zoneassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "route53zoneassociations"}

var route53zoneassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "Route53ZoneAssociation"}

// Get takes name of the route53ZoneAssociation, and returns the corresponding route53ZoneAssociation object, and an error if there is any.
func (c *FakeRoute53ZoneAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.Route53ZoneAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(route53zoneassociationsResource, c.ns, name), &v1alpha1.Route53ZoneAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53ZoneAssociation), err
}

// List takes label and field selectors, and returns the list of Route53ZoneAssociations that match those selectors.
func (c *FakeRoute53ZoneAssociations) List(opts v1.ListOptions) (result *v1alpha1.Route53ZoneAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(route53zoneassociationsResource, route53zoneassociationsKind, c.ns, opts), &v1alpha1.Route53ZoneAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Route53ZoneAssociationList{ListMeta: obj.(*v1alpha1.Route53ZoneAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.Route53ZoneAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested route53ZoneAssociations.
func (c *FakeRoute53ZoneAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(route53zoneassociationsResource, c.ns, opts))

}

// Create takes the representation of a route53ZoneAssociation and creates it.  Returns the server's representation of the route53ZoneAssociation, and an error, if there is any.
func (c *FakeRoute53ZoneAssociations) Create(route53ZoneAssociation *v1alpha1.Route53ZoneAssociation) (result *v1alpha1.Route53ZoneAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(route53zoneassociationsResource, c.ns, route53ZoneAssociation), &v1alpha1.Route53ZoneAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53ZoneAssociation), err
}

// Update takes the representation of a route53ZoneAssociation and updates it. Returns the server's representation of the route53ZoneAssociation, and an error, if there is any.
func (c *FakeRoute53ZoneAssociations) Update(route53ZoneAssociation *v1alpha1.Route53ZoneAssociation) (result *v1alpha1.Route53ZoneAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(route53zoneassociationsResource, c.ns, route53ZoneAssociation), &v1alpha1.Route53ZoneAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53ZoneAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoute53ZoneAssociations) UpdateStatus(route53ZoneAssociation *v1alpha1.Route53ZoneAssociation) (*v1alpha1.Route53ZoneAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(route53zoneassociationsResource, "status", c.ns, route53ZoneAssociation), &v1alpha1.Route53ZoneAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53ZoneAssociation), err
}

// Delete takes name of the route53ZoneAssociation and deletes it. Returns an error if one occurs.
func (c *FakeRoute53ZoneAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(route53zoneassociationsResource, c.ns, name), &v1alpha1.Route53ZoneAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoute53ZoneAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(route53zoneassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.Route53ZoneAssociationList{})
	return err
}

// Patch applies the patch and returns the patched route53ZoneAssociation.
func (c *FakeRoute53ZoneAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Route53ZoneAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(route53zoneassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Route53ZoneAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Route53ZoneAssociation), err
}
