// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computezonevmextensionpolicy


type ComputeZoneVmExtensionPolicyInstanceSelectorsLabelSelector struct {
	// A map of key-value pairs representing VM labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/compute_zone_vm_extension_policy#inclusion_labels ComputeZoneVmExtensionPolicy#inclusion_labels}
	InclusionLabels *map[string]*string `field:"optional" json:"inclusionLabels" yaml:"inclusionLabels"`
}

