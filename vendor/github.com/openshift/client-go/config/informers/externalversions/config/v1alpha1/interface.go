// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openshift/client-go/config/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Backups returns a BackupInformer.
	Backups() BackupInformer
	// ClusterImagePolicies returns a ClusterImagePolicyInformer.
	ClusterImagePolicies() ClusterImagePolicyInformer
	// ClusterMonitorings returns a ClusterMonitoringInformer.
	ClusterMonitorings() ClusterMonitoringInformer
	// ImagePolicies returns a ImagePolicyInformer.
	ImagePolicies() ImagePolicyInformer
	// InsightsDataGathers returns a InsightsDataGatherInformer.
	InsightsDataGathers() InsightsDataGatherInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Backups returns a BackupInformer.
func (v *version) Backups() BackupInformer {
	return &backupInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterImagePolicies returns a ClusterImagePolicyInformer.
func (v *version) ClusterImagePolicies() ClusterImagePolicyInformer {
	return &clusterImagePolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterMonitorings returns a ClusterMonitoringInformer.
func (v *version) ClusterMonitorings() ClusterMonitoringInformer {
	return &clusterMonitoringInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ImagePolicies returns a ImagePolicyInformer.
func (v *version) ImagePolicies() ImagePolicyInformer {
	return &imagePolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InsightsDataGathers returns a InsightsDataGatherInformer.
func (v *version) InsightsDataGathers() InsightsDataGatherInformer {
	return &insightsDataGatherInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
