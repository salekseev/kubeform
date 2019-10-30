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

// FakeGlueClassifiers implements GlueClassifierInterface
type FakeGlueClassifiers struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var glueclassifiersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "glueclassifiers"}

var glueclassifiersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "GlueClassifier"}

// Get takes name of the glueClassifier, and returns the corresponding glueClassifier object, and an error if there is any.
func (c *FakeGlueClassifiers) Get(name string, options v1.GetOptions) (result *v1alpha1.GlueClassifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(glueclassifiersResource, c.ns, name), &v1alpha1.GlueClassifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueClassifier), err
}

// List takes label and field selectors, and returns the list of GlueClassifiers that match those selectors.
func (c *FakeGlueClassifiers) List(opts v1.ListOptions) (result *v1alpha1.GlueClassifierList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(glueclassifiersResource, glueclassifiersKind, c.ns, opts), &v1alpha1.GlueClassifierList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GlueClassifierList{ListMeta: obj.(*v1alpha1.GlueClassifierList).ListMeta}
	for _, item := range obj.(*v1alpha1.GlueClassifierList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested glueClassifiers.
func (c *FakeGlueClassifiers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(glueclassifiersResource, c.ns, opts))

}

// Create takes the representation of a glueClassifier and creates it.  Returns the server's representation of the glueClassifier, and an error, if there is any.
func (c *FakeGlueClassifiers) Create(glueClassifier *v1alpha1.GlueClassifier) (result *v1alpha1.GlueClassifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(glueclassifiersResource, c.ns, glueClassifier), &v1alpha1.GlueClassifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueClassifier), err
}

// Update takes the representation of a glueClassifier and updates it. Returns the server's representation of the glueClassifier, and an error, if there is any.
func (c *FakeGlueClassifiers) Update(glueClassifier *v1alpha1.GlueClassifier) (result *v1alpha1.GlueClassifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(glueclassifiersResource, c.ns, glueClassifier), &v1alpha1.GlueClassifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueClassifier), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlueClassifiers) UpdateStatus(glueClassifier *v1alpha1.GlueClassifier) (*v1alpha1.GlueClassifier, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(glueclassifiersResource, "status", c.ns, glueClassifier), &v1alpha1.GlueClassifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueClassifier), err
}

// Delete takes name of the glueClassifier and deletes it. Returns an error if one occurs.
func (c *FakeGlueClassifiers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(glueclassifiersResource, c.ns, name), &v1alpha1.GlueClassifier{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlueClassifiers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(glueclassifiersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GlueClassifierList{})
	return err
}

// Patch applies the patch and returns the patched glueClassifier.
func (c *FakeGlueClassifiers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GlueClassifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(glueclassifiersResource, c.ns, name, pt, data, subresources...), &v1alpha1.GlueClassifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GlueClassifier), err
}
