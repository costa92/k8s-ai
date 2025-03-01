

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	applyconfigurationscorev1 "github.com/costa92/k8s-ai/pkg/generated/applyconfigurations/core/v1"
	scheme "github.com/costa92/k8s-ai/pkg/generated/clientset/versioned/scheme"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ServicesGetter has a method to return a ServiceInterface.
// A group's client should implement this interface.
type ServicesGetter interface {
	Services(namespace string) ServiceInterface
}

// ServiceInterface has methods to work with Service resources.
type ServiceInterface interface {
	Create(ctx context.Context, service *corev1.Service, opts metav1.CreateOptions) (*corev1.Service, error)
	Update(ctx context.Context, service *corev1.Service, opts metav1.UpdateOptions) (*corev1.Service, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, service *corev1.Service, opts metav1.UpdateOptions) (*corev1.Service, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*corev1.Service, error)
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.ServiceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *corev1.Service, err error)
	Apply(ctx context.Context, service *applyconfigurationscorev1.ServiceApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.Service, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, service *applyconfigurationscorev1.ServiceApplyConfiguration, opts metav1.ApplyOptions) (result *corev1.Service, err error)
	ServiceExpansion
}

// services implements ServiceInterface
type services struct {
	*gentype.ClientWithListAndApply[*corev1.Service, *corev1.ServiceList, *applyconfigurationscorev1.ServiceApplyConfiguration]
}

// newServices returns a Services
func newServices(c *CoreV1Client, namespace string) *services {
	return &services{
		gentype.NewClientWithListAndApply[*corev1.Service, *corev1.ServiceList, *applyconfigurationscorev1.ServiceApplyConfiguration](
			"services",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *corev1.Service { return &corev1.Service{} },
			func() *corev1.ServiceList { return &corev1.ServiceList{} },
			gentype.PrefersProtobuf[*corev1.Service](),
		),
	}
}
