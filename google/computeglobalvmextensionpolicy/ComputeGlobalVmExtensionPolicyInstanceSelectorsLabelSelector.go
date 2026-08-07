// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeglobalvmextensionpolicy


type ComputeGlobalVmExtensionPolicyInstanceSelectorsLabelSelector struct {
	// Labels as key value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/compute_global_vm_extension_policy#inclusion_labels ComputeGlobalVmExtensionPolicy#inclusion_labels}
	InclusionLabels *map[string]*string `field:"optional" json:"inclusionLabels" yaml:"inclusionLabels"`
}

