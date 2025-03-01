

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	corev1 "github.com/costa92/k8s-ai/pkg/generated/applyconfigurations/core/v1"
	typedcorev1 "github.com/costa92/k8s-ai/pkg/generated/clientset/versioned/typed/core/v1"
	v1 "k8s.io/api/core/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakePodTemplates implements PodTemplateInterface
type fakePodTemplates struct {
	*gentype.FakeClientWithListAndApply[*v1.PodTemplate, *v1.PodTemplateList, *corev1.PodTemplateApplyConfiguration]
	Fake *FakeCoreV1
}

func newFakePodTemplates(fake *FakeCoreV1, namespace string) typedcorev1.PodTemplateInterface {
	return &fakePodTemplates{
		gentype.NewFakeClientWithListAndApply[*v1.PodTemplate, *v1.PodTemplateList, *corev1.PodTemplateApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("podtemplates"),
			v1.SchemeGroupVersion.WithKind("PodTemplate"),
			func() *v1.PodTemplate { return &v1.PodTemplate{} },
			func() *v1.PodTemplateList { return &v1.PodTemplateList{} },
			func(dst, src *v1.PodTemplateList) { dst.ListMeta = src.ListMeta },
			func(list *v1.PodTemplateList) []*v1.PodTemplate { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.PodTemplateList, items []*v1.PodTemplate) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
