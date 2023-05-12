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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/k3s-io/k3s/pkg/apis/k3s.cattle.io/v1"
	scheme "github.com/k3s-io/k3s/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AddonsGetter has a method to return a AddonInterface.
// A group's client should implement this interface.
type AddonsGetter interface {
	Addons(namespace string) AddonInterface
}

// AddonInterface has methods to work with Addon resources.
type AddonInterface interface {
	Create(ctx context.Context, addon *v1.Addon, opts metav1.CreateOptions) (*v1.Addon, error)
	Update(ctx context.Context, addon *v1.Addon, opts metav1.UpdateOptions) (*v1.Addon, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Addon, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AddonList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Addon, err error)
	AddonExpansion
}

// addons implements AddonInterface
type addons struct {
	client rest.Interface
	ns     string
}

// newAddons returns a Addons
func newAddons(c *K3sV1Client, namespace string) *addons {
	return &addons{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the addon, and returns the corresponding addon object, and an error if there is any.
func (c *addons) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Addon, err error) {
	result = &v1.Addon{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("addons").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Addons that match those selectors.
func (c *addons) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AddonList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AddonList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("addons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested addons.
func (c *addons) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("addons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a addon and creates it.  Returns the server's representation of the addon, and an error, if there is any.
func (c *addons) Create(ctx context.Context, addon *v1.Addon, opts metav1.CreateOptions) (result *v1.Addon, err error) {
	result = &v1.Addon{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("addons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(addon).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a addon and updates it. Returns the server's representation of the addon, and an error, if there is any.
func (c *addons) Update(ctx context.Context, addon *v1.Addon, opts metav1.UpdateOptions) (result *v1.Addon, err error) {
	result = &v1.Addon{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("addons").
		Name(addon.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(addon).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the addon and deletes it. Returns an error if one occurs.
func (c *addons) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("addons").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *addons) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("addons").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched addon.
func (c *addons) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Addon, err error) {
	result = &v1.Addon{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("addons").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
