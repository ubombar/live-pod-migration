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

	v1alpha1 "github.com/ubombar/live-pod-migration/pkg/apis/livepodmigration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeColdMigrations implements ColdMigrationInterface
type FakeColdMigrations struct {
	Fake *FakeLivepodmigrationV1alpha1
	ns   string
}

var coldmigrationsResource = schema.GroupVersionResource{Group: "livepodmigration.edgenet.io", Version: "v1alpha1", Resource: "coldmigrations"}

var coldmigrationsKind = schema.GroupVersionKind{Group: "livepodmigration.edgenet.io", Version: "v1alpha1", Kind: "ColdMigration"}

// Get takes name of the coldMigration, and returns the corresponding coldMigration object, and an error if there is any.
func (c *FakeColdMigrations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ColdMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(coldmigrationsResource, c.ns, name), &v1alpha1.ColdMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ColdMigration), err
}

// List takes label and field selectors, and returns the list of ColdMigrations that match those selectors.
func (c *FakeColdMigrations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ColdMigrationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(coldmigrationsResource, coldmigrationsKind, c.ns, opts), &v1alpha1.ColdMigrationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ColdMigrationList{ListMeta: obj.(*v1alpha1.ColdMigrationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ColdMigrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested coldMigrations.
func (c *FakeColdMigrations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(coldmigrationsResource, c.ns, opts))

}

// Create takes the representation of a coldMigration and creates it.  Returns the server's representation of the coldMigration, and an error, if there is any.
func (c *FakeColdMigrations) Create(ctx context.Context, coldMigration *v1alpha1.ColdMigration, opts v1.CreateOptions) (result *v1alpha1.ColdMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(coldmigrationsResource, c.ns, coldMigration), &v1alpha1.ColdMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ColdMigration), err
}

// Update takes the representation of a coldMigration and updates it. Returns the server's representation of the coldMigration, and an error, if there is any.
func (c *FakeColdMigrations) Update(ctx context.Context, coldMigration *v1alpha1.ColdMigration, opts v1.UpdateOptions) (result *v1alpha1.ColdMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(coldmigrationsResource, c.ns, coldMigration), &v1alpha1.ColdMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ColdMigration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeColdMigrations) UpdateStatus(ctx context.Context, coldMigration *v1alpha1.ColdMigration, opts v1.UpdateOptions) (*v1alpha1.ColdMigration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(coldmigrationsResource, "status", c.ns, coldMigration), &v1alpha1.ColdMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ColdMigration), err
}

// Delete takes name of the coldMigration and deletes it. Returns an error if one occurs.
func (c *FakeColdMigrations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(coldmigrationsResource, c.ns, name, opts), &v1alpha1.ColdMigration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeColdMigrations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(coldmigrationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ColdMigrationList{})
	return err
}

// Patch applies the patch and returns the patched coldMigration.
func (c *FakeColdMigrations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ColdMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(coldmigrationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ColdMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ColdMigration), err
}