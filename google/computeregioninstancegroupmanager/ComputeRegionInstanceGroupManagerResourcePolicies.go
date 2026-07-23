// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerResourcePolicies struct {
	// The URL of the workload policy that is specified for this managed instance group.
	//
	// It can be a full or partial URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/compute_region_instance_group_manager#workload_policy ComputeRegionInstanceGroupManager#workload_policy}
	WorkloadPolicy *string `field:"optional" json:"workloadPolicy" yaml:"workloadPolicy"`
}

