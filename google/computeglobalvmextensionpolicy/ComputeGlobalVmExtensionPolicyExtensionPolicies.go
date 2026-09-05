// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computeglobalvmextensionpolicy


type ComputeGlobalVmExtensionPolicyExtensionPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_global_vm_extension_policy#extension_name ComputeGlobalVmExtensionPolicy#extension_name}.
	ExtensionName *string `field:"required" json:"extensionName" yaml:"extensionName"`
	// The version pinning for the extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_global_vm_extension_policy#pinned_version ComputeGlobalVmExtensionPolicy#pinned_version}
	PinnedVersion *string `field:"optional" json:"pinnedVersion" yaml:"pinnedVersion"`
	// String configuration payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/compute_global_vm_extension_policy#string_config ComputeGlobalVmExtensionPolicy#string_config}
	StringConfig *string `field:"optional" json:"stringConfig" yaml:"stringConfig"`
}

