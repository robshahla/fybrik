// +build !ignore_autogenerated

// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"fybrik.io/fybrik/pkg/serde"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ApplicationDetails) DeepCopyInto(out *ApplicationDetails) {
	{
		in := &in
		*out = make(ApplicationDetails, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDetails.
func (in ApplicationDetails) DeepCopy() ApplicationDetails {
	if in == nil {
		return nil
	}
	out := new(ApplicationDetails)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Blueprint) DeepCopyInto(out *Blueprint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Blueprint.
func (in *Blueprint) DeepCopy() *Blueprint {
	if in == nil {
		return nil
	}
	out := new(Blueprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Blueprint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintList) DeepCopyInto(out *BlueprintList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Blueprint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintList.
func (in *BlueprintList) DeepCopy() *BlueprintList {
	if in == nil {
		return nil
	}
	out := new(BlueprintList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlueprintList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintSpec) DeepCopyInto(out *BlueprintSpec) {
	*out = *in
	in.Flow.DeepCopyInto(&out.Flow)
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]ComponentTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintSpec.
func (in *BlueprintSpec) DeepCopy() *BlueprintSpec {
	if in == nil {
		return nil
	}
	out := new(BlueprintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintStatus) DeepCopyInto(out *BlueprintStatus) {
	*out = *in
	out.ObservedState = in.ObservedState
	if in.Releases != nil {
		in, out := &in.Releases, &out.Releases
		*out = make(map[string]int64, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintStatus.
func (in *BlueprintStatus) DeepCopy() *BlueprintStatus {
	if in == nil {
		return nil
	}
	out := new(BlueprintStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogRequirements) DeepCopyInto(out *CatalogRequirements) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogRequirements.
func (in *CatalogRequirements) DeepCopy() *CatalogRequirements {
	if in == nil {
		return nil
	}
	out := new(CatalogRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpec) DeepCopyInto(out *ChartSpec) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpec.
func (in *ChartSpec) DeepCopy() *ChartSpec {
	if in == nil {
		return nil
	}
	out := new(ChartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentTemplate) DeepCopyInto(out *ComponentTemplate) {
	*out = *in
	in.Chart.DeepCopyInto(&out.Chart)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentTemplate.
func (in *ComponentTemplate) DeepCopy() *ComponentTemplate {
	if in == nil {
		return nil
	}
	out := new(ComponentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopyModuleArgs) DeepCopyInto(out *CopyModuleArgs) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]serde.Arbitrary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopyModuleArgs.
func (in *CopyModuleArgs) DeepCopy() *CopyModuleArgs {
	if in == nil {
		return nil
	}
	out := new(CopyModuleArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopyRequirements) DeepCopyInto(out *CopyRequirements) {
	*out = *in
	out.Catalog = in.Catalog
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopyRequirements.
func (in *CopyRequirements) DeepCopy() *CopyRequirements {
	if in == nil {
		return nil
	}
	out := new(CopyRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataContext) DeepCopyInto(out *DataContext) {
	*out = *in
	out.Requirements = in.Requirements
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataContext.
func (in *DataContext) DeepCopy() *DataContext {
	if in == nil {
		return nil
	}
	out := new(DataContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataFlow) DeepCopyInto(out *DataFlow) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]FlowStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataFlow.
func (in *DataFlow) DeepCopy() *DataFlow {
	if in == nil {
		return nil
	}
	out := new(DataFlow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataRequirements) DeepCopyInto(out *DataRequirements) {
	*out = *in
	out.Interface = in.Interface
	out.Copy = in.Copy
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataRequirements.
func (in *DataRequirements) DeepCopy() *DataRequirements {
	if in == nil {
		return nil
	}
	out := new(DataRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataStore) DeepCopyInto(out *DataStore) {
	*out = *in
	out.Vault = in.Vault
	in.Connection.DeepCopyInto(&out.Connection)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataStore.
func (in *DataStore) DeepCopy() *DataStore {
	if in == nil {
		return nil
	}
	out := new(DataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetDetails) DeepCopyInto(out *DatasetDetails) {
	*out = *in
	in.Details.DeepCopyInto(&out.Details)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetDetails.
func (in *DatasetDetails) DeepCopy() *DatasetDetails {
	if in == nil {
		return nil
	}
	out := new(DatasetDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dependency) DeepCopyInto(out *Dependency) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dependency.
func (in *Dependency) DeepCopy() *Dependency {
	if in == nil {
		return nil
	}
	out := new(Dependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpec) DeepCopyInto(out *EndpointSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpec.
func (in *EndpointSpec) DeepCopy() *EndpointSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowStep) DeepCopyInto(out *FlowStep) {
	*out = *in
	in.Arguments.DeepCopyInto(&out.Arguments)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowStep.
func (in *FlowStep) DeepCopy() *FlowStep {
	if in == nil {
		return nil
	}
	out := new(FlowStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikApplication) DeepCopyInto(out *FybrikApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikApplication.
func (in *FybrikApplication) DeepCopy() *FybrikApplication {
	if in == nil {
		return nil
	}
	out := new(FybrikApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikApplicationList) DeepCopyInto(out *FybrikApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FybrikApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikApplicationList.
func (in *FybrikApplicationList) DeepCopy() *FybrikApplicationList {
	if in == nil {
		return nil
	}
	out := new(FybrikApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikApplicationSpec) DeepCopyInto(out *FybrikApplicationSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.AppInfo != nil {
		in, out := &in.AppInfo, &out.AppInfo
		*out = make(ApplicationDetails, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataContext, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikApplicationSpec.
func (in *FybrikApplicationSpec) DeepCopy() *FybrikApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(FybrikApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikApplicationStatus) DeepCopyInto(out *FybrikApplicationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	if in.CatalogedAssets != nil {
		in, out := &in.CatalogedAssets, &out.CatalogedAssets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Generated != nil {
		in, out := &in.Generated, &out.Generated
		*out = new(ResourceReference)
		**out = **in
	}
	if in.ProvisionedStorage != nil {
		in, out := &in.ProvisionedStorage, &out.ProvisionedStorage
		*out = make(map[string]DatasetDetails, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ReadEndpointsMap != nil {
		in, out := &in.ReadEndpointsMap, &out.ReadEndpointsMap
		*out = make(map[string]EndpointSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikApplicationStatus.
func (in *FybrikApplicationStatus) DeepCopy() *FybrikApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(FybrikApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikModule) DeepCopyInto(out *FybrikModule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikModule.
func (in *FybrikModule) DeepCopy() *FybrikModule {
	if in == nil {
		return nil
	}
	out := new(FybrikModule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikModule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikModuleList) DeepCopyInto(out *FybrikModuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FybrikModule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikModuleList.
func (in *FybrikModuleList) DeepCopy() *FybrikModuleList {
	if in == nil {
		return nil
	}
	out := new(FybrikModuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikModuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikModuleSpec) DeepCopyInto(out *FybrikModuleSpec) {
	*out = *in
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]Dependency, len(*in))
		copy(*out, *in)
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]ModuleCapability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Chart.DeepCopyInto(&out.Chart)
	if in.StatusIndicators != nil {
		in, out := &in.StatusIndicators, &out.StatusIndicators
		*out = make([]ResourceStatusIndicator, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikModuleSpec.
func (in *FybrikModuleSpec) DeepCopy() *FybrikModuleSpec {
	if in == nil {
		return nil
	}
	out := new(FybrikModuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikStorageAccount) DeepCopyInto(out *FybrikStorageAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikStorageAccount.
func (in *FybrikStorageAccount) DeepCopy() *FybrikStorageAccount {
	if in == nil {
		return nil
	}
	out := new(FybrikStorageAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikStorageAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikStorageAccountList) DeepCopyInto(out *FybrikStorageAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FybrikStorageAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikStorageAccountList.
func (in *FybrikStorageAccountList) DeepCopy() *FybrikStorageAccountList {
	if in == nil {
		return nil
	}
	out := new(FybrikStorageAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FybrikStorageAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikStorageAccountSpec) DeepCopyInto(out *FybrikStorageAccountSpec) {
	*out = *in
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikStorageAccountSpec.
func (in *FybrikStorageAccountSpec) DeepCopy() *FybrikStorageAccountSpec {
	if in == nil {
		return nil
	}
	out := new(FybrikStorageAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FybrikStorageAccountStatus) DeepCopyInto(out *FybrikStorageAccountStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FybrikStorageAccountStatus.
func (in *FybrikStorageAccountStatus) DeepCopy() *FybrikStorageAccountStatus {
	if in == nil {
		return nil
	}
	out := new(FybrikStorageAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceDetails) DeepCopyInto(out *InterfaceDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceDetails.
func (in *InterfaceDetails) DeepCopy() *InterfaceDetails {
	if in == nil {
		return nil
	}
	out := new(InterfaceDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaBlueprint) DeepCopyInto(out *MetaBlueprint) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaBlueprint.
func (in *MetaBlueprint) DeepCopy() *MetaBlueprint {
	if in == nil {
		return nil
	}
	out := new(MetaBlueprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleAPI) DeepCopyInto(out *ModuleAPI) {
	*out = *in
	out.InterfaceDetails = in.InterfaceDetails
	out.Endpoint = in.Endpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleAPI.
func (in *ModuleAPI) DeepCopy() *ModuleAPI {
	if in == nil {
		return nil
	}
	out := new(ModuleAPI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleArguments) DeepCopyInto(out *ModuleArguments) {
	*out = *in
	if in.Copy != nil {
		in, out := &in.Copy, &out.Copy
		*out = new(CopyModuleArgs)
		(*in).DeepCopyInto(*out)
	}
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = make([]ReadModuleArgs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = make([]WriteModuleArgs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleArguments.
func (in *ModuleArguments) DeepCopy() *ModuleArguments {
	if in == nil {
		return nil
	}
	out := new(ModuleArguments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleCapability) DeepCopyInto(out *ModuleCapability) {
	*out = *in
	if in.SupportedInterfaces != nil {
		in, out := &in.SupportedInterfaces, &out.SupportedInterfaces
		*out = make([]ModuleInOut, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(ModuleAPI)
		**out = **in
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]SupportedAction, len(*in))
		copy(*out, *in)
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make([]Plugin, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleCapability.
func (in *ModuleCapability) DeepCopy() *ModuleCapability {
	if in == nil {
		return nil
	}
	out := new(ModuleCapability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleInOut) DeepCopyInto(out *ModuleInOut) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(InterfaceDetails)
		**out = **in
	}
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(InterfaceDetails)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleInOut.
func (in *ModuleInOut) DeepCopy() *ModuleInOut {
	if in == nil {
		return nil
	}
	out := new(ModuleInOut)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservedState) DeepCopyInto(out *ObservedState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservedState.
func (in *ObservedState) DeepCopy() *ObservedState {
	if in == nil {
		return nil
	}
	out := new(ObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plotter) DeepCopyInto(out *Plotter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plotter.
func (in *Plotter) DeepCopy() *Plotter {
	if in == nil {
		return nil
	}
	out := new(Plotter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Plotter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlotterList) DeepCopyInto(out *PlotterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Plotter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlotterList.
func (in *PlotterList) DeepCopy() *PlotterList {
	if in == nil {
		return nil
	}
	out := new(PlotterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlotterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlotterSpec) DeepCopyInto(out *PlotterSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Blueprints != nil {
		in, out := &in.Blueprints, &out.Blueprints
		*out = make(map[string]BlueprintSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlotterSpec.
func (in *PlotterSpec) DeepCopy() *PlotterSpec {
	if in == nil {
		return nil
	}
	out := new(PlotterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlotterStatus) DeepCopyInto(out *PlotterStatus) {
	*out = *in
	out.ObservedState = in.ObservedState
	if in.Blueprints != nil {
		in, out := &in.Blueprints, &out.Blueprints
		*out = make(map[string]MetaBlueprint, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ReadyTimestamp != nil {
		in, out := &in.ReadyTimestamp, &out.ReadyTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlotterStatus.
func (in *PlotterStatus) DeepCopy() *PlotterStatus {
	if in == nil {
		return nil
	}
	out := new(PlotterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadModuleArgs) DeepCopyInto(out *ReadModuleArgs) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]serde.Arbitrary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadModuleArgs.
func (in *ReadModuleArgs) DeepCopy() *ReadModuleArgs {
	if in == nil {
		return nil
	}
	out := new(ReadModuleArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceReference) DeepCopyInto(out *ResourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceReference.
func (in *ResourceReference) DeepCopy() *ResourceReference {
	if in == nil {
		return nil
	}
	out := new(ResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatusIndicator) DeepCopyInto(out *ResourceStatusIndicator) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatusIndicator.
func (in *ResourceStatusIndicator) DeepCopy() *ResourceStatusIndicator {
	if in == nil {
		return nil
	}
	out := new(ResourceStatusIndicator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	in.WorkloadSelector.DeepCopyInto(&out.WorkloadSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SupportedAction) DeepCopyInto(out *SupportedAction) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SupportedAction.
func (in *SupportedAction) DeepCopy() *SupportedAction {
	if in == nil {
		return nil
	}
	out := new(SupportedAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault) DeepCopyInto(out *Vault) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault.
func (in *Vault) DeepCopy() *Vault {
	if in == nil {
		return nil
	}
	out := new(Vault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WriteModuleArgs) DeepCopyInto(out *WriteModuleArgs) {
	*out = *in
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]serde.Arbitrary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WriteModuleArgs.
func (in *WriteModuleArgs) DeepCopy() *WriteModuleArgs {
	if in == nil {
		return nil
	}
	out := new(WriteModuleArgs)
	in.DeepCopyInto(out)
	return out
}
