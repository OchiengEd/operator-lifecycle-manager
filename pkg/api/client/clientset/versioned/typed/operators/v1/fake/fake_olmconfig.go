/*
Copyright Red Hat, Inc.

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

	operatorsv1 "github.com/operator-framework/api/pkg/operators/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOLMConfigs implements OLMConfigInterface
type FakeOLMConfigs struct {
	Fake *FakeOperatorsV1
}

var olmconfigsResource = schema.GroupVersionResource{Group: "operators.coreos.com", Version: "v1", Resource: "olmconfigs"}

var olmconfigsKind = schema.GroupVersionKind{Group: "operators.coreos.com", Version: "v1", Kind: "OLMConfig"}

// Get takes name of the oLMConfig, and returns the corresponding oLMConfig object, and an error if there is any.
func (c *FakeOLMConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorsv1.OLMConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(olmconfigsResource, name), &operatorsv1.OLMConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OLMConfig), err
}

// List takes label and field selectors, and returns the list of OLMConfigs that match those selectors.
func (c *FakeOLMConfigs) List(ctx context.Context, opts v1.ListOptions) (result *operatorsv1.OLMConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(olmconfigsResource, olmconfigsKind, opts), &operatorsv1.OLMConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorsv1.OLMConfigList{ListMeta: obj.(*operatorsv1.OLMConfigList).ListMeta}
	for _, item := range obj.(*operatorsv1.OLMConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oLMConfigs.
func (c *FakeOLMConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(olmconfigsResource, opts))
}

// Create takes the representation of a oLMConfig and creates it.  Returns the server's representation of the oLMConfig, and an error, if there is any.
func (c *FakeOLMConfigs) Create(ctx context.Context, oLMConfig *operatorsv1.OLMConfig, opts v1.CreateOptions) (result *operatorsv1.OLMConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(olmconfigsResource, oLMConfig), &operatorsv1.OLMConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OLMConfig), err
}

// Update takes the representation of a oLMConfig and updates it. Returns the server's representation of the oLMConfig, and an error, if there is any.
func (c *FakeOLMConfigs) Update(ctx context.Context, oLMConfig *operatorsv1.OLMConfig, opts v1.UpdateOptions) (result *operatorsv1.OLMConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(olmconfigsResource, oLMConfig), &operatorsv1.OLMConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OLMConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOLMConfigs) UpdateStatus(ctx context.Context, oLMConfig *operatorsv1.OLMConfig, opts v1.UpdateOptions) (*operatorsv1.OLMConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(olmconfigsResource, "status", oLMConfig), &operatorsv1.OLMConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OLMConfig), err
}

// Delete takes name of the oLMConfig and deletes it. Returns an error if one occurs.
func (c *FakeOLMConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(olmconfigsResource, name), &operatorsv1.OLMConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOLMConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(olmconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorsv1.OLMConfigList{})
	return err
}

// Patch applies the patch and returns the patched oLMConfig.
func (c *FakeOLMConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorsv1.OLMConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(olmconfigsResource, name, pt, data, subresources...), &operatorsv1.OLMConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorsv1.OLMConfig), err
}
