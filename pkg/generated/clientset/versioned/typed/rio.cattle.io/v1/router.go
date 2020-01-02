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

package v1

import (
	"time"

	v1 "github.com/rancher/rio/pkg/apis/rio.cattle.io/v1"
	scheme "github.com/rancher/rio/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RoutersGetter has a method to return a RouterInterface.
// A group's client should implement this interface.
type RoutersGetter interface {
	Routers(namespace string) RouterInterface
}

// RouterInterface has methods to work with Router resources.
type RouterInterface interface {
	Create(*v1.Router) (*v1.Router, error)
	Update(*v1.Router) (*v1.Router, error)
	UpdateStatus(*v1.Router) (*v1.Router, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Router, error)
	List(opts metav1.ListOptions) (*v1.RouterList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Router, err error)
	RouterExpansion
}

// routers implements RouterInterface
type routers struct {
	client rest.Interface
	ns     string
}

// newRouters returns a Routers
func newRouters(c *RioV1Client, namespace string) *routers {
	return &routers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the router, and returns the corresponding router object, and an error if there is any.
func (c *routers) Get(name string, options metav1.GetOptions) (result *v1.Router, err error) {
	result = &v1.Router{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("routers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Routers that match those selectors.
func (c *routers) List(opts metav1.ListOptions) (result *v1.RouterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.RouterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("routers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested routers.
func (c *routers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("routers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a router and creates it.  Returns the server's representation of the router, and an error, if there is any.
func (c *routers) Create(router *v1.Router) (result *v1.Router, err error) {
	result = &v1.Router{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("routers").
		Body(router).
		Do().
		Into(result)
	return
}

// Update takes the representation of a router and updates it. Returns the server's representation of the router, and an error, if there is any.
func (c *routers) Update(router *v1.Router) (result *v1.Router, err error) {
	result = &v1.Router{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("routers").
		Name(router.Name).
		Body(router).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *routers) UpdateStatus(router *v1.Router) (result *v1.Router, err error) {
	result = &v1.Router{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("routers").
		Name(router.Name).
		SubResource("status").
		Body(router).
		Do().
		Into(result)
	return
}

// Delete takes name of the router and deletes it. Returns an error if one occurs.
func (c *routers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("routers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *routers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("routers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched router.
func (c *routers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Router, err error) {
	result = &v1.Router{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("routers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
