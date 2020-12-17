// +build !ignore_autogenerated

// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package models

import (
	strfmt "github.com/go-openapi/strfmt"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BPFMapStatus) DeepCopyInto(out *BPFMapStatus) {
	*out = *in
	if in.Maps != nil {
		in, out := &in.Maps, &out.Maps
		*out = make([]*BPFMapProperties, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BPFMapProperties)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BPFMapStatus.
func (in *BPFMapStatus) DeepCopy() *BPFMapStatus {
	if in == nil {
		return nil
	}
	out := new(BPFMapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthManager) DeepCopyInto(out *BandwidthManager) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthManager.
func (in *BandwidthManager) DeepCopy() *BandwidthManager {
	if in == nil {
		return nil
	}
	out := new(BandwidthManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CIDRPolicy) DeepCopyInto(out *CIDRPolicy) {
	*out = *in
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]*PolicyRule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PolicyRule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]*PolicyRule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PolicyRule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIDRPolicy.
func (in *CIDRPolicy) DeepCopy() *CIDRPolicy {
	if in == nil {
		return nil
	}
	out := new(CIDRPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClockSource) DeepCopyInto(out *ClockSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClockSource.
func (in *ClockSource) DeepCopy() *ClockSource {
	if in == nil {
		return nil
	}
	out := new(ClockSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterMeshStatus) DeepCopyInto(out *ClusterMeshStatus) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]*RemoteCluster, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RemoteCluster)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterMeshStatus.
func (in *ClusterMeshStatus) DeepCopy() *ClusterMeshStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterMeshStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.CiliumHealth != nil {
		in, out := &in.CiliumHealth, &out.CiliumHealth
		*out = new(Status)
		**out = **in
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]*NodeElement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeElement)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerStatus) DeepCopyInto(out *ControllerStatus) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(ControllerStatusConfiguration)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ControllerStatusStatus)
		(*in).DeepCopyInto(*out)
	}
	in.UUID.DeepCopyInto(&out.UUID)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerStatus.
func (in *ControllerStatus) DeepCopy() *ControllerStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerStatusConfiguration) DeepCopyInto(out *ControllerStatusConfiguration) {
	*out = *in
	in.ErrorRetryBase.DeepCopyInto(&out.ErrorRetryBase)
	in.Interval.DeepCopyInto(&out.Interval)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerStatusConfiguration.
func (in *ControllerStatusConfiguration) DeepCopy() *ControllerStatusConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerStatusConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerStatusStatus) DeepCopyInto(out *ControllerStatusStatus) {
	*out = *in
	in.LastFailureTimestamp.DeepCopyInto(&out.LastFailureTimestamp)
	in.LastSuccessTimestamp.DeepCopyInto(&out.LastSuccessTimestamp)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerStatusStatus.
func (in *ControllerStatusStatus) DeepCopy() *ControllerStatusStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerStatusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointPolicy) DeepCopyInto(out *EndpointPolicy) {
	*out = *in
	if in.AllowedEgressIdentities != nil {
		in, out := &in.AllowedEgressIdentities, &out.AllowedEgressIdentities
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.AllowedIngressIdentities != nil {
		in, out := &in.AllowedIngressIdentities, &out.AllowedIngressIdentities
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.CidrPolicy != nil {
		in, out := &in.CidrPolicy, &out.CidrPolicy
		*out = new(CIDRPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.DeniedEgressIdentities != nil {
		in, out := &in.DeniedEgressIdentities, &out.DeniedEgressIdentities
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.DeniedIngressIdentities != nil {
		in, out := &in.DeniedIngressIdentities, &out.DeniedIngressIdentities
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.L4 != nil {
		in, out := &in.L4, &out.L4
		*out = new(L4Policy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointPolicy.
func (in *EndpointPolicy) DeepCopy() *EndpointPolicy {
	if in == nil {
		return nil
	}
	out := new(EndpointPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostRouting) DeepCopyInto(out *HostRouting) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostRouting.
func (in *HostRouting) DeepCopy() *HostRouting {
	if in == nil {
		return nil
	}
	out := new(HostRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubbleStatus) DeepCopyInto(out *HubbleStatus) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(HubbleStatusMetrics)
		**out = **in
	}
	if in.Observer != nil {
		in, out := &in.Observer, &out.Observer
		*out = new(HubbleStatusObserver)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubbleStatus.
func (in *HubbleStatus) DeepCopy() *HubbleStatus {
	if in == nil {
		return nil
	}
	out := new(HubbleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HubbleStatusObserver) DeepCopyInto(out *HubbleStatusObserver) {
	*out = *in
	in.Uptime.DeepCopyInto(&out.Uptime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HubbleStatusObserver.
func (in *HubbleStatusObserver) DeepCopy() *HubbleStatusObserver {
	if in == nil {
		return nil
	}
	out := new(HubbleStatusObserver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMStatus) DeepCopyInto(out *IPAMStatus) {
	*out = *in
	if in.Allocations != nil {
		in, out := &in.Allocations, &out.Allocations
		*out = make(AllocationMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IPV4 != nil {
		in, out := &in.IPV4, &out.IPV4
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IPV6 != nil {
		in, out := &in.IPV6, &out.IPV6
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMStatus.
func (in *IPAMStatus) DeepCopy() *IPAMStatus {
	if in == nil {
		return nil
	}
	out := new(IPAMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sStatus) DeepCopyInto(out *K8sStatus) {
	*out = *in
	if in.K8sAPIVersions != nil {
		in, out := &in.K8sAPIVersions, &out.K8sAPIVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sStatus.
func (in *K8sStatus) DeepCopy() *K8sStatus {
	if in == nil {
		return nil
	}
	out := new(K8sStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyReplacement) DeepCopyInto(out *KubeProxyReplacement) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]*KubeProxyReplacementDevicesItems0, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(KubeProxyReplacementDevicesItems0)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(KubeProxyReplacementFeatures)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyReplacement.
func (in *KubeProxyReplacement) DeepCopy() *KubeProxyReplacement {
	if in == nil {
		return nil
	}
	out := new(KubeProxyReplacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyReplacementDevicesItems0) DeepCopyInto(out *KubeProxyReplacementDevicesItems0) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyReplacementDevicesItems0.
func (in *KubeProxyReplacementDevicesItems0) DeepCopy() *KubeProxyReplacementDevicesItems0 {
	if in == nil {
		return nil
	}
	out := new(KubeProxyReplacementDevicesItems0)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyReplacementFeatures) DeepCopyInto(out *KubeProxyReplacementFeatures) {
	*out = *in
	if in.ExternalIPs != nil {
		in, out := &in.ExternalIPs, &out.ExternalIPs
		*out = new(KubeProxyReplacementFeaturesExternalIPs)
		**out = **in
	}
	if in.HostPort != nil {
		in, out := &in.HostPort, &out.HostPort
		*out = new(KubeProxyReplacementFeaturesHostPort)
		**out = **in
	}
	if in.HostReachableServices != nil {
		in, out := &in.HostReachableServices, &out.HostReachableServices
		*out = new(KubeProxyReplacementFeaturesHostReachableServices)
		(*in).DeepCopyInto(*out)
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(KubeProxyReplacementFeaturesNodePort)
		**out = **in
	}
	if in.SessionAffinity != nil {
		in, out := &in.SessionAffinity, &out.SessionAffinity
		*out = new(KubeProxyReplacementFeaturesSessionAffinity)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyReplacementFeatures.
func (in *KubeProxyReplacementFeatures) DeepCopy() *KubeProxyReplacementFeatures {
	if in == nil {
		return nil
	}
	out := new(KubeProxyReplacementFeatures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyReplacementFeaturesExternalIPs) DeepCopyInto(out *KubeProxyReplacementFeaturesExternalIPs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyReplacementFeaturesExternalIPs.
func (in *KubeProxyReplacementFeaturesExternalIPs) DeepCopy() *KubeProxyReplacementFeaturesExternalIPs {
	if in == nil {
		return nil
	}
	out := new(KubeProxyReplacementFeaturesExternalIPs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyReplacementFeaturesHostPort) DeepCopyInto(out *KubeProxyReplacementFeaturesHostPort) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyReplacementFeaturesHostPort.
func (in *KubeProxyReplacementFeaturesHostPort) DeepCopy() *KubeProxyReplacementFeaturesHostPort {
	if in == nil {
		return nil
	}
	out := new(KubeProxyReplacementFeaturesHostPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyReplacementFeaturesHostReachableServices) DeepCopyInto(out *KubeProxyReplacementFeaturesHostReachableServices) {
	*out = *in
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyReplacementFeaturesHostReachableServices.
func (in *KubeProxyReplacementFeaturesHostReachableServices) DeepCopy() *KubeProxyReplacementFeaturesHostReachableServices {
	if in == nil {
		return nil
	}
	out := new(KubeProxyReplacementFeaturesHostReachableServices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyReplacementFeaturesNodePort) DeepCopyInto(out *KubeProxyReplacementFeaturesNodePort) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyReplacementFeaturesNodePort.
func (in *KubeProxyReplacementFeaturesNodePort) DeepCopy() *KubeProxyReplacementFeaturesNodePort {
	if in == nil {
		return nil
	}
	out := new(KubeProxyReplacementFeaturesNodePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyReplacementFeaturesSessionAffinity) DeepCopyInto(out *KubeProxyReplacementFeaturesSessionAffinity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyReplacementFeaturesSessionAffinity.
func (in *KubeProxyReplacementFeaturesSessionAffinity) DeepCopy() *KubeProxyReplacementFeaturesSessionAffinity {
	if in == nil {
		return nil
	}
	out := new(KubeProxyReplacementFeaturesSessionAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4Policy) DeepCopyInto(out *L4Policy) {
	*out = *in
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]*PolicyRule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PolicyRule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]*PolicyRule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PolicyRule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L4Policy.
func (in *L4Policy) DeepCopy() *L4Policy {
	if in == nil {
		return nil
	}
	out := new(L4Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Masquerading) DeepCopyInto(out *Masquerading) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(MasqueradingEnabled)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Masquerading.
func (in *Masquerading) DeepCopy() *Masquerading {
	if in == nil {
		return nil
	}
	out := new(Masquerading)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAddressing) DeepCopyInto(out *NodeAddressing) {
	*out = *in
	if in.IPV4 != nil {
		in, out := &in.IPV4, &out.IPV4
		*out = new(NodeAddressingElement)
		**out = **in
	}
	if in.IPV6 != nil {
		in, out := &in.IPV6, &out.IPV6
		*out = new(NodeAddressingElement)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAddressing.
func (in *NodeAddressing) DeepCopy() *NodeAddressing {
	if in == nil {
		return nil
	}
	out := new(NodeAddressing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeElement) DeepCopyInto(out *NodeElement) {
	*out = *in
	if in.HealthEndpointAddress != nil {
		in, out := &in.HealthEndpointAddress, &out.HealthEndpointAddress
		*out = new(NodeAddressing)
		(*in).DeepCopyInto(*out)
	}
	if in.PrimaryAddress != nil {
		in, out := &in.PrimaryAddress, &out.PrimaryAddress
		*out = new(NodeAddressing)
		(*in).DeepCopyInto(*out)
	}
	if in.SecondaryAddresses != nil {
		in, out := &in.SecondaryAddresses, &out.SecondaryAddresses
		*out = make([]*NodeAddressingElement, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(NodeAddressingElement)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeElement.
func (in *NodeElement) DeepCopy() *NodeElement {
	if in == nil {
		return nil
	}
	out := new(NodeElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRule) DeepCopyInto(out *PolicyRule) {
	*out = *in
	if in.DerivedFromRules != nil {
		in, out := &in.DerivedFromRules, &out.DerivedFromRules
		*out = make([][]string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRule.
func (in *PolicyRule) DeepCopy() *PolicyRule {
	if in == nil {
		return nil
	}
	out := new(PolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyStatistics) DeepCopyInto(out *ProxyStatistics) {
	*out = *in
	if in.Statistics != nil {
		in, out := &in.Statistics, &out.Statistics
		*out = new(RequestResponseStatistics)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyStatistics.
func (in *ProxyStatistics) DeepCopy() *ProxyStatistics {
	if in == nil {
		return nil
	}
	out := new(ProxyStatistics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyStatus) DeepCopyInto(out *ProxyStatus) {
	*out = *in
	if in.Redirects != nil {
		in, out := &in.Redirects, &out.Redirects
		*out = make([]*ProxyRedirect, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ProxyRedirect)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyStatus.
func (in *ProxyStatus) DeepCopy() *ProxyStatus {
	if in == nil {
		return nil
	}
	out := new(ProxyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteCluster) DeepCopyInto(out *RemoteCluster) {
	*out = *in
	in.LastFailure.DeepCopyInto(&out.LastFailure)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteCluster.
func (in *RemoteCluster) DeepCopy() *RemoteCluster {
	if in == nil {
		return nil
	}
	out := new(RemoteCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequestResponseStatistics) DeepCopyInto(out *RequestResponseStatistics) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(MessageForwardingStatistics)
		**out = **in
	}
	if in.Responses != nil {
		in, out := &in.Responses, &out.Responses
		*out = new(MessageForwardingStatistics)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequestResponseStatistics.
func (in *RequestResponseStatistics) DeepCopy() *RequestResponseStatistics {
	if in == nil {
		return nil
	}
	out := new(RequestResponseStatistics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusResponse) DeepCopyInto(out *StatusResponse) {
	*out = *in
	if in.BandwidthManager != nil {
		in, out := &in.BandwidthManager, &out.BandwidthManager
		*out = new(BandwidthManager)
		(*in).DeepCopyInto(*out)
	}
	if in.BpfMaps != nil {
		in, out := &in.BpfMaps, &out.BpfMaps
		*out = new(BPFMapStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Cilium != nil {
		in, out := &in.Cilium, &out.Cilium
		*out = new(Status)
		**out = **in
	}
	if in.ClockSource != nil {
		in, out := &in.ClockSource, &out.ClockSource
		*out = new(ClockSource)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(ClusterStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterMesh != nil {
		in, out := &in.ClusterMesh, &out.ClusterMesh
		*out = new(ClusterMeshStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerRuntime != nil {
		in, out := &in.ContainerRuntime, &out.ContainerRuntime
		*out = new(Status)
		**out = **in
	}
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make(ControllerStatuses, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ControllerStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.HostRouting != nil {
		in, out := &in.HostRouting, &out.HostRouting
		*out = new(HostRouting)
		**out = **in
	}
	if in.Hubble != nil {
		in, out := &in.Hubble, &out.Hubble
		*out = new(HubbleStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Ipam != nil {
		in, out := &in.Ipam, &out.Ipam
		*out = new(IPAMStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeProxyReplacement != nil {
		in, out := &in.KubeProxyReplacement, &out.KubeProxyReplacement
		*out = new(KubeProxyReplacement)
		(*in).DeepCopyInto(*out)
	}
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(K8sStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Kvstore != nil {
		in, out := &in.Kvstore, &out.Kvstore
		*out = new(Status)
		**out = **in
	}
	if in.Masquerading != nil {
		in, out := &in.Masquerading, &out.Masquerading
		*out = new(Masquerading)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeMonitor != nil {
		in, out := &in.NodeMonitor, &out.NodeMonitor
		*out = new(MonitorStatus)
		**out = **in
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(ProxyStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Stale != nil {
		in, out := &in.Stale, &out.Stale
		*out = make(map[string]strfmt.DateTime, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusResponse.
func (in *StatusResponse) DeepCopy() *StatusResponse {
	if in == nil {
		return nil
	}
	out := new(StatusResponse)
	in.DeepCopyInto(out)
	return out
}
