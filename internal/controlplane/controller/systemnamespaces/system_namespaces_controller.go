package systemnamespaces

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/tools/cache"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"
	versioned "k8s.io/client-go/kubernetes"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/klog/v2"
)

//listers "k8s.io/client-go/listers/core/v1"

// Controller ensure system namespaces exist.
type Controller struct {
	client versioned.Interface

	namespaceLister listers.NamespaceLister
	namespaceSynced cache.InformerSynced

	systemNamespaces []string
	interval         time.Duration
}

// NewController creates a new Controller to ensure system namespaces exist.
func NewController(clientset versioned.Interface, namespaceInformer coreinformers.NamespaceInformer) *Controller {
	systemNamespaces := []string{metav1.NamespaceSystem, metav1.NamespacePublic, v1.NamespaceNodeLease, metav1.NamespaceDefault}
	interval := 1 * time.Minute

	return &Controller{
		client:           clientset,
		namespaceLister:  namespaceInformer.Lister(),
		namespaceSynced:  namespaceInformer.Informer().HasSynced,
		systemNamespaces: systemNamespaces,
		interval:         interval,
	}
}

// Run starts one worker.
func (c *Controller) Run(stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	defer klog.Infof("Shutting down system namespaces controller")

	klog.Infof("Starting system namespaces controller")

	if !cache.WaitForCacheSync(stopCh, c.namespaceSynced) {
		utilruntime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}

	go wait.Until(c.sync, c.interval, stopCh)

	<-stopCh
}

func (c *Controller) sync() {
	// Loop the system namespace list, and create them if they do not exist
	for _, ns := range c.systemNamespaces {
		if err := c.createNamespaceIfNeeded(ns); err != nil {
			utilruntime.HandleError(fmt.Errorf("unable to create required kubernetes system Namespace %s: %v", ns, err))
		}
	}
}

func (c *Controller) createNamespaceIfNeeded(ns string) error {
	if _, err := c.namespaceLister.Get(ns); err == nil {
		// the namespace already exists
		return nil
	}
	newNs := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ns,
			Namespace: "",
		},
	}
	_, err := c.client.CoreV1().Namespaces().Create(context.TODO(), newNs, metav1.CreateOptions{})
	if err != nil && errors.IsAlreadyExists(err) {
		err = nil
	}
	return err
}
