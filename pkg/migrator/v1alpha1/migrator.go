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
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
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

		recorder: recorder,
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

	if ok := cache.WaitForCacheSync(stopCh, m.livePodMigrationSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	go wait.Until(m.runWorker, time.Second, stopCh)

	<-stopCh

	return nil
}

func (m *Migrator) runWorker() {
	// TODO: Implement the mechanism for peer to peer migration
}
