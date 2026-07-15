// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerInstanceLifecyclePolicyOnRepair struct {
	// Specifies whether the MIG can change a VM's zone during a repair.
	//
	// If "YES", MIG can select a different zone for the VM during a repair. Else if "NO", MIG cannot change a VM's zone during a repair. The default value of allow_changing_zone is "NO".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/compute_region_instance_group_manager#allow_changing_zone ComputeRegionInstanceGroupManager#allow_changing_zone}
	AllowChangingZone *string `field:"optional" json:"allowChangingZone" yaml:"allowChangingZone"`
}

