// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computezonevmextensionpolicy


type ComputeZoneVmExtensionPolicyInstanceSelectors struct {
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_zone_vm_extension_policy#label_selector ComputeZoneVmExtensionPolicy#label_selector}
	LabelSelector *ComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
}

