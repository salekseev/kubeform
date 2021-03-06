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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRedisInstances implements RedisInstanceInterface
type FakeRedisInstances struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var redisinstancesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "redisinstances"}

var redisinstancesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "RedisInstance"}

// Get takes name of the redisInstance, and returns the corresponding redisInstance object, and an error if there is any.
func (c *FakeRedisInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.RedisInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redisinstancesResource, c.ns, name), &v1alpha1.RedisInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisInstance), err
}

// List takes label and field selectors, and returns the list of RedisInstances that match those selectors.
func (c *FakeRedisInstances) List(opts v1.ListOptions) (result *v1alpha1.RedisInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redisinstancesResource, redisinstancesKind, c.ns, opts), &v1alpha1.RedisInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RedisInstanceList{ListMeta: obj.(*v1alpha1.RedisInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.RedisInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redisInstances.
func (c *FakeRedisInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redisinstancesResource, c.ns, opts))

}

// Create takes the representation of a redisInstance and creates it.  Returns the server's representation of the redisInstance, and an error, if there is any.
func (c *FakeRedisInstances) Create(redisInstance *v1alpha1.RedisInstance) (result *v1alpha1.RedisInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redisinstancesResource, c.ns, redisInstance), &v1alpha1.RedisInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisInstance), err
}

// Update takes the representation of a redisInstance and updates it. Returns the server's representation of the redisInstance, and an error, if there is any.
func (c *FakeRedisInstances) Update(redisInstance *v1alpha1.RedisInstance) (result *v1alpha1.RedisInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redisinstancesResource, c.ns, redisInstance), &v1alpha1.RedisInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRedisInstances) UpdateStatus(redisInstance *v1alpha1.RedisInstance) (*v1alpha1.RedisInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(redisinstancesResource, "status", c.ns, redisInstance), &v1alpha1.RedisInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisInstance), err
}

// Delete takes name of the redisInstance and deletes it. Returns an error if one occurs.
func (c *FakeRedisInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(redisinstancesResource, c.ns, name), &v1alpha1.RedisInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedisInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redisinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RedisInstanceList{})
	return err
}

// Patch applies the patch and returns the patched redisInstance.
func (c *FakeRedisInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redisinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RedisInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedisInstance), err
}
