/*
Copyright The Kubernetes Authors.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/kwok/pkg/apis/v1alpha1"
)

// FakePortForwards implements PortForwardInterface
type FakePortForwards struct {
	Fake *FakeKwokV1alpha1
	ns   string
}

var portforwardsResource = v1alpha1.SchemeGroupVersion.WithResource("portforwards")

var portforwardsKind = v1alpha1.SchemeGroupVersion.WithKind("PortForward")

// Get takes name of the portForward, and returns the corresponding portForward object, and an error if there is any.
func (c *FakePortForwards) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PortForward, err error) {
	emptyResult := &v1alpha1.PortForward{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(portforwardsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PortForward), err
}

// List takes label and field selectors, and returns the list of PortForwards that match those selectors.
func (c *FakePortForwards) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PortForwardList, err error) {
	emptyResult := &v1alpha1.PortForwardList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(portforwardsResource, portforwardsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PortForwardList{ListMeta: obj.(*v1alpha1.PortForwardList).ListMeta}
	for _, item := range obj.(*v1alpha1.PortForwardList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested portForwards.
func (c *FakePortForwards) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(portforwardsResource, c.ns, opts))

}

// Create takes the representation of a portForward and creates it.  Returns the server's representation of the portForward, and an error, if there is any.
func (c *FakePortForwards) Create(ctx context.Context, portForward *v1alpha1.PortForward, opts v1.CreateOptions) (result *v1alpha1.PortForward, err error) {
	emptyResult := &v1alpha1.PortForward{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(portforwardsResource, c.ns, portForward, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PortForward), err
}

// Update takes the representation of a portForward and updates it. Returns the server's representation of the portForward, and an error, if there is any.
func (c *FakePortForwards) Update(ctx context.Context, portForward *v1alpha1.PortForward, opts v1.UpdateOptions) (result *v1alpha1.PortForward, err error) {
	emptyResult := &v1alpha1.PortForward{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(portforwardsResource, c.ns, portForward, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PortForward), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePortForwards) UpdateStatus(ctx context.Context, portForward *v1alpha1.PortForward, opts v1.UpdateOptions) (result *v1alpha1.PortForward, err error) {
	emptyResult := &v1alpha1.PortForward{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(portforwardsResource, "status", c.ns, portForward, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PortForward), err
}

// Delete takes name of the portForward and deletes it. Returns an error if one occurs.
func (c *FakePortForwards) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(portforwardsResource, c.ns, name, opts), &v1alpha1.PortForward{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePortForwards) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(portforwardsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PortForwardList{})
	return err
}

// Patch applies the patch and returns the patched portForward.
func (c *FakePortForwards) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PortForward, err error) {
	emptyResult := &v1alpha1.PortForward{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(portforwardsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.PortForward), err
}
