/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/apache/submarine/submarine-cloud-v2/pkg/apis/submarine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSubmarines implements SubmarineInterface
type FakeSubmarines struct {
	Fake *FakeSubmarineV1alpha1
	ns   string
}

var submarinesResource = schema.GroupVersionResource{Group: "submarine.apache.org", Version: "v1alpha1", Resource: "submarines"}

var submarinesKind = schema.GroupVersionKind{Group: "submarine.apache.org", Version: "v1alpha1", Kind: "Submarine"}

// Get takes name of the submarine, and returns the corresponding submarine object, and an error if there is any.
func (c *FakeSubmarines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Submarine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(submarinesResource, c.ns, name), &v1alpha1.Submarine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Submarine), err
}

// List takes label and field selectors, and returns the list of Submarines that match those selectors.
func (c *FakeSubmarines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubmarineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(submarinesResource, submarinesKind, c.ns, opts), &v1alpha1.SubmarineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubmarineList{ListMeta: obj.(*v1alpha1.SubmarineList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubmarineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested submarines.
func (c *FakeSubmarines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(submarinesResource, c.ns, opts))

}

// Create takes the representation of a submarine and creates it.  Returns the server's representation of the submarine, and an error, if there is any.
func (c *FakeSubmarines) Create(ctx context.Context, submarine *v1alpha1.Submarine, opts v1.CreateOptions) (result *v1alpha1.Submarine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(submarinesResource, c.ns, submarine), &v1alpha1.Submarine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Submarine), err
}

// Update takes the representation of a submarine and updates it. Returns the server's representation of the submarine, and an error, if there is any.
func (c *FakeSubmarines) Update(ctx context.Context, submarine *v1alpha1.Submarine, opts v1.UpdateOptions) (result *v1alpha1.Submarine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(submarinesResource, c.ns, submarine), &v1alpha1.Submarine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Submarine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubmarines) UpdateStatus(ctx context.Context, submarine *v1alpha1.Submarine, opts v1.UpdateOptions) (*v1alpha1.Submarine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(submarinesResource, "status", c.ns, submarine), &v1alpha1.Submarine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Submarine), err
}

// Delete takes name of the submarine and deletes it. Returns an error if one occurs.
func (c *FakeSubmarines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(submarinesResource, c.ns, name), &v1alpha1.Submarine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubmarines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(submarinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubmarineList{})
	return err
}

// Patch applies the patch and returns the patched submarine.
func (c *FakeSubmarines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Submarine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(submarinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Submarine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Submarine), err
}
