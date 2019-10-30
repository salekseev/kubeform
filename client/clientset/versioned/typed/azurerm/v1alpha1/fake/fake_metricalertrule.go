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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMetricAlertrules implements MetricAlertruleInterface
type FakeMetricAlertrules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var metricalertrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "metricalertrules"}

var metricalertrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MetricAlertrule"}

// Get takes name of the metricAlertrule, and returns the corresponding metricAlertrule object, and an error if there is any.
func (c *FakeMetricAlertrules) Get(name string, options v1.GetOptions) (result *v1alpha1.MetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(metricalertrulesResource, c.ns, name), &v1alpha1.MetricAlertrule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricAlertrule), err
}

// List takes label and field selectors, and returns the list of MetricAlertrules that match those selectors.
func (c *FakeMetricAlertrules) List(opts v1.ListOptions) (result *v1alpha1.MetricAlertruleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(metricalertrulesResource, metricalertrulesKind, c.ns, opts), &v1alpha1.MetricAlertruleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MetricAlertruleList{ListMeta: obj.(*v1alpha1.MetricAlertruleList).ListMeta}
	for _, item := range obj.(*v1alpha1.MetricAlertruleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested metricAlertrules.
func (c *FakeMetricAlertrules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(metricalertrulesResource, c.ns, opts))

}

// Create takes the representation of a metricAlertrule and creates it.  Returns the server's representation of the metricAlertrule, and an error, if there is any.
func (c *FakeMetricAlertrules) Create(metricAlertrule *v1alpha1.MetricAlertrule) (result *v1alpha1.MetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(metricalertrulesResource, c.ns, metricAlertrule), &v1alpha1.MetricAlertrule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricAlertrule), err
}

// Update takes the representation of a metricAlertrule and updates it. Returns the server's representation of the metricAlertrule, and an error, if there is any.
func (c *FakeMetricAlertrules) Update(metricAlertrule *v1alpha1.MetricAlertrule) (result *v1alpha1.MetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(metricalertrulesResource, c.ns, metricAlertrule), &v1alpha1.MetricAlertrule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricAlertrule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMetricAlertrules) UpdateStatus(metricAlertrule *v1alpha1.MetricAlertrule) (*v1alpha1.MetricAlertrule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(metricalertrulesResource, "status", c.ns, metricAlertrule), &v1alpha1.MetricAlertrule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricAlertrule), err
}

// Delete takes name of the metricAlertrule and deletes it. Returns an error if one occurs.
func (c *FakeMetricAlertrules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(metricalertrulesResource, c.ns, name), &v1alpha1.MetricAlertrule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMetricAlertrules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(metricalertrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MetricAlertruleList{})
	return err
}

// Patch applies the patch and returns the patched metricAlertrule.
func (c *FakeMetricAlertrules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MetricAlertrule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(metricalertrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MetricAlertrule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricAlertrule), err
}
