

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	corev1 "github.com/costa92/k8s-ai/pkg/generated/applyconfigurations/core/v1"
	typedcorev1 "github.com/costa92/k8s-ai/pkg/generated/clientset/versioned/typed/core/v1"
	v1 "k8s.io/api/core/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeLimitRanges implements LimitRangeInterface
type fakeLimitRanges struct {
	*gentype.FakeClientWithListAndApply[*v1.LimitRange, *v1.LimitRangeList, *corev1.LimitRangeApplyConfiguration]
	Fake *FakeCoreV1
}

func newFakeLimitRanges(fake *FakeCoreV1, namespace string) typedcorev1.LimitRangeInterface {
	return &fakeLimitRanges{
		gentype.NewFakeClientWithListAndApply[*v1.LimitRange, *v1.LimitRangeList, *corev1.LimitRangeApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("limitranges"),
			v1.SchemeGroupVersion.WithKind("LimitRange"),
			func() *v1.LimitRange { return &v1.LimitRange{} },
			func() *v1.LimitRangeList { return &v1.LimitRangeList{} },
			func(dst, src *v1.LimitRangeList) { dst.ListMeta = src.ListMeta },
			func(list *v1.LimitRangeList) []*v1.LimitRange { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.LimitRangeList, items []*v1.LimitRange) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
