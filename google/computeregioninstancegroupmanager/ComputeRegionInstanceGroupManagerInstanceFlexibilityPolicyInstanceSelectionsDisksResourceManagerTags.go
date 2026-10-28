// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelectionsDisksResourceManagerTags struct {
	// The unique key of the resource manager tag to assign to disks. Keys must be in the format tagKeys/{tag_key_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/compute_region_instance_group_manager#key ComputeRegionInstanceGroupManager#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the resource manager tag to assign to disks. Values must be in the format tagValues/456.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/compute_region_instance_group_manager#value ComputeRegionInstanceGroupManager#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

