//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1alpha1 "k8s.io/component-base/config/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BastionControllerConfiguration) DeepCopyInto(out *BastionControllerConfiguration) {
	*out = *in
	if in.MaxLifetime != nil {
		in, out := &in.MaxLifetime, &out.MaxLifetime
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BastionControllerConfiguration.
func (in *BastionControllerConfiguration) DeepCopy() *BastionControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(BastionControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProfileControllerConfiguration) DeepCopyInto(out *CloudProfileControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProfileControllerConfiguration.
func (in *CloudProfileControllerConfiguration) DeepCopy() *CloudProfileControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(CloudProfileControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerDeploymentControllerConfiguration) DeepCopyInto(out *ControllerDeploymentControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerDeploymentControllerConfiguration.
func (in *ControllerDeploymentControllerConfiguration) DeepCopy() *ControllerDeploymentControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerDeploymentControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerConfiguration) DeepCopyInto(out *ControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.GardenClientConnection = in.GardenClientConnection
	in.Controllers.DeepCopyInto(&out.Controllers)
	if in.LeaderElection != nil {
		in, out := &in.LeaderElection, &out.LeaderElection
		*out = new(configv1alpha1.LeaderElectionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	out.Server = in.Server
	if in.Debugging != nil {
		in, out := &in.Debugging, &out.Debugging
		*out = new(configv1alpha1.DebuggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerConfiguration.
func (in *ControllerManagerConfiguration) DeepCopy() *ControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerControllerConfiguration) DeepCopyInto(out *ControllerManagerControllerConfiguration) {
	*out = *in
	if in.Bastion != nil {
		in, out := &in.Bastion, &out.Bastion
		*out = new(BastionControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudProfile != nil {
		in, out := &in.CloudProfile, &out.CloudProfile
		*out = new(CloudProfileControllerConfiguration)
		**out = **in
	}
	if in.ControllerDeployment != nil {
		in, out := &in.ControllerDeployment, &out.ControllerDeployment
		*out = new(ControllerDeploymentControllerConfiguration)
		**out = **in
	}
	if in.ControllerRegistration != nil {
		in, out := &in.ControllerRegistration, &out.ControllerRegistration
		*out = new(ControllerRegistrationControllerConfiguration)
		**out = **in
	}
	if in.Event != nil {
		in, out := &in.Event, &out.Event
		*out = new(EventControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ExposureClass != nil {
		in, out := &in.ExposureClass, &out.ExposureClass
		*out = new(ExposureClassControllerConfiguration)
		**out = **in
	}
	if in.Plant != nil {
		in, out := &in.Plant, &out.Plant
		*out = new(PlantControllerConfiguration)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(ProjectControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(QuotaControllerConfiguration)
		**out = **in
	}
	if in.SecretBinding != nil {
		in, out := &in.SecretBinding, &out.SecretBinding
		*out = new(SecretBindingControllerConfiguration)
		**out = **in
	}
	if in.SecretBindingProvider != nil {
		in, out := &in.SecretBindingProvider, &out.SecretBindingProvider
		*out = new(SecretBindingProviderControllerConfiguration)
		**out = **in
	}
	if in.Seed != nil {
		in, out := &in.Seed, &out.Seed
		*out = new(SeedControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	in.ShootMaintenance.DeepCopyInto(&out.ShootMaintenance)
	out.ShootQuota = in.ShootQuota
	out.ShootHibernation = in.ShootHibernation
	if in.ShootReference != nil {
		in, out := &in.ShootReference, &out.ShootReference
		*out = new(ShootReferenceControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ShootRetry != nil {
		in, out := &in.ShootRetry, &out.ShootRetry
		*out = new(ShootRetryControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ShootConditions != nil {
		in, out := &in.ShootConditions, &out.ShootConditions
		*out = new(ShootConditionsControllerConfiguration)
		**out = **in
	}
	if in.ShootStatusLabel != nil {
		in, out := &in.ShootStatusLabel, &out.ShootStatusLabel
		*out = new(ShootStatusLabelControllerConfiguration)
		**out = **in
	}
	if in.ManagedSeedSet != nil {
		in, out := &in.ManagedSeedSet, &out.ManagedSeedSet
		*out = new(ManagedSeedSetControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerControllerConfiguration.
func (in *ControllerManagerControllerConfiguration) DeepCopy() *ControllerManagerControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRegistrationControllerConfiguration) DeepCopyInto(out *ControllerRegistrationControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRegistrationControllerConfiguration.
func (in *ControllerRegistrationControllerConfiguration) DeepCopy() *ControllerRegistrationControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerRegistrationControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventControllerConfiguration) DeepCopyInto(out *EventControllerConfiguration) {
	*out = *in
	if in.TTLNonShootEvents != nil {
		in, out := &in.TTLNonShootEvents, &out.TTLNonShootEvents
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventControllerConfiguration.
func (in *EventControllerConfiguration) DeepCopy() *EventControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(EventControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExposureClassControllerConfiguration) DeepCopyInto(out *ExposureClassControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExposureClassControllerConfiguration.
func (in *ExposureClassControllerConfiguration) DeepCopy() *ExposureClassControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ExposureClassControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSServer) DeepCopyInto(out *HTTPSServer) {
	*out = *in
	out.Server = in.Server
	out.TLS = in.TLS
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSServer.
func (in *HTTPSServer) DeepCopy() *HTTPSServer {
	if in == nil {
		return nil
	}
	out := new(HTTPSServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedSeedSetControllerConfiguration) DeepCopyInto(out *ManagedSeedSetControllerConfiguration) {
	*out = *in
	if in.MaxShootRetries != nil {
		in, out := &in.MaxShootRetries, &out.MaxShootRetries
		*out = new(int)
		**out = **in
	}
	out.SyncPeriod = in.SyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedSeedSetControllerConfiguration.
func (in *ManagedSeedSetControllerConfiguration) DeepCopy() *ManagedSeedSetControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ManagedSeedSetControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlantControllerConfiguration) DeepCopyInto(out *PlantControllerConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlantControllerConfiguration.
func (in *PlantControllerConfiguration) DeepCopy() *PlantControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(PlantControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectControllerConfiguration) DeepCopyInto(out *ProjectControllerConfiguration) {
	*out = *in
	if in.MinimumLifetimeDays != nil {
		in, out := &in.MinimumLifetimeDays, &out.MinimumLifetimeDays
		*out = new(int)
		**out = **in
	}
	if in.Quotas != nil {
		in, out := &in.Quotas, &out.Quotas
		*out = make([]QuotaConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StaleGracePeriodDays != nil {
		in, out := &in.StaleGracePeriodDays, &out.StaleGracePeriodDays
		*out = new(int)
		**out = **in
	}
	if in.StaleExpirationTimeDays != nil {
		in, out := &in.StaleExpirationTimeDays, &out.StaleExpirationTimeDays
		*out = new(int)
		**out = **in
	}
	if in.StaleSyncPeriod != nil {
		in, out := &in.StaleSyncPeriod, &out.StaleSyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectControllerConfiguration.
func (in *ProjectControllerConfiguration) DeepCopy() *ProjectControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ProjectControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaConfiguration) DeepCopyInto(out *QuotaConfiguration) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	if in.ProjectSelector != nil {
		in, out := &in.ProjectSelector, &out.ProjectSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaConfiguration.
func (in *QuotaConfiguration) DeepCopy() *QuotaConfiguration {
	if in == nil {
		return nil
	}
	out := new(QuotaConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaControllerConfiguration) DeepCopyInto(out *QuotaControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaControllerConfiguration.
func (in *QuotaControllerConfiguration) DeepCopy() *QuotaControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(QuotaControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBindingControllerConfiguration) DeepCopyInto(out *SecretBindingControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBindingControllerConfiguration.
func (in *SecretBindingControllerConfiguration) DeepCopy() *SecretBindingControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SecretBindingControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBindingProviderControllerConfiguration) DeepCopyInto(out *SecretBindingProviderControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBindingProviderControllerConfiguration.
func (in *SecretBindingProviderControllerConfiguration) DeepCopy() *SecretBindingProviderControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SecretBindingProviderControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedControllerConfiguration) DeepCopyInto(out *SeedControllerConfiguration) {
	*out = *in
	if in.MonitorPeriod != nil {
		in, out := &in.MonitorPeriod, &out.MonitorPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.ShootMonitorPeriod != nil {
		in, out := &in.ShootMonitorPeriod, &out.ShootMonitorPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	out.SyncPeriod = in.SyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedControllerConfiguration.
func (in *SeedControllerConfiguration) DeepCopy() *SeedControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SeedControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	out.HTTP = in.HTTP
	out.HTTPS = in.HTTPS
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootConditionsControllerConfiguration) DeepCopyInto(out *ShootConditionsControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootConditionsControllerConfiguration.
func (in *ShootConditionsControllerConfiguration) DeepCopy() *ShootConditionsControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootConditionsControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootHibernationControllerConfiguration) DeepCopyInto(out *ShootHibernationControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootHibernationControllerConfiguration.
func (in *ShootHibernationControllerConfiguration) DeepCopy() *ShootHibernationControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootHibernationControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootMaintenanceControllerConfiguration) DeepCopyInto(out *ShootMaintenanceControllerConfiguration) {
	*out = *in
	if in.EnableShootControlPlaneRestarter != nil {
		in, out := &in.EnableShootControlPlaneRestarter, &out.EnableShootControlPlaneRestarter
		*out = new(bool)
		**out = **in
	}
	if in.EnableShootCoreAddonRestarter != nil {
		in, out := &in.EnableShootCoreAddonRestarter, &out.EnableShootCoreAddonRestarter
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootMaintenanceControllerConfiguration.
func (in *ShootMaintenanceControllerConfiguration) DeepCopy() *ShootMaintenanceControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootMaintenanceControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootQuotaControllerConfiguration) DeepCopyInto(out *ShootQuotaControllerConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootQuotaControllerConfiguration.
func (in *ShootQuotaControllerConfiguration) DeepCopy() *ShootQuotaControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootQuotaControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootReferenceControllerConfiguration) DeepCopyInto(out *ShootReferenceControllerConfiguration) {
	*out = *in
	if in.ProtectAuditPolicyConfigMaps != nil {
		in, out := &in.ProtectAuditPolicyConfigMaps, &out.ProtectAuditPolicyConfigMaps
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootReferenceControllerConfiguration.
func (in *ShootReferenceControllerConfiguration) DeepCopy() *ShootReferenceControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootReferenceControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootRetryControllerConfiguration) DeepCopyInto(out *ShootRetryControllerConfiguration) {
	*out = *in
	if in.RetryPeriod != nil {
		in, out := &in.RetryPeriod, &out.RetryPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootRetryControllerConfiguration.
func (in *ShootRetryControllerConfiguration) DeepCopy() *ShootRetryControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootRetryControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootStatusLabelControllerConfiguration) DeepCopyInto(out *ShootStatusLabelControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootStatusLabelControllerConfiguration.
func (in *ShootStatusLabelControllerConfiguration) DeepCopy() *ShootStatusLabelControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootStatusLabelControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSServer) DeepCopyInto(out *TLSServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSServer.
func (in *TLSServer) DeepCopy() *TLSServer {
	if in == nil {
		return nil
	}
	out := new(TLSServer)
	in.DeepCopyInto(out)
	return out
}
