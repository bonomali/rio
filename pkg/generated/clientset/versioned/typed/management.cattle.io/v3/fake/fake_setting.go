/*
Copyright 2020 Rancher Labs.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	v3 "github.com/rancher/rio/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSettings implements SettingInterface
type FakeSettings struct {
	Fake *FakeManagementV3
}

var settingsResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "settings"}

var settingsKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Setting"}

// Get takes name of the setting, and returns the corresponding setting object, and an error if there is any.
func (c *FakeSettings) Get(name string, options v1.GetOptions) (result *v3.Setting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(settingsResource, name), &v3.Setting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.Setting), err
}

// List takes label and field selectors, and returns the list of Settings that match those selectors.
func (c *FakeSettings) List(opts v1.ListOptions) (result *v3.SettingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(settingsResource, settingsKind, opts), &v3.SettingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.SettingList{ListMeta: obj.(*v3.SettingList).ListMeta}
	for _, item := range obj.(*v3.SettingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested settings.
func (c *FakeSettings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(settingsResource, opts))
}

// Create takes the representation of a setting and creates it.  Returns the server's representation of the setting, and an error, if there is any.
func (c *FakeSettings) Create(setting *v3.Setting) (result *v3.Setting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(settingsResource, setting), &v3.Setting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.Setting), err
}

// Update takes the representation of a setting and updates it. Returns the server's representation of the setting, and an error, if there is any.
func (c *FakeSettings) Update(setting *v3.Setting) (result *v3.Setting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(settingsResource, setting), &v3.Setting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.Setting), err
}

// Delete takes name of the setting and deletes it. Returns an error if one occurs.
func (c *FakeSettings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(settingsResource, name), &v3.Setting{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSettings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(settingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v3.SettingList{})
	return err
}

// Patch applies the patch and returns the patched setting.
func (c *FakeSettings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.Setting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(settingsResource, name, pt, data, subresources...), &v3.Setting{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.Setting), err
}
