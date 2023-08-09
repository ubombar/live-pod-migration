// package v1alpha1

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"time"

// 	"github.com/ubombar/live-pod-migration/pkg/apis/livepodmigration/v1alpha1"
// 	"github.com/ubombar/live-pod-migration/pkg/generated/clientset/versioned/fake"
// 	informers "github.com/ubombar/live-pod-migration/pkg/generated/informers/externalversions"
// 	v1 "k8s.io/api/core/v1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	k8sinformers "k8s.io/client-go/informers"
// 	k8sfake "k8s.io/client-go/kubernetes/fake"
// 	"k8s.io/client-go/tools/record"
// )

// var (
// 	alwaysReady        = func() bool { return true }
// 	noResyncPeriodFunc = func() time.Duration { return 0 }
// )

// func newController(stopCh chan struct{}) (*Controller, *k8sfake.Clientset, *fake.Clientset) {
// 	kubeclientset := k8sfake.NewSimpleClientset()
// 	lpmclientset := fake.NewSimpleClientset()

// 	kubeinformers := k8sinformers.NewSharedInformerFactory(kubeclientset, noResyncPeriodFunc())
// 	lpminformers := informers.NewSharedInformerFactory(lpmclientset, noResyncPeriodFunc())

// 	c := NewController(
// 		kubeclientset,
// 		lpmclientset,
// 		kubeinformers.Core().V1().Pods(),
// 		lpminformers.Livepodmigration().V1alpha1().LivePodMigrations())

// 	c.podsSynced = alwaysReady
// 	c.livePodMigrationSynced = alwaysReady
// 	c.recorder = &record.FakeRecorder{}

// 	kubeinformers.Start(stopCh)
// 	lpminformers.Start(stopCh)
// 	c.Run(2, stopCh)

// 	return c, kubeclientset, lpmclientset
// }

// func TestInitiation(t *testing.T) {
// 	stopCh := make(chan struct{})
// 	defer close(stopCh)

// 	_, kubeclientset, lpmclientset := newController(stopCh)

// 	kubeclientset.CoreV1().Nodes().Create(context.Background(), &v1.Node{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: "node1",
// 		},
// 		Spec: v1.NodeSpec{
// 			Unschedulable: false,
// 		},
// 	}, metav1.CreateOptions{})

// 	kubeclientset.CoreV1().Nodes().Create(context.Background(), &v1.Node{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: "node2",
// 		},
// 		Spec: v1.NodeSpec{
// 			Unschedulable: false,
// 		},
// 	}, metav1.CreateOptions{})

// 	kubeclientset.CoreV1().Pods("test").Create(context.Background(), &v1.Pod{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: "test-pod",
// 		},
// 		Status: v1.PodStatus{
// 			NominatedNodeName: "node1",
// 			Phase:             v1.PodRunning,
// 		},
// 	}, metav1.CreateOptions{})

// 	lpmclientset.LivepodmigrationV1alpha1().LivePodMigrations("default").Create(context.Background(), &v1alpha1.LivePodMigration{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name: "migration-test",
// 		},
// 		Spec: v1alpha1.LivePodMigrationSpec{
// 			PodNamespace:    "test",
// 			PodName:         "test-pod",
// 			DestinationNode: "node2",
// 			ServiceName:     "service1",
// 		},
// 	}, metav1.CreateOptions{})

// 	time.Sleep(time.Second)

// 	lpm, err := lpmclientset.LivepodmigrationV1alpha1().LivePodMigrations("default").Get(context.Background(), "migration-test", metav1.GetOptions{})

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if lpm.Status.MigrationStatus != v1alpha1.MigrationStatusPending {
// 		t.Error(fmt.Errorf("status not pending"))
// 	}

// }

// func TestMain(m *testing.M) {
// 	m.Run()
// }
