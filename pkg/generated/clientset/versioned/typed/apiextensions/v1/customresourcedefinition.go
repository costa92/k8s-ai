

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	applyconfigurationsapiextensionsv1 "github.com/costa92/k8s-ai/pkg/generated/applyconfigurations/apiextensions/v1"
	scheme "github.com/costa92/k8s-ai/pkg/generated/clientset/versioned/scheme"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// CustomResourceDefinitionsGetter has a method to return a CustomResourceDefinitionInterface.
// A group's client should implement this interface.
type CustomResourceDefinitionsGetter interface {
	CustomResourceDefinitions() CustomResourceDefinitionInterface
}

// CustomResourceDefinitionInterface has methods to work with CustomResourceDefinition resources.
type CustomResourceDefinitionInterface interface {
	Create(ctx context.Context, customResourceDefinition *apiextensionsv1.CustomResourceDefinition, opts metav1.CreateOptions) (*apiextensionsv1.CustomResourceDefinition, error)
	Update(ctx context.Context, customResourceDefinition *apiextensionsv1.CustomResourceDefinition, opts metav1.UpdateOptions) (*apiextensionsv1.CustomResourceDefinition, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, customResourceDefinition *apiextensionsv1.CustomResourceDefinition, opts metav1.UpdateOptions) (*apiextensionsv1.CustomResourceDefinition, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*apiextensionsv1.CustomResourceDefinition, error)
	List(ctx context.Context, opts metav1.ListOptions) (*apiextensionsv1.CustomResourceDefinitionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *apiextensionsv1.CustomResourceDefinition, err error)
	Apply(ctx context.Context, customResourceDefinition *applyconfigurationsapiextensionsv1.CustomResourceDefinitionApplyConfiguration, opts metav1.ApplyOptions) (result *apiextensionsv1.CustomResourceDefinition, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, customResourceDefinition *applyconfigurationsapiextensionsv1.CustomResourceDefinitionApplyConfiguration, opts metav1.ApplyOptions) (result *apiextensionsv1.CustomResourceDefinition, err error)
	CustomResourceDefinitionExpansion
}

// customResourceDefinitions implements CustomResourceDefinitionInterface
type customResourceDefinitions struct {
	*gentype.ClientWithListAndApply[*apiextensionsv1.CustomResourceDefinition, *apiextensionsv1.CustomResourceDefinitionList, *applyconfigurationsapiextensionsv1.CustomResourceDefinitionApplyConfiguration]
}

// newCustomResourceDefinitions returns a CustomResourceDefinitions
func newCustomResourceDefinitions(c *ApiextensionsV1Client) *customResourceDefinitions {
	return &customResourceDefinitions{
		gentype.NewClientWithListAndApply[*apiextensionsv1.CustomResourceDefinition, *apiextensionsv1.CustomResourceDefinitionList, *applyconfigurationsapiextensionsv1.CustomResourceDefinitionApplyConfiguration](
			"customresourcedefinitions",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *apiextensionsv1.CustomResourceDefinition { return &apiextensionsv1.CustomResourceDefinition{} },
			func() *apiextensionsv1.CustomResourceDefinitionList {
				return &apiextensionsv1.CustomResourceDefinitionList{}
			},
			gentype.PrefersProtobuf[*apiextensionsv1.CustomResourceDefinition](),
		),
	}
}
