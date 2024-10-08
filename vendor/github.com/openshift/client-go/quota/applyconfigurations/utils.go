// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1 "github.com/openshift/api/quota/v1"
	internal "github.com/openshift/client-go/quota/applyconfigurations/internal"
	quotav1 "github.com/openshift/client-go/quota/applyconfigurations/quota/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=quota.openshift.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("ClusterResourceQuota"):
		return &quotav1.ClusterResourceQuotaApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterResourceQuotaSelector"):
		return &quotav1.ClusterResourceQuotaSelectorApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterResourceQuotaSpec"):
		return &quotav1.ClusterResourceQuotaSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ClusterResourceQuotaStatus"):
		return &quotav1.ClusterResourceQuotaStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ResourceQuotaStatusByNamespace"):
		return &quotav1.ResourceQuotaStatusByNamespaceApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
