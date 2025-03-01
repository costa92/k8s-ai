

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	corev1 "github.com/costa92/k8s-ai/pkg/generated/applyconfigurations/core/v1"
	typedcorev1 "github.com/costa92/k8s-ai/pkg/generated/clientset/versioned/typed/core/v1"
	v1 "k8s.io/api/core/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeNodes implements NodeInterface
type fakeNodes struct {
	*gentype.FakeClientWithListAndApply[*v1.Node, *v1.NodeList, *corev1.NodeApplyConfiguration]
	Fake *FakeCoreV1
}

func newFakeNodes(fake *FakeCoreV1) typedcorev1.NodeInterface {
	return &fakeNodes{
		gentype.NewFakeClientWithListAndApply[*v1.Node, *v1.NodeList, *corev1.NodeApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("nodes"),
			v1.SchemeGroupVersion.WithKind("Node"),
			func() *v1.Node { return &v1.Node{} },
			func() *v1.NodeList { return &v1.NodeList{} },
			func(dst, src *v1.NodeList) { dst.ListMeta = src.ListMeta },
			func(list *v1.NodeList) []*v1.Node { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.NodeList, items []*v1.Node) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
