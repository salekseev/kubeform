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

// FakeGameliftGameSessionQueues implements GameliftGameSessionQueueInterface
type FakeGameliftGameSessionQueues struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var gameliftgamesessionqueuesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "gameliftgamesessionqueues"}

var gameliftgamesessionqueuesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "GameliftGameSessionQueue"}

// Get takes name of the gameliftGameSessionQueue, and returns the corresponding gameliftGameSessionQueue object, and an error if there is any.
func (c *FakeGameliftGameSessionQueues) Get(name string, options v1.GetOptions) (result *v1alpha1.GameliftGameSessionQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gameliftgamesessionqueuesResource, c.ns, name), &v1alpha1.GameliftGameSessionQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameliftGameSessionQueue), err
}

// List takes label and field selectors, and returns the list of GameliftGameSessionQueues that match those selectors.
func (c *FakeGameliftGameSessionQueues) List(opts v1.ListOptions) (result *v1alpha1.GameliftGameSessionQueueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gameliftgamesessionqueuesResource, gameliftgamesessionqueuesKind, c.ns, opts), &v1alpha1.GameliftGameSessionQueueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GameliftGameSessionQueueList{ListMeta: obj.(*v1alpha1.GameliftGameSessionQueueList).ListMeta}
	for _, item := range obj.(*v1alpha1.GameliftGameSessionQueueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gameliftGameSessionQueues.
func (c *FakeGameliftGameSessionQueues) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gameliftgamesessionqueuesResource, c.ns, opts))

}

// Create takes the representation of a gameliftGameSessionQueue and creates it.  Returns the server's representation of the gameliftGameSessionQueue, and an error, if there is any.
func (c *FakeGameliftGameSessionQueues) Create(gameliftGameSessionQueue *v1alpha1.GameliftGameSessionQueue) (result *v1alpha1.GameliftGameSessionQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gameliftgamesessionqueuesResource, c.ns, gameliftGameSessionQueue), &v1alpha1.GameliftGameSessionQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameliftGameSessionQueue), err
}

// Update takes the representation of a gameliftGameSessionQueue and updates it. Returns the server's representation of the gameliftGameSessionQueue, and an error, if there is any.
func (c *FakeGameliftGameSessionQueues) Update(gameliftGameSessionQueue *v1alpha1.GameliftGameSessionQueue) (result *v1alpha1.GameliftGameSessionQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gameliftgamesessionqueuesResource, c.ns, gameliftGameSessionQueue), &v1alpha1.GameliftGameSessionQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameliftGameSessionQueue), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGameliftGameSessionQueues) UpdateStatus(gameliftGameSessionQueue *v1alpha1.GameliftGameSessionQueue) (*v1alpha1.GameliftGameSessionQueue, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gameliftgamesessionqueuesResource, "status", c.ns, gameliftGameSessionQueue), &v1alpha1.GameliftGameSessionQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameliftGameSessionQueue), err
}

// Delete takes name of the gameliftGameSessionQueue and deletes it. Returns an error if one occurs.
func (c *FakeGameliftGameSessionQueues) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gameliftgamesessionqueuesResource, c.ns, name), &v1alpha1.GameliftGameSessionQueue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGameliftGameSessionQueues) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gameliftgamesessionqueuesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GameliftGameSessionQueueList{})
	return err
}

// Patch applies the patch and returns the patched gameliftGameSessionQueue.
func (c *FakeGameliftGameSessionQueues) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GameliftGameSessionQueue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gameliftgamesessionqueuesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GameliftGameSessionQueue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GameliftGameSessionQueue), err
}
