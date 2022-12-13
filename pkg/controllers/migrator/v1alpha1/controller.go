// /*
// Copyright 2017 The Kubernetes Authors.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

package v1alpha1

import (
	"fmt"
	"time"

	clientset "github.com/ubombar/live-pod-migration/pkg/generated/clientset/versioned"
	livepodmigrationscheme "github.com/ubombar/live-pod-migration/pkg/generated/clientset/versioned/scheme"
	informers "github.com/ubombar/live-pod-migration/pkg/generated/informers/externalversions/livepodmigration/v1alpha1"
	listers "github.com/ubombar/live-pod-migration/pkg/generated/listers/livepodmigration/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

const controllerAgentName = "livepodmigration-migrator"

const (
	LivePodMigrationNamespace     = "livepodmigration"
	LivePodMigrationMigratorLabel = "migrator.livepodmigration.edgenet.io"
)

// Controller is the controller implementation for LivePodMigration resources
type Migrator struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// sampleclientset is a clientset for our own API group
	livepodmigrationclientset clientset.Interface

	livePodMigrationsLister listers.LivePodMigrationLister
	livePodMigrationSynced  cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface

	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

func NewMigrator(
	kubeclientset kubernetes.Interface,
	lpmclientset clientset.Interface,
	livePodMigrationInformer informers.LivePodMigrationInformer) *Migrator {

	utilruntime.Must(livepodmigrationscheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})

	recorder := eventBroadcaster.NewRecorder(livepodmigrationscheme.Scheme, corev1.EventSource{})

	controller := &Migrator{
		kubeclientset:             kubeclientset,
		livepodmigrationclientset: lpmclientset,

		livePodMigrationsLister: livePodMigrationInformer.Lister(),
		livePodMigrationSynced:  livePodMigrationInformer.Informer().HasSynced,

		workqueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), controllerAgentName),
		recorder:  recorder,
	}

	// Setting up event handlers
	klog.Info("Setting up event handlers")

	livePodMigrationInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		UpdateFunc: func(oldObj, newObj interface{}) {
			// Check for updates here
		},
	})

	return controller
}

func (m *Migrator) Run(stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer m.workqueue.ShutDown()

	if ok := cache.WaitForCacheSync(stopCh, m.livePodMigrationSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	go wait.Until(m.runWorker, time.Second, stopCh)

	<-stopCh

	return nil
}

func (m *Migrator) runWorker() {
	for m.processNextWorkItem() {

	}
}

func (m *Migrator) processNextWorkItem() bool {
	obj, shutdown := m.workqueue.Get()

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer m.workqueue.Done(obj)
		var key string
		var ok bool
		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			m.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncHandler, passing it the namespace/name string of the
		// Foo resource to be synced.
		if err := m.syncHandler(key); err != nil {
			// Put the item back on the workqueue to handle any transient errors.
			m.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		m.workqueue.Forget(obj)
		klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

func (m *Migrator) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)

	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// Get the Foo resource with this namespace/name
	lpm, err := m.livePodMigrationsLister.LivePodMigrations(namespace).Get(name)
	if err != nil {
		// The Foo resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("lpm '%s' in work queue no longer exists", key))
			return nil
		}
	}

	// TODO: Migrator
	fmt.Printf("lpm: %v\n", lpm)

	return nil
}
