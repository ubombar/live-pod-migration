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
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

const controllerAgentName = "livepodmigration-controller"

const (
	LivePodMigrationNamespace     = "livepodmigration"
	LivePodMigrationMigratorLabel = "migrator.livepodmigration.edgenet.io"
)

// Controller is the controller implementation for LivePodMigration resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// sampleclientset is a clientset for our own API group
	livepodmigrationclientset clientset.Interface

	podsLister corelisters.PodLister
	podsSynced cache.InformerSynced

	livePodMigrationRequestLister listers.LivePodMigrationRequestLister
	livePodMigrationRequestSynced cache.InformerSynced

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

func NewController(
	kubeclientset kubernetes.Interface,
	lpmclientset clientset.Interface,
	podInformer coreinformers.PodInformer,
	livePodMigrationInformer informers.LivePodMigrationRequestInformer) *Controller {

	utilruntime.Must(livepodmigrationscheme.AddToScheme(scheme.Scheme))
	klog.V(4).Info("Creating event broadcaster")

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})

	recorder := eventBroadcaster.NewRecorder(livepodmigrationscheme.Scheme, corev1.EventSource{})

	controller := &Controller{
		kubeclientset:             kubeclientset,
		livepodmigrationclientset: lpmclientset,

		podsLister: podInformer.Lister(),
		podsSynced: livePodMigrationInformer.Informer().HasSynced,

		livePodMigrationRequestLister: livePodMigrationInformer.Lister(),
		livePodMigrationRequestSynced: livePodMigrationInformer.Informer().HasSynced,

		workqueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), controllerAgentName),
		recorder:  recorder,
	}

	// Setting up event handlers
	klog.Info("Setting up event handlers")

	livePodMigrationInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueLivePodMigrationRequest,
		UpdateFunc: func(oldObj, newObj interface{}) {
			controller.workqueue.Add(newObj)
		},
	})

	return controller
}

func (c *Controller) enqueueLivePodMigrationRequest(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}

func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	if ok := cache.WaitForCacheSync(stopCh, c.podsSynced, c.livePodMigrationRequestSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	klog.V(4).Infoln("Starting workers")
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	return nil
}

func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

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
		defer c.workqueue.Done(obj)
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
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncHandler, passing it the namespace/name string of the
		// Foo resource to be synced.
		if err := c.syncLivePodMigrationRequestHandler(key); err != nil {
			// Put the item back on the workqueue to handle any transient errors.
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		c.workqueue.Forget(obj)
		klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

func (c *Controller) syncLivePodMigrationRequestHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	fmt.Printf("namespace: %v\n", namespace)
	fmt.Printf("name: %v\n", name)

	// // Get the Foo resource with this namespace/name
	// lpm, err := c.livePodMigrationsLister.LivePodMigrations(namespace).Get(name)
	// if err != nil {
	// 	// The Foo resource may no longer exist, in which case we stop
	// 	// processing.
	// 	if errors.IsNotFound(err) {
	// 		utilruntime.HandleError(fmt.Errorf("lpm '%s' in work queue no longer exists", key))
	// 		return nil
	// 	}
	// }

	// lpmCopy := lpm.DeepCopy()

	// if lpm.Status.MigrationStatus == "" {
	// 	// If the lpm is just created
	// 	err := c.checkEligibilityOfMigration(lpmCopy)

	// 	lpmCopy.Status = v1alphav1types.LivePodMigrationStatus{
	// 		MigrationStatus:  v1alphav1types.MigrationStatusPending,
	// 		MigrationMessage: "",
	// 		PodAccessible:    true,
	// 		CheckpointFile:   "",
	// 	}

	// 	if err != nil {
	// 		lpmCopy.Status.MigrationStatus = v1alphav1types.MigrationStatusError
	// 		lpmCopy.Status.MigrationMessage = fmt.Sprint(err)
	// 	}

	// 	_, err = c.livepodmigrationclientset.LivepodmigrationV1alpha1().LivePodMigrations(namespace).UpdateStatus(context.Background(), lpmCopy, v1.UpdateOptions{})

	// 	if err != nil {
	// 		utilruntime.HandleError(fmt.Errorf("cannot update the lpm object '%s'", err))
	// 		return nil
	// 	}

	// 	c.recorder.Event(lpmCopy, corev1.EventTypeNormal, "SuccessSynced", fmt.Sprintf("We have updated the object on namespace %s \n", namespace))
	// } else if lpm.Status.MigrationStatus == v1alphav1types.MigrationStatusCheckpointing {
	// 	c.recorder.Event(lpmCopy, corev1.EventTypeNormal, "SuccessSynced", "Checkpointing")
	// } else if lpm.Status.MigrationStatus == v1alphav1types.MigrationStatusTransferring {
	// 	c.recorder.Event(lpmCopy, corev1.EventTypeNormal, "SuccessSynced", "Transferring")
	// } else if lpm.Status.MigrationStatus == v1alphav1types.MigrationStatusRestoring {
	// 	c.recorder.Event(lpmCopy, corev1.EventTypeNormal, "SuccessSynced", "Restoring")
	// } else if lpm.Status.MigrationStatus == v1alphav1types.MigrationStatusCleaning {
	// 	c.recorder.Event(lpmCopy, corev1.EventTypeNormal, "SuccessSynced", "Cleaning")
	// } else if lpm.Status.MigrationStatus == v1alphav1types.MigrationStatusCompleted {
	// 	c.recorder.Event(lpmCopy, corev1.EventTypeNormal, "SuccessSynced", "Completed")
	// } else if lpm.Status.MigrationStatus == v1alphav1types.MigrationStatusError {
	// 	c.recorder.Event(lpmCopy, corev1.EventTypeNormal, "SuccessSynced", "Error")
	// } else {
	// 	// Unknown status, drop
	// 	c.recorder.Event(lpmCopy, corev1.EventTypeNormal, "SuccessSynced", "Uknown status received")
	// }

	return nil
}
