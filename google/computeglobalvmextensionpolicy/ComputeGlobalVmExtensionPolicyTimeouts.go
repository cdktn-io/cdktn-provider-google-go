// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeglobalvmextensionpolicy


type ComputeGlobalVmExtensionPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_global_vm_extension_policy#create ComputeGlobalVmExtensionPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_global_vm_extension_policy#delete ComputeGlobalVmExtensionPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/compute_global_vm_extension_policy#update ComputeGlobalVmExtensionPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

