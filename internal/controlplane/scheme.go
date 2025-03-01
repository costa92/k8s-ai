package controlplane

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/api/legacyscheme"
)

func init() {
	// we need to add the options to empty v1
	metav1.AddToGroupVersion(legacyscheme.Scheme, schema.GroupVersion{Version: "v1"})
	// we need to add the options to empty v1
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	legacyscheme.Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}
