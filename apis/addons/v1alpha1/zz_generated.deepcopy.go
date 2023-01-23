// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddOnStatusCondition) DeepCopyInto(out *AddOnStatusCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddOnStatusCondition.
func (in *AddOnStatusCondition) DeepCopy() *AddOnStatusCondition {
	if in == nil {
		return nil
	}
	out := new(AddOnStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalCatalogSource) DeepCopyInto(out *AdditionalCatalogSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalCatalogSource.
func (in *AdditionalCatalogSource) DeepCopy() *AdditionalCatalogSource {
	if in == nil {
		return nil
	}
	out := new(AdditionalCatalogSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Addon) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallOLMAllNamespaces) DeepCopyInto(out *AddonInstallOLMAllNamespaces) {
	*out = *in
	in.AddonInstallOLMCommon.DeepCopyInto(&out.AddonInstallOLMCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallOLMAllNamespaces.
func (in *AddonInstallOLMAllNamespaces) DeepCopy() *AddonInstallOLMAllNamespaces {
	if in == nil {
		return nil
	}
	out := new(AddonInstallOLMAllNamespaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallOLMCommon) DeepCopyInto(out *AddonInstallOLMCommon) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(SubscriptionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalCatalogSources != nil {
		in, out := &in.AdditionalCatalogSources, &out.AdditionalCatalogSources
		*out = make([]AdditionalCatalogSource, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallOLMCommon.
func (in *AddonInstallOLMCommon) DeepCopy() *AddonInstallOLMCommon {
	if in == nil {
		return nil
	}
	out := new(AddonInstallOLMCommon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallOLMOwnNamespace) DeepCopyInto(out *AddonInstallOLMOwnNamespace) {
	*out = *in
	in.AddonInstallOLMCommon.DeepCopyInto(&out.AddonInstallOLMCommon)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallOLMOwnNamespace.
func (in *AddonInstallOLMOwnNamespace) DeepCopy() *AddonInstallOLMOwnNamespace {
	if in == nil {
		return nil
	}
	out := new(AddonInstallOLMOwnNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallSpec) DeepCopyInto(out *AddonInstallSpec) {
	*out = *in
	if in.OLMAllNamespaces != nil {
		in, out := &in.OLMAllNamespaces, &out.OLMAllNamespaces
		*out = new(AddonInstallOLMAllNamespaces)
		(*in).DeepCopyInto(*out)
	}
	if in.OLMOwnNamespace != nil {
		in, out := &in.OLMOwnNamespace, &out.OLMOwnNamespace
		*out = new(AddonInstallOLMOwnNamespace)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallSpec.
func (in *AddonInstallSpec) DeepCopy() *AddonInstallSpec {
	if in == nil {
		return nil
	}
	out := new(AddonInstallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstance) DeepCopyInto(out *AddonInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstance.
func (in *AddonInstance) DeepCopy() *AddonInstance {
	if in == nil {
		return nil
	}
	out := new(AddonInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstanceList) DeepCopyInto(out *AddonInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddonInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstanceList.
func (in *AddonInstanceList) DeepCopy() *AddonInstanceList {
	if in == nil {
		return nil
	}
	out := new(AddonInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstanceSpec) DeepCopyInto(out *AddonInstanceSpec) {
	*out = *in
	out.HeartbeatUpdatePeriod = in.HeartbeatUpdatePeriod
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstanceSpec.
func (in *AddonInstanceSpec) DeepCopy() *AddonInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(AddonInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstanceStatus) DeepCopyInto(out *AddonInstanceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstanceStatus.
func (in *AddonInstanceStatus) DeepCopy() *AddonInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(AddonInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonList) DeepCopyInto(out *AddonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Addon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonList.
func (in *AddonList) DeepCopy() *AddonList {
	if in == nil {
		return nil
	}
	out := new(AddonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonNamespace) DeepCopyInto(out *AddonNamespace) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonNamespace.
func (in *AddonNamespace) DeepCopy() *AddonNamespace {
	if in == nil {
		return nil
	}
	out := new(AddonNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonOperator) DeepCopyInto(out *AddonOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonOperator.
func (in *AddonOperator) DeepCopy() *AddonOperator {
	if in == nil {
		return nil
	}
	out := new(AddonOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonOperatorFeatureToggles) DeepCopyInto(out *AddonOperatorFeatureToggles) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonOperatorFeatureToggles.
func (in *AddonOperatorFeatureToggles) DeepCopy() *AddonOperatorFeatureToggles {
	if in == nil {
		return nil
	}
	out := new(AddonOperatorFeatureToggles)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonOperatorList) DeepCopyInto(out *AddonOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddonOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonOperatorList.
func (in *AddonOperatorList) DeepCopy() *AddonOperatorList {
	if in == nil {
		return nil
	}
	out := new(AddonOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonOperatorOCM) DeepCopyInto(out *AddonOperatorOCM) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonOperatorOCM.
func (in *AddonOperatorOCM) DeepCopy() *AddonOperatorOCM {
	if in == nil {
		return nil
	}
	out := new(AddonOperatorOCM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonOperatorSpec) DeepCopyInto(out *AddonOperatorSpec) {
	*out = *in
	out.FeatureToggles = in.FeatureToggles
	if in.OCM != nil {
		in, out := &in.OCM, &out.OCM
		*out = new(AddonOperatorOCM)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonOperatorSpec.
func (in *AddonOperatorSpec) DeepCopy() *AddonOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(AddonOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonOperatorStatus) DeepCopyInto(out *AddonOperatorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LastHeartbeatTime.DeepCopyInto(&out.LastHeartbeatTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonOperatorStatus.
func (in *AddonOperatorStatus) DeepCopy() *AddonOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(AddonOperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonSecretPropagation) DeepCopyInto(out *AddonSecretPropagation) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]AddonSecretPropagationReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonSecretPropagation.
func (in *AddonSecretPropagation) DeepCopy() *AddonSecretPropagation {
	if in == nil {
		return nil
	}
	out := new(AddonSecretPropagation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonSecretPropagationReference) DeepCopyInto(out *AddonSecretPropagationReference) {
	*out = *in
	out.SourceSecret = in.SourceSecret
	out.DestinationSecret = in.DestinationSecret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonSecretPropagationReference.
func (in *AddonSecretPropagationReference) DeepCopy() *AddonSecretPropagationReference {
	if in == nil {
		return nil
	}
	out := new(AddonSecretPropagationReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonSpec) DeepCopyInto(out *AddonSpec) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]AddonNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CommonLabels != nil {
		in, out := &in.CommonLabels, &out.CommonLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CommonAnnotations != nil {
		in, out := &in.CommonAnnotations, &out.CommonAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Install.DeepCopyInto(&out.Install)
	if in.UpgradePolicy != nil {
		in, out := &in.UpgradePolicy, &out.UpgradePolicy
		*out = new(AddonUpgradePolicy)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(MonitoringSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretPropagation != nil {
		in, out := &in.SecretPropagation, &out.SecretPropagation
		*out = new(AddonSecretPropagation)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonSpec.
func (in *AddonSpec) DeepCopy() *AddonSpec {
	if in == nil {
		return nil
	}
	out := new(AddonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonStatus) DeepCopyInto(out *AddonStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UpgradePolicy != nil {
		in, out := &in.UpgradePolicy, &out.UpgradePolicy
		*out = new(AddonUpgradePolicyStatus)
		**out = **in
	}
	if in.OCMReportedStatus != nil {
		in, out := &in.OCMReportedStatus, &out.OCMReportedStatus
		*out = new(OCMAddOnStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonStatus.
func (in *AddonStatus) DeepCopy() *AddonStatus {
	if in == nil {
		return nil
	}
	out := new(AddonStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonUpgradePolicy) DeepCopyInto(out *AddonUpgradePolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonUpgradePolicy.
func (in *AddonUpgradePolicy) DeepCopy() *AddonUpgradePolicy {
	if in == nil {
		return nil
	}
	out := new(AddonUpgradePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonUpgradePolicyStatus) DeepCopyInto(out *AddonUpgradePolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonUpgradePolicyStatus.
func (in *AddonUpgradePolicyStatus) DeepCopy() *AddonUpgradePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(AddonUpgradePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretReference) DeepCopyInto(out *ClusterSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretReference.
func (in *ClusterSecretReference) DeepCopy() *ClusterSecretReference {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvObject) DeepCopyInto(out *EnvObject) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvObject.
func (in *EnvObject) DeepCopy() *EnvObject {
	if in == nil {
		return nil
	}
	out := new(EnvObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringFederationSpec) DeepCopyInto(out *MonitoringFederationSpec) {
	*out = *in
	if in.MatchNames != nil {
		in, out := &in.MatchNames, &out.MatchNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MatchLabels != nil {
		in, out := &in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringFederationSpec.
func (in *MonitoringFederationSpec) DeepCopy() *MonitoringFederationSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringFederationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringSpec) DeepCopyInto(out *MonitoringSpec) {
	*out = *in
	if in.Federation != nil {
		in, out := &in.Federation, &out.Federation
		*out = new(MonitoringFederationSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringSpec.
func (in *MonitoringSpec) DeepCopy() *MonitoringSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCMAddOnStatus) DeepCopyInto(out *OCMAddOnStatus) {
	*out = *in
	if in.StatusConditions != nil {
		in, out := &in.StatusConditions, &out.StatusConditions
		*out = make([]AddOnStatusCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCMAddOnStatus.
func (in *OCMAddOnStatus) DeepCopy() *OCMAddOnStatus {
	if in == nil {
		return nil
	}
	out := new(OCMAddOnStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionConfig) DeepCopyInto(out *SubscriptionConfig) {
	*out = *in
	if in.EnvironmentVariables != nil {
		in, out := &in.EnvironmentVariables, &out.EnvironmentVariables
		*out = make([]EnvObject, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionConfig.
func (in *SubscriptionConfig) DeepCopy() *SubscriptionConfig {
	if in == nil {
		return nil
	}
	out := new(SubscriptionConfig)
	in.DeepCopyInto(out)
	return out
}
