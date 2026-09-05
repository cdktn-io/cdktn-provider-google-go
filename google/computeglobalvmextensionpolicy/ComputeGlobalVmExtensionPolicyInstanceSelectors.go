// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeglobalvmextensionpolicy


type ComputeGlobalVmExtensionPolicyInstanceSelectors struct {
	// label_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_global_vm_extension_policy#label_selector ComputeGlobalVmExtensionPolicy#label_selector}
	LabelSelector *ComputeGlobalVmExtensionPolicyInstanceSelectorsLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
}

